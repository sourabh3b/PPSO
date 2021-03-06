package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"fmt"
	"math/rand"
	"time"
	"github.com/PPSO/plot"
)


const (
	MaxIteration   = 100
	Dimension      = 2
	PopulationSize = 50
	StartRange     = -5.0
	EndRange       = 5.0
	W              = 0.5
)



type Particle struct {
	Pos       []float64
	Vel       []float64
	Fitness   float64
	Pbest     []float64
	Lbest     []float64
	Neighbors []*Particle
}



func randomPos() float64 {
	var randValue float64
	rand.Seed(time.Now().UnixNano())
	randValue = StartRange + (EndRange-StartRange)*rand.Float64()
	return randValue
}


func (p *Particle) Initialize(dim, population int) {
	p.Pos = make([]float64, dim)
	p.Vel = make([]float64, dim)
	p.Pbest = make([]float64, dim)
	p.Lbest = make([]float64, dim)

	for i, _ := range p.Pos {
		tmp := randomPos()
		p.Pos[i] = tmp
	}
	copy(p.Pbest, p.Pos)
}


//This is a basic x^2
func evaluate(pos []float64) float64 {
	var result float64 = 0.0
	for _, x := range pos {
		result += x * x
	}
	return result
}


func advance(p Particle) ([]float64, []float64) {
	dim := len(p.Pos)
	rho1 := 0.14 * rand.Float64()
	rho2 := 0.14 * rand.Float64()
	newPos := make([]float64, dim)
	newVel := make([]float64, dim)

	for i := 0; i < dim; i++ {
		newPos[i] = p.Pos[i] + p.Vel[i]
		newVel[i] = W * p.Vel[i]
		newVel[i] += rho1 * (p.Pbest[i] - p.Pos[i])
		newVel[i] += rho2 * (p.Lbest[i] - p.Pos[i])
	}

	return newPos, newVel
}


//TestRoute - test route
func TestRoute(w http.ResponseWriter, r *http.Request) {
	//render := render.New()
	fmt.Fprint(w, "Hello World !")
	//render.JSON(w, http.StatusOK, nil)
	return
}


func main() {
	log.Info("Starting PPSO ....")
	//http.HandleFunc("/test", TestRoute)
	//http.ListenAndServe(":8889", nil)
	//todo : create API to calculate PSO fitness function value


	var swarm []Particle
	var bestParticle *Particle

	//Initialization
	swarm = make([]Particle, PopulationSize)

	//filling swarm
	for i := range swarm {
		p := Particle{}
		p.Initialize(Dimension, PopulationSize)
		p.Fitness = evaluate(p.Pos)
		swarm[i] = p
	}


	// Adding  Neighbors
	for i := range swarm {
		for j := range swarm {
			if i != j {
				swarm[i].Neighbors = append(swarm[i].Neighbors, &swarm[j])
			}
		}
	}

	// Pick up the particle which has the best fitness
	bestParticle = &swarm[0]
	for i := range swarm {
		if swarm[i].Fitness < bestParticle.Fitness {
			bestParticle = &swarm[i]
		}
	}
	for i := range swarm {
		for j := range swarm[i].Lbest {
			swarm[i].Lbest[j] = bestParticle.Lbest[j]
		}
	}


	var plotXPoints []float64
	var plotYPoints []float64
	//while a termination criterion is not met:
	for n := 0; n < MaxIteration; n++ {
		// Update the particle's velocity:
		for i, p := range swarm {
			swarm[i].Pos, swarm[i].Vel = advance(p)
		}

		// Update Personal Best
		for i, p := range swarm {
			fitness := evaluate(p.Pos)
			swarm[i].Fitness = fitness
			pbestFitness := evaluate(p.Pbest)
			if fitness < pbestFitness {
				for j := range swarm[i].Pos {
					swarm[i].Pbest[j] = swarm[i].Pos[j]
				}
			}
		}

		// Update Local Best
		bestParticle = &swarm[0]
		for i := range swarm {
			if swarm[i].Fitness < bestParticle.Fitness {
				bestParticle = &swarm[i]
			}
		}
		for i := range swarm {
			for j := range swarm[i].Lbest {
				swarm[i].Lbest[j] = bestParticle.Lbest[j]
			}
		}

		// Output
		//log.Info("Fitness Value  : " , n , "  - ", bestParticle.Fitness)
		//log.Info(n)
		//fmt.Println(bestParticle.Fitness)

		//taking X points with Y points
		plotXPoints = append(plotXPoints,float64(n))
		plotYPoints = append(plotYPoints,bestParticle.Fitness)

	}

	//fmt.Println("plotYPoints = ", plotYPoints)
	//fmt.Println("plotXPoints = ", plotXPoints)
	//plot graph
	plot.PlotGraph(plotXPoints,plotYPoints)

	log.Info("Graph Plot Done")
}


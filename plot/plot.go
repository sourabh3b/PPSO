package plot

import (
	"github.com/Arafatk/glot"
	"github.com/PPSO/util"
	"github.com/PPSO/constants"
	"github.com/satori/go.uuid"
)

//plotGraph - plots the 2D Graph for the points
func PlotGraph(xPoints []float64, yPoints []float64) {
	dimensions := 2
	// The dimensions supported by the plot
	persist := false
	debug := false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	pointGroupName := constants.FitnessValue
	style := "lines"
	points := [][]float64{xPoints, yPoints}
	// Adding a point group
	plot.AddPointGroup(pointGroupName, style, points)
	// A plot type used to make points/ curves and customize and save them as an image.
	plot.SetTitle(constants.PSOPlotTitle)
	// Optional: Setting the title of the plot
	plot.SetXLabel(constants.XAxisLabel)
	plot.SetYLabel(constants.YAxisLabel)
	// Optional: Setting label for X and Y axis


	_,maxX := util.MinMax(xPoints)
	_,maxY := util.MinMax(yPoints)

	plot.SetXrange(0, int(maxX))
	plot.SetYrange(0, int(maxY) * 2)

	uuidN,_ := uuid.NewV1()
	destinationFile := constants.GraphFolderName + constants.Slash + constants.PPSO + uuidN.String() + constants.DotPNG
	plot.SavePlot(destinationFile)
}


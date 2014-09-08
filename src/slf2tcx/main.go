package main

import (
	"fmt"
	"gpx"
	"log"
	"os"
	"slf"
	"tcx"
)

func main() {
	var workoutPath string
	var trackPath string
	var ansPath string
	switch len(os.Args) {
	case 3:
		workoutPath = os.Args[1]
		ansPath = os.Args[2]
	case 4:
		workoutPath = os.Args[1]
		trackPath = os.Args[2]
		ansPath = os.Args[3]
	default:
		fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], " input.slf [replace_trk.gpx] output.tcx\n")
		os.Exit(1)
	}
	var ans *tcx.TrainingCenterDatabase = new(tcx.TrainingCenterDatabase)
	var track *gpx.Gpx
	var workout *slf.Log
	var err error
	// Load SLF workout
	if workout, err = slf.Load(workoutPath); err != nil {
		log.Fatal(err)
	}
	ans.FromLog(workout)
	// Load GPX track
	if trackPath != "" {
		if track, err = gpx.Load(trackPath); err != nil {
			log.Fatal(err)
		}
		ans.ReplaceTrack(track)
	}
	// Save TCX
	if err = ans.Save(ansPath); err != nil {
		log.Fatal(err)
	}
}
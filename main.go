package main

import (
	"fmt"
	"os"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	buildDate string
)

func main() {
	fmt.Printf("Drone Email Plugin built at %s\n", buildDate)

	system := drone.System{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("system", &system)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.Recipients) == 0 {
		vargs.Recipients = []string{
			build.Email,
		}
	}

	if vargs.Port == 0 {
		vargs.Port = 587
	}

	err := Send(&Context{
		System: system,
		Repo:   repo,
		Build:  build,
		Vargs:  vargs,
	})

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
		return
	}
}

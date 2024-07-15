package main

import (
	"os"

	phpredishandler "github.com/initializ-buildpacks/php-redis-session-handler"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	"github.com/paketo-buildpacks/packit/v2/servicebindings"
)

func main() {
	logEmitter := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))
	serviceResolver := servicebindings.NewResolver()

	packit.Run(
		phpredishandler.Detect(
			serviceResolver,
		),
		phpredishandler.Build(
			phpredishandler.NewRedisConfigParser(),
			serviceResolver,
			phpredishandler.NewRedisConfigWriter(logEmitter),
			logEmitter,
		),
	)
}

package bootstrap

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(fx.Provide(
	NewLogger,
	NewApplication,
	NewHttpServer,
))
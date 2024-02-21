//go:build !notracing

package fraudserv

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("fraudserv")

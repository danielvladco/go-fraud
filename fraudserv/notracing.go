//go:build notracing

package fraudserv

import "go.opentelemetry.io/otel/trace/noop"

var tracer = noop.NewTracerProvider().Tracer("fraudserv")

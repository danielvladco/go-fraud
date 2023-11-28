//go:build tracing

package attribute

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func String(k string, v string) trace.SpanStartEventOption {
	return attribute.String(k, v)
}

func Int(k string, v int) trace.SpanStartEventOption {
	return attribute.Int(k, v)
}

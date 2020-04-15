// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package converter

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector/consumer"
	"github.com/open-telemetry/opentelemetry-collector/consumer/consumerdata"
	"github.com/open-telemetry/opentelemetry-collector/consumer/pdata"
	"github.com/open-telemetry/opentelemetry-collector/internal/data"
	"github.com/open-telemetry/opentelemetry-collector/translator/internaldata"
)

// NewInternalToOCTraceConverter creates new internalToOCTraceConverter that takes TraceConsumer and
// implements ConsumeTrace interface.
func NewInternalToOCTraceConverter(tc consumer.TraceConsumerOld) consumer.TraceConsumer {
	return &internalToOCTraceConverter{tc}
}

// internalToOCTraceConverter is a internal to oc translation shim that takes TraceConsumer and
// implements ConsumeTrace interface.
type internalToOCTraceConverter struct {
	traceConsumer consumer.TraceConsumerOld
}

// ConsumeTrace takes new-style data.TraceData method, converts it to OC and uses old-style ConsumeTraceData method
// to process the trace data.
func (tc *internalToOCTraceConverter) ConsumeTrace(ctx context.Context, td pdata.TraceData) error {
	ocTraces := internaldata.TraceDataToOC(td)
	for i := range ocTraces {
		err := tc.traceConsumer.ConsumeTraceData(ctx, ocTraces[i])
		if err != nil {
			return err
		}
	}
	return nil
}

var _ consumer.TraceConsumer = (*internalToOCTraceConverter)(nil)

// NewInternalToOCMetricsConverter creates new internalToOCMetricsConverter that takes MetricsConsumer and
// implements ConsumeTrace interface.
func NewInternalToOCMetricsConverter(tc consumer.MetricsConsumerOld) consumer.MetricsConsumer {
	return &internalToOCMetricsConverter{tc}
}

// internalToOCMetricsConverter is a internal to oc translation shim that takes MetricsConsumer and
// implements ConsumeMetrics interface.
type internalToOCMetricsConverter struct {
	metricsConsumer consumer.MetricsConsumerOld
}

// ConsumeMetrics takes new-style data.MetricData method, converts it to OC and uses old-style ConsumeMetricsData method
// to process the metrics data.
func (tc *internalToOCMetricsConverter) ConsumeMetrics(ctx context.Context, td data.MetricData) error {
	ocMetrics := internaldata.MetricDataToOC(td)
	for i := range ocMetrics {
		err := tc.metricsConsumer.ConsumeMetricsData(ctx, ocMetrics[i])
		if err != nil {
			return err
		}
	}
	return nil
}

var _ consumer.MetricsConsumer = (*internalToOCMetricsConverter)(nil)

// NewOCToInternalTraceConverter creates new ocToInternalTraceConverter that takes TraceConsumer and
// implements ConsumeTrace interface.
func NewOCToInternalTraceConverter(tc consumer.TraceConsumer) consumer.TraceConsumerOld {
	return &ocToInternalTraceConverter{tc}
}

// ocToInternalTraceConverter is a internal to oc translation shim that takes TraceConsumer and
// implements ConsumeTrace interface.
type ocToInternalTraceConverter struct {
	traceConsumer consumer.TraceConsumer
}

// ConsumeTrace takes new-style data.TraceData method, converts it to OC and uses old-style ConsumeTraceData method
// to process the trace data.
func (tc *ocToInternalTraceConverter) ConsumeTraceData(ctx context.Context, td consumerdata.TraceData) error {
	traceData := internaldata.OCToTraceData(td)
	err := tc.traceConsumer.ConsumeTrace(ctx, traceData)
	if err != nil {
		return err
	}

	return nil
}

var _ consumer.TraceConsumerOld = (*ocToInternalTraceConverter)(nil)

// NewOCToInternalMetricsConverter creates new ocToInternalMetricsConverter that takes MetricsConsumer and
// implements ConsumeTrace interface.
func NewOCToInternalMetricsConverter(tc consumer.MetricsConsumer) consumer.MetricsConsumerOld {
	return &ocToInternalMetricsConverter{tc}
}

// ocToInternalMetricsConverter is a internal to oc translation shim that takes MetricsConsumer and
// implements ConsumeMetrics interface.
type ocToInternalMetricsConverter struct {
	metricsConsumer consumer.MetricsConsumer
}

// ConsumeMetrics takes new-style data.MetricData method, converts it to OC and uses old-style ConsumeMetricsData method
// to process the metrics data.
func (tc *ocToInternalMetricsConverter) ConsumeMetricsData(ctx context.Context, td consumerdata.MetricsData) error {
	metricsData := internaldata.OCToMetricData(td)
	err := tc.metricsConsumer.ConsumeMetrics(ctx, metricsData)
	if err != nil {
		return err
	}
	return nil
}

var _ consumer.MetricsConsumerOld = (*ocToInternalMetricsConverter)(nil)
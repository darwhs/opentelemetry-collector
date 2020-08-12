package testbed

// MockDataConsumer is an interface that
//keeps the count of number of events received by mock receiver.
//This is mainly useful for the Exporters that are not have the matching receiver
type MockTraceDataConsumer interface {
	// MockConsumeTraceData receives metrics and traces and just count the number of event received.
	MockConsumeTraceData(spansCount int) error
}

type MockMetricDataConsumer interface {
	// MockConsumeMetricData receives metrics and traces and just count the number of event received.
	MockConsumeMetricData(metricsCount int) error
}

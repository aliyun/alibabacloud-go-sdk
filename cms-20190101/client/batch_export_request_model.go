// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchExportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCursor(v string) *BatchExportRequest
	GetCursor() *string
	SetLength(v int32) *BatchExportRequest
	GetLength() *int32
	SetMeasurements(v []*string) *BatchExportRequest
	GetMeasurements() []*string
	SetMetric(v string) *BatchExportRequest
	GetMetric() *string
	SetNamespace(v string) *BatchExportRequest
	GetNamespace() *string
}

type BatchExportRequest struct {
	// When you call this operation in a loop to export data, you must specify the value of `Cursor`. You can obtain the value of `Cursor` by using the following methods:
	//
	// - When you call this operation for the first time, you must first call the Cursor operation to obtain the value of `Cursor`. For more information, see [Cursor](https://help.aliyun.com/document_detail/2330730.html).
	//
	// - When you call this operation again, you can obtain the value of `Cursor` from the response of the last call.
	//
	// This parameter is required.
	//
	// example:
	//
	// eyJidWNrZXRzIjo0LCJjdXJzb3IiOiIxNjQxNDU0MzIwMDAwMWUxY2YxNWY0NTU0MTliZjllYTY4OWQ2ODI1OTU1Yzc1NmZjMDQ2OTMxMzczMzM2MzUzMTMxMzEzMzM0MzMzODM5MzEzMTMwMjQyYzY5MmQ3NTY2MzYzMjY3NmI2ZjM5MzU2YjY4MzAzMTYyNzg3MTcwNjkzMTM3MjQyYyIsImN1cnNvclZlcnNpb24iOiJxdWVyeSIsImVuZFRpbWUiOjE2NDE0NTQ3OTU4MjMsImV4cG9ydEVuZFRpbWUiOjE2NDE0NTQ3OTU4MjMsImV4cG9ydFN0YXJ0VGltZSI6MTY0MTQ1NDE5NTgyMywiZXhwcmVzc1JhbmdlIjpmYWxzZSwiaGFzTmV4dCI6dHJ1ZSwiaW5wdXRNZXRyaWMiOiJDUFVVdGlsaXphdGlvbiIsImlucHV0TmFtZXNwYWNlIjoiYWNzX2Vjc19kYXNoYm9hcmQiLCJsaW1pdCI6MTAwMCwibG9nVGltZU1vZGUiOnRydWUsIm1hdGNoZXJzIjp7ImNoYWluIjpbeyJsYWJlbCI6InVzZXJJZCIsIm9wZXJhdG9yIjoiRVFVQUxTIiwidmFsdWUiOiIxNzM2NTExMTM0Mzg5MTEwIn1dfSwibWV0cmljIjoiQ1BVVXRpbGl6YXRpb24iLCJtZXRyaWNUeXBlIjoiTUVUUklDIiwibmFtZXNwYWNlIjoiYWNzX2Vjc19kYXNoYm9hcmQiLCJuZXh0UGtBZGFwdGVyIjp7fSwib2Zmc2V0IjowLCJwYXJlbnRVaWQiOjEyNzA2NzY2Nzk1NDY3MDQsInN0YXJ0VGltZSI6MTY0MTQ1NDE5NTgyMywic3RlcCI6LTEsInRpbWVvdXQiOjEyMCwid2luZG93Ijo2****
	Cursor *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	// The maximum number of data entries to return each time.
	//
	// Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// The measurements that are used to customize the returned data. By default, all measurements are returned.
	//
	// For example, the metric `cpu_idle` of the cloud service `acs_ecs_dashboard` has three measurement columns: `Average`, `Maximum`, and `Minimum`. If you only need to return the `Average` and `Maximum` columns, set this parameter to the array `["Average", "Maximum"]`.
	//
	// For information about how to obtain the measurements of a metric of a cloud service, see the `statistics` column of [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	Measurements []*string `json:"Measurements,omitempty" xml:"Measurements,omitempty" type:"Repeated"`
	// The name of the metric of the cloud service.
	//
	// For information about how to obtain the name of a metric of a cloud service, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > This parameter must be the same as the request parameter `Metric` in the Cursor operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_idle
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The data namespace of the cloud service.
	//
	// For information about how to obtain the data namespace of a cloud service, see [Appendix 1: Metrics](https://help.aliyun.com/document_detail/163515.html).
	//
	// > This parameter must be the same as the request parameter `Namespace` in the Cursor operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchExportRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchExportRequest) GoString() string {
	return s.String()
}

func (s *BatchExportRequest) GetCursor() *string {
	return s.Cursor
}

func (s *BatchExportRequest) GetLength() *int32 {
	return s.Length
}

func (s *BatchExportRequest) GetMeasurements() []*string {
	return s.Measurements
}

func (s *BatchExportRequest) GetMetric() *string {
	return s.Metric
}

func (s *BatchExportRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchExportRequest) SetCursor(v string) *BatchExportRequest {
	s.Cursor = &v
	return s
}

func (s *BatchExportRequest) SetLength(v int32) *BatchExportRequest {
	s.Length = &v
	return s
}

func (s *BatchExportRequest) SetMeasurements(v []*string) *BatchExportRequest {
	s.Measurements = v
	return s
}

func (s *BatchExportRequest) SetMetric(v string) *BatchExportRequest {
	s.Metric = &v
	return s
}

func (s *BatchExportRequest) SetNamespace(v string) *BatchExportRequest {
	s.Namespace = &v
	return s
}

func (s *BatchExportRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBenchmarkTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListBenchmarkTaskRequest
	GetFilter() *string
	SetModelId(v string) *ListBenchmarkTaskRequest
	GetModelId() *string
	SetOrder(v string) *ListBenchmarkTaskRequest
	GetOrder() *string
	SetPageNumber(v string) *ListBenchmarkTaskRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListBenchmarkTaskRequest
	GetPageSize() *string
	SetRequestMethod(v string) *ListBenchmarkTaskRequest
	GetRequestMethod() *string
	SetServiceName(v string) *ListBenchmarkTaskRequest
	GetServiceName() *string
	SetSort(v string) *ListBenchmarkTaskRequest
	GetSort() *string
	SetStatus(v string) *ListBenchmarkTaskRequest
	GetStatus() *string
}

type ListBenchmarkTaskRequest struct {
	// The filter parameter. If specified, the system searches for records that match the specified value in the EAS service name and stress testing task name.
	//
	// example:
	//
	// test_bench
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The model ID.
	//
	// example:
	//
	// Qwen3.5-27B
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The sort order. Valid values:
	//
	// - desc: Descending order. This is the default value.
	//
	// - asc: Ascending order.
	//
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// The page number of the stress testing task list. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of stress testing tasks to return on each page. Default value: 100.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request method for stress testing. Valid values:
	//
	// - concurrency: Fixed concurrency mode.
	//
	// - requestRate: Fixed request rate mode.
	//
	// - peakThroughput: Peak throughput mode.
	//
	// example:
	//
	// requestRate
	RequestMethod *string `json:"RequestMethod,omitempty" xml:"RequestMethod,omitempty"`
	// The name of the EAS service that corresponds to the stress testing task. For more information about how to query the service name, see [ListServices](https://help.aliyun.com/document_detail/412109.html).
	//
	// example:
	//
	// test_bench_srv
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The sort field.
	//
	// example:
	//
	// CreateTime
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	// The execution status used to filter stress testing tasks.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListBenchmarkTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *ListBenchmarkTaskRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListBenchmarkTaskRequest) GetModelId() *string {
	return s.ModelId
}

func (s *ListBenchmarkTaskRequest) GetOrder() *string {
	return s.Order
}

func (s *ListBenchmarkTaskRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListBenchmarkTaskRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListBenchmarkTaskRequest) GetRequestMethod() *string {
	return s.RequestMethod
}

func (s *ListBenchmarkTaskRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListBenchmarkTaskRequest) GetSort() *string {
	return s.Sort
}

func (s *ListBenchmarkTaskRequest) GetStatus() *string {
	return s.Status
}

func (s *ListBenchmarkTaskRequest) SetFilter(v string) *ListBenchmarkTaskRequest {
	s.Filter = &v
	return s
}

func (s *ListBenchmarkTaskRequest) SetModelId(v string) *ListBenchmarkTaskRequest {
	s.ModelId = &v
	return s
}

func (s *ListBenchmarkTaskRequest) SetOrder(v string) *ListBenchmarkTaskRequest {
	s.Order = &v
	return s
}

func (s *ListBenchmarkTaskRequest) SetPageNumber(v string) *ListBenchmarkTaskRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBenchmarkTaskRequest) SetPageSize(v string) *ListBenchmarkTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListBenchmarkTaskRequest) SetRequestMethod(v string) *ListBenchmarkTaskRequest {
	s.RequestMethod = &v
	return s
}

func (s *ListBenchmarkTaskRequest) SetServiceName(v string) *ListBenchmarkTaskRequest {
	s.ServiceName = &v
	return s
}

func (s *ListBenchmarkTaskRequest) SetSort(v string) *ListBenchmarkTaskRequest {
	s.Sort = &v
	return s
}

func (s *ListBenchmarkTaskRequest) SetStatus(v string) *ListBenchmarkTaskRequest {
	s.Status = &v
	return s
}

func (s *ListBenchmarkTaskRequest) Validate() error {
	return dara.Validate(s)
}

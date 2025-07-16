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
	SetPageNumber(v string) *ListBenchmarkTaskRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListBenchmarkTaskRequest
	GetPageSize() *string
	SetServiceName(v string) *ListBenchmarkTaskRequest
	GetServiceName() *string
}

type ListBenchmarkTaskRequest struct {
	// The keyword used to query required stress testing tasks. If this parameter is specified, the system returns stress testing tasks based on the names of the stress testing tasks in the matched Elastic Algorithm Service (EAS).
	//
	// example:
	//
	// test_bench
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the EAS service that corresponds to the stress testing task. For more information about how to query the service name, see [ListServices](https://help.aliyun.com/document_detail/412109.html).
	//
	// example:
	//
	// test_bench_srv
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
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

func (s *ListBenchmarkTaskRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListBenchmarkTaskRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListBenchmarkTaskRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListBenchmarkTaskRequest) SetFilter(v string) *ListBenchmarkTaskRequest {
	s.Filter = &v
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

func (s *ListBenchmarkTaskRequest) SetServiceName(v string) *ListBenchmarkTaskRequest {
	s.ServiceName = &v
	return s
}

func (s *ListBenchmarkTaskRequest) Validate() error {
	return dara.Validate(s)
}

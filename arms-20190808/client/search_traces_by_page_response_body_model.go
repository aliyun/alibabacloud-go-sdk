// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchTracesByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *SearchTracesByPageResponseBodyPageBean) *SearchTracesByPageResponseBody
	GetPageBean() *SearchTracesByPageResponseBodyPageBean
	SetRequestId(v string) *SearchTracesByPageResponseBody
	GetRequestId() *string
}

type SearchTracesByPageResponseBody struct {
	// The struct returned.
	PageBean *SearchTracesByPageResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4C518054-852F-4023-ABC1-4AF95FF7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchTracesByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageResponseBody) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBody) GetPageBean() *SearchTracesByPageResponseBodyPageBean {
	return s.PageBean
}

func (s *SearchTracesByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchTracesByPageResponseBody) SetPageBean(v *SearchTracesByPageResponseBodyPageBean) *SearchTracesByPageResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchTracesByPageResponseBody) SetRequestId(v string) *SearchTracesByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchTracesByPageResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchTracesByPageResponseBodyPageBean struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1601
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The details of the returned traces.
	TraceInfos []*SearchTracesByPageResponseBodyPageBeanTraceInfos `json:"TraceInfos,omitempty" xml:"TraceInfos,omitempty" type:"Repeated"`
}

func (s SearchTracesByPageResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBodyPageBean) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchTracesByPageResponseBodyPageBean) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchTracesByPageResponseBodyPageBean) GetTotal() *int32 {
	return s.Total
}

func (s *SearchTracesByPageResponseBodyPageBean) GetTraceInfos() []*SearchTracesByPageResponseBodyPageBeanTraceInfos {
	return s.TraceInfos
}

func (s *SearchTracesByPageResponseBodyPageBean) SetPageNumber(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetPageSize(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetTotal(v int32) *SearchTracesByPageResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) SetTraceInfos(v []*SearchTracesByPageResponseBodyPageBeanTraceInfos) *SearchTracesByPageResponseBodyPageBean {
	s.TraceInfos = v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBean) Validate() error {
	return dara.Validate(s)
}

type SearchTracesByPageResponseBodyPageBeanTraceInfos struct {
	// The amount of time consumed by the trace. Unit: milliseconds.
	//
	// example:
	//
	// 679
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The name of the traced span.
	//
	// example:
	//
	// /demo/queryException/12
	OperationName *string `json:"OperationName,omitempty" xml:"OperationName,omitempty"`
	// The IP address of the host where the application resides.
	//
	// example:
	//
	// 172.20.XX.XX
	ServiceIp *string `json:"ServiceIp,omitempty" xml:"ServiceIp,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// arms-k8s-demo-subcomponent
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The span ID of the trace.
	//
	// example:
	//
	// be3d6dcf5750e***
	SpanID *string `json:"SpanID,omitempty" xml:"SpanID,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 1595174436994
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The trace ID.
	//
	// example:
	//
	// ac1400a115951744369947025d****
	TraceID *string `json:"TraceID,omitempty" xml:"TraceID,omitempty"`
}

func (s SearchTracesByPageResponseBodyPageBeanTraceInfos) String() string {
	return dara.Prettify(s)
}

func (s SearchTracesByPageResponseBodyPageBeanTraceInfos) GoString() string {
	return s.String()
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetDuration() *int64 {
	return s.Duration
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetOperationName() *string {
	return s.OperationName
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetServiceIp() *string {
	return s.ServiceIp
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetServiceName() *string {
	return s.ServiceName
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetSpanID() *string {
	return s.SpanID
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) GetTraceID() *string {
	return s.TraceID
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetDuration(v int64) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.Duration = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetOperationName(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.OperationName = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetServiceIp(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.ServiceIp = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetServiceName(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.ServiceName = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetSpanID(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.SpanID = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetTimestamp(v int64) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.Timestamp = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) SetTraceID(v string) *SearchTracesByPageResponseBodyPageBeanTraceInfos {
	s.TraceID = &v
	return s
}

func (s *SearchTracesByPageResponseBodyPageBeanTraceInfos) Validate() error {
	return dara.Validate(s)
}

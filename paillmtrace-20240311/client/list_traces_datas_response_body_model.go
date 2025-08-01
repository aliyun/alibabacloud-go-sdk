// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTracesDatasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTracesDatasResponseBody
	GetCode() *string
	SetMessage(v string) *ListTracesDatasResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListTracesDatasResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTracesDatasResponseBody
	GetTotalCount() *int32
	SetTraces(v []interface{}) *ListTracesDatasResponseBody
	GetTraces() []interface{}
}

type ListTracesDatasResponseBody struct {
	// The internal error code. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// ExecutionFailure
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. This parameter is returned only when an error occurs.
	//
	// example:
	//
	// failed to get trace data
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// POP request id
	//
	// example:
	//
	// 6A87228C-969A-1381-98CF-AE07AE630FA5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of traces that meet the condition.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The JSON array with each element being a trace\\"s JSON string. Length of the array is equal to or less than the page size parameter value.
	Traces []interface{} `json:"Traces,omitempty" xml:"Traces,omitempty" type:"Repeated"`
}

func (s ListTracesDatasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTracesDatasResponseBody) GoString() string {
	return s.String()
}

func (s *ListTracesDatasResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTracesDatasResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTracesDatasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTracesDatasResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTracesDatasResponseBody) GetTraces() []interface{} {
	return s.Traces
}

func (s *ListTracesDatasResponseBody) SetCode(v string) *ListTracesDatasResponseBody {
	s.Code = &v
	return s
}

func (s *ListTracesDatasResponseBody) SetMessage(v string) *ListTracesDatasResponseBody {
	s.Message = &v
	return s
}

func (s *ListTracesDatasResponseBody) SetRequestId(v string) *ListTracesDatasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTracesDatasResponseBody) SetTotalCount(v int32) *ListTracesDatasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTracesDatasResponseBody) SetTraces(v []interface{}) *ListTracesDatasResponseBody {
	s.Traces = v
	return s
}

func (s *ListTracesDatasResponseBody) Validate() error {
	return dara.Validate(s)
}

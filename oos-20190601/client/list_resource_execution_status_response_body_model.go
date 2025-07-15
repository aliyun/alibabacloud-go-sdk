// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceExecutionStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceExecutionStatusResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceExecutionStatusResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListResourceExecutionStatusResponseBody
	GetRequestId() *string
	SetResourceExecutionStatus(v []*ListResourceExecutionStatusResponseBodyResourceExecutionStatus) *ListResourceExecutionStatusResponseBody
	GetResourceExecutionStatus() []*ListResourceExecutionStatusResponseBodyResourceExecutionStatus
}

type ListResourceExecutionStatusResponseBody struct {
	// The number of entries returned on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// MTRBMDc0NjAtRUJFNy00N0NBLTk3NTctMTJDQzQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ED571CBD-9F96-419D-B919-CF340BBCA157
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution information of the resource.
	ResourceExecutionStatus []*ListResourceExecutionStatusResponseBodyResourceExecutionStatus `json:"ResourceExecutionStatus,omitempty" xml:"ResourceExecutionStatus,omitempty" type:"Repeated"`
}

func (s ListResourceExecutionStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExecutionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceExecutionStatusResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceExecutionStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceExecutionStatusResponseBody) GetResourceExecutionStatus() []*ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	return s.ResourceExecutionStatus
}

func (s *ListResourceExecutionStatusResponseBody) SetMaxResults(v int32) *ListResourceExecutionStatusResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBody) SetNextToken(v string) *ListResourceExecutionStatusResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBody) SetRequestId(v string) *ListResourceExecutionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBody) SetResourceExecutionStatus(v []*ListResourceExecutionStatusResponseBodyResourceExecutionStatus) *ListResourceExecutionStatusResponseBody {
	s.ResourceExecutionStatus = v
	return s
}

func (s *ListResourceExecutionStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListResourceExecutionStatusResponseBodyResourceExecutionStatus struct {
	// The ID of the execution.
	//
	// example:
	//
	// exec-6be6d6ff805349d9ac13
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The time when the execution started running.
	//
	// example:
	//
	// 2020-11-13T08:48:34Z
	ExecutionTime *string `json:"ExecutionTime,omitempty" xml:"ExecutionTime,omitempty"`
	// The output of the template.
	//
	// example:
	//
	// { 				"commandOutput": "hello\\n" 			}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// i-bp1e1bxxxxxxxxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The status of the execution.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListResourceExecutionStatusResponseBodyResourceExecutionStatus) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExecutionStatusResponseBodyResourceExecutionStatus) GoString() string {
	return s.String()
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) GetExecutionTime() *string {
	return s.ExecutionTime
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) GetOutputs() *string {
	return s.Outputs
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) GetStatus() *string {
	return s.Status
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetExecutionId(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.ExecutionId = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetExecutionTime(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.ExecutionTime = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetOutputs(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.Outputs = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetResourceId(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.ResourceId = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) SetStatus(v string) *ListResourceExecutionStatusResponseBodyResourceExecutionStatus {
	s.Status = &v
	return s
}

func (s *ListResourceExecutionStatusResponseBodyResourceExecutionStatus) Validate() error {
	return dara.Validate(s)
}

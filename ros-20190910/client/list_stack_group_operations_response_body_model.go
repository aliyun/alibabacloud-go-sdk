// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackGroupOperationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListStackGroupOperationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStackGroupOperationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListStackGroupOperationsResponseBody
	GetRequestId() *string
	SetStackGroupOperations(v []*ListStackGroupOperationsResponseBodyStackGroupOperations) *ListStackGroupOperationsResponseBody
	GetStackGroupOperations() []*ListStackGroupOperationsResponseBodyStackGroupOperations
	SetTotalCount(v int32) *ListStackGroupOperationsResponseBody
	GetTotalCount() *int32
}

type ListStackGroupOperationsResponseBody struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stack group operations.
	StackGroupOperations []*ListStackGroupOperationsResponseBodyStackGroupOperations `json:"StackGroupOperations,omitempty" xml:"StackGroupOperations,omitempty" type:"Repeated"`
	// The total number of stack group operations.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStackGroupOperationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStackGroupOperationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStackGroupOperationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStackGroupOperationsResponseBody) GetStackGroupOperations() []*ListStackGroupOperationsResponseBodyStackGroupOperations {
	return s.StackGroupOperations
}

func (s *ListStackGroupOperationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListStackGroupOperationsResponseBody) SetPageNumber(v int32) *ListStackGroupOperationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetPageSize(v int32) *ListStackGroupOperationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetRequestId(v string) *ListStackGroupOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetStackGroupOperations(v []*ListStackGroupOperationsResponseBodyStackGroupOperations) *ListStackGroupOperationsResponseBody {
	s.StackGroupOperations = v
	return s
}

func (s *ListStackGroupOperationsResponseBody) SetTotalCount(v int32) *ListStackGroupOperationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStackGroupOperationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListStackGroupOperationsResponseBodyStackGroupOperations struct {
	// The operation type.
	//
	// Valid values:
	//
	// 	- CREATE
	//
	// 	- UPDATE
	//
	// 	- DELETE
	//
	// 	- DETECT_DRIFT
	//
	// example:
	//
	// CREATE
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The time when the operation was initiated.
	//
	// example:
	//
	// 2020-01-20T09:22:36.000000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the operation ended.
	//
	// example:
	//
	// 2020-01-20T09:22:41.000000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The description of the operation.
	//
	// example:
	//
	// Create stack instance in hangzhou
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The operation ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the stack group.
	//
	// example:
	//
	// fd0ddef9-9540-4b42-a464-94f77835****
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The name of the stack group.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The state of the operation.
	//
	// Valid values:
	//
	// 	- RUNNING
	//
	// 	- SUCCEEDED
	//
	// 	- FAILED
	//
	// 	- STOPPING
	//
	// 	- STOPPED
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListStackGroupOperationsResponseBodyStackGroupOperations) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupOperationsResponseBodyStackGroupOperations) GoString() string {
	return s.String()
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) GetAction() *string {
	return s.Action
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) GetEndTime() *string {
	return s.EndTime
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) GetOperationDescription() *string {
	return s.OperationDescription
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) GetOperationId() *string {
	return s.OperationId
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) GetStackGroupId() *string {
	return s.StackGroupId
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) GetStatus() *string {
	return s.Status
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetAction(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.Action = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetCreateTime(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.CreateTime = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetEndTime(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.EndTime = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetOperationDescription(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.OperationDescription = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetOperationId(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.OperationId = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetStackGroupId(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.StackGroupId = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetStackGroupName(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.StackGroupName = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) SetStatus(v string) *ListStackGroupOperationsResponseBodyStackGroupOperations {
	s.Status = &v
	return s
}

func (s *ListStackGroupOperationsResponseBodyStackGroupOperations) Validate() error {
	return dara.Validate(s)
}

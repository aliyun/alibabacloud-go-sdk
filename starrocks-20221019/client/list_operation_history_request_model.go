// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *ListOperationHistoryRequest
	GetEndTime() *int64
	SetInstanceId(v string) *ListOperationHistoryRequest
	GetInstanceId() *string
	SetOperationId(v string) *ListOperationHistoryRequest
	GetOperationId() *string
	SetOperationStatus(v string) *ListOperationHistoryRequest
	GetOperationStatus() *string
	SetOperationType(v string) *ListOperationHistoryRequest
	GetOperationType() *string
	SetPageNumber(v int32) *ListOperationHistoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOperationHistoryRequest
	GetPageSize() *int32
	SetStartTime(v int64) *ListOperationHistoryRequest
	GetStartTime() *int64
}

type ListOperationHistoryRequest struct {
	// example:
	//
	// 1742179008000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// op-f49743caa809****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// example:
	//
	// COMPLETED
	OperationStatus *string `json:"OperationStatus,omitempty" xml:"OperationStatus,omitempty"`
	// example:
	//
	// update_configuration
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1742179008000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListOperationHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationHistoryRequest) GoString() string {
	return s.String()
}

func (s *ListOperationHistoryRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListOperationHistoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationHistoryRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *ListOperationHistoryRequest) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *ListOperationHistoryRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ListOperationHistoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOperationHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOperationHistoryRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListOperationHistoryRequest) SetEndTime(v int64) *ListOperationHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *ListOperationHistoryRequest) SetInstanceId(v string) *ListOperationHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOperationHistoryRequest) SetOperationId(v string) *ListOperationHistoryRequest {
	s.OperationId = &v
	return s
}

func (s *ListOperationHistoryRequest) SetOperationStatus(v string) *ListOperationHistoryRequest {
	s.OperationStatus = &v
	return s
}

func (s *ListOperationHistoryRequest) SetOperationType(v string) *ListOperationHistoryRequest {
	s.OperationType = &v
	return s
}

func (s *ListOperationHistoryRequest) SetPageNumber(v int32) *ListOperationHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOperationHistoryRequest) SetPageSize(v int32) *ListOperationHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationHistoryRequest) SetStartTime(v int64) *ListOperationHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *ListOperationHistoryRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperation interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *Operation
	GetClusterId() *string
	SetCreateTime(v int64) *Operation
	GetCreateTime() *int64
	SetDescription(v string) *Operation
	GetDescription() *string
	SetEndTime(v int64) *Operation
	GetEndTime() *int64
	SetOperationId(v string) *Operation
	GetOperationId() *string
	SetOperationState(v string) *Operation
	GetOperationState() *string
	SetOperationType(v string) *Operation
	GetOperationType() *string
	SetStartTime(v int64) *Operation
	GetStartTime() *int64
	SetStateChangeReason(v *OperationStateChangeReason) *Operation
	GetStateChangeReason() *OperationStateChangeReason
}

type Operation struct {
	// 集群ID。
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 创建时间。
	//
	// example:
	//
	// 1628589439114
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 描述。
	//
	// example:
	//
	// start
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 结束时间。
	//
	// example:
	//
	// 1628589439114
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 操作ID。
	//
	// example:
	//
	// op-13c37a77c505****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// 操作状态。
	//
	// example:
	//
	// COMPLETED
	OperationState *string `json:"OperationState,omitempty" xml:"OperationState,omitempty"`
	// 操作类型。
	//
	// example:
	//
	// CLUSTER
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	// 开始时间。
	//
	// example:
	//
	// 1628589439114
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 状态转换原因。
	StateChangeReason *OperationStateChangeReason `json:"StateChangeReason,omitempty" xml:"StateChangeReason,omitempty"`
}

func (s Operation) String() string {
	return dara.Prettify(s)
}

func (s Operation) GoString() string {
	return s.String()
}

func (s *Operation) GetClusterId() *string {
	return s.ClusterId
}

func (s *Operation) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Operation) GetDescription() *string {
	return s.Description
}

func (s *Operation) GetEndTime() *int64 {
	return s.EndTime
}

func (s *Operation) GetOperationId() *string {
	return s.OperationId
}

func (s *Operation) GetOperationState() *string {
	return s.OperationState
}

func (s *Operation) GetOperationType() *string {
	return s.OperationType
}

func (s *Operation) GetStartTime() *int64 {
	return s.StartTime
}

func (s *Operation) GetStateChangeReason() *OperationStateChangeReason {
	return s.StateChangeReason
}

func (s *Operation) SetClusterId(v string) *Operation {
	s.ClusterId = &v
	return s
}

func (s *Operation) SetCreateTime(v int64) *Operation {
	s.CreateTime = &v
	return s
}

func (s *Operation) SetDescription(v string) *Operation {
	s.Description = &v
	return s
}

func (s *Operation) SetEndTime(v int64) *Operation {
	s.EndTime = &v
	return s
}

func (s *Operation) SetOperationId(v string) *Operation {
	s.OperationId = &v
	return s
}

func (s *Operation) SetOperationState(v string) *Operation {
	s.OperationState = &v
	return s
}

func (s *Operation) SetOperationType(v string) *Operation {
	s.OperationType = &v
	return s
}

func (s *Operation) SetStartTime(v int64) *Operation {
	s.StartTime = &v
	return s
}

func (s *Operation) SetStateChangeReason(v *OperationStateChangeReason) *Operation {
	s.StateChangeReason = v
	return s
}

func (s *Operation) Validate() error {
	return dara.Validate(s)
}

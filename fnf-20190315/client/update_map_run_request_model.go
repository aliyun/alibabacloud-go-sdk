// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMapRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int64) *UpdateMapRunRequest
	GetConcurrency() *int64
	SetExecutionName(v string) *UpdateMapRunRequest
	GetExecutionName() *string
	SetFlowName(v string) *UpdateMapRunRequest
	GetFlowName() *string
	SetMapRunName(v string) *UpdateMapRunRequest
	GetMapRunName() *string
	SetRequestId(v string) *UpdateMapRunRequest
	GetRequestId() *string
	SetToleratedFailedCount(v int64) *UpdateMapRunRequest
	GetToleratedFailedCount() *int64
	SetToleratedFailedPercentage(v float32) *UpdateMapRunRequest
	GetToleratedFailedPercentage() *float32
}

type UpdateMapRunRequest struct {
	// example:
	//
	// 1
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my_exec_name
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my_flow_name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c39142f1345b196d678333c41f113100
	MapRunName *string `json:"MapRunName,omitempty" xml:"MapRunName,omitempty"`
	// example:
	//
	// 3A44E113-9962-5B0B-AB92-14060EFE3164
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	ToleratedFailedCount *int64 `json:"ToleratedFailedCount,omitempty" xml:"ToleratedFailedCount,omitempty"`
	// example:
	//
	// 20
	ToleratedFailedPercentage *float32 `json:"ToleratedFailedPercentage,omitempty" xml:"ToleratedFailedPercentage,omitempty"`
}

func (s UpdateMapRunRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMapRunRequest) GoString() string {
	return s.String()
}

func (s *UpdateMapRunRequest) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *UpdateMapRunRequest) GetExecutionName() *string {
	return s.ExecutionName
}

func (s *UpdateMapRunRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *UpdateMapRunRequest) GetMapRunName() *string {
	return s.MapRunName
}

func (s *UpdateMapRunRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMapRunRequest) GetToleratedFailedCount() *int64 {
	return s.ToleratedFailedCount
}

func (s *UpdateMapRunRequest) GetToleratedFailedPercentage() *float32 {
	return s.ToleratedFailedPercentage
}

func (s *UpdateMapRunRequest) SetConcurrency(v int64) *UpdateMapRunRequest {
	s.Concurrency = &v
	return s
}

func (s *UpdateMapRunRequest) SetExecutionName(v string) *UpdateMapRunRequest {
	s.ExecutionName = &v
	return s
}

func (s *UpdateMapRunRequest) SetFlowName(v string) *UpdateMapRunRequest {
	s.FlowName = &v
	return s
}

func (s *UpdateMapRunRequest) SetMapRunName(v string) *UpdateMapRunRequest {
	s.MapRunName = &v
	return s
}

func (s *UpdateMapRunRequest) SetRequestId(v string) *UpdateMapRunRequest {
	s.RequestId = &v
	return s
}

func (s *UpdateMapRunRequest) SetToleratedFailedCount(v int64) *UpdateMapRunRequest {
	s.ToleratedFailedCount = &v
	return s
}

func (s *UpdateMapRunRequest) SetToleratedFailedPercentage(v float32) *UpdateMapRunRequest {
	s.ToleratedFailedPercentage = &v
	return s
}

func (s *UpdateMapRunRequest) Validate() error {
	return dara.Validate(s)
}

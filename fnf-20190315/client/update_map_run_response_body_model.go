// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMapRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int64) *UpdateMapRunResponseBody
	GetConcurrency() *int64
	SetExecutionName(v string) *UpdateMapRunResponseBody
	GetExecutionName() *string
	SetFlowName(v string) *UpdateMapRunResponseBody
	GetFlowName() *string
	SetMapRunName(v string) *UpdateMapRunResponseBody
	GetMapRunName() *string
	SetRequestId(v string) *UpdateMapRunResponseBody
	GetRequestId() *string
	SetToleratedFailedCount(v int64) *UpdateMapRunResponseBody
	GetToleratedFailedCount() *int64
	SetToleratedFailedPercentage(v float32) *UpdateMapRunResponseBody
	GetToleratedFailedPercentage() *float32
}

type UpdateMapRunResponseBody struct {
	// example:
	//
	// 1
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// example:
	//
	// my_exec_name
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// example:
	//
	// my_flow_name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// example:
	//
	// c39142f1345b196d678333c41f113000
	MapRunName *string `json:"MapRunName,omitempty" xml:"MapRunName,omitempty"`
	// Id of the request
	//
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

func (s UpdateMapRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMapRunResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMapRunResponseBody) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *UpdateMapRunResponseBody) GetExecutionName() *string {
	return s.ExecutionName
}

func (s *UpdateMapRunResponseBody) GetFlowName() *string {
	return s.FlowName
}

func (s *UpdateMapRunResponseBody) GetMapRunName() *string {
	return s.MapRunName
}

func (s *UpdateMapRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMapRunResponseBody) GetToleratedFailedCount() *int64 {
	return s.ToleratedFailedCount
}

func (s *UpdateMapRunResponseBody) GetToleratedFailedPercentage() *float32 {
	return s.ToleratedFailedPercentage
}

func (s *UpdateMapRunResponseBody) SetConcurrency(v int64) *UpdateMapRunResponseBody {
	s.Concurrency = &v
	return s
}

func (s *UpdateMapRunResponseBody) SetExecutionName(v string) *UpdateMapRunResponseBody {
	s.ExecutionName = &v
	return s
}

func (s *UpdateMapRunResponseBody) SetFlowName(v string) *UpdateMapRunResponseBody {
	s.FlowName = &v
	return s
}

func (s *UpdateMapRunResponseBody) SetMapRunName(v string) *UpdateMapRunResponseBody {
	s.MapRunName = &v
	return s
}

func (s *UpdateMapRunResponseBody) SetRequestId(v string) *UpdateMapRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMapRunResponseBody) SetToleratedFailedCount(v int64) *UpdateMapRunResponseBody {
	s.ToleratedFailedCount = &v
	return s
}

func (s *UpdateMapRunResponseBody) SetToleratedFailedPercentage(v float32) *UpdateMapRunResponseBody {
	s.ToleratedFailedPercentage = &v
	return s
}

func (s *UpdateMapRunResponseBody) Validate() error {
	return dara.Validate(s)
}

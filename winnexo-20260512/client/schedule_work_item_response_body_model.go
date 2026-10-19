// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduleWorkItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ScheduleWorkItemResponseBody
	GetCode() *string
	SetCreated(v bool) *ScheduleWorkItemResponseBody
	GetCreated() *bool
	SetMessage(v string) *ScheduleWorkItemResponseBody
	GetMessage() *string
	SetPlanId(v string) *ScheduleWorkItemResponseBody
	GetPlanId() *string
	SetRequestId(v string) *ScheduleWorkItemResponseBody
	GetRequestId() *string
	SetRevision(v int64) *ScheduleWorkItemResponseBody
	GetRevision() *int64
	SetRunId(v string) *ScheduleWorkItemResponseBody
	GetRunId() *string
	SetRunSeq(v int64) *ScheduleWorkItemResponseBody
	GetRunSeq() *int64
	SetWorkItemId(v string) *ScheduleWorkItemResponseBody
	GetWorkItemId() *string
}

type ScheduleWorkItemResponseBody struct {
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 本次是否新创建 Plan
	//
	// example:
	//
	// false
	Created *bool `json:"created,omitempty" xml:"created,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 物理计划ID
	//
	// example:
	//
	// example
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 修订号
	//
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// 运行ID
	//
	// example:
	//
	// example
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// 运行序号
	//
	// example:
	//
	// 1
	RunSeq *int64 `json:"runSeq,omitempty" xml:"runSeq,omitempty"`
	// 工作项ID
	//
	// example:
	//
	// example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s ScheduleWorkItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScheduleWorkItemResponseBody) GoString() string {
	return s.String()
}

func (s *ScheduleWorkItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *ScheduleWorkItemResponseBody) GetCreated() *bool {
	return s.Created
}

func (s *ScheduleWorkItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ScheduleWorkItemResponseBody) GetPlanId() *string {
	return s.PlanId
}

func (s *ScheduleWorkItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScheduleWorkItemResponseBody) GetRevision() *int64 {
	return s.Revision
}

func (s *ScheduleWorkItemResponseBody) GetRunId() *string {
	return s.RunId
}

func (s *ScheduleWorkItemResponseBody) GetRunSeq() *int64 {
	return s.RunSeq
}

func (s *ScheduleWorkItemResponseBody) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ScheduleWorkItemResponseBody) SetCode(v string) *ScheduleWorkItemResponseBody {
	s.Code = &v
	return s
}

func (s *ScheduleWorkItemResponseBody) SetCreated(v bool) *ScheduleWorkItemResponseBody {
	s.Created = &v
	return s
}

func (s *ScheduleWorkItemResponseBody) SetMessage(v string) *ScheduleWorkItemResponseBody {
	s.Message = &v
	return s
}

func (s *ScheduleWorkItemResponseBody) SetPlanId(v string) *ScheduleWorkItemResponseBody {
	s.PlanId = &v
	return s
}

func (s *ScheduleWorkItemResponseBody) SetRequestId(v string) *ScheduleWorkItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScheduleWorkItemResponseBody) SetRevision(v int64) *ScheduleWorkItemResponseBody {
	s.Revision = &v
	return s
}

func (s *ScheduleWorkItemResponseBody) SetRunId(v string) *ScheduleWorkItemResponseBody {
	s.RunId = &v
	return s
}

func (s *ScheduleWorkItemResponseBody) SetRunSeq(v int64) *ScheduleWorkItemResponseBody {
	s.RunSeq = &v
	return s
}

func (s *ScheduleWorkItemResponseBody) SetWorkItemId(v string) *ScheduleWorkItemResponseBody {
	s.WorkItemId = &v
	return s
}

func (s *ScheduleWorkItemResponseBody) Validate() error {
	return dara.Validate(s)
}

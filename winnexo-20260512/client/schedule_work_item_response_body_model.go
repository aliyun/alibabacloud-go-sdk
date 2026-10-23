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
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether a new plan is created in this request.
	//
	// example:
	//
	// false
	Created *bool `json:"created,omitempty" xml:"created,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The ID of the physical plan.
	//
	// example:
	//
	// example
	PlanId *string `json:"planId,omitempty" xml:"planId,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The revision number.
	//
	// example:
	//
	// 1
	Revision *int64 `json:"revision,omitempty" xml:"revision,omitempty"`
	// The run ID.
	//
	// example:
	//
	// example
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// The run sequence number.
	//
	// example:
	//
	// 1
	RunSeq *int64 `json:"runSeq,omitempty" xml:"runSeq,omitempty"`
	// The ID of the work item.
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

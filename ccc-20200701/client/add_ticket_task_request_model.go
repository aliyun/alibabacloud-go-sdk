// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTicketTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssignee(v string) *AddTicketTaskRequest
	GetAssignee() *string
	SetComment(v string) *AddTicketTaskRequest
	GetComment() *string
	SetInstanceId(v string) *AddTicketTaskRequest
	GetInstanceId() *string
	SetPosition(v string) *AddTicketTaskRequest
	GetPosition() *string
	SetTaskId(v string) *AddTicketTaskRequest
	GetTaskId() *string
	SetTicketId(v string) *AddTicketTaskRequest
	GetTicketId() *string
}

type AddTicketTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// assignee@ccc-test
	Assignee *string `json:"Assignee,omitempty" xml:"Assignee,omitempty"`
	// This parameter is required.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// After
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3b8c672f-48f6-45f5-bf41-9d4cb2b4a716
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3d26b90a-c5d2-4b09-8219-********
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s AddTicketTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTicketTaskRequest) GoString() string {
	return s.String()
}

func (s *AddTicketTaskRequest) GetAssignee() *string {
	return s.Assignee
}

func (s *AddTicketTaskRequest) GetComment() *string {
	return s.Comment
}

func (s *AddTicketTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddTicketTaskRequest) GetPosition() *string {
	return s.Position
}

func (s *AddTicketTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AddTicketTaskRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *AddTicketTaskRequest) SetAssignee(v string) *AddTicketTaskRequest {
	s.Assignee = &v
	return s
}

func (s *AddTicketTaskRequest) SetComment(v string) *AddTicketTaskRequest {
	s.Comment = &v
	return s
}

func (s *AddTicketTaskRequest) SetInstanceId(v string) *AddTicketTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *AddTicketTaskRequest) SetPosition(v string) *AddTicketTaskRequest {
	s.Position = &v
	return s
}

func (s *AddTicketTaskRequest) SetTaskId(v string) *AddTicketTaskRequest {
	s.TaskId = &v
	return s
}

func (s *AddTicketTaskRequest) SetTicketId(v string) *AddTicketTaskRequest {
	s.TicketId = &v
	return s
}

func (s *AddTicketTaskRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishTicketTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *FinishTicketTaskRequest
	GetComment() *string
	SetInstanceId(v string) *FinishTicketTaskRequest
	GetInstanceId() *string
	SetTaskId(v string) *FinishTicketTaskRequest
	GetTaskId() *string
	SetTicketId(v string) *FinishTicketTaskRequest
	GetTicketId() *string
}

type FinishTicketTaskRequest struct {
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
	// b52a34dc-f514-4600-9c39-3cf657167c97
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// f2c6722b-cd13-442d-bf10-22a07c70d6d5
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s FinishTicketTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s FinishTicketTaskRequest) GoString() string {
	return s.String()
}

func (s *FinishTicketTaskRequest) GetComment() *string {
	return s.Comment
}

func (s *FinishTicketTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *FinishTicketTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *FinishTicketTaskRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *FinishTicketTaskRequest) SetComment(v string) *FinishTicketTaskRequest {
	s.Comment = &v
	return s
}

func (s *FinishTicketTaskRequest) SetInstanceId(v string) *FinishTicketTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *FinishTicketTaskRequest) SetTaskId(v string) *FinishTicketTaskRequest {
	s.TaskId = &v
	return s
}

func (s *FinishTicketTaskRequest) SetTicketId(v string) *FinishTicketTaskRequest {
	s.TicketId = &v
	return s
}

func (s *FinishTicketTaskRequest) Validate() error {
	return dara.Validate(s)
}

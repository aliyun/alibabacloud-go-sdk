// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferTicketTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssignee(v string) *TransferTicketTaskRequest
	GetAssignee() *string
	SetComment(v string) *TransferTicketTaskRequest
	GetComment() *string
	SetInstanceId(v string) *TransferTicketTaskRequest
	GetInstanceId() *string
	SetTaskId(v string) *TransferTicketTaskRequest
	GetTaskId() *string
	SetTicketId(v string) *TransferTicketTaskRequest
	GetTicketId() *string
}

type TransferTicketTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// assignee@ccc-test
	Assignee *string `json:"Assignee,omitempty" xml:"Assignee,omitempty"`
	Comment  *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// f780ade8-3ca9-458b-b067-63077946a570
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 5491d3b4-14ee-4341-b5f1-db2c78beddeb
	TicketId *string `json:"TicketId,omitempty" xml:"TicketId,omitempty"`
}

func (s TransferTicketTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferTicketTaskRequest) GoString() string {
	return s.String()
}

func (s *TransferTicketTaskRequest) GetAssignee() *string {
	return s.Assignee
}

func (s *TransferTicketTaskRequest) GetComment() *string {
	return s.Comment
}

func (s *TransferTicketTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TransferTicketTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *TransferTicketTaskRequest) GetTicketId() *string {
	return s.TicketId
}

func (s *TransferTicketTaskRequest) SetAssignee(v string) *TransferTicketTaskRequest {
	s.Assignee = &v
	return s
}

func (s *TransferTicketTaskRequest) SetComment(v string) *TransferTicketTaskRequest {
	s.Comment = &v
	return s
}

func (s *TransferTicketTaskRequest) SetInstanceId(v string) *TransferTicketTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *TransferTicketTaskRequest) SetTaskId(v string) *TransferTicketTaskRequest {
	s.TaskId = &v
	return s
}

func (s *TransferTicketTaskRequest) SetTicketId(v string) *TransferTicketTaskRequest {
	s.TicketId = &v
	return s
}

func (s *TransferTicketTaskRequest) Validate() error {
	return dara.Validate(s)
}

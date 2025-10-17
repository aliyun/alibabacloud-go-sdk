// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetP2PToken(v string) *GetConnectionTicketResponseBody
	GetP2PToken() *string
	SetRequestId(v string) *GetConnectionTicketResponseBody
	GetRequestId() *string
	SetTaskCode(v string) *GetConnectionTicketResponseBody
	GetTaskCode() *string
	SetTaskId(v string) *GetConnectionTicketResponseBody
	GetTaskId() *string
	SetTaskMessage(v string) *GetConnectionTicketResponseBody
	GetTaskMessage() *string
	SetTaskStatus(v string) *GetConnectionTicketResponseBody
	GetTaskStatus() *string
	SetTicket(v string) *GetConnectionTicketResponseBody
	GetTicket() *string
}

type GetConnectionTicketResponseBody struct {
	P2PToken *string `json:"P2PToken,omitempty" xml:"P2PToken,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskCode  *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	// example:
	//
	// 2afbad19-778a-4fc5-9674-1f19c63862da
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskMessage *string `json:"TaskMessage,omitempty" xml:"TaskMessage,omitempty"`
	// example:
	//
	// FINISHED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// W0VuY29kaW5nXQ0KSW5wdXRFbmNvZGluZz1V********
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetConnectionTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetConnectionTicketResponseBody) GetP2PToken() *string {
	return s.P2PToken
}

func (s *GetConnectionTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConnectionTicketResponseBody) GetTaskCode() *string {
	return s.TaskCode
}

func (s *GetConnectionTicketResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetConnectionTicketResponseBody) GetTaskMessage() *string {
	return s.TaskMessage
}

func (s *GetConnectionTicketResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetConnectionTicketResponseBody) GetTicket() *string {
	return s.Ticket
}

func (s *GetConnectionTicketResponseBody) SetP2PToken(v string) *GetConnectionTicketResponseBody {
	s.P2PToken = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetRequestId(v string) *GetConnectionTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskCode(v string) *GetConnectionTicketResponseBody {
	s.TaskCode = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskId(v string) *GetConnectionTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskMessage(v string) *GetConnectionTicketResponseBody {
	s.TaskMessage = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTaskStatus(v string) *GetConnectionTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetConnectionTicketResponseBody) SetTicket(v string) *GetConnectionTicketResponseBody {
	s.Ticket = &v
	return s
}

func (s *GetConnectionTicketResponseBody) Validate() error {
	return dara.Validate(s)
}

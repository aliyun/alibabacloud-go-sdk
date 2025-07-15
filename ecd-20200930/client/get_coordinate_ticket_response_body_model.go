// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCoordinateTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCoId(v string) *GetCoordinateTicketResponseBody
	GetCoId() *string
	SetRequestId(v string) *GetCoordinateTicketResponseBody
	GetRequestId() *string
	SetTaskId(v string) *GetCoordinateTicketResponseBody
	GetTaskId() *string
	SetTaskStatus(v string) *GetCoordinateTicketResponseBody
	GetTaskStatus() *string
	SetTicket(v string) *GetCoordinateTicketResponseBody
	GetTicket() *string
}

type GetCoordinateTicketResponseBody struct {
	// The ID of the stream collaboration.
	//
	// example:
	//
	// co-0sot77uale3****
	CoId *string `json:"CoId,omitempty" xml:"CoId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the cloud computer connection task.
	//
	// example:
	//
	// 39cc15e5-6998-4b9f-9b2c-7a4cc3e2****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task status.
	//
	// Possible values:
	//
	// 	- Finished
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Failed
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- Running
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// Finished
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The credentials of the stream collaboration.
	//
	// example:
	//
	// W0VuY29kaW5nXQ0KSW5wdXRFbmNvZGluZz1V********
	Ticket *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
}

func (s GetCoordinateTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCoordinateTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetCoordinateTicketResponseBody) GetCoId() *string {
	return s.CoId
}

func (s *GetCoordinateTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCoordinateTicketResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCoordinateTicketResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *GetCoordinateTicketResponseBody) GetTicket() *string {
	return s.Ticket
}

func (s *GetCoordinateTicketResponseBody) SetCoId(v string) *GetCoordinateTicketResponseBody {
	s.CoId = &v
	return s
}

func (s *GetCoordinateTicketResponseBody) SetRequestId(v string) *GetCoordinateTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCoordinateTicketResponseBody) SetTaskId(v string) *GetCoordinateTicketResponseBody {
	s.TaskId = &v
	return s
}

func (s *GetCoordinateTicketResponseBody) SetTaskStatus(v string) *GetCoordinateTicketResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *GetCoordinateTicketResponseBody) SetTicket(v string) *GetCoordinateTicketResponseBody {
	s.Ticket = &v
	return s
}

func (s *GetCoordinateTicketResponseBody) Validate() error {
	return dara.Validate(s)
}

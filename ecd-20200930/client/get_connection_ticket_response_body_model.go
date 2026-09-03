// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *GetConnectionTicketResponseBody
	GetDesktopId() *string
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
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ticket used to connect to the user instance. Before you use the ticket, Base64-decode its content and save it as an xxx.ica file, and then open the file. The following Python code provides an example:
	//
	// ```
	//
	// import base64
	//
	// response = {
	//
	//     "Ticket": "W0VuY29kaW5nXQ0KSW5wdXRFbmNvZGluZz1V********",
	//
	//     "RequestId": "1CBAFFAB-B697-4049-A9B1-67E1FC5F****",
	//
	// }
	//
	// f = open (\\"xxx.ica\\", \\"w\\")
	//
	// out = base64.b64decode(response[\\"Ticket\\"])
	//
	// f.write(out)
	//
	// f.close()
	//
	// ```
	//
	// example:
	//
	// W0VuY29kaW5nXQ0KSW5wdXRFbmNvZGluZz1V********
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	// The ID of the cloud computer connection task.
	//
	// example:
	//
	// 2afbad19-778a-4fc5-9674-1f19c638****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The ID of the cloud computer connection task.
	//
	// example:
	//
	// 2afbad19-778a-4fc5-9674-1f19c638****
	TaskMessage *string `json:"TaskMessage,omitempty" xml:"TaskMessage,omitempty"`
	// The task status.
	//
	// example:
	//
	// FINISHED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The connection credential for the cloud computer.
	//
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

func (s *GetConnectionTicketResponseBody) GetDesktopId() *string {
	return s.DesktopId
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

func (s *GetConnectionTicketResponseBody) SetDesktopId(v string) *GetConnectionTicketResponseBody {
	s.DesktopId = &v
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

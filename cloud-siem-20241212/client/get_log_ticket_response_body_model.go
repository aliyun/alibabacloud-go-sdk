// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogTicket(v string) *GetLogTicketResponseBody
	GetLogTicket() *string
	SetRequestId(v string) *GetLogTicketResponseBody
	GetRequestId() *string
}

type GetLogTicketResponseBody struct {
	// example:
	//
	// *******。
	LogTicket *string `json:"LogTicket,omitempty" xml:"LogTicket,omitempty"`
	// example:
	//
	// 173326*******。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLogTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLogTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetLogTicketResponseBody) GetLogTicket() *string {
	return s.LogTicket
}

func (s *GetLogTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLogTicketResponseBody) SetLogTicket(v string) *GetLogTicketResponseBody {
	s.LogTicket = &v
	return s
}

func (s *GetLogTicketResponseBody) SetRequestId(v string) *GetLogTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLogTicketResponseBody) Validate() error {
	return dara.Validate(s)
}

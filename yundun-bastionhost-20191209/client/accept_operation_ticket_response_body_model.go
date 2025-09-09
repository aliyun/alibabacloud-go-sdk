// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAcceptOperationTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AcceptOperationTicketResponseBody
	GetRequestId() *string
}

type AcceptOperationTicketResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AcceptOperationTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AcceptOperationTicketResponseBody) GoString() string {
	return s.String()
}

func (s *AcceptOperationTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AcceptOperationTicketResponseBody) SetRequestId(v string) *AcceptOperationTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *AcceptOperationTicketResponseBody) Validate() error {
	return dara.Validate(s)
}

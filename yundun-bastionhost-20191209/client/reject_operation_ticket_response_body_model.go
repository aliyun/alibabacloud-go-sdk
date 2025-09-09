// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectOperationTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RejectOperationTicketResponseBody
	GetRequestId() *string
}

type RejectOperationTicketResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RejectOperationTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RejectOperationTicketResponseBody) GoString() string {
	return s.String()
}

func (s *RejectOperationTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RejectOperationTicketResponseBody) SetRequestId(v string) *RejectOperationTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *RejectOperationTicketResponseBody) Validate() error {
	return dara.Validate(s)
}

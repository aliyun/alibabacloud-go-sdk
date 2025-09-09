// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOperationTicketResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOperationTicketId(v string) *CreateOperationTicketResponseBody
	GetOperationTicketId() *string
	SetRequestId(v string) *CreateOperationTicketResponseBody
	GetRequestId() *string
}

type CreateOperationTicketResponseBody struct {
	// example:
	//
	// 2
	OperationTicketId *string `json:"OperationTicketId,omitempty" xml:"OperationTicketId,omitempty"`
	// example:
	//
	// 0ECCC399-4D35-48A7-8379-5C6180E66235
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOperationTicketResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOperationTicketResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOperationTicketResponseBody) GetOperationTicketId() *string {
	return s.OperationTicketId
}

func (s *CreateOperationTicketResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOperationTicketResponseBody) SetOperationTicketId(v string) *CreateOperationTicketResponseBody {
	s.OperationTicketId = &v
	return s
}

func (s *CreateOperationTicketResponseBody) SetRequestId(v string) *CreateOperationTicketResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOperationTicketResponseBody) Validate() error {
	return dara.Validate(s)
}

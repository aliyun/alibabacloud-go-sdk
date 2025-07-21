// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReceiverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReceiverId(v string) *CreateReceiverResponseBody
	GetReceiverId() *string
	SetRequestId(v string) *CreateReceiverResponseBody
	GetRequestId() *string
}

type CreateReceiverResponseBody struct {
	// Receiver list ID
	//
	// example:
	//
	// 7312e09b8fffc5c7b2e2fbf5b6dc2073
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateReceiverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateReceiverResponseBody) GoString() string {
	return s.String()
}

func (s *CreateReceiverResponseBody) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *CreateReceiverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateReceiverResponseBody) SetReceiverId(v string) *CreateReceiverResponseBody {
	s.ReceiverId = &v
	return s
}

func (s *CreateReceiverResponseBody) SetRequestId(v string) *CreateReceiverResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateReceiverResponseBody) Validate() error {
	return dara.Validate(s)
}

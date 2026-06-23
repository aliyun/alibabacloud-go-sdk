// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageBodyMD5(v string) *SendMessageResponseBody
	GetMessageBodyMD5() *string
	SetMessageId(v string) *SendMessageResponseBody
	GetMessageId() *string
	SetReceiptHandle(v string) *SendMessageResponseBody
	GetReceiptHandle() *string
}

type SendMessageResponseBody struct {
	MessageBodyMD5 *string `json:"MessageBodyMD5,omitempty" xml:"MessageBodyMD5,omitempty"`
	MessageId      *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	ReceiptHandle  *string `json:"ReceiptHandle,omitempty" xml:"ReceiptHandle,omitempty"`
}

func (s SendMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendMessageResponseBody) GetMessageBodyMD5() *string {
	return s.MessageBodyMD5
}

func (s *SendMessageResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *SendMessageResponseBody) GetReceiptHandle() *string {
	return s.ReceiptHandle
}

func (s *SendMessageResponseBody) SetMessageBodyMD5(v string) *SendMessageResponseBody {
	s.MessageBodyMD5 = &v
	return s
}

func (s *SendMessageResponseBody) SetMessageId(v string) *SendMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendMessageResponseBody) SetReceiptHandle(v string) *SendMessageResponseBody {
	s.ReceiptHandle = &v
	return s
}

func (s *SendMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

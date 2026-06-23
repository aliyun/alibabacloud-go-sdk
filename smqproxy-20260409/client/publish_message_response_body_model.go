// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageBodyMD5(v string) *PublishMessageResponseBody
	GetMessageBodyMD5() *string
	SetMessageId(v string) *PublishMessageResponseBody
	GetMessageId() *string
}

type PublishMessageResponseBody struct {
	MessageBodyMD5 *string `json:"MessageBodyMD5,omitempty" xml:"MessageBodyMD5,omitempty"`
	MessageId      *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s PublishMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PublishMessageResponseBody) GetMessageBodyMD5() *string {
	return s.MessageBodyMD5
}

func (s *PublishMessageResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *PublishMessageResponseBody) SetMessageBodyMD5(v string) *PublishMessageResponseBody {
	s.MessageBodyMD5 = &v
	return s
}

func (s *PublishMessageResponseBody) SetMessageId(v string) *PublishMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *PublishMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSendMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*BatchSendMessageResponseBodyMessages) *BatchSendMessageResponseBody
	GetMessages() []*BatchSendMessageResponseBodyMessages
}

type BatchSendMessageResponseBody struct {
	Messages []*BatchSendMessageResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
}

func (s BatchSendMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMessageResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSendMessageResponseBody) GetMessages() []*BatchSendMessageResponseBodyMessages {
	return s.Messages
}

func (s *BatchSendMessageResponseBody) SetMessages(v []*BatchSendMessageResponseBodyMessages) *BatchSendMessageResponseBody {
	s.Messages = v
	return s
}

func (s *BatchSendMessageResponseBody) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchSendMessageResponseBodyMessages struct {
	MessageBodyMD5 *string `json:"MessageBodyMD5,omitempty" xml:"MessageBodyMD5,omitempty"`
	MessageId      *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s BatchSendMessageResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s BatchSendMessageResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *BatchSendMessageResponseBodyMessages) GetMessageBodyMD5() *string {
	return s.MessageBodyMD5
}

func (s *BatchSendMessageResponseBodyMessages) GetMessageId() *string {
	return s.MessageId
}

func (s *BatchSendMessageResponseBodyMessages) SetMessageBodyMD5(v string) *BatchSendMessageResponseBodyMessages {
	s.MessageBodyMD5 = &v
	return s
}

func (s *BatchSendMessageResponseBodyMessages) SetMessageId(v string) *BatchSendMessageResponseBodyMessages {
	s.MessageId = &v
	return s
}

func (s *BatchSendMessageResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}

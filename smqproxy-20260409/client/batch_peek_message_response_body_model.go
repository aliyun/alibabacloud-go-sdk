// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPeekMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*BatchPeekMessageResponseBodyMessages) *BatchPeekMessageResponseBody
	GetMessages() []*BatchPeekMessageResponseBodyMessages
}

type BatchPeekMessageResponseBody struct {
	Messages []*BatchPeekMessageResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
}

func (s BatchPeekMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchPeekMessageResponseBody) GoString() string {
	return s.String()
}

func (s *BatchPeekMessageResponseBody) GetMessages() []*BatchPeekMessageResponseBodyMessages {
	return s.Messages
}

func (s *BatchPeekMessageResponseBody) SetMessages(v []*BatchPeekMessageResponseBodyMessages) *BatchPeekMessageResponseBody {
	s.Messages = v
	return s
}

func (s *BatchPeekMessageResponseBody) Validate() error {
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

type BatchPeekMessageResponseBodyMessages struct {
	DequeueCount     *int64  `json:"DequeueCount,omitempty" xml:"DequeueCount,omitempty"`
	EnqueueTime      *int64  `json:"EnqueueTime,omitempty" xml:"EnqueueTime,omitempty"`
	FirstDequeueTime *int64  `json:"FirstDequeueTime,omitempty" xml:"FirstDequeueTime,omitempty"`
	MessageBody      *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	MessageBodyMD5   *string `json:"MessageBodyMD5,omitempty" xml:"MessageBodyMD5,omitempty"`
	MessageGroupId   *string `json:"MessageGroupId,omitempty" xml:"MessageGroupId,omitempty"`
	MessageId        *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	Priority         *int64  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	UserProperties   *string `json:"UserProperties,omitempty" xml:"UserProperties,omitempty"`
}

func (s BatchPeekMessageResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s BatchPeekMessageResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *BatchPeekMessageResponseBodyMessages) GetDequeueCount() *int64 {
	return s.DequeueCount
}

func (s *BatchPeekMessageResponseBodyMessages) GetEnqueueTime() *int64 {
	return s.EnqueueTime
}

func (s *BatchPeekMessageResponseBodyMessages) GetFirstDequeueTime() *int64 {
	return s.FirstDequeueTime
}

func (s *BatchPeekMessageResponseBodyMessages) GetMessageBody() *string {
	return s.MessageBody
}

func (s *BatchPeekMessageResponseBodyMessages) GetMessageBodyMD5() *string {
	return s.MessageBodyMD5
}

func (s *BatchPeekMessageResponseBodyMessages) GetMessageGroupId() *string {
	return s.MessageGroupId
}

func (s *BatchPeekMessageResponseBodyMessages) GetMessageId() *string {
	return s.MessageId
}

func (s *BatchPeekMessageResponseBodyMessages) GetPriority() *int64 {
	return s.Priority
}

func (s *BatchPeekMessageResponseBodyMessages) GetUserProperties() *string {
	return s.UserProperties
}

func (s *BatchPeekMessageResponseBodyMessages) SetDequeueCount(v int64) *BatchPeekMessageResponseBodyMessages {
	s.DequeueCount = &v
	return s
}

func (s *BatchPeekMessageResponseBodyMessages) SetEnqueueTime(v int64) *BatchPeekMessageResponseBodyMessages {
	s.EnqueueTime = &v
	return s
}

func (s *BatchPeekMessageResponseBodyMessages) SetFirstDequeueTime(v int64) *BatchPeekMessageResponseBodyMessages {
	s.FirstDequeueTime = &v
	return s
}

func (s *BatchPeekMessageResponseBodyMessages) SetMessageBody(v string) *BatchPeekMessageResponseBodyMessages {
	s.MessageBody = &v
	return s
}

func (s *BatchPeekMessageResponseBodyMessages) SetMessageBodyMD5(v string) *BatchPeekMessageResponseBodyMessages {
	s.MessageBodyMD5 = &v
	return s
}

func (s *BatchPeekMessageResponseBodyMessages) SetMessageGroupId(v string) *BatchPeekMessageResponseBodyMessages {
	s.MessageGroupId = &v
	return s
}

func (s *BatchPeekMessageResponseBodyMessages) SetMessageId(v string) *BatchPeekMessageResponseBodyMessages {
	s.MessageId = &v
	return s
}

func (s *BatchPeekMessageResponseBodyMessages) SetPriority(v int64) *BatchPeekMessageResponseBodyMessages {
	s.Priority = &v
	return s
}

func (s *BatchPeekMessageResponseBodyMessages) SetUserProperties(v string) *BatchPeekMessageResponseBodyMessages {
	s.UserProperties = &v
	return s
}

func (s *BatchPeekMessageResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}

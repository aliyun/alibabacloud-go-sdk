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
	// The message list.
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
	// The total number of times the message has been consumed.
	//
	// example:
	//
	// 5
	DequeueCount *int64 `json:"DequeueCount,omitempty" xml:"DequeueCount,omitempty"`
	// The time when the message was sent to the queue. The value is a UNIX timestamp that starts from 1970-01-01 00:00:00.000.
	//
	// example:
	//
	// 1250700979248
	EnqueueTime *int64 `json:"EnqueueTime,omitempty" xml:"EnqueueTime,omitempty"`
	// The time when the message was first consumed. The value is a UNIX timestamp in milliseconds that starts from 1970-01-01 00:00:00.000. If DequeueCount is 0, this value is the same as EnqueueTime.
	//
	// example:
	//
	// 1250700979348
	FirstDequeueTime *int64 `json:"FirstDequeueTime,omitempty" xml:"FirstDequeueTime,omitempty"`
	// The message body.
	//
	// example:
	//
	// This is a test message
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	// The MD5 hash of the message body.
	//
	// example:
	//
	// F9360F391579E71CA77BC5D50242****
	MessageBodyMD5 *string `json:"MessageBodyMD5,omitempty" xml:"MessageBodyMD5,omitempty"`
	// The message group ID (FIFO queues only).
	//
	// example:
	//
	// test-group
	MessageGroupId *string `json:"MessageGroupId,omitempty" xml:"MessageGroupId,omitempty"`
	// The message ID.
	//
	// example:
	//
	// C0A80255D7D1163042EA3F2814D055F1
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The priority. Valid values: 1 to 16. Default value: 8. A value of 1 indicates the highest priority.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The user-defined properties (JSON-formatted string).
	//
	// example:
	//
	// {"properties1":"value"}
	UserProperties map[string]*MessagesUserPropertiesValue `json:"UserProperties,omitempty" xml:"UserProperties,omitempty"`
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

func (s *BatchPeekMessageResponseBodyMessages) GetUserProperties() map[string]*MessagesUserPropertiesValue {
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

func (s *BatchPeekMessageResponseBodyMessages) SetUserProperties(v map[string]*MessagesUserPropertiesValue) *BatchPeekMessageResponseBodyMessages {
	s.UserProperties = v
	return s
}

func (s *BatchPeekMessageResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}

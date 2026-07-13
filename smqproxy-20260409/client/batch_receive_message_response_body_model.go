// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchReceiveMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessages(v []*BatchReceiveMessageResponseBodyMessages) *BatchReceiveMessageResponseBody
	GetMessages() []*BatchReceiveMessageResponseBodyMessages
}

type BatchReceiveMessageResponseBody struct {
	Messages []*BatchReceiveMessageResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
}

func (s BatchReceiveMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchReceiveMessageResponseBody) GoString() string {
	return s.String()
}

func (s *BatchReceiveMessageResponseBody) GetMessages() []*BatchReceiveMessageResponseBodyMessages {
	return s.Messages
}

func (s *BatchReceiveMessageResponseBody) SetMessages(v []*BatchReceiveMessageResponseBodyMessages) *BatchReceiveMessageResponseBody {
	s.Messages = v
	return s
}

func (s *BatchReceiveMessageResponseBody) Validate() error {
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

type BatchReceiveMessageResponseBodyMessages struct {
	// example:
	//
	// 1
	DequeueCount *int64 `json:"DequeueCount,omitempty" xml:"DequeueCount,omitempty"`
	// example:
	//
	// 1250700979348
	EnqueueTime *int64 `json:"EnqueueTime,omitempty" xml:"EnqueueTime,omitempty"`
	// example:
	//
	// 1250700979348
	FirstDequeueTime *int64 `json:"FirstDequeueTime,omitempty" xml:"FirstDequeueTime,omitempty"`
	// example:
	//
	// This is test message 1.
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	// example:
	//
	// C5DD56A39F5F7BB8B3337C6D11B6D8BE
	MessageBodyMD5 *string `json:"MessageBodyMD5,omitempty" xml:"MessageBodyMD5,omitempty"`
	// example:
	//
	// test-group
	MessageGroupId *string `json:"MessageGroupId,omitempty" xml:"MessageGroupId,omitempty"`
	// example:
	//
	// 5F290C926D472878214D9529A8FA200000001
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 1250700979348
	NextVisibleTime *int64 `json:"NextVisibleTime,omitempty" xml:"NextVisibleTime,omitempty"`
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 1-ODU4OTkzNDU5My0xNDMyNzI3ODI3LTItOA==
	ReceiptHandle *string `json:"ReceiptHandle,omitempty" xml:"ReceiptHandle,omitempty"`
	// example:
	//
	// {"properties1":"value"}
	UserProperties map[string]*MessagesUserPropertiesValue `json:"UserProperties,omitempty" xml:"UserProperties,omitempty"`
}

func (s BatchReceiveMessageResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s BatchReceiveMessageResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *BatchReceiveMessageResponseBodyMessages) GetDequeueCount() *int64 {
	return s.DequeueCount
}

func (s *BatchReceiveMessageResponseBodyMessages) GetEnqueueTime() *int64 {
	return s.EnqueueTime
}

func (s *BatchReceiveMessageResponseBodyMessages) GetFirstDequeueTime() *int64 {
	return s.FirstDequeueTime
}

func (s *BatchReceiveMessageResponseBodyMessages) GetMessageBody() *string {
	return s.MessageBody
}

func (s *BatchReceiveMessageResponseBodyMessages) GetMessageBodyMD5() *string {
	return s.MessageBodyMD5
}

func (s *BatchReceiveMessageResponseBodyMessages) GetMessageGroupId() *string {
	return s.MessageGroupId
}

func (s *BatchReceiveMessageResponseBodyMessages) GetMessageId() *string {
	return s.MessageId
}

func (s *BatchReceiveMessageResponseBodyMessages) GetNextVisibleTime() *int64 {
	return s.NextVisibleTime
}

func (s *BatchReceiveMessageResponseBodyMessages) GetPriority() *int64 {
	return s.Priority
}

func (s *BatchReceiveMessageResponseBodyMessages) GetReceiptHandle() *string {
	return s.ReceiptHandle
}

func (s *BatchReceiveMessageResponseBodyMessages) GetUserProperties() map[string]*MessagesUserPropertiesValue {
	return s.UserProperties
}

func (s *BatchReceiveMessageResponseBodyMessages) SetDequeueCount(v int64) *BatchReceiveMessageResponseBodyMessages {
	s.DequeueCount = &v
	return s
}

func (s *BatchReceiveMessageResponseBodyMessages) SetEnqueueTime(v int64) *BatchReceiveMessageResponseBodyMessages {
	s.EnqueueTime = &v
	return s
}

func (s *BatchReceiveMessageResponseBodyMessages) SetFirstDequeueTime(v int64) *BatchReceiveMessageResponseBodyMessages {
	s.FirstDequeueTime = &v
	return s
}

func (s *BatchReceiveMessageResponseBodyMessages) SetMessageBody(v string) *BatchReceiveMessageResponseBodyMessages {
	s.MessageBody = &v
	return s
}

func (s *BatchReceiveMessageResponseBodyMessages) SetMessageBodyMD5(v string) *BatchReceiveMessageResponseBodyMessages {
	s.MessageBodyMD5 = &v
	return s
}

func (s *BatchReceiveMessageResponseBodyMessages) SetMessageGroupId(v string) *BatchReceiveMessageResponseBodyMessages {
	s.MessageGroupId = &v
	return s
}

func (s *BatchReceiveMessageResponseBodyMessages) SetMessageId(v string) *BatchReceiveMessageResponseBodyMessages {
	s.MessageId = &v
	return s
}

func (s *BatchReceiveMessageResponseBodyMessages) SetNextVisibleTime(v int64) *BatchReceiveMessageResponseBodyMessages {
	s.NextVisibleTime = &v
	return s
}

func (s *BatchReceiveMessageResponseBodyMessages) SetPriority(v int64) *BatchReceiveMessageResponseBodyMessages {
	s.Priority = &v
	return s
}

func (s *BatchReceiveMessageResponseBodyMessages) SetReceiptHandle(v string) *BatchReceiveMessageResponseBodyMessages {
	s.ReceiptHandle = &v
	return s
}

func (s *BatchReceiveMessageResponseBodyMessages) SetUserProperties(v map[string]*MessagesUserPropertiesValue) *BatchReceiveMessageResponseBodyMessages {
	s.UserProperties = v
	return s
}

func (s *BatchReceiveMessageResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}

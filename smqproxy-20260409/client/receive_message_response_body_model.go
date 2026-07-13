// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiveMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDequeueCount(v int64) *ReceiveMessageResponseBody
	GetDequeueCount() *int64
	SetEnqueueTime(v int64) *ReceiveMessageResponseBody
	GetEnqueueTime() *int64
	SetFirstDequeueTime(v int64) *ReceiveMessageResponseBody
	GetFirstDequeueTime() *int64
	SetMessageBody(v string) *ReceiveMessageResponseBody
	GetMessageBody() *string
	SetMessageBodyMD5(v string) *ReceiveMessageResponseBody
	GetMessageBodyMD5() *string
	SetMessageGroupId(v string) *ReceiveMessageResponseBody
	GetMessageGroupId() *string
	SetMessageId(v string) *ReceiveMessageResponseBody
	GetMessageId() *string
	SetNextVisibleTime(v int64) *ReceiveMessageResponseBody
	GetNextVisibleTime() *int64
	SetPriority(v int64) *ReceiveMessageResponseBody
	GetPriority() *int64
	SetReceiptHandle(v string) *ReceiveMessageResponseBody
	GetReceiptHandle() *string
	SetUserProperties(v map[string]*UserPropertiesValue) *ReceiveMessageResponseBody
	GetUserProperties() map[string]*UserPropertiesValue
}

type ReceiveMessageResponseBody struct {
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
	// Hello MNS
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
	// 5F290C926D472878-2-14D9529A8FA-200000001
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
	// 1-ODU4OTkzNDU5My0xNDM1MTk3NjAwLTItNg==
	ReceiptHandle *string `json:"ReceiptHandle,omitempty" xml:"ReceiptHandle,omitempty"`
	// example:
	//
	// {"xxx":"value"}
	UserProperties map[string]*UserPropertiesValue `json:"UserProperties,omitempty" xml:"UserProperties,omitempty"`
}

func (s ReceiveMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReceiveMessageResponseBody) GoString() string {
	return s.String()
}

func (s *ReceiveMessageResponseBody) GetDequeueCount() *int64 {
	return s.DequeueCount
}

func (s *ReceiveMessageResponseBody) GetEnqueueTime() *int64 {
	return s.EnqueueTime
}

func (s *ReceiveMessageResponseBody) GetFirstDequeueTime() *int64 {
	return s.FirstDequeueTime
}

func (s *ReceiveMessageResponseBody) GetMessageBody() *string {
	return s.MessageBody
}

func (s *ReceiveMessageResponseBody) GetMessageBodyMD5() *string {
	return s.MessageBodyMD5
}

func (s *ReceiveMessageResponseBody) GetMessageGroupId() *string {
	return s.MessageGroupId
}

func (s *ReceiveMessageResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *ReceiveMessageResponseBody) GetNextVisibleTime() *int64 {
	return s.NextVisibleTime
}

func (s *ReceiveMessageResponseBody) GetPriority() *int64 {
	return s.Priority
}

func (s *ReceiveMessageResponseBody) GetReceiptHandle() *string {
	return s.ReceiptHandle
}

func (s *ReceiveMessageResponseBody) GetUserProperties() map[string]*UserPropertiesValue {
	return s.UserProperties
}

func (s *ReceiveMessageResponseBody) SetDequeueCount(v int64) *ReceiveMessageResponseBody {
	s.DequeueCount = &v
	return s
}

func (s *ReceiveMessageResponseBody) SetEnqueueTime(v int64) *ReceiveMessageResponseBody {
	s.EnqueueTime = &v
	return s
}

func (s *ReceiveMessageResponseBody) SetFirstDequeueTime(v int64) *ReceiveMessageResponseBody {
	s.FirstDequeueTime = &v
	return s
}

func (s *ReceiveMessageResponseBody) SetMessageBody(v string) *ReceiveMessageResponseBody {
	s.MessageBody = &v
	return s
}

func (s *ReceiveMessageResponseBody) SetMessageBodyMD5(v string) *ReceiveMessageResponseBody {
	s.MessageBodyMD5 = &v
	return s
}

func (s *ReceiveMessageResponseBody) SetMessageGroupId(v string) *ReceiveMessageResponseBody {
	s.MessageGroupId = &v
	return s
}

func (s *ReceiveMessageResponseBody) SetMessageId(v string) *ReceiveMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *ReceiveMessageResponseBody) SetNextVisibleTime(v int64) *ReceiveMessageResponseBody {
	s.NextVisibleTime = &v
	return s
}

func (s *ReceiveMessageResponseBody) SetPriority(v int64) *ReceiveMessageResponseBody {
	s.Priority = &v
	return s
}

func (s *ReceiveMessageResponseBody) SetReceiptHandle(v string) *ReceiveMessageResponseBody {
	s.ReceiptHandle = &v
	return s
}

func (s *ReceiveMessageResponseBody) SetUserProperties(v map[string]*UserPropertiesValue) *ReceiveMessageResponseBody {
	s.UserProperties = v
	return s
}

func (s *ReceiveMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

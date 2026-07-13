// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPeekMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDequeueCount(v int64) *PeekMessageResponseBody
	GetDequeueCount() *int64
	SetEnqueueTime(v int64) *PeekMessageResponseBody
	GetEnqueueTime() *int64
	SetFirstDequeueTime(v int64) *PeekMessageResponseBody
	GetFirstDequeueTime() *int64
	SetMessageBody(v string) *PeekMessageResponseBody
	GetMessageBody() *string
	SetMessageBodyMD5(v string) *PeekMessageResponseBody
	GetMessageBodyMD5() *string
	SetMessageGroupId(v string) *PeekMessageResponseBody
	GetMessageGroupId() *string
	SetMessageId(v string) *PeekMessageResponseBody
	GetMessageId() *string
	SetPriority(v int64) *PeekMessageResponseBody
	GetPriority() *int64
	SetUserProperties(v map[string]*UserPropertiesValue) *PeekMessageResponseBody
	GetUserProperties() map[string]*UserPropertiesValue
}

type PeekMessageResponseBody struct {
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
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// {"xxx":"value"}
	UserProperties map[string]*UserPropertiesValue `json:"UserProperties,omitempty" xml:"UserProperties,omitempty"`
}

func (s PeekMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PeekMessageResponseBody) GoString() string {
	return s.String()
}

func (s *PeekMessageResponseBody) GetDequeueCount() *int64 {
	return s.DequeueCount
}

func (s *PeekMessageResponseBody) GetEnqueueTime() *int64 {
	return s.EnqueueTime
}

func (s *PeekMessageResponseBody) GetFirstDequeueTime() *int64 {
	return s.FirstDequeueTime
}

func (s *PeekMessageResponseBody) GetMessageBody() *string {
	return s.MessageBody
}

func (s *PeekMessageResponseBody) GetMessageBodyMD5() *string {
	return s.MessageBodyMD5
}

func (s *PeekMessageResponseBody) GetMessageGroupId() *string {
	return s.MessageGroupId
}

func (s *PeekMessageResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *PeekMessageResponseBody) GetPriority() *int64 {
	return s.Priority
}

func (s *PeekMessageResponseBody) GetUserProperties() map[string]*UserPropertiesValue {
	return s.UserProperties
}

func (s *PeekMessageResponseBody) SetDequeueCount(v int64) *PeekMessageResponseBody {
	s.DequeueCount = &v
	return s
}

func (s *PeekMessageResponseBody) SetEnqueueTime(v int64) *PeekMessageResponseBody {
	s.EnqueueTime = &v
	return s
}

func (s *PeekMessageResponseBody) SetFirstDequeueTime(v int64) *PeekMessageResponseBody {
	s.FirstDequeueTime = &v
	return s
}

func (s *PeekMessageResponseBody) SetMessageBody(v string) *PeekMessageResponseBody {
	s.MessageBody = &v
	return s
}

func (s *PeekMessageResponseBody) SetMessageBodyMD5(v string) *PeekMessageResponseBody {
	s.MessageBodyMD5 = &v
	return s
}

func (s *PeekMessageResponseBody) SetMessageGroupId(v string) *PeekMessageResponseBody {
	s.MessageGroupId = &v
	return s
}

func (s *PeekMessageResponseBody) SetMessageId(v string) *PeekMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *PeekMessageResponseBody) SetPriority(v int64) *PeekMessageResponseBody {
	s.Priority = &v
	return s
}

func (s *PeekMessageResponseBody) SetUserProperties(v map[string]*UserPropertiesValue) *PeekMessageResponseBody {
	s.UserProperties = v
	return s
}

func (s *PeekMessageResponseBody) Validate() error {
	return dara.Validate(s)
}

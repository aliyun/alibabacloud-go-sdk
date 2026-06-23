// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDelaySeconds(v int32) *SendMessageRequest
	GetDelaySeconds() *int32
	SetMessageBody(v string) *SendMessageRequest
	GetMessageBody() *string
	SetMessageGroupId(v string) *SendMessageRequest
	GetMessageGroupId() *string
	SetPriority(v int32) *SendMessageRequest
	GetPriority() *int32
	SetUserProperties(v string) *SendMessageRequest
	GetUserProperties() *string
}

type SendMessageRequest struct {
	// example:
	//
	// 0
	DelaySeconds *int32 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// example:
	//
	// "Hello MNS"
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	// example:
	//
	// group-123
	MessageGroupId *string `json:"MessageGroupId,omitempty" xml:"MessageGroupId,omitempty"`
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// {"key1":"value1", "key2":"value2"}
	UserProperties *string `json:"UserProperties,omitempty" xml:"UserProperties,omitempty"`
}

func (s SendMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendMessageRequest) GoString() string {
	return s.String()
}

func (s *SendMessageRequest) GetDelaySeconds() *int32 {
	return s.DelaySeconds
}

func (s *SendMessageRequest) GetMessageBody() *string {
	return s.MessageBody
}

func (s *SendMessageRequest) GetMessageGroupId() *string {
	return s.MessageGroupId
}

func (s *SendMessageRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *SendMessageRequest) GetUserProperties() *string {
	return s.UserProperties
}

func (s *SendMessageRequest) SetDelaySeconds(v int32) *SendMessageRequest {
	s.DelaySeconds = &v
	return s
}

func (s *SendMessageRequest) SetMessageBody(v string) *SendMessageRequest {
	s.MessageBody = &v
	return s
}

func (s *SendMessageRequest) SetMessageGroupId(v string) *SendMessageRequest {
	s.MessageGroupId = &v
	return s
}

func (s *SendMessageRequest) SetPriority(v int32) *SendMessageRequest {
	s.Priority = &v
	return s
}

func (s *SendMessageRequest) SetUserProperties(v string) *SendMessageRequest {
	s.UserProperties = &v
	return s
}

func (s *SendMessageRequest) Validate() error {
	return dara.Validate(s)
}

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
	SetUserProperties(v map[string]*UserPropertiesValue) *SendMessageRequest
	GetUserProperties() map[string]*UserPropertiesValue
}

type SendMessageRequest struct {
	// The delay period after which the message can be consumed. Unit: seconds. Maximum value: 7 days.
	//
	// example:
	//
	// 0
	DelaySeconds *int32 `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	// The message body. We recommend that you Base64-encode the message body in advance to avoid issues with special characters.
	//
	// example:
	//
	// "Hello MNS"
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	// The message group ID. Messages within the same group are delivered in order. This parameter is valid only for FIFO queues.
	//
	// example:
	//
	// group-123
	MessageGroupId *string `json:"MessageGroupId,omitempty" xml:"MessageGroupId,omitempty"`
	// The message priority. A smaller value indicates a higher priority. The value 1 indicates the highest priority.
	//
	// example:
	//
	// 2
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The user-defined properties.
	//
	// example:
	//
	// {"key1":"value1", "key2":"value2"}
	UserProperties map[string]*UserPropertiesValue `json:"UserProperties,omitempty" xml:"UserProperties,omitempty"`
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

func (s *SendMessageRequest) GetUserProperties() map[string]*UserPropertiesValue {
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

func (s *SendMessageRequest) SetUserProperties(v map[string]*UserPropertiesValue) *SendMessageRequest {
	s.UserProperties = v
	return s
}

func (s *SendMessageRequest) Validate() error {
	return dara.Validate(s)
}

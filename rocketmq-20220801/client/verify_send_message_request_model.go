// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySendMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryTimeStamp(v int64) *VerifySendMessageRequest
	GetDeliveryTimeStamp() *int64
	SetLiteTopicName(v string) *VerifySendMessageRequest
	GetLiteTopicName() *string
	SetMessage(v string) *VerifySendMessageRequest
	GetMessage() *string
	SetMessageGroup(v string) *VerifySendMessageRequest
	GetMessageGroup() *string
	SetMessageKey(v string) *VerifySendMessageRequest
	GetMessageKey() *string
	SetMessageTag(v string) *VerifySendMessageRequest
	GetMessageTag() *string
	SetUserProperties(v map[string]interface{}) *VerifySendMessageRequest
	GetUserProperties() map[string]interface{}
}

type VerifySendMessageRequest struct {
	// example:
	//
	// 1773718320000
	DeliveryTimeStamp *int64 `json:"deliveryTimeStamp,omitempty" xml:"deliveryTimeStamp,omitempty"`
	// example:
	//
	// abc
	LiteTopicName *string `json:"liteTopicName,omitempty" xml:"liteTopicName,omitempty"`
	// The message body.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// testMessageGroup
	MessageGroup *string `json:"messageGroup,omitempty" xml:"messageGroup,omitempty"`
	// The message key.
	//
	// example:
	//
	// xx
	MessageKey *string `json:"messageKey,omitempty" xml:"messageKey,omitempty"`
	// The message tag.
	//
	// example:
	//
	// xx
	MessageTag     *string                `json:"messageTag,omitempty" xml:"messageTag,omitempty"`
	UserProperties map[string]interface{} `json:"userProperties,omitempty" xml:"userProperties,omitempty"`
}

func (s VerifySendMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifySendMessageRequest) GoString() string {
	return s.String()
}

func (s *VerifySendMessageRequest) GetDeliveryTimeStamp() *int64 {
	return s.DeliveryTimeStamp
}

func (s *VerifySendMessageRequest) GetLiteTopicName() *string {
	return s.LiteTopicName
}

func (s *VerifySendMessageRequest) GetMessage() *string {
	return s.Message
}

func (s *VerifySendMessageRequest) GetMessageGroup() *string {
	return s.MessageGroup
}

func (s *VerifySendMessageRequest) GetMessageKey() *string {
	return s.MessageKey
}

func (s *VerifySendMessageRequest) GetMessageTag() *string {
	return s.MessageTag
}

func (s *VerifySendMessageRequest) GetUserProperties() map[string]interface{} {
	return s.UserProperties
}

func (s *VerifySendMessageRequest) SetDeliveryTimeStamp(v int64) *VerifySendMessageRequest {
	s.DeliveryTimeStamp = &v
	return s
}

func (s *VerifySendMessageRequest) SetLiteTopicName(v string) *VerifySendMessageRequest {
	s.LiteTopicName = &v
	return s
}

func (s *VerifySendMessageRequest) SetMessage(v string) *VerifySendMessageRequest {
	s.Message = &v
	return s
}

func (s *VerifySendMessageRequest) SetMessageGroup(v string) *VerifySendMessageRequest {
	s.MessageGroup = &v
	return s
}

func (s *VerifySendMessageRequest) SetMessageKey(v string) *VerifySendMessageRequest {
	s.MessageKey = &v
	return s
}

func (s *VerifySendMessageRequest) SetMessageTag(v string) *VerifySendMessageRequest {
	s.MessageTag = &v
	return s
}

func (s *VerifySendMessageRequest) SetUserProperties(v map[string]interface{}) *VerifySendMessageRequest {
	s.UserProperties = v
	return s
}

func (s *VerifySendMessageRequest) Validate() error {
	return dara.Validate(s)
}

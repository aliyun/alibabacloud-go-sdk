// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMessageAttributes(v *PublishMessageRequestMessageAttributes) *PublishMessageRequest
	GetMessageAttributes() *PublishMessageRequestMessageAttributes
	SetMessageBody(v string) *PublishMessageRequest
	GetMessageBody() *string
	SetMessageGroupId(v string) *PublishMessageRequest
	GetMessageGroupId() *string
	SetMessageTag(v string) *PublishMessageRequest
	GetMessageTag() *string
}

type PublishMessageRequest struct {
	// **Message attributes**. You do not need to set this parameter when pushing messages to queues or HTTP endpoints.
	MessageAttributes *PublishMessageRequestMessageAttributes `json:"MessageAttributes,omitempty" xml:"MessageAttributes,omitempty" type:"Struct"`
	// **Message content**. Base64-encode the content beforehand.
	//
	// Maximum size: 64 KB.
	//
	// example:
	//
	// hello topic
	MessageBody    *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
	MessageGroupId *string `json:"MessageGroupId,omitempty" xml:"MessageGroupId,omitempty"`
	// **Message tag**, used for message filtering on subscriptions.
	//
	// Maximum length: 64 characters.
	//
	// example:
	//
	// order-event
	MessageTag *string `json:"MessageTag,omitempty" xml:"MessageTag,omitempty"`
}

func (s PublishMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishMessageRequest) GoString() string {
	return s.String()
}

func (s *PublishMessageRequest) GetMessageAttributes() *PublishMessageRequestMessageAttributes {
	return s.MessageAttributes
}

func (s *PublishMessageRequest) GetMessageBody() *string {
	return s.MessageBody
}

func (s *PublishMessageRequest) GetMessageGroupId() *string {
	return s.MessageGroupId
}

func (s *PublishMessageRequest) GetMessageTag() *string {
	return s.MessageTag
}

func (s *PublishMessageRequest) SetMessageAttributes(v *PublishMessageRequestMessageAttributes) *PublishMessageRequest {
	s.MessageAttributes = v
	return s
}

func (s *PublishMessageRequest) SetMessageBody(v string) *PublishMessageRequest {
	s.MessageBody = &v
	return s
}

func (s *PublishMessageRequest) SetMessageGroupId(v string) *PublishMessageRequest {
	s.MessageGroupId = &v
	return s
}

func (s *PublishMessageRequest) SetMessageTag(v string) *PublishMessageRequest {
	s.MessageTag = &v
	return s
}

func (s *PublishMessageRequest) Validate() error {
	if s.MessageAttributes != nil {
		if err := s.MessageAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PublishMessageRequestMessageAttributes struct {
	// **Email push attributes*	- (JSON-formatted string).
	//
	// example:
	//
	// For more information, see https://www.alibabacloud.com/help/en/direct-mail/singlesendmail.
	DirectMail *string `json:"DirectMail,omitempty" xml:"DirectMail,omitempty"`
	// SMS push attributes (JSON-formatted string), including `FreeSignName`, `TemplateCode`, `Type`, `Receiver`, and `SmsParams`.
	//
	// example:
	//
	// {"FreeSignName":"Alibaba Cloud","TemplateCode":"SMS_123456","Type":"singleContent","Receiver":"13800000000","SmsParams":"{\\"code\\":\\"1234\\"}"}
	DirectSMS *string `json:"DirectSMS,omitempty" xml:"DirectSMS,omitempty"`
	// Mobile push attributes (JSON-formatted string).
	//
	// example:
	//
	// Sample mobile push attributes value.
	Push *string `json:"Push,omitempty" xml:"Push,omitempty"`
}

func (s PublishMessageRequestMessageAttributes) String() string {
	return dara.Prettify(s)
}

func (s PublishMessageRequestMessageAttributes) GoString() string {
	return s.String()
}

func (s *PublishMessageRequestMessageAttributes) GetDirectMail() *string {
	return s.DirectMail
}

func (s *PublishMessageRequestMessageAttributes) GetDirectSMS() *string {
	return s.DirectSMS
}

func (s *PublishMessageRequestMessageAttributes) GetPush() *string {
	return s.Push
}

func (s *PublishMessageRequestMessageAttributes) SetDirectMail(v string) *PublishMessageRequestMessageAttributes {
	s.DirectMail = &v
	return s
}

func (s *PublishMessageRequestMessageAttributes) SetDirectSMS(v string) *PublishMessageRequestMessageAttributes {
	s.DirectSMS = &v
	return s
}

func (s *PublishMessageRequestMessageAttributes) SetPush(v string) *PublishMessageRequestMessageAttributes {
	s.Push = &v
	return s
}

func (s *PublishMessageRequestMessageAttributes) Validate() error {
	return dara.Validate(s)
}

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
	SetMessageTag(v string) *PublishMessageRequest
	GetMessageTag() *string
}

type PublishMessageRequest struct {
	MessageAttributes *PublishMessageRequestMessageAttributes `json:"MessageAttributes,omitempty" xml:"MessageAttributes,omitempty" type:"Struct"`
	// example:
	//
	// hello topic
	MessageBody *string `json:"MessageBody,omitempty" xml:"MessageBody,omitempty"`
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
	// example:
	//
	// 详见 https://help.aliyun.com/zh/direct-mail/singlesendmail
	DirectMail *string `json:"DirectMail,omitempty" xml:"DirectMail,omitempty"`
	// example:
	//
	// {"FreeSignName":"阿里云","TemplateCode":"SMS_123456","Type":"singleContent","Receiver":"13800000000","SmsParams":"{\\"code\\":\\"1234\\"}"}
	DirectSMS *string `json:"DirectSMS,omitempty" xml:"DirectSMS,omitempty"`
	// example:
	//
	// 移动推送属性示例值
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

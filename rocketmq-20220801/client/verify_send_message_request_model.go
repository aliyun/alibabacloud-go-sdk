// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySendMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLiteTopicName(v string) *VerifySendMessageRequest
	GetLiteTopicName() *string
	SetMessage(v string) *VerifySendMessageRequest
	GetMessage() *string
	SetMessageKey(v string) *VerifySendMessageRequest
	GetMessageKey() *string
	SetMessageTag(v string) *VerifySendMessageRequest
	GetMessageTag() *string
}

type VerifySendMessageRequest struct {
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
	MessageTag *string `json:"messageTag,omitempty" xml:"messageTag,omitempty"`
}

func (s VerifySendMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s VerifySendMessageRequest) GoString() string {
	return s.String()
}

func (s *VerifySendMessageRequest) GetLiteTopicName() *string {
	return s.LiteTopicName
}

func (s *VerifySendMessageRequest) GetMessage() *string {
	return s.Message
}

func (s *VerifySendMessageRequest) GetMessageKey() *string {
	return s.MessageKey
}

func (s *VerifySendMessageRequest) GetMessageTag() *string {
	return s.MessageTag
}

func (s *VerifySendMessageRequest) SetLiteTopicName(v string) *VerifySendMessageRequest {
	s.LiteTopicName = &v
	return s
}

func (s *VerifySendMessageRequest) SetMessage(v string) *VerifySendMessageRequest {
	s.Message = &v
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

func (s *VerifySendMessageRequest) Validate() error {
	return dara.Validate(s)
}

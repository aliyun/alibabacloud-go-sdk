// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateChatbotInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChatbotInstanceId(v string) *AssociateChatbotInstanceRequest
	GetChatbotInstanceId() *string
	SetChatbotName(v string) *AssociateChatbotInstanceRequest
	GetChatbotName() *string
	SetInstanceId(v string) *AssociateChatbotInstanceRequest
	GetInstanceId() *string
	SetNluServiceParamsJson(v string) *AssociateChatbotInstanceRequest
	GetNluServiceParamsJson() *string
	SetNluServiceType(v string) *AssociateChatbotInstanceRequest
	GetNluServiceType() *string
	SetUnionSource(v string) *AssociateChatbotInstanceRequest
	GetUnionSource() *string
}

type AssociateChatbotInstanceRequest struct {
	// example:
	//
	// chatbot-720edd02b66a
	ChatbotInstanceId *string `json:"ChatbotInstanceId,omitempty" xml:"ChatbotInstanceId,omitempty"`
	ChatbotName       *string `json:"ChatbotName,omitempty" xml:"ChatbotName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NluServiceParamsJson *string `json:"NluServiceParamsJson,omitempty" xml:"NluServiceParamsJson,omitempty"`
	NluServiceType       *string `json:"NluServiceType,omitempty" xml:"NluServiceType,omitempty"`
	UnionSource          *string `json:"UnionSource,omitempty" xml:"UnionSource,omitempty"`
}

func (s AssociateChatbotInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateChatbotInstanceRequest) GoString() string {
	return s.String()
}

func (s *AssociateChatbotInstanceRequest) GetChatbotInstanceId() *string {
	return s.ChatbotInstanceId
}

func (s *AssociateChatbotInstanceRequest) GetChatbotName() *string {
	return s.ChatbotName
}

func (s *AssociateChatbotInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AssociateChatbotInstanceRequest) GetNluServiceParamsJson() *string {
	return s.NluServiceParamsJson
}

func (s *AssociateChatbotInstanceRequest) GetNluServiceType() *string {
	return s.NluServiceType
}

func (s *AssociateChatbotInstanceRequest) GetUnionSource() *string {
	return s.UnionSource
}

func (s *AssociateChatbotInstanceRequest) SetChatbotInstanceId(v string) *AssociateChatbotInstanceRequest {
	s.ChatbotInstanceId = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetChatbotName(v string) *AssociateChatbotInstanceRequest {
	s.ChatbotName = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetInstanceId(v string) *AssociateChatbotInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetNluServiceParamsJson(v string) *AssociateChatbotInstanceRequest {
	s.NluServiceParamsJson = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetNluServiceType(v string) *AssociateChatbotInstanceRequest {
	s.NluServiceType = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) SetUnionSource(v string) *AssociateChatbotInstanceRequest {
	s.UnionSource = &v
	return s
}

func (s *AssociateChatbotInstanceRequest) Validate() error {
	return dara.Validate(s)
}

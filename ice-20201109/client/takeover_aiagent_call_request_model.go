// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTakeoverAIAgentCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHumanAgentUserId(v string) *TakeoverAIAgentCallRequest
	GetHumanAgentUserId() *string
	SetInstanceId(v string) *TakeoverAIAgentCallRequest
	GetInstanceId() *string
	SetRequireToken(v bool) *TakeoverAIAgentCallRequest
	GetRequireToken() *bool
}

type TakeoverAIAgentCallRequest struct {
	// The ID of the human agent that will take over the AI agent (UserId in ARTC). If you do not specify this parameter, it is automatically generated and returned.
	//
	// example:
	//
	// uid2
	HumanAgentUserId *string `json:"HumanAgentUserId,omitempty" xml:"HumanAgentUserId,omitempty"`
	// The ID of the AI agent that will be taken over.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to return the ARTC token. Default value: false.
	//
	// example:
	//
	// false
	RequireToken *bool `json:"RequireToken,omitempty" xml:"RequireToken,omitempty"`
}

func (s TakeoverAIAgentCallRequest) String() string {
	return dara.Prettify(s)
}

func (s TakeoverAIAgentCallRequest) GoString() string {
	return s.String()
}

func (s *TakeoverAIAgentCallRequest) GetHumanAgentUserId() *string {
	return s.HumanAgentUserId
}

func (s *TakeoverAIAgentCallRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TakeoverAIAgentCallRequest) GetRequireToken() *bool {
	return s.RequireToken
}

func (s *TakeoverAIAgentCallRequest) SetHumanAgentUserId(v string) *TakeoverAIAgentCallRequest {
	s.HumanAgentUserId = &v
	return s
}

func (s *TakeoverAIAgentCallRequest) SetInstanceId(v string) *TakeoverAIAgentCallRequest {
	s.InstanceId = &v
	return s
}

func (s *TakeoverAIAgentCallRequest) SetRequireToken(v bool) *TakeoverAIAgentCallRequest {
	s.RequireToken = &v
	return s
}

func (s *TakeoverAIAgentCallRequest) Validate() error {
	return dara.Validate(s)
}

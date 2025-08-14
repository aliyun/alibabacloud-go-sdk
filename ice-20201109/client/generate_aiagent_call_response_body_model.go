// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAIAgentCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentUserId(v string) *GenerateAIAgentCallResponseBody
	GetAIAgentUserId() *string
	SetChannelId(v string) *GenerateAIAgentCallResponseBody
	GetChannelId() *string
	SetInstanceId(v string) *GenerateAIAgentCallResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *GenerateAIAgentCallResponseBody
	GetRequestId() *string
	SetToken(v string) *GenerateAIAgentCallResponseBody
	GetToken() *string
	SetUserId(v string) *GenerateAIAgentCallResponseBody
	GetUserId() *string
}

type GenerateAIAgentCallResponseBody struct {
	// The username of the AI agent in the Alibaba Real-Time Communication (ARTC) channel.
	//
	// example:
	//
	// 877ae632caae49b1afc81c2e8194ffb4
	AIAgentUserId *string `json:"AIAgentUserId,omitempty" xml:"AIAgentUserId,omitempty"`
	// The ARTC channel ID.
	//
	// example:
	//
	// 70f22d5784194938a7e387052f2b3208
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The ID of the AI agent.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ARTC token of the client.
	//
	// example:
	//
	// eyJhcHBpZCI6ICIxMjM0MTIzNxxxxx
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The username in the ARTC channel.
	//
	// example:
	//
	// user123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GenerateAIAgentCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateAIAgentCallResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateAIAgentCallResponseBody) GetAIAgentUserId() *string {
	return s.AIAgentUserId
}

func (s *GenerateAIAgentCallResponseBody) GetChannelId() *string {
	return s.ChannelId
}

func (s *GenerateAIAgentCallResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GenerateAIAgentCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateAIAgentCallResponseBody) GetToken() *string {
	return s.Token
}

func (s *GenerateAIAgentCallResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GenerateAIAgentCallResponseBody) SetAIAgentUserId(v string) *GenerateAIAgentCallResponseBody {
	s.AIAgentUserId = &v
	return s
}

func (s *GenerateAIAgentCallResponseBody) SetChannelId(v string) *GenerateAIAgentCallResponseBody {
	s.ChannelId = &v
	return s
}

func (s *GenerateAIAgentCallResponseBody) SetInstanceId(v string) *GenerateAIAgentCallResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GenerateAIAgentCallResponseBody) SetRequestId(v string) *GenerateAIAgentCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateAIAgentCallResponseBody) SetToken(v string) *GenerateAIAgentCallResponseBody {
	s.Token = &v
	return s
}

func (s *GenerateAIAgentCallResponseBody) SetUserId(v string) *GenerateAIAgentCallResponseBody {
	s.UserId = &v
	return s
}

func (s *GenerateAIAgentCallResponseBody) Validate() error {
	return dara.Validate(s)
}

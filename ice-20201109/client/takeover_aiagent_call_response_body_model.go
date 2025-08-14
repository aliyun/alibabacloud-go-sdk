// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTakeoverAIAgentCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *TakeoverAIAgentCallResponseBody
	GetChannelId() *string
	SetHumanAgentUserId(v string) *TakeoverAIAgentCallResponseBody
	GetHumanAgentUserId() *string
	SetRequestId(v string) *TakeoverAIAgentCallResponseBody
	GetRequestId() *string
	SetToken(v string) *TakeoverAIAgentCallResponseBody
	GetToken() *string
}

type TakeoverAIAgentCallResponseBody struct {
	// The ID of the ARTC channel.
	//
	// example:
	//
	// 70f22d5784194938a7e387052f2b3208
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The ID of the human agent.
	//
	// example:
	//
	// uid2
	HumanAgentUserId *string `json:"HumanAgentUserId,omitempty" xml:"HumanAgentUserId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ARTC token.
	//
	// example:
	//
	// eyJhcHBpZCI6ICIxMjM0MTIzNxxxxx
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s TakeoverAIAgentCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TakeoverAIAgentCallResponseBody) GoString() string {
	return s.String()
}

func (s *TakeoverAIAgentCallResponseBody) GetChannelId() *string {
	return s.ChannelId
}

func (s *TakeoverAIAgentCallResponseBody) GetHumanAgentUserId() *string {
	return s.HumanAgentUserId
}

func (s *TakeoverAIAgentCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TakeoverAIAgentCallResponseBody) GetToken() *string {
	return s.Token
}

func (s *TakeoverAIAgentCallResponseBody) SetChannelId(v string) *TakeoverAIAgentCallResponseBody {
	s.ChannelId = &v
	return s
}

func (s *TakeoverAIAgentCallResponseBody) SetHumanAgentUserId(v string) *TakeoverAIAgentCallResponseBody {
	s.HumanAgentUserId = &v
	return s
}

func (s *TakeoverAIAgentCallResponseBody) SetRequestId(v string) *TakeoverAIAgentCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *TakeoverAIAgentCallResponseBody) SetToken(v string) *TakeoverAIAgentCallResponseBody {
	s.Token = &v
	return s
}

func (s *TakeoverAIAgentCallResponseBody) Validate() error {
	return dara.Validate(s)
}

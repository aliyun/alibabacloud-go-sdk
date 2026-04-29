// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindPolarClawAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *UnbindPolarClawAgentRequest
	GetAgentId() *string
	SetApplicationId(v string) *UnbindPolarClawAgentRequest
	GetApplicationId() *string
	SetChannel(v string) *UnbindPolarClawAgentRequest
	GetChannel() *string
	SetChannelAccountId(v string) *UnbindPolarClawAgentRequest
	GetChannelAccountId() *string
}

type UnbindPolarClawAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// feishu
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// default
	ChannelAccountId *string `json:"ChannelAccountId,omitempty" xml:"ChannelAccountId,omitempty"`
}

func (s UnbindPolarClawAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindPolarClawAgentRequest) GoString() string {
	return s.String()
}

func (s *UnbindPolarClawAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *UnbindPolarClawAgentRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *UnbindPolarClawAgentRequest) GetChannel() *string {
	return s.Channel
}

func (s *UnbindPolarClawAgentRequest) GetChannelAccountId() *string {
	return s.ChannelAccountId
}

func (s *UnbindPolarClawAgentRequest) SetAgentId(v string) *UnbindPolarClawAgentRequest {
	s.AgentId = &v
	return s
}

func (s *UnbindPolarClawAgentRequest) SetApplicationId(v string) *UnbindPolarClawAgentRequest {
	s.ApplicationId = &v
	return s
}

func (s *UnbindPolarClawAgentRequest) SetChannel(v string) *UnbindPolarClawAgentRequest {
	s.Channel = &v
	return s
}

func (s *UnbindPolarClawAgentRequest) SetChannelAccountId(v string) *UnbindPolarClawAgentRequest {
	s.ChannelAccountId = &v
	return s
}

func (s *UnbindPolarClawAgentRequest) Validate() error {
	return dara.Validate(s)
}

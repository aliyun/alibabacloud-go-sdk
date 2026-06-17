// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindPolarClawAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v string) *BindPolarClawAgentRequest
	GetAgentId() *string
	SetApplicationId(v string) *BindPolarClawAgentRequest
	GetApplicationId() *string
	SetChannel(v string) *BindPolarClawAgentRequest
	GetChannel() *string
	SetChannelAccountId(v string) *BindPolarClawAgentRequest
	GetChannelAccountId() *string
}

type BindPolarClawAgentRequest struct {
	// The agent ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// work
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The channel ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// feishu
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The account ID for the channel. If omitted, the default value is \\"default\\".
	//
	// example:
	//
	// default
	ChannelAccountId *string `json:"ChannelAccountId,omitempty" xml:"ChannelAccountId,omitempty"`
}

func (s BindPolarClawAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s BindPolarClawAgentRequest) GoString() string {
	return s.String()
}

func (s *BindPolarClawAgentRequest) GetAgentId() *string {
	return s.AgentId
}

func (s *BindPolarClawAgentRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *BindPolarClawAgentRequest) GetChannel() *string {
	return s.Channel
}

func (s *BindPolarClawAgentRequest) GetChannelAccountId() *string {
	return s.ChannelAccountId
}

func (s *BindPolarClawAgentRequest) SetAgentId(v string) *BindPolarClawAgentRequest {
	s.AgentId = &v
	return s
}

func (s *BindPolarClawAgentRequest) SetApplicationId(v string) *BindPolarClawAgentRequest {
	s.ApplicationId = &v
	return s
}

func (s *BindPolarClawAgentRequest) SetChannel(v string) *BindPolarClawAgentRequest {
	s.Channel = &v
	return s
}

func (s *BindPolarClawAgentRequest) SetChannelAccountId(v string) *BindPolarClawAgentRequest {
	s.ChannelAccountId = &v
	return s
}

func (s *BindPolarClawAgentRequest) Validate() error {
	return dara.Validate(s)
}

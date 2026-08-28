// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentIMChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateAgentIMChannelRequestBody) *UpdateAgentIMChannelRequest
	GetBody() *UpdateAgentIMChannelRequestBody
	SetClientToken(v string) *UpdateAgentIMChannelRequest
	GetClientToken() *string
}

type UpdateAgentIMChannelRequest struct {
	// The request body.
	Body *UpdateAgentIMChannelRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The reserved idempotency token. The backend does not provide persistent idempotency guarantees in this phase.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateAgentIMChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelRequest) GetBody() *UpdateAgentIMChannelRequestBody {
	return s.Body
}

func (s *UpdateAgentIMChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAgentIMChannelRequest) SetBody(v *UpdateAgentIMChannelRequestBody) *UpdateAgentIMChannelRequest {
	s.Body = v
	return s
}

func (s *UpdateAgentIMChannelRequest) SetClientToken(v string) *UpdateAgentIMChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAgentIMChannelRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentIMChannelRequestBody struct {
	// The channel behavior configuration. When provided, the entire configuration is replaced. An empty object restores default values.
	ChannelConfig *UpdateAgentIMChannelRequestBodyChannelConfig `json:"channelConfig,omitempty" xml:"channelConfig,omitempty" type:"Struct"`
	// Specifies whether to enable the IM channel. Default value: true (when created).
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The ID of the bound ServiceEndpoint. The endpoint must belong to the specified agent and its current version, be in the ready state, and have a public endpoint address.
	//
	// example:
	//
	// se-1
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty" xml:"serviceEndpointId,omitempty"`
}

func (s UpdateAgentIMChannelRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelRequestBody) GetChannelConfig() *UpdateAgentIMChannelRequestBodyChannelConfig {
	return s.ChannelConfig
}

func (s *UpdateAgentIMChannelRequestBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateAgentIMChannelRequestBody) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *UpdateAgentIMChannelRequestBody) SetChannelConfig(v *UpdateAgentIMChannelRequestBodyChannelConfig) *UpdateAgentIMChannelRequestBody {
	s.ChannelConfig = v
	return s
}

func (s *UpdateAgentIMChannelRequestBody) SetEnabled(v bool) *UpdateAgentIMChannelRequestBody {
	s.Enabled = &v
	return s
}

func (s *UpdateAgentIMChannelRequestBody) SetServiceEndpointId(v string) *UpdateAgentIMChannelRequestBody {
	s.ServiceEndpointId = &v
	return s
}

func (s *UpdateAgentIMChannelRequestBody) Validate() error {
	if s.ChannelConfig != nil {
		if err := s.ChannelConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentIMChannelRequestBodyChannelConfig struct {
	// Specifies whether to display the thinking process in IM messages. Default value: false.
	ShowThinking *bool `json:"showThinking,omitempty" xml:"showThinking,omitempty"`
	// Specifies whether to display the tool calling process in IM messages. Default value: false.
	ShowToolCalls *bool `json:"showToolCalls,omitempty" xml:"showToolCalls,omitempty"`
}

func (s UpdateAgentIMChannelRequestBodyChannelConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentIMChannelRequestBodyChannelConfig) GoString() string {
	return s.String()
}

func (s *UpdateAgentIMChannelRequestBodyChannelConfig) GetShowThinking() *bool {
	return s.ShowThinking
}

func (s *UpdateAgentIMChannelRequestBodyChannelConfig) GetShowToolCalls() *bool {
	return s.ShowToolCalls
}

func (s *UpdateAgentIMChannelRequestBodyChannelConfig) SetShowThinking(v bool) *UpdateAgentIMChannelRequestBodyChannelConfig {
	s.ShowThinking = &v
	return s
}

func (s *UpdateAgentIMChannelRequestBodyChannelConfig) SetShowToolCalls(v bool) *UpdateAgentIMChannelRequestBodyChannelConfig {
	s.ShowToolCalls = &v
	return s
}

func (s *UpdateAgentIMChannelRequestBodyChannelConfig) Validate() error {
	return dara.Validate(s)
}

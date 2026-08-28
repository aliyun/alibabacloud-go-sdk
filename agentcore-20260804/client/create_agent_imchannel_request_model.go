// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgentIMChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateAgentIMChannelRequestBody) *CreateAgentIMChannelRequest
	GetBody() *CreateAgentIMChannelRequestBody
	SetClientToken(v string) *CreateAgentIMChannelRequest
	GetClientToken() *string
}

type CreateAgentIMChannelRequest struct {
	// The request body.
	Body *CreateAgentIMChannelRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// A reserved idempotency token. The backend does not provide persistent idempotency guarantees in the current phase.
	//
	// example:
	//
	// client-token-1
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateAgentIMChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentIMChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentIMChannelRequest) GetBody() *CreateAgentIMChannelRequestBody {
	return s.Body
}

func (s *CreateAgentIMChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAgentIMChannelRequest) SetBody(v *CreateAgentIMChannelRequestBody) *CreateAgentIMChannelRequest {
	s.Body = v
	return s
}

func (s *CreateAgentIMChannelRequest) SetClientToken(v string) *CreateAgentIMChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAgentIMChannelRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentIMChannelRequestBody struct {
	// The channel behavior configuration.
	ChannelConfig *CreateAgentIMChannelRequestBodyChannelConfig `json:"channelConfig,omitempty" xml:"channelConfig,omitempty" type:"Struct"`
	// The IM channel type. Valid values:
	//
	// - DINGTALK: DingTalk.
	//
	// - FEISHU: Lark.
	//
	// - WECOM: WeCom.
	//
	// This parameter is required.
	//
	// example:
	//
	// DINGTALK
	ChannelType *string `json:"channelType,omitempty" xml:"channelType,omitempty"`
	// The channel credentials. All fields must be provided and field values must be non-empty strings. DingTalk uses clientID and clientSecret. Lark uses appId and appSecret. WeCom uses botId and secret.
	//
	// This parameter is required.
	Credential map[string]*string `json:"credential,omitempty" xml:"credential,omitempty"`
	// Specifies whether to enable the IM channel. Default value: true.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The ID of the ServiceEndpoint to bind. The endpoint must belong to the specified agent and its current version, be in the ready state, and have a public network address.
	//
	// This parameter is required.
	//
	// example:
	//
	// se-1
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty" xml:"serviceEndpointId,omitempty"`
}

func (s CreateAgentIMChannelRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentIMChannelRequestBody) GoString() string {
	return s.String()
}

func (s *CreateAgentIMChannelRequestBody) GetChannelConfig() *CreateAgentIMChannelRequestBodyChannelConfig {
	return s.ChannelConfig
}

func (s *CreateAgentIMChannelRequestBody) GetChannelType() *string {
	return s.ChannelType
}

func (s *CreateAgentIMChannelRequestBody) GetCredential() map[string]*string {
	return s.Credential
}

func (s *CreateAgentIMChannelRequestBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateAgentIMChannelRequestBody) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *CreateAgentIMChannelRequestBody) SetChannelConfig(v *CreateAgentIMChannelRequestBodyChannelConfig) *CreateAgentIMChannelRequestBody {
	s.ChannelConfig = v
	return s
}

func (s *CreateAgentIMChannelRequestBody) SetChannelType(v string) *CreateAgentIMChannelRequestBody {
	s.ChannelType = &v
	return s
}

func (s *CreateAgentIMChannelRequestBody) SetCredential(v map[string]*string) *CreateAgentIMChannelRequestBody {
	s.Credential = v
	return s
}

func (s *CreateAgentIMChannelRequestBody) SetEnabled(v bool) *CreateAgentIMChannelRequestBody {
	s.Enabled = &v
	return s
}

func (s *CreateAgentIMChannelRequestBody) SetServiceEndpointId(v string) *CreateAgentIMChannelRequestBody {
	s.ServiceEndpointId = &v
	return s
}

func (s *CreateAgentIMChannelRequestBody) Validate() error {
	if s.ChannelConfig != nil {
		if err := s.ChannelConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgentIMChannelRequestBodyChannelConfig struct {
	// Specifies whether to display the thinking process in IM messages. Default value: false.
	ShowThinking *bool `json:"showThinking,omitempty" xml:"showThinking,omitempty"`
	// Specifies whether to display the tool calling process in IM messages. Default value: false.
	ShowToolCalls *bool `json:"showToolCalls,omitempty" xml:"showToolCalls,omitempty"`
}

func (s CreateAgentIMChannelRequestBodyChannelConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateAgentIMChannelRequestBodyChannelConfig) GoString() string {
	return s.String()
}

func (s *CreateAgentIMChannelRequestBodyChannelConfig) GetShowThinking() *bool {
	return s.ShowThinking
}

func (s *CreateAgentIMChannelRequestBodyChannelConfig) GetShowToolCalls() *bool {
	return s.ShowToolCalls
}

func (s *CreateAgentIMChannelRequestBodyChannelConfig) SetShowThinking(v bool) *CreateAgentIMChannelRequestBodyChannelConfig {
	s.ShowThinking = &v
	return s
}

func (s *CreateAgentIMChannelRequestBodyChannelConfig) SetShowToolCalls(v bool) *CreateAgentIMChannelRequestBodyChannelConfig {
	s.ShowToolCalls = &v
	return s
}

func (s *CreateAgentIMChannelRequestBodyChannelConfig) Validate() error {
	return dara.Validate(s)
}

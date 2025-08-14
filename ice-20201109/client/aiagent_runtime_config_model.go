// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAIAgentRuntimeConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAgentUserId(v string) *AIAgentRuntimeConfig
	GetAgentUserId() *string
	SetAuthToken(v string) *AIAgentRuntimeConfig
	GetAuthToken() *string
	SetAvatarChat3D(v *AIAgentRuntimeConfigAvatarChat3D) *AIAgentRuntimeConfig
	GetAvatarChat3D() *AIAgentRuntimeConfigAvatarChat3D
	SetChannelId(v string) *AIAgentRuntimeConfig
	GetChannelId() *string
	SetVisionChat(v *AIAgentRuntimeConfigVisionChat) *AIAgentRuntimeConfig
	GetVisionChat() *AIAgentRuntimeConfigVisionChat
	SetVoiceChat(v *AIAgentRuntimeConfigVoiceChat) *AIAgentRuntimeConfig
	GetVoiceChat() *AIAgentRuntimeConfigVoiceChat
}

type AIAgentRuntimeConfig struct {
	AgentUserId *string `json:"AgentUserId,omitempty" xml:"AgentUserId,omitempty"`
	AuthToken   *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	// Deprecated
	AvatarChat3D *AIAgentRuntimeConfigAvatarChat3D `json:"AvatarChat3D,omitempty" xml:"AvatarChat3D,omitempty" type:"Struct"`
	ChannelId    *string                           `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// Deprecated
	VisionChat *AIAgentRuntimeConfigVisionChat `json:"VisionChat,omitempty" xml:"VisionChat,omitempty" type:"Struct"`
	// Deprecated
	VoiceChat *AIAgentRuntimeConfigVoiceChat `json:"VoiceChat,omitempty" xml:"VoiceChat,omitempty" type:"Struct"`
}

func (s AIAgentRuntimeConfig) String() string {
	return dara.Prettify(s)
}

func (s AIAgentRuntimeConfig) GoString() string {
	return s.String()
}

func (s *AIAgentRuntimeConfig) GetAgentUserId() *string {
	return s.AgentUserId
}

func (s *AIAgentRuntimeConfig) GetAuthToken() *string {
	return s.AuthToken
}

func (s *AIAgentRuntimeConfig) GetAvatarChat3D() *AIAgentRuntimeConfigAvatarChat3D {
	return s.AvatarChat3D
}

func (s *AIAgentRuntimeConfig) GetChannelId() *string {
	return s.ChannelId
}

func (s *AIAgentRuntimeConfig) GetVisionChat() *AIAgentRuntimeConfigVisionChat {
	return s.VisionChat
}

func (s *AIAgentRuntimeConfig) GetVoiceChat() *AIAgentRuntimeConfigVoiceChat {
	return s.VoiceChat
}

func (s *AIAgentRuntimeConfig) SetAgentUserId(v string) *AIAgentRuntimeConfig {
	s.AgentUserId = &v
	return s
}

func (s *AIAgentRuntimeConfig) SetAuthToken(v string) *AIAgentRuntimeConfig {
	s.AuthToken = &v
	return s
}

func (s *AIAgentRuntimeConfig) SetAvatarChat3D(v *AIAgentRuntimeConfigAvatarChat3D) *AIAgentRuntimeConfig {
	s.AvatarChat3D = v
	return s
}

func (s *AIAgentRuntimeConfig) SetChannelId(v string) *AIAgentRuntimeConfig {
	s.ChannelId = &v
	return s
}

func (s *AIAgentRuntimeConfig) SetVisionChat(v *AIAgentRuntimeConfigVisionChat) *AIAgentRuntimeConfig {
	s.VisionChat = v
	return s
}

func (s *AIAgentRuntimeConfig) SetVoiceChat(v *AIAgentRuntimeConfigVoiceChat) *AIAgentRuntimeConfig {
	s.VoiceChat = v
	return s
}

func (s *AIAgentRuntimeConfig) Validate() error {
	return dara.Validate(s)
}

type AIAgentRuntimeConfigAvatarChat3D struct {
	AgentUserId *string `json:"AgentUserId,omitempty" xml:"AgentUserId,omitempty"`
	AuthToken   *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s AIAgentRuntimeConfigAvatarChat3D) String() string {
	return dara.Prettify(s)
}

func (s AIAgentRuntimeConfigAvatarChat3D) GoString() string {
	return s.String()
}

func (s *AIAgentRuntimeConfigAvatarChat3D) GetAgentUserId() *string {
	return s.AgentUserId
}

func (s *AIAgentRuntimeConfigAvatarChat3D) GetAuthToken() *string {
	return s.AuthToken
}

func (s *AIAgentRuntimeConfigAvatarChat3D) GetChannelId() *string {
	return s.ChannelId
}

func (s *AIAgentRuntimeConfigAvatarChat3D) SetAgentUserId(v string) *AIAgentRuntimeConfigAvatarChat3D {
	s.AgentUserId = &v
	return s
}

func (s *AIAgentRuntimeConfigAvatarChat3D) SetAuthToken(v string) *AIAgentRuntimeConfigAvatarChat3D {
	s.AuthToken = &v
	return s
}

func (s *AIAgentRuntimeConfigAvatarChat3D) SetChannelId(v string) *AIAgentRuntimeConfigAvatarChat3D {
	s.ChannelId = &v
	return s
}

func (s *AIAgentRuntimeConfigAvatarChat3D) Validate() error {
	return dara.Validate(s)
}

type AIAgentRuntimeConfigVisionChat struct {
	AgentUserId *string `json:"AgentUserId,omitempty" xml:"AgentUserId,omitempty"`
	AuthToken   *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s AIAgentRuntimeConfigVisionChat) String() string {
	return dara.Prettify(s)
}

func (s AIAgentRuntimeConfigVisionChat) GoString() string {
	return s.String()
}

func (s *AIAgentRuntimeConfigVisionChat) GetAgentUserId() *string {
	return s.AgentUserId
}

func (s *AIAgentRuntimeConfigVisionChat) GetAuthToken() *string {
	return s.AuthToken
}

func (s *AIAgentRuntimeConfigVisionChat) GetChannelId() *string {
	return s.ChannelId
}

func (s *AIAgentRuntimeConfigVisionChat) SetAgentUserId(v string) *AIAgentRuntimeConfigVisionChat {
	s.AgentUserId = &v
	return s
}

func (s *AIAgentRuntimeConfigVisionChat) SetAuthToken(v string) *AIAgentRuntimeConfigVisionChat {
	s.AuthToken = &v
	return s
}

func (s *AIAgentRuntimeConfigVisionChat) SetChannelId(v string) *AIAgentRuntimeConfigVisionChat {
	s.ChannelId = &v
	return s
}

func (s *AIAgentRuntimeConfigVisionChat) Validate() error {
	return dara.Validate(s)
}

type AIAgentRuntimeConfigVoiceChat struct {
	AgentUserId *string `json:"AgentUserId,omitempty" xml:"AgentUserId,omitempty"`
	AuthToken   *string `json:"AuthToken,omitempty" xml:"AuthToken,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s AIAgentRuntimeConfigVoiceChat) String() string {
	return dara.Prettify(s)
}

func (s AIAgentRuntimeConfigVoiceChat) GoString() string {
	return s.String()
}

func (s *AIAgentRuntimeConfigVoiceChat) GetAgentUserId() *string {
	return s.AgentUserId
}

func (s *AIAgentRuntimeConfigVoiceChat) GetAuthToken() *string {
	return s.AuthToken
}

func (s *AIAgentRuntimeConfigVoiceChat) GetChannelId() *string {
	return s.ChannelId
}

func (s *AIAgentRuntimeConfigVoiceChat) SetAgentUserId(v string) *AIAgentRuntimeConfigVoiceChat {
	s.AgentUserId = &v
	return s
}

func (s *AIAgentRuntimeConfigVoiceChat) SetAuthToken(v string) *AIAgentRuntimeConfigVoiceChat {
	s.AuthToken = &v
	return s
}

func (s *AIAgentRuntimeConfigVoiceChat) SetChannelId(v string) *AIAgentRuntimeConfigVoiceChat {
	s.ChannelId = &v
	return s
}

func (s *AIAgentRuntimeConfigVoiceChat) Validate() error {
	return dara.Validate(s)
}

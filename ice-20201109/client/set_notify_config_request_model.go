// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetNotifyConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *SetNotifyConfigRequest
	GetAIAgentId() *string
	SetAudioOssPath(v string) *SetNotifyConfigRequest
	GetAudioOssPath() *string
	SetCallbackUrl(v string) *SetNotifyConfigRequest
	GetCallbackUrl() *string
	SetEnableAudioRecording(v bool) *SetNotifyConfigRequest
	GetEnableAudioRecording() *bool
	SetEnableNotify(v bool) *SetNotifyConfigRequest
	GetEnableNotify() *bool
	SetEventTypes(v string) *SetNotifyConfigRequest
	GetEventTypes() *string
	SetToken(v string) *SetNotifyConfigRequest
	GetToken() *string
}

type SetNotifyConfigRequest struct {
	// The ID of the AI agent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	AIAgentId    *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
	AudioOssPath *string `json:"AudioOssPath,omitempty" xml:"AudioOssPath,omitempty"`
	// The URL for receiving callback notifications. By default, this parameter is left empty.
	//
	// example:
	//
	// http://customer.com/callback
	CallbackUrl          *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	EnableAudioRecording *bool   `json:"EnableAudioRecording,omitempty" xml:"EnableAudioRecording,omitempty"`
	// Specifies whether to enable event notifications.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableNotify *bool `json:"EnableNotify,omitempty" xml:"EnableNotify,omitempty"`
	// The event types. If you do not specify this parameter, all event types are selected.
	//
	// 	- agent_start
	//
	// 	- agent_stop
	//
	// 	- error
	//
	// example:
	//
	// agent_start,agent_stop,error
	EventTypes *string `json:"EventTypes,omitempty" xml:"EventTypes,omitempty"`
	// The authentication token for callback. The token is carried in the Authorization header of a callback request. By default, this parameter is left empty.
	//
	// example:
	//
	// eyJhcHBpZCI6ICIxMjM0MTIzNxxxxx
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s SetNotifyConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetNotifyConfigRequest) GoString() string {
	return s.String()
}

func (s *SetNotifyConfigRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *SetNotifyConfigRequest) GetAudioOssPath() *string {
	return s.AudioOssPath
}

func (s *SetNotifyConfigRequest) GetCallbackUrl() *string {
	return s.CallbackUrl
}

func (s *SetNotifyConfigRequest) GetEnableAudioRecording() *bool {
	return s.EnableAudioRecording
}

func (s *SetNotifyConfigRequest) GetEnableNotify() *bool {
	return s.EnableNotify
}

func (s *SetNotifyConfigRequest) GetEventTypes() *string {
	return s.EventTypes
}

func (s *SetNotifyConfigRequest) GetToken() *string {
	return s.Token
}

func (s *SetNotifyConfigRequest) SetAIAgentId(v string) *SetNotifyConfigRequest {
	s.AIAgentId = &v
	return s
}

func (s *SetNotifyConfigRequest) SetAudioOssPath(v string) *SetNotifyConfigRequest {
	s.AudioOssPath = &v
	return s
}

func (s *SetNotifyConfigRequest) SetCallbackUrl(v string) *SetNotifyConfigRequest {
	s.CallbackUrl = &v
	return s
}

func (s *SetNotifyConfigRequest) SetEnableAudioRecording(v bool) *SetNotifyConfigRequest {
	s.EnableAudioRecording = &v
	return s
}

func (s *SetNotifyConfigRequest) SetEnableNotify(v bool) *SetNotifyConfigRequest {
	s.EnableNotify = &v
	return s
}

func (s *SetNotifyConfigRequest) SetEventTypes(v string) *SetNotifyConfigRequest {
	s.EventTypes = &v
	return s
}

func (s *SetNotifyConfigRequest) SetToken(v string) *SetNotifyConfigRequest {
	s.Token = &v
	return s
}

func (s *SetNotifyConfigRequest) Validate() error {
	return dara.Validate(s)
}

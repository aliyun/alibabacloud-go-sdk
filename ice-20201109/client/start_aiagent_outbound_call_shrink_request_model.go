// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAgentOutboundCallShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *StartAIAgentOutboundCallShrinkRequest
	GetAIAgentId() *string
	SetCalledNumber(v string) *StartAIAgentOutboundCallShrinkRequest
	GetCalledNumber() *string
	SetCallerNumber(v string) *StartAIAgentOutboundCallShrinkRequest
	GetCallerNumber() *string
	SetConfigShrink(v string) *StartAIAgentOutboundCallShrinkRequest
	GetConfigShrink() *string
	SetImsAIAgentFreeObCall(v string) *StartAIAgentOutboundCallShrinkRequest
	GetImsAIAgentFreeObCall() *string
	SetSessionId(v string) *StartAIAgentOutboundCallShrinkRequest
	GetSessionId() *string
	SetUserData(v string) *StartAIAgentOutboundCallShrinkRequest
	GetUserData() *string
}

type StartAIAgentOutboundCallShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ***********e4f309379701645f4****
	AIAgentId *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 173*****533
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 183*****333
	CallerNumber         *string `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	ConfigShrink         *string `json:"Config,omitempty" xml:"Config,omitempty"`
	ImsAIAgentFreeObCall *string `json:"ImsAIAgentFreeObCall,omitempty" xml:"ImsAIAgentFreeObCall,omitempty"`
	// example:
	//
	// f213fbc005e4f309379701645f4****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s StartAIAgentOutboundCallShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAIAgentOutboundCallShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartAIAgentOutboundCallShrinkRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *StartAIAgentOutboundCallShrinkRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *StartAIAgentOutboundCallShrinkRequest) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *StartAIAgentOutboundCallShrinkRequest) GetConfigShrink() *string {
	return s.ConfigShrink
}

func (s *StartAIAgentOutboundCallShrinkRequest) GetImsAIAgentFreeObCall() *string {
	return s.ImsAIAgentFreeObCall
}

func (s *StartAIAgentOutboundCallShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StartAIAgentOutboundCallShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *StartAIAgentOutboundCallShrinkRequest) SetAIAgentId(v string) *StartAIAgentOutboundCallShrinkRequest {
	s.AIAgentId = &v
	return s
}

func (s *StartAIAgentOutboundCallShrinkRequest) SetCalledNumber(v string) *StartAIAgentOutboundCallShrinkRequest {
	s.CalledNumber = &v
	return s
}

func (s *StartAIAgentOutboundCallShrinkRequest) SetCallerNumber(v string) *StartAIAgentOutboundCallShrinkRequest {
	s.CallerNumber = &v
	return s
}

func (s *StartAIAgentOutboundCallShrinkRequest) SetConfigShrink(v string) *StartAIAgentOutboundCallShrinkRequest {
	s.ConfigShrink = &v
	return s
}

func (s *StartAIAgentOutboundCallShrinkRequest) SetImsAIAgentFreeObCall(v string) *StartAIAgentOutboundCallShrinkRequest {
	s.ImsAIAgentFreeObCall = &v
	return s
}

func (s *StartAIAgentOutboundCallShrinkRequest) SetSessionId(v string) *StartAIAgentOutboundCallShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *StartAIAgentOutboundCallShrinkRequest) SetUserData(v string) *StartAIAgentOutboundCallShrinkRequest {
	s.UserData = &v
	return s
}

func (s *StartAIAgentOutboundCallShrinkRequest) Validate() error {
	return dara.Validate(s)
}

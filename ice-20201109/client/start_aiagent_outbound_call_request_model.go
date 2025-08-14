// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAIAgentOutboundCallRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *StartAIAgentOutboundCallRequest
	GetAIAgentId() *string
	SetCalledNumber(v string) *StartAIAgentOutboundCallRequest
	GetCalledNumber() *string
	SetCallerNumber(v string) *StartAIAgentOutboundCallRequest
	GetCallerNumber() *string
	SetConfig(v *AIAgentOutboundCallConfig) *StartAIAgentOutboundCallRequest
	GetConfig() *AIAgentOutboundCallConfig
	SetImsAIAgentFreeObCall(v string) *StartAIAgentOutboundCallRequest
	GetImsAIAgentFreeObCall() *string
	SetSessionId(v string) *StartAIAgentOutboundCallRequest
	GetSessionId() *string
	SetUserData(v string) *StartAIAgentOutboundCallRequest
	GetUserData() *string
}

type StartAIAgentOutboundCallRequest struct {
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
	CallerNumber         *string                    `json:"CallerNumber,omitempty" xml:"CallerNumber,omitempty"`
	Config               *AIAgentOutboundCallConfig `json:"Config,omitempty" xml:"Config,omitempty"`
	ImsAIAgentFreeObCall *string                    `json:"ImsAIAgentFreeObCall,omitempty" xml:"ImsAIAgentFreeObCall,omitempty"`
	// example:
	//
	// f213fbc005e4f309379701645f4****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s StartAIAgentOutboundCallRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAIAgentOutboundCallRequest) GoString() string {
	return s.String()
}

func (s *StartAIAgentOutboundCallRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *StartAIAgentOutboundCallRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *StartAIAgentOutboundCallRequest) GetCallerNumber() *string {
	return s.CallerNumber
}

func (s *StartAIAgentOutboundCallRequest) GetConfig() *AIAgentOutboundCallConfig {
	return s.Config
}

func (s *StartAIAgentOutboundCallRequest) GetImsAIAgentFreeObCall() *string {
	return s.ImsAIAgentFreeObCall
}

func (s *StartAIAgentOutboundCallRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *StartAIAgentOutboundCallRequest) GetUserData() *string {
	return s.UserData
}

func (s *StartAIAgentOutboundCallRequest) SetAIAgentId(v string) *StartAIAgentOutboundCallRequest {
	s.AIAgentId = &v
	return s
}

func (s *StartAIAgentOutboundCallRequest) SetCalledNumber(v string) *StartAIAgentOutboundCallRequest {
	s.CalledNumber = &v
	return s
}

func (s *StartAIAgentOutboundCallRequest) SetCallerNumber(v string) *StartAIAgentOutboundCallRequest {
	s.CallerNumber = &v
	return s
}

func (s *StartAIAgentOutboundCallRequest) SetConfig(v *AIAgentOutboundCallConfig) *StartAIAgentOutboundCallRequest {
	s.Config = v
	return s
}

func (s *StartAIAgentOutboundCallRequest) SetImsAIAgentFreeObCall(v string) *StartAIAgentOutboundCallRequest {
	s.ImsAIAgentFreeObCall = &v
	return s
}

func (s *StartAIAgentOutboundCallRequest) SetSessionId(v string) *StartAIAgentOutboundCallRequest {
	s.SessionId = &v
	return s
}

func (s *StartAIAgentOutboundCallRequest) SetUserData(v string) *StartAIAgentOutboundCallRequest {
	s.UserData = &v
	return s
}

func (s *StartAIAgentOutboundCallRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIAgentVideoAuditTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *SubmitAIAgentVideoAuditTaskShrinkRequest
	GetAIAgentId() *string
	SetAuditInterval(v int32) *SubmitAIAgentVideoAuditTaskShrinkRequest
	GetAuditInterval() *int32
	SetCallbackConfigShrink(v string) *SubmitAIAgentVideoAuditTaskShrinkRequest
	GetCallbackConfigShrink() *string
	SetCapturePoliciesShrink(v string) *SubmitAIAgentVideoAuditTaskShrinkRequest
	GetCapturePoliciesShrink() *string
	SetInputShrink(v string) *SubmitAIAgentVideoAuditTaskShrinkRequest
	GetInputShrink() *string
	SetUserData(v string) *SubmitAIAgentVideoAuditTaskShrinkRequest
	GetUserData() *string
}

type SubmitAIAgentVideoAuditTaskShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// *******3b3d94abda22******
	AIAgentId *string `json:"AIAgentId,omitempty" xml:"AIAgentId,omitempty"`
	// example:
	//
	// 3000
	AuditInterval *int32 `json:"AuditInterval,omitempty" xml:"AuditInterval,omitempty"`
	// example:
	//
	// {"Url":"https://yourcallback","Token":"yourtoken"}
	CallbackConfigShrink *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
	// This parameter is required.
	CapturePoliciesShrink *string `json:"CapturePolicies,omitempty" xml:"CapturePolicies,omitempty"`
	// This parameter is required.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIAgentVideoAuditTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIAgentVideoAuditTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) GetAuditInterval() *int32 {
	return s.AuditInterval
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) GetCallbackConfigShrink() *string {
	return s.CallbackConfigShrink
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) GetCapturePoliciesShrink() *string {
	return s.CapturePoliciesShrink
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) SetAIAgentId(v string) *SubmitAIAgentVideoAuditTaskShrinkRequest {
	s.AIAgentId = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) SetAuditInterval(v int32) *SubmitAIAgentVideoAuditTaskShrinkRequest {
	s.AuditInterval = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) SetCallbackConfigShrink(v string) *SubmitAIAgentVideoAuditTaskShrinkRequest {
	s.CallbackConfigShrink = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) SetCapturePoliciesShrink(v string) *SubmitAIAgentVideoAuditTaskShrinkRequest {
	s.CapturePoliciesShrink = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) SetInputShrink(v string) *SubmitAIAgentVideoAuditTaskShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) SetUserData(v string) *SubmitAIAgentVideoAuditTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

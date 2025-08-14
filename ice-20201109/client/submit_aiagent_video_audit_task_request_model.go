// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAIAgentVideoAuditTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAIAgentId(v string) *SubmitAIAgentVideoAuditTaskRequest
	GetAIAgentId() *string
	SetAuditInterval(v int32) *SubmitAIAgentVideoAuditTaskRequest
	GetAuditInterval() *int32
	SetCallbackConfig(v *SubmitAIAgentVideoAuditTaskRequestCallbackConfig) *SubmitAIAgentVideoAuditTaskRequest
	GetCallbackConfig() *SubmitAIAgentVideoAuditTaskRequestCallbackConfig
	SetCapturePolicies(v []*SubmitAIAgentVideoAuditTaskRequestCapturePolicies) *SubmitAIAgentVideoAuditTaskRequest
	GetCapturePolicies() []*SubmitAIAgentVideoAuditTaskRequestCapturePolicies
	SetInput(v *SubmitAIAgentVideoAuditTaskRequestInput) *SubmitAIAgentVideoAuditTaskRequest
	GetInput() *SubmitAIAgentVideoAuditTaskRequestInput
	SetUserData(v string) *SubmitAIAgentVideoAuditTaskRequest
	GetUserData() *string
}

type SubmitAIAgentVideoAuditTaskRequest struct {
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
	CallbackConfig *SubmitAIAgentVideoAuditTaskRequestCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
	// This parameter is required.
	CapturePolicies []*SubmitAIAgentVideoAuditTaskRequestCapturePolicies `json:"CapturePolicies,omitempty" xml:"CapturePolicies,omitempty" type:"Repeated"`
	// This parameter is required.
	Input    *SubmitAIAgentVideoAuditTaskRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	UserData *string                                  `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitAIAgentVideoAuditTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIAgentVideoAuditTaskRequest) GoString() string {
	return s.String()
}

func (s *SubmitAIAgentVideoAuditTaskRequest) GetAIAgentId() *string {
	return s.AIAgentId
}

func (s *SubmitAIAgentVideoAuditTaskRequest) GetAuditInterval() *int32 {
	return s.AuditInterval
}

func (s *SubmitAIAgentVideoAuditTaskRequest) GetCallbackConfig() *SubmitAIAgentVideoAuditTaskRequestCallbackConfig {
	return s.CallbackConfig
}

func (s *SubmitAIAgentVideoAuditTaskRequest) GetCapturePolicies() []*SubmitAIAgentVideoAuditTaskRequestCapturePolicies {
	return s.CapturePolicies
}

func (s *SubmitAIAgentVideoAuditTaskRequest) GetInput() *SubmitAIAgentVideoAuditTaskRequestInput {
	return s.Input
}

func (s *SubmitAIAgentVideoAuditTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitAIAgentVideoAuditTaskRequest) SetAIAgentId(v string) *SubmitAIAgentVideoAuditTaskRequest {
	s.AIAgentId = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequest) SetAuditInterval(v int32) *SubmitAIAgentVideoAuditTaskRequest {
	s.AuditInterval = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequest) SetCallbackConfig(v *SubmitAIAgentVideoAuditTaskRequestCallbackConfig) *SubmitAIAgentVideoAuditTaskRequest {
	s.CallbackConfig = v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequest) SetCapturePolicies(v []*SubmitAIAgentVideoAuditTaskRequestCapturePolicies) *SubmitAIAgentVideoAuditTaskRequest {
	s.CapturePolicies = v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequest) SetInput(v *SubmitAIAgentVideoAuditTaskRequestInput) *SubmitAIAgentVideoAuditTaskRequest {
	s.Input = v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequest) SetUserData(v string) *SubmitAIAgentVideoAuditTaskRequest {
	s.UserData = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequest) Validate() error {
	return dara.Validate(s)
}

type SubmitAIAgentVideoAuditTaskRequestCallbackConfig struct {
	// example:
	//
	// Bearer Token
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// https://yourcallback
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s SubmitAIAgentVideoAuditTaskRequestCallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIAgentVideoAuditTaskRequestCallbackConfig) GoString() string {
	return s.String()
}

func (s *SubmitAIAgentVideoAuditTaskRequestCallbackConfig) GetToken() *string {
	return s.Token
}

func (s *SubmitAIAgentVideoAuditTaskRequestCallbackConfig) GetUrl() *string {
	return s.Url
}

func (s *SubmitAIAgentVideoAuditTaskRequestCallbackConfig) SetToken(v string) *SubmitAIAgentVideoAuditTaskRequestCallbackConfig {
	s.Token = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequestCallbackConfig) SetUrl(v string) *SubmitAIAgentVideoAuditTaskRequestCallbackConfig {
	s.Url = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequestCallbackConfig) Validate() error {
	return dara.Validate(s)
}

type SubmitAIAgentVideoAuditTaskRequestCapturePolicies struct {
	// example:
	//
	// 10
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2
	FrameCount *int32  `json:"FrameCount,omitempty" xml:"FrameCount,omitempty"`
	Prompt     *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 0
	StartTime *int32 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SubmitAIAgentVideoAuditTaskRequestCapturePolicies) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIAgentVideoAuditTaskRequestCapturePolicies) GoString() string {
	return s.String()
}

func (s *SubmitAIAgentVideoAuditTaskRequestCapturePolicies) GetDuration() *int32 {
	return s.Duration
}

func (s *SubmitAIAgentVideoAuditTaskRequestCapturePolicies) GetFrameCount() *int32 {
	return s.FrameCount
}

func (s *SubmitAIAgentVideoAuditTaskRequestCapturePolicies) GetPrompt() *string {
	return s.Prompt
}

func (s *SubmitAIAgentVideoAuditTaskRequestCapturePolicies) GetStartTime() *int32 {
	return s.StartTime
}

func (s *SubmitAIAgentVideoAuditTaskRequestCapturePolicies) SetDuration(v int32) *SubmitAIAgentVideoAuditTaskRequestCapturePolicies {
	s.Duration = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequestCapturePolicies) SetFrameCount(v int32) *SubmitAIAgentVideoAuditTaskRequestCapturePolicies {
	s.FrameCount = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequestCapturePolicies) SetPrompt(v string) *SubmitAIAgentVideoAuditTaskRequestCapturePolicies {
	s.Prompt = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequestCapturePolicies) SetStartTime(v int32) *SubmitAIAgentVideoAuditTaskRequestCapturePolicies {
	s.StartTime = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequestCapturePolicies) Validate() error {
	return dara.Validate(s)
}

type SubmitAIAgentVideoAuditTaskRequestInput struct {
	// example:
	//
	// http://my-bucket.cn-shanghai.aliyuncs.com/object-id.mp4
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// example:
	//
	// OSS
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitAIAgentVideoAuditTaskRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitAIAgentVideoAuditTaskRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitAIAgentVideoAuditTaskRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitAIAgentVideoAuditTaskRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitAIAgentVideoAuditTaskRequestInput) SetMedia(v string) *SubmitAIAgentVideoAuditTaskRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequestInput) SetType(v string) *SubmitAIAgentVideoAuditTaskRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitAIAgentVideoAuditTaskRequestInput) Validate() error {
	return dara.Validate(s)
}

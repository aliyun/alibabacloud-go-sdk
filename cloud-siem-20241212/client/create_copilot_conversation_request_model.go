// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCopilotConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateCopilotConversationRequest
	GetLang() *string
	SetModel(v string) *CreateCopilotConversationRequest
	GetModel() *string
	SetPlanMode(v string) *CreateCopilotConversationRequest
	GetPlanMode() *string
	SetProjectId(v string) *CreateCopilotConversationRequest
	GetProjectId() *string
	SetRegionId(v string) *CreateCopilotConversationRequest
	GetRegionId() *string
	SetThinkingMode(v string) *CreateCopilotConversationRequest
	GetThinkingMode() *string
	SetTitle(v string) *CreateCopilotConversationRequest
	GetTitle() *string
	SetTraceId(v string) *CreateCopilotConversationRequest
	GetTraceId() *string
}

type CreateCopilotConversationRequest struct {
	// The language of the response messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The model identifier.
	//
	// example:
	//
	// qwen3.7-max
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The planning configuration.
	//
	// example:
	//
	// enabled
	PlanMode *string `json:"PlanMode,omitempty" xml:"PlanMode,omitempty"`
	// The security project ID.
	//
	// example:
	//
	// project-incident-****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region where the data management center for threat analysis is located. Select the management center region based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Assets are located in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are located outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The thinking mode configuration.
	//
	// example:
	//
	// enabled
	ThinkingMode *string `json:"ThinkingMode,omitempty" xml:"ThinkingMode,omitempty"`
	// The session title.
	//
	// example:
	//
	// Security incident investigation
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The trace ID of the created span object.
	//
	// example:
	//
	// 550e8400e29b41d4a71644665544****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateCopilotConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCopilotConversationRequest) GoString() string {
	return s.String()
}

func (s *CreateCopilotConversationRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateCopilotConversationRequest) GetModel() *string {
	return s.Model
}

func (s *CreateCopilotConversationRequest) GetPlanMode() *string {
	return s.PlanMode
}

func (s *CreateCopilotConversationRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateCopilotConversationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCopilotConversationRequest) GetThinkingMode() *string {
	return s.ThinkingMode
}

func (s *CreateCopilotConversationRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateCopilotConversationRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateCopilotConversationRequest) SetLang(v string) *CreateCopilotConversationRequest {
	s.Lang = &v
	return s
}

func (s *CreateCopilotConversationRequest) SetModel(v string) *CreateCopilotConversationRequest {
	s.Model = &v
	return s
}

func (s *CreateCopilotConversationRequest) SetPlanMode(v string) *CreateCopilotConversationRequest {
	s.PlanMode = &v
	return s
}

func (s *CreateCopilotConversationRequest) SetProjectId(v string) *CreateCopilotConversationRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateCopilotConversationRequest) SetRegionId(v string) *CreateCopilotConversationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCopilotConversationRequest) SetThinkingMode(v string) *CreateCopilotConversationRequest {
	s.ThinkingMode = &v
	return s
}

func (s *CreateCopilotConversationRequest) SetTitle(v string) *CreateCopilotConversationRequest {
	s.Title = &v
	return s
}

func (s *CreateCopilotConversationRequest) SetTraceId(v string) *CreateCopilotConversationRequest {
	s.TraceId = &v
	return s
}

func (s *CreateCopilotConversationRequest) Validate() error {
	return dara.Validate(s)
}

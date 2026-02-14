// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *CreateCustomAgentShrinkRequest
	GetDMSUnit() *string
	SetDataJson(v string) *CreateCustomAgentShrinkRequest
	GetDataJson() *string
	SetDescription(v string) *CreateCustomAgentShrinkRequest
	GetDescription() *string
	SetExecutionConfigShrink(v string) *CreateCustomAgentShrinkRequest
	GetExecutionConfigShrink() *string
	SetInstruction(v string) *CreateCustomAgentShrinkRequest
	GetInstruction() *string
	SetKnowledge(v string) *CreateCustomAgentShrinkRequest
	GetKnowledge() *string
	SetKnowledgeConfigListShrink(v string) *CreateCustomAgentShrinkRequest
	GetKnowledgeConfigListShrink() *string
	SetName(v string) *CreateCustomAgentShrinkRequest
	GetName() *string
	SetScheduleTaskConfigShrink(v string) *CreateCustomAgentShrinkRequest
	GetScheduleTaskConfigShrink() *string
	SetTextReportConfig(v string) *CreateCustomAgentShrinkRequest
	GetTextReportConfig() *string
	SetWebReportConfig(v string) *CreateCustomAgentShrinkRequest
	GetWebReportConfig() *string
	SetWorkspaceId(v string) *CreateCustomAgentShrinkRequest
	GetWorkspaceId() *string
}

type CreateCustomAgentShrinkRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit                   *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	DataJson                  *string `json:"DataJson,omitempty" xml:"DataJson,omitempty"`
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ExecutionConfigShrink     *string `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty"`
	Instruction               *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	Knowledge                 *string `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	KnowledgeConfigListShrink *string `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty"`
	Name                      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ScheduleTaskConfigShrink  *string `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty"`
	TextReportConfig          *string `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	WebReportConfig           *string `json:"WebReportConfig,omitempty" xml:"WebReportConfig,omitempty"`
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateCustomAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentShrinkRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *CreateCustomAgentShrinkRequest) GetDataJson() *string {
	return s.DataJson
}

func (s *CreateCustomAgentShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomAgentShrinkRequest) GetExecutionConfigShrink() *string {
	return s.ExecutionConfigShrink
}

func (s *CreateCustomAgentShrinkRequest) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateCustomAgentShrinkRequest) GetKnowledge() *string {
	return s.Knowledge
}

func (s *CreateCustomAgentShrinkRequest) GetKnowledgeConfigListShrink() *string {
	return s.KnowledgeConfigListShrink
}

func (s *CreateCustomAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentShrinkRequest) GetScheduleTaskConfigShrink() *string {
	return s.ScheduleTaskConfigShrink
}

func (s *CreateCustomAgentShrinkRequest) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *CreateCustomAgentShrinkRequest) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *CreateCustomAgentShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCustomAgentShrinkRequest) SetDMSUnit(v string) *CreateCustomAgentShrinkRequest {
	s.DMSUnit = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetDataJson(v string) *CreateCustomAgentShrinkRequest {
	s.DataJson = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetDescription(v string) *CreateCustomAgentShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetExecutionConfigShrink(v string) *CreateCustomAgentShrinkRequest {
	s.ExecutionConfigShrink = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetInstruction(v string) *CreateCustomAgentShrinkRequest {
	s.Instruction = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetKnowledge(v string) *CreateCustomAgentShrinkRequest {
	s.Knowledge = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetKnowledgeConfigListShrink(v string) *CreateCustomAgentShrinkRequest {
	s.KnowledgeConfigListShrink = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetName(v string) *CreateCustomAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetScheduleTaskConfigShrink(v string) *CreateCustomAgentShrinkRequest {
	s.ScheduleTaskConfigShrink = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetTextReportConfig(v string) *CreateCustomAgentShrinkRequest {
	s.TextReportConfig = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetWebReportConfig(v string) *CreateCustomAgentShrinkRequest {
	s.WebReportConfig = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetWorkspaceId(v string) *CreateCustomAgentShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}

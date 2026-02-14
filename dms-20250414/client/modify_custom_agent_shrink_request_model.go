// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAgentId(v string) *ModifyCustomAgentShrinkRequest
	GetCustomAgentId() *string
	SetDMSUnit(v string) *ModifyCustomAgentShrinkRequest
	GetDMSUnit() *string
	SetDataJson(v string) *ModifyCustomAgentShrinkRequest
	GetDataJson() *string
	SetDescription(v string) *ModifyCustomAgentShrinkRequest
	GetDescription() *string
	SetExecutionConfigShrink(v string) *ModifyCustomAgentShrinkRequest
	GetExecutionConfigShrink() *string
	SetInstruction(v string) *ModifyCustomAgentShrinkRequest
	GetInstruction() *string
	SetKnowledge(v string) *ModifyCustomAgentShrinkRequest
	GetKnowledge() *string
	SetKnowledgeConfigListShrink(v string) *ModifyCustomAgentShrinkRequest
	GetKnowledgeConfigListShrink() *string
	SetName(v string) *ModifyCustomAgentShrinkRequest
	GetName() *string
	SetScheduleTaskConfigShrink(v string) *ModifyCustomAgentShrinkRequest
	GetScheduleTaskConfigShrink() *string
	SetTextReportConfig(v string) *ModifyCustomAgentShrinkRequest
	GetTextReportConfig() *string
	SetWebReportConfig(v string) *ModifyCustomAgentShrinkRequest
	GetWebReportConfig() *string
	SetWorkspaceId(v string) *ModifyCustomAgentShrinkRequest
	GetWorkspaceId() *string
}

type ModifyCustomAgentShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
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

func (s ModifyCustomAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentShrinkRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *ModifyCustomAgentShrinkRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *ModifyCustomAgentShrinkRequest) GetDataJson() *string {
	return s.DataJson
}

func (s *ModifyCustomAgentShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCustomAgentShrinkRequest) GetExecutionConfigShrink() *string {
	return s.ExecutionConfigShrink
}

func (s *ModifyCustomAgentShrinkRequest) GetInstruction() *string {
	return s.Instruction
}

func (s *ModifyCustomAgentShrinkRequest) GetKnowledge() *string {
	return s.Knowledge
}

func (s *ModifyCustomAgentShrinkRequest) GetKnowledgeConfigListShrink() *string {
	return s.KnowledgeConfigListShrink
}

func (s *ModifyCustomAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCustomAgentShrinkRequest) GetScheduleTaskConfigShrink() *string {
	return s.ScheduleTaskConfigShrink
}

func (s *ModifyCustomAgentShrinkRequest) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *ModifyCustomAgentShrinkRequest) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *ModifyCustomAgentShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ModifyCustomAgentShrinkRequest) SetCustomAgentId(v string) *ModifyCustomAgentShrinkRequest {
	s.CustomAgentId = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetDMSUnit(v string) *ModifyCustomAgentShrinkRequest {
	s.DMSUnit = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetDataJson(v string) *ModifyCustomAgentShrinkRequest {
	s.DataJson = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetDescription(v string) *ModifyCustomAgentShrinkRequest {
	s.Description = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetExecutionConfigShrink(v string) *ModifyCustomAgentShrinkRequest {
	s.ExecutionConfigShrink = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetInstruction(v string) *ModifyCustomAgentShrinkRequest {
	s.Instruction = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetKnowledge(v string) *ModifyCustomAgentShrinkRequest {
	s.Knowledge = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetKnowledgeConfigListShrink(v string) *ModifyCustomAgentShrinkRequest {
	s.KnowledgeConfigListShrink = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetName(v string) *ModifyCustomAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetScheduleTaskConfigShrink(v string) *ModifyCustomAgentShrinkRequest {
	s.ScheduleTaskConfigShrink = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetTextReportConfig(v string) *ModifyCustomAgentShrinkRequest {
	s.TextReportConfig = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetWebReportConfig(v string) *ModifyCustomAgentShrinkRequest {
	s.WebReportConfig = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetWorkspaceId(v string) *ModifyCustomAgentShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}

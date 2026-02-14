// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *CreateCustomAgentRequest
	GetDMSUnit() *string
	SetDataJson(v string) *CreateCustomAgentRequest
	GetDataJson() *string
	SetDescription(v string) *CreateCustomAgentRequest
	GetDescription() *string
	SetExecutionConfig(v *CreateCustomAgentRequestExecutionConfig) *CreateCustomAgentRequest
	GetExecutionConfig() *CreateCustomAgentRequestExecutionConfig
	SetInstruction(v string) *CreateCustomAgentRequest
	GetInstruction() *string
	SetKnowledge(v string) *CreateCustomAgentRequest
	GetKnowledge() *string
	SetKnowledgeConfigList(v []*CreateCustomAgentRequestKnowledgeConfigList) *CreateCustomAgentRequest
	GetKnowledgeConfigList() []*CreateCustomAgentRequestKnowledgeConfigList
	SetName(v string) *CreateCustomAgentRequest
	GetName() *string
	SetScheduleTaskConfig(v *CreateCustomAgentRequestScheduleTaskConfig) *CreateCustomAgentRequest
	GetScheduleTaskConfig() *CreateCustomAgentRequestScheduleTaskConfig
	SetTextReportConfig(v string) *CreateCustomAgentRequest
	GetTextReportConfig() *string
	SetWebReportConfig(v string) *CreateCustomAgentRequest
	GetWebReportConfig() *string
	SetWorkspaceId(v string) *CreateCustomAgentRequest
	GetWorkspaceId() *string
}

type CreateCustomAgentRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit             *string                                        `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	DataJson            *string                                        `json:"DataJson,omitempty" xml:"DataJson,omitempty"`
	Description         *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	ExecutionConfig     *CreateCustomAgentRequestExecutionConfig       `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
	Instruction         *string                                        `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	Knowledge           *string                                        `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	KnowledgeConfigList []*CreateCustomAgentRequestKnowledgeConfigList `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	Name                *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	ScheduleTaskConfig  *CreateCustomAgentRequestScheduleTaskConfig    `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty" type:"Struct"`
	TextReportConfig    *string                                        `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	WebReportConfig     *string                                        `json:"WebReportConfig,omitempty" xml:"WebReportConfig,omitempty"`
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *CreateCustomAgentRequest) GetDataJson() *string {
	return s.DataJson
}

func (s *CreateCustomAgentRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomAgentRequest) GetExecutionConfig() *CreateCustomAgentRequestExecutionConfig {
	return s.ExecutionConfig
}

func (s *CreateCustomAgentRequest) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateCustomAgentRequest) GetKnowledge() *string {
	return s.Knowledge
}

func (s *CreateCustomAgentRequest) GetKnowledgeConfigList() []*CreateCustomAgentRequestKnowledgeConfigList {
	return s.KnowledgeConfigList
}

func (s *CreateCustomAgentRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentRequest) GetScheduleTaskConfig() *CreateCustomAgentRequestScheduleTaskConfig {
	return s.ScheduleTaskConfig
}

func (s *CreateCustomAgentRequest) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *CreateCustomAgentRequest) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *CreateCustomAgentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCustomAgentRequest) SetDMSUnit(v string) *CreateCustomAgentRequest {
	s.DMSUnit = &v
	return s
}

func (s *CreateCustomAgentRequest) SetDataJson(v string) *CreateCustomAgentRequest {
	s.DataJson = &v
	return s
}

func (s *CreateCustomAgentRequest) SetDescription(v string) *CreateCustomAgentRequest {
	s.Description = &v
	return s
}

func (s *CreateCustomAgentRequest) SetExecutionConfig(v *CreateCustomAgentRequestExecutionConfig) *CreateCustomAgentRequest {
	s.ExecutionConfig = v
	return s
}

func (s *CreateCustomAgentRequest) SetInstruction(v string) *CreateCustomAgentRequest {
	s.Instruction = &v
	return s
}

func (s *CreateCustomAgentRequest) SetKnowledge(v string) *CreateCustomAgentRequest {
	s.Knowledge = &v
	return s
}

func (s *CreateCustomAgentRequest) SetKnowledgeConfigList(v []*CreateCustomAgentRequestKnowledgeConfigList) *CreateCustomAgentRequest {
	s.KnowledgeConfigList = v
	return s
}

func (s *CreateCustomAgentRequest) SetName(v string) *CreateCustomAgentRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentRequest) SetScheduleTaskConfig(v *CreateCustomAgentRequestScheduleTaskConfig) *CreateCustomAgentRequest {
	s.ScheduleTaskConfig = v
	return s
}

func (s *CreateCustomAgentRequest) SetTextReportConfig(v string) *CreateCustomAgentRequest {
	s.TextReportConfig = &v
	return s
}

func (s *CreateCustomAgentRequest) SetWebReportConfig(v string) *CreateCustomAgentRequest {
	s.WebReportConfig = &v
	return s
}

func (s *CreateCustomAgentRequest) SetWorkspaceId(v string) *CreateCustomAgentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCustomAgentRequest) Validate() error {
	if s.ExecutionConfig != nil {
		if err := s.ExecutionConfig.Validate(); err != nil {
			return err
		}
	}
	if s.KnowledgeConfigList != nil {
		for _, item := range s.KnowledgeConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScheduleTaskConfig != nil {
		if err := s.ScheduleTaskConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomAgentRequestExecutionConfig struct {
	// example:
	//
	// true
	SkipAskHuman *bool `json:"SkipAskHuman,omitempty" xml:"SkipAskHuman,omitempty"`
	// example:
	//
	// true
	SkipPlan *bool `json:"SkipPlan,omitempty" xml:"SkipPlan,omitempty"`
	// example:
	//
	// true
	SkipSqlConfirm *bool `json:"SkipSqlConfirm,omitempty" xml:"SkipSqlConfirm,omitempty"`
	// example:
	//
	// true
	SkipWebReportConfirm *bool `json:"SkipWebReportConfirm,omitempty" xml:"SkipWebReportConfirm,omitempty"`
}

func (s CreateCustomAgentRequestExecutionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestExecutionConfig) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestExecutionConfig) GetSkipAskHuman() *bool {
	return s.SkipAskHuman
}

func (s *CreateCustomAgentRequestExecutionConfig) GetSkipPlan() *bool {
	return s.SkipPlan
}

func (s *CreateCustomAgentRequestExecutionConfig) GetSkipSqlConfirm() *bool {
	return s.SkipSqlConfirm
}

func (s *CreateCustomAgentRequestExecutionConfig) GetSkipWebReportConfirm() *bool {
	return s.SkipWebReportConfirm
}

func (s *CreateCustomAgentRequestExecutionConfig) SetSkipAskHuman(v bool) *CreateCustomAgentRequestExecutionConfig {
	s.SkipAskHuman = &v
	return s
}

func (s *CreateCustomAgentRequestExecutionConfig) SetSkipPlan(v bool) *CreateCustomAgentRequestExecutionConfig {
	s.SkipPlan = &v
	return s
}

func (s *CreateCustomAgentRequestExecutionConfig) SetSkipSqlConfirm(v bool) *CreateCustomAgentRequestExecutionConfig {
	s.SkipSqlConfirm = &v
	return s
}

func (s *CreateCustomAgentRequestExecutionConfig) SetSkipWebReportConfirm(v bool) *CreateCustomAgentRequestExecutionConfig {
	s.SkipWebReportConfirm = &v
	return s
}

func (s *CreateCustomAgentRequestExecutionConfig) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestKnowledgeConfigList struct {
	// example:
	//
	// mcp
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// nhdpt9adf6ac**********ca
	McpServerId *string `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
}

func (s CreateCustomAgentRequestKnowledgeConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestKnowledgeConfigList) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestKnowledgeConfigList) GetAccessType() *string {
	return s.AccessType
}

func (s *CreateCustomAgentRequestKnowledgeConfigList) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *CreateCustomAgentRequestKnowledgeConfigList) SetAccessType(v string) *CreateCustomAgentRequestKnowledgeConfigList {
	s.AccessType = &v
	return s
}

func (s *CreateCustomAgentRequestKnowledgeConfigList) SetMcpServerId(v string) *CreateCustomAgentRequestKnowledgeConfigList {
	s.McpServerId = &v
	return s
}

func (s *CreateCustomAgentRequestKnowledgeConfigList) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestScheduleTaskConfig struct {
	// example:
	//
	// 0 0 0 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 4m24*****mg7j2v
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
}

func (s CreateCustomAgentRequestScheduleTaskConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestScheduleTaskConfig) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestScheduleTaskConfig) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateCustomAgentRequestScheduleTaskConfig) GetQuery() *string {
	return s.Query
}

func (s *CreateCustomAgentRequestScheduleTaskConfig) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *CreateCustomAgentRequestScheduleTaskConfig) SetCronExpression(v string) *CreateCustomAgentRequestScheduleTaskConfig {
	s.CronExpression = &v
	return s
}

func (s *CreateCustomAgentRequestScheduleTaskConfig) SetQuery(v string) *CreateCustomAgentRequestScheduleTaskConfig {
	s.Query = &v
	return s
}

func (s *CreateCustomAgentRequestScheduleTaskConfig) SetRelatedSessionId(v string) *CreateCustomAgentRequestScheduleTaskConfig {
	s.RelatedSessionId = &v
	return s
}

func (s *CreateCustomAgentRequestScheduleTaskConfig) Validate() error {
	return dara.Validate(s)
}

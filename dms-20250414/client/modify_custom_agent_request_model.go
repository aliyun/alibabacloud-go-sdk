// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomAgentId(v string) *ModifyCustomAgentRequest
	GetCustomAgentId() *string
	SetDMSUnit(v string) *ModifyCustomAgentRequest
	GetDMSUnit() *string
	SetDataJson(v string) *ModifyCustomAgentRequest
	GetDataJson() *string
	SetDescription(v string) *ModifyCustomAgentRequest
	GetDescription() *string
	SetExecutionConfig(v *ModifyCustomAgentRequestExecutionConfig) *ModifyCustomAgentRequest
	GetExecutionConfig() *ModifyCustomAgentRequestExecutionConfig
	SetInstruction(v string) *ModifyCustomAgentRequest
	GetInstruction() *string
	SetKnowledge(v string) *ModifyCustomAgentRequest
	GetKnowledge() *string
	SetKnowledgeConfigList(v []*ModifyCustomAgentRequestKnowledgeConfigList) *ModifyCustomAgentRequest
	GetKnowledgeConfigList() []*ModifyCustomAgentRequestKnowledgeConfigList
	SetName(v string) *ModifyCustomAgentRequest
	GetName() *string
	SetScheduleTaskConfig(v *ModifyCustomAgentRequestScheduleTaskConfig) *ModifyCustomAgentRequest
	GetScheduleTaskConfig() *ModifyCustomAgentRequestScheduleTaskConfig
	SetTextReportConfig(v string) *ModifyCustomAgentRequest
	GetTextReportConfig() *string
	SetWebReportConfig(v string) *ModifyCustomAgentRequest
	GetWebReportConfig() *string
	SetWorkspaceId(v string) *ModifyCustomAgentRequest
	GetWorkspaceId() *string
}

type ModifyCustomAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// example:
	//
	// cn-hangzhou
	DMSUnit             *string                                        `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	DataJson            *string                                        `json:"DataJson,omitempty" xml:"DataJson,omitempty"`
	Description         *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	ExecutionConfig     *ModifyCustomAgentRequestExecutionConfig       `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
	Instruction         *string                                        `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	Knowledge           *string                                        `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	KnowledgeConfigList []*ModifyCustomAgentRequestKnowledgeConfigList `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	Name                *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	ScheduleTaskConfig  *ModifyCustomAgentRequestScheduleTaskConfig    `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty" type:"Struct"`
	TextReportConfig    *string                                        `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	WebReportConfig     *string                                        `json:"WebReportConfig,omitempty" xml:"WebReportConfig,omitempty"`
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ModifyCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentRequest) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *ModifyCustomAgentRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *ModifyCustomAgentRequest) GetDataJson() *string {
	return s.DataJson
}

func (s *ModifyCustomAgentRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCustomAgentRequest) GetExecutionConfig() *ModifyCustomAgentRequestExecutionConfig {
	return s.ExecutionConfig
}

func (s *ModifyCustomAgentRequest) GetInstruction() *string {
	return s.Instruction
}

func (s *ModifyCustomAgentRequest) GetKnowledge() *string {
	return s.Knowledge
}

func (s *ModifyCustomAgentRequest) GetKnowledgeConfigList() []*ModifyCustomAgentRequestKnowledgeConfigList {
	return s.KnowledgeConfigList
}

func (s *ModifyCustomAgentRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCustomAgentRequest) GetScheduleTaskConfig() *ModifyCustomAgentRequestScheduleTaskConfig {
	return s.ScheduleTaskConfig
}

func (s *ModifyCustomAgentRequest) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *ModifyCustomAgentRequest) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *ModifyCustomAgentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ModifyCustomAgentRequest) SetCustomAgentId(v string) *ModifyCustomAgentRequest {
	s.CustomAgentId = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetDMSUnit(v string) *ModifyCustomAgentRequest {
	s.DMSUnit = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetDataJson(v string) *ModifyCustomAgentRequest {
	s.DataJson = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetDescription(v string) *ModifyCustomAgentRequest {
	s.Description = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetExecutionConfig(v *ModifyCustomAgentRequestExecutionConfig) *ModifyCustomAgentRequest {
	s.ExecutionConfig = v
	return s
}

func (s *ModifyCustomAgentRequest) SetInstruction(v string) *ModifyCustomAgentRequest {
	s.Instruction = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetKnowledge(v string) *ModifyCustomAgentRequest {
	s.Knowledge = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetKnowledgeConfigList(v []*ModifyCustomAgentRequestKnowledgeConfigList) *ModifyCustomAgentRequest {
	s.KnowledgeConfigList = v
	return s
}

func (s *ModifyCustomAgentRequest) SetName(v string) *ModifyCustomAgentRequest {
	s.Name = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetScheduleTaskConfig(v *ModifyCustomAgentRequestScheduleTaskConfig) *ModifyCustomAgentRequest {
	s.ScheduleTaskConfig = v
	return s
}

func (s *ModifyCustomAgentRequest) SetTextReportConfig(v string) *ModifyCustomAgentRequest {
	s.TextReportConfig = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetWebReportConfig(v string) *ModifyCustomAgentRequest {
	s.WebReportConfig = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetWorkspaceId(v string) *ModifyCustomAgentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ModifyCustomAgentRequest) Validate() error {
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

type ModifyCustomAgentRequestExecutionConfig struct {
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

func (s ModifyCustomAgentRequestExecutionConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentRequestExecutionConfig) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentRequestExecutionConfig) GetSkipAskHuman() *bool {
	return s.SkipAskHuman
}

func (s *ModifyCustomAgentRequestExecutionConfig) GetSkipPlan() *bool {
	return s.SkipPlan
}

func (s *ModifyCustomAgentRequestExecutionConfig) GetSkipSqlConfirm() *bool {
	return s.SkipSqlConfirm
}

func (s *ModifyCustomAgentRequestExecutionConfig) GetSkipWebReportConfirm() *bool {
	return s.SkipWebReportConfirm
}

func (s *ModifyCustomAgentRequestExecutionConfig) SetSkipAskHuman(v bool) *ModifyCustomAgentRequestExecutionConfig {
	s.SkipAskHuman = &v
	return s
}

func (s *ModifyCustomAgentRequestExecutionConfig) SetSkipPlan(v bool) *ModifyCustomAgentRequestExecutionConfig {
	s.SkipPlan = &v
	return s
}

func (s *ModifyCustomAgentRequestExecutionConfig) SetSkipSqlConfirm(v bool) *ModifyCustomAgentRequestExecutionConfig {
	s.SkipSqlConfirm = &v
	return s
}

func (s *ModifyCustomAgentRequestExecutionConfig) SetSkipWebReportConfirm(v bool) *ModifyCustomAgentRequestExecutionConfig {
	s.SkipWebReportConfirm = &v
	return s
}

func (s *ModifyCustomAgentRequestExecutionConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyCustomAgentRequestKnowledgeConfigList struct {
	// example:
	//
	// mcp
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// example:
	//
	// nhdpt9adf6ac**********ca
	McpServerId *string `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
}

func (s ModifyCustomAgentRequestKnowledgeConfigList) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentRequestKnowledgeConfigList) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentRequestKnowledgeConfigList) GetAccessType() *string {
	return s.AccessType
}

func (s *ModifyCustomAgentRequestKnowledgeConfigList) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ModifyCustomAgentRequestKnowledgeConfigList) SetAccessType(v string) *ModifyCustomAgentRequestKnowledgeConfigList {
	s.AccessType = &v
	return s
}

func (s *ModifyCustomAgentRequestKnowledgeConfigList) SetMcpServerId(v string) *ModifyCustomAgentRequestKnowledgeConfigList {
	s.McpServerId = &v
	return s
}

func (s *ModifyCustomAgentRequestKnowledgeConfigList) Validate() error {
	return dara.Validate(s)
}

type ModifyCustomAgentRequestScheduleTaskConfig struct {
	// example:
	//
	// 0 0 0,1 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	Query          *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 4m24*****mg7j2v
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
}

func (s ModifyCustomAgentRequestScheduleTaskConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentRequestScheduleTaskConfig) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentRequestScheduleTaskConfig) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyCustomAgentRequestScheduleTaskConfig) GetQuery() *string {
	return s.Query
}

func (s *ModifyCustomAgentRequestScheduleTaskConfig) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *ModifyCustomAgentRequestScheduleTaskConfig) SetCronExpression(v string) *ModifyCustomAgentRequestScheduleTaskConfig {
	s.CronExpression = &v
	return s
}

func (s *ModifyCustomAgentRequestScheduleTaskConfig) SetQuery(v string) *ModifyCustomAgentRequestScheduleTaskConfig {
	s.Query = &v
	return s
}

func (s *ModifyCustomAgentRequestScheduleTaskConfig) SetRelatedSessionId(v string) *ModifyCustomAgentRequestScheduleTaskConfig {
	s.RelatedSessionId = &v
	return s
}

func (s *ModifyCustomAgentRequestScheduleTaskConfig) Validate() error {
	return dara.Validate(s)
}

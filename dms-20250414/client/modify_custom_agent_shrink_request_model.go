// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackConfigShrink(v string) *ModifyCustomAgentShrinkRequest
	GetCallbackConfigShrink() *string
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
	SetKnowledgeSemanticConfigListShrink(v string) *ModifyCustomAgentShrinkRequest
	GetKnowledgeSemanticConfigListShrink() *string
	SetName(v string) *ModifyCustomAgentShrinkRequest
	GetName() *string
	SetRelatedSessionId(v string) *ModifyCustomAgentShrinkRequest
	GetRelatedSessionId() *string
	SetScheduleTaskConfigShrink(v string) *ModifyCustomAgentShrinkRequest
	GetScheduleTaskConfigShrink() *string
	SetTextReportConfig(v string) *ModifyCustomAgentShrinkRequest
	GetTextReportConfig() *string
	SetUserSpecifiedSkillListShrink(v string) *ModifyCustomAgentShrinkRequest
	GetUserSpecifiedSkillListShrink() *string
	SetWebReportConfig(v string) *ModifyCustomAgentShrinkRequest
	GetWebReportConfig() *string
	SetWebReportTheme(v string) *ModifyCustomAgentShrinkRequest
	GetWebReportTheme() *string
	SetWorkspaceId(v string) *ModifyCustomAgentShrinkRequest
	GetWorkspaceId() *string
}

type ModifyCustomAgentShrinkRequest struct {
	CallbackConfigShrink *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
	// The custom agent ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The specified data scope in **JSON string format**.
	//
	// - Common parameter description
	//
	//   - tableFlag: true indicates a specified data scope
	//
	//   - scope: personal is a fixed value
	//
	//   - personal: pass parameters for file or database types
	//
	// **File type**. Pass parameters in the following format:
	//
	// - DataSourceType: remote_data_center is a fixed value
	//
	// - FileId: the file ID
	//
	// - Database: the database name returned by the ListDataCenterTable operation, which is usually the file name
	//
	// - Tables: the table name returned by the ListDataCenterTable operation
	//
	// - TableIds: the TableId returned by the ListDataCenterTable operation
	//
	// - RegionId: the current region
	//
	// ```
	//
	// {
	//
	//   "tableFlag": true,
	//
	//   "scope": "personal",
	//
	//   "personal": {
	//
	//     "DataSourceType": "remote_data_center",
	//
	//     "FileId": "f-f0jksn001ibmkoo********6v2zn6",
	//
	//     "Database": "diamonds.csv",
	//
	//     "Tables": [
	//
	//       "diamonds"
	//
	//     ],
	//
	//     "TableIds": [
	//
	//       "35hfn94pxl********50pi"
	//
	//     ],
	//
	//     "RegionId": "cn-hangzhou"
	//
	//   }
	//
	// }
	//
	// ```
	//
	// **Database type**. Pass parameters in the following format:
	//
	// - DataSourceType: database is a fixed value
	//
	// - DmsInstanceId: the DMS instance ID returned by the data center operation
	//
	// - DmsDatabaseId: the DMS database ID returned by the data center operation
	//
	// - FileId: the instance name (deprecated)
	//
	// - DbName: the database name returned by the data center operation
	//
	// - Database: the database name returned by the data center operation
	//
	// - Tables: the table name returned by the data center operation
	//
	// - TableIds: the TableId returned by the data center operation
	//
	// - Engine: the engine type (mysql or postgresql)
	//
	// - RegionId: the current region
	//
	// ```
	//
	// {
	//
	//   "tableFlag": true,
	//
	//   "scope": "personal",
	//
	//   "personal": {
	//
	//     "DataSourceType": "database",
	//
	//     "DmsInstanceId": "284***8",
	//
	//     "DmsDatabaseId": "769***45",
	//
	//     "FileId": "pgm-bp15095e*******6t",
	//
	//     "DbName": "pg_catalog",
	//
	//     "Database": "pg_catalog",
	//
	//     "Tables": [
	//
	//       "pg_aggregate"
	//
	//     ],
	//
	//     "TableIds": [
	//
	//       "5263****31"
	//
	//     ],
	//
	//     "Engine": "postgresql",
	//
	//     "RegionId": "cn-hangzhou"
	//
	//   }
	//
	// }
	//
	// ```
	//
	// example:
	//
	// {
	//
	//   "tableFlag" : true,
	//
	//   "scope" : "personal",
	//
	//   "personal" : {
	//
	//     "DataSourceType" : "remote_data_center",
	//
	//     "FileId" : "f-5qlrwaw10********s3gpw1z",
	//
	//     "Database" : "TestTable******.xlsx",
	//
	//     "Tables" : [ "Sheet1" ],
	//
	//     "TableIds" : [ "******" ],
	//
	//     "RegionId" : "cn-hangzhou"
	//
	//   }
	//
	// }
	DataJson *string `json:"DataJson,omitempty" xml:"DataJson,omitempty"`
	// The description of the custom agent.
	//
	// example:
	//
	// AgentTestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution configuration.
	ExecutionConfigShrink *string `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty"`
	// The instruction.
	//
	// example:
	//
	// Analysis framework:
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// The knowledge.
	//
	// example:
	//
	// Core metric definitions:
	//
	// 1. GMV (Gross Merchandise Volume) refers to the total order amount, including both paid and unpaid orders.
	//
	// 2. Order volume is the number of valid orders placed per day.
	//
	// 3. UV (Unique Visitors) refers to the deduplicated number of users who visit the website or app.
	//
	// 4. Conversion rate = number of paid orders / UV, reflecting traffic conversion efficiency.
	Knowledge *string `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	// The external knowledge bases.
	KnowledgeConfigListShrink         *string `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty"`
	KnowledgeSemanticConfigListShrink *string `json:"KnowledgeSemanticConfigList,omitempty" xml:"KnowledgeSemanticConfigList,omitempty"`
	// The name of the custom agent.
	//
	// example:
	//
	// AgentTestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the referenced historical session.
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The scheduled task configuration.
	ScheduleTaskConfigShrink *string `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty"`
	// The text report format.
	//
	// example:
	//
	// The text report requires all numbers to be written in Chinese characters instead of Arabic numerals
	TextReportConfig             *string `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	UserSpecifiedSkillListShrink *string `json:"UserSpecifiedSkillList,omitempty" xml:"UserSpecifiedSkillList,omitempty"`
	// The web report format.
	//
	// example:
	//
	// The web report requires all numbers to be written in Chinese characters instead of Arabic numerals
	WebReportConfig *string `json:"WebReportConfig,omitempty" xml:"WebReportConfig,omitempty"`
	WebReportTheme  *string `json:"WebReportTheme,omitempty" xml:"WebReportTheme,omitempty"`
	// The workspace ID.
	//
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

func (s *ModifyCustomAgentShrinkRequest) GetCallbackConfigShrink() *string {
	return s.CallbackConfigShrink
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

func (s *ModifyCustomAgentShrinkRequest) GetKnowledgeSemanticConfigListShrink() *string {
	return s.KnowledgeSemanticConfigListShrink
}

func (s *ModifyCustomAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCustomAgentShrinkRequest) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *ModifyCustomAgentShrinkRequest) GetScheduleTaskConfigShrink() *string {
	return s.ScheduleTaskConfigShrink
}

func (s *ModifyCustomAgentShrinkRequest) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *ModifyCustomAgentShrinkRequest) GetUserSpecifiedSkillListShrink() *string {
	return s.UserSpecifiedSkillListShrink
}

func (s *ModifyCustomAgentShrinkRequest) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *ModifyCustomAgentShrinkRequest) GetWebReportTheme() *string {
	return s.WebReportTheme
}

func (s *ModifyCustomAgentShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ModifyCustomAgentShrinkRequest) SetCallbackConfigShrink(v string) *ModifyCustomAgentShrinkRequest {
	s.CallbackConfigShrink = &v
	return s
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

func (s *ModifyCustomAgentShrinkRequest) SetKnowledgeSemanticConfigListShrink(v string) *ModifyCustomAgentShrinkRequest {
	s.KnowledgeSemanticConfigListShrink = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetName(v string) *ModifyCustomAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetRelatedSessionId(v string) *ModifyCustomAgentShrinkRequest {
	s.RelatedSessionId = &v
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

func (s *ModifyCustomAgentShrinkRequest) SetUserSpecifiedSkillListShrink(v string) *ModifyCustomAgentShrinkRequest {
	s.UserSpecifiedSkillListShrink = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetWebReportConfig(v string) *ModifyCustomAgentShrinkRequest {
	s.WebReportConfig = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetWebReportTheme(v string) *ModifyCustomAgentShrinkRequest {
	s.WebReportTheme = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) SetWorkspaceId(v string) *ModifyCustomAgentShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ModifyCustomAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}

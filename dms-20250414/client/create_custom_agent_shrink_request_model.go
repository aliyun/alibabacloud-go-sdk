// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *CreateCustomAgentShrinkRequest
	GetAgentType() *string
	SetCallbackConfigShrink(v string) *CreateCustomAgentShrinkRequest
	GetCallbackConfigShrink() *string
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
	SetKnowledgeSemanticConfigListShrink(v string) *CreateCustomAgentShrinkRequest
	GetKnowledgeSemanticConfigListShrink() *string
	SetName(v string) *CreateCustomAgentShrinkRequest
	GetName() *string
	SetRecommendQuestionListShrink(v string) *CreateCustomAgentShrinkRequest
	GetRecommendQuestionListShrink() *string
	SetRelatedSessionId(v string) *CreateCustomAgentShrinkRequest
	GetRelatedSessionId() *string
	SetScheduleTaskConfigShrink(v string) *CreateCustomAgentShrinkRequest
	GetScheduleTaskConfigShrink() *string
	SetTextReportConfig(v string) *CreateCustomAgentShrinkRequest
	GetTextReportConfig() *string
	SetUserSpecifiedSkillListShrink(v string) *CreateCustomAgentShrinkRequest
	GetUserSpecifiedSkillListShrink() *string
	SetWebReportConfig(v string) *CreateCustomAgentShrinkRequest
	GetWebReportConfig() *string
	SetWebReportTheme(v string) *CreateCustomAgentShrinkRequest
	GetWebReportTheme() *string
	SetWorkspaceId(v string) *CreateCustomAgentShrinkRequest
	GetWorkspaceId() *string
}

type CreateCustomAgentShrinkRequest struct {
	// The type of the custom agent. Valid values:
	//
	// - Default: The built-in custom agent. Each workspace can have at most one agent of this type, and it cannot be deleted.
	//
	// - Common: A regular custom agent.
	//
	// example:
	//
	// Common
	AgentType            *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	CallbackConfigShrink *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The specified data scope in **JSON character string format**.
	//
	// - General parameter description:
	//
	//   - tableFlag: Set to true to specify the data scope.
	//
	//   - scope: Set to personal as a fixed value.
	//
	//   - personal: The parameters for file and database types.
	//
	// **File type**: Pass parameters in the following format:
	//
	// - DataSourceType: Set to remote_data_center as a fixed value.
	//
	// - FileId: The file ID.
	//
	// - Database: The database name returned by the ListDataCenterTable operation, which is usually the file name.
	//
	// - Tables: The table names returned by the ListDataCenterTable operation.
	//
	// - TableIds: The table IDs returned by the ListDataCenterTable operation.
	//
	// - RegionId: The current region.
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
	// **Database type**: Pass parameters in the following format:
	//
	// - DataSourceType: Set to database as a fixed value.
	//
	// - DmsInstanceId: The DMS instance ID returned by the data center operation.
	//
	// - DmsDatabaseId: The DMS database ID returned by the data center operation.
	//
	// - FileId: The instance name. (Deprecated)
	//
	// - DbName: The database name returned by the data center operation.
	//
	// - Database: The database name returned by the data center operation.
	//
	// - Tables: The table names returned by the data center operation.
	//
	// - TableIds: The table IDs returned by the data center operation.
	//
	// - Engine: The engine type, such as mysql or postgresql.
	//
	// - RegionId: The current region.
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
	//     "Database" : "Sample table******.xlsx",
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
	// Sample agent description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution configuration.
	ExecutionConfigShrink *string `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty"`
	// The instructions. Maximum length: 10,000 characters.
	//
	// example:
	//
	// Core metric definitions:
	//
	// 1. GMV (Gross Merchandise Volume) refers to the total sum of order amounts, including both paid and unpaid orders.
	//
	// 2. Order volume refers to the number of valid orders placed per day.
	//
	// 3. UV (Unique Visitors) refers to the number of deduplicated users visiting the website or app.
	//
	// 4. Conversion rate = Number of paid orders / UV, reflecting the traffic conversion efficiency
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// The knowledge. Maximum length: 10,000 characters.
	//
	// example:
	//
	// Core metric definitions:
	//
	// 1. GMV (Gross Merchandise Volume) refers to the total sum of order amounts, including both paid and unpaid orders.
	//
	// 2. Order volume refers to the number of valid orders placed per day.
	//
	// 3. UV (Unique Visitors) refers to the number of deduplicated users visiting the website or app.
	//
	// 4. Conversion rate = Number of paid orders / UV, reflecting the traffic conversion efficiency
	Knowledge *string `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	// The external knowledge bases.
	KnowledgeConfigListShrink         *string `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty"`
	KnowledgeSemanticConfigListShrink *string `json:"KnowledgeSemanticConfigList,omitempty" xml:"KnowledgeSemanticConfigList,omitempty"`
	// The name of the custom agent.
	//
	// example:
	//
	// Sample agent name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of recommended questions. You can configure up to 10 recommended questions.
	RecommendQuestionListShrink *string `json:"RecommendQuestionList,omitempty" xml:"RecommendQuestionList,omitempty"`
	// The ID of the referenced historical session.
	//
	// example:
	//
	// 99crainsjnveo*******vynid
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The scheduled task configuration.
	ScheduleTaskConfigShrink *string `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty"`
	// The text report format.
	//
	// example:
	//
	// For the text report, use Chinese characters for all numbers instead of Arabic numerals
	TextReportConfig *string `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	// The list of user-specified skills.
	UserSpecifiedSkillListShrink *string `json:"UserSpecifiedSkillList,omitempty" xml:"UserSpecifiedSkillList,omitempty"`
	// The web report format.
	//
	// example:
	//
	// For the web report, use Chinese characters for all numbers instead of Arabic numerals
	WebReportConfig *string `json:"WebReportConfig,omitempty" xml:"WebReportConfig,omitempty"`
	// The web report theme.
	//
	// example:
	//
	// Default
	WebReportTheme *string `json:"WebReportTheme,omitempty" xml:"WebReportTheme,omitempty"`
	// The workspace ID.
	//
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

func (s *CreateCustomAgentShrinkRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *CreateCustomAgentShrinkRequest) GetCallbackConfigShrink() *string {
	return s.CallbackConfigShrink
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

func (s *CreateCustomAgentShrinkRequest) GetKnowledgeSemanticConfigListShrink() *string {
	return s.KnowledgeSemanticConfigListShrink
}

func (s *CreateCustomAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentShrinkRequest) GetRecommendQuestionListShrink() *string {
	return s.RecommendQuestionListShrink
}

func (s *CreateCustomAgentShrinkRequest) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *CreateCustomAgentShrinkRequest) GetScheduleTaskConfigShrink() *string {
	return s.ScheduleTaskConfigShrink
}

func (s *CreateCustomAgentShrinkRequest) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *CreateCustomAgentShrinkRequest) GetUserSpecifiedSkillListShrink() *string {
	return s.UserSpecifiedSkillListShrink
}

func (s *CreateCustomAgentShrinkRequest) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *CreateCustomAgentShrinkRequest) GetWebReportTheme() *string {
	return s.WebReportTheme
}

func (s *CreateCustomAgentShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCustomAgentShrinkRequest) SetAgentType(v string) *CreateCustomAgentShrinkRequest {
	s.AgentType = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetCallbackConfigShrink(v string) *CreateCustomAgentShrinkRequest {
	s.CallbackConfigShrink = &v
	return s
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

func (s *CreateCustomAgentShrinkRequest) SetKnowledgeSemanticConfigListShrink(v string) *CreateCustomAgentShrinkRequest {
	s.KnowledgeSemanticConfigListShrink = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetName(v string) *CreateCustomAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetRecommendQuestionListShrink(v string) *CreateCustomAgentShrinkRequest {
	s.RecommendQuestionListShrink = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetRelatedSessionId(v string) *CreateCustomAgentShrinkRequest {
	s.RelatedSessionId = &v
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

func (s *CreateCustomAgentShrinkRequest) SetUserSpecifiedSkillListShrink(v string) *CreateCustomAgentShrinkRequest {
	s.UserSpecifiedSkillListShrink = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetWebReportConfig(v string) *CreateCustomAgentShrinkRequest {
	s.WebReportConfig = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetWebReportTheme(v string) *CreateCustomAgentShrinkRequest {
	s.WebReportTheme = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) SetWorkspaceId(v string) *CreateCustomAgentShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetRelatedSessionId(v string) *CreateCustomAgentShrinkRequest
	GetRelatedSessionId() *string
	SetScheduleTaskConfigShrink(v string) *CreateCustomAgentShrinkRequest
	GetScheduleTaskConfigShrink() *string
	SetTextReportConfig(v string) *CreateCustomAgentShrinkRequest
	GetTextReportConfig() *string
	SetWebReportConfig(v string) *CreateCustomAgentShrinkRequest
	GetWebReportConfig() *string
	SetWebReportTheme(v string) *CreateCustomAgentShrinkRequest
	GetWebReportTheme() *string
	SetWorkspaceId(v string) *CreateCustomAgentShrinkRequest
	GetWorkspaceId() *string
}

type CreateCustomAgentShrinkRequest struct {
	CallbackConfigShrink *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The specified data range in **JSON string format**.
	//
	// - Common parameter description
	//
	//   - tableFlag: true indicates a specified data range.
	//
	//   - scope: personal is a fixed value.
	//
	//   - personal: pass parameters for file or database types.
	//
	// **File type**. Pass parameters in the following format:
	//
	// - DataSourceType: remote_data_center is a fixed value.
	//
	// - FileId: The file ID.
	//
	// - Database: The database name returned by the ListDataCenterTable operation, which is usually the file name.
	//
	// - Tables: The table name returned by the ListDataCenterTable operation.
	//
	// - TableIds: The TableId returned by the ListDataCenterTable operation.
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
	// **Database type**. Pass parameters as follows:
	//
	// - DataSourceType: database is a fixed value.
	//
	// - DmsInstanceId: The DMS instance ID returned by the data center operation.
	//
	// - DmsDatabaseId: The DMS database ID returned by the data center operation.
	//
	// - FileId: The instance name (deprecated).
	//
	// - DbName: The database name returned by the data center operation.
	//
	// - Database: The database name returned by the data center operation.
	//
	// - Tables: The table name returned by the data center operation.
	//
	// - TableIds: The TableId returned by the data center operation.
	//
	// - Engine: The engine type (mysql or postgresql).
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
	// Core metric definitions:
	//
	// 1. GMV (Gross Merchandise Volume) refers to the total order amount, including both paid and unpaid orders;
	//
	// 2. Order volume is the number of valid orders placed per day;
	//
	// 3. UV (Unique Visitors) refers to the deduplicated number of users who visit the website or app;
	//
	// 4. Conversion rate = number of paid orders / UV, reflecting traffic conversion efficiency;
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
	// The external knowledge base configurations.
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
	TextReportConfig *string `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
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

func (s CreateCustomAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentShrinkRequest) GoString() string {
	return s.String()
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

func (s *CreateCustomAgentShrinkRequest) GetRelatedSessionId() *string {
	return s.RelatedSessionId
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

func (s *CreateCustomAgentShrinkRequest) GetWebReportTheme() *string {
	return s.WebReportTheme
}

func (s *CreateCustomAgentShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
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

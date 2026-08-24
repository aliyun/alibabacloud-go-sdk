// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackConfig(v *ModifyCustomAgentRequestCallbackConfig) *ModifyCustomAgentRequest
	GetCallbackConfig() *ModifyCustomAgentRequestCallbackConfig
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
	SetKnowledgeSemanticConfigList(v []*ModifyCustomAgentRequestKnowledgeSemanticConfigList) *ModifyCustomAgentRequest
	GetKnowledgeSemanticConfigList() []*ModifyCustomAgentRequestKnowledgeSemanticConfigList
	SetName(v string) *ModifyCustomAgentRequest
	GetName() *string
	SetRelatedSessionId(v string) *ModifyCustomAgentRequest
	GetRelatedSessionId() *string
	SetScheduleTaskConfig(v *ModifyCustomAgentRequestScheduleTaskConfig) *ModifyCustomAgentRequest
	GetScheduleTaskConfig() *ModifyCustomAgentRequestScheduleTaskConfig
	SetTextReportConfig(v string) *ModifyCustomAgentRequest
	GetTextReportConfig() *string
	SetUserSpecifiedSkillList(v []*string) *ModifyCustomAgentRequest
	GetUserSpecifiedSkillList() []*string
	SetWebReportConfig(v string) *ModifyCustomAgentRequest
	GetWebReportConfig() *string
	SetWebReportTheme(v string) *ModifyCustomAgentRequest
	GetWebReportTheme() *string
	SetWorkspaceId(v string) *ModifyCustomAgentRequest
	GetWorkspaceId() *string
}

type ModifyCustomAgentRequest struct {
	CallbackConfig *ModifyCustomAgentRequestCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
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
	ExecutionConfig *ModifyCustomAgentRequestExecutionConfig `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
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
	KnowledgeConfigList         []*ModifyCustomAgentRequestKnowledgeConfigList         `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	KnowledgeSemanticConfigList []*ModifyCustomAgentRequestKnowledgeSemanticConfigList `json:"KnowledgeSemanticConfigList,omitempty" xml:"KnowledgeSemanticConfigList,omitempty" type:"Repeated"`
	// The name of the custom agent.
	//
	// example:
	//
	// AgentTestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the referenced historical session.
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The scheduled task configuration.
	ScheduleTaskConfig *ModifyCustomAgentRequestScheduleTaskConfig `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty" type:"Struct"`
	// The text report format.
	//
	// example:
	//
	// The text report requires all numbers to be written in Chinese characters instead of Arabic numerals
	TextReportConfig       *string   `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	UserSpecifiedSkillList []*string `json:"UserSpecifiedSkillList,omitempty" xml:"UserSpecifiedSkillList,omitempty" type:"Repeated"`
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

func (s ModifyCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentRequest) GetCallbackConfig() *ModifyCustomAgentRequestCallbackConfig {
	return s.CallbackConfig
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

func (s *ModifyCustomAgentRequest) GetKnowledgeSemanticConfigList() []*ModifyCustomAgentRequestKnowledgeSemanticConfigList {
	return s.KnowledgeSemanticConfigList
}

func (s *ModifyCustomAgentRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCustomAgentRequest) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *ModifyCustomAgentRequest) GetScheduleTaskConfig() *ModifyCustomAgentRequestScheduleTaskConfig {
	return s.ScheduleTaskConfig
}

func (s *ModifyCustomAgentRequest) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *ModifyCustomAgentRequest) GetUserSpecifiedSkillList() []*string {
	return s.UserSpecifiedSkillList
}

func (s *ModifyCustomAgentRequest) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *ModifyCustomAgentRequest) GetWebReportTheme() *string {
	return s.WebReportTheme
}

func (s *ModifyCustomAgentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ModifyCustomAgentRequest) SetCallbackConfig(v *ModifyCustomAgentRequestCallbackConfig) *ModifyCustomAgentRequest {
	s.CallbackConfig = v
	return s
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

func (s *ModifyCustomAgentRequest) SetKnowledgeSemanticConfigList(v []*ModifyCustomAgentRequestKnowledgeSemanticConfigList) *ModifyCustomAgentRequest {
	s.KnowledgeSemanticConfigList = v
	return s
}

func (s *ModifyCustomAgentRequest) SetName(v string) *ModifyCustomAgentRequest {
	s.Name = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetRelatedSessionId(v string) *ModifyCustomAgentRequest {
	s.RelatedSessionId = &v
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

func (s *ModifyCustomAgentRequest) SetUserSpecifiedSkillList(v []*string) *ModifyCustomAgentRequest {
	s.UserSpecifiedSkillList = v
	return s
}

func (s *ModifyCustomAgentRequest) SetWebReportConfig(v string) *ModifyCustomAgentRequest {
	s.WebReportConfig = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetWebReportTheme(v string) *ModifyCustomAgentRequest {
	s.WebReportTheme = &v
	return s
}

func (s *ModifyCustomAgentRequest) SetWorkspaceId(v string) *ModifyCustomAgentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ModifyCustomAgentRequest) Validate() error {
	if s.CallbackConfig != nil {
		if err := s.CallbackConfig.Validate(); err != nil {
			return err
		}
	}
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
	if s.KnowledgeSemanticConfigList != nil {
		for _, item := range s.KnowledgeSemanticConfigList {
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

type ModifyCustomAgentRequestCallbackConfig struct {
	CallbackArgs   *string `json:"CallbackArgs,omitempty" xml:"CallbackArgs,omitempty"`
	CallbackPrompt *string `json:"CallbackPrompt,omitempty" xml:"CallbackPrompt,omitempty"`
	CallbackTime   *int32  `json:"CallbackTime,omitempty" xml:"CallbackTime,omitempty"`
	ToolId         *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyCustomAgentRequestCallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentRequestCallbackConfig) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentRequestCallbackConfig) GetCallbackArgs() *string {
	return s.CallbackArgs
}

func (s *ModifyCustomAgentRequestCallbackConfig) GetCallbackPrompt() *string {
	return s.CallbackPrompt
}

func (s *ModifyCustomAgentRequestCallbackConfig) GetCallbackTime() *int32 {
	return s.CallbackTime
}

func (s *ModifyCustomAgentRequestCallbackConfig) GetToolId() *string {
	return s.ToolId
}

func (s *ModifyCustomAgentRequestCallbackConfig) GetType() *string {
	return s.Type
}

func (s *ModifyCustomAgentRequestCallbackConfig) SetCallbackArgs(v string) *ModifyCustomAgentRequestCallbackConfig {
	s.CallbackArgs = &v
	return s
}

func (s *ModifyCustomAgentRequestCallbackConfig) SetCallbackPrompt(v string) *ModifyCustomAgentRequestCallbackConfig {
	s.CallbackPrompt = &v
	return s
}

func (s *ModifyCustomAgentRequestCallbackConfig) SetCallbackTime(v int32) *ModifyCustomAgentRequestCallbackConfig {
	s.CallbackTime = &v
	return s
}

func (s *ModifyCustomAgentRequestCallbackConfig) SetToolId(v string) *ModifyCustomAgentRequestCallbackConfig {
	s.ToolId = &v
	return s
}

func (s *ModifyCustomAgentRequestCallbackConfig) SetType(v string) *ModifyCustomAgentRequestCallbackConfig {
	s.Type = &v
	return s
}

func (s *ModifyCustomAgentRequestCallbackConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyCustomAgentRequestExecutionConfig struct {
	ForbiddenAppendDataSource *bool `json:"ForbiddenAppendDataSource,omitempty" xml:"ForbiddenAppendDataSource,omitempty"`
	// Specifies whether to disable user inquiries during the process.
	//
	// example:
	//
	// true
	SkipAskHuman *bool `json:"SkipAskHuman,omitempty" xml:"SkipAskHuman,omitempty"`
	// Specifies whether to skip the plan confirmation step.
	//
	// example:
	//
	// true
	SkipPlan *bool `json:"SkipPlan,omitempty" xml:"SkipPlan,omitempty"`
	// Specifies whether to skip all SQL confirmations.
	//
	// example:
	//
	// true
	SkipSqlConfirm *bool `json:"SkipSqlConfirm,omitempty" xml:"SkipSqlConfirm,omitempty"`
	// Specifies whether to skip the web report rendering confirmation.
	//
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

func (s *ModifyCustomAgentRequestExecutionConfig) GetForbiddenAppendDataSource() *bool {
	return s.ForbiddenAppendDataSource
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

func (s *ModifyCustomAgentRequestExecutionConfig) SetForbiddenAppendDataSource(v bool) *ModifyCustomAgentRequestExecutionConfig {
	s.ForbiddenAppendDataSource = &v
	return s
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
	// The access type.
	//
	// example:
	//
	// mcp
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	KbUuid     *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// The ID of the MCP server.
	//
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

func (s *ModifyCustomAgentRequestKnowledgeConfigList) GetKbUuid() *string {
	return s.KbUuid
}

func (s *ModifyCustomAgentRequestKnowledgeConfigList) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ModifyCustomAgentRequestKnowledgeConfigList) SetAccessType(v string) *ModifyCustomAgentRequestKnowledgeConfigList {
	s.AccessType = &v
	return s
}

func (s *ModifyCustomAgentRequestKnowledgeConfigList) SetKbUuid(v string) *ModifyCustomAgentRequestKnowledgeConfigList {
	s.KbUuid = &v
	return s
}

func (s *ModifyCustomAgentRequestKnowledgeConfigList) SetMcpServerId(v string) *ModifyCustomAgentRequestKnowledgeConfigList {
	s.McpServerId = &v
	return s
}

func (s *ModifyCustomAgentRequestKnowledgeConfigList) Validate() error {
	return dara.Validate(s)
}

type ModifyCustomAgentRequestKnowledgeSemanticConfigList struct {
	DbId          *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyCustomAgentRequestKnowledgeSemanticConfigList) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentRequestKnowledgeSemanticConfigList) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentRequestKnowledgeSemanticConfigList) GetDbId() *string {
	return s.DbId
}

func (s *ModifyCustomAgentRequestKnowledgeSemanticConfigList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCustomAgentRequestKnowledgeSemanticConfigList) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *ModifyCustomAgentRequestKnowledgeSemanticConfigList) GetType() *string {
	return s.Type
}

func (s *ModifyCustomAgentRequestKnowledgeSemanticConfigList) SetDbId(v string) *ModifyCustomAgentRequestKnowledgeSemanticConfigList {
	s.DbId = &v
	return s
}

func (s *ModifyCustomAgentRequestKnowledgeSemanticConfigList) SetInstanceId(v string) *ModifyCustomAgentRequestKnowledgeSemanticConfigList {
	s.InstanceId = &v
	return s
}

func (s *ModifyCustomAgentRequestKnowledgeSemanticConfigList) SetKnowledgeUuid(v string) *ModifyCustomAgentRequestKnowledgeSemanticConfigList {
	s.KnowledgeUuid = &v
	return s
}

func (s *ModifyCustomAgentRequestKnowledgeSemanticConfigList) SetType(v string) *ModifyCustomAgentRequestKnowledgeSemanticConfigList {
	s.Type = &v
	return s
}

func (s *ModifyCustomAgentRequestKnowledgeSemanticConfigList) Validate() error {
	return dara.Validate(s)
}

type ModifyCustomAgentRequestScheduleTaskConfig struct {
	// The cron expression for the time-based scheduling.
	//
	// example:
	//
	// 0 0 0,1 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The query for the scheduled task.
	//
	// example:
	//
	// Analyze this data and provide a briefing
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of the referenced historical session.
	//
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

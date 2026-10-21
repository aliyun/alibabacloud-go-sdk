// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentType(v string) *CreateCustomAgentRequest
	GetAgentType() *string
	SetCallbackConfig(v *CreateCustomAgentRequestCallbackConfig) *CreateCustomAgentRequest
	GetCallbackConfig() *CreateCustomAgentRequestCallbackConfig
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
	SetKnowledgeSemanticConfigList(v []*CreateCustomAgentRequestKnowledgeSemanticConfigList) *CreateCustomAgentRequest
	GetKnowledgeSemanticConfigList() []*CreateCustomAgentRequestKnowledgeSemanticConfigList
	SetName(v string) *CreateCustomAgentRequest
	GetName() *string
	SetRecommendQuestionList(v []*string) *CreateCustomAgentRequest
	GetRecommendQuestionList() []*string
	SetRelatedSessionId(v string) *CreateCustomAgentRequest
	GetRelatedSessionId() *string
	SetScheduleTaskConfig(v *CreateCustomAgentRequestScheduleTaskConfig) *CreateCustomAgentRequest
	GetScheduleTaskConfig() *CreateCustomAgentRequestScheduleTaskConfig
	SetTextReportConfig(v string) *CreateCustomAgentRequest
	GetTextReportConfig() *string
	SetUserSpecifiedSkillList(v []*string) *CreateCustomAgentRequest
	GetUserSpecifiedSkillList() []*string
	SetWebReportConfig(v string) *CreateCustomAgentRequest
	GetWebReportConfig() *string
	SetWebReportTheme(v string) *CreateCustomAgentRequest
	GetWebReportTheme() *string
	SetWorkspaceId(v string) *CreateCustomAgentRequest
	GetWorkspaceId() *string
}

type CreateCustomAgentRequest struct {
	// The type of the custom agent. Valid values:
	//
	// - Default: the system-built-in custom agent. Each workspace can have at most one Default agent, and it cannot be deleted.
	//
	// - Common: a regular custom agent.
	//
	// example:
	//
	// Common
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The callback feature configuration.
	CallbackConfig *CreateCustomAgentRequestCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The specified data scope, in **JSON character string format**.
	//
	// - General parameter descriptions
	//
	//   - tableFlag: true specifies the data scope.
	//
	//   - scope: personal is a fixed value.
	//
	//   - personal: the parameters for file and database types.
	//
	// **File type**: pass parameters in the following format:
	//
	// - DataSourceType: remote_data_center is a fixed value.
	//
	// - FileId: the file ID.
	//
	// - Database: the database name returned by the ListDataCenterTable operation, which is usually the file name.
	//
	// - Tables: the table names returned by the ListDataCenterTable operation.
	//
	// - TableIds: the table IDs returned by the ListDataCenterTable operation.
	//
	// - RegionId: the current region.
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
	// **Database type**: pass parameters in the following format:
	//
	// - DataSourceType: database is a fixed value.
	//
	// - DmsInstanceId: the DMS instance ID returned by the data center API.
	//
	// - DmsDatabaseId: the DMS database ID returned by the data center API.
	//
	// - FileId: the instance name (**[Deprecated]**).
	//
	// - DbName: the database name returned by the data center API.
	//
	// - Database: the database name returned by the data center API.
	//
	// - Tables: the table names returned by the data center API.
	//
	// - TableIds: the table IDs returned by the data center API.
	//
	// - Engine: the database engine type, such as mysql or postgresql.
	//
	// - RegionId: the current region.
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
	//     "Database" : "Test table******.xlsx",
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
	// Agent test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution configuration.
	ExecutionConfig *CreateCustomAgentRequestExecutionConfig `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
	// The instructions.
	//
	// - Input limit: up to 10,000 characters.
	//
	// example:
	//
	// Core metric definitions:
	//
	// 1. GMV (Gross Merchandise Volume) refers to the total order amount, including both paid and unpaid orders.
	//
	// 2. Order volume refers to the number of valid daily orders.
	//
	// 3. UV (Unique Visitors) refers to the number of deduplicated users visiting the website or app.
	//
	// 4. Conversion rate = Paid orders / UV, reflecting the traffic conversion efficiency
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// The knowledge.
	//
	// - Input limit: up to 10,000 characters.
	//
	// example:
	//
	// Core metric definitions:
	//
	// 1. GMV (Gross Merchandise Volume) refers to the total order amount, including both paid and unpaid orders.
	//
	// 2. Order volume refers to the number of valid daily orders.
	//
	// 3. UV (Unique Visitors) refers to the number of deduplicated users visiting the website or app.
	//
	// 4. Conversion rate = Paid orders / UV, reflecting the traffic conversion efficiency
	Knowledge *string `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	// The external knowledge bases.
	KnowledgeConfigList []*CreateCustomAgentRequestKnowledgeConfigList `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	// The knowledge semantic configuration list.
	KnowledgeSemanticConfigList []*CreateCustomAgentRequestKnowledgeSemanticConfigList `json:"KnowledgeSemanticConfigList,omitempty" xml:"KnowledgeSemanticConfigList,omitempty" type:"Repeated"`
	// The name of the custom agent.
	//
	// example:
	//
	// Agent test name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of recommended questions.
	//
	// - You can configure up to 10 recommended questions.
	RecommendQuestionList []*string `json:"RecommendQuestionList,omitempty" xml:"RecommendQuestionList,omitempty" type:"Repeated"`
	// The ID of the referenced historical session.
	//
	// example:
	//
	// 99crainsjnveo*******vynid
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The scheduled task configuration.
	ScheduleTaskConfig *CreateCustomAgentRequestScheduleTaskConfig `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty" type:"Struct"`
	// The text report format.
	//
	// example:
	//
	// The text report requires that all numbers be written as Chinese numerals instead of Arabic numerals
	TextReportConfig *string `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	// The list of user-specified skills.
	UserSpecifiedSkillList []*string `json:"UserSpecifiedSkillList,omitempty" xml:"UserSpecifiedSkillList,omitempty" type:"Repeated"`
	// The web report format.
	//
	// example:
	//
	// The web report requires that all numbers be written as Chinese numerals instead of Arabic numerals
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

func (s CreateCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequest) GetAgentType() *string {
	return s.AgentType
}

func (s *CreateCustomAgentRequest) GetCallbackConfig() *CreateCustomAgentRequestCallbackConfig {
	return s.CallbackConfig
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

func (s *CreateCustomAgentRequest) GetKnowledgeSemanticConfigList() []*CreateCustomAgentRequestKnowledgeSemanticConfigList {
	return s.KnowledgeSemanticConfigList
}

func (s *CreateCustomAgentRequest) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentRequest) GetRecommendQuestionList() []*string {
	return s.RecommendQuestionList
}

func (s *CreateCustomAgentRequest) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *CreateCustomAgentRequest) GetScheduleTaskConfig() *CreateCustomAgentRequestScheduleTaskConfig {
	return s.ScheduleTaskConfig
}

func (s *CreateCustomAgentRequest) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *CreateCustomAgentRequest) GetUserSpecifiedSkillList() []*string {
	return s.UserSpecifiedSkillList
}

func (s *CreateCustomAgentRequest) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *CreateCustomAgentRequest) GetWebReportTheme() *string {
	return s.WebReportTheme
}

func (s *CreateCustomAgentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCustomAgentRequest) SetAgentType(v string) *CreateCustomAgentRequest {
	s.AgentType = &v
	return s
}

func (s *CreateCustomAgentRequest) SetCallbackConfig(v *CreateCustomAgentRequestCallbackConfig) *CreateCustomAgentRequest {
	s.CallbackConfig = v
	return s
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

func (s *CreateCustomAgentRequest) SetKnowledgeSemanticConfigList(v []*CreateCustomAgentRequestKnowledgeSemanticConfigList) *CreateCustomAgentRequest {
	s.KnowledgeSemanticConfigList = v
	return s
}

func (s *CreateCustomAgentRequest) SetName(v string) *CreateCustomAgentRequest {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentRequest) SetRecommendQuestionList(v []*string) *CreateCustomAgentRequest {
	s.RecommendQuestionList = v
	return s
}

func (s *CreateCustomAgentRequest) SetRelatedSessionId(v string) *CreateCustomAgentRequest {
	s.RelatedSessionId = &v
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

func (s *CreateCustomAgentRequest) SetUserSpecifiedSkillList(v []*string) *CreateCustomAgentRequest {
	s.UserSpecifiedSkillList = v
	return s
}

func (s *CreateCustomAgentRequest) SetWebReportConfig(v string) *CreateCustomAgentRequest {
	s.WebReportConfig = &v
	return s
}

func (s *CreateCustomAgentRequest) SetWebReportTheme(v string) *CreateCustomAgentRequest {
	s.WebReportTheme = &v
	return s
}

func (s *CreateCustomAgentRequest) SetWorkspaceId(v string) *CreateCustomAgentRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCustomAgentRequest) Validate() error {
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

type CreateCustomAgentRequestCallbackConfig struct {
	// The parameters to invoke the callback tool.
	//
	// example:
	//
	// {"ID":"test"}
	CallbackArgs *string `json:"CallbackArgs,omitempty" xml:"CallbackArgs,omitempty"`
	// The prompt for generating the parameters to invoke the callback tool.
	//
	// example:
	//
	// send web report to 123456@163.com
	CallbackPrompt *string `json:"CallbackPrompt,omitempty" xml:"CallbackPrompt,omitempty"`
	// The timing to invoke the callback tool.
	//
	// example:
	//
	// 0
	CallbackTime *int32 `json:"CallbackTime,omitempty" xml:"CallbackTime,omitempty"`
	// The callback tool identifier.
	//
	// example:
	//
	// Mcp serverId (6te...g7a)
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	// The callback tool type.
	//
	// example:
	//
	// Mcp
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCustomAgentRequestCallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestCallbackConfig) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestCallbackConfig) GetCallbackArgs() *string {
	return s.CallbackArgs
}

func (s *CreateCustomAgentRequestCallbackConfig) GetCallbackPrompt() *string {
	return s.CallbackPrompt
}

func (s *CreateCustomAgentRequestCallbackConfig) GetCallbackTime() *int32 {
	return s.CallbackTime
}

func (s *CreateCustomAgentRequestCallbackConfig) GetToolId() *string {
	return s.ToolId
}

func (s *CreateCustomAgentRequestCallbackConfig) GetType() *string {
	return s.Type
}

func (s *CreateCustomAgentRequestCallbackConfig) SetCallbackArgs(v string) *CreateCustomAgentRequestCallbackConfig {
	s.CallbackArgs = &v
	return s
}

func (s *CreateCustomAgentRequestCallbackConfig) SetCallbackPrompt(v string) *CreateCustomAgentRequestCallbackConfig {
	s.CallbackPrompt = &v
	return s
}

func (s *CreateCustomAgentRequestCallbackConfig) SetCallbackTime(v int32) *CreateCustomAgentRequestCallbackConfig {
	s.CallbackTime = &v
	return s
}

func (s *CreateCustomAgentRequestCallbackConfig) SetToolId(v string) *CreateCustomAgentRequestCallbackConfig {
	s.ToolId = &v
	return s
}

func (s *CreateCustomAgentRequestCallbackConfig) SetType(v string) *CreateCustomAgentRequestCallbackConfig {
	s.Type = &v
	return s
}

func (s *CreateCustomAgentRequestCallbackConfig) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestExecutionConfig struct {
	// Specifies whether to prohibit appending data sources in a session.
	//
	// example:
	//
	// true
	ForbiddenAppendDataSource *bool `json:"ForbiddenAppendDataSource,omitempty" xml:"ForbiddenAppendDataSource,omitempty"`
	// Specifies whether to prohibit user inquiries during the process.
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
	// Specifies whether to skip the confirmation for generating a web report.
	//
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

func (s *CreateCustomAgentRequestExecutionConfig) GetForbiddenAppendDataSource() *bool {
	return s.ForbiddenAppendDataSource
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

func (s *CreateCustomAgentRequestExecutionConfig) SetForbiddenAppendDataSource(v bool) *CreateCustomAgentRequestExecutionConfig {
	s.ForbiddenAppendDataSource = &v
	return s
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
	// The access type. If you do not specify this parameter, the default value is mcp. Valid values:
	//
	// - mcp: access through MCP. You must pass the McpServerId parameter.
	//
	// - dms_kb: access through DMS knowledge base. You must pass the KbUuid parameter.
	//
	// example:
	//
	// mcp
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The knowledge base ID.
	//
	// example:
	//
	// kb-HZ-6anhdzx7hddf9*******obx
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// The MCP server ID.
	//
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

func (s *CreateCustomAgentRequestKnowledgeConfigList) GetKbUuid() *string {
	return s.KbUuid
}

func (s *CreateCustomAgentRequestKnowledgeConfigList) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *CreateCustomAgentRequestKnowledgeConfigList) SetAccessType(v string) *CreateCustomAgentRequestKnowledgeConfigList {
	s.AccessType = &v
	return s
}

func (s *CreateCustomAgentRequestKnowledgeConfigList) SetKbUuid(v string) *CreateCustomAgentRequestKnowledgeConfigList {
	s.KbUuid = &v
	return s
}

func (s *CreateCustomAgentRequestKnowledgeConfigList) SetMcpServerId(v string) *CreateCustomAgentRequestKnowledgeConfigList {
	s.McpServerId = &v
	return s
}

func (s *CreateCustomAgentRequestKnowledgeConfigList) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestKnowledgeSemanticConfigList struct {
	// The DMS database ID.
	//
	// example:
	//
	// 81512473
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The DMS instance ID.
	//
	// example:
	//
	// 1532071
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The knowledge UUID.
	//
	// example:
	//
	// 7ca8f50d782*******02e87
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
	// The knowledge semantic type. Valid values:
	//
	// - sql_template: SQL template.
	//
	// - ossie_model: Ossie semantic view.
	//
	// example:
	//
	// sql_template
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCustomAgentRequestKnowledgeSemanticConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequestKnowledgeSemanticConfigList) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentRequestKnowledgeSemanticConfigList) GetDbId() *string {
	return s.DbId
}

func (s *CreateCustomAgentRequestKnowledgeSemanticConfigList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCustomAgentRequestKnowledgeSemanticConfigList) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *CreateCustomAgentRequestKnowledgeSemanticConfigList) GetType() *string {
	return s.Type
}

func (s *CreateCustomAgentRequestKnowledgeSemanticConfigList) SetDbId(v string) *CreateCustomAgentRequestKnowledgeSemanticConfigList {
	s.DbId = &v
	return s
}

func (s *CreateCustomAgentRequestKnowledgeSemanticConfigList) SetInstanceId(v string) *CreateCustomAgentRequestKnowledgeSemanticConfigList {
	s.InstanceId = &v
	return s
}

func (s *CreateCustomAgentRequestKnowledgeSemanticConfigList) SetKnowledgeUuid(v string) *CreateCustomAgentRequestKnowledgeSemanticConfigList {
	s.KnowledgeUuid = &v
	return s
}

func (s *CreateCustomAgentRequestKnowledgeSemanticConfigList) SetType(v string) *CreateCustomAgentRequestKnowledgeSemanticConfigList {
	s.Type = &v
	return s
}

func (s *CreateCustomAgentRequestKnowledgeSemanticConfigList) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentRequestScheduleTaskConfig struct {
	// The cron expression for time-based scheduling.
	//
	// example:
	//
	// 0 0 0 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The query for the scheduled task.
	//
	// example:
	//
	// Analyze this data and provide a brief report
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of the referenced historical session.
	//
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

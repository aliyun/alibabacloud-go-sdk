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
	SetName(v string) *ModifyCustomAgentRequest
	GetName() *string
	SetRelatedSessionId(v string) *ModifyCustomAgentRequest
	GetRelatedSessionId() *string
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
	// The callback configuration.
	CallbackConfig *ModifyCustomAgentRequestCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
	// The ID of the custom agent.
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
	// The data scope for the agent, specified in a **JSON-formatted string**.
	//
	// - General parameters:
	//
	//   - `tableFlag`: Set this to `true` to specify the data scope.
	//
	//   - `scope`: The value must be `personal`.
	//
	//   - `personal`: The parameters for files or databases.
	//
	// **For files**, use the following parameters:
	//
	// - `DataSourceType`: The value must be `remote_data_center`.
	//
	// - `FileId`: The file ID.
	//
	// - `Database`: The database name returned by the `ListDataCenterTable` operation. This is typically the file name.
	//
	// - `Tables`: The table names returned by the `ListDataCenterTable` operation.
	//
	// - `TableIds`: The table IDs returned by the `ListDataCenterTable` operation.
	//
	// - `RegionId`: The current region.
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
	// **For databases**, use the following parameters:
	//
	// - `DataSourceType`: The value must be `database`.
	//
	// - `DmsInstanceId`: The ID of the DMS instance, which is returned by the data center API.
	//
	// - `DmsDatabaseId`: The ID of the DMS database, which is returned by the data center API.
	//
	// - `FileId`: The instance name. This parameter is deprecated.
	//
	// - `DbName`: The database name returned by the data center API.
	//
	// - `Database`: The database name returned by the data center API.
	//
	// - `Tables`: The table names returned by the data center API.
	//
	// - `TableIds`: The table IDs returned by the data center API.
	//
	// - `Engine`: The database engine type. Valid values: `mysql` and `postgresql`.
	//
	// - `RegionId`: The current region.
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
	//     "Database" : "测试表格******.xlsx",
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
	// Agent测试描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The execution configuration.
	ExecutionConfig *ModifyCustomAgentRequestExecutionConfig `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
	// The system instruction for the custom agent.
	//
	// - The maximum length is 10,000 characters.
	//
	// example:
	//
	// 分析框架：
	//
	// 1、需按日、周、月维度监控核心指标（GMV、订单量、UV、转化率），分析趋势变化及同比/环比波动；
	//
	// 2、划分新老客、渠道、地域进行拆解，识别增长来源与短板；
	//
	// 3、结合用户行为路径（浏览→加购→支付）开展漏斗分析，定位流失环节；
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// A text-based knowledge base for the custom agent.
	//
	// - The maximum length is 10,000 characters.
	//
	// example:
	//
	// 核心指标定义：
	//
	// 1、GMV（成交总额）指订单金额总和，含已支付及未支付成功订单；
	//
	// 2、订单量为每日有效下单笔数；
	//
	// 3、UV（独立访客）指访问网站或APP的去重用户数；
	//
	// 4、转化率=支付订单数 / UV，反映流量转化效率；
	Knowledge *string `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	// The configurations for the external knowledge base.
	KnowledgeConfigList []*ModifyCustomAgentRequestKnowledgeConfigList `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	// The name of the custom agent.
	//
	// example:
	//
	// Agent测试名称
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The configuration for the scheduled task.
	ScheduleTaskConfig *ModifyCustomAgentRequestScheduleTaskConfig `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty" type:"Struct"`
	// The formatting instructions for the text report.
	//
	// - The maximum length is 10,000 characters.
	//
	// example:
	//
	// 文字报告要求所有数字不使用阿拉伯数字，全部转为中文数字
	TextReportConfig *string `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	// The formatting instructions for the web report.
	//
	// - The maximum length is 50,000 characters.
	//
	// example:
	//
	// 网页报告要求所有数字不使用阿拉伯数字，全部转为中文数字
	WebReportConfig *string `json:"WebReportConfig,omitempty" xml:"WebReportConfig,omitempty"`
	// The ID of the workspace.
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

func (s *ModifyCustomAgentRequest) GetWebReportConfig() *string {
	return s.WebReportConfig
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

func (s *ModifyCustomAgentRequest) SetWebReportConfig(v string) *ModifyCustomAgentRequest {
	s.WebReportConfig = &v
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
	if s.ScheduleTaskConfig != nil {
		if err := s.ScheduleTaskConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCustomAgentRequestCallbackConfig struct {
	// The arguments for the callback.
	CallbackArgs *string `json:"CallbackArgs,omitempty" xml:"CallbackArgs,omitempty"`
	// The prompt to use for the callback.
	CallbackPrompt *string `json:"CallbackPrompt,omitempty" xml:"CallbackPrompt,omitempty"`
	// The timestamp of the callback.
	CallbackTime *int32 `json:"CallbackTime,omitempty" xml:"CallbackTime,omitempty"`
	// The ID of the tool to call.
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	// The callback type.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// Specifies whether to prevent the agent from asking for user input during execution.
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
	// Specifies whether to skip all SQL confirmation steps.
	//
	// example:
	//
	// true
	SkipSqlConfirm *bool `json:"SkipSqlConfirm,omitempty" xml:"SkipSqlConfirm,omitempty"`
	// Specifies whether to skip the confirmation for web report generation.
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
	// The access type.
	//
	// - `mcp`: Connects via the MCP service.
	//
	// example:
	//
	// mcp
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The UUID of the knowledge base.
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
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

type ModifyCustomAgentRequestScheduleTaskConfig struct {
	// The cron expression for the scheduled task.
	//
	// example:
	//
	// 0 0 0,1 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The query for the scheduled task.
	//
	// example:
	//
	// 分析一下这份数据，给出简报
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of a previous session to use as a reference.
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

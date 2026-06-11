// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentRequest interface {
	dara.Model
	String() string
	GoString() string
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
	SetName(v string) *CreateCustomAgentRequest
	GetName() *string
	SetRelatedSessionId(v string) *CreateCustomAgentRequest
	GetRelatedSessionId() *string
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
	CallbackConfig *CreateCustomAgentRequestCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
	// The ID of the DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The data range, specified as a **JSON string**.
	//
	// - General parameters
	//
	//   - `tableFlag`: Set to `true` to specify a data range.
	//
	//   - `scope`: The value must be `personal`.
	//
	//   - `personal`: Contains the parameters for a file or database.
	//
	// **File type**: Use the following parameters.
	//
	// - `DataSourceType`: The value must be `remote_data_center`.
	//
	// - `FileId`: The ID of the file.
	//
	// - `Database`: The name of the database returned by the `ListDataCenterTable` operation. This is typically the file name.
	//
	// - `Tables`: The names of the tables returned by the `ListDataCenterTable` operation.
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
	// **Database type**: Use the following parameters.
	//
	// - `DataSourceType`: The value must be `database`.
	//
	// - `DmsInstanceId`: The ID of the DMS instance returned by the data center API.
	//
	// - `DmsDatabaseId`: The ID of the DMS database returned by the data center API.
	//
	// - `FileId`: The instance name. This parameter is deprecated.
	//
	// - `DbName`: The name of the database returned by the data center API.
	//
	// - `Database`: The name of the database returned by the data center API.
	//
	// - `Tables`: The names of the tables returned by the data center API.
	//
	// - `TableIds`: The table IDs returned by the data center API.
	//
	// - `Engine`: The database engine. Valid values: `mysql` and `postgresql`.
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
	ExecutionConfig *CreateCustomAgentRequestExecutionConfig `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
	// The instruction for the custom agent.
	//
	// - Maximum length: 10,000 characters.
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
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// The knowledge for the custom agent.
	//
	// - Maximum length: 10,000 characters.
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
	// A list of external knowledge bases.
	KnowledgeConfigList []*CreateCustomAgentRequestKnowledgeConfigList `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	// The name of the custom agent.
	//
	// example:
	//
	// Agent测试名称
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The configuration for the scheduled task.
	ScheduleTaskConfig *CreateCustomAgentRequestScheduleTaskConfig `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty" type:"Struct"`
	// The formatting requirements for the text report.
	//
	// example:
	//
	// 文字报告要求所有数字不使用阿拉伯数字，全部转为中文数字
	TextReportConfig *string `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	// The formatting requirements for the web report.
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

func (s CreateCustomAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentRequest) GoString() string {
	return s.String()
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

func (s *CreateCustomAgentRequest) GetName() *string {
	return s.Name
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

func (s *CreateCustomAgentRequest) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *CreateCustomAgentRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
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

func (s *CreateCustomAgentRequest) SetName(v string) *CreateCustomAgentRequest {
	s.Name = &v
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

func (s *CreateCustomAgentRequest) SetWebReportConfig(v string) *CreateCustomAgentRequest {
	s.WebReportConfig = &v
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
	if s.ScheduleTaskConfig != nil {
		if err := s.ScheduleTaskConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomAgentRequestCallbackConfig struct {
	CallbackArgs   *string `json:"CallbackArgs,omitempty" xml:"CallbackArgs,omitempty"`
	CallbackPrompt *string `json:"CallbackPrompt,omitempty" xml:"CallbackPrompt,omitempty"`
	CallbackTime   *int32  `json:"CallbackTime,omitempty" xml:"CallbackTime,omitempty"`
	ToolId         *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// Specifies whether to skip asking the user for input during execution.
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
	// Specifies whether to skip the web report confirmation step.
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
	// The access type.
	//
	// - `mcp`: Access via an MCP server.
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

type CreateCustomAgentRequestScheduleTaskConfig struct {
	// The cron expression for the scheduled task.
	//
	// example:
	//
	// 0 0 0 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The query for the scheduled task.
	//
	// example:
	//
	// 分析一下这份数据，给出简报
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of a previous session to use for reference.
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

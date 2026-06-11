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
	SetName(v string) *ModifyCustomAgentShrinkRequest
	GetName() *string
	SetRelatedSessionId(v string) *ModifyCustomAgentShrinkRequest
	GetRelatedSessionId() *string
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
	// The callback configuration.
	CallbackConfigShrink *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
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
	ExecutionConfigShrink *string `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty"`
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
	KnowledgeConfigListShrink *string `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty"`
	// The name of the custom agent.
	//
	// example:
	//
	// Agent测试名称
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The configuration for the scheduled task.
	ScheduleTaskConfigShrink *string `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty"`
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

func (s *ModifyCustomAgentShrinkRequest) GetWebReportConfig() *string {
	return s.WebReportConfig
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

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
	SetWorkspaceId(v string) *CreateCustomAgentShrinkRequest
	GetWorkspaceId() *string
}

type CreateCustomAgentShrinkRequest struct {
	CallbackConfigShrink *string `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty"`
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
	ExecutionConfigShrink *string `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty"`
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

func (s *CreateCustomAgentShrinkRequest) SetWorkspaceId(v string) *CreateCustomAgentShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCustomAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}

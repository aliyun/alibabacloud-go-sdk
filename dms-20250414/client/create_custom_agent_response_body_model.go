// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateCustomAgentResponseBodyData) *CreateCustomAgentResponseBody
	GetData() *CreateCustomAgentResponseBodyData
	SetErrorCode(v string) *CreateCustomAgentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateCustomAgentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateCustomAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCustomAgentResponseBody
	GetSuccess() *bool
}

type CreateCustomAgentResponseBody struct {
	// The response data.
	Data *CreateCustomAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBody) GetData() *CreateCustomAgentResponseBodyData {
	return s.Data
}

func (s *CreateCustomAgentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateCustomAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomAgentResponseBody) SetData(v *CreateCustomAgentResponseBodyData) *CreateCustomAgentResponseBody {
	s.Data = v
	return s
}

func (s *CreateCustomAgentResponseBody) SetErrorCode(v string) *CreateCustomAgentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetErrorMessage(v string) *CreateCustomAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetRequestId(v string) *CreateCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomAgentResponseBody) SetSuccess(v bool) *CreateCustomAgentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomAgentResponseBodyData struct {
	// The primary Alibaba Cloud account ID.
	//
	// example:
	//
	// 16738266********
	AliyunParentUid *string `json:"AliyunParentUid,omitempty" xml:"AliyunParentUid,omitempty"`
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 20372822********
	AliyunUid      *string                                          `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	CallbackConfig *CreateCustomAgentResponseBodyDataCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
	// The username of the creator.
	//
	// example:
	//
	// HaoY*****
	CreatorUserName *string `json:"CreatorUserName,omitempty" xml:"CreatorUserName,omitempty"`
	// The custom agent ID.
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
	// The specified data scope, in a JSON string format.
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
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	// The execution configuration.
	ExecutionConfig *CreateCustomAgentResponseBodyDataExecutionConfig `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
	// The creation time.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The instruction for the agent.
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
	// Indicates whether a scheduled task is configured for the agent.
	//
	// example:
	//
	// false
	IsScheduleTask *bool `json:"IsScheduleTask,omitempty" xml:"IsScheduleTask,omitempty"`
	// The knowledge for the agent.
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
	KnowledgeConfigList []*CreateCustomAgentResponseBodyDataKnowledgeConfigList `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	// The ID of the user who last modified the agent.
	//
	// example:
	//
	// 20372822********
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The username of the user who last modified the agent.
	//
	// example:
	//
	// HaoY*****
	ModifierUserName *string `json:"ModifierUserName,omitempty" xml:"ModifierUserName,omitempty"`
	// The name of the custom agent.
	//
	// example:
	//
	// Agent测试名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The next run time of the scheduled task, formatted as a UNIX timestamp.
	//
	// - Timestamp format
	//
	// example:
	//
	// 1767715200
	NextRuntime *int64 `json:"NextRuntime,omitempty" xml:"NextRuntime,omitempty"`
	// The offline time.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	OfflineTime *string `json:"OfflineTime,omitempty" xml:"OfflineTime,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The release time.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The configuration of the scheduled task.
	ScheduleTaskConfig *CreateCustomAgentResponseBodyDataScheduleTaskConfig `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty" type:"Struct"`
	// The status of the custom agent.
	//
	// example:
	//
	// NEW
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The format of the text report.
	//
	// example:
	//
	// 文字报告要求所有数字不使用阿拉伯数字，全部转为中文数字
	TextReportConfig *string `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	// The format of the web report.
	//
	// example:
	//
	// 网页报告要求所有数字不使用阿拉伯数字，全部转为中文数字
	WebReportConfig *string `json:"WebReportConfig,omitempty" xml:"WebReportConfig,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateCustomAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyData) GetAliyunParentUid() *string {
	return s.AliyunParentUid
}

func (s *CreateCustomAgentResponseBodyData) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *CreateCustomAgentResponseBodyData) GetCallbackConfig() *CreateCustomAgentResponseBodyDataCallbackConfig {
	return s.CallbackConfig
}

func (s *CreateCustomAgentResponseBodyData) GetCreatorUserName() *string {
	return s.CreatorUserName
}

func (s *CreateCustomAgentResponseBodyData) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *CreateCustomAgentResponseBodyData) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *CreateCustomAgentResponseBodyData) GetDataJson() *string {
	return s.DataJson
}

func (s *CreateCustomAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomAgentResponseBodyData) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *CreateCustomAgentResponseBodyData) GetExecutionConfig() *CreateCustomAgentResponseBodyDataExecutionConfig {
	return s.ExecutionConfig
}

func (s *CreateCustomAgentResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CreateCustomAgentResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateCustomAgentResponseBodyData) GetInstruction() *string {
	return s.Instruction
}

func (s *CreateCustomAgentResponseBodyData) GetIsScheduleTask() *bool {
	return s.IsScheduleTask
}

func (s *CreateCustomAgentResponseBodyData) GetKnowledge() *string {
	return s.Knowledge
}

func (s *CreateCustomAgentResponseBodyData) GetKnowledgeConfigList() []*CreateCustomAgentResponseBodyDataKnowledgeConfigList {
	return s.KnowledgeConfigList
}

func (s *CreateCustomAgentResponseBodyData) GetModifier() *string {
	return s.Modifier
}

func (s *CreateCustomAgentResponseBodyData) GetModifierUserName() *string {
	return s.ModifierUserName
}

func (s *CreateCustomAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateCustomAgentResponseBodyData) GetNextRuntime() *int64 {
	return s.NextRuntime
}

func (s *CreateCustomAgentResponseBodyData) GetOfflineTime() *string {
	return s.OfflineTime
}

func (s *CreateCustomAgentResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *CreateCustomAgentResponseBodyData) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *CreateCustomAgentResponseBodyData) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *CreateCustomAgentResponseBodyData) GetScheduleTaskConfig() *CreateCustomAgentResponseBodyDataScheduleTaskConfig {
	return s.ScheduleTaskConfig
}

func (s *CreateCustomAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateCustomAgentResponseBodyData) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *CreateCustomAgentResponseBodyData) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *CreateCustomAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCustomAgentResponseBodyData) SetAliyunParentUid(v string) *CreateCustomAgentResponseBodyData {
	s.AliyunParentUid = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetAliyunUid(v string) *CreateCustomAgentResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetCallbackConfig(v *CreateCustomAgentResponseBodyDataCallbackConfig) *CreateCustomAgentResponseBodyData {
	s.CallbackConfig = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetCreatorUserName(v string) *CreateCustomAgentResponseBodyData {
	s.CreatorUserName = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetCustomAgentId(v string) *CreateCustomAgentResponseBodyData {
	s.CustomAgentId = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetDMSUnit(v string) *CreateCustomAgentResponseBodyData {
	s.DMSUnit = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetDataJson(v string) *CreateCustomAgentResponseBodyData {
	s.DataJson = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetDescription(v string) *CreateCustomAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetDmsUnit(v string) *CreateCustomAgentResponseBodyData {
	s.DmsUnit = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetExecutionConfig(v *CreateCustomAgentResponseBodyDataExecutionConfig) *CreateCustomAgentResponseBodyData {
	s.ExecutionConfig = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetGmtCreated(v string) *CreateCustomAgentResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetGmtModified(v string) *CreateCustomAgentResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetInstruction(v string) *CreateCustomAgentResponseBodyData {
	s.Instruction = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetIsScheduleTask(v bool) *CreateCustomAgentResponseBodyData {
	s.IsScheduleTask = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetKnowledge(v string) *CreateCustomAgentResponseBodyData {
	s.Knowledge = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetKnowledgeConfigList(v []*CreateCustomAgentResponseBodyDataKnowledgeConfigList) *CreateCustomAgentResponseBodyData {
	s.KnowledgeConfigList = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetModifier(v string) *CreateCustomAgentResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetModifierUserName(v string) *CreateCustomAgentResponseBodyData {
	s.ModifierUserName = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetName(v string) *CreateCustomAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetNextRuntime(v int64) *CreateCustomAgentResponseBodyData {
	s.NextRuntime = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetOfflineTime(v string) *CreateCustomAgentResponseBodyData {
	s.OfflineTime = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetRegion(v string) *CreateCustomAgentResponseBodyData {
	s.Region = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetRelatedSessionId(v string) *CreateCustomAgentResponseBodyData {
	s.RelatedSessionId = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetReleaseTime(v string) *CreateCustomAgentResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetScheduleTaskConfig(v *CreateCustomAgentResponseBodyDataScheduleTaskConfig) *CreateCustomAgentResponseBodyData {
	s.ScheduleTaskConfig = v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetStatus(v string) *CreateCustomAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetTextReportConfig(v string) *CreateCustomAgentResponseBodyData {
	s.TextReportConfig = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetWebReportConfig(v string) *CreateCustomAgentResponseBodyData {
	s.WebReportConfig = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) SetWorkspaceId(v string) *CreateCustomAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCustomAgentResponseBodyData) Validate() error {
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

type CreateCustomAgentResponseBodyDataCallbackConfig struct {
	CallbackArgs   *string `json:"CallbackArgs,omitempty" xml:"CallbackArgs,omitempty"`
	CallbackPrompt *string `json:"CallbackPrompt,omitempty" xml:"CallbackPrompt,omitempty"`
	CallbackTime   *int32  `json:"CallbackTime,omitempty" xml:"CallbackTime,omitempty"`
	ToolId         *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataCallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataCallbackConfig) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataCallbackConfig) GetCallbackArgs() *string {
	return s.CallbackArgs
}

func (s *CreateCustomAgentResponseBodyDataCallbackConfig) GetCallbackPrompt() *string {
	return s.CallbackPrompt
}

func (s *CreateCustomAgentResponseBodyDataCallbackConfig) GetCallbackTime() *int32 {
	return s.CallbackTime
}

func (s *CreateCustomAgentResponseBodyDataCallbackConfig) GetToolId() *string {
	return s.ToolId
}

func (s *CreateCustomAgentResponseBodyDataCallbackConfig) GetType() *string {
	return s.Type
}

func (s *CreateCustomAgentResponseBodyDataCallbackConfig) SetCallbackArgs(v string) *CreateCustomAgentResponseBodyDataCallbackConfig {
	s.CallbackArgs = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataCallbackConfig) SetCallbackPrompt(v string) *CreateCustomAgentResponseBodyDataCallbackConfig {
	s.CallbackPrompt = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataCallbackConfig) SetCallbackTime(v int32) *CreateCustomAgentResponseBodyDataCallbackConfig {
	s.CallbackTime = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataCallbackConfig) SetToolId(v string) *CreateCustomAgentResponseBodyDataCallbackConfig {
	s.ToolId = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataCallbackConfig) SetType(v string) *CreateCustomAgentResponseBodyDataCallbackConfig {
	s.Type = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataCallbackConfig) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataExecutionConfig struct {
	// Indicates whether to prevent the agent from asking the user for clarification during execution.
	//
	// example:
	//
	// true
	SkipAskHuman *bool `json:"SkipAskHuman,omitempty" xml:"SkipAskHuman,omitempty"`
	// Indicates whether to skip the plan confirmation step.
	//
	// example:
	//
	// true
	SkipPlan *bool `json:"SkipPlan,omitempty" xml:"SkipPlan,omitempty"`
	// Indicates whether to skip all SQL confirmations.
	//
	// example:
	//
	// true
	SkipSqlConfirm *bool `json:"SkipSqlConfirm,omitempty" xml:"SkipSqlConfirm,omitempty"`
	// Indicates whether to skip the web report generation confirmation.
	//
	// example:
	//
	// true
	SkipWebReportConfirm *bool `json:"SkipWebReportConfirm,omitempty" xml:"SkipWebReportConfirm,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataExecutionConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataExecutionConfig) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataExecutionConfig) GetSkipAskHuman() *bool {
	return s.SkipAskHuman
}

func (s *CreateCustomAgentResponseBodyDataExecutionConfig) GetSkipPlan() *bool {
	return s.SkipPlan
}

func (s *CreateCustomAgentResponseBodyDataExecutionConfig) GetSkipSqlConfirm() *bool {
	return s.SkipSqlConfirm
}

func (s *CreateCustomAgentResponseBodyDataExecutionConfig) GetSkipWebReportConfirm() *bool {
	return s.SkipWebReportConfirm
}

func (s *CreateCustomAgentResponseBodyDataExecutionConfig) SetSkipAskHuman(v bool) *CreateCustomAgentResponseBodyDataExecutionConfig {
	s.SkipAskHuman = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataExecutionConfig) SetSkipPlan(v bool) *CreateCustomAgentResponseBodyDataExecutionConfig {
	s.SkipPlan = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataExecutionConfig) SetSkipSqlConfirm(v bool) *CreateCustomAgentResponseBodyDataExecutionConfig {
	s.SkipSqlConfirm = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataExecutionConfig) SetSkipWebReportConfirm(v bool) *CreateCustomAgentResponseBodyDataExecutionConfig {
	s.SkipWebReportConfirm = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataExecutionConfig) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataKnowledgeConfigList struct {
	// The access method.
	//
	// - `mcp`: Connects using MCP.
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

func (s CreateCustomAgentResponseBodyDataKnowledgeConfigList) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataKnowledgeConfigList) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataKnowledgeConfigList) GetAccessType() *string {
	return s.AccessType
}

func (s *CreateCustomAgentResponseBodyDataKnowledgeConfigList) GetKbUuid() *string {
	return s.KbUuid
}

func (s *CreateCustomAgentResponseBodyDataKnowledgeConfigList) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *CreateCustomAgentResponseBodyDataKnowledgeConfigList) SetAccessType(v string) *CreateCustomAgentResponseBodyDataKnowledgeConfigList {
	s.AccessType = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataKnowledgeConfigList) SetKbUuid(v string) *CreateCustomAgentResponseBodyDataKnowledgeConfigList {
	s.KbUuid = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataKnowledgeConfigList) SetMcpServerId(v string) *CreateCustomAgentResponseBodyDataKnowledgeConfigList {
	s.McpServerId = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataKnowledgeConfigList) Validate() error {
	return dara.Validate(s)
}

type CreateCustomAgentResponseBodyDataScheduleTaskConfig struct {
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
	// The ID of a previous session to be used as a template for the scheduled task.
	//
	// example:
	//
	// 4m24*****mg7j2v
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
}

func (s CreateCustomAgentResponseBodyDataScheduleTaskConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentResponseBodyDataScheduleTaskConfig) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentResponseBodyDataScheduleTaskConfig) GetCronExpression() *string {
	return s.CronExpression
}

func (s *CreateCustomAgentResponseBodyDataScheduleTaskConfig) GetQuery() *string {
	return s.Query
}

func (s *CreateCustomAgentResponseBodyDataScheduleTaskConfig) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *CreateCustomAgentResponseBodyDataScheduleTaskConfig) SetCronExpression(v string) *CreateCustomAgentResponseBodyDataScheduleTaskConfig {
	s.CronExpression = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataScheduleTaskConfig) SetQuery(v string) *CreateCustomAgentResponseBodyDataScheduleTaskConfig {
	s.Query = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataScheduleTaskConfig) SetRelatedSessionId(v string) *CreateCustomAgentResponseBodyDataScheduleTaskConfig {
	s.RelatedSessionId = &v
	return s
}

func (s *CreateCustomAgentResponseBodyDataScheduleTaskConfig) Validate() error {
	return dara.Validate(s)
}

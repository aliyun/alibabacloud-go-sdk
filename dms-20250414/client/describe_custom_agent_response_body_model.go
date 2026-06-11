// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeCustomAgentResponseBodyData) *DescribeCustomAgentResponseBody
	GetData() *DescribeCustomAgentResponseBodyData
	SetErrorCode(v string) *DescribeCustomAgentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeCustomAgentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DescribeCustomAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCustomAgentResponseBody
	GetSuccess() *bool
}

type DescribeCustomAgentResponseBody struct {
	// The details of the custom agent.
	Data *DescribeCustomAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request fails.
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
	// The ID of the request.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentResponseBody) GetData() *DescribeCustomAgentResponseBodyData {
	return s.Data
}

func (s *DescribeCustomAgentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeCustomAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCustomAgentResponseBody) SetData(v *DescribeCustomAgentResponseBodyData) *DescribeCustomAgentResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomAgentResponseBody) SetErrorCode(v string) *DescribeCustomAgentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCustomAgentResponseBody) SetErrorMessage(v string) *DescribeCustomAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeCustomAgentResponseBody) SetRequestId(v string) *DescribeCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomAgentResponseBody) SetSuccess(v bool) *DescribeCustomAgentResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCustomAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCustomAgentResponseBodyData struct {
	// The ID of the parent Alibaba Cloud account.
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
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The callback configuration.
	CallbackConfig *DescribeCustomAgentResponseBodyDataCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
	// The username of the creator.
	//
	// example:
	//
	// HaoY*****
	CreatorUserName *string `json:"CreatorUserName,omitempty" xml:"CreatorUserName,omitempty"`
	// The ID of the custom agent.
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
	// The data scope, formatted as a JSON string.
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
	// Indicates whether this is the default agent.
	DefaultAgent *int32 `json:"DefaultAgent,omitempty" xml:"DefaultAgent,omitempty"`
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
	ExecutionConfig *DescribeCustomAgentResponseBodyDataExecutionConfig `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
	// The time when the agent was created.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the agent was last modified.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The instruction for the agent\\"s analysis.
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
	// Indicates whether a scheduled task is configured.
	//
	// example:
	//
	// true
	IsScheduleTask *bool `json:"IsScheduleTask,omitempty" xml:"IsScheduleTask,omitempty"`
	// The domain knowledge for the agent.
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
	// The knowledge configurations.
	KnowledgeConfigList []*DescribeCustomAgentResponseBodyDataKnowledgeConfigList `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	// The ID of the user who last modified the agent.
	//
	// example:
	//
	// 20372822********
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The username of the modifier.
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
	// The next scheduled execution time.
	//
	// example:
	//
	// 1767715200
	NextRuntime *int64 `json:"NextRuntime,omitempty" xml:"NextRuntime,omitempty"`
	// The time when the agent was taken offline.
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
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the related historical session.
	//
	// example:
	//
	// 5xyz...
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The release time.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The configuration of the scheduled task.
	ScheduleTaskConfig *DescribeCustomAgentResponseBodyDataScheduleTaskConfig `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty" type:"Struct"`
	// The status of the custom agent. Valid values:
	//
	// example:
	//
	// RELEASED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s DescribeCustomAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentResponseBodyData) GetAliyunParentUid() *string {
	return s.AliyunParentUid
}

func (s *DescribeCustomAgentResponseBodyData) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *DescribeCustomAgentResponseBodyData) GetCallbackConfig() *DescribeCustomAgentResponseBodyDataCallbackConfig {
	return s.CallbackConfig
}

func (s *DescribeCustomAgentResponseBodyData) GetCreatorUserName() *string {
	return s.CreatorUserName
}

func (s *DescribeCustomAgentResponseBodyData) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *DescribeCustomAgentResponseBodyData) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *DescribeCustomAgentResponseBodyData) GetDataJson() *string {
	return s.DataJson
}

func (s *DescribeCustomAgentResponseBodyData) GetDefaultAgent() *int32 {
	return s.DefaultAgent
}

func (s *DescribeCustomAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeCustomAgentResponseBodyData) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *DescribeCustomAgentResponseBodyData) GetExecutionConfig() *DescribeCustomAgentResponseBodyDataExecutionConfig {
	return s.ExecutionConfig
}

func (s *DescribeCustomAgentResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeCustomAgentResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCustomAgentResponseBodyData) GetInstruction() *string {
	return s.Instruction
}

func (s *DescribeCustomAgentResponseBodyData) GetIsScheduleTask() *bool {
	return s.IsScheduleTask
}

func (s *DescribeCustomAgentResponseBodyData) GetKnowledge() *string {
	return s.Knowledge
}

func (s *DescribeCustomAgentResponseBodyData) GetKnowledgeConfigList() []*DescribeCustomAgentResponseBodyDataKnowledgeConfigList {
	return s.KnowledgeConfigList
}

func (s *DescribeCustomAgentResponseBodyData) GetModifier() *string {
	return s.Modifier
}

func (s *DescribeCustomAgentResponseBodyData) GetModifierUserName() *string {
	return s.ModifierUserName
}

func (s *DescribeCustomAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeCustomAgentResponseBodyData) GetNextRuntime() *int64 {
	return s.NextRuntime
}

func (s *DescribeCustomAgentResponseBodyData) GetOfflineTime() *string {
	return s.OfflineTime
}

func (s *DescribeCustomAgentResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *DescribeCustomAgentResponseBodyData) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *DescribeCustomAgentResponseBodyData) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *DescribeCustomAgentResponseBodyData) GetScheduleTaskConfig() *DescribeCustomAgentResponseBodyDataScheduleTaskConfig {
	return s.ScheduleTaskConfig
}

func (s *DescribeCustomAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeCustomAgentResponseBodyData) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *DescribeCustomAgentResponseBodyData) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *DescribeCustomAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeCustomAgentResponseBodyData) SetAliyunParentUid(v string) *DescribeCustomAgentResponseBodyData {
	s.AliyunParentUid = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetAliyunUid(v string) *DescribeCustomAgentResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetCallbackConfig(v *DescribeCustomAgentResponseBodyDataCallbackConfig) *DescribeCustomAgentResponseBodyData {
	s.CallbackConfig = v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetCreatorUserName(v string) *DescribeCustomAgentResponseBodyData {
	s.CreatorUserName = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetCustomAgentId(v string) *DescribeCustomAgentResponseBodyData {
	s.CustomAgentId = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetDMSUnit(v string) *DescribeCustomAgentResponseBodyData {
	s.DMSUnit = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetDataJson(v string) *DescribeCustomAgentResponseBodyData {
	s.DataJson = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetDefaultAgent(v int32) *DescribeCustomAgentResponseBodyData {
	s.DefaultAgent = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetDescription(v string) *DescribeCustomAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetDmsUnit(v string) *DescribeCustomAgentResponseBodyData {
	s.DmsUnit = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetExecutionConfig(v *DescribeCustomAgentResponseBodyDataExecutionConfig) *DescribeCustomAgentResponseBodyData {
	s.ExecutionConfig = v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetGmtCreated(v string) *DescribeCustomAgentResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetGmtModified(v string) *DescribeCustomAgentResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetInstruction(v string) *DescribeCustomAgentResponseBodyData {
	s.Instruction = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetIsScheduleTask(v bool) *DescribeCustomAgentResponseBodyData {
	s.IsScheduleTask = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetKnowledge(v string) *DescribeCustomAgentResponseBodyData {
	s.Knowledge = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetKnowledgeConfigList(v []*DescribeCustomAgentResponseBodyDataKnowledgeConfigList) *DescribeCustomAgentResponseBodyData {
	s.KnowledgeConfigList = v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetModifier(v string) *DescribeCustomAgentResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetModifierUserName(v string) *DescribeCustomAgentResponseBodyData {
	s.ModifierUserName = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetName(v string) *DescribeCustomAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetNextRuntime(v int64) *DescribeCustomAgentResponseBodyData {
	s.NextRuntime = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetOfflineTime(v string) *DescribeCustomAgentResponseBodyData {
	s.OfflineTime = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetRegion(v string) *DescribeCustomAgentResponseBodyData {
	s.Region = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetRelatedSessionId(v string) *DescribeCustomAgentResponseBodyData {
	s.RelatedSessionId = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetReleaseTime(v string) *DescribeCustomAgentResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetScheduleTaskConfig(v *DescribeCustomAgentResponseBodyDataScheduleTaskConfig) *DescribeCustomAgentResponseBodyData {
	s.ScheduleTaskConfig = v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetStatus(v string) *DescribeCustomAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetTextReportConfig(v string) *DescribeCustomAgentResponseBodyData {
	s.TextReportConfig = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetWebReportConfig(v string) *DescribeCustomAgentResponseBodyData {
	s.WebReportConfig = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetWorkspaceId(v string) *DescribeCustomAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) Validate() error {
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

type DescribeCustomAgentResponseBodyDataCallbackConfig struct {
	// The callback arguments.
	CallbackArgs *string `json:"CallbackArgs,omitempty" xml:"CallbackArgs,omitempty"`
	// The callback prompt.
	CallbackPrompt *string `json:"CallbackPrompt,omitempty" xml:"CallbackPrompt,omitempty"`
	// The callback time.
	CallbackTime *int32 `json:"CallbackTime,omitempty" xml:"CallbackTime,omitempty"`
	// The tool ID.
	ToolId *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	// The callback type.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCustomAgentResponseBodyDataCallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentResponseBodyDataCallbackConfig) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentResponseBodyDataCallbackConfig) GetCallbackArgs() *string {
	return s.CallbackArgs
}

func (s *DescribeCustomAgentResponseBodyDataCallbackConfig) GetCallbackPrompt() *string {
	return s.CallbackPrompt
}

func (s *DescribeCustomAgentResponseBodyDataCallbackConfig) GetCallbackTime() *int32 {
	return s.CallbackTime
}

func (s *DescribeCustomAgentResponseBodyDataCallbackConfig) GetToolId() *string {
	return s.ToolId
}

func (s *DescribeCustomAgentResponseBodyDataCallbackConfig) GetType() *string {
	return s.Type
}

func (s *DescribeCustomAgentResponseBodyDataCallbackConfig) SetCallbackArgs(v string) *DescribeCustomAgentResponseBodyDataCallbackConfig {
	s.CallbackArgs = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataCallbackConfig) SetCallbackPrompt(v string) *DescribeCustomAgentResponseBodyDataCallbackConfig {
	s.CallbackPrompt = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataCallbackConfig) SetCallbackTime(v int32) *DescribeCustomAgentResponseBodyDataCallbackConfig {
	s.CallbackTime = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataCallbackConfig) SetToolId(v string) *DescribeCustomAgentResponseBodyDataCallbackConfig {
	s.ToolId = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataCallbackConfig) SetType(v string) *DescribeCustomAgentResponseBodyDataCallbackConfig {
	s.Type = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataCallbackConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomAgentResponseBodyDataExecutionConfig struct {
	// Specifies whether to disable prompts that require human intervention.
	//
	// example:
	//
	// false
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
	// Specifies whether to skip confirmation before a web report is generated.
	//
	// example:
	//
	// false
	SkipWebReportConfirm *bool `json:"SkipWebReportConfirm,omitempty" xml:"SkipWebReportConfirm,omitempty"`
}

func (s DescribeCustomAgentResponseBodyDataExecutionConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentResponseBodyDataExecutionConfig) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentResponseBodyDataExecutionConfig) GetSkipAskHuman() *bool {
	return s.SkipAskHuman
}

func (s *DescribeCustomAgentResponseBodyDataExecutionConfig) GetSkipPlan() *bool {
	return s.SkipPlan
}

func (s *DescribeCustomAgentResponseBodyDataExecutionConfig) GetSkipSqlConfirm() *bool {
	return s.SkipSqlConfirm
}

func (s *DescribeCustomAgentResponseBodyDataExecutionConfig) GetSkipWebReportConfirm() *bool {
	return s.SkipWebReportConfirm
}

func (s *DescribeCustomAgentResponseBodyDataExecutionConfig) SetSkipAskHuman(v bool) *DescribeCustomAgentResponseBodyDataExecutionConfig {
	s.SkipAskHuman = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataExecutionConfig) SetSkipPlan(v bool) *DescribeCustomAgentResponseBodyDataExecutionConfig {
	s.SkipPlan = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataExecutionConfig) SetSkipSqlConfirm(v bool) *DescribeCustomAgentResponseBodyDataExecutionConfig {
	s.SkipSqlConfirm = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataExecutionConfig) SetSkipWebReportConfirm(v bool) *DescribeCustomAgentResponseBodyDataExecutionConfig {
	s.SkipWebReportConfirm = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataExecutionConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomAgentResponseBodyDataKnowledgeConfigList struct {
	// The access type.
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The UUID of the knowledge base.
	KbUuid *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// The ID of the MCP server.
	McpServerId *string `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
}

func (s DescribeCustomAgentResponseBodyDataKnowledgeConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentResponseBodyDataKnowledgeConfigList) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeConfigList) GetAccessType() *string {
	return s.AccessType
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeConfigList) GetKbUuid() *string {
	return s.KbUuid
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeConfigList) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeConfigList) SetAccessType(v string) *DescribeCustomAgentResponseBodyDataKnowledgeConfigList {
	s.AccessType = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeConfigList) SetKbUuid(v string) *DescribeCustomAgentResponseBodyDataKnowledgeConfigList {
	s.KbUuid = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeConfigList) SetMcpServerId(v string) *DescribeCustomAgentResponseBodyDataKnowledgeConfigList {
	s.McpServerId = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomAgentResponseBodyDataScheduleTaskConfig struct {
	// The cron expression for the scheduled task.
	//
	// example:
	//
	// 0 0 0 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The task query.
	//
	// example:
	//
	// 分析一下这份数据，给出简报
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of the related historical session.
	//
	// example:
	//
	// 4m24*****mg7j2v
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
}

func (s DescribeCustomAgentResponseBodyDataScheduleTaskConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentResponseBodyDataScheduleTaskConfig) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentResponseBodyDataScheduleTaskConfig) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeCustomAgentResponseBodyDataScheduleTaskConfig) GetQuery() *string {
	return s.Query
}

func (s *DescribeCustomAgentResponseBodyDataScheduleTaskConfig) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *DescribeCustomAgentResponseBodyDataScheduleTaskConfig) SetCronExpression(v string) *DescribeCustomAgentResponseBodyDataScheduleTaskConfig {
	s.CronExpression = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataScheduleTaskConfig) SetQuery(v string) *DescribeCustomAgentResponseBodyDataScheduleTaskConfig {
	s.Query = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataScheduleTaskConfig) SetRelatedSessionId(v string) *DescribeCustomAgentResponseBodyDataScheduleTaskConfig {
	s.RelatedSessionId = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataScheduleTaskConfig) Validate() error {
	return dara.Validate(s)
}

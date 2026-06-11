// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListCustomAgentResponseBodyData) *ListCustomAgentResponseBody
	GetData() *ListCustomAgentResponseBodyData
	SetErrorCode(v string) *ListCustomAgentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListCustomAgentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListCustomAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomAgentResponseBody
	GetSuccess() *bool
}

type ListCustomAgentResponseBody struct {
	// The returned data.
	Data *ListCustomAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomAgentResponseBody) GetData() *ListCustomAgentResponseBodyData {
	return s.Data
}

func (s *ListCustomAgentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListCustomAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomAgentResponseBody) SetData(v *ListCustomAgentResponseBodyData) *ListCustomAgentResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomAgentResponseBody) SetErrorCode(v string) *ListCustomAgentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListCustomAgentResponseBody) SetErrorMessage(v string) *ListCustomAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListCustomAgentResponseBody) SetRequestId(v string) *ListCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomAgentResponseBody) SetSuccess(v bool) *ListCustomAgentResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomAgentResponseBodyData struct {
	// A list of custom agent objects.
	Content []*ListCustomAgentResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 5
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int64 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListCustomAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomAgentResponseBodyData) GetContent() []*ListCustomAgentResponseBodyDataContent {
	return s.Content
}

func (s *ListCustomAgentResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCustomAgentResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCustomAgentResponseBodyData) GetTotalElements() *int64 {
	return s.TotalElements
}

func (s *ListCustomAgentResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *ListCustomAgentResponseBodyData) SetContent(v []*ListCustomAgentResponseBodyDataContent) *ListCustomAgentResponseBodyData {
	s.Content = v
	return s
}

func (s *ListCustomAgentResponseBodyData) SetPageNumber(v int64) *ListCustomAgentResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCustomAgentResponseBodyData) SetPageSize(v int64) *ListCustomAgentResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCustomAgentResponseBodyData) SetTotalElements(v int64) *ListCustomAgentResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListCustomAgentResponseBodyData) SetTotalPages(v int64) *ListCustomAgentResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListCustomAgentResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomAgentResponseBodyDataContent struct {
	// The parent Alibaba Cloud account ID.
	//
	// example:
	//
	// 16738266********
	AliyunParentId *string `json:"AliyunParentId,omitempty" xml:"AliyunParentId,omitempty"`
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 20372822********
	AliyunUid      *string                                               `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	CallbackConfig *ListCustomAgentResponseBodyDataContentCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
	// The name of the creator.
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
	// The data scope, specified as a JSON string.
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
	DataJson     *string `json:"DataJson,omitempty" xml:"DataJson,omitempty"`
	DefaultAgent *int32  `json:"DefaultAgent,omitempty" xml:"DefaultAgent,omitempty"`
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
	ExecutionConfig *ListCustomAgentResponseBodyDataContentExecutionConfig `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
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
	// The instructions.
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
	// Indicates whether the agent is configured with a recurring task.
	//
	// example:
	//
	// true
	IsScheduleTask *bool `json:"IsScheduleTask,omitempty" xml:"IsScheduleTask,omitempty"`
	// The provided knowledge.
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
	Knowledge           *string                                                      `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	KnowledgeConfigList []*ListCustomAgentResponseBodyDataContentKnowledgeConfigList `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	// The user who last modified the agent.
	//
	// example:
	//
	// 20372822********
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The name of the user who last modified the agent.
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
	// If a recurring task is configured, this indicates its next scheduled runtime.
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
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The time when the agent was published.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The recurring task configuration.
	ScheduleTaskConfig *ListCustomAgentResponseBodyDataContentScheduleTaskConfig `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty" type:"Struct"`
	// The status of the custom agent.
	//
	// example:
	//
	// RELEASED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The formatting instructions for the text report.
	//
	// example:
	//
	// 文字报告要求所有数字不使用阿拉伯数字，全部转为中文数字
	TextReportConfig *string `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	// The formatting instructions for the web report.
	//
	// example:
	//
	// 网页报告要求所有数字不使用阿拉伯数字，全部转为中文数字
	WebReportConfig *string `json:"WebReportConfig,omitempty" xml:"WebReportConfig,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListCustomAgentResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListCustomAgentResponseBodyDataContent) GetAliyunParentId() *string {
	return s.AliyunParentId
}

func (s *ListCustomAgentResponseBodyDataContent) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ListCustomAgentResponseBodyDataContent) GetCallbackConfig() *ListCustomAgentResponseBodyDataContentCallbackConfig {
	return s.CallbackConfig
}

func (s *ListCustomAgentResponseBodyDataContent) GetCreatorUserName() *string {
	return s.CreatorUserName
}

func (s *ListCustomAgentResponseBodyDataContent) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *ListCustomAgentResponseBodyDataContent) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *ListCustomAgentResponseBodyDataContent) GetDataJson() *string {
	return s.DataJson
}

func (s *ListCustomAgentResponseBodyDataContent) GetDefaultAgent() *int32 {
	return s.DefaultAgent
}

func (s *ListCustomAgentResponseBodyDataContent) GetDescription() *string {
	return s.Description
}

func (s *ListCustomAgentResponseBodyDataContent) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *ListCustomAgentResponseBodyDataContent) GetExecutionConfig() *ListCustomAgentResponseBodyDataContentExecutionConfig {
	return s.ExecutionConfig
}

func (s *ListCustomAgentResponseBodyDataContent) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListCustomAgentResponseBodyDataContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListCustomAgentResponseBodyDataContent) GetInstruction() *string {
	return s.Instruction
}

func (s *ListCustomAgentResponseBodyDataContent) GetIsScheduleTask() *bool {
	return s.IsScheduleTask
}

func (s *ListCustomAgentResponseBodyDataContent) GetKnowledge() *string {
	return s.Knowledge
}

func (s *ListCustomAgentResponseBodyDataContent) GetKnowledgeConfigList() []*ListCustomAgentResponseBodyDataContentKnowledgeConfigList {
	return s.KnowledgeConfigList
}

func (s *ListCustomAgentResponseBodyDataContent) GetModifier() *string {
	return s.Modifier
}

func (s *ListCustomAgentResponseBodyDataContent) GetModifierUserName() *string {
	return s.ModifierUserName
}

func (s *ListCustomAgentResponseBodyDataContent) GetName() *string {
	return s.Name
}

func (s *ListCustomAgentResponseBodyDataContent) GetNextRuntime() *int64 {
	return s.NextRuntime
}

func (s *ListCustomAgentResponseBodyDataContent) GetOfflineTime() *string {
	return s.OfflineTime
}

func (s *ListCustomAgentResponseBodyDataContent) GetRegion() *string {
	return s.Region
}

func (s *ListCustomAgentResponseBodyDataContent) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *ListCustomAgentResponseBodyDataContent) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *ListCustomAgentResponseBodyDataContent) GetScheduleTaskConfig() *ListCustomAgentResponseBodyDataContentScheduleTaskConfig {
	return s.ScheduleTaskConfig
}

func (s *ListCustomAgentResponseBodyDataContent) GetStatus() *string {
	return s.Status
}

func (s *ListCustomAgentResponseBodyDataContent) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *ListCustomAgentResponseBodyDataContent) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *ListCustomAgentResponseBodyDataContent) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListCustomAgentResponseBodyDataContent) SetAliyunParentId(v string) *ListCustomAgentResponseBodyDataContent {
	s.AliyunParentId = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetAliyunUid(v string) *ListCustomAgentResponseBodyDataContent {
	s.AliyunUid = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetCallbackConfig(v *ListCustomAgentResponseBodyDataContentCallbackConfig) *ListCustomAgentResponseBodyDataContent {
	s.CallbackConfig = v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetCreatorUserName(v string) *ListCustomAgentResponseBodyDataContent {
	s.CreatorUserName = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetCustomAgentId(v string) *ListCustomAgentResponseBodyDataContent {
	s.CustomAgentId = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetDMSUnit(v string) *ListCustomAgentResponseBodyDataContent {
	s.DMSUnit = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetDataJson(v string) *ListCustomAgentResponseBodyDataContent {
	s.DataJson = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetDefaultAgent(v int32) *ListCustomAgentResponseBodyDataContent {
	s.DefaultAgent = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetDescription(v string) *ListCustomAgentResponseBodyDataContent {
	s.Description = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetDmsUnit(v string) *ListCustomAgentResponseBodyDataContent {
	s.DmsUnit = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetExecutionConfig(v *ListCustomAgentResponseBodyDataContentExecutionConfig) *ListCustomAgentResponseBodyDataContent {
	s.ExecutionConfig = v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetGmtCreated(v string) *ListCustomAgentResponseBodyDataContent {
	s.GmtCreated = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetGmtModified(v string) *ListCustomAgentResponseBodyDataContent {
	s.GmtModified = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetInstruction(v string) *ListCustomAgentResponseBodyDataContent {
	s.Instruction = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetIsScheduleTask(v bool) *ListCustomAgentResponseBodyDataContent {
	s.IsScheduleTask = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetKnowledge(v string) *ListCustomAgentResponseBodyDataContent {
	s.Knowledge = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetKnowledgeConfigList(v []*ListCustomAgentResponseBodyDataContentKnowledgeConfigList) *ListCustomAgentResponseBodyDataContent {
	s.KnowledgeConfigList = v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetModifier(v string) *ListCustomAgentResponseBodyDataContent {
	s.Modifier = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetModifierUserName(v string) *ListCustomAgentResponseBodyDataContent {
	s.ModifierUserName = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetName(v string) *ListCustomAgentResponseBodyDataContent {
	s.Name = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetNextRuntime(v int64) *ListCustomAgentResponseBodyDataContent {
	s.NextRuntime = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetOfflineTime(v string) *ListCustomAgentResponseBodyDataContent {
	s.OfflineTime = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetRegion(v string) *ListCustomAgentResponseBodyDataContent {
	s.Region = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetRelatedSessionId(v string) *ListCustomAgentResponseBodyDataContent {
	s.RelatedSessionId = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetReleaseTime(v string) *ListCustomAgentResponseBodyDataContent {
	s.ReleaseTime = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetScheduleTaskConfig(v *ListCustomAgentResponseBodyDataContentScheduleTaskConfig) *ListCustomAgentResponseBodyDataContent {
	s.ScheduleTaskConfig = v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetStatus(v string) *ListCustomAgentResponseBodyDataContent {
	s.Status = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetTextReportConfig(v string) *ListCustomAgentResponseBodyDataContent {
	s.TextReportConfig = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetWebReportConfig(v string) *ListCustomAgentResponseBodyDataContent {
	s.WebReportConfig = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetWorkspaceId(v string) *ListCustomAgentResponseBodyDataContent {
	s.WorkspaceId = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) Validate() error {
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

type ListCustomAgentResponseBodyDataContentCallbackConfig struct {
	CallbackArgs   *string `json:"CallbackArgs,omitempty" xml:"CallbackArgs,omitempty"`
	CallbackPrompt *string `json:"CallbackPrompt,omitempty" xml:"CallbackPrompt,omitempty"`
	CallbackTime   *int32  `json:"CallbackTime,omitempty" xml:"CallbackTime,omitempty"`
	ToolId         *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCustomAgentResponseBodyDataContentCallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentResponseBodyDataContentCallbackConfig) GoString() string {
	return s.String()
}

func (s *ListCustomAgentResponseBodyDataContentCallbackConfig) GetCallbackArgs() *string {
	return s.CallbackArgs
}

func (s *ListCustomAgentResponseBodyDataContentCallbackConfig) GetCallbackPrompt() *string {
	return s.CallbackPrompt
}

func (s *ListCustomAgentResponseBodyDataContentCallbackConfig) GetCallbackTime() *int32 {
	return s.CallbackTime
}

func (s *ListCustomAgentResponseBodyDataContentCallbackConfig) GetToolId() *string {
	return s.ToolId
}

func (s *ListCustomAgentResponseBodyDataContentCallbackConfig) GetType() *string {
	return s.Type
}

func (s *ListCustomAgentResponseBodyDataContentCallbackConfig) SetCallbackArgs(v string) *ListCustomAgentResponseBodyDataContentCallbackConfig {
	s.CallbackArgs = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentCallbackConfig) SetCallbackPrompt(v string) *ListCustomAgentResponseBodyDataContentCallbackConfig {
	s.CallbackPrompt = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentCallbackConfig) SetCallbackTime(v int32) *ListCustomAgentResponseBodyDataContentCallbackConfig {
	s.CallbackTime = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentCallbackConfig) SetToolId(v string) *ListCustomAgentResponseBodyDataContentCallbackConfig {
	s.ToolId = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentCallbackConfig) SetType(v string) *ListCustomAgentResponseBodyDataContentCallbackConfig {
	s.Type = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentCallbackConfig) Validate() error {
	return dara.Validate(s)
}

type ListCustomAgentResponseBodyDataContentExecutionConfig struct {
	// Indicates whether to prevent user inquiries during the process.
	//
	// example:
	//
	// false
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
	// Indicates whether to skip the confirmation for generating a web report.
	//
	// example:
	//
	// false
	SkipWebReportConfirm *bool `json:"SkipWebReportConfirm,omitempty" xml:"SkipWebReportConfirm,omitempty"`
}

func (s ListCustomAgentResponseBodyDataContentExecutionConfig) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentResponseBodyDataContentExecutionConfig) GoString() string {
	return s.String()
}

func (s *ListCustomAgentResponseBodyDataContentExecutionConfig) GetSkipAskHuman() *bool {
	return s.SkipAskHuman
}

func (s *ListCustomAgentResponseBodyDataContentExecutionConfig) GetSkipPlan() *bool {
	return s.SkipPlan
}

func (s *ListCustomAgentResponseBodyDataContentExecutionConfig) GetSkipSqlConfirm() *bool {
	return s.SkipSqlConfirm
}

func (s *ListCustomAgentResponseBodyDataContentExecutionConfig) GetSkipWebReportConfirm() *bool {
	return s.SkipWebReportConfirm
}

func (s *ListCustomAgentResponseBodyDataContentExecutionConfig) SetSkipAskHuman(v bool) *ListCustomAgentResponseBodyDataContentExecutionConfig {
	s.SkipAskHuman = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentExecutionConfig) SetSkipPlan(v bool) *ListCustomAgentResponseBodyDataContentExecutionConfig {
	s.SkipPlan = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentExecutionConfig) SetSkipSqlConfirm(v bool) *ListCustomAgentResponseBodyDataContentExecutionConfig {
	s.SkipSqlConfirm = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentExecutionConfig) SetSkipWebReportConfirm(v bool) *ListCustomAgentResponseBodyDataContentExecutionConfig {
	s.SkipWebReportConfirm = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentExecutionConfig) Validate() error {
	return dara.Validate(s)
}

type ListCustomAgentResponseBodyDataContentKnowledgeConfigList struct {
	AccessType  *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	KbUuid      *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	McpServerId *string `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
}

func (s ListCustomAgentResponseBodyDataContentKnowledgeConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentResponseBodyDataContentKnowledgeConfigList) GoString() string {
	return s.String()
}

func (s *ListCustomAgentResponseBodyDataContentKnowledgeConfigList) GetAccessType() *string {
	return s.AccessType
}

func (s *ListCustomAgentResponseBodyDataContentKnowledgeConfigList) GetKbUuid() *string {
	return s.KbUuid
}

func (s *ListCustomAgentResponseBodyDataContentKnowledgeConfigList) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ListCustomAgentResponseBodyDataContentKnowledgeConfigList) SetAccessType(v string) *ListCustomAgentResponseBodyDataContentKnowledgeConfigList {
	s.AccessType = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentKnowledgeConfigList) SetKbUuid(v string) *ListCustomAgentResponseBodyDataContentKnowledgeConfigList {
	s.KbUuid = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentKnowledgeConfigList) SetMcpServerId(v string) *ListCustomAgentResponseBodyDataContentKnowledgeConfigList {
	s.McpServerId = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentKnowledgeConfigList) Validate() error {
	return dara.Validate(s)
}

type ListCustomAgentResponseBodyDataContentScheduleTaskConfig struct {
	// The cron expression for the recurring task.
	//
	// example:
	//
	// 0 0 0 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The query for the recurring task.
	//
	// example:
	//
	// 分析一下这份数据，给出简报
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The ID of the referenced historical session.
	//
	// example:
	//
	// 4m24*****mg7j2v
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
}

func (s ListCustomAgentResponseBodyDataContentScheduleTaskConfig) String() string {
	return dara.Prettify(s)
}

func (s ListCustomAgentResponseBodyDataContentScheduleTaskConfig) GoString() string {
	return s.String()
}

func (s *ListCustomAgentResponseBodyDataContentScheduleTaskConfig) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ListCustomAgentResponseBodyDataContentScheduleTaskConfig) GetQuery() *string {
	return s.Query
}

func (s *ListCustomAgentResponseBodyDataContentScheduleTaskConfig) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *ListCustomAgentResponseBodyDataContentScheduleTaskConfig) SetCronExpression(v string) *ListCustomAgentResponseBodyDataContentScheduleTaskConfig {
	s.CronExpression = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentScheduleTaskConfig) SetQuery(v string) *ListCustomAgentResponseBodyDataContentScheduleTaskConfig {
	s.Query = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentScheduleTaskConfig) SetRelatedSessionId(v string) *ListCustomAgentResponseBodyDataContentScheduleTaskConfig {
	s.RelatedSessionId = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentScheduleTaskConfig) Validate() error {
	return dara.Validate(s)
}

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
	// The response struct.
	Data *DescribeCustomAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
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
	// The Alibaba Cloud account ID of the parent account.
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
	AliyunUid      *string                                            `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	CallbackConfig *DescribeCustomAgentResponseBodyDataCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
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
	// The specified data scope in JSON string format.
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
	//     "Database" : "Test spreadsheet******.xlsx",
	//
	//     "Tables" : [ "Sheet1" ],
	//
	//     "TableIds" : [ "******" ],
	//
	//     "RegionId" : "ap-southeast-1"
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
	// Agent test description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	// The execution configuration.
	ExecutionConfig *DescribeCustomAgentResponseBodyDataExecutionConfig `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
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
	// The instruction.
	//
	// example:
	//
	// Analysis framework:
	//
	// 1. Monitor core metrics (GMV, order volume, UV, conversion rate) by day, week, and month dimensions, and analyze trends and year-over-year/month-over-month fluctuations;
	//
	// 2. Segment by new/existing customers, channels, and regions to identify growth sources and weaknesses;
	//
	// 3. Conduct funnel analysis based on user behavior paths (browse → add to cart → payment) to identify drop-off points;
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// Specifies whether a periodic task is configured.
	//
	// example:
	//
	// true
	IsScheduleTask *bool `json:"IsScheduleTask,omitempty" xml:"IsScheduleTask,omitempty"`
	// The knowledge.
	//
	// example:
	//
	// Core metric definitions:
	//
	// 1. GMV (Gross Merchandise Volume) refers to the total order amount, including paid and unpaid orders;
	//
	// 2. Order volume is the number of valid orders placed per day;
	//
	// 3. UV (Unique Visitors) refers to the deduplicated number of users who visit the website or app;
	//
	// 4. Conversion rate = number of paid orders / UV, reflecting traffic conversion efficiency;
	Knowledge                   *string                                                           `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	KnowledgeConfigList         []*DescribeCustomAgentResponseBodyDataKnowledgeConfigList         `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	KnowledgeSemanticConfigList []*DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList `json:"KnowledgeSemanticConfigList,omitempty" xml:"KnowledgeSemanticConfigList,omitempty" type:"Repeated"`
	// The modifier.
	//
	// example:
	//
	// 20372822********
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The name of the modifier.
	//
	// example:
	//
	// HaoY*****
	ModifierUserName *string `json:"ModifierUserName,omitempty" xml:"ModifierUserName,omitempty"`
	// The name of the custom agent.
	//
	// example:
	//
	// Agent test name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The next run time of the periodic task.
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
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The referenced historical session ID.
	//
	// example:
	//
	// 5xyz...
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The publish time.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The periodic task configuration.
	ScheduleTaskConfig *DescribeCustomAgentResponseBodyDataScheduleTaskConfig `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty" type:"Struct"`
	// The status of the custom agent.
	//
	// example:
	//
	// RELEASED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The text report format.
	//
	// example:
	//
	// The text report requires all numbers to be converted from Arabic numerals to Chinese numerals
	TextReportConfig *string `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	// The web report format.
	//
	// example:
	//
	// The web report requires all numbers to be converted from Arabic numerals to Chinese numerals
	WebReportConfig *string `json:"WebReportConfig,omitempty" xml:"WebReportConfig,omitempty"`
	WebReportTheme  *string `json:"WebReportTheme,omitempty" xml:"WebReportTheme,omitempty"`
	// The workspace ID.
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

func (s *DescribeCustomAgentResponseBodyData) GetKnowledgeSemanticConfigList() []*DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList {
	return s.KnowledgeSemanticConfigList
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

func (s *DescribeCustomAgentResponseBodyData) GetWebReportTheme() *string {
	return s.WebReportTheme
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

func (s *DescribeCustomAgentResponseBodyData) SetKnowledgeSemanticConfigList(v []*DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) *DescribeCustomAgentResponseBodyData {
	s.KnowledgeSemanticConfigList = v
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

func (s *DescribeCustomAgentResponseBodyData) SetWebReportTheme(v string) *DescribeCustomAgentResponseBodyData {
	s.WebReportTheme = &v
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

type DescribeCustomAgentResponseBodyDataCallbackConfig struct {
	CallbackArgs   *string `json:"CallbackArgs,omitempty" xml:"CallbackArgs,omitempty"`
	CallbackPrompt *string `json:"CallbackPrompt,omitempty" xml:"CallbackPrompt,omitempty"`
	CallbackTime   *int32  `json:"CallbackTime,omitempty" xml:"CallbackTime,omitempty"`
	ToolId         *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// Specifies whether to disable user inquiries during the process.
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
	// Specifies whether to skip the web report drawing confirmation.
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
	AccessType  *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	KbUuid      *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
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

type DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList struct {
	DbId          *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) GoString() string {
	return s.String()
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) GetDbId() *string {
	return s.DbId
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) GetType() *string {
	return s.Type
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) SetDbId(v string) *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList {
	s.DbId = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) SetInstanceId(v string) *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList {
	s.InstanceId = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) SetKnowledgeUuid(v string) *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList {
	s.KnowledgeUuid = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) SetType(v string) *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList {
	s.Type = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeSemanticConfigList) Validate() error {
	return dara.Validate(s)
}

type DescribeCustomAgentResponseBodyDataScheduleTaskConfig struct {
	// The cron expression for timed scheduling.
	//
	// example:
	//
	// 0 0 0 ? 	- 1-7
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The query of the periodic task.
	//
	// example:
	//
	// Analyze this data and provide a brief report
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The referenced historical session ID.
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

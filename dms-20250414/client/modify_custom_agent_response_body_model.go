// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyCustomAgentResponseBodyData) *ModifyCustomAgentResponseBody
	GetData() *ModifyCustomAgentResponseBodyData
	SetErrorCode(v string) *ModifyCustomAgentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ModifyCustomAgentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ModifyCustomAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyCustomAgentResponseBody
	GetSuccess() *bool
}

type ModifyCustomAgentResponseBody struct {
	// The returned struct.
	Data *ModifyCustomAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyCustomAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentResponseBody) GetData() *ModifyCustomAgentResponseBodyData {
	return s.Data
}

func (s *ModifyCustomAgentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModifyCustomAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyCustomAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCustomAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyCustomAgentResponseBody) SetData(v *ModifyCustomAgentResponseBodyData) *ModifyCustomAgentResponseBody {
	s.Data = v
	return s
}

func (s *ModifyCustomAgentResponseBody) SetErrorCode(v string) *ModifyCustomAgentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyCustomAgentResponseBody) SetErrorMessage(v string) *ModifyCustomAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyCustomAgentResponseBody) SetRequestId(v string) *ModifyCustomAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCustomAgentResponseBody) SetSuccess(v bool) *ModifyCustomAgentResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyCustomAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCustomAgentResponseBodyData struct {
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
	AliyunUid      *string                                          `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	CallbackConfig *ModifyCustomAgentResponseBodyDataCallbackConfig `json:"CallbackConfig,omitempty" xml:"CallbackConfig,omitempty" type:"Struct"`
	// The creator name.
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
	// The specified data range in JSON string format.
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
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DmsUnit *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	// The execution configuration.
	ExecutionConfig *ModifyCustomAgentResponseBodyDataExecutionConfig `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
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
	// 1. Monitor core metrics (GMV, order volume, UV, conversion rate) by day, week, and month dimensions, and analyze trends and year-over-year/month-over-month fluctuations.
	//
	// 2. Segment by new/existing customers, channels, and regions to identify growth sources and weaknesses.
	//
	// 3. Conduct funnel analysis based on user behavior paths (browse → add to cart → payment) to identify drop-off points
	Instruction *string `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	// Indicates whether a periodic task is configured.
	//
	// example:
	//
	// false
	IsScheduleTask *bool `json:"IsScheduleTask,omitempty" xml:"IsScheduleTask,omitempty"`
	// The knowledge.
	//
	// example:
	//
	// Core metric definitions:
	//
	// 1. GMV (Gross Merchandise Volume) refers to the total order amount, including paid and unpaid orders.
	//
	// 2. Order volume is the number of valid orders placed per day.
	//
	// 3. UV (Unique Visitors) refers to the deduplicated number of users who visit the website or app.
	//
	// 4. Conversion rate = number of paid orders / UV, reflecting traffic conversion efficiency
	Knowledge *string `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	// The external knowledge base list.
	KnowledgeConfigList         []*ModifyCustomAgentResponseBodyDataKnowledgeConfigList         `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	KnowledgeSemanticConfigList []*ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList `json:"KnowledgeSemanticConfigList,omitempty" xml:"KnowledgeSemanticConfigList,omitempty" type:"Repeated"`
	// The modifier.
	//
	// example:
	//
	// 20372822********
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The modifier name.
	//
	// example:
	//
	// HaoY*****
	ModifierUserName *string `json:"ModifierUserName,omitempty" xml:"ModifierUserName,omitempty"`
	// The name of the custom agent.
	//
	// example:
	//
	// AgentTestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The next run time of the periodic task.
	//
	// - In timestamp format.
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
	// The ID of the referenced historical session.
	RelatedSessionId *string `json:"RelatedSessionId,omitempty" xml:"RelatedSessionId,omitempty"`
	// The publish time.
	//
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The periodic task configuration.
	ScheduleTaskConfig *ModifyCustomAgentResponseBodyDataScheduleTaskConfig `json:"ScheduleTaskConfig,omitempty" xml:"ScheduleTaskConfig,omitempty" type:"Struct"`
	// The custom agent status.
	//
	// example:
	//
	// RELEASED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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

func (s ModifyCustomAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentResponseBodyData) GetAliyunParentUid() *string {
	return s.AliyunParentUid
}

func (s *ModifyCustomAgentResponseBodyData) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ModifyCustomAgentResponseBodyData) GetCallbackConfig() *ModifyCustomAgentResponseBodyDataCallbackConfig {
	return s.CallbackConfig
}

func (s *ModifyCustomAgentResponseBodyData) GetCreatorUserName() *string {
	return s.CreatorUserName
}

func (s *ModifyCustomAgentResponseBodyData) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *ModifyCustomAgentResponseBodyData) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *ModifyCustomAgentResponseBodyData) GetDataJson() *string {
	return s.DataJson
}

func (s *ModifyCustomAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ModifyCustomAgentResponseBodyData) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *ModifyCustomAgentResponseBodyData) GetExecutionConfig() *ModifyCustomAgentResponseBodyDataExecutionConfig {
	return s.ExecutionConfig
}

func (s *ModifyCustomAgentResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ModifyCustomAgentResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ModifyCustomAgentResponseBodyData) GetInstruction() *string {
	return s.Instruction
}

func (s *ModifyCustomAgentResponseBodyData) GetIsScheduleTask() *bool {
	return s.IsScheduleTask
}

func (s *ModifyCustomAgentResponseBodyData) GetKnowledge() *string {
	return s.Knowledge
}

func (s *ModifyCustomAgentResponseBodyData) GetKnowledgeConfigList() []*ModifyCustomAgentResponseBodyDataKnowledgeConfigList {
	return s.KnowledgeConfigList
}

func (s *ModifyCustomAgentResponseBodyData) GetKnowledgeSemanticConfigList() []*ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList {
	return s.KnowledgeSemanticConfigList
}

func (s *ModifyCustomAgentResponseBodyData) GetModifier() *string {
	return s.Modifier
}

func (s *ModifyCustomAgentResponseBodyData) GetModifierUserName() *string {
	return s.ModifierUserName
}

func (s *ModifyCustomAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ModifyCustomAgentResponseBodyData) GetNextRuntime() *int64 {
	return s.NextRuntime
}

func (s *ModifyCustomAgentResponseBodyData) GetOfflineTime() *string {
	return s.OfflineTime
}

func (s *ModifyCustomAgentResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *ModifyCustomAgentResponseBodyData) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *ModifyCustomAgentResponseBodyData) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *ModifyCustomAgentResponseBodyData) GetScheduleTaskConfig() *ModifyCustomAgentResponseBodyDataScheduleTaskConfig {
	return s.ScheduleTaskConfig
}

func (s *ModifyCustomAgentResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ModifyCustomAgentResponseBodyData) GetTextReportConfig() *string {
	return s.TextReportConfig
}

func (s *ModifyCustomAgentResponseBodyData) GetUserSpecifiedSkillList() []*string {
	return s.UserSpecifiedSkillList
}

func (s *ModifyCustomAgentResponseBodyData) GetWebReportConfig() *string {
	return s.WebReportConfig
}

func (s *ModifyCustomAgentResponseBodyData) GetWebReportTheme() *string {
	return s.WebReportTheme
}

func (s *ModifyCustomAgentResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ModifyCustomAgentResponseBodyData) SetAliyunParentUid(v string) *ModifyCustomAgentResponseBodyData {
	s.AliyunParentUid = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetAliyunUid(v string) *ModifyCustomAgentResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetCallbackConfig(v *ModifyCustomAgentResponseBodyDataCallbackConfig) *ModifyCustomAgentResponseBodyData {
	s.CallbackConfig = v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetCreatorUserName(v string) *ModifyCustomAgentResponseBodyData {
	s.CreatorUserName = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetCustomAgentId(v string) *ModifyCustomAgentResponseBodyData {
	s.CustomAgentId = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetDMSUnit(v string) *ModifyCustomAgentResponseBodyData {
	s.DMSUnit = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetDataJson(v string) *ModifyCustomAgentResponseBodyData {
	s.DataJson = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetDescription(v string) *ModifyCustomAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetDmsUnit(v string) *ModifyCustomAgentResponseBodyData {
	s.DmsUnit = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetExecutionConfig(v *ModifyCustomAgentResponseBodyDataExecutionConfig) *ModifyCustomAgentResponseBodyData {
	s.ExecutionConfig = v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetGmtCreated(v string) *ModifyCustomAgentResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetGmtModified(v string) *ModifyCustomAgentResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetInstruction(v string) *ModifyCustomAgentResponseBodyData {
	s.Instruction = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetIsScheduleTask(v bool) *ModifyCustomAgentResponseBodyData {
	s.IsScheduleTask = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetKnowledge(v string) *ModifyCustomAgentResponseBodyData {
	s.Knowledge = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetKnowledgeConfigList(v []*ModifyCustomAgentResponseBodyDataKnowledgeConfigList) *ModifyCustomAgentResponseBodyData {
	s.KnowledgeConfigList = v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetKnowledgeSemanticConfigList(v []*ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) *ModifyCustomAgentResponseBodyData {
	s.KnowledgeSemanticConfigList = v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetModifier(v string) *ModifyCustomAgentResponseBodyData {
	s.Modifier = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetModifierUserName(v string) *ModifyCustomAgentResponseBodyData {
	s.ModifierUserName = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetName(v string) *ModifyCustomAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetNextRuntime(v int64) *ModifyCustomAgentResponseBodyData {
	s.NextRuntime = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetOfflineTime(v string) *ModifyCustomAgentResponseBodyData {
	s.OfflineTime = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetRegion(v string) *ModifyCustomAgentResponseBodyData {
	s.Region = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetRelatedSessionId(v string) *ModifyCustomAgentResponseBodyData {
	s.RelatedSessionId = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetReleaseTime(v string) *ModifyCustomAgentResponseBodyData {
	s.ReleaseTime = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetScheduleTaskConfig(v *ModifyCustomAgentResponseBodyDataScheduleTaskConfig) *ModifyCustomAgentResponseBodyData {
	s.ScheduleTaskConfig = v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetStatus(v string) *ModifyCustomAgentResponseBodyData {
	s.Status = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetTextReportConfig(v string) *ModifyCustomAgentResponseBodyData {
	s.TextReportConfig = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetUserSpecifiedSkillList(v []*string) *ModifyCustomAgentResponseBodyData {
	s.UserSpecifiedSkillList = v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetWebReportConfig(v string) *ModifyCustomAgentResponseBodyData {
	s.WebReportConfig = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetWebReportTheme(v string) *ModifyCustomAgentResponseBodyData {
	s.WebReportTheme = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) SetWorkspaceId(v string) *ModifyCustomAgentResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyData) Validate() error {
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

type ModifyCustomAgentResponseBodyDataCallbackConfig struct {
	CallbackArgs   *string `json:"CallbackArgs,omitempty" xml:"CallbackArgs,omitempty"`
	CallbackPrompt *string `json:"CallbackPrompt,omitempty" xml:"CallbackPrompt,omitempty"`
	CallbackTime   *int32  `json:"CallbackTime,omitempty" xml:"CallbackTime,omitempty"`
	ToolId         *string `json:"ToolId,omitempty" xml:"ToolId,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyCustomAgentResponseBodyDataCallbackConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentResponseBodyDataCallbackConfig) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentResponseBodyDataCallbackConfig) GetCallbackArgs() *string {
	return s.CallbackArgs
}

func (s *ModifyCustomAgentResponseBodyDataCallbackConfig) GetCallbackPrompt() *string {
	return s.CallbackPrompt
}

func (s *ModifyCustomAgentResponseBodyDataCallbackConfig) GetCallbackTime() *int32 {
	return s.CallbackTime
}

func (s *ModifyCustomAgentResponseBodyDataCallbackConfig) GetToolId() *string {
	return s.ToolId
}

func (s *ModifyCustomAgentResponseBodyDataCallbackConfig) GetType() *string {
	return s.Type
}

func (s *ModifyCustomAgentResponseBodyDataCallbackConfig) SetCallbackArgs(v string) *ModifyCustomAgentResponseBodyDataCallbackConfig {
	s.CallbackArgs = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataCallbackConfig) SetCallbackPrompt(v string) *ModifyCustomAgentResponseBodyDataCallbackConfig {
	s.CallbackPrompt = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataCallbackConfig) SetCallbackTime(v int32) *ModifyCustomAgentResponseBodyDataCallbackConfig {
	s.CallbackTime = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataCallbackConfig) SetToolId(v string) *ModifyCustomAgentResponseBodyDataCallbackConfig {
	s.ToolId = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataCallbackConfig) SetType(v string) *ModifyCustomAgentResponseBodyDataCallbackConfig {
	s.Type = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataCallbackConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyCustomAgentResponseBodyDataExecutionConfig struct {
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

func (s ModifyCustomAgentResponseBodyDataExecutionConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentResponseBodyDataExecutionConfig) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentResponseBodyDataExecutionConfig) GetForbiddenAppendDataSource() *bool {
	return s.ForbiddenAppendDataSource
}

func (s *ModifyCustomAgentResponseBodyDataExecutionConfig) GetSkipAskHuman() *bool {
	return s.SkipAskHuman
}

func (s *ModifyCustomAgentResponseBodyDataExecutionConfig) GetSkipPlan() *bool {
	return s.SkipPlan
}

func (s *ModifyCustomAgentResponseBodyDataExecutionConfig) GetSkipSqlConfirm() *bool {
	return s.SkipSqlConfirm
}

func (s *ModifyCustomAgentResponseBodyDataExecutionConfig) GetSkipWebReportConfirm() *bool {
	return s.SkipWebReportConfirm
}

func (s *ModifyCustomAgentResponseBodyDataExecutionConfig) SetForbiddenAppendDataSource(v bool) *ModifyCustomAgentResponseBodyDataExecutionConfig {
	s.ForbiddenAppendDataSource = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataExecutionConfig) SetSkipAskHuman(v bool) *ModifyCustomAgentResponseBodyDataExecutionConfig {
	s.SkipAskHuman = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataExecutionConfig) SetSkipPlan(v bool) *ModifyCustomAgentResponseBodyDataExecutionConfig {
	s.SkipPlan = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataExecutionConfig) SetSkipSqlConfirm(v bool) *ModifyCustomAgentResponseBodyDataExecutionConfig {
	s.SkipSqlConfirm = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataExecutionConfig) SetSkipWebReportConfirm(v bool) *ModifyCustomAgentResponseBodyDataExecutionConfig {
	s.SkipWebReportConfirm = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataExecutionConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyCustomAgentResponseBodyDataKnowledgeConfigList struct {
	// The access type. Valid values:
	//
	// - mcp: accessed through MCP.
	//
	// example:
	//
	// mcp
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	KbUuid     *string `json:"KbUuid,omitempty" xml:"KbUuid,omitempty"`
	// The ID of the MCP Server.
	//
	// example:
	//
	// nhdpt9adf6ac**********ca
	McpServerId *string `json:"McpServerId,omitempty" xml:"McpServerId,omitempty"`
}

func (s ModifyCustomAgentResponseBodyDataKnowledgeConfigList) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentResponseBodyDataKnowledgeConfigList) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeConfigList) GetAccessType() *string {
	return s.AccessType
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeConfigList) GetKbUuid() *string {
	return s.KbUuid
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeConfigList) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeConfigList) SetAccessType(v string) *ModifyCustomAgentResponseBodyDataKnowledgeConfigList {
	s.AccessType = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeConfigList) SetKbUuid(v string) *ModifyCustomAgentResponseBodyDataKnowledgeConfigList {
	s.KbUuid = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeConfigList) SetMcpServerId(v string) *ModifyCustomAgentResponseBodyDataKnowledgeConfigList {
	s.McpServerId = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeConfigList) Validate() error {
	return dara.Validate(s)
}

type ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList struct {
	DbId          *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KnowledgeUuid *string `json:"KnowledgeUuid,omitempty" xml:"KnowledgeUuid,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) GetDbId() *string {
	return s.DbId
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) GetKnowledgeUuid() *string {
	return s.KnowledgeUuid
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) GetType() *string {
	return s.Type
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) SetDbId(v string) *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList {
	s.DbId = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) SetInstanceId(v string) *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList {
	s.InstanceId = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) SetKnowledgeUuid(v string) *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList {
	s.KnowledgeUuid = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) SetType(v string) *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList {
	s.Type = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataKnowledgeSemanticConfigList) Validate() error {
	return dara.Validate(s)
}

type ModifyCustomAgentResponseBodyDataScheduleTaskConfig struct {
	// The cron expression for timed scheduling.
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

func (s ModifyCustomAgentResponseBodyDataScheduleTaskConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomAgentResponseBodyDataScheduleTaskConfig) GoString() string {
	return s.String()
}

func (s *ModifyCustomAgentResponseBodyDataScheduleTaskConfig) GetCronExpression() *string {
	return s.CronExpression
}

func (s *ModifyCustomAgentResponseBodyDataScheduleTaskConfig) GetQuery() *string {
	return s.Query
}

func (s *ModifyCustomAgentResponseBodyDataScheduleTaskConfig) GetRelatedSessionId() *string {
	return s.RelatedSessionId
}

func (s *ModifyCustomAgentResponseBodyDataScheduleTaskConfig) SetCronExpression(v string) *ModifyCustomAgentResponseBodyDataScheduleTaskConfig {
	s.CronExpression = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataScheduleTaskConfig) SetQuery(v string) *ModifyCustomAgentResponseBodyDataScheduleTaskConfig {
	s.Query = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataScheduleTaskConfig) SetRelatedSessionId(v string) *ModifyCustomAgentResponseBodyDataScheduleTaskConfig {
	s.RelatedSessionId = &v
	return s
}

func (s *ModifyCustomAgentResponseBodyDataScheduleTaskConfig) Validate() error {
	return dara.Validate(s)
}

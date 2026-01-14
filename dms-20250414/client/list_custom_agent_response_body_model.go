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
	Data *ListCustomAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	Content []*ListCustomAgentResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 5
	TotalElements *int64 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
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
	// example:
	//
	// 16738266********
	AliyunParentId *string `json:"AliyunParentId,omitempty" xml:"AliyunParentId,omitempty"`
	// example:
	//
	// 20372822********
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// example:
	//
	// HaoY*****
	CreatorUserName *string `json:"CreatorUserName,omitempty" xml:"CreatorUserName,omitempty"`
	// example:
	//
	// ca-4y3ca4khkcu**********ysf
	CustomAgentId *string `json:"CustomAgentId,omitempty" xml:"CustomAgentId,omitempty"`
	DataJson      *string `json:"DataJson,omitempty" xml:"DataJson,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// cn-hangzhou
	DmsUnit         *string                                                `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	ExecutionConfig *ListCustomAgentResponseBodyDataContentExecutionConfig `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtModified         *string                                                      `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Instruction         *string                                                      `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	Knowledge           *string                                                      `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	KnowledgeConfigList []*ListCustomAgentResponseBodyDataContentKnowledgeConfigList `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
	// example:
	//
	// 20372822********
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// HaoY*****
	ModifierUserName *string `json:"ModifierUserName,omitempty" xml:"ModifierUserName,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	OfflineTime *string `json:"OfflineTime,omitempty" xml:"OfflineTime,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// example:
	//
	// RELEASED
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TextReportConfig *string `json:"TextReportConfig,omitempty" xml:"TextReportConfig,omitempty"`
	WebReportConfig  *string `json:"WebReportConfig,omitempty" xml:"WebReportConfig,omitempty"`
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

func (s *ListCustomAgentResponseBodyDataContent) GetCreatorUserName() *string {
	return s.CreatorUserName
}

func (s *ListCustomAgentResponseBodyDataContent) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *ListCustomAgentResponseBodyDataContent) GetDataJson() *string {
	return s.DataJson
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

func (s *ListCustomAgentResponseBodyDataContent) GetOfflineTime() *string {
	return s.OfflineTime
}

func (s *ListCustomAgentResponseBodyDataContent) GetRegion() *string {
	return s.Region
}

func (s *ListCustomAgentResponseBodyDataContent) GetReleaseTime() *string {
	return s.ReleaseTime
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

func (s *ListCustomAgentResponseBodyDataContent) SetCreatorUserName(v string) *ListCustomAgentResponseBodyDataContent {
	s.CreatorUserName = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetCustomAgentId(v string) *ListCustomAgentResponseBodyDataContent {
	s.CustomAgentId = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetDataJson(v string) *ListCustomAgentResponseBodyDataContent {
	s.DataJson = &v
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

func (s *ListCustomAgentResponseBodyDataContent) SetOfflineTime(v string) *ListCustomAgentResponseBodyDataContent {
	s.OfflineTime = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetRegion(v string) *ListCustomAgentResponseBodyDataContent {
	s.Region = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContent) SetReleaseTime(v string) *ListCustomAgentResponseBodyDataContent {
	s.ReleaseTime = &v
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
	return nil
}

type ListCustomAgentResponseBodyDataContentExecutionConfig struct {
	SkipAskHuman *bool `json:"SkipAskHuman,omitempty" xml:"SkipAskHuman,omitempty"`
	// example:
	//
	// true
	SkipPlan             *bool `json:"SkipPlan,omitempty" xml:"SkipPlan,omitempty"`
	SkipSqlConfirm       *bool `json:"SkipSqlConfirm,omitempty" xml:"SkipSqlConfirm,omitempty"`
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

func (s *ListCustomAgentResponseBodyDataContentKnowledgeConfigList) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ListCustomAgentResponseBodyDataContentKnowledgeConfigList) SetAccessType(v string) *ListCustomAgentResponseBodyDataContentKnowledgeConfigList {
	s.AccessType = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentKnowledgeConfigList) SetMcpServerId(v string) *ListCustomAgentResponseBodyDataContentKnowledgeConfigList {
	s.McpServerId = &v
	return s
}

func (s *ListCustomAgentResponseBodyDataContentKnowledgeConfigList) Validate() error {
	return dara.Validate(s)
}

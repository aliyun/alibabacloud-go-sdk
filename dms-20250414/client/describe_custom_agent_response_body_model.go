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
	Data *DescribeCustomAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// example:
	//
	// 16738266********
	AliyunParentUid *string `json:"AliyunParentUid,omitempty" xml:"AliyunParentUid,omitempty"`
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
	DmsUnit         *string                                             `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	ExecutionConfig *DescribeCustomAgentResponseBodyDataExecutionConfig `json:"ExecutionConfig,omitempty" xml:"ExecutionConfig,omitempty" type:"Struct"`
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2025-12-11T14:04:32.000+00:00
	GmtModified         *string                                                   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Instruction         *string                                                   `json:"Instruction,omitempty" xml:"Instruction,omitempty"`
	Knowledge           *string                                                   `json:"Knowledge,omitempty" xml:"Knowledge,omitempty"`
	KnowledgeConfigList []*DescribeCustomAgentResponseBodyDataKnowledgeConfigList `json:"KnowledgeConfigList,omitempty" xml:"KnowledgeConfigList,omitempty" type:"Repeated"`
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

func (s *DescribeCustomAgentResponseBodyData) GetCreatorUserName() *string {
	return s.CreatorUserName
}

func (s *DescribeCustomAgentResponseBodyData) GetCustomAgentId() *string {
	return s.CustomAgentId
}

func (s *DescribeCustomAgentResponseBodyData) GetDataJson() *string {
	return s.DataJson
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

func (s *DescribeCustomAgentResponseBodyData) GetOfflineTime() *string {
	return s.OfflineTime
}

func (s *DescribeCustomAgentResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *DescribeCustomAgentResponseBodyData) GetReleaseTime() *string {
	return s.ReleaseTime
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

func (s *DescribeCustomAgentResponseBodyData) SetCreatorUserName(v string) *DescribeCustomAgentResponseBodyData {
	s.CreatorUserName = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetCustomAgentId(v string) *DescribeCustomAgentResponseBodyData {
	s.CustomAgentId = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetDataJson(v string) *DescribeCustomAgentResponseBodyData {
	s.DataJson = &v
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

func (s *DescribeCustomAgentResponseBodyData) SetOfflineTime(v string) *DescribeCustomAgentResponseBodyData {
	s.OfflineTime = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetRegion(v string) *DescribeCustomAgentResponseBodyData {
	s.Region = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyData) SetReleaseTime(v string) *DescribeCustomAgentResponseBodyData {
	s.ReleaseTime = &v
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

type DescribeCustomAgentResponseBodyDataExecutionConfig struct {
	SkipAskHuman *bool `json:"SkipAskHuman,omitempty" xml:"SkipAskHuman,omitempty"`
	// example:
	//
	// true
	SkipPlan             *bool `json:"SkipPlan,omitempty" xml:"SkipPlan,omitempty"`
	SkipSqlConfirm       *bool `json:"SkipSqlConfirm,omitempty" xml:"SkipSqlConfirm,omitempty"`
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

func (s *DescribeCustomAgentResponseBodyDataKnowledgeConfigList) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeConfigList) SetAccessType(v string) *DescribeCustomAgentResponseBodyDataKnowledgeConfigList {
	s.AccessType = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeConfigList) SetMcpServerId(v string) *DescribeCustomAgentResponseBodyDataKnowledgeConfigList {
	s.McpServerId = &v
	return s
}

func (s *DescribeCustomAgentResponseBodyDataKnowledgeConfigList) Validate() error {
	return dara.Validate(s)
}

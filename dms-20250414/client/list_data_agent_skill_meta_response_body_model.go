// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentSkillMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataAgentSkillMetaResponseBodyData) *ListDataAgentSkillMetaResponseBody
	GetData() *ListDataAgentSkillMetaResponseBodyData
	SetErrorCode(v string) *ListDataAgentSkillMetaResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataAgentSkillMetaResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDataAgentSkillMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataAgentSkillMetaResponseBody
	GetSuccess() *bool
}

type ListDataAgentSkillMetaResponseBody struct {
	// The response struct.
	Data *ListDataAgentSkillMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned when the request is abnormal.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
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

func (s ListDataAgentSkillMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentSkillMetaResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataAgentSkillMetaResponseBody) GetData() *ListDataAgentSkillMetaResponseBodyData {
	return s.Data
}

func (s *ListDataAgentSkillMetaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataAgentSkillMetaResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataAgentSkillMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataAgentSkillMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataAgentSkillMetaResponseBody) SetData(v *ListDataAgentSkillMetaResponseBodyData) *ListDataAgentSkillMetaResponseBody {
	s.Data = v
	return s
}

func (s *ListDataAgentSkillMetaResponseBody) SetErrorCode(v string) *ListDataAgentSkillMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBody) SetErrorMessage(v string) *ListDataAgentSkillMetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBody) SetRequestId(v string) *ListDataAgentSkillMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBody) SetSuccess(v bool) *ListDataAgentSkillMetaResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataAgentSkillMetaResponseBodyData struct {
	// The list of data content.
	Content []*ListDataAgentSkillMetaResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
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

func (s ListDataAgentSkillMetaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentSkillMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataAgentSkillMetaResponseBodyData) GetContent() []*ListDataAgentSkillMetaResponseBodyDataContent {
	return s.Content
}

func (s *ListDataAgentSkillMetaResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDataAgentSkillMetaResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDataAgentSkillMetaResponseBodyData) GetTotalElements() *int64 {
	return s.TotalElements
}

func (s *ListDataAgentSkillMetaResponseBodyData) GetTotalPages() *int64 {
	return s.TotalPages
}

func (s *ListDataAgentSkillMetaResponseBodyData) SetContent(v []*ListDataAgentSkillMetaResponseBodyDataContent) *ListDataAgentSkillMetaResponseBodyData {
	s.Content = v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyData) SetPageNumber(v int64) *ListDataAgentSkillMetaResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyData) SetPageSize(v int64) *ListDataAgentSkillMetaResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyData) SetTotalElements(v int64) *ListDataAgentSkillMetaResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyData) SetTotalPages(v int64) *ListDataAgentSkillMetaResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyData) Validate() error {
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

type ListDataAgentSkillMetaResponseBodyDataContent struct {
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
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// The name of the creator.
	//
	// example:
	//
	// HaoY*****
	CreatorUserName *string `json:"CreatorUserName,omitempty" xml:"CreatorUserName,omitempty"`
	// The skill description.
	//
	// example:
	//
	// This is a demo skill description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the skill is available. Valid values: true and false.
	//
	// example:
	//
	// true
	Enabled *int32 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
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
	// The skill parsing error message. This parameter is returned when the skill status is INVALID.
	//
	// example:
	//
	// SKILL.md file not exist.
	ParseError *string `json:"ParseError,omitempty" xml:"ParseError,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The source of the skill. Valid values:
	//
	// - User: a skill uploaded by the user.
	//
	// - Agent: a skill derived from Agent analysis.
	//
	// example:
	//
	// User
	SkillFrom *string `json:"SkillFrom,omitempty" xml:"SkillFrom,omitempty"`
	// The skill ID.
	//
	// example:
	//
	// ski-04pomiln*************j0
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The skill name.
	//
	// example:
	//
	// data-query-skill
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
	// The skill status. Valid values:
	//
	// - INIT: not ready.
	//
	// - ACTIVE: active.
	//
	// - INVALID: invalid.
	//
	// example:
	//
	// ACTIVE
	SkillStatus *string `json:"SkillStatus,omitempty" xml:"SkillStatus,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 56kv1pvl9uvt9**********bb
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataAgentSkillMetaResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentSkillMetaResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetAliyunParentUid() *string {
	return s.AliyunParentUid
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetCreatorUserName() *string {
	return s.CreatorUserName
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetDescription() *string {
	return s.Description
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetEnabled() *int32 {
	return s.Enabled
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetParseError() *string {
	return s.ParseError
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetRegion() *string {
	return s.Region
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetSkillFrom() *string {
	return s.SkillFrom
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetSkillId() *string {
	return s.SkillId
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetSkillName() *string {
	return s.SkillName
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetSkillStatus() *string {
	return s.SkillStatus
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetAliyunParentUid(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.AliyunParentUid = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetAliyunUid(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.AliyunUid = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetCreatorUserName(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.CreatorUserName = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetDescription(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.Description = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetEnabled(v int32) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.Enabled = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetGmtCreated(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.GmtCreated = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetGmtModified(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.GmtModified = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetParseError(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.ParseError = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetRegion(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.Region = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetSkillFrom(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.SkillFrom = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetSkillId(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.SkillId = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetSkillName(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.SkillName = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetSkillStatus(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.SkillStatus = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) SetWorkspaceId(v string) *ListDataAgentSkillMetaResponseBodyDataContent {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataAgentSkillMetaResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}

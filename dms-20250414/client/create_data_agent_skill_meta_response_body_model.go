// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentSkillMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDataAgentSkillMetaResponseBodyData) *CreateDataAgentSkillMetaResponseBody
	GetData() *CreateDataAgentSkillMetaResponseBodyData
	SetErrorCode(v string) *CreateDataAgentSkillMetaResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataAgentSkillMetaResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataAgentSkillMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataAgentSkillMetaResponseBody
	GetSuccess() *bool
}

type CreateDataAgentSkillMetaResponseBody struct {
	// The response struct.
	Data *CreateDataAgentSkillMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when a system-level request failure occurs.
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

func (s CreateDataAgentSkillMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentSkillMetaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataAgentSkillMetaResponseBody) GetData() *CreateDataAgentSkillMetaResponseBodyData {
	return s.Data
}

func (s *CreateDataAgentSkillMetaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataAgentSkillMetaResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataAgentSkillMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataAgentSkillMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataAgentSkillMetaResponseBody) SetData(v *CreateDataAgentSkillMetaResponseBodyData) *CreateDataAgentSkillMetaResponseBody {
	s.Data = v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBody) SetErrorCode(v string) *CreateDataAgentSkillMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBody) SetErrorMessage(v string) *CreateDataAgentSkillMetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBody) SetRequestId(v string) *CreateDataAgentSkillMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBody) SetSuccess(v bool) *CreateDataAgentSkillMetaResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataAgentSkillMetaResponseBodyData struct {
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
	// The creator name.
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
	// The skill parsing error message.
	//
	// - When the skill status is INVALID, the parsing error message is returned.
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
	// The skill source.
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
	// - INIT: Not ready.
	//
	// - ACTIVE: Active.
	//
	// - INVALID: Invalid.
	//
	// example:
	//
	// ACTIVE
	SkillStatus *string `json:"SkillStatus,omitempty" xml:"SkillStatus,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 11if52e44**********edbv6
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataAgentSkillMetaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentSkillMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetAliyunParentUid() *string {
	return s.AliyunParentUid
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetCreatorUserName() *string {
	return s.CreatorUserName
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetEnabled() *int32 {
	return s.Enabled
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetParseError() *string {
	return s.ParseError
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetSkillFrom() *string {
	return s.SkillFrom
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetSkillId() *string {
	return s.SkillId
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetSkillName() *string {
	return s.SkillName
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetSkillStatus() *string {
	return s.SkillStatus
}

func (s *CreateDataAgentSkillMetaResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetAliyunParentUid(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.AliyunParentUid = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetAliyunUid(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.AliyunUid = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetCreatorUserName(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.CreatorUserName = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetDescription(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.Description = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetEnabled(v int32) *CreateDataAgentSkillMetaResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetGmtCreated(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetGmtModified(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetParseError(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.ParseError = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetRegion(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.Region = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetSkillFrom(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.SkillFrom = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetSkillId(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.SkillId = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetSkillName(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.SkillName = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetSkillStatus(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.SkillStatus = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) SetWorkspaceId(v string) *CreateDataAgentSkillMetaResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataAgentSkillMetaResponseBodyData) Validate() error {
	return dara.Validate(s)
}

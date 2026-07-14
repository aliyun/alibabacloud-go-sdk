// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListChatGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListChatGroupResponseBody
	GetCode() *string
	SetData(v *ListChatGroupResponseBodyData) *ListChatGroupResponseBody
	GetData() *ListChatGroupResponseBodyData
	SetMessage(v string) *ListChatGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListChatGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListChatGroupResponseBody
	GetSuccess() *bool
}

type ListChatGroupResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code. Valid values:
	//
	// - OK: The request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListChatGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The prompt message. This parameter has a value when an exception is returned.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListChatGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListChatGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatGroupResponseBody) GetData() *ListChatGroupResponseBodyData {
	return s.Data
}

func (s *ListChatGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListChatGroupResponseBody) SetAccessDeniedDetail(v string) *ListChatGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListChatGroupResponseBody) SetCode(v string) *ListChatGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatGroupResponseBody) SetData(v *ListChatGroupResponseBodyData) *ListChatGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListChatGroupResponseBody) SetMessage(v string) *ListChatGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatGroupResponseBody) SetRequestId(v string) *ListChatGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatGroupResponseBody) SetSuccess(v bool) *ListChatGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ListChatGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChatGroupResponseBodyData struct {
	// The group list.
	List []*ListChatGroupResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 51
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListChatGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChatGroupResponseBodyData) GetList() []*ListChatGroupResponseBodyDataList {
	return s.List
}

func (s *ListChatGroupResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListChatGroupResponseBodyData) SetList(v []*ListChatGroupResponseBodyDataList) *ListChatGroupResponseBodyData {
	s.List = v
	return s
}

func (s *ListChatGroupResponseBodyData) SetTotal(v int64) *ListChatGroupResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListChatGroupResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChatGroupResponseBodyDataList struct {
	// The business phone number.
	//
	// example:
	//
	// 8613800**
	BusinessNumber *string `json:"BusinessNumber,omitempty" xml:"BusinessNumber,omitempty"`
	// The role of the bot in the group.
	//
	// example:
	//
	// example
	BusinessRole *string `json:"BusinessRole,omitempty" xml:"BusinessRole,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The update time.
	//
	// example:
	//
	// 94
	GmtModifier *int64 `json:"GmtModifier,omitempty" xml:"GmtModifier,omitempty"`
	// The group ID.
	//
	// example:
	//
	// EA30d***
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The group link.
	//
	// example:
	//
	// example
	GroupLink *string `json:"GroupLink,omitempty" xml:"GroupLink,omitempty"`
	// The group status.
	//
	// example:
	//
	// ACTIVE
	GroupStatus *string `json:"GroupStatus,omitempty" xml:"GroupStatus,omitempty"`
	// The group type.
	//
	// example:
	//
	// example
	GroupType *string `json:"GroupType,omitempty" xml:"GroupType,omitempty"`
	// The invite link.
	//
	// example:
	//
	// https://chat.whatsapp.com/***
	InviteLink *string `json:"InviteLink,omitempty" xml:"InviteLink,omitempty"`
	// The group profile picture.
	ProfilePictureFile *string `json:"ProfilePictureFile,omitempty" xml:"ProfilePictureFile,omitempty"`
	// The group subject.
	//
	// example:
	//
	// This is a test subject
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	// The number of group members.
	//
	// example:
	//
	// 35
	TotalParticipantCount *int64 `json:"TotalParticipantCount,omitempty" xml:"TotalParticipantCount,omitempty"`
}

func (s ListChatGroupResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListChatGroupResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListChatGroupResponseBodyDataList) GetBusinessNumber() *string {
	return s.BusinessNumber
}

func (s *ListChatGroupResponseBodyDataList) GetBusinessRole() *string {
	return s.BusinessRole
}

func (s *ListChatGroupResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListChatGroupResponseBodyDataList) GetGmtModifier() *int64 {
	return s.GmtModifier
}

func (s *ListChatGroupResponseBodyDataList) GetGroupId() *string {
	return s.GroupId
}

func (s *ListChatGroupResponseBodyDataList) GetGroupLink() *string {
	return s.GroupLink
}

func (s *ListChatGroupResponseBodyDataList) GetGroupStatus() *string {
	return s.GroupStatus
}

func (s *ListChatGroupResponseBodyDataList) GetGroupType() *string {
	return s.GroupType
}

func (s *ListChatGroupResponseBodyDataList) GetInviteLink() *string {
	return s.InviteLink
}

func (s *ListChatGroupResponseBodyDataList) GetProfilePictureFile() *string {
	return s.ProfilePictureFile
}

func (s *ListChatGroupResponseBodyDataList) GetSubject() *string {
	return s.Subject
}

func (s *ListChatGroupResponseBodyDataList) GetTotalParticipantCount() *int64 {
	return s.TotalParticipantCount
}

func (s *ListChatGroupResponseBodyDataList) SetBusinessNumber(v string) *ListChatGroupResponseBodyDataList {
	s.BusinessNumber = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetBusinessRole(v string) *ListChatGroupResponseBodyDataList {
	s.BusinessRole = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetDescription(v string) *ListChatGroupResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetGmtModifier(v int64) *ListChatGroupResponseBodyDataList {
	s.GmtModifier = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetGroupId(v string) *ListChatGroupResponseBodyDataList {
	s.GroupId = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetGroupLink(v string) *ListChatGroupResponseBodyDataList {
	s.GroupLink = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetGroupStatus(v string) *ListChatGroupResponseBodyDataList {
	s.GroupStatus = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetGroupType(v string) *ListChatGroupResponseBodyDataList {
	s.GroupType = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetInviteLink(v string) *ListChatGroupResponseBodyDataList {
	s.InviteLink = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetProfilePictureFile(v string) *ListChatGroupResponseBodyDataList {
	s.ProfilePictureFile = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetSubject(v string) *ListChatGroupResponseBodyDataList {
	s.Subject = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) SetTotalParticipantCount(v int64) *ListChatGroupResponseBodyDataList {
	s.TotalParticipantCount = &v
	return s
}

func (s *ListChatGroupResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

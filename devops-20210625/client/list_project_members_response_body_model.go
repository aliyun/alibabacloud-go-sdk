// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListProjectMembersResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListProjectMembersResponseBody
	GetErrorMsg() *string
	SetMembers(v []*ListProjectMembersResponseBodyMembers) *ListProjectMembersResponseBody
	GetMembers() []*ListProjectMembersResponseBodyMembers
	SetRequestId(v string) *ListProjectMembersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProjectMembersResponseBody
	GetSuccess() *bool
}

type ListProjectMembersResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string                                  `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Members  []*ListProjectMembersResponseBodyMembers `json:"members,omitempty" xml:"members,omitempty" type:"Repeated"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListProjectMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListProjectMembersResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListProjectMembersResponseBody) GetMembers() []*ListProjectMembersResponseBodyMembers {
	return s.Members
}

func (s *ListProjectMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectMembersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectMembersResponseBody) SetErrorCode(v string) *ListProjectMembersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectMembersResponseBody) SetErrorMsg(v string) *ListProjectMembersResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListProjectMembersResponseBody) SetMembers(v []*ListProjectMembersResponseBodyMembers) *ListProjectMembersResponseBody {
	s.Members = v
	return s
}

func (s *ListProjectMembersResponseBody) SetRequestId(v string) *ListProjectMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectMembersResponseBody) SetSuccess(v bool) *ListProjectMembersResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectMembersResponseBody) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectMembersResponseBodyMembers struct {
	// example:
	//
	// null
	Account *string `json:"account,omitempty" xml:"account,omitempty"`
	// example:
	//
	// https://xxxxxx.png
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// null
	DingTalkId *string `json:"dingTalkId,omitempty" xml:"dingTalkId,omitempty"`
	// example:
	//
	// 名称
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// mingcheng
	DisplayNickName *string `json:"displayNickName,omitempty" xml:"displayNickName,omitempty"`
	// example:
	//
	// 名称
	DisplayRealName *string                                        `json:"displayRealName,omitempty" xml:"displayRealName,omitempty"`
	Division        *ListProjectMembersResponseBodyMembersDivision `json:"division,omitempty" xml:"division,omitempty" type:"Struct"`
	// example:
	//
	// accountsxxxx@mail.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// null
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// 19xxxx31947xxxx
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 135xxxxxxxxx
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// example:
	//
	// name
	NameEn *string `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	// example:
	//
	// 昵称
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// nicheng
	NickNamePinyin       *string                                                    `json:"nickNamePinyin,omitempty" xml:"nickNamePinyin,omitempty"`
	OrganizationUserInfo *ListProjectMembersResponseBodyMembersOrganizationUserInfo `json:"organizationUserInfo,omitempty" xml:"organizationUserInfo,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	RealName *string `json:"realName,omitempty" xml:"realName,omitempty"`
	// example:
	//
	// xxx
	RealNamePinyin *string `json:"realNamePinyin,omitempty" xml:"realNamePinyin,omitempty"`
	// 角色名称
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
	// example:
	//
	// User
	Stamp *string `json:"stamp,omitempty" xml:"stamp,omitempty"`
	// example:
	//
	// null
	TbRoleId *string `json:"tbRoleId,omitempty" xml:"tbRoleId,omitempty"`
}

func (s ListProjectMembersResponseBodyMembers) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyMembers) GetAccount() *string {
	return s.Account
}

func (s *ListProjectMembersResponseBodyMembers) GetAvatar() *string {
	return s.Avatar
}

func (s *ListProjectMembersResponseBodyMembers) GetDingTalkId() *string {
	return s.DingTalkId
}

func (s *ListProjectMembersResponseBodyMembers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListProjectMembersResponseBodyMembers) GetDisplayNickName() *string {
	return s.DisplayNickName
}

func (s *ListProjectMembersResponseBodyMembers) GetDisplayRealName() *string {
	return s.DisplayRealName
}

func (s *ListProjectMembersResponseBodyMembers) GetDivision() *ListProjectMembersResponseBodyMembersDivision {
	return s.Division
}

func (s *ListProjectMembersResponseBodyMembers) GetEmail() *string {
	return s.Email
}

func (s *ListProjectMembersResponseBodyMembers) GetGender() *string {
	return s.Gender
}

func (s *ListProjectMembersResponseBodyMembers) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListProjectMembersResponseBodyMembers) GetMobile() *string {
	return s.Mobile
}

func (s *ListProjectMembersResponseBodyMembers) GetNameEn() *string {
	return s.NameEn
}

func (s *ListProjectMembersResponseBodyMembers) GetNickName() *string {
	return s.NickName
}

func (s *ListProjectMembersResponseBodyMembers) GetNickNamePinyin() *string {
	return s.NickNamePinyin
}

func (s *ListProjectMembersResponseBodyMembers) GetOrganizationUserInfo() *ListProjectMembersResponseBodyMembersOrganizationUserInfo {
	return s.OrganizationUserInfo
}

func (s *ListProjectMembersResponseBodyMembers) GetRealName() *string {
	return s.RealName
}

func (s *ListProjectMembersResponseBodyMembers) GetRealNamePinyin() *string {
	return s.RealNamePinyin
}

func (s *ListProjectMembersResponseBodyMembers) GetRoleName() *string {
	return s.RoleName
}

func (s *ListProjectMembersResponseBodyMembers) GetStamp() *string {
	return s.Stamp
}

func (s *ListProjectMembersResponseBodyMembers) GetTbRoleId() *string {
	return s.TbRoleId
}

func (s *ListProjectMembersResponseBodyMembers) SetAccount(v string) *ListProjectMembersResponseBodyMembers {
	s.Account = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetAvatar(v string) *ListProjectMembersResponseBodyMembers {
	s.Avatar = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetDingTalkId(v string) *ListProjectMembersResponseBodyMembers {
	s.DingTalkId = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetDisplayName(v string) *ListProjectMembersResponseBodyMembers {
	s.DisplayName = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetDisplayNickName(v string) *ListProjectMembersResponseBodyMembers {
	s.DisplayNickName = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetDisplayRealName(v string) *ListProjectMembersResponseBodyMembers {
	s.DisplayRealName = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetDivision(v *ListProjectMembersResponseBodyMembersDivision) *ListProjectMembersResponseBodyMembers {
	s.Division = v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetEmail(v string) *ListProjectMembersResponseBodyMembers {
	s.Email = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetGender(v string) *ListProjectMembersResponseBodyMembers {
	s.Gender = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetIdentifier(v string) *ListProjectMembersResponseBodyMembers {
	s.Identifier = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetMobile(v string) *ListProjectMembersResponseBodyMembers {
	s.Mobile = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetNameEn(v string) *ListProjectMembersResponseBodyMembers {
	s.NameEn = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetNickName(v string) *ListProjectMembersResponseBodyMembers {
	s.NickName = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetNickNamePinyin(v string) *ListProjectMembersResponseBodyMembers {
	s.NickNamePinyin = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetOrganizationUserInfo(v *ListProjectMembersResponseBodyMembersOrganizationUserInfo) *ListProjectMembersResponseBodyMembers {
	s.OrganizationUserInfo = v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetRealName(v string) *ListProjectMembersResponseBodyMembers {
	s.RealName = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetRealNamePinyin(v string) *ListProjectMembersResponseBodyMembers {
	s.RealNamePinyin = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetRoleName(v string) *ListProjectMembersResponseBodyMembers {
	s.RoleName = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetStamp(v string) *ListProjectMembersResponseBodyMembers {
	s.Stamp = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) SetTbRoleId(v string) *ListProjectMembersResponseBodyMembers {
	s.TbRoleId = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembers) Validate() error {
	if s.Division != nil {
		if err := s.Division.Validate(); err != nil {
			return err
		}
	}
	if s.OrganizationUserInfo != nil {
		if err := s.OrganizationUserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectMembersResponseBodyMembersDivision struct {
	// example:
	//
	// 1345xxxxxxx947xxxx
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
}

func (s ListProjectMembersResponseBodyMembersDivision) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBodyMembersDivision) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyMembersDivision) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListProjectMembersResponseBodyMembersDivision) SetIdentifier(v string) *ListProjectMembersResponseBodyMembersDivision {
	s.Identifier = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembersDivision) Validate() error {
	return dara.Validate(s)
}

type ListProjectMembersResponseBodyMembersOrganizationUserInfo struct {
	// example:
	//
	// 5e7xxxxb3cd3711dd6xxx2c
	OrganizationIdentifier *string `json:"organizationIdentifier,omitempty" xml:"organizationIdentifier,omitempty"`
}

func (s ListProjectMembersResponseBodyMembersOrganizationUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersResponseBodyMembersOrganizationUserInfo) GoString() string {
	return s.String()
}

func (s *ListProjectMembersResponseBodyMembersOrganizationUserInfo) GetOrganizationIdentifier() *string {
	return s.OrganizationIdentifier
}

func (s *ListProjectMembersResponseBodyMembersOrganizationUserInfo) SetOrganizationIdentifier(v string) *ListProjectMembersResponseBodyMembersOrganizationUserInfo {
	s.OrganizationIdentifier = &v
	return s
}

func (s *ListProjectMembersResponseBodyMembersOrganizationUserInfo) Validate() error {
	return dara.Validate(s)
}

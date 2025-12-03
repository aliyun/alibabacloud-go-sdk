// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetWorkitemTime(v *CreateWorkitemRecordResponseBodyWorkitemTime) *CreateWorkitemRecordResponseBody
	GetWorkitemTime() *CreateWorkitemRecordResponseBodyWorkitemTime
	SetErrorCode(v string) *CreateWorkitemRecordResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateWorkitemRecordResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateWorkitemRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkitemRecordResponseBody
	GetSuccess() *bool
}

type CreateWorkitemRecordResponseBody struct {
	WorkitemTime *CreateWorkitemRecordResponseBodyWorkitemTime `json:"WorkitemTime,omitempty" xml:"WorkitemTime,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateWorkitemRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRecordResponseBody) GetWorkitemTime() *CreateWorkitemRecordResponseBodyWorkitemTime {
	return s.WorkitemTime
}

func (s *CreateWorkitemRecordResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateWorkitemRecordResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateWorkitemRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkitemRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkitemRecordResponseBody) SetWorkitemTime(v *CreateWorkitemRecordResponseBodyWorkitemTime) *CreateWorkitemRecordResponseBody {
	s.WorkitemTime = v
	return s
}

func (s *CreateWorkitemRecordResponseBody) SetErrorCode(v string) *CreateWorkitemRecordResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWorkitemRecordResponseBody) SetErrorMsg(v string) *CreateWorkitemRecordResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateWorkitemRecordResponseBody) SetRequestId(v string) *CreateWorkitemRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkitemRecordResponseBody) SetSuccess(v bool) *CreateWorkitemRecordResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkitemRecordResponseBody) Validate() error {
	if s.WorkitemTime != nil {
		if err := s.WorkitemTime.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkitemRecordResponseBodyWorkitemTime struct {
	// example:
	//
	// 12
	ActualTime  *int64  `json:"actualTime,omitempty" xml:"actualTime,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1667205617089
	GmtEnd *int64 `json:"gmtEnd,omitempty" xml:"gmtEnd,omitempty"`
	// example:
	//
	// 1667205617061
	GmtStart *int64 `json:"gmtStart,omitempty" xml:"gmtStart,omitempty"`
	// example:
	//
	// deafe5f33xxxxx6a259d8dafd
	Identifier *string                                                 `json:"identifier,omitempty" xml:"identifier,omitempty"`
	RecordUser *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser `json:"recordUser,omitempty" xml:"recordUser,omitempty" type:"Struct"`
	// example:
	//
	// deafe5f33xxxxx6a259d8dafd
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 9144ef6b72d8exxxxx9e61a4d0
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemRecordResponseBodyWorkitemTime) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemRecordResponseBodyWorkitemTime) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) GetActualTime() *int64 {
	return s.ActualTime
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) GetGmtEnd() *int64 {
	return s.GmtEnd
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) GetGmtStart() *int64 {
	return s.GmtStart
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) GetRecordUser() *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	return s.RecordUser
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) GetType() *string {
	return s.Type
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetActualTime(v int64) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.ActualTime = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetDescription(v string) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.Description = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetGmtEnd(v int64) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.GmtEnd = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetGmtStart(v int64) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.GmtStart = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetIdentifier(v string) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetRecordUser(v *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.RecordUser = v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetType(v string) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.Type = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) SetWorkitemIdentifier(v string) *CreateWorkitemRecordResponseBodyWorkitemTime {
	s.WorkitemIdentifier = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTime) Validate() error {
	if s.RecordUser != nil {
		if err := s.RecordUser.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser struct {
	// example:
	//
	// 1316458xxxxx41068
	Account *string `json:"account,omitempty" xml:"account,omitempty"`
	// example:
	//
	// https://tcs-devops.aliyuncs.com/thumbnail/112m7150e256dbba0d6456bafbb65c81f90d/w/200/h/200
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// xx12322
	DingTalkId      *string `json:"dingTalkId,omitempty" xml:"dingTalkId,omitempty"`
	DisplayName     *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	DisplayNickName *string `json:"displayNickName,omitempty" xml:"displayNickName,omitempty"`
	DisplayRealName *string `json:"displayRealName,omitempty" xml:"displayRealName,omitempty"`
	// example:
	//
	// xxx@xxx.com
	Email  *string `json:"email,omitempty" xml:"email,omitempty"`
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// deafe5f33xxxxx6a259d8dafd
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// true/false
	IsDisabled *bool   `json:"isDisabled,omitempty" xml:"isDisabled,omitempty"`
	Mobile     *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// example:
	//
	// wangwu
	NameEn   *string `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	NickName *string `json:"nickName,omitempty" xml:"nickName,omitempty"`
	// example:
	//
	// wangwu
	NickNamePinyin *string `json:"nickNamePinyin,omitempty" xml:"nickNamePinyin,omitempty"`
	RealName       *string `json:"realName,omitempty" xml:"realName,omitempty"`
	// example:
	//
	// wangwu
	RealNamePinyin *string `json:"realNamePinyin,omitempty" xml:"realNamePinyin,omitempty"`
	Stamp          *string `json:"stamp,omitempty" xml:"stamp,omitempty"`
	// example:
	//
	// 1e9903d8b3f1xxxxxf9286ef5
	TbRoleId *string `json:"tbRoleId,omitempty" xml:"tbRoleId,omitempty"`
}

func (s CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GoString() string {
	return s.String()
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetAccount() *string {
	return s.Account
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetAvatar() *string {
	return s.Avatar
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetDingTalkId() *string {
	return s.DingTalkId
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetDisplayNickName() *string {
	return s.DisplayNickName
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetDisplayRealName() *string {
	return s.DisplayRealName
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetEmail() *string {
	return s.Email
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetGender() *string {
	return s.Gender
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetIsDisabled() *bool {
	return s.IsDisabled
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetMobile() *string {
	return s.Mobile
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetNameEn() *string {
	return s.NameEn
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetNickName() *string {
	return s.NickName
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetNickNamePinyin() *string {
	return s.NickNamePinyin
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetRealName() *string {
	return s.RealName
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetRealNamePinyin() *string {
	return s.RealNamePinyin
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetStamp() *string {
	return s.Stamp
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) GetTbRoleId() *string {
	return s.TbRoleId
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetAccount(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Account = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetAvatar(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Avatar = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetDingTalkId(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.DingTalkId = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetDisplayName(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.DisplayName = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetDisplayNickName(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.DisplayNickName = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetDisplayRealName(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.DisplayRealName = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetEmail(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Email = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetGender(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Gender = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetIdentifier(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetIsDisabled(v bool) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.IsDisabled = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetMobile(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Mobile = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetNameEn(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.NameEn = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetNickName(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.NickName = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetNickNamePinyin(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.NickNamePinyin = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetRealName(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.RealName = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetRealNamePinyin(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.RealNamePinyin = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetStamp(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.Stamp = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) SetTbRoleId(v string) *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser {
	s.TbRoleId = &v
	return s
}

func (s *CreateWorkitemRecordResponseBodyWorkitemTimeRecordUser) Validate() error {
	return dara.Validate(s)
}

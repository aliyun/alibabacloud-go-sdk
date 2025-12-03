// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkitemEstimateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetWorkitemTimeEstimate(v *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) *CreateWorkitemEstimateResponseBody
	GetWorkitemTimeEstimate() *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate
	SetErrorCode(v string) *CreateWorkitemEstimateResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateWorkitemEstimateResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *CreateWorkitemEstimateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkitemEstimateResponseBody
	GetSuccess() *bool
}

type CreateWorkitemEstimateResponseBody struct {
	WorkitemTimeEstimate *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate `json:"WorkitemTimeEstimate,omitempty" xml:"WorkitemTimeEstimate,omitempty" type:"Struct"`
	// example:
	//
	// Openapi.RequestError
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
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateWorkitemEstimateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemEstimateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkitemEstimateResponseBody) GetWorkitemTimeEstimate() *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	return s.WorkitemTimeEstimate
}

func (s *CreateWorkitemEstimateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateWorkitemEstimateResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateWorkitemEstimateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkitemEstimateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkitemEstimateResponseBody) SetWorkitemTimeEstimate(v *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) *CreateWorkitemEstimateResponseBody {
	s.WorkitemTimeEstimate = v
	return s
}

func (s *CreateWorkitemEstimateResponseBody) SetErrorCode(v string) *CreateWorkitemEstimateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBody) SetErrorMsg(v string) *CreateWorkitemEstimateResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBody) SetRequestId(v string) *CreateWorkitemEstimateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBody) SetSuccess(v bool) *CreateWorkitemEstimateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBody) Validate() error {
	if s.WorkitemTimeEstimate != nil {
		if err := s.WorkitemTimeEstimate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// deafe5f33xxxxx6a259d8dafd
	Identifier *string                                                           `json:"identifier,omitempty" xml:"identifier,omitempty"`
	RecordUser *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser `json:"recordUser,omitempty" xml:"recordUser,omitempty" type:"Struct"`
	// example:
	//
	// 11
	SpentTime *int64 `json:"spentTime,omitempty" xml:"spentTime,omitempty"`
	// example:
	//
	// 9144ef6b72d8exxxxx9e61a4d0
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1e9903d8b3f1xxxxxf9286ef5
	WorkitemIdentifier *string `json:"workitemIdentifier,omitempty" xml:"workitemIdentifier,omitempty"`
}

func (s CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) GoString() string {
	return s.String()
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetRecordUser() *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	return s.RecordUser
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetSpentTime() *int64 {
	return s.SpentTime
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetType() *string {
	return s.Type
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) GetWorkitemIdentifier() *string {
	return s.WorkitemIdentifier
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetDescription(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.Description = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetIdentifier(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetRecordUser(v *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.RecordUser = v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetSpentTime(v int64) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.SpentTime = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetType(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.Type = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) SetWorkitemIdentifier(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate {
	s.WorkitemIdentifier = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimate) Validate() error {
	if s.RecordUser != nil {
		if err := s.RecordUser.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser struct {
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
	// fdsad12xxx11
	DingTalkId  *string `json:"dingTalkId,omitempty" xml:"dingTalkId,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// wangwu
	DisplayNickName *string `json:"displayNickName,omitempty" xml:"displayNickName,omitempty"`
	// example:
	//
	// wangwu
	DisplayRealName *string `json:"displayRealName,omitempty" xml:"displayRealName,omitempty"`
	// example:
	//
	// xxxxxxcc@xxx.com
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
	// 9144ef6b72d8exxxxx9e61a4d0
	TbRoleId *string `json:"tbRoleId,omitempty" xml:"tbRoleId,omitempty"`
}

func (s CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GoString() string {
	return s.String()
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetAccount() *string {
	return s.Account
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetAvatar() *string {
	return s.Avatar
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetDingTalkId() *string {
	return s.DingTalkId
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetDisplayNickName() *string {
	return s.DisplayNickName
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetDisplayRealName() *string {
	return s.DisplayRealName
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetEmail() *string {
	return s.Email
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetGender() *string {
	return s.Gender
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetIsDisabled() *bool {
	return s.IsDisabled
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetMobile() *string {
	return s.Mobile
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetNameEn() *string {
	return s.NameEn
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetNickName() *string {
	return s.NickName
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetNickNamePinyin() *string {
	return s.NickNamePinyin
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetRealName() *string {
	return s.RealName
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetRealNamePinyin() *string {
	return s.RealNamePinyin
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetStamp() *string {
	return s.Stamp
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) GetTbRoleId() *string {
	return s.TbRoleId
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetAccount(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Account = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetAvatar(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Avatar = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetDingTalkId(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.DingTalkId = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetDisplayName(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.DisplayName = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetDisplayNickName(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.DisplayNickName = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetDisplayRealName(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.DisplayRealName = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetEmail(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Email = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetGender(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Gender = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetIdentifier(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Identifier = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetIsDisabled(v bool) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.IsDisabled = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetMobile(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Mobile = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetNameEn(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.NameEn = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetNickName(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.NickName = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetNickNamePinyin(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.NickNamePinyin = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetRealName(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.RealName = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetRealNamePinyin(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.RealNamePinyin = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetStamp(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.Stamp = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) SetTbRoleId(v string) *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser {
	s.TbRoleId = &v
	return s
}

func (s *CreateWorkitemEstimateResponseBodyWorkitemTimeEstimateRecordUser) Validate() error {
	return dara.Validate(s)
}

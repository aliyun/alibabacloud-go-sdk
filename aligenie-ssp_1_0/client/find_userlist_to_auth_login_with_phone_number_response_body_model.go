// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindUserlistToAuthLoginWithPhoneNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *FindUserlistToAuthLoginWithPhoneNumberResponseBody
	GetCode() *int32
	SetDataObj(v *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj) *FindUserlistToAuthLoginWithPhoneNumberResponseBody
	GetDataObj() *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj
	SetMessage(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBody
	GetRequestId() *string
	SetResult(v *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult) *FindUserlistToAuthLoginWithPhoneNumberResponseBody
	GetResult() *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult
	SetSuccess(v bool) *FindUserlistToAuthLoginWithPhoneNumberResponseBody
	GetSuccess() *bool
}

type FindUserlistToAuthLoginWithPhoneNumberResponseBody struct {
	// example:
	//
	// 200
	Code    *int32                                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	DataObj *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj `json:"DataObj,omitempty" xml:"DataObj,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	Success   *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) GetDataObj() *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj {
	return s.DataObj
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) GetResult() *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult {
	return s.Result
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetCode(v int32) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetDataObj(v *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.DataObj = v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetMessage(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetRequestId(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetResult(v *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.Result = v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) SetSuccess(v bool) *FindUserlistToAuthLoginWithPhoneNumberResponseBody {
	s.Success = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBody) Validate() error {
	return dara.Validate(s)
}

type FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj struct {
	// example:
	//
	// dbe2eb4458302b9246c6da17fbc95f4b
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj) String() string {
	return dara.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj) GetSessionId() *string {
	return s.SessionId
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj) SetSessionId(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj {
	s.SessionId = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyDataObj) Validate() error {
	return dara.Validate(s)
}

type FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult struct {
	UserListToAuthLogin []*FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin `json:"UserListToAuthLogin,omitempty" xml:"UserListToAuthLogin,omitempty" type:"Repeated"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult) GetUserListToAuthLogin() []*FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin {
	return s.UserListToAuthLogin
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult) SetUserListToAuthLogin(v []*FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult {
	s.UserListToAuthLogin = v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin struct {
	// example:
	//
	// https://xxx
	Avatar *string `json:"Avatar,omitempty" xml:"Avatar,omitempty"`
	// example:
	//
	// KsVgypxAipf+xNECMZV2ONMcheqiIoEGFvgx+T8s1oV6/euTK9+ImYvLVPsSqFDh
	EncryptedUserIdentifier *string `json:"EncryptedUserIdentifier,omitempty" xml:"EncryptedUserIdentifier,omitempty"`
	// example:
	//
	// PHONE_NUMBER_BINDING_WITH_ALIGENIE：phoneNumber
	FindingType *string `json:"FindingType,omitempty" xml:"FindingType,omitempty"`
	// example:
	//
	// XXX
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// example:
	//
	// ALIGENIE
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) String() string {
	return dara.Prettify(s)
}

func (s FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) GoString() string {
	return s.String()
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) GetAvatar() *string {
	return s.Avatar
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) GetEncryptedUserIdentifier() *string {
	return s.EncryptedUserIdentifier
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) GetFindingType() *string {
	return s.FindingType
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) GetNickname() *string {
	return s.Nickname
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) GetUserType() *string {
	return s.UserType
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) SetAvatar(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin {
	s.Avatar = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) SetEncryptedUserIdentifier(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin {
	s.EncryptedUserIdentifier = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) SetFindingType(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin {
	s.FindingType = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) SetNickname(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin {
	s.Nickname = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) SetUserType(v string) *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin {
	s.UserType = &v
	return s
}

func (s *FindUserlistToAuthLoginWithPhoneNumberResponseBodyResultUserListToAuthLogin) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendTestByTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *SendTestByTemplateRequest
	GetAccountName() *string
	SetBirthday(v string) *SendTestByTemplateRequest
	GetBirthday() *string
	SetEmail(v string) *SendTestByTemplateRequest
	GetEmail() *string
	SetGender(v string) *SendTestByTemplateRequest
	GetGender() *string
	SetMobile(v string) *SendTestByTemplateRequest
	GetMobile() *string
	SetNickName(v string) *SendTestByTemplateRequest
	GetNickName() *string
	SetOwnerId(v int64) *SendTestByTemplateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SendTestByTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SendTestByTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v int32) *SendTestByTemplateRequest
	GetTemplateId() *int32
	SetTemplateParams(v string) *SendTestByTemplateRequest
	GetTemplateParams() *string
	SetUserName(v string) *SendTestByTemplateRequest
	GetUserName() *string
}

type SendTestByTemplateRequest struct {
	// The sender address. The length cannot exceed 60 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 账号+@+域名
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The birthday. The length cannot exceed 30 characters.
	//
	// example:
	//
	// 2000/01/01
	Birthday *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	// The recipient email address. The length cannot exceed 60 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 账号+@+域名
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The gender. The length cannot exceed 30 characters.
	//
	// example:
	//
	// 先生
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// The mobile number. The length cannot exceed 30 characters.
	//
	// example:
	//
	// 1380000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// The nickname. The length cannot exceed 30 characters.
	//
	// example:
	//
	// 昵称
	NickName             *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// {"age":"20","nickName":"tom"}
	TemplateParams *string `json:"TemplateParams,omitempty" xml:"TemplateParams,omitempty"`
	// The user name. The length cannot exceed 30 characters.
	//
	// example:
	//
	// 姓名
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s SendTestByTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SendTestByTemplateRequest) GoString() string {
	return s.String()
}

func (s *SendTestByTemplateRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *SendTestByTemplateRequest) GetBirthday() *string {
	return s.Birthday
}

func (s *SendTestByTemplateRequest) GetEmail() *string {
	return s.Email
}

func (s *SendTestByTemplateRequest) GetGender() *string {
	return s.Gender
}

func (s *SendTestByTemplateRequest) GetMobile() *string {
	return s.Mobile
}

func (s *SendTestByTemplateRequest) GetNickName() *string {
	return s.NickName
}

func (s *SendTestByTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SendTestByTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SendTestByTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SendTestByTemplateRequest) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *SendTestByTemplateRequest) GetTemplateParams() *string {
	return s.TemplateParams
}

func (s *SendTestByTemplateRequest) GetUserName() *string {
	return s.UserName
}

func (s *SendTestByTemplateRequest) SetAccountName(v string) *SendTestByTemplateRequest {
	s.AccountName = &v
	return s
}

func (s *SendTestByTemplateRequest) SetBirthday(v string) *SendTestByTemplateRequest {
	s.Birthday = &v
	return s
}

func (s *SendTestByTemplateRequest) SetEmail(v string) *SendTestByTemplateRequest {
	s.Email = &v
	return s
}

func (s *SendTestByTemplateRequest) SetGender(v string) *SendTestByTemplateRequest {
	s.Gender = &v
	return s
}

func (s *SendTestByTemplateRequest) SetMobile(v string) *SendTestByTemplateRequest {
	s.Mobile = &v
	return s
}

func (s *SendTestByTemplateRequest) SetNickName(v string) *SendTestByTemplateRequest {
	s.NickName = &v
	return s
}

func (s *SendTestByTemplateRequest) SetOwnerId(v int64) *SendTestByTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *SendTestByTemplateRequest) SetResourceOwnerAccount(v string) *SendTestByTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendTestByTemplateRequest) SetResourceOwnerId(v int64) *SendTestByTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SendTestByTemplateRequest) SetTemplateId(v int32) *SendTestByTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *SendTestByTemplateRequest) SetTemplateParams(v string) *SendTestByTemplateRequest {
	s.TemplateParams = &v
	return s
}

func (s *SendTestByTemplateRequest) SetUserName(v string) *SendTestByTemplateRequest {
	s.UserName = &v
	return s
}

func (s *SendTestByTemplateRequest) Validate() error {
	return dara.Validate(s)
}

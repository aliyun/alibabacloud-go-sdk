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
	SetUserName(v string) *SendTestByTemplateRequest
	GetUserName() *string
}

type SendTestByTemplateRequest struct {
	// Sender address, with a maximum length of 60 characters
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Birthday, with a maximum length of 30 characters
	//
	// example:
	//
	// 2000/01/01
	Birthday *string `json:"Birthday,omitempty" xml:"Birthday,omitempty"`
	// Recipient address, with a maximum length of 60 characters
	//
	// This parameter is required.
	//
	// example:
	//
	// test1@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Gender, with a maximum length of 30 characters
	//
	// example:
	//
	// doctor
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// Mobile, with a maximum length of 30 characters
	//
	// example:
	//
	// 1380000****
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	// NickName, with a maximum length of 30 characters
	//
	// example:
	//
	// LC
	NickName             *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Template ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// UserName, with a maximum length of 30 characters
	//
	// example:
	//
	// Lucy
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

func (s *SendTestByTemplateRequest) SetUserName(v string) *SendTestByTemplateRequest {
	s.UserName = &v
	return s
}

func (s *SendTestByTemplateRequest) Validate() error {
	return dara.Validate(s)
}

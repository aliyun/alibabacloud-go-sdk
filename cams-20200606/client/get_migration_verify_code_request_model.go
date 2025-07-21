// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationVerifyCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustSpaceId(v string) *GetMigrationVerifyCodeRequest
	GetCustSpaceId() *string
	SetLocale(v string) *GetMigrationVerifyCodeRequest
	GetLocale() *string
	SetMethod(v string) *GetMigrationVerifyCodeRequest
	GetMethod() *string
	SetOwnerId(v int64) *GetMigrationVerifyCodeRequest
	GetOwnerId() *int64
	SetPhoneNumber(v string) *GetMigrationVerifyCodeRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *GetMigrationVerifyCodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetMigrationVerifyCodeRequest
	GetResourceOwnerId() *int64
}

type GetMigrationVerifyCodeRequest struct {
	// The space ID of the user under the independent software vendor (ISV) account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 293483938849493
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// The language.
	//
	// This parameter is required.
	//
	// example:
	//
	// zh_CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
	// The method to obtain the verification code. Valid values: SMS and VOICE.
	//
	// This parameter is required.
	//
	// example:
	//
	// sms
	Method  *string `json:"Method,omitempty" xml:"Method,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetMigrationVerifyCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationVerifyCodeRequest) GoString() string {
	return s.String()
}

func (s *GetMigrationVerifyCodeRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *GetMigrationVerifyCodeRequest) GetLocale() *string {
	return s.Locale
}

func (s *GetMigrationVerifyCodeRequest) GetMethod() *string {
	return s.Method
}

func (s *GetMigrationVerifyCodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetMigrationVerifyCodeRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetMigrationVerifyCodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetMigrationVerifyCodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetMigrationVerifyCodeRequest) SetCustSpaceId(v string) *GetMigrationVerifyCodeRequest {
	s.CustSpaceId = &v
	return s
}

func (s *GetMigrationVerifyCodeRequest) SetLocale(v string) *GetMigrationVerifyCodeRequest {
	s.Locale = &v
	return s
}

func (s *GetMigrationVerifyCodeRequest) SetMethod(v string) *GetMigrationVerifyCodeRequest {
	s.Method = &v
	return s
}

func (s *GetMigrationVerifyCodeRequest) SetOwnerId(v int64) *GetMigrationVerifyCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMigrationVerifyCodeRequest) SetPhoneNumber(v string) *GetMigrationVerifyCodeRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetMigrationVerifyCodeRequest) SetResourceOwnerAccount(v string) *GetMigrationVerifyCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetMigrationVerifyCodeRequest) SetResourceOwnerId(v int64) *GetMigrationVerifyCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetMigrationVerifyCodeRequest) Validate() error {
	return dara.Validate(s)
}

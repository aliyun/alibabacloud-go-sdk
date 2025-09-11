// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidPhoneCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertifyCode(v string) *ValidPhoneCodeRequest
	GetCertifyCode() *string
	SetOwnerId(v int64) *ValidPhoneCodeRequest
	GetOwnerId() *int64
	SetPhoneNo(v string) *ValidPhoneCodeRequest
	GetPhoneNo() *string
	SetResourceOwnerAccount(v string) *ValidPhoneCodeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ValidPhoneCodeRequest
	GetResourceOwnerId() *int64
}

type ValidPhoneCodeRequest struct {
	// 验证码
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	CertifyCode *string `json:"CertifyCode,omitempty" xml:"CertifyCode,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 手机号
	//
	// This parameter is required.
	//
	// example:
	//
	// 137****1234
	PhoneNo              *string `json:"PhoneNo,omitempty" xml:"PhoneNo,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ValidPhoneCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidPhoneCodeRequest) GoString() string {
	return s.String()
}

func (s *ValidPhoneCodeRequest) GetCertifyCode() *string {
	return s.CertifyCode
}

func (s *ValidPhoneCodeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ValidPhoneCodeRequest) GetPhoneNo() *string {
	return s.PhoneNo
}

func (s *ValidPhoneCodeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ValidPhoneCodeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ValidPhoneCodeRequest) SetCertifyCode(v string) *ValidPhoneCodeRequest {
	s.CertifyCode = &v
	return s
}

func (s *ValidPhoneCodeRequest) SetOwnerId(v int64) *ValidPhoneCodeRequest {
	s.OwnerId = &v
	return s
}

func (s *ValidPhoneCodeRequest) SetPhoneNo(v string) *ValidPhoneCodeRequest {
	s.PhoneNo = &v
	return s
}

func (s *ValidPhoneCodeRequest) SetResourceOwnerAccount(v string) *ValidPhoneCodeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ValidPhoneCodeRequest) SetResourceOwnerId(v int64) *ValidPhoneCodeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ValidPhoneCodeRequest) Validate() error {
	return dara.Validate(s)
}

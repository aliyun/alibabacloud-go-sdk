// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCertNoTwoElementVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *CertNoTwoElementVerificationRequest
	GetAuthCode() *string
	SetCertName(v string) *CertNoTwoElementVerificationRequest
	GetCertName() *string
	SetCertNo(v string) *CertNoTwoElementVerificationRequest
	GetCertNo() *string
	SetOwnerId(v int64) *CertNoTwoElementVerificationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CertNoTwoElementVerificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CertNoTwoElementVerificationRequest
	GetResourceOwnerId() *int64
}

type CertNoTwoElementVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// QASDW@#**
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 331021200001010000
	CertNo               *string `json:"CertNo,omitempty" xml:"CertNo,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CertNoTwoElementVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s CertNoTwoElementVerificationRequest) GoString() string {
	return s.String()
}

func (s *CertNoTwoElementVerificationRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *CertNoTwoElementVerificationRequest) GetCertName() *string {
	return s.CertName
}

func (s *CertNoTwoElementVerificationRequest) GetCertNo() *string {
	return s.CertNo
}

func (s *CertNoTwoElementVerificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CertNoTwoElementVerificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CertNoTwoElementVerificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CertNoTwoElementVerificationRequest) SetAuthCode(v string) *CertNoTwoElementVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *CertNoTwoElementVerificationRequest) SetCertName(v string) *CertNoTwoElementVerificationRequest {
	s.CertName = &v
	return s
}

func (s *CertNoTwoElementVerificationRequest) SetCertNo(v string) *CertNoTwoElementVerificationRequest {
	s.CertNo = &v
	return s
}

func (s *CertNoTwoElementVerificationRequest) SetOwnerId(v int64) *CertNoTwoElementVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *CertNoTwoElementVerificationRequest) SetResourceOwnerAccount(v string) *CertNoTwoElementVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CertNoTwoElementVerificationRequest) SetResourceOwnerId(v int64) *CertNoTwoElementVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CertNoTwoElementVerificationRequest) Validate() error {
	return dara.Validate(s)
}

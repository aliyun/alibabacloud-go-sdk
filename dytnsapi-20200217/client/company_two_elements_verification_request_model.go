// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompanyTwoElementsVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *CompanyTwoElementsVerificationRequest
	GetAuthCode() *string
	SetEpCertName(v string) *CompanyTwoElementsVerificationRequest
	GetEpCertName() *string
	SetEpCertNo(v string) *CompanyTwoElementsVerificationRequest
	GetEpCertNo() *string
	SetOwnerId(v int64) *CompanyTwoElementsVerificationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CompanyTwoElementsVerificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CompanyTwoElementsVerificationRequest
	GetResourceOwnerId() *int64
}

type CompanyTwoElementsVerificationRequest struct {
	// The authorization code.
	//
	// >  On the [My Applications](https://dytns.console.aliyun.com/analysis/apply) page in the [Cell Phone Number Service console](https://dytns.console.aliyun.com/overview?spm=a2c4g.608385.0.0.79847f8b3awqUC), you can obtain the authorization code (also known as authorization ID).
	//
	// This parameter is required.
	//
	// example:
	//
	// Dd1r***4id
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The enterprise name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	EpCertName *string `json:"EpCertName,omitempty" xml:"EpCertName,omitempty"`
	// The business license number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9242032*******J627
	EpCertNo             *string `json:"EpCertNo,omitempty" xml:"EpCertNo,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CompanyTwoElementsVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s CompanyTwoElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *CompanyTwoElementsVerificationRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *CompanyTwoElementsVerificationRequest) GetEpCertName() *string {
	return s.EpCertName
}

func (s *CompanyTwoElementsVerificationRequest) GetEpCertNo() *string {
	return s.EpCertNo
}

func (s *CompanyTwoElementsVerificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CompanyTwoElementsVerificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CompanyTwoElementsVerificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CompanyTwoElementsVerificationRequest) SetAuthCode(v string) *CompanyTwoElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *CompanyTwoElementsVerificationRequest) SetEpCertName(v string) *CompanyTwoElementsVerificationRequest {
	s.EpCertName = &v
	return s
}

func (s *CompanyTwoElementsVerificationRequest) SetEpCertNo(v string) *CompanyTwoElementsVerificationRequest {
	s.EpCertNo = &v
	return s
}

func (s *CompanyTwoElementsVerificationRequest) SetOwnerId(v int64) *CompanyTwoElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *CompanyTwoElementsVerificationRequest) SetResourceOwnerAccount(v string) *CompanyTwoElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CompanyTwoElementsVerificationRequest) SetResourceOwnerId(v int64) *CompanyTwoElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CompanyTwoElementsVerificationRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompanyFourElementsVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *CompanyFourElementsVerificationRequest
	GetAuthCode() *string
	SetEpCertName(v string) *CompanyFourElementsVerificationRequest
	GetEpCertName() *string
	SetEpCertNo(v string) *CompanyFourElementsVerificationRequest
	GetEpCertNo() *string
	SetLegalPersonCertName(v string) *CompanyFourElementsVerificationRequest
	GetLegalPersonCertName() *string
	SetLegalPersonCertNo(v string) *CompanyFourElementsVerificationRequest
	GetLegalPersonCertNo() *string
	SetOwnerId(v int64) *CompanyFourElementsVerificationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CompanyFourElementsVerificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CompanyFourElementsVerificationRequest
	GetResourceOwnerId() *int64
}

type CompanyFourElementsVerificationRequest struct {
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
	EpCertNo *string `json:"EpCertNo,omitempty" xml:"EpCertNo,omitempty"`
	// The name of the legal representative.
	//
	// >  If an enterprise has multiple legal representatives, separate them with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	LegalPersonCertName *string `json:"LegalPersonCertName,omitempty" xml:"LegalPersonCertName,omitempty"`
	// The ID card number of the legal representative.
	//
	// >  If an enterprise has multiple legal representatives, separate the ID card numbers with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 331021********0011
	LegalPersonCertNo    *string `json:"LegalPersonCertNo,omitempty" xml:"LegalPersonCertNo,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CompanyFourElementsVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s CompanyFourElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *CompanyFourElementsVerificationRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *CompanyFourElementsVerificationRequest) GetEpCertName() *string {
	return s.EpCertName
}

func (s *CompanyFourElementsVerificationRequest) GetEpCertNo() *string {
	return s.EpCertNo
}

func (s *CompanyFourElementsVerificationRequest) GetLegalPersonCertName() *string {
	return s.LegalPersonCertName
}

func (s *CompanyFourElementsVerificationRequest) GetLegalPersonCertNo() *string {
	return s.LegalPersonCertNo
}

func (s *CompanyFourElementsVerificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CompanyFourElementsVerificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CompanyFourElementsVerificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CompanyFourElementsVerificationRequest) SetAuthCode(v string) *CompanyFourElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetEpCertName(v string) *CompanyFourElementsVerificationRequest {
	s.EpCertName = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetEpCertNo(v string) *CompanyFourElementsVerificationRequest {
	s.EpCertNo = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetLegalPersonCertName(v string) *CompanyFourElementsVerificationRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetLegalPersonCertNo(v string) *CompanyFourElementsVerificationRequest {
	s.LegalPersonCertNo = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetOwnerId(v int64) *CompanyFourElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetResourceOwnerAccount(v string) *CompanyFourElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) SetResourceOwnerId(v int64) *CompanyFourElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CompanyFourElementsVerificationRequest) Validate() error {
	return dara.Validate(s)
}

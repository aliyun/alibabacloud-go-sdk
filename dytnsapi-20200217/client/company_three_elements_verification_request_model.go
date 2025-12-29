// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompanyThreeElementsVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *CompanyThreeElementsVerificationRequest
	GetAuthCode() *string
	SetEpCertName(v string) *CompanyThreeElementsVerificationRequest
	GetEpCertName() *string
	SetEpCertNo(v string) *CompanyThreeElementsVerificationRequest
	GetEpCertNo() *string
	SetLegalPersonCertName(v string) *CompanyThreeElementsVerificationRequest
	GetLegalPersonCertName() *string
	SetOwnerId(v int64) *CompanyThreeElementsVerificationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CompanyThreeElementsVerificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CompanyThreeElementsVerificationRequest
	GetResourceOwnerId() *int64
}

type CompanyThreeElementsVerificationRequest struct {
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
	// 示例值
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
	// 示例值示例值
	LegalPersonCertName  *string `json:"LegalPersonCertName,omitempty" xml:"LegalPersonCertName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CompanyThreeElementsVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s CompanyThreeElementsVerificationRequest) GoString() string {
	return s.String()
}

func (s *CompanyThreeElementsVerificationRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *CompanyThreeElementsVerificationRequest) GetEpCertName() *string {
	return s.EpCertName
}

func (s *CompanyThreeElementsVerificationRequest) GetEpCertNo() *string {
	return s.EpCertNo
}

func (s *CompanyThreeElementsVerificationRequest) GetLegalPersonCertName() *string {
	return s.LegalPersonCertName
}

func (s *CompanyThreeElementsVerificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CompanyThreeElementsVerificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CompanyThreeElementsVerificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CompanyThreeElementsVerificationRequest) SetAuthCode(v string) *CompanyThreeElementsVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetEpCertName(v string) *CompanyThreeElementsVerificationRequest {
	s.EpCertName = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetEpCertNo(v string) *CompanyThreeElementsVerificationRequest {
	s.EpCertNo = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetLegalPersonCertName(v string) *CompanyThreeElementsVerificationRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetOwnerId(v int64) *CompanyThreeElementsVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetResourceOwnerAccount(v string) *CompanyThreeElementsVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) SetResourceOwnerId(v int64) *CompanyThreeElementsVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CompanyThreeElementsVerificationRequest) Validate() error {
	return dara.Validate(s)
}

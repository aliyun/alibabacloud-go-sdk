// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRealNameCertificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountCertifyType(v string) *CreateRealNameCertificationRequest
	GetAccountCertifyType() *string
	SetCorporateLicenseNumber(v string) *CreateRealNameCertificationRequest
	GetCorporateLicenseNumber() *string
	SetCorporateName(v string) *CreateRealNameCertificationRequest
	GetCorporateName() *string
	SetLicenseNumber(v string) *CreateRealNameCertificationRequest
	GetLicenseNumber() *string
	SetLicenseType(v string) *CreateRealNameCertificationRequest
	GetLicenseType() *string
	SetName(v string) *CreateRealNameCertificationRequest
	GetName() *string
	SetPk(v string) *CreateRealNameCertificationRequest
	GetPk() *string
}

type CreateRealNameCertificationRequest struct {
	AccountCertifyType     *string `json:"AccountCertifyType,omitempty" xml:"AccountCertifyType,omitempty"`
	CorporateLicenseNumber *string `json:"CorporateLicenseNumber,omitempty" xml:"CorporateLicenseNumber,omitempty"`
	CorporateName          *string `json:"CorporateName,omitempty" xml:"CorporateName,omitempty"`
	LicenseNumber          *string `json:"LicenseNumber,omitempty" xml:"LicenseNumber,omitempty"`
	LicenseType            *string `json:"LicenseType,omitempty" xml:"LicenseType,omitempty"`
	Name                   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Pk                     *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s CreateRealNameCertificationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRealNameCertificationRequest) GoString() string {
	return s.String()
}

func (s *CreateRealNameCertificationRequest) GetAccountCertifyType() *string {
	return s.AccountCertifyType
}

func (s *CreateRealNameCertificationRequest) GetCorporateLicenseNumber() *string {
	return s.CorporateLicenseNumber
}

func (s *CreateRealNameCertificationRequest) GetCorporateName() *string {
	return s.CorporateName
}

func (s *CreateRealNameCertificationRequest) GetLicenseNumber() *string {
	return s.LicenseNumber
}

func (s *CreateRealNameCertificationRequest) GetLicenseType() *string {
	return s.LicenseType
}

func (s *CreateRealNameCertificationRequest) GetName() *string {
	return s.Name
}

func (s *CreateRealNameCertificationRequest) GetPk() *string {
	return s.Pk
}

func (s *CreateRealNameCertificationRequest) SetAccountCertifyType(v string) *CreateRealNameCertificationRequest {
	s.AccountCertifyType = &v
	return s
}

func (s *CreateRealNameCertificationRequest) SetCorporateLicenseNumber(v string) *CreateRealNameCertificationRequest {
	s.CorporateLicenseNumber = &v
	return s
}

func (s *CreateRealNameCertificationRequest) SetCorporateName(v string) *CreateRealNameCertificationRequest {
	s.CorporateName = &v
	return s
}

func (s *CreateRealNameCertificationRequest) SetLicenseNumber(v string) *CreateRealNameCertificationRequest {
	s.LicenseNumber = &v
	return s
}

func (s *CreateRealNameCertificationRequest) SetLicenseType(v string) *CreateRealNameCertificationRequest {
	s.LicenseType = &v
	return s
}

func (s *CreateRealNameCertificationRequest) SetName(v string) *CreateRealNameCertificationRequest {
	s.Name = &v
	return s
}

func (s *CreateRealNameCertificationRequest) SetPk(v string) *CreateRealNameCertificationRequest {
	s.Pk = &v
	return s
}

func (s *CreateRealNameCertificationRequest) Validate() error {
	return dara.Validate(s)
}

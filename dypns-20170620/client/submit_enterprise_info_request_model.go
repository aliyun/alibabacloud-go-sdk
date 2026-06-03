// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEnterpriseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessLicenseAddress(v string) *SubmitEnterpriseInfoRequest
	GetBusinessLicenseAddress() *string
	SetBusinessLicensePicture(v string) *SubmitEnterpriseInfoRequest
	GetBusinessLicensePicture() *string
	SetEnterpriseId(v int64) *SubmitEnterpriseInfoRequest
	GetEnterpriseId() *int64
	SetEnterpriseName(v string) *SubmitEnterpriseInfoRequest
	GetEnterpriseName() *string
	SetLegalPersonCertNumber(v string) *SubmitEnterpriseInfoRequest
	GetLegalPersonCertNumber() *string
	SetLegalPersonCertPicture(v string) *SubmitEnterpriseInfoRequest
	GetLegalPersonCertPicture() *string
	SetLegalPersonName(v string) *SubmitEnterpriseInfoRequest
	GetLegalPersonName() *string
	SetManagerCertNumber(v string) *SubmitEnterpriseInfoRequest
	GetManagerCertNumber() *string
	SetManagerCertPicture(v string) *SubmitEnterpriseInfoRequest
	GetManagerCertPicture() *string
	SetManagerContactNumber(v string) *SubmitEnterpriseInfoRequest
	GetManagerContactNumber() *string
	SetManagerName(v string) *SubmitEnterpriseInfoRequest
	GetManagerName() *string
	SetNumberApplicationLetterPicture(v string) *SubmitEnterpriseInfoRequest
	GetNumberApplicationLetterPicture() *string
	SetOrganizationCode(v string) *SubmitEnterpriseInfoRequest
	GetOrganizationCode() *string
	SetOwnerId(v int64) *SubmitEnterpriseInfoRequest
	GetOwnerId() *int64
	SetRemark(v string) *SubmitEnterpriseInfoRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *SubmitEnterpriseInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SubmitEnterpriseInfoRequest
	GetResourceOwnerId() *int64
	SetUndertakingPicture(v string) *SubmitEnterpriseInfoRequest
	GetUndertakingPicture() *string
}

type SubmitEnterpriseInfoRequest struct {
	// This parameter is required.
	BusinessLicenseAddress *string `json:"BusinessLicenseAddress,omitempty" xml:"BusinessLicenseAddress,omitempty"`
	// This parameter is required.
	BusinessLicensePicture *string `json:"BusinessLicensePicture,omitempty" xml:"BusinessLicensePicture,omitempty"`
	EnterpriseId           *int64  `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// This parameter is required.
	EnterpriseName *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
	// This parameter is required.
	LegalPersonCertNumber *string `json:"LegalPersonCertNumber,omitempty" xml:"LegalPersonCertNumber,omitempty"`
	// This parameter is required.
	LegalPersonCertPicture *string `json:"LegalPersonCertPicture,omitempty" xml:"LegalPersonCertPicture,omitempty"`
	// This parameter is required.
	LegalPersonName *string `json:"LegalPersonName,omitempty" xml:"LegalPersonName,omitempty"`
	// This parameter is required.
	ManagerCertNumber *string `json:"ManagerCertNumber,omitempty" xml:"ManagerCertNumber,omitempty"`
	// This parameter is required.
	ManagerCertPicture *string `json:"ManagerCertPicture,omitempty" xml:"ManagerCertPicture,omitempty"`
	// This parameter is required.
	ManagerContactNumber *string `json:"ManagerContactNumber,omitempty" xml:"ManagerContactNumber,omitempty"`
	// This parameter is required.
	ManagerName *string `json:"ManagerName,omitempty" xml:"ManagerName,omitempty"`
	// This parameter is required.
	NumberApplicationLetterPicture *string `json:"NumberApplicationLetterPicture,omitempty" xml:"NumberApplicationLetterPicture,omitempty"`
	// This parameter is required.
	OrganizationCode     *string `json:"OrganizationCode,omitempty" xml:"OrganizationCode,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	UndertakingPicture *string `json:"UndertakingPicture,omitempty" xml:"UndertakingPicture,omitempty"`
}

func (s SubmitEnterpriseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitEnterpriseInfoRequest) GoString() string {
	return s.String()
}

func (s *SubmitEnterpriseInfoRequest) GetBusinessLicenseAddress() *string {
	return s.BusinessLicenseAddress
}

func (s *SubmitEnterpriseInfoRequest) GetBusinessLicensePicture() *string {
	return s.BusinessLicensePicture
}

func (s *SubmitEnterpriseInfoRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *SubmitEnterpriseInfoRequest) GetEnterpriseName() *string {
	return s.EnterpriseName
}

func (s *SubmitEnterpriseInfoRequest) GetLegalPersonCertNumber() *string {
	return s.LegalPersonCertNumber
}

func (s *SubmitEnterpriseInfoRequest) GetLegalPersonCertPicture() *string {
	return s.LegalPersonCertPicture
}

func (s *SubmitEnterpriseInfoRequest) GetLegalPersonName() *string {
	return s.LegalPersonName
}

func (s *SubmitEnterpriseInfoRequest) GetManagerCertNumber() *string {
	return s.ManagerCertNumber
}

func (s *SubmitEnterpriseInfoRequest) GetManagerCertPicture() *string {
	return s.ManagerCertPicture
}

func (s *SubmitEnterpriseInfoRequest) GetManagerContactNumber() *string {
	return s.ManagerContactNumber
}

func (s *SubmitEnterpriseInfoRequest) GetManagerName() *string {
	return s.ManagerName
}

func (s *SubmitEnterpriseInfoRequest) GetNumberApplicationLetterPicture() *string {
	return s.NumberApplicationLetterPicture
}

func (s *SubmitEnterpriseInfoRequest) GetOrganizationCode() *string {
	return s.OrganizationCode
}

func (s *SubmitEnterpriseInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SubmitEnterpriseInfoRequest) GetRemark() *string {
	return s.Remark
}

func (s *SubmitEnterpriseInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SubmitEnterpriseInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SubmitEnterpriseInfoRequest) GetUndertakingPicture() *string {
	return s.UndertakingPicture
}

func (s *SubmitEnterpriseInfoRequest) SetBusinessLicenseAddress(v string) *SubmitEnterpriseInfoRequest {
	s.BusinessLicenseAddress = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetBusinessLicensePicture(v string) *SubmitEnterpriseInfoRequest {
	s.BusinessLicensePicture = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetEnterpriseId(v int64) *SubmitEnterpriseInfoRequest {
	s.EnterpriseId = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetEnterpriseName(v string) *SubmitEnterpriseInfoRequest {
	s.EnterpriseName = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetLegalPersonCertNumber(v string) *SubmitEnterpriseInfoRequest {
	s.LegalPersonCertNumber = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetLegalPersonCertPicture(v string) *SubmitEnterpriseInfoRequest {
	s.LegalPersonCertPicture = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetLegalPersonName(v string) *SubmitEnterpriseInfoRequest {
	s.LegalPersonName = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetManagerCertNumber(v string) *SubmitEnterpriseInfoRequest {
	s.ManagerCertNumber = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetManagerCertPicture(v string) *SubmitEnterpriseInfoRequest {
	s.ManagerCertPicture = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetManagerContactNumber(v string) *SubmitEnterpriseInfoRequest {
	s.ManagerContactNumber = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetManagerName(v string) *SubmitEnterpriseInfoRequest {
	s.ManagerName = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetNumberApplicationLetterPicture(v string) *SubmitEnterpriseInfoRequest {
	s.NumberApplicationLetterPicture = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetOrganizationCode(v string) *SubmitEnterpriseInfoRequest {
	s.OrganizationCode = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetOwnerId(v int64) *SubmitEnterpriseInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetRemark(v string) *SubmitEnterpriseInfoRequest {
	s.Remark = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetResourceOwnerAccount(v string) *SubmitEnterpriseInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetResourceOwnerId(v int64) *SubmitEnterpriseInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) SetUndertakingPicture(v string) *SubmitEnterpriseInfoRequest {
	s.UndertakingPicture = &v
	return s
}

func (s *SubmitEnterpriseInfoRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppUrl(v string) *CreateSmsSignRequest
	GetAppUrl() *string
	SetBusinessLicenseCert(v string) *CreateSmsSignRequest
	GetBusinessLicenseCert() *string
	SetCertType(v string) *CreateSmsSignRequest
	GetCertType() *string
	SetExtendMessage(v string) *CreateSmsSignRequest
	GetExtendMessage() *string
	SetOldSmsSignName(v string) *CreateSmsSignRequest
	GetOldSmsSignName() *string
	SetOrderId(v string) *CreateSmsSignRequest
	GetOrderId() *string
	SetOrganizationCodeCert(v string) *CreateSmsSignRequest
	GetOrganizationCodeCert() *string
	SetOwnerId(v int64) *CreateSmsSignRequest
	GetOwnerId() *int64
	SetProdCode(v string) *CreateSmsSignRequest
	GetProdCode() *string
	SetResourceOwnerAccount(v string) *CreateSmsSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsSignRequest
	GetResourceOwnerId() *int64
	SetSmsSignName(v string) *CreateSmsSignRequest
	GetSmsSignName() *string
	SetSmsSignRemark(v string) *CreateSmsSignRequest
	GetSmsSignRemark() *string
	SetSmsSignSource(v string) *CreateSmsSignRequest
	GetSmsSignSource() *string
	SetTaxRegistrationCert(v string) *CreateSmsSignRequest
	GetTaxRegistrationCert() *string
}

type CreateSmsSignRequest struct {
	// This parameter is required.
	AppUrl *string `json:"AppUrl,omitempty" xml:"AppUrl,omitempty"`
	// This parameter is required.
	BusinessLicenseCert *string `json:"BusinessLicenseCert,omitempty" xml:"BusinessLicenseCert,omitempty"`
	// This parameter is required.
	CertType             *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	ExtendMessage        *string `json:"ExtendMessage,omitempty" xml:"ExtendMessage,omitempty"`
	OldSmsSignName       *string `json:"OldSmsSignName,omitempty" xml:"OldSmsSignName,omitempty"`
	OrderId              *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OrganizationCodeCert *string `json:"OrganizationCodeCert,omitempty" xml:"OrganizationCodeCert,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SmsSignName *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	// This parameter is required.
	SmsSignRemark *string `json:"SmsSignRemark,omitempty" xml:"SmsSignRemark,omitempty"`
	// This parameter is required.
	SmsSignSource       *string `json:"SmsSignSource,omitempty" xml:"SmsSignSource,omitempty"`
	TaxRegistrationCert *string `json:"TaxRegistrationCert,omitempty" xml:"TaxRegistrationCert,omitempty"`
}

func (s CreateSmsSignRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsSignRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsSignRequest) GetAppUrl() *string {
	return s.AppUrl
}

func (s *CreateSmsSignRequest) GetBusinessLicenseCert() *string {
	return s.BusinessLicenseCert
}

func (s *CreateSmsSignRequest) GetCertType() *string {
	return s.CertType
}

func (s *CreateSmsSignRequest) GetExtendMessage() *string {
	return s.ExtendMessage
}

func (s *CreateSmsSignRequest) GetOldSmsSignName() *string {
	return s.OldSmsSignName
}

func (s *CreateSmsSignRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateSmsSignRequest) GetOrganizationCodeCert() *string {
	return s.OrganizationCodeCert
}

func (s *CreateSmsSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsSignRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *CreateSmsSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmsSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmsSignRequest) GetSmsSignName() *string {
	return s.SmsSignName
}

func (s *CreateSmsSignRequest) GetSmsSignRemark() *string {
	return s.SmsSignRemark
}

func (s *CreateSmsSignRequest) GetSmsSignSource() *string {
	return s.SmsSignSource
}

func (s *CreateSmsSignRequest) GetTaxRegistrationCert() *string {
	return s.TaxRegistrationCert
}

func (s *CreateSmsSignRequest) SetAppUrl(v string) *CreateSmsSignRequest {
	s.AppUrl = &v
	return s
}

func (s *CreateSmsSignRequest) SetBusinessLicenseCert(v string) *CreateSmsSignRequest {
	s.BusinessLicenseCert = &v
	return s
}

func (s *CreateSmsSignRequest) SetCertType(v string) *CreateSmsSignRequest {
	s.CertType = &v
	return s
}

func (s *CreateSmsSignRequest) SetExtendMessage(v string) *CreateSmsSignRequest {
	s.ExtendMessage = &v
	return s
}

func (s *CreateSmsSignRequest) SetOldSmsSignName(v string) *CreateSmsSignRequest {
	s.OldSmsSignName = &v
	return s
}

func (s *CreateSmsSignRequest) SetOrderId(v string) *CreateSmsSignRequest {
	s.OrderId = &v
	return s
}

func (s *CreateSmsSignRequest) SetOrganizationCodeCert(v string) *CreateSmsSignRequest {
	s.OrganizationCodeCert = &v
	return s
}

func (s *CreateSmsSignRequest) SetOwnerId(v int64) *CreateSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsSignRequest) SetProdCode(v string) *CreateSmsSignRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateSmsSignRequest) SetResourceOwnerAccount(v string) *CreateSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsSignRequest) SetResourceOwnerId(v int64) *CreateSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsSignRequest) SetSmsSignName(v string) *CreateSmsSignRequest {
	s.SmsSignName = &v
	return s
}

func (s *CreateSmsSignRequest) SetSmsSignRemark(v string) *CreateSmsSignRequest {
	s.SmsSignRemark = &v
	return s
}

func (s *CreateSmsSignRequest) SetSmsSignSource(v string) *CreateSmsSignRequest {
	s.SmsSignSource = &v
	return s
}

func (s *CreateSmsSignRequest) SetTaxRegistrationCert(v string) *CreateSmsSignRequest {
	s.TaxRegistrationCert = &v
	return s
}

func (s *CreateSmsSignRequest) Validate() error {
	return dara.Validate(s)
}

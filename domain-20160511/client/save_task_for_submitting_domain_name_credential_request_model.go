// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainNameCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredential(v string) *SaveTaskForSubmittingDomainNameCredentialRequest
	GetCredential() *string
	SetCredentialNo(v string) *SaveTaskForSubmittingDomainNameCredentialRequest
	GetCredentialNo() *string
	SetCredentialType(v string) *SaveTaskForSubmittingDomainNameCredentialRequest
	GetCredentialType() *string
	SetDomainName(v string) *SaveTaskForSubmittingDomainNameCredentialRequest
	GetDomainName() *string
	SetLang(v string) *SaveTaskForSubmittingDomainNameCredentialRequest
	GetLang() *string
	SetSaleId(v string) *SaveTaskForSubmittingDomainNameCredentialRequest
	GetSaleId() *string
	SetUserClientIp(v string) *SaveTaskForSubmittingDomainNameCredentialRequest
	GetUserClientIp() *string
}

type SaveTaskForSubmittingDomainNameCredentialRequest struct {
	// This parameter is required.
	Credential *string `json:"Credential,omitempty" xml:"Credential,omitempty"`
	// This parameter is required.
	CredentialNo   *string `json:"CredentialNo,omitempty" xml:"CredentialNo,omitempty"`
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// This parameter is required.
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SaleId       *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForSubmittingDomainNameCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainNameCredentialRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) GetCredential() *string {
	return s.Credential
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) GetCredentialNo() *string {
	return s.CredentialNo
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) GetCredentialType() *string {
	return s.CredentialType
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) GetSaleId() *string {
	return s.SaleId
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) SetCredential(v string) *SaveTaskForSubmittingDomainNameCredentialRequest {
	s.Credential = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) SetCredentialNo(v string) *SaveTaskForSubmittingDomainNameCredentialRequest {
	s.CredentialNo = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) SetCredentialType(v string) *SaveTaskForSubmittingDomainNameCredentialRequest {
	s.CredentialType = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) SetDomainName(v string) *SaveTaskForSubmittingDomainNameCredentialRequest {
	s.DomainName = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) SetLang(v string) *SaveTaskForSubmittingDomainNameCredentialRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) SetSaleId(v string) *SaveTaskForSubmittingDomainNameCredentialRequest {
	s.SaleId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) SetUserClientIp(v string) *SaveTaskForSubmittingDomainNameCredentialRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialRequest) Validate() error {
	return dara.Validate(s)
}

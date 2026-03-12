// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactTemplateId(v int64) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest
	GetContactTemplateId() *int64
	SetDomainName(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest
	GetDomainName() *string
	SetLang(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest
	GetLang() *string
	SetSaleId(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest
	GetSaleId() *string
	SetUserClientIp(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest
	GetUserClientIp() *string
}

type SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest struct {
	// This parameter is required.
	ContactTemplateId *int64 `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	// This parameter is required.
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SaleId       *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) GetContactTemplateId() *int64 {
	return s.ContactTemplateId
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) GetSaleId() *string {
	return s.SaleId
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) SetContactTemplateId(v int64) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest {
	s.ContactTemplateId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) SetDomainName(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest {
	s.DomainName = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) SetLang(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) SetSaleId(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest {
	s.SaleId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) SetUserClientIp(v string) *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveTaskForSubmittingDomainNameCredentialByTemplateIdRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveContactTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCCity(v string) *SaveContactTemplateRequest
	GetCCity() *string
	SetCCompany(v string) *SaveContactTemplateRequest
	GetCCompany() *string
	SetCCountry(v string) *SaveContactTemplateRequest
	GetCCountry() *string
	SetCName(v string) *SaveContactTemplateRequest
	GetCName() *string
	SetCProvince(v string) *SaveContactTemplateRequest
	GetCProvince() *string
	SetCVenu(v string) *SaveContactTemplateRequest
	GetCVenu() *string
	SetContactTemplateId(v int64) *SaveContactTemplateRequest
	GetContactTemplateId() *int64
	SetDefaultTemplate(v bool) *SaveContactTemplateRequest
	GetDefaultTemplate() *bool
	SetECity(v string) *SaveContactTemplateRequest
	GetECity() *string
	SetECompany(v string) *SaveContactTemplateRequest
	GetECompany() *string
	SetEName(v string) *SaveContactTemplateRequest
	GetEName() *string
	SetEProvince(v string) *SaveContactTemplateRequest
	GetEProvince() *string
	SetEVenu(v string) *SaveContactTemplateRequest
	GetEVenu() *string
	SetEmail(v string) *SaveContactTemplateRequest
	GetEmail() *string
	SetLang(v string) *SaveContactTemplateRequest
	GetLang() *string
	SetPostalCode(v string) *SaveContactTemplateRequest
	GetPostalCode() *string
	SetRegType(v string) *SaveContactTemplateRequest
	GetRegType() *string
	SetTelArea(v string) *SaveContactTemplateRequest
	GetTelArea() *string
	SetTelExt(v string) *SaveContactTemplateRequest
	GetTelExt() *string
	SetTelMain(v string) *SaveContactTemplateRequest
	GetTelMain() *string
	SetUserClientIp(v string) *SaveContactTemplateRequest
	GetUserClientIp() *string
}

type SaveContactTemplateRequest struct {
	CCity             *string `json:"CCity,omitempty" xml:"CCity,omitempty"`
	CCompany          *string `json:"CCompany,omitempty" xml:"CCompany,omitempty"`
	CCountry          *string `json:"CCountry,omitempty" xml:"CCountry,omitempty"`
	CName             *string `json:"CName,omitempty" xml:"CName,omitempty"`
	CProvince         *string `json:"CProvince,omitempty" xml:"CProvince,omitempty"`
	CVenu             *string `json:"CVenu,omitempty" xml:"CVenu,omitempty"`
	ContactTemplateId *int64  `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	DefaultTemplate   *bool   `json:"DefaultTemplate,omitempty" xml:"DefaultTemplate,omitempty"`
	ECity             *string `json:"ECity,omitempty" xml:"ECity,omitempty"`
	ECompany          *string `json:"ECompany,omitempty" xml:"ECompany,omitempty"`
	EName             *string `json:"EName,omitempty" xml:"EName,omitempty"`
	EProvince         *string `json:"EProvince,omitempty" xml:"EProvince,omitempty"`
	EVenu             *string `json:"EVenu,omitempty" xml:"EVenu,omitempty"`
	Email             *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PostalCode        *string `json:"PostalCode,omitempty" xml:"PostalCode,omitempty"`
	RegType           *string `json:"RegType,omitempty" xml:"RegType,omitempty"`
	TelArea           *string `json:"TelArea,omitempty" xml:"TelArea,omitempty"`
	TelExt            *string `json:"TelExt,omitempty" xml:"TelExt,omitempty"`
	TelMain           *string `json:"TelMain,omitempty" xml:"TelMain,omitempty"`
	UserClientIp      *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveContactTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveContactTemplateRequest) GoString() string {
	return s.String()
}

func (s *SaveContactTemplateRequest) GetCCity() *string {
	return s.CCity
}

func (s *SaveContactTemplateRequest) GetCCompany() *string {
	return s.CCompany
}

func (s *SaveContactTemplateRequest) GetCCountry() *string {
	return s.CCountry
}

func (s *SaveContactTemplateRequest) GetCName() *string {
	return s.CName
}

func (s *SaveContactTemplateRequest) GetCProvince() *string {
	return s.CProvince
}

func (s *SaveContactTemplateRequest) GetCVenu() *string {
	return s.CVenu
}

func (s *SaveContactTemplateRequest) GetContactTemplateId() *int64 {
	return s.ContactTemplateId
}

func (s *SaveContactTemplateRequest) GetDefaultTemplate() *bool {
	return s.DefaultTemplate
}

func (s *SaveContactTemplateRequest) GetECity() *string {
	return s.ECity
}

func (s *SaveContactTemplateRequest) GetECompany() *string {
	return s.ECompany
}

func (s *SaveContactTemplateRequest) GetEName() *string {
	return s.EName
}

func (s *SaveContactTemplateRequest) GetEProvince() *string {
	return s.EProvince
}

func (s *SaveContactTemplateRequest) GetEVenu() *string {
	return s.EVenu
}

func (s *SaveContactTemplateRequest) GetEmail() *string {
	return s.Email
}

func (s *SaveContactTemplateRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveContactTemplateRequest) GetPostalCode() *string {
	return s.PostalCode
}

func (s *SaveContactTemplateRequest) GetRegType() *string {
	return s.RegType
}

func (s *SaveContactTemplateRequest) GetTelArea() *string {
	return s.TelArea
}

func (s *SaveContactTemplateRequest) GetTelExt() *string {
	return s.TelExt
}

func (s *SaveContactTemplateRequest) GetTelMain() *string {
	return s.TelMain
}

func (s *SaveContactTemplateRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveContactTemplateRequest) SetCCity(v string) *SaveContactTemplateRequest {
	s.CCity = &v
	return s
}

func (s *SaveContactTemplateRequest) SetCCompany(v string) *SaveContactTemplateRequest {
	s.CCompany = &v
	return s
}

func (s *SaveContactTemplateRequest) SetCCountry(v string) *SaveContactTemplateRequest {
	s.CCountry = &v
	return s
}

func (s *SaveContactTemplateRequest) SetCName(v string) *SaveContactTemplateRequest {
	s.CName = &v
	return s
}

func (s *SaveContactTemplateRequest) SetCProvince(v string) *SaveContactTemplateRequest {
	s.CProvince = &v
	return s
}

func (s *SaveContactTemplateRequest) SetCVenu(v string) *SaveContactTemplateRequest {
	s.CVenu = &v
	return s
}

func (s *SaveContactTemplateRequest) SetContactTemplateId(v int64) *SaveContactTemplateRequest {
	s.ContactTemplateId = &v
	return s
}

func (s *SaveContactTemplateRequest) SetDefaultTemplate(v bool) *SaveContactTemplateRequest {
	s.DefaultTemplate = &v
	return s
}

func (s *SaveContactTemplateRequest) SetECity(v string) *SaveContactTemplateRequest {
	s.ECity = &v
	return s
}

func (s *SaveContactTemplateRequest) SetECompany(v string) *SaveContactTemplateRequest {
	s.ECompany = &v
	return s
}

func (s *SaveContactTemplateRequest) SetEName(v string) *SaveContactTemplateRequest {
	s.EName = &v
	return s
}

func (s *SaveContactTemplateRequest) SetEProvince(v string) *SaveContactTemplateRequest {
	s.EProvince = &v
	return s
}

func (s *SaveContactTemplateRequest) SetEVenu(v string) *SaveContactTemplateRequest {
	s.EVenu = &v
	return s
}

func (s *SaveContactTemplateRequest) SetEmail(v string) *SaveContactTemplateRequest {
	s.Email = &v
	return s
}

func (s *SaveContactTemplateRequest) SetLang(v string) *SaveContactTemplateRequest {
	s.Lang = &v
	return s
}

func (s *SaveContactTemplateRequest) SetPostalCode(v string) *SaveContactTemplateRequest {
	s.PostalCode = &v
	return s
}

func (s *SaveContactTemplateRequest) SetRegType(v string) *SaveContactTemplateRequest {
	s.RegType = &v
	return s
}

func (s *SaveContactTemplateRequest) SetTelArea(v string) *SaveContactTemplateRequest {
	s.TelArea = &v
	return s
}

func (s *SaveContactTemplateRequest) SetTelExt(v string) *SaveContactTemplateRequest {
	s.TelExt = &v
	return s
}

func (s *SaveContactTemplateRequest) SetTelMain(v string) *SaveContactTemplateRequest {
	s.TelMain = &v
	return s
}

func (s *SaveContactTemplateRequest) SetUserClientIp(v string) *SaveContactTemplateRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveContactTemplateRequest) Validate() error {
	return dara.Validate(s)
}

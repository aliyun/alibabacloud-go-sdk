// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingContactByTemplateIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddTransferLock(v bool) *SaveTaskForUpdatingContactByTemplateIdRequest
	GetAddTransferLock() *bool
	SetContactTemplateId(v int64) *SaveTaskForUpdatingContactByTemplateIdRequest
	GetContactTemplateId() *int64
	SetContactType(v string) *SaveTaskForUpdatingContactByTemplateIdRequest
	GetContactType() *string
	SetDomainName(v string) *SaveTaskForUpdatingContactByTemplateIdRequest
	GetDomainName() *string
	SetLang(v string) *SaveTaskForUpdatingContactByTemplateIdRequest
	GetLang() *string
	SetSaleId(v string) *SaveTaskForUpdatingContactByTemplateIdRequest
	GetSaleId() *string
	SetUserClientIp(v string) *SaveTaskForUpdatingContactByTemplateIdRequest
	GetUserClientIp() *string
}

type SaveTaskForUpdatingContactByTemplateIdRequest struct {
	AddTransferLock *bool `json:"AddTransferLock,omitempty" xml:"AddTransferLock,omitempty"`
	// This parameter is required.
	ContactTemplateId *int64 `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	// This parameter is required.
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// This parameter is required.
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SaleId       *string `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForUpdatingContactByTemplateIdRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingContactByTemplateIdRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) GetAddTransferLock() *bool {
	return s.AddTransferLock
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) GetContactTemplateId() *int64 {
	return s.ContactTemplateId
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) GetContactType() *string {
	return s.ContactType
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) GetSaleId() *string {
	return s.SaleId
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) SetAddTransferLock(v bool) *SaveTaskForUpdatingContactByTemplateIdRequest {
	s.AddTransferLock = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) SetContactTemplateId(v int64) *SaveTaskForUpdatingContactByTemplateIdRequest {
	s.ContactTemplateId = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) SetContactType(v string) *SaveTaskForUpdatingContactByTemplateIdRequest {
	s.ContactType = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) SetDomainName(v string) *SaveTaskForUpdatingContactByTemplateIdRequest {
	s.DomainName = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) SetLang(v string) *SaveTaskForUpdatingContactByTemplateIdRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) SetSaleId(v string) *SaveTaskForUpdatingContactByTemplateIdRequest {
	s.SaleId = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) SetUserClientIp(v string) *SaveTaskForUpdatingContactByTemplateIdRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdRequest) Validate() error {
	return dara.Validate(s)
}

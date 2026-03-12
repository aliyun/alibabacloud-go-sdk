// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingContactByTempateIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddTransferLock(v bool) *SaveTaskForUpdatingContactByTempateIdRequest
	GetAddTransferLock() *bool
	SetContactTemplateId(v int64) *SaveTaskForUpdatingContactByTempateIdRequest
	GetContactTemplateId() *int64
	SetContactType(v string) *SaveTaskForUpdatingContactByTempateIdRequest
	GetContactType() *string
	SetDomainName(v string) *SaveTaskForUpdatingContactByTempateIdRequest
	GetDomainName() *string
	SetLang(v string) *SaveTaskForUpdatingContactByTempateIdRequest
	GetLang() *string
	SetSaleId(v string) *SaveTaskForUpdatingContactByTempateIdRequest
	GetSaleId() *string
	SetUserClientIp(v string) *SaveTaskForUpdatingContactByTempateIdRequest
	GetUserClientIp() *string
}

type SaveTaskForUpdatingContactByTempateIdRequest struct {
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

func (s SaveTaskForUpdatingContactByTempateIdRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingContactByTempateIdRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) GetAddTransferLock() *bool {
	return s.AddTransferLock
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) GetContactTemplateId() *int64 {
	return s.ContactTemplateId
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) GetContactType() *string {
	return s.ContactType
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) GetSaleId() *string {
	return s.SaleId
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) SetAddTransferLock(v bool) *SaveTaskForUpdatingContactByTempateIdRequest {
	s.AddTransferLock = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) SetContactTemplateId(v int64) *SaveTaskForUpdatingContactByTempateIdRequest {
	s.ContactTemplateId = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) SetContactType(v string) *SaveTaskForUpdatingContactByTempateIdRequest {
	s.ContactType = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) SetDomainName(v string) *SaveTaskForUpdatingContactByTempateIdRequest {
	s.DomainName = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) SetLang(v string) *SaveTaskForUpdatingContactByTempateIdRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) SetSaleId(v string) *SaveTaskForUpdatingContactByTempateIdRequest {
	s.SaleId = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) SetUserClientIp(v string) *SaveTaskForUpdatingContactByTempateIdRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdRequest) Validate() error {
	return dara.Validate(s)
}

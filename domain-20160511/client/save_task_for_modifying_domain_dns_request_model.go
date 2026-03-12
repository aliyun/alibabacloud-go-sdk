// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForModifyingDomainDnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunDns(v bool) *SaveTaskForModifyingDomainDnsRequest
	GetAliyunDns() *bool
	SetDnsList(v []*string) *SaveTaskForModifyingDomainDnsRequest
	GetDnsList() []*string
	SetDomainName(v string) *SaveTaskForModifyingDomainDnsRequest
	GetDomainName() *string
	SetLang(v string) *SaveTaskForModifyingDomainDnsRequest
	GetLang() *string
	SetSaleId(v string) *SaveTaskForModifyingDomainDnsRequest
	GetSaleId() *string
	SetUserClientIp(v string) *SaveTaskForModifyingDomainDnsRequest
	GetUserClientIp() *string
}

type SaveTaskForModifyingDomainDnsRequest struct {
	// This parameter is required.
	AliyunDns    *bool     `json:"AliyunDns,omitempty" xml:"AliyunDns,omitempty"`
	DnsList      []*string `json:"DnsList,omitempty" xml:"DnsList,omitempty" type:"Repeated"`
	DomainName   *string   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	SaleId       *string   `json:"SaleId,omitempty" xml:"SaleId,omitempty"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveTaskForModifyingDomainDnsRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForModifyingDomainDnsRequest) GoString() string {
	return s.String()
}

func (s *SaveTaskForModifyingDomainDnsRequest) GetAliyunDns() *bool {
	return s.AliyunDns
}

func (s *SaveTaskForModifyingDomainDnsRequest) GetDnsList() []*string {
	return s.DnsList
}

func (s *SaveTaskForModifyingDomainDnsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveTaskForModifyingDomainDnsRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveTaskForModifyingDomainDnsRequest) GetSaleId() *string {
	return s.SaleId
}

func (s *SaveTaskForModifyingDomainDnsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveTaskForModifyingDomainDnsRequest) SetAliyunDns(v bool) *SaveTaskForModifyingDomainDnsRequest {
	s.AliyunDns = &v
	return s
}

func (s *SaveTaskForModifyingDomainDnsRequest) SetDnsList(v []*string) *SaveTaskForModifyingDomainDnsRequest {
	s.DnsList = v
	return s
}

func (s *SaveTaskForModifyingDomainDnsRequest) SetDomainName(v string) *SaveTaskForModifyingDomainDnsRequest {
	s.DomainName = &v
	return s
}

func (s *SaveTaskForModifyingDomainDnsRequest) SetLang(v string) *SaveTaskForModifyingDomainDnsRequest {
	s.Lang = &v
	return s
}

func (s *SaveTaskForModifyingDomainDnsRequest) SetSaleId(v string) *SaveTaskForModifyingDomainDnsRequest {
	s.SaleId = &v
	return s
}

func (s *SaveTaskForModifyingDomainDnsRequest) SetUserClientIp(v string) *SaveTaskForModifyingDomainDnsRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveTaskForModifyingDomainDnsRequest) Validate() error {
	return dara.Validate(s)
}

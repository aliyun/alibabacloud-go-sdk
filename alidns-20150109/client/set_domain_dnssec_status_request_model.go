// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainDnssecStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetDomainDnssecStatusRequest
	GetDomainName() *string
	SetLang(v string) *SetDomainDnssecStatusRequest
	GetLang() *string
	SetStatus(v string) *SetDomainDnssecStatusRequest
	GetStatus() *string
}

type SetDomainDnssecStatusRequest struct {
	// The domain name for which you want to enable the DNSSEC. Only the users of the paid editions of Alibaba Cloud DNS can enable this feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The DNSSEC status. Valid values:
	//
	// 	- ON: enables DNSSEC for the domain name.
	//
	// 	- OFF: disables DNSSEC for the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// ON
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetDomainDnssecStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDomainDnssecStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDomainDnssecStatusRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetDomainDnssecStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *SetDomainDnssecStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *SetDomainDnssecStatusRequest) SetDomainName(v string) *SetDomainDnssecStatusRequest {
	s.DomainName = &v
	return s
}

func (s *SetDomainDnssecStatusRequest) SetLang(v string) *SetDomainDnssecStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetDomainDnssecStatusRequest) SetStatus(v string) *SetDomainDnssecStatusRequest {
	s.Status = &v
	return s
}

func (s *SetDomainDnssecStatusRequest) Validate() error {
	return dara.Validate(s)
}

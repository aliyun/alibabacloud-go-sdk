// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHichinaDomainDNSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ModifyHichinaDomainDNSRequest
	GetDomainName() *string
	SetLang(v string) *ModifyHichinaDomainDNSRequest
	GetLang() *string
	SetUserClientIp(v string) *ModifyHichinaDomainDNSRequest
	GetUserClientIp() *string
}

type ModifyHichinaDomainDNSRequest struct {
	// The domain name.
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
	// Default value: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 192.0.2.0
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ModifyHichinaDomainDNSRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHichinaDomainDNSRequest) GoString() string {
	return s.String()
}

func (s *ModifyHichinaDomainDNSRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyHichinaDomainDNSRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyHichinaDomainDNSRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *ModifyHichinaDomainDNSRequest) SetDomainName(v string) *ModifyHichinaDomainDNSRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyHichinaDomainDNSRequest) SetLang(v string) *ModifyHichinaDomainDNSRequest {
	s.Lang = &v
	return s
}

func (s *ModifyHichinaDomainDNSRequest) SetUserClientIp(v string) *ModifyHichinaDomainDNSRequest {
	s.UserClientIp = &v
	return s
}

func (s *ModifyHichinaDomainDNSRequest) Validate() error {
	return dara.Validate(s)
}

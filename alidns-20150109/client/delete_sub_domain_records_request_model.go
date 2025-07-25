// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubDomainRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteSubDomainRecordsRequest
	GetDomainName() *string
	SetLang(v string) *DeleteSubDomainRecordsRequest
	GetLang() *string
	SetRR(v string) *DeleteSubDomainRecordsRequest
	GetRR() *string
	SetType(v string) *DeleteSubDomainRecordsRequest
	GetType() *string
	SetUserClientIp(v string) *DeleteSubDomainRecordsRequest
	GetUserClientIp() *string
}

type DeleteSubDomainRecordsRequest struct {
	// The domain name. You can call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The hostname field in the DNS record.
	//
	// For example, if you want to resolve @.example.com, you must set this parameter to an at sign (@) instead of leaving it empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	RR *string `json:"RR,omitempty" xml:"RR,omitempty"`
	// The type of DNS records. If you do not specify this parameter, all types of DNS records corresponding to the subdomain are returned.
	//
	// Valid values: **A, MX, CNAME, TXT, REDIRECT_URL, FORWORD_URL, NS, AAAA, and SRV**. The value is not case-sensitive.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 192.0.2.0
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteSubDomainRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubDomainRecordsRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubDomainRecordsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteSubDomainRecordsRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteSubDomainRecordsRequest) GetRR() *string {
	return s.RR
}

func (s *DeleteSubDomainRecordsRequest) GetType() *string {
	return s.Type
}

func (s *DeleteSubDomainRecordsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DeleteSubDomainRecordsRequest) SetDomainName(v string) *DeleteSubDomainRecordsRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteSubDomainRecordsRequest) SetLang(v string) *DeleteSubDomainRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DeleteSubDomainRecordsRequest) SetRR(v string) *DeleteSubDomainRecordsRequest {
	s.RR = &v
	return s
}

func (s *DeleteSubDomainRecordsRequest) SetType(v string) *DeleteSubDomainRecordsRequest {
	s.Type = &v
	return s
}

func (s *DeleteSubDomainRecordsRequest) SetUserClientIp(v string) *DeleteSubDomainRecordsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteSubDomainRecordsRequest) Validate() error {
	return dara.Validate(s)
}

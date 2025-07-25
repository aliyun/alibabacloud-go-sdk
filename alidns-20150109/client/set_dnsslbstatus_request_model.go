// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDNSSLBStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SetDNSSLBStatusRequest
	GetDomainName() *string
	SetLang(v string) *SetDNSSLBStatusRequest
	GetLang() *string
	SetLine(v string) *SetDNSSLBStatusRequest
	GetLine() *string
	SetOpen(v bool) *SetDNSSLBStatusRequest
	GetOpen() *bool
	SetSubDomain(v string) *SetDNSSLBStatusRequest
	GetSubDomain() *string
	SetType(v string) *SetDNSSLBStatusRequest
	GetType() *string
	SetUserClientIp(v string) *SetDNSSLBStatusRequest
	GetUserClientIp() *string
}

type SetDNSSLBStatusRequest struct {
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the content within the request and response. Default: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The DNS resolution line. The line can be the default line, China Telecom, and China Mobile.
	//
	// example:
	//
	// China Mobile.
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// Specifies whether to enable or disable weighted round-robin. Valid values:
	//
	// 	- **true*	- (default): enables weighted round-robin.
	//
	// 	- **false**: disables weighted round-robin.
	//
	// example:
	//
	// true
	Open *bool `json:"Open,omitempty" xml:"Open,omitempty"`
	// The subdomain name for which you want to enable weighted round-robin. Set the parameter to @.example.com instead of example.com.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The type of the Domain Name System (DNS) record. Valid values: A and AAAA. Default value: A.
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

func (s SetDNSSLBStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDNSSLBStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDNSSLBStatusRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SetDNSSLBStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *SetDNSSLBStatusRequest) GetLine() *string {
	return s.Line
}

func (s *SetDNSSLBStatusRequest) GetOpen() *bool {
	return s.Open
}

func (s *SetDNSSLBStatusRequest) GetSubDomain() *string {
	return s.SubDomain
}

func (s *SetDNSSLBStatusRequest) GetType() *string {
	return s.Type
}

func (s *SetDNSSLBStatusRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SetDNSSLBStatusRequest) SetDomainName(v string) *SetDNSSLBStatusRequest {
	s.DomainName = &v
	return s
}

func (s *SetDNSSLBStatusRequest) SetLang(v string) *SetDNSSLBStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetDNSSLBStatusRequest) SetLine(v string) *SetDNSSLBStatusRequest {
	s.Line = &v
	return s
}

func (s *SetDNSSLBStatusRequest) SetOpen(v bool) *SetDNSSLBStatusRequest {
	s.Open = &v
	return s
}

func (s *SetDNSSLBStatusRequest) SetSubDomain(v string) *SetDNSSLBStatusRequest {
	s.SubDomain = &v
	return s
}

func (s *SetDNSSLBStatusRequest) SetType(v string) *SetDNSSLBStatusRequest {
	s.Type = &v
	return s
}

func (s *SetDNSSLBStatusRequest) SetUserClientIp(v string) *SetDNSSLBStatusRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetDNSSLBStatusRequest) Validate() error {
	return dara.Validate(s)
}

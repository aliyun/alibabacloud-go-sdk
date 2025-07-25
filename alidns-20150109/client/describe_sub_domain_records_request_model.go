// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubDomainRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeSubDomainRecordsRequest
	GetDomainName() *string
	SetLang(v string) *DescribeSubDomainRecordsRequest
	GetLang() *string
	SetLine(v string) *DescribeSubDomainRecordsRequest
	GetLine() *string
	SetPageNumber(v int64) *DescribeSubDomainRecordsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSubDomainRecordsRequest
	GetPageSize() *int64
	SetSubDomain(v string) *DescribeSubDomainRecordsRequest
	GetSubDomain() *string
	SetType(v string) *DescribeSubDomainRecordsRequest
	GetType() *string
	SetUserClientIp(v string) *DescribeSubDomainRecordsRequest
	GetUserClientIp() *string
}

type DescribeSubDomainRecordsRequest struct {
	// The domain name.
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
	// The DNS resolution line.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: **1 to 100**. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// If you set SubDomain to `a.www.example.com` and leave
	//
	// DomainName empty, the system returns the DNS records that contain the hostname `a.www` for the domain name example.com. If you set SubDomain to a.www.example.com and set DomainName to www.example.com, the system returns the DNS records that contain the hostname `a` for the domain name www.example.com. If you set SubDomain to a.www.example.com and set DomainName to a.www.example.com, the system returns the DNS records that contain the hostname `@` for the domain name a.www.example.com.
	//
	// This parameter is required.
	//
	// example:
	//
	// a.www.example.com
	SubDomain *string `json:"SubDomain,omitempty" xml:"SubDomain,omitempty"`
	// The type of DNS records. If you do not specify this parameter, all types of DNS records for the subdomain name are returned.
	//
	// Valid values: **A, MX, CNAME, TXT, REDIRECT_URL, FORWORD_URL, NS, AAAA, and SRV**.
	//
	// example:
	//
	// MX
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 192.0.2.0
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeSubDomainRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubDomainRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubDomainRecordsRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeSubDomainRecordsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSubDomainRecordsRequest) GetLine() *string {
	return s.Line
}

func (s *DescribeSubDomainRecordsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSubDomainRecordsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSubDomainRecordsRequest) GetSubDomain() *string {
	return s.SubDomain
}

func (s *DescribeSubDomainRecordsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeSubDomainRecordsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeSubDomainRecordsRequest) SetDomainName(v string) *DescribeSubDomainRecordsRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetLang(v string) *DescribeSubDomainRecordsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetLine(v string) *DescribeSubDomainRecordsRequest {
	s.Line = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetPageNumber(v int64) *DescribeSubDomainRecordsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetPageSize(v int64) *DescribeSubDomainRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetSubDomain(v string) *DescribeSubDomainRecordsRequest {
	s.SubDomain = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetType(v string) *DescribeSubDomainRecordsRequest {
	s.Type = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) SetUserClientIp(v string) *DescribeSubDomainRecordsRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeSubDomainRecordsRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainByDomainNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryDomainByDomainNameRequest
	GetDomainName() *string
	SetLang(v string) *QueryDomainByDomainNameRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryDomainByDomainNameRequest
	GetUserClientIp() *string
}

type QueryDomainByDomainNameRequest struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the error message to return if the request fails. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainByDomainNameRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainByDomainNameRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainByDomainNameRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainByDomainNameRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryDomainByDomainNameRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainByDomainNameRequest) SetDomainName(v string) *QueryDomainByDomainNameRequest {
	s.DomainName = &v
	return s
}

func (s *QueryDomainByDomainNameRequest) SetLang(v string) *QueryDomainByDomainNameRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainByDomainNameRequest) SetUserClientIp(v string) *QueryDomainByDomainNameRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainByDomainNameRequest) Validate() error {
	return dara.Validate(s)
}

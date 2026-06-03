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
	// This parameter is required.
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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

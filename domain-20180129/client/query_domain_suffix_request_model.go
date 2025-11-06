// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainSuffixRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *QueryDomainSuffixRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryDomainSuffixRequest
	GetUserClientIp() *string
}

type QueryDomainSuffixRequest struct {
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainSuffixRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSuffixRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainSuffixRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryDomainSuffixRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainSuffixRequest) SetLang(v string) *QueryDomainSuffixRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainSuffixRequest) SetUserClientIp(v string) *QueryDomainSuffixRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainSuffixRequest) Validate() error {
	return dara.Validate(s)
}

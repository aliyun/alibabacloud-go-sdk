// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainRealNameVerificationInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *QueryDomainRealNameVerificationInfoRequest
	GetDomainName() *string
	SetFetchImage(v bool) *QueryDomainRealNameVerificationInfoRequest
	GetFetchImage() *bool
	SetLang(v string) *QueryDomainRealNameVerificationInfoRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryDomainRealNameVerificationInfoRequest
	GetUserClientIp() *string
}

type QueryDomainRealNameVerificationInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// false
	FetchImage *bool `json:"FetchImage,omitempty" xml:"FetchImage,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainRealNameVerificationInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainRealNameVerificationInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainRealNameVerificationInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainRealNameVerificationInfoRequest) GetFetchImage() *bool {
	return s.FetchImage
}

func (s *QueryDomainRealNameVerificationInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryDomainRealNameVerificationInfoRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainRealNameVerificationInfoRequest) SetDomainName(v string) *QueryDomainRealNameVerificationInfoRequest {
	s.DomainName = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoRequest) SetFetchImage(v bool) *QueryDomainRealNameVerificationInfoRequest {
	s.FetchImage = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoRequest) SetLang(v string) *QueryDomainRealNameVerificationInfoRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoRequest) SetUserClientIp(v string) *QueryDomainRealNameVerificationInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainRealNameVerificationInfoRequest) Validate() error {
	return dara.Validate(s)
}

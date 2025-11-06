// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDomainSunriseClaimRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CheckDomainSunriseClaimRequest
	GetDomainName() *string
	SetLang(v string) *CheckDomainSunriseClaimRequest
	GetLang() *string
	SetUserClientIp(v string) *CheckDomainSunriseClaimRequest
	GetUserClientIp() *string
}

type CheckDomainSunriseClaimRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CheckDomainSunriseClaimRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainSunriseClaimRequest) GoString() string {
	return s.String()
}

func (s *CheckDomainSunriseClaimRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CheckDomainSunriseClaimRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckDomainSunriseClaimRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *CheckDomainSunriseClaimRequest) SetDomainName(v string) *CheckDomainSunriseClaimRequest {
	s.DomainName = &v
	return s
}

func (s *CheckDomainSunriseClaimRequest) SetLang(v string) *CheckDomainSunriseClaimRequest {
	s.Lang = &v
	return s
}

func (s *CheckDomainSunriseClaimRequest) SetUserClientIp(v string) *CheckDomainSunriseClaimRequest {
	s.UserClientIp = &v
	return s
}

func (s *CheckDomainSunriseClaimRequest) Validate() error {
	return dara.Validate(s)
}

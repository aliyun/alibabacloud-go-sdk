// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CheckDomainRequest
	GetDomainName() *string
	SetLang(v string) *CheckDomainRequest
	GetLang() *string
	SetUserClientIp(v string) *CheckDomainRequest
	GetUserClientIp() *string
}

type CheckDomainRequest struct {
	// This parameter is required.
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s CheckDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainRequest) GoString() string {
	return s.String()
}

func (s *CheckDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CheckDomainRequest) GetLang() *string {
	return s.Lang
}

func (s *CheckDomainRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *CheckDomainRequest) SetDomainName(v string) *CheckDomainRequest {
	s.DomainName = &v
	return s
}

func (s *CheckDomainRequest) SetLang(v string) *CheckDomainRequest {
	s.Lang = &v
	return s
}

func (s *CheckDomainRequest) SetUserClientIp(v string) *CheckDomainRequest {
	s.UserClientIp = &v
	return s
}

func (s *CheckDomainRequest) Validate() error {
	return dara.Validate(s)
}

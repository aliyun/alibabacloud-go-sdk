// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainAdminDivisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *QueryDomainAdminDivisionRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryDomainAdminDivisionRequest
	GetUserClientIp() *string
}

type QueryDomainAdminDivisionRequest struct {
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainAdminDivisionRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainAdminDivisionRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainAdminDivisionRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryDomainAdminDivisionRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainAdminDivisionRequest) SetLang(v string) *QueryDomainAdminDivisionRequest {
	s.Lang = &v
	return s
}

func (s *QueryDomainAdminDivisionRequest) SetUserClientIp(v string) *QueryDomainAdminDivisionRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainAdminDivisionRequest) Validate() error {
	return dara.Validate(s)
}

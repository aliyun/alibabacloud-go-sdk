// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContactInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactType(v string) *QueryContactInfoRequest
	GetContactType() *string
	SetDomainName(v string) *QueryContactInfoRequest
	GetDomainName() *string
	SetLang(v string) *QueryContactInfoRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryContactInfoRequest
	GetUserClientIp() *string
}

type QueryContactInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// admin
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
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

func (s QueryContactInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryContactInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryContactInfoRequest) GetContactType() *string {
	return s.ContactType
}

func (s *QueryContactInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryContactInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryContactInfoRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryContactInfoRequest) SetContactType(v string) *QueryContactInfoRequest {
	s.ContactType = &v
	return s
}

func (s *QueryContactInfoRequest) SetDomainName(v string) *QueryContactInfoRequest {
	s.DomainName = &v
	return s
}

func (s *QueryContactInfoRequest) SetLang(v string) *QueryContactInfoRequest {
	s.Lang = &v
	return s
}

func (s *QueryContactInfoRequest) SetUserClientIp(v string) *QueryContactInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryContactInfoRequest) Validate() error {
	return dara.Validate(s)
}

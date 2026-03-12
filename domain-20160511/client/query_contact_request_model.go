// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactType(v string) *QueryContactRequest
	GetContactType() *string
	SetDomainName(v string) *QueryContactRequest
	GetDomainName() *string
	SetLang(v string) *QueryContactRequest
	GetLang() *string
	SetUserClientIp(v string) *QueryContactRequest
	GetUserClientIp() *string
}

type QueryContactRequest struct {
	// This parameter is required.
	ContactType *string `json:"ContactType,omitempty" xml:"ContactType,omitempty"`
	// This parameter is required.
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang         *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryContactRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryContactRequest) GoString() string {
	return s.String()
}

func (s *QueryContactRequest) GetContactType() *string {
	return s.ContactType
}

func (s *QueryContactRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryContactRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryContactRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryContactRequest) SetContactType(v string) *QueryContactRequest {
	s.ContactType = &v
	return s
}

func (s *QueryContactRequest) SetDomainName(v string) *QueryContactRequest {
	s.DomainName = &v
	return s
}

func (s *QueryContactRequest) SetLang(v string) *QueryContactRequest {
	s.Lang = &v
	return s
}

func (s *QueryContactRequest) SetUserClientIp(v string) *QueryContactRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryContactRequest) Validate() error {
	return dara.Validate(s)
}

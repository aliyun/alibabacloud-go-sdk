// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmTransferInEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v []*string) *ConfirmTransferInEmailRequest
	GetDomainName() []*string
	SetEmail(v string) *ConfirmTransferInEmailRequest
	GetEmail() *string
	SetLang(v string) *ConfirmTransferInEmailRequest
	GetLang() *string
	SetUserClientIp(v string) *ConfirmTransferInEmailRequest
	GetUserClientIp() *string
}

type ConfirmTransferInEmailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// abc.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test@test.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s ConfirmTransferInEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmTransferInEmailRequest) GoString() string {
	return s.String()
}

func (s *ConfirmTransferInEmailRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *ConfirmTransferInEmailRequest) GetEmail() *string {
	return s.Email
}

func (s *ConfirmTransferInEmailRequest) GetLang() *string {
	return s.Lang
}

func (s *ConfirmTransferInEmailRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *ConfirmTransferInEmailRequest) SetDomainName(v []*string) *ConfirmTransferInEmailRequest {
	s.DomainName = v
	return s
}

func (s *ConfirmTransferInEmailRequest) SetEmail(v string) *ConfirmTransferInEmailRequest {
	s.Email = &v
	return s
}

func (s *ConfirmTransferInEmailRequest) SetLang(v string) *ConfirmTransferInEmailRequest {
	s.Lang = &v
	return s
}

func (s *ConfirmTransferInEmailRequest) SetUserClientIp(v string) *ConfirmTransferInEmailRequest {
	s.UserClientIp = &v
	return s
}

func (s *ConfirmTransferInEmailRequest) Validate() error {
	return dara.Validate(s)
}

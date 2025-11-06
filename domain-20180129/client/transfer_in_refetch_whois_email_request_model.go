// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInRefetchWhoisEmailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *TransferInRefetchWhoisEmailRequest
	GetDomainName() *string
	SetLang(v string) *TransferInRefetchWhoisEmailRequest
	GetLang() *string
	SetUserClientIp(v string) *TransferInRefetchWhoisEmailRequest
	GetUserClientIp() *string
}

type TransferInRefetchWhoisEmailRequest struct {
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

func (s TransferInRefetchWhoisEmailRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferInRefetchWhoisEmailRequest) GoString() string {
	return s.String()
}

func (s *TransferInRefetchWhoisEmailRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *TransferInRefetchWhoisEmailRequest) GetLang() *string {
	return s.Lang
}

func (s *TransferInRefetchWhoisEmailRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *TransferInRefetchWhoisEmailRequest) SetDomainName(v string) *TransferInRefetchWhoisEmailRequest {
	s.DomainName = &v
	return s
}

func (s *TransferInRefetchWhoisEmailRequest) SetLang(v string) *TransferInRefetchWhoisEmailRequest {
	s.Lang = &v
	return s
}

func (s *TransferInRefetchWhoisEmailRequest) SetUserClientIp(v string) *TransferInRefetchWhoisEmailRequest {
	s.UserClientIp = &v
	return s
}

func (s *TransferInRefetchWhoisEmailRequest) Validate() error {
	return dara.Validate(s)
}

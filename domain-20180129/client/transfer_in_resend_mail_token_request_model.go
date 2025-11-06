// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInResendMailTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *TransferInResendMailTokenRequest
	GetDomainName() *string
	SetLang(v string) *TransferInResendMailTokenRequest
	GetLang() *string
	SetUserClientIp(v string) *TransferInResendMailTokenRequest
	GetUserClientIp() *string
}

type TransferInResendMailTokenRequest struct {
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

func (s TransferInResendMailTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferInResendMailTokenRequest) GoString() string {
	return s.String()
}

func (s *TransferInResendMailTokenRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *TransferInResendMailTokenRequest) GetLang() *string {
	return s.Lang
}

func (s *TransferInResendMailTokenRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *TransferInResendMailTokenRequest) SetDomainName(v string) *TransferInResendMailTokenRequest {
	s.DomainName = &v
	return s
}

func (s *TransferInResendMailTokenRequest) SetLang(v string) *TransferInResendMailTokenRequest {
	s.Lang = &v
	return s
}

func (s *TransferInResendMailTokenRequest) SetUserClientIp(v string) *TransferInResendMailTokenRequest {
	s.UserClientIp = &v
	return s
}

func (s *TransferInResendMailTokenRequest) Validate() error {
	return dara.Validate(s)
}

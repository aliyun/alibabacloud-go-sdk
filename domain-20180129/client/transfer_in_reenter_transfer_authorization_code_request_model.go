// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferInReenterTransferAuthorizationCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *TransferInReenterTransferAuthorizationCodeRequest
	GetDomainName() *string
	SetLang(v string) *TransferInReenterTransferAuthorizationCodeRequest
	GetLang() *string
	SetTransferAuthorizationCode(v string) *TransferInReenterTransferAuthorizationCodeRequest
	GetTransferAuthorizationCode() *string
	SetUserClientIp(v string) *TransferInReenterTransferAuthorizationCodeRequest
	GetUserClientIp() *string
}

type TransferInReenterTransferAuthorizationCodeRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// testCode
	TransferAuthorizationCode *string `json:"TransferAuthorizationCode,omitempty" xml:"TransferAuthorizationCode,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s TransferInReenterTransferAuthorizationCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferInReenterTransferAuthorizationCodeRequest) GoString() string {
	return s.String()
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) GetLang() *string {
	return s.Lang
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) GetTransferAuthorizationCode() *string {
	return s.TransferAuthorizationCode
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) SetDomainName(v string) *TransferInReenterTransferAuthorizationCodeRequest {
	s.DomainName = &v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) SetLang(v string) *TransferInReenterTransferAuthorizationCodeRequest {
	s.Lang = &v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) SetTransferAuthorizationCode(v string) *TransferInReenterTransferAuthorizationCodeRequest {
	s.TransferAuthorizationCode = &v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) SetUserClientIp(v string) *TransferInReenterTransferAuthorizationCodeRequest {
	s.UserClientIp = &v
	return s
}

func (s *TransferInReenterTransferAuthorizationCodeRequest) Validate() error {
	return dara.Validate(s)
}

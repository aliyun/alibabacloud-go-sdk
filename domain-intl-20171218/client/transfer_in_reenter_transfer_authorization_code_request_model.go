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
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Lang       *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	TransferAuthorizationCode *string `json:"TransferAuthorizationCode,omitempty" xml:"TransferAuthorizationCode,omitempty"`
	UserClientIp              *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
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

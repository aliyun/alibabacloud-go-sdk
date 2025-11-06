// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForQueryingTransferAuthorizationCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForQueryingTransferAuthorizationCodeRequest struct {
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

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) SetDomainName(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) SetLang(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) SetUserClientIp(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeRequest) Validate() error {
	return dara.Validate(s)
}

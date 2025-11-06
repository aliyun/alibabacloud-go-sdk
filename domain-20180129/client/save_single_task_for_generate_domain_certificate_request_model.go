// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForGenerateDomainCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForGenerateDomainCertificateRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForGenerateDomainCertificateRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForGenerateDomainCertificateRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForGenerateDomainCertificateRequest struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the error message to return if the request fails. Valid values:
	//
	// 	- **zh**: Chinese.
	//
	// 	- **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForGenerateDomainCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForGenerateDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForGenerateDomainCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForGenerateDomainCertificateRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForGenerateDomainCertificateRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForGenerateDomainCertificateRequest) SetDomainName(v string) *SaveSingleTaskForGenerateDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForGenerateDomainCertificateRequest) SetLang(v string) *SaveSingleTaskForGenerateDomainCertificateRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForGenerateDomainCertificateRequest) SetUserClientIp(v string) *SaveSingleTaskForGenerateDomainCertificateRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForGenerateDomainCertificateRequest) Validate() error {
	return dara.Validate(s)
}

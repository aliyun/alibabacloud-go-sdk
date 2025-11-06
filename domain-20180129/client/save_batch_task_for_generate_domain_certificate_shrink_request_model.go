// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForGenerateDomainCertificateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNamesShrink(v string) *SaveBatchTaskForGenerateDomainCertificateShrinkRequest
	GetDomainNamesShrink() *string
	SetLang(v string) *SaveBatchTaskForGenerateDomainCertificateShrinkRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveBatchTaskForGenerateDomainCertificateShrinkRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForGenerateDomainCertificateShrinkRequest struct {
	// The domain names.
	//
	// This parameter is required.
	DomainNamesShrink *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
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

func (s SaveBatchTaskForGenerateDomainCertificateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForGenerateDomainCertificateShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForGenerateDomainCertificateShrinkRequest) GetDomainNamesShrink() *string {
	return s.DomainNamesShrink
}

func (s *SaveBatchTaskForGenerateDomainCertificateShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForGenerateDomainCertificateShrinkRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForGenerateDomainCertificateShrinkRequest) SetDomainNamesShrink(v string) *SaveBatchTaskForGenerateDomainCertificateShrinkRequest {
	s.DomainNamesShrink = &v
	return s
}

func (s *SaveBatchTaskForGenerateDomainCertificateShrinkRequest) SetLang(v string) *SaveBatchTaskForGenerateDomainCertificateShrinkRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForGenerateDomainCertificateShrinkRequest) SetUserClientIp(v string) *SaveBatchTaskForGenerateDomainCertificateShrinkRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForGenerateDomainCertificateShrinkRequest) Validate() error {
	return dara.Validate(s)
}

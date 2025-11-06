// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForGenerateDomainCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v []*string) *SaveBatchTaskForGenerateDomainCertificateRequest
	GetDomainNames() []*string
	SetLang(v string) *SaveBatchTaskForGenerateDomainCertificateRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveBatchTaskForGenerateDomainCertificateRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForGenerateDomainCertificateRequest struct {
	// The domain names.
	//
	// This parameter is required.
	DomainNames []*string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty" type:"Repeated"`
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

func (s SaveBatchTaskForGenerateDomainCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForGenerateDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForGenerateDomainCertificateRequest) GetDomainNames() []*string {
	return s.DomainNames
}

func (s *SaveBatchTaskForGenerateDomainCertificateRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForGenerateDomainCertificateRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForGenerateDomainCertificateRequest) SetDomainNames(v []*string) *SaveBatchTaskForGenerateDomainCertificateRequest {
	s.DomainNames = v
	return s
}

func (s *SaveBatchTaskForGenerateDomainCertificateRequest) SetLang(v string) *SaveBatchTaskForGenerateDomainCertificateRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForGenerateDomainCertificateRequest) SetUserClientIp(v string) *SaveBatchTaskForGenerateDomainCertificateRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForGenerateDomainCertificateRequest) Validate() error {
	return dara.Validate(s)
}

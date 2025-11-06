// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDomainNameProxyServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *SaveSingleTaskForDomainNameProxyServiceRequest
	GetDomainName() *string
	SetLang(v string) *SaveSingleTaskForDomainNameProxyServiceRequest
	GetLang() *string
	SetStatus(v bool) *SaveSingleTaskForDomainNameProxyServiceRequest
	GetStatus() *bool
	SetUserClientIp(v string) *SaveSingleTaskForDomainNameProxyServiceRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForDomainNameProxyServiceRequest struct {
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
	// false
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForDomainNameProxyServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDomainNameProxyServiceRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) GetStatus() *bool {
	return s.Status
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) SetDomainName(v string) *SaveSingleTaskForDomainNameProxyServiceRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) SetLang(v string) *SaveSingleTaskForDomainNameProxyServiceRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) SetStatus(v bool) *SaveSingleTaskForDomainNameProxyServiceRequest {
	s.Status = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) SetUserClientIp(v string) *SaveSingleTaskForDomainNameProxyServiceRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForDomainNameProxyServiceRequest) Validate() error {
	return dara.Validate(s)
}

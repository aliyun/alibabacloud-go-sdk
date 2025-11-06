// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForDomainNameProxyServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v []*string) *SaveBatchTaskForDomainNameProxyServiceRequest
	GetDomainName() []*string
	SetLang(v string) *SaveBatchTaskForDomainNameProxyServiceRequest
	GetLang() *string
	SetStatus(v bool) *SaveBatchTaskForDomainNameProxyServiceRequest
	GetStatus() *bool
	SetUserClientIp(v string) *SaveBatchTaskForDomainNameProxyServiceRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForDomainNameProxyServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test1.com,test2.com,test3.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
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

func (s SaveBatchTaskForDomainNameProxyServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForDomainNameProxyServiceRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) GetStatus() *bool {
	return s.Status
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetDomainName(v []*string) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetLang(v string) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetStatus(v bool) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.Status = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) SetUserClientIp(v string) *SaveBatchTaskForDomainNameProxyServiceRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForDomainNameProxyServiceRequest) Validate() error {
	return dara.Validate(s)
}

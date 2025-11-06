// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForModifyingDomainDnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunDns(v bool) *SaveBatchTaskForModifyingDomainDnsRequest
	GetAliyunDns() *bool
	SetDomainName(v []*string) *SaveBatchTaskForModifyingDomainDnsRequest
	GetDomainName() []*string
	SetDomainNameServer(v []*string) *SaveBatchTaskForModifyingDomainDnsRequest
	GetDomainNameServer() []*string
	SetLang(v string) *SaveBatchTaskForModifyingDomainDnsRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveBatchTaskForModifyingDomainDnsRequest
	GetUserClientIp() *string
}

type SaveBatchTaskForModifyingDomainDnsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// false
	AliyunDns *bool `json:"AliyunDns,omitempty" xml:"AliyunDns,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// example:
	//
	// ns1.test.com
	DomainNameServer []*string `json:"DomainNameServer,omitempty" xml:"DomainNameServer,omitempty" type:"Repeated"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveBatchTaskForModifyingDomainDnsRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForModifyingDomainDnsRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) GetAliyunDns() *bool {
	return s.AliyunDns
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) GetDomainName() []*string {
	return s.DomainName
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) GetDomainNameServer() []*string {
	return s.DomainNameServer
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) SetAliyunDns(v bool) *SaveBatchTaskForModifyingDomainDnsRequest {
	s.AliyunDns = &v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) SetDomainName(v []*string) *SaveBatchTaskForModifyingDomainDnsRequest {
	s.DomainName = v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) SetDomainNameServer(v []*string) *SaveBatchTaskForModifyingDomainDnsRequest {
	s.DomainNameServer = v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) SetLang(v string) *SaveBatchTaskForModifyingDomainDnsRequest {
	s.Lang = &v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) SetUserClientIp(v string) *SaveBatchTaskForModifyingDomainDnsRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveBatchTaskForModifyingDomainDnsRequest) Validate() error {
	return dara.Validate(s)
}

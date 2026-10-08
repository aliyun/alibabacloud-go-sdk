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
	// Specifies whether to use Alibaba Cloud DNS servers. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	AliyunDns *bool `json:"AliyunDns,omitempty" xml:"AliyunDns,omitempty"`
	// The domain names.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName []*string `json:"DomainName,omitempty" xml:"DomainName,omitempty" type:"Repeated"`
	// The new DNS servers. This parameter is required if **AliyunDns*	- is set to **false**.
	//
	// example:
	//
	// ns1.test.com
	DomainNameServer []*string `json:"DomainNameServer,omitempty" xml:"DomainNameServer,omitempty" type:"Repeated"`
	// The language of API error messages. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value: **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The user IP address. You can set this parameter to **127.0.0.1**.
	//
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForModifyingDnsHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDnsName(v string) *SaveSingleTaskForModifyingDnsHostRequest
	GetDnsName() *string
	SetInstanceId(v string) *SaveSingleTaskForModifyingDnsHostRequest
	GetInstanceId() *string
	SetIp(v []*string) *SaveSingleTaskForModifyingDnsHostRequest
	GetIp() []*string
	SetLang(v string) *SaveSingleTaskForModifyingDnsHostRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForModifyingDnsHostRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForModifyingDnsHostRequest struct {
	// DNS name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dns1
	DnsName *string `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	// Domain instance ID, which can be obtained by invoking the QueryDomainList API.
	//
	// This parameter is required.
	//
	// example:
	//
	// S123456789
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// List of IP addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// 218.xx.xx.236
	Ip []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	// Language for error messages returned by the API. Valid values:
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
	// User IP address.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForModifyingDnsHostRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForModifyingDnsHostRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) GetDnsName() *string {
	return s.DnsName
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) GetIp() []*string {
	return s.Ip
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) SetDnsName(v string) *SaveSingleTaskForModifyingDnsHostRequest {
	s.DnsName = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) SetInstanceId(v string) *SaveSingleTaskForModifyingDnsHostRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) SetIp(v []*string) *SaveSingleTaskForModifyingDnsHostRequest {
	s.Ip = v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) SetLang(v string) *SaveSingleTaskForModifyingDnsHostRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) SetUserClientIp(v string) *SaveSingleTaskForModifyingDnsHostRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostRequest) Validate() error {
	return dara.Validate(s)
}

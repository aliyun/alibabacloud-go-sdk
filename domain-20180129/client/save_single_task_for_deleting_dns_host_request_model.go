// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDeletingDnsHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDnsName(v string) *SaveSingleTaskForDeletingDnsHostRequest
	GetDnsName() *string
	SetInstanceId(v string) *SaveSingleTaskForDeletingDnsHostRequest
	GetInstanceId() *string
	SetLang(v string) *SaveSingleTaskForDeletingDnsHostRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForDeletingDnsHostRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForDeletingDnsHostRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dns1
	DnsName *string `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// S2019270W570xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForDeletingDnsHostRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDeletingDnsHostRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) GetDnsName() *string {
	return s.DnsName
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) SetDnsName(v string) *SaveSingleTaskForDeletingDnsHostRequest {
	s.DnsName = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) SetInstanceId(v string) *SaveSingleTaskForDeletingDnsHostRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) SetLang(v string) *SaveSingleTaskForDeletingDnsHostRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) SetUserClientIp(v string) *SaveSingleTaskForDeletingDnsHostRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingDnsHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDnsName(v string) *SaveSingleTaskForCreatingDnsHostRequest
	GetDnsName() *string
	SetInstanceId(v string) *SaveSingleTaskForCreatingDnsHostRequest
	GetInstanceId() *string
	SetIp(v []*string) *SaveSingleTaskForCreatingDnsHostRequest
	GetIp() []*string
	SetLang(v string) *SaveSingleTaskForCreatingDnsHostRequest
	GetLang() *string
	SetUserClientIp(v string) *SaveSingleTaskForCreatingDnsHostRequest
	GetUserClientIp() *string
}

type SaveSingleTaskForCreatingDnsHostRequest struct {
	// This parameter is required.
	DnsName *string `json:"DnsName,omitempty" xml:"DnsName,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Ip           []*string `json:"Ip,omitempty" xml:"Ip,omitempty" type:"Repeated"`
	Lang         *string   `json:"Lang,omitempty" xml:"Lang,omitempty"`
	UserClientIp *string   `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SaveSingleTaskForCreatingDnsHostRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingDnsHostRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) GetDnsName() *string {
	return s.DnsName
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) GetIp() []*string {
	return s.Ip
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) GetLang() *string {
	return s.Lang
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) SetDnsName(v string) *SaveSingleTaskForCreatingDnsHostRequest {
	s.DnsName = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) SetInstanceId(v string) *SaveSingleTaskForCreatingDnsHostRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) SetIp(v []*string) *SaveSingleTaskForCreatingDnsHostRequest {
	s.Ip = v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) SetLang(v string) *SaveSingleTaskForCreatingDnsHostRequest {
	s.Lang = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) SetUserClientIp(v string) *SaveSingleTaskForCreatingDnsHostRequest {
	s.UserClientIp = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateCloudGtmInstanceConfigRequest
	GetAcceptLanguage() *string
	SetChargeType(v string) *CreateCloudGtmInstanceConfigRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateCloudGtmInstanceConfigRequest
	GetClientToken() *string
	SetEnableStatus(v string) *CreateCloudGtmInstanceConfigRequest
	GetEnableStatus() *string
	SetInstanceId(v string) *CreateCloudGtmInstanceConfigRequest
	GetInstanceId() *string
	SetRemark(v string) *CreateCloudGtmInstanceConfigRequest
	GetRemark() *string
	SetScheduleHostname(v string) *CreateCloudGtmInstanceConfigRequest
	GetScheduleHostname() *string
	SetScheduleRrType(v string) *CreateCloudGtmInstanceConfigRequest
	GetScheduleRrType() *string
	SetScheduleZoneMode(v string) *CreateCloudGtmInstanceConfigRequest
	GetScheduleZoneMode() *string
	SetScheduleZoneName(v string) *CreateCloudGtmInstanceConfigRequest
	GetScheduleZoneName() *string
	SetTtl(v int32) *CreateCloudGtmInstanceConfigRequest
	GetTtl() *int32
}

type CreateCloudGtmInstanceConfigRequest struct {
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The billing method for the instance configuration:
	//
	// - prepay: Subscription. This is the default value.
	//
	// - postpay: Pay-as-you-go.
	//
	// example:
	//
	// postpay
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can specify a custom value, but you must make sure that the value is unique among different requests. The token can contain up to 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The status of the domain name instance:
	//
	// - enable: Enabled. The intelligent scheduling policy of the GTM instance is active.
	//
	// - disable: Disabled. The intelligent scheduling policy of the GTM instance is unavailable.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The unique ID of the GTM 3.0 instance.
	//
	// example:
	//
	// gtm-cn-jmp3qnw**03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The remark.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The host record of the GTM access domain name.
	//
	// example:
	//
	// www
	ScheduleHostname *string `json:"ScheduleHostname,omitempty" xml:"ScheduleHostname,omitempty"`
	// The DNS record type of the access domain name:
	//
	// - A: IPv4 address
	//
	// - AAAA: IPv6 address
	//
	// - CNAME: Canonical name
	//
	// example:
	//
	// A
	ScheduleRrType *string `json:"ScheduleRrType,omitempty" xml:"ScheduleRrType,omitempty"`
	// The configuration mode for the access domain name:
	//
	// - sys_assign: The system assigns a domain name. This mode is no longer supported.
	//
	// - custom: Custom mode. Select a domain name under the account that owns the instance and enter a host record to generate the access domain name.
	//
	// example:
	//
	// custom
	ScheduleZoneMode *string `json:"ScheduleZoneMode,omitempty" xml:"ScheduleZoneMode,omitempty"`
	// The zone name, which is the parent zone of the GTM access domain name. This is typically a domain name hosted in the Alibaba Cloud DNS console under the account that owns the GTM instance. Primary domains and subdomains are supported.
	//
	// example:
	//
	// example.com
	ScheduleZoneName *string `json:"ScheduleZoneName,omitempty" xml:"ScheduleZoneName,omitempty"`
	// The global Time to Live (TTL) in seconds. This is the TTL for the access domain name that resolves to an address in an address pool. This value affects how long the DNS record is cached on a carrier\\"s local DNS server. You can specify a custom TTL.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s CreateCloudGtmInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmInstanceConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateCloudGtmInstanceConfigRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateCloudGtmInstanceConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCloudGtmInstanceConfigRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *CreateCloudGtmInstanceConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCloudGtmInstanceConfigRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateCloudGtmInstanceConfigRequest) GetScheduleHostname() *string {
	return s.ScheduleHostname
}

func (s *CreateCloudGtmInstanceConfigRequest) GetScheduleRrType() *string {
	return s.ScheduleRrType
}

func (s *CreateCloudGtmInstanceConfigRequest) GetScheduleZoneMode() *string {
	return s.ScheduleZoneMode
}

func (s *CreateCloudGtmInstanceConfigRequest) GetScheduleZoneName() *string {
	return s.ScheduleZoneName
}

func (s *CreateCloudGtmInstanceConfigRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *CreateCloudGtmInstanceConfigRequest) SetAcceptLanguage(v string) *CreateCloudGtmInstanceConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigRequest) SetChargeType(v string) *CreateCloudGtmInstanceConfigRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigRequest) SetClientToken(v string) *CreateCloudGtmInstanceConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigRequest) SetEnableStatus(v string) *CreateCloudGtmInstanceConfigRequest {
	s.EnableStatus = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigRequest) SetInstanceId(v string) *CreateCloudGtmInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigRequest) SetRemark(v string) *CreateCloudGtmInstanceConfigRequest {
	s.Remark = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigRequest) SetScheduleHostname(v string) *CreateCloudGtmInstanceConfigRequest {
	s.ScheduleHostname = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigRequest) SetScheduleRrType(v string) *CreateCloudGtmInstanceConfigRequest {
	s.ScheduleRrType = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigRequest) SetScheduleZoneMode(v string) *CreateCloudGtmInstanceConfigRequest {
	s.ScheduleZoneMode = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigRequest) SetScheduleZoneName(v string) *CreateCloudGtmInstanceConfigRequest {
	s.ScheduleZoneName = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigRequest) SetTtl(v int32) *CreateCloudGtmInstanceConfigRequest {
	s.Ttl = &v
	return s
}

func (s *CreateCloudGtmInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}

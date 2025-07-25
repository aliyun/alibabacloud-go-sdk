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
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can specify a custom value for this parameter, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enabling state of the access domain name. Valid values:
	//
	// 	- enable: The access domain name is enabled and the intelligent scheduling policy of the corresponding GTM instance takes effect.
	//
	// 	- disable: The access domain name is disabled and the intelligent scheduling policy of the corresponding GTM instance is unavailable.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The ID of the Global Traffic Manager (GTM) 3.0 instance. This ID uniquely identifies a GTM 3.0 instance.
	//
	// example:
	//
	// gtm-cn-jmp3qnw**03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The description of the access domain name.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The hostname of the access domain name.
	//
	// example:
	//
	// www
	ScheduleHostname *string `json:"ScheduleHostname,omitempty" xml:"ScheduleHostname,omitempty"`
	// The type of the Domain Name System (DNS) record configured for the access domain name. Valid values:
	//
	// 	- A: IPv4 address
	//
	// 	- AAAA: IPv6 address
	//
	// 	- CNAME: domain name
	//
	// example:
	//
	// A
	ScheduleRrType *string `json:"ScheduleRrType,omitempty" xml:"ScheduleRrType,omitempty"`
	// The configuration mode of the access domain name. Valid values:
	//
	// 	- sys_assign: system allocation. This mode is not supported.
	//
	// 	- custom: custom allocation. You must select a zone within the account to which the instance belongs and enter a hostname to generate an access domain name.
	//
	// example:
	//
	// custom
	ScheduleZoneMode *string `json:"ScheduleZoneMode,omitempty" xml:"ScheduleZoneMode,omitempty"`
	// The name of the parent zone for the access domain name configured in GTM. In most cases, the value of this parameter is the name of a zone hosted by Alibaba Cloud DNS. This zone belongs to the account to which the GTM instance belongs. You can specify the name of a zone or subzone.
	//
	// example:
	//
	// example.com
	ScheduleZoneName *string `json:"ScheduleZoneName,omitempty" xml:"ScheduleZoneName,omitempty"`
	// The global time to live (TTL) period. Unit: seconds. The global TTL period affects how long the DNS records that map the access domain name to the addresses in the address pools are cached in the local DNS servers of Internet service providers (ISPs). You can specify a custom value.
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

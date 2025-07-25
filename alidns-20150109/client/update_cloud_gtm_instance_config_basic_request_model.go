// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigBasicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigBasicRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *UpdateCloudGtmInstanceConfigBasicRequest
	GetClientToken() *string
	SetConfigId(v string) *UpdateCloudGtmInstanceConfigBasicRequest
	GetConfigId() *string
	SetInstanceId(v string) *UpdateCloudGtmInstanceConfigBasicRequest
	GetInstanceId() *string
	SetScheduleHostname(v string) *UpdateCloudGtmInstanceConfigBasicRequest
	GetScheduleHostname() *string
	SetScheduleZoneName(v string) *UpdateCloudGtmInstanceConfigBasicRequest
	GetScheduleZoneName() *string
	SetTtl(v int32) *UpdateCloudGtmInstanceConfigBasicRequest
	GetTtl() *int32
}

type UpdateCloudGtmInstanceConfigBasicRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US*	- (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration ID of the access domain name. Two configuration IDs exist when the access domain name is bound to the same GTM instance but an A record and an AAAA record are configured for the access domain name. The configuration ID uniquely identifies a configuration.
	//
	// You can call the [ListCloudGtmInstanceConfigs](https://help.aliyun.com/document_detail/2797349.html) operation to query the value of ConfigId for the access domain name.
	//
	// example:
	//
	// Config-000**11
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the GTM 3.0 instance for which you want to modify the TTL configuration.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Host record of the domain accessed by GTM.
	//
	// example:
	//
	// www
	ScheduleHostname *string `json:"ScheduleHostname,omitempty" xml:"ScheduleHostname,omitempty"`
	// The zone (such as example.com) or subzone (such as a.example.com) of the GTM access domain name. In most cases, the zone or subzone is hosted in Authoritative DNS Resolution of the Alibaba Cloud DNS console within the account to which the GTM instance belongs.
	//
	// example:
	//
	// example.com
	ScheduleZoneName *string `json:"ScheduleZoneName,omitempty" xml:"ScheduleZoneName,omitempty"`
	// The global TTL value, in seconds. The global TTL value affects how long the DNS records that map the access domain name to the addresses in the address pools are cached in the local DNS servers of Internet service providers (ISPs).
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigBasicRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigBasicRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) GetScheduleHostname() *string {
	return s.ScheduleHostname
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) GetScheduleZoneName() *string {
	return s.ScheduleZoneName
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigBasicRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) SetClientToken(v string) *UpdateCloudGtmInstanceConfigBasicRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) SetConfigId(v string) *UpdateCloudGtmInstanceConfigBasicRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) SetInstanceId(v string) *UpdateCloudGtmInstanceConfigBasicRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) SetScheduleHostname(v string) *UpdateCloudGtmInstanceConfigBasicRequest {
	s.ScheduleHostname = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) SetScheduleZoneName(v string) *UpdateCloudGtmInstanceConfigBasicRequest {
	s.ScheduleZoneName = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) SetTtl(v int32) *UpdateCloudGtmInstanceConfigBasicRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigBasicRequest) Validate() error {
	return dara.Validate(s)
}

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
	// - **zh-CN**: Chinese.
	//
	// - **en-US*	- (default): English.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// A client-generated token that you can use to ensure the idempotence of the request. Make sure that the token is unique among different requests. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the domain name instance configuration. For the same access domain name and GTM instance, you can configure both A and AAAA records. This results in two domain name instance configurations for the same GTM instance. The ConfigId uniquely identifies the configuration object that you want to modify.
	//
	// Call the [ListCloudGtmInstanceConfigs](~~ListCloudGtmInstanceConfigs~~) operation to query the ConfigId of a domain name instance.
	//
	// example:
	//
	// Config-000****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the GTM 3.0 instance that you want to modify.
	//
	// example:
	//
	// gtm-cn-wwo3a3h****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The host record of the GTM access domain name.
	//
	// example:
	//
	// www
	ScheduleHostname *string `json:"ScheduleHostname,omitempty" xml:"ScheduleHostname,omitempty"`
	// The root domain (such as example.com) or subdomain (such as a.example.com) of the GTM access domain name. This is usually a domain name that is hosted in the authoritative zone of the Alibaba Cloud DNS console under the account that owns the GTM instance.
	//
	// example:
	//
	// example.com
	ScheduleZoneName *string `json:"ScheduleZoneName,omitempty" xml:"ScheduleZoneName,omitempty"`
	// The global Time to Live (TTL) in seconds. This is the TTL for the DNS record that resolves the access domain name to an address in an address pool. The TTL affects how long the DNS record is cached on a carrier\\"s Local DNS server.
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

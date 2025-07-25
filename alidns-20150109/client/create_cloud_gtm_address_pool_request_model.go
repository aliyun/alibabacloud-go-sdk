// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *CreateCloudGtmAddressPoolRequest
	GetAcceptLanguage() *string
	SetAddressPoolName(v string) *CreateCloudGtmAddressPoolRequest
	GetAddressPoolName() *string
	SetAddressPoolType(v string) *CreateCloudGtmAddressPoolRequest
	GetAddressPoolType() *string
	SetClientToken(v string) *CreateCloudGtmAddressPoolRequest
	GetClientToken() *string
	SetEnableStatus(v string) *CreateCloudGtmAddressPoolRequest
	GetEnableStatus() *string
	SetHealthJudgement(v string) *CreateCloudGtmAddressPoolRequest
	GetHealthJudgement() *string
	SetRemark(v string) *CreateCloudGtmAddressPoolRequest
	GetRemark() *string
}

type CreateCloudGtmAddressPoolRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Address pool name, helping users distinguish the purpose of address pools.
	//
	// example:
	//
	// Address pool-1
	AddressPoolName *string `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	// The type of the address pool. Valid values:
	//
	// 	- IPv4: IPv4 addresses are returned for Domain Name System (DNS) resolution.
	//
	// 	- IPv6: IPv6 addresses are returned for DNS resolution.
	//
	// 	- domain: Domain names are returned for DNS resolution.
	//
	// example:
	//
	// IPv4
	AddressPoolType *string `json:"AddressPoolType,omitempty" xml:"AddressPoolType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enabling state of the address pool. Valid values:
	//
	// 	- enable: The address pool is enabled, and the addresses in the address pool are returned for DNS resolution when the health check results are normal.
	//
	// 	- disable: The address pool is disabled, and the addresses in the address pool are not returned for DNS resolution regardless of whether the health check results are normal or not.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The condition for determining the health status of the address pool. Valid values:
	//
	// 	- any_ok: At least one address in the address pool is available.
	//
	// 	- p30_ok: At least 30% of the addresses in the address pool are available.
	//
	// 	- p50_ok: At least 50% of the addresses in the address pool are available.
	//
	// 	- p70_ok: At least 70% of the addresses in the address pool are available.
	//
	// 	- all_ok: All addresses in the address pool are available.
	//
	// example:
	//
	// any_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// Remarks for the address pool, helping users distinguish the usage scenarios of different address pools.
	//
	// example:
	//
	// app
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateCloudGtmAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmAddressPoolRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *CreateCloudGtmAddressPoolRequest) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *CreateCloudGtmAddressPoolRequest) GetAddressPoolType() *string {
	return s.AddressPoolType
}

func (s *CreateCloudGtmAddressPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCloudGtmAddressPoolRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *CreateCloudGtmAddressPoolRequest) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *CreateCloudGtmAddressPoolRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateCloudGtmAddressPoolRequest) SetAcceptLanguage(v string) *CreateCloudGtmAddressPoolRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *CreateCloudGtmAddressPoolRequest) SetAddressPoolName(v string) *CreateCloudGtmAddressPoolRequest {
	s.AddressPoolName = &v
	return s
}

func (s *CreateCloudGtmAddressPoolRequest) SetAddressPoolType(v string) *CreateCloudGtmAddressPoolRequest {
	s.AddressPoolType = &v
	return s
}

func (s *CreateCloudGtmAddressPoolRequest) SetClientToken(v string) *CreateCloudGtmAddressPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCloudGtmAddressPoolRequest) SetEnableStatus(v string) *CreateCloudGtmAddressPoolRequest {
	s.EnableStatus = &v
	return s
}

func (s *CreateCloudGtmAddressPoolRequest) SetHealthJudgement(v string) *CreateCloudGtmAddressPoolRequest {
	s.HealthJudgement = &v
	return s
}

func (s *CreateCloudGtmAddressPoolRequest) SetRemark(v string) *CreateCloudGtmAddressPoolRequest {
	s.Remark = &v
	return s
}

func (s *CreateCloudGtmAddressPoolRequest) Validate() error {
	return dara.Validate(s)
}

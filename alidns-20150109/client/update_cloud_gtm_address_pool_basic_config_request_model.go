// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolBasicConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmAddressPoolBasicConfigRequest
	GetAcceptLanguage() *string
	SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolBasicConfigRequest
	GetAddressPoolId() *string
	SetAddressPoolName(v string) *UpdateCloudGtmAddressPoolBasicConfigRequest
	GetAddressPoolName() *string
	SetClientToken(v string) *UpdateCloudGtmAddressPoolBasicConfigRequest
	GetClientToken() *string
	SetHealthJudgement(v string) *UpdateCloudGtmAddressPoolBasicConfigRequest
	GetHealthJudgement() *string
}

type UpdateCloudGtmAddressPoolBasicConfigRequest struct {
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
	// The ID of the address pool. This ID uniquely identifies the address pool.
	//
	// example:
	//
	// pool-89528023225442**16
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// Address pool name, helping users distinguish the purpose of address pools.
	//
	// example:
	//
	// app
	AddressPoolName *string `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
}

func (s UpdateCloudGtmAddressPoolBasicConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolBasicConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolBasicConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmAddressPoolBasicConfigRequest) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *UpdateCloudGtmAddressPoolBasicConfigRequest) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *UpdateCloudGtmAddressPoolBasicConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmAddressPoolBasicConfigRequest) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *UpdateCloudGtmAddressPoolBasicConfigRequest) SetAcceptLanguage(v string) *UpdateCloudGtmAddressPoolBasicConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolBasicConfigRequest) SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolBasicConfigRequest {
	s.AddressPoolId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolBasicConfigRequest) SetAddressPoolName(v string) *UpdateCloudGtmAddressPoolBasicConfigRequest {
	s.AddressPoolName = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolBasicConfigRequest) SetClientToken(v string) *UpdateCloudGtmAddressPoolBasicConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolBasicConfigRequest) SetHealthJudgement(v string) *UpdateCloudGtmAddressPoolBasicConfigRequest {
	s.HealthJudgement = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolBasicConfigRequest) Validate() error {
	return dara.Validate(s)
}

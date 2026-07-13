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
	// - zh-CN: Chinese
	//
	// - en-US (default): English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the address pool.
	//
	// example:
	//
	// pool-89528023225442****
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// app
	AddressPoolName *string `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	// The client token that is used to ensure the idempotence of the request. The client generates this token to make sure that each request is unique. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The condition for determining the health status of the address pool.
	//
	// - any_ok: At least one address in the address pool is active.
	//
	// - p30_ok: At least 30% of the addresses in the address pool are active.
	//
	// - p50_ok: At least 50% of the addresses in the address pool are active.
	//
	// - p70_ok: At least 70% of the addresses in the address pool are active.
	//
	// - all_ok: All addresses in the address pool are active.
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

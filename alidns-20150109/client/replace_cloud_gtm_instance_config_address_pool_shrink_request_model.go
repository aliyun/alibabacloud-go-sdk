// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest
	GetAcceptLanguage() *string
	SetAddressPoolsShrink(v string) *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest
	GetAddressPoolsShrink() *string
	SetClientToken(v string) *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest
	GetClientToken() *string
	SetConfigId(v string) *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest
	GetConfigId() *string
	SetInstanceId(v string) *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest
	GetInstanceId() *string
}

type ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest struct {
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
	// The address pools.
	AddressPoolsShrink *string `json:"AddressPools,omitempty" xml:"AddressPools,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration ID of the access domain name. Two configuration IDs exist when the access domain name is bound to the same GTM instance but an A record and an AAAA record are configured for the access domain name. The configuration ID uniquely identifies a configuration.
	//
	// You can call the [ListCloudGtmInstanceConfigs](~~ListCloudGtmInstanceConfigs~~) operation to query the configuration ID of the access domain name.
	//
	// example:
	//
	// Config-000**11
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the GTM 3.0 instance for which you want to change address pools.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) GetAddressPoolsShrink() *string {
	return s.AddressPoolsShrink
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) SetAcceptLanguage(v string) *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) SetAddressPoolsShrink(v string) *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest {
	s.AddressPoolsShrink = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) SetClientToken(v string) *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) SetConfigId(v string) *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest {
	s.ConfigId = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) SetInstanceId(v string) *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ReplaceCloudGtmInstanceConfigAddressPoolShrinkRequest) Validate() error {
	return dara.Validate(s)
}

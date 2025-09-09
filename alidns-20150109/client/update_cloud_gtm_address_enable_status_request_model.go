// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressEnableStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmAddressEnableStatusRequest
	GetAcceptLanguage() *string
	SetAddressId(v string) *UpdateCloudGtmAddressEnableStatusRequest
	GetAddressId() *string
	SetClientToken(v string) *UpdateCloudGtmAddressEnableStatusRequest
	GetClientToken() *string
	SetEnableStatus(v string) *UpdateCloudGtmAddressEnableStatusRequest
	GetEnableStatus() *string
}

type UpdateCloudGtmAddressEnableStatusRequest struct {
	// The language of the returned results. Valid values:
	//
	// - zh-CN: Chinese
	//
	// - en-US: English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the address. This ID uniquely identifies the address.
	//
	// This parameter is required.
	//
	// example:
	//
	// addr-89518218114368**92
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can specify a custom value for this parameter, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enabling state of the address. Valid values:
	//
	// 	- enable: The address is enabled and the address can be used for Domain Name System (DNS) resolution if the address passes health checks.
	//
	// 	- disable: The address is disabled and the address cannot be used for DNS resolution regardless of whether the address passes health checks or not.
	//
	// This parameter is required.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
}

func (s UpdateCloudGtmAddressEnableStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressEnableStatusRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmAddressEnableStatusRequest) GetAddressId() *string {
	return s.AddressId
}

func (s *UpdateCloudGtmAddressEnableStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmAddressEnableStatusRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *UpdateCloudGtmAddressEnableStatusRequest) SetAcceptLanguage(v string) *UpdateCloudGtmAddressEnableStatusRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmAddressEnableStatusRequest) SetAddressId(v string) *UpdateCloudGtmAddressEnableStatusRequest {
	s.AddressId = &v
	return s
}

func (s *UpdateCloudGtmAddressEnableStatusRequest) SetClientToken(v string) *UpdateCloudGtmAddressEnableStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmAddressEnableStatusRequest) SetEnableStatus(v string) *UpdateCloudGtmAddressEnableStatusRequest {
	s.EnableStatus = &v
	return s
}

func (s *UpdateCloudGtmAddressEnableStatusRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolEnableStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmAddressPoolEnableStatusRequest
	GetAcceptLanguage() *string
	SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolEnableStatusRequest
	GetAddressPoolId() *string
	SetClientToken(v string) *UpdateCloudGtmAddressPoolEnableStatusRequest
	GetClientToken() *string
	SetEnableStatus(v string) *UpdateCloudGtmAddressPoolEnableStatusRequest
	GetEnableStatus() *string
}

type UpdateCloudGtmAddressPoolEnableStatusRequest struct {
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
}

func (s UpdateCloudGtmAddressPoolEnableStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolEnableStatusRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmAddressPoolEnableStatusRequest) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *UpdateCloudGtmAddressPoolEnableStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmAddressPoolEnableStatusRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *UpdateCloudGtmAddressPoolEnableStatusRequest) SetAcceptLanguage(v string) *UpdateCloudGtmAddressPoolEnableStatusRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolEnableStatusRequest) SetAddressPoolId(v string) *UpdateCloudGtmAddressPoolEnableStatusRequest {
	s.AddressPoolId = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolEnableStatusRequest) SetClientToken(v string) *UpdateCloudGtmAddressPoolEnableStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolEnableStatusRequest) SetEnableStatus(v string) *UpdateCloudGtmAddressPoolEnableStatusRequest {
	s.EnableStatus = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolEnableStatusRequest) Validate() error {
	return dara.Validate(s)
}

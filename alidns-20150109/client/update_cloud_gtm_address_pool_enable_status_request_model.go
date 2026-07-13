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
	// A client token to ensure the idempotence of the request. Generate a unique value from your client for this parameter. The client token can contain only ASCII characters and must be no more than 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enabled status of the address pool:
	//
	// - enable: Enables the address pool. If the health check is normal, the address pool is included in DNS resolution.
	//
	// - disable: Disables the address pool. The address pool is not included in DNS resolution, regardless of its health check status.
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

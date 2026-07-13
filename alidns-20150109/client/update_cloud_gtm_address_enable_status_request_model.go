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
	// The language of the response. Valid values:
	//
	// - zh-CN: Chinese
	//
	// - en-US: English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the address.
	//
	// This parameter is required.
	//
	// example:
	//
	// addr-89518218114368****
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// A client-generated token that is used to ensure the idempotence of the request. Make sure that the token is unique among different requests. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The enabled status of the address:
	//
	// - enable: The address can be used for DNS resolution if its health check is normal.
	//
	// - disable: The address cannot be used for DNS resolution, regardless of its health check status.
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

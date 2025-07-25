// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressManualAvailableStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmAddressManualAvailableStatusRequest
	GetAcceptLanguage() *string
	SetAddressId(v string) *UpdateCloudGtmAddressManualAvailableStatusRequest
	GetAddressId() *string
	SetAvailableMode(v string) *UpdateCloudGtmAddressManualAvailableStatusRequest
	GetAvailableMode() *string
	SetClientToken(v string) *UpdateCloudGtmAddressManualAvailableStatusRequest
	GetClientToken() *string
	SetManualAvailableStatus(v string) *UpdateCloudGtmAddressManualAvailableStatusRequest
	GetManualAvailableStatus() *string
}

type UpdateCloudGtmAddressManualAvailableStatusRequest struct {
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
	// The ID of the address. This ID uniquely identifies the address.
	//
	// This parameter is required.
	//
	// example:
	//
	// addr-89518218114368**92
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The failover mode that is used when address exceptions are identified. Valid values:
	//
	// 	- auto: the automatic mode. The system determines whether to return an address based on health check results. If the address fails health checks, the system does not return the address. If the address passes health checks, the system returns the address.
	//
	// 	- manual: the manual mode. If an address is in the unavailable state, the address is not returned for DNS requests even if the address passes health checks. If an address is in the available state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// example:
	//
	// manual
	AvailableMode *string `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can specify a custom value for this parameter, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The availability state of the address when AvailableMode is set to manual. Valid values:
	//
	// 	- available: The address is normal. In this state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// 	- unavailable: The address is abnormal. In this state, the address is not returned for DNS requests even if the address passes health checks.
	//
	// example:
	//
	// available
	ManualAvailableStatus *string `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
}

func (s UpdateCloudGtmAddressManualAvailableStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressManualAvailableStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressManualAvailableStatusRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmAddressManualAvailableStatusRequest) GetAddressId() *string {
	return s.AddressId
}

func (s *UpdateCloudGtmAddressManualAvailableStatusRequest) GetAvailableMode() *string {
	return s.AvailableMode
}

func (s *UpdateCloudGtmAddressManualAvailableStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmAddressManualAvailableStatusRequest) GetManualAvailableStatus() *string {
	return s.ManualAvailableStatus
}

func (s *UpdateCloudGtmAddressManualAvailableStatusRequest) SetAcceptLanguage(v string) *UpdateCloudGtmAddressManualAvailableStatusRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmAddressManualAvailableStatusRequest) SetAddressId(v string) *UpdateCloudGtmAddressManualAvailableStatusRequest {
	s.AddressId = &v
	return s
}

func (s *UpdateCloudGtmAddressManualAvailableStatusRequest) SetAvailableMode(v string) *UpdateCloudGtmAddressManualAvailableStatusRequest {
	s.AvailableMode = &v
	return s
}

func (s *UpdateCloudGtmAddressManualAvailableStatusRequest) SetClientToken(v string) *UpdateCloudGtmAddressManualAvailableStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmAddressManualAvailableStatusRequest) SetManualAvailableStatus(v string) *UpdateCloudGtmAddressManualAvailableStatusRequest {
	s.ManualAvailableStatus = &v
	return s
}

func (s *UpdateCloudGtmAddressManualAvailableStatusRequest) Validate() error {
	return dara.Validate(s)
}

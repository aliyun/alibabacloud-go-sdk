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
	// - zh-CN: Chinese
	//
	// - en-US (Default): English
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
	// The failover method for the address. Valid values:
	//
	// - auto: Automatic. The address status is determined by health check results. If a health check fails, DNS resolution stops. If it succeeds, DNS resolution resumes.
	//
	// - manual: Manual. You control the address status. If set to abnormal, DNS resolution stops and does not resume even if health checks succeed. If set to normal, DNS resolution resumes. If a health check fails, an alert is triggered but resolution does not stop.
	//
	// example:
	//
	// manual
	AvailableMode *string `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	// A client-generated token that is used to ensure the idempotence of the request. Make sure that the token is unique among different requests. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The availability status of the address. This parameter takes effect only when AvailableMode is set to manual. Valid values:
	//
	// - available: Normal. The address is available. If a health check fails, an alert is triggered but DNS resolution continues.
	//
	// - unavailable: Abnormal. DNS resolution for the address is stopped. It does not resume even if health checks succeed.
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

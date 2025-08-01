// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceCheckShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayServiceCheckShrinkRequest
	GetAcceptLanguage() *string
	SetCheck(v bool) *UpdateGatewayServiceCheckShrinkRequest
	GetCheck() *bool
	SetExpectedStatusesShrink(v string) *UpdateGatewayServiceCheckShrinkRequest
	GetExpectedStatusesShrink() *string
	SetGatewayUniqueId(v string) *UpdateGatewayServiceCheckShrinkRequest
	GetGatewayUniqueId() *string
	SetHealthyThreshold(v int32) *UpdateGatewayServiceCheckShrinkRequest
	GetHealthyThreshold() *int32
	SetHttpHost(v string) *UpdateGatewayServiceCheckShrinkRequest
	GetHttpHost() *string
	SetHttpPath(v string) *UpdateGatewayServiceCheckShrinkRequest
	GetHttpPath() *string
	SetInterval(v int32) *UpdateGatewayServiceCheckShrinkRequest
	GetInterval() *int32
	SetProtocol(v string) *UpdateGatewayServiceCheckShrinkRequest
	GetProtocol() *string
	SetServiceId(v string) *UpdateGatewayServiceCheckShrinkRequest
	GetServiceId() *string
	SetTimeout(v int32) *UpdateGatewayServiceCheckShrinkRequest
	GetTimeout() *int32
	SetUnhealthyThreshold(v int32) *UpdateGatewayServiceCheckShrinkRequest
	GetUnhealthyThreshold() *int32
}

type UpdateGatewayServiceCheckShrinkRequest struct {
	// The language in which you want to display the results. Valid values: zh and en. zh indicates Chinese, which is the default value. en indicates English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// Specifies whether to enable the health check.
	//
	// example:
	//
	// true
	Check *bool `json:"Check,omitempty" xml:"Check,omitempty"`
	// The expected status code, which is required if the health check protocol is HTTP.
	ExpectedStatusesShrink *string `json:"ExpectedStatuses,omitempty" xml:"ExpectedStatuses,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-8d410698bd7f4628ab76b*****72dd1d
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The healthy threshold of the health check.
	//
	// example:
	//
	// 2
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The health check domain name, which is optional if the health check protocol is HTTP.
	//
	// example:
	//
	// example.com
	HttpHost *string `json:"HttpHost,omitempty" xml:"HttpHost,omitempty"`
	// The health check path, which is required if the health check protocol is HTTP.
	//
	// example:
	//
	// /healthz
	HttpPath *string `json:"HttpPath,omitempty" xml:"HttpPath,omitempty"`
	// The interval at which the health check is performed.
	//
	// example:
	//
	// 2
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The health check protocol. Valid values:
	//
	// 	- HTTP
	//
	// 	- TCP
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 12
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The timeout period of responses to the health check. Unit: seconds.
	//
	// example:
	//
	// 5
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The unhealthy threshold of the health check.
	//
	// example:
	//
	// 2
	UnhealthyThreshold *int32 `json:"UnhealthyThreshold,omitempty" xml:"UnhealthyThreshold,omitempty"`
}

func (s UpdateGatewayServiceCheckShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceCheckShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetCheck() *bool {
	return s.Check
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetExpectedStatusesShrink() *string {
	return s.ExpectedStatusesShrink
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetHttpHost() *string {
	return s.HttpHost
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetHttpPath() *string {
	return s.HttpPath
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateGatewayServiceCheckShrinkRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayServiceCheckShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetCheck(v bool) *UpdateGatewayServiceCheckShrinkRequest {
	s.Check = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetExpectedStatusesShrink(v string) *UpdateGatewayServiceCheckShrinkRequest {
	s.ExpectedStatusesShrink = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayServiceCheckShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetHealthyThreshold(v int32) *UpdateGatewayServiceCheckShrinkRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetHttpHost(v string) *UpdateGatewayServiceCheckShrinkRequest {
	s.HttpHost = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetHttpPath(v string) *UpdateGatewayServiceCheckShrinkRequest {
	s.HttpPath = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetInterval(v int32) *UpdateGatewayServiceCheckShrinkRequest {
	s.Interval = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetProtocol(v string) *UpdateGatewayServiceCheckShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetServiceId(v string) *UpdateGatewayServiceCheckShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetTimeout(v int32) *UpdateGatewayServiceCheckShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) SetUnhealthyThreshold(v int32) *UpdateGatewayServiceCheckShrinkRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *UpdateGatewayServiceCheckShrinkRequest) Validate() error {
	return dara.Validate(s)
}

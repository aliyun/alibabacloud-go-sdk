// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayServiceCheckRequest
	GetAcceptLanguage() *string
	SetCheck(v bool) *UpdateGatewayServiceCheckRequest
	GetCheck() *bool
	SetExpectedStatuses(v []*int32) *UpdateGatewayServiceCheckRequest
	GetExpectedStatuses() []*int32
	SetGatewayUniqueId(v string) *UpdateGatewayServiceCheckRequest
	GetGatewayUniqueId() *string
	SetHealthyThreshold(v int32) *UpdateGatewayServiceCheckRequest
	GetHealthyThreshold() *int32
	SetHttpHost(v string) *UpdateGatewayServiceCheckRequest
	GetHttpHost() *string
	SetHttpPath(v string) *UpdateGatewayServiceCheckRequest
	GetHttpPath() *string
	SetInterval(v int32) *UpdateGatewayServiceCheckRequest
	GetInterval() *int32
	SetProtocol(v string) *UpdateGatewayServiceCheckRequest
	GetProtocol() *string
	SetServiceId(v string) *UpdateGatewayServiceCheckRequest
	GetServiceId() *string
	SetTimeout(v int32) *UpdateGatewayServiceCheckRequest
	GetTimeout() *int32
	SetUnhealthyThreshold(v int32) *UpdateGatewayServiceCheckRequest
	GetUnhealthyThreshold() *int32
}

type UpdateGatewayServiceCheckRequest struct {
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
	ExpectedStatuses []*int32 `json:"ExpectedStatuses,omitempty" xml:"ExpectedStatuses,omitempty" type:"Repeated"`
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

func (s UpdateGatewayServiceCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceCheckRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceCheckRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayServiceCheckRequest) GetCheck() *bool {
	return s.Check
}

func (s *UpdateGatewayServiceCheckRequest) GetExpectedStatuses() []*int32 {
	return s.ExpectedStatuses
}

func (s *UpdateGatewayServiceCheckRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayServiceCheckRequest) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *UpdateGatewayServiceCheckRequest) GetHttpHost() *string {
	return s.HttpHost
}

func (s *UpdateGatewayServiceCheckRequest) GetHttpPath() *string {
	return s.HttpPath
}

func (s *UpdateGatewayServiceCheckRequest) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateGatewayServiceCheckRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateGatewayServiceCheckRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *UpdateGatewayServiceCheckRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateGatewayServiceCheckRequest) GetUnhealthyThreshold() *int32 {
	return s.UnhealthyThreshold
}

func (s *UpdateGatewayServiceCheckRequest) SetAcceptLanguage(v string) *UpdateGatewayServiceCheckRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) SetCheck(v bool) *UpdateGatewayServiceCheckRequest {
	s.Check = &v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) SetExpectedStatuses(v []*int32) *UpdateGatewayServiceCheckRequest {
	s.ExpectedStatuses = v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) SetGatewayUniqueId(v string) *UpdateGatewayServiceCheckRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) SetHealthyThreshold(v int32) *UpdateGatewayServiceCheckRequest {
	s.HealthyThreshold = &v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) SetHttpHost(v string) *UpdateGatewayServiceCheckRequest {
	s.HttpHost = &v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) SetHttpPath(v string) *UpdateGatewayServiceCheckRequest {
	s.HttpPath = &v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) SetInterval(v int32) *UpdateGatewayServiceCheckRequest {
	s.Interval = &v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) SetProtocol(v string) *UpdateGatewayServiceCheckRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) SetServiceId(v string) *UpdateGatewayServiceCheckRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) SetTimeout(v int32) *UpdateGatewayServiceCheckRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) SetUnhealthyThreshold(v int32) *UpdateGatewayServiceCheckRequest {
	s.UnhealthyThreshold = &v
	return s
}

func (s *UpdateGatewayServiceCheckRequest) Validate() error {
	return dara.Validate(s)
}

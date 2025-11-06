// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceTrafficPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayServiceTrafficPolicyRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *UpdateGatewayServiceTrafficPolicyRequest
	GetGatewayId() *int64
	SetGatewayTrafficPolicy(v *TrafficPolicy) *UpdateGatewayServiceTrafficPolicyRequest
	GetGatewayTrafficPolicy() *TrafficPolicy
	SetGatewayUniqueId(v string) *UpdateGatewayServiceTrafficPolicyRequest
	GetGatewayUniqueId() *string
	SetServiceId(v int64) *UpdateGatewayServiceTrafficPolicyRequest
	GetServiceId() *int64
}

type UpdateGatewayServiceTrafficPolicyRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh-CN*	- (default): Chinese
	//
	// 	- **en-US**: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 429
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The traffic policy of the gateway.
	//
	// This parameter is required.
	GatewayTrafficPolicy *TrafficPolicy `json:"GatewayTrafficPolicy,omitempty" xml:"GatewayTrafficPolicy,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-75c5036c083e4f89ba8ef9fafff2e902
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// 411
	ServiceId *int64 `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s UpdateGatewayServiceTrafficPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceTrafficPolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceTrafficPolicyRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayServiceTrafficPolicyRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayServiceTrafficPolicyRequest) GetGatewayTrafficPolicy() *TrafficPolicy {
	return s.GatewayTrafficPolicy
}

func (s *UpdateGatewayServiceTrafficPolicyRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayServiceTrafficPolicyRequest) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *UpdateGatewayServiceTrafficPolicyRequest) SetAcceptLanguage(v string) *UpdateGatewayServiceTrafficPolicyRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyRequest) SetGatewayId(v int64) *UpdateGatewayServiceTrafficPolicyRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyRequest) SetGatewayTrafficPolicy(v *TrafficPolicy) *UpdateGatewayServiceTrafficPolicyRequest {
	s.GatewayTrafficPolicy = v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyRequest) SetGatewayUniqueId(v string) *UpdateGatewayServiceTrafficPolicyRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyRequest) SetServiceId(v int64) *UpdateGatewayServiceTrafficPolicyRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyRequest) Validate() error {
	if s.GatewayTrafficPolicy != nil {
		if err := s.GatewayTrafficPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

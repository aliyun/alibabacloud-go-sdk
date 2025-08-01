// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayServiceTrafficPolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayServiceTrafficPolicyShrinkRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *UpdateGatewayServiceTrafficPolicyShrinkRequest
	GetGatewayId() *int64
	SetGatewayTrafficPolicyShrink(v string) *UpdateGatewayServiceTrafficPolicyShrinkRequest
	GetGatewayTrafficPolicyShrink() *string
	SetGatewayUniqueId(v string) *UpdateGatewayServiceTrafficPolicyShrinkRequest
	GetGatewayUniqueId() *string
	SetServiceId(v int64) *UpdateGatewayServiceTrafficPolicyShrinkRequest
	GetServiceId() *int64
}

type UpdateGatewayServiceTrafficPolicyShrinkRequest struct {
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
	GatewayTrafficPolicyShrink *string `json:"GatewayTrafficPolicy,omitempty" xml:"GatewayTrafficPolicy,omitempty"`
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

func (s UpdateGatewayServiceTrafficPolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayServiceTrafficPolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayServiceTrafficPolicyShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayServiceTrafficPolicyShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayServiceTrafficPolicyShrinkRequest) GetGatewayTrafficPolicyShrink() *string {
	return s.GatewayTrafficPolicyShrink
}

func (s *UpdateGatewayServiceTrafficPolicyShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayServiceTrafficPolicyShrinkRequest) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *UpdateGatewayServiceTrafficPolicyShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayServiceTrafficPolicyShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyShrinkRequest) SetGatewayId(v int64) *UpdateGatewayServiceTrafficPolicyShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyShrinkRequest) SetGatewayTrafficPolicyShrink(v string) *UpdateGatewayServiceTrafficPolicyShrinkRequest {
	s.GatewayTrafficPolicyShrink = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayServiceTrafficPolicyShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyShrinkRequest) SetServiceId(v int64) *UpdateGatewayServiceTrafficPolicyShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *UpdateGatewayServiceTrafficPolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}

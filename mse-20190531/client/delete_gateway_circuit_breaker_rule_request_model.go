// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayCircuitBreakerRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewayCircuitBreakerRuleRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *DeleteGatewayCircuitBreakerRuleRequest
	GetGatewayUniqueId() *string
	SetRouteId(v int64) *DeleteGatewayCircuitBreakerRuleRequest
	GetRouteId() *int64
	SetRuleId(v int64) *DeleteGatewayCircuitBreakerRuleRequest
	GetRuleId() *int64
}

type DeleteGatewayCircuitBreakerRuleRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gw-1cef5440bf2d484db419fb264d4f****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11151
	RouteId *int64 `json:"RouteId,omitempty" xml:"RouteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteGatewayCircuitBreakerRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayCircuitBreakerRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayCircuitBreakerRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewayCircuitBreakerRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayCircuitBreakerRuleRequest) GetRouteId() *int64 {
	return s.RouteId
}

func (s *DeleteGatewayCircuitBreakerRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteGatewayCircuitBreakerRuleRequest) SetAcceptLanguage(v string) *DeleteGatewayCircuitBreakerRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewayCircuitBreakerRuleRequest) SetGatewayUniqueId(v string) *DeleteGatewayCircuitBreakerRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayCircuitBreakerRuleRequest) SetRouteId(v int64) *DeleteGatewayCircuitBreakerRuleRequest {
	s.RouteId = &v
	return s
}

func (s *DeleteGatewayCircuitBreakerRuleRequest) SetRuleId(v int64) *DeleteGatewayCircuitBreakerRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteGatewayCircuitBreakerRuleRequest) Validate() error {
	return dara.Validate(s)
}

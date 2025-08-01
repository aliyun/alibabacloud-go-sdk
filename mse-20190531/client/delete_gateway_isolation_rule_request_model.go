// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayIsolationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewayIsolationRuleRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *DeleteGatewayIsolationRuleRequest
	GetGatewayUniqueId() *string
	SetRouteId(v int64) *DeleteGatewayIsolationRuleRequest
	GetRouteId() *int64
	SetRuleId(v int64) *DeleteGatewayIsolationRuleRequest
	GetRuleId() *int64
}

type DeleteGatewayIsolationRuleRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// gw-e2d226bba4b2445c9e29fa7f8216****
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

func (s DeleteGatewayIsolationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayIsolationRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayIsolationRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewayIsolationRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayIsolationRuleRequest) GetRouteId() *int64 {
	return s.RouteId
}

func (s *DeleteGatewayIsolationRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteGatewayIsolationRuleRequest) SetAcceptLanguage(v string) *DeleteGatewayIsolationRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewayIsolationRuleRequest) SetGatewayUniqueId(v string) *DeleteGatewayIsolationRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayIsolationRuleRequest) SetRouteId(v int64) *DeleteGatewayIsolationRuleRequest {
	s.RouteId = &v
	return s
}

func (s *DeleteGatewayIsolationRuleRequest) SetRuleId(v int64) *DeleteGatewayIsolationRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteGatewayIsolationRuleRequest) Validate() error {
	return dara.Validate(s)
}

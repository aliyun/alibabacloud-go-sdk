// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGatewayFlowRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteGatewayFlowRuleRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *DeleteGatewayFlowRuleRequest
	GetGatewayUniqueId() *string
	SetRouteId(v int64) *DeleteGatewayFlowRuleRequest
	GetRouteId() *int64
	SetRuleId(v int64) *DeleteGatewayFlowRuleRequest
	GetRuleId() *int64
}

type DeleteGatewayFlowRuleRequest struct {
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

func (s DeleteGatewayFlowRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGatewayFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteGatewayFlowRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteGatewayFlowRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *DeleteGatewayFlowRuleRequest) GetRouteId() *int64 {
	return s.RouteId
}

func (s *DeleteGatewayFlowRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DeleteGatewayFlowRuleRequest) SetAcceptLanguage(v string) *DeleteGatewayFlowRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteGatewayFlowRuleRequest) SetGatewayUniqueId(v string) *DeleteGatewayFlowRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *DeleteGatewayFlowRuleRequest) SetRouteId(v int64) *DeleteGatewayFlowRuleRequest {
	s.RouteId = &v
	return s
}

func (s *DeleteGatewayFlowRuleRequest) SetRuleId(v int64) *DeleteGatewayFlowRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteGatewayFlowRuleRequest) Validate() error {
	return dara.Validate(s)
}

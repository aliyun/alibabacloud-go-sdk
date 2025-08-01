// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayCircuitBreakerRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayCircuitBreakerRuleRequest
	GetAcceptLanguage() *string
	SetFilterParams(v string) *ListGatewayCircuitBreakerRuleRequest
	GetFilterParams() *string
}

type ListGatewayCircuitBreakerRuleRequest struct {
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// just for POP
	//
	// This parameter is required.
	//
	// example:
	//
	// param
	FilterParams *string `json:"FilterParams,omitempty" xml:"FilterParams,omitempty"`
}

func (s ListGatewayCircuitBreakerRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayCircuitBreakerRuleRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayCircuitBreakerRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayCircuitBreakerRuleRequest) GetFilterParams() *string {
	return s.FilterParams
}

func (s *ListGatewayCircuitBreakerRuleRequest) SetAcceptLanguage(v string) *ListGatewayCircuitBreakerRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleRequest) SetFilterParams(v string) *ListGatewayCircuitBreakerRuleRequest {
	s.FilterParams = &v
	return s
}

func (s *ListGatewayCircuitBreakerRuleRequest) Validate() error {
	return dara.Validate(s)
}

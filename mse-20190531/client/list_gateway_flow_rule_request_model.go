// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayFlowRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayFlowRuleRequest
	GetAcceptLanguage() *string
	SetFilterParams(v string) *ListGatewayFlowRuleRequest
	GetFilterParams() *string
}

type ListGatewayFlowRuleRequest struct {
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

func (s ListGatewayFlowRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayFlowRuleRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayFlowRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayFlowRuleRequest) GetFilterParams() *string {
	return s.FilterParams
}

func (s *ListGatewayFlowRuleRequest) SetAcceptLanguage(v string) *ListGatewayFlowRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayFlowRuleRequest) SetFilterParams(v string) *ListGatewayFlowRuleRequest {
	s.FilterParams = &v
	return s
}

func (s *ListGatewayFlowRuleRequest) Validate() error {
	return dara.Validate(s)
}

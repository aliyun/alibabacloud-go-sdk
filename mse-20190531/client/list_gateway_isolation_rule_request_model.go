// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayIsolationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListGatewayIsolationRuleRequest
	GetAcceptLanguage() *string
	SetFilterParams(v string) *ListGatewayIsolationRuleRequest
	GetFilterParams() *string
}

type ListGatewayIsolationRuleRequest struct {
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

func (s ListGatewayIsolationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayIsolationRuleRequest) GoString() string {
	return s.String()
}

func (s *ListGatewayIsolationRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListGatewayIsolationRuleRequest) GetFilterParams() *string {
	return s.FilterParams
}

func (s *ListGatewayIsolationRuleRequest) SetAcceptLanguage(v string) *ListGatewayIsolationRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListGatewayIsolationRuleRequest) SetFilterParams(v string) *ListGatewayIsolationRuleRequest {
	s.FilterParams = &v
	return s
}

func (s *ListGatewayIsolationRuleRequest) Validate() error {
	return dara.Validate(s)
}

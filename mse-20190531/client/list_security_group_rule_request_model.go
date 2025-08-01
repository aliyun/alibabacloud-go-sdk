// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityGroupRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *ListSecurityGroupRuleRequest
	GetAcceptLanguage() *string
	SetGatewayUniqueId(v string) *ListSecurityGroupRuleRequest
	GetGatewayUniqueId() *string
}

type ListSecurityGroupRuleRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The unique ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-83b0ddb569434f82b9fe8e4c60c40f7c
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
}

func (s ListSecurityGroupRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *ListSecurityGroupRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *ListSecurityGroupRuleRequest) SetAcceptLanguage(v string) *ListSecurityGroupRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *ListSecurityGroupRuleRequest) SetGatewayUniqueId(v string) *ListSecurityGroupRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *ListSecurityGroupRuleRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSecurityGroupRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddSecurityGroupRuleRequest
	GetAcceptLanguage() *string
	SetDescription(v string) *AddSecurityGroupRuleRequest
	GetDescription() *string
	SetGatewayUniqueId(v string) *AddSecurityGroupRuleRequest
	GetGatewayUniqueId() *string
	SetPortRange(v string) *AddSecurityGroupRuleRequest
	GetPortRange() *string
	SetSecurityGroupId(v string) *AddSecurityGroupRuleRequest
	GetSecurityGroupId() *string
}

type AddSecurityGroupRuleRequest struct {
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
	// The description.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the gateway.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-c9bc5afd61014165bd58f621b491****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The range of port numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1/65535
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-wz929kxhcdpw9z8idqd8
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s AddSecurityGroupRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupRuleRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddSecurityGroupRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *AddSecurityGroupRuleRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddSecurityGroupRuleRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *AddSecurityGroupRuleRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AddSecurityGroupRuleRequest) SetAcceptLanguage(v string) *AddSecurityGroupRuleRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddSecurityGroupRuleRequest) SetDescription(v string) *AddSecurityGroupRuleRequest {
	s.Description = &v
	return s
}

func (s *AddSecurityGroupRuleRequest) SetGatewayUniqueId(v string) *AddSecurityGroupRuleRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddSecurityGroupRuleRequest) SetPortRange(v string) *AddSecurityGroupRuleRequest {
	s.PortRange = &v
	return s
}

func (s *AddSecurityGroupRuleRequest) SetSecurityGroupId(v string) *AddSecurityGroupRuleRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AddSecurityGroupRuleRequest) Validate() error {
	return dara.Validate(s)
}

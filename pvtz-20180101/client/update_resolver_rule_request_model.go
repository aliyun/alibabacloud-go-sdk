// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResolverRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *UpdateResolverRuleRequest
	GetEndpointId() *string
	SetForwardIp(v []*UpdateResolverRuleRequestForwardIp) *UpdateResolverRuleRequest
	GetForwardIp() []*UpdateResolverRuleRequestForwardIp
	SetLang(v string) *UpdateResolverRuleRequest
	GetLang() *string
	SetName(v string) *UpdateResolverRuleRequest
	GetName() *string
	SetRuleId(v string) *UpdateResolverRuleRequest
	GetRuleId() *string
}

type UpdateResolverRuleRequest struct {
	// The endpoint ID.
	//
	// example:
	//
	// hr****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The IP addresses and ports of the external Domain Name System (DNS) servers. Enter the IP addresses and ports of the destination servers to which the DNS requests are forwarded. You can enter up to six IP addresses and ports. Both private and public IP addresses are supported.
	//
	// >  If you specify public IP addresses as the IP addresses of the external DNS servers and Elastic Compute Service (ECS) instances in the outbound virtual private cloud (VPC) are not assigned public IP addresses, you need to activate NAT Gateway for the VPC and create and manage SNAT entries on a NAT gateway.
	ForwardIp []*UpdateResolverRuleRequestForwardIp `json:"ForwardIp,omitempty" xml:"ForwardIp,omitempty" type:"Repeated"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the forwarding rule.
	//
	// example:
	//
	// forward rule-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the forwarding rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// hr****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateResolverRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResolverRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateResolverRuleRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *UpdateResolverRuleRequest) GetForwardIp() []*UpdateResolverRuleRequestForwardIp {
	return s.ForwardIp
}

func (s *UpdateResolverRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateResolverRuleRequest) GetName() *string {
	return s.Name
}

func (s *UpdateResolverRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *UpdateResolverRuleRequest) SetEndpointId(v string) *UpdateResolverRuleRequest {
	s.EndpointId = &v
	return s
}

func (s *UpdateResolverRuleRequest) SetForwardIp(v []*UpdateResolverRuleRequestForwardIp) *UpdateResolverRuleRequest {
	s.ForwardIp = v
	return s
}

func (s *UpdateResolverRuleRequest) SetLang(v string) *UpdateResolverRuleRequest {
	s.Lang = &v
	return s
}

func (s *UpdateResolverRuleRequest) SetName(v string) *UpdateResolverRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateResolverRuleRequest) SetRuleId(v string) *UpdateResolverRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateResolverRuleRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateResolverRuleRequestForwardIp struct {
	// The IP address of the destination server.
	//
	// >  You cannot specify the following IP addresses as the IP addresses of the external DNS servers because the IP addresses are reserved by the system: 100.100.2.136 to 100.100.2.138, and 100.100.2.116 to 100.100.2.118.
	//
	// example:
	//
	// 172.16.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port of the destination server.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s UpdateResolverRuleRequestForwardIp) String() string {
	return dara.Prettify(s)
}

func (s UpdateResolverRuleRequestForwardIp) GoString() string {
	return s.String()
}

func (s *UpdateResolverRuleRequestForwardIp) GetIp() *string {
	return s.Ip
}

func (s *UpdateResolverRuleRequestForwardIp) GetPort() *int32 {
	return s.Port
}

func (s *UpdateResolverRuleRequestForwardIp) SetIp(v string) *UpdateResolverRuleRequestForwardIp {
	s.Ip = &v
	return s
}

func (s *UpdateResolverRuleRequestForwardIp) SetPort(v int32) *UpdateResolverRuleRequestForwardIp {
	s.Port = &v
	return s
}

func (s *UpdateResolverRuleRequestForwardIp) Validate() error {
	return dara.Validate(s)
}

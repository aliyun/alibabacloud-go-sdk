// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkRuleAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkRules(v string) *DescribeNetworkRuleAttributesRequest
	GetNetworkRules() *string
}

type DescribeNetworkRuleAttributesRequest struct {
	// The details of the port forwarding rule. This parameter is a JSON string. The string contains the following fields:
	//
	// 	- **InstanceId**: the ID of the instance. This field is required and must be of the STRING type.
	//
	// 	- **Protocol**: the forwarding protocol. This field is required and must be of the STRING type. Valid values: **tcp*	- and **udp**.
	//
	// 	- **FrontendPort**: the forwarding port. This field is required and must be of the INTEGER type.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"ddoscoo-cn-mp91j1ao****","Protocol":"tcp","FrontendPort":8080}]
	NetworkRules *string `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty"`
}

func (s DescribeNetworkRuleAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkRuleAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkRuleAttributesRequest) GetNetworkRules() *string {
	return s.NetworkRules
}

func (s *DescribeNetworkRuleAttributesRequest) SetNetworkRules(v string) *DescribeNetworkRuleAttributesRequest {
	s.NetworkRules = &v
	return s
}

func (s *DescribeNetworkRuleAttributesRequest) Validate() error {
	return dara.Validate(s)
}

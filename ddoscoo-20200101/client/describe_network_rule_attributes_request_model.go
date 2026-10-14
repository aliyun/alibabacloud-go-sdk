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
	// The port forwarding rules to query. This parameter is a JSON string with the following structure:
	//
	// - **InstanceId**: String type. Required. The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// - **Protocol**: String type. Required. The forwarding protocol type. Valid values: **tcp**, **udp**.
	//
	// - **FrontendPort**: Integer type. Required. The forwarding port.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkRule(v string) *DeleteNetworkRuleRequest
	GetNetworkRule() *string
}

type DeleteNetworkRuleRequest struct {
	// An array that consists of the information about the port forwarding rule. This parameter is a JSON string. The string contains the following fields:
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
	NetworkRule *string `json:"NetworkRule,omitempty" xml:"NetworkRule,omitempty"`
}

func (s DeleteNetworkRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteNetworkRuleRequest) GetNetworkRule() *string {
	return s.NetworkRule
}

func (s *DeleteNetworkRuleRequest) SetNetworkRule(v string) *DeleteNetworkRuleRequest {
	s.NetworkRule = &v
	return s
}

func (s *DeleteNetworkRuleRequest) Validate() error {
	return dara.Validate(s)
}

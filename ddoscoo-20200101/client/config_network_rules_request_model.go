// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigNetworkRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkRules(v string) *ConfigNetworkRulesRequest
	GetNetworkRules() *string
}

type ConfigNetworkRulesRequest struct {
	// The details of the port forwarding rule. This parameter is a JSON string. The string contains the following fields:
	//
	// 	- **InstanceId**: the ID of the instance. This field is required and must be of the STRING type.
	//
	// 	- **Protocol**: the forwarding protocol. This field is required and must be of the STRING type. Valid values: **tcp*	- and **udp**.
	//
	// 	- **FrontendPort**: the forwarding port. This field is required and must be of the INTEGER type.
	//
	// 	- **BackendPort**: the port of the origin server. This field is required and must be of the INTEGER type.
	//
	// 	- **RealServers**: the IP addresses of the origin server. This field is required and must be a JSON array. You can specify up to 20 IP addresses.
	//
	// > You can modify only the value of **RealServers*	- when you modify a port forwarding rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"ddoscoo-cn-mp91j1ao****","Protocol":"tcp","FrontendPort":8080,"BackendPort":8080,"RealServers":["1.1.1.1","2.2.2.2","3.3.3.3"]}]
	NetworkRules *string `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty"`
}

func (s ConfigNetworkRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigNetworkRulesRequest) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRulesRequest) GetNetworkRules() *string {
	return s.NetworkRules
}

func (s *ConfigNetworkRulesRequest) SetNetworkRules(v string) *ConfigNetworkRulesRequest {
	s.NetworkRules = &v
	return s
}

func (s *ConfigNetworkRulesRequest) Validate() error {
	return dara.Validate(s)
}

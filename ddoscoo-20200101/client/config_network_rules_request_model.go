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
	// The details of the port forwarding rules. The value is a JSON format character string that contains the following structure:
	//
	// - **InstanceId**: String. Required. The instance ID of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// - **Protocol**: String. Required. The forwarding protocol type. Valid values: **tcp**, **udp**.
	//
	// - **FrontendPort**: Integer. Required. The Redirection Port.
	//
	// - **BackendPort**: Integer. Required. The origin server port.
	//
	// - **RealServers**: JSON array. Required. The list of origin server IP addresses. A maximum of 20 IP addresses are supported.
	//
	// > When you edit port forwarding rules, you can modify only **RealServers**, which specifies the origin server IP addresses.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RulePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListeners(v string) *ConfigLayer4RulePolicyRequest
	GetListeners() *string
}

type ConfigLayer4RulePolicyRequest struct {
	// The port forwarding rules.
	//
	// This parameter is a string that represents a JSON array converted from a character string. Each element in the JSON array represents a created port forwarding rule. You can specify up to one port forwarding rule at a time.
	//
	// > You can invoke [DescribeNetworkRules](https://help.aliyun.com/document_detail/157484.html) to query all created port forwarding rules.
	//
	// Each port forwarding rule contains the following fields:
	//
	// - **InstanceId**: string. Required. The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// - **Protocol**: string. Required. The forwarding protocol type. Valid values: **tcp**, **udp**.
	//
	// - **FrontendPort**: integer. Required. The Redirection Port.
	//
	// - **BackendPort**: integer. Required. The origin server port.
	//
	// - **PriRealServers**: JSON array. Required. The list of primary origin server IP addresses. Each element in the JSON array represents a primary origin server IP address. You can specify up to 20 primary origin server IP addresses.
	//
	//     Each primary origin server IP address contains the following fields:
	//
	//     - **RealServer**: string. Required. The primary origin server IP address.
	//
	// - **SecRealServers**: JSON array. Required. The list of secondary origin server IP addresses. Each element in the JSON array represents a secondary origin server IP address. You can specify up to 20 secondary origin server IP addresses.
	//
	//     Each secondary origin server IP address contains the following fields:
	//
	//
	//
	//     - **RealServer**: string. Required. The secondary origin server IP address.
	//
	//
	//
	// - **CurrentRsIndex**: integer. Required. The type of origin server to activate. Valid values:
	//
	//     - **1**: The primary origin server settings take effect. Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the primary origin server IP addresses.
	//
	//     - **2**: The secondary origin server settings take effect. Anti-DDoS Pro or Anti-DDoS Premium forwards service traffic to the secondary origin server IP addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{\\"InstanceId\\":\\"ddosDip-sg-4hr2b3l****\\",\\"Protocol\\":\\"udp\\",\\"FrontendPort\\":2020,\\"BackendPort\\":2022,\\"PriRealServers\\":[{\\"RealServer\\":\\"192.0.2.1\\"},{\\"RealServer\\":\\"192.0.2.2\\"}],\\"SecRealServers\\":[{\\"RealServer\\":\\"192.0.2.3\\"},{\\"RealServer\\":\\"192.0.2.4\\"}],\\"CurrentRsIndex\\":1}]
	Listeners *string `json:"Listeners,omitempty" xml:"Listeners,omitempty"`
}

func (s ConfigLayer4RulePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RulePolicyRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RulePolicyRequest) GetListeners() *string {
	return s.Listeners
}

func (s *ConfigLayer4RulePolicyRequest) SetListeners(v string) *ConfigLayer4RulePolicyRequest {
	s.Listeners = &v
	return s
}

func (s *ConfigLayer4RulePolicyRequest) Validate() error {
	return dara.Validate(s)
}

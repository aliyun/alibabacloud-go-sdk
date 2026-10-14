// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpMode(v string) *CreateNetworkRulesRequest
	GetIpMode() *string
	SetNetworkRules(v string) *CreateNetworkRulesRequest
	GetNetworkRules() *string
}

type CreateNetworkRulesRequest struct {
	IpMode *string `json:"IpMode,omitempty" xml:"IpMode,omitempty"`
	// The details of the port forwarding rules, expressed as a character string in JSON format with the following structure:
	//
	// - **InstanceId**: String. Required. The Anti-DDoS Pro or Anti-DDoS Premium instance ID.
	//
	// - **Protocol**: String. Required. The forwarding protocol type. Valid values: **tcp**, **udp**.
	//
	// - **FrontendPort**: Integer. Required. The Redirection Port.
	//
	// - **BackendPort**: Integer. Required. The origin server port.
	//
	// - **RealServers**: JSON array. Required. The list of origin server IP addresses. A maximum of 20 IP addresses are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"InstanceId":"ddoscoo-cn-mp91j1ao****","Protocol":"tcp","FrontendPort":8080,"BackendPort":8080,"RealServers":["1.1.1.1","2.2.2.2"]}]
	NetworkRules *string `json:"NetworkRules,omitempty" xml:"NetworkRules,omitempty"`
}

func (s CreateNetworkRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkRulesRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkRulesRequest) GetIpMode() *string {
	return s.IpMode
}

func (s *CreateNetworkRulesRequest) GetNetworkRules() *string {
	return s.NetworkRules
}

func (s *CreateNetworkRulesRequest) SetIpMode(v string) *CreateNetworkRulesRequest {
	s.IpMode = &v
	return s
}

func (s *CreateNetworkRulesRequest) SetNetworkRules(v string) *CreateNetworkRulesRequest {
	s.NetworkRules = &v
	return s
}

func (s *CreateNetworkRulesRequest) Validate() error {
	return dara.Validate(s)
}

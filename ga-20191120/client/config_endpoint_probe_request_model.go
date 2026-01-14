// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigEndpointProbeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ConfigEndpointProbeRequest
	GetClientToken() *string
	SetEnable(v string) *ConfigEndpointProbeRequest
	GetEnable() *string
	SetEndpoint(v string) *ConfigEndpointProbeRequest
	GetEndpoint() *string
	SetEndpointGroupId(v string) *ConfigEndpointProbeRequest
	GetEndpointGroupId() *string
	SetEndpointType(v string) *ConfigEndpointProbeRequest
	GetEndpointType() *string
	SetProbePort(v string) *ConfigEndpointProbeRequest
	GetProbePort() *string
	SetProbeProtocol(v string) *ConfigEndpointProbeRequest
	GetProbeProtocol() *string
	SetRegionId(v string) *ConfigEndpointProbeRequest
	GetRegionId() *string
}

type ConfigEndpointProbeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable latency monitoring. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// 127.0.XX.XX
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The endpoint group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	EndpointGroupId *string `json:"EndpointGroupId,omitempty" xml:"EndpointGroupId,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Ip:*	- a custom IP address.
	//
	// 	- **Domain:*	- a custom domain name.
	//
	// 	- **EIP:*	- an Alibaba Cloud elastic IP address (EIP).
	//
	// 	- **PublicIp:*	- an Alibaba Cloud public IP address.
	//
	// This parameter is required.
	//
	// example:
	//
	// Ip
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The port that is used to monitor latency. Valid values: **0*	- to **65535**.
	//
	// example:
	//
	// 80
	ProbePort *string `json:"ProbePort,omitempty" xml:"ProbePort,omitempty"`
	// The protocol that is used to monitor latency. Valid values:
	//
	// 	- **tcp:*	- TCP.
	//
	// 	- **icmp:*	- ICMP.
	//
	// example:
	//
	// tcp
	ProbeProtocol *string `json:"ProbeProtocol,omitempty" xml:"ProbeProtocol,omitempty"`
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigEndpointProbeRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigEndpointProbeRequest) GoString() string {
	return s.String()
}

func (s *ConfigEndpointProbeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ConfigEndpointProbeRequest) GetEnable() *string {
	return s.Enable
}

func (s *ConfigEndpointProbeRequest) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ConfigEndpointProbeRequest) GetEndpointGroupId() *string {
	return s.EndpointGroupId
}

func (s *ConfigEndpointProbeRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *ConfigEndpointProbeRequest) GetProbePort() *string {
	return s.ProbePort
}

func (s *ConfigEndpointProbeRequest) GetProbeProtocol() *string {
	return s.ProbeProtocol
}

func (s *ConfigEndpointProbeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigEndpointProbeRequest) SetClientToken(v string) *ConfigEndpointProbeRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetEnable(v string) *ConfigEndpointProbeRequest {
	s.Enable = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetEndpoint(v string) *ConfigEndpointProbeRequest {
	s.Endpoint = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetEndpointGroupId(v string) *ConfigEndpointProbeRequest {
	s.EndpointGroupId = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetEndpointType(v string) *ConfigEndpointProbeRequest {
	s.EndpointType = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetProbePort(v string) *ConfigEndpointProbeRequest {
	s.ProbePort = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetProbeProtocol(v string) *ConfigEndpointProbeRequest {
	s.ProbeProtocol = &v
	return s
}

func (s *ConfigEndpointProbeRequest) SetRegionId(v string) *ConfigEndpointProbeRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigEndpointProbeRequest) Validate() error {
	return dara.Validate(s)
}

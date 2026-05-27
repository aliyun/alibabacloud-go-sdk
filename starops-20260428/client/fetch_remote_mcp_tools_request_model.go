// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFetchRemoteMcpToolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnection(v *FetchRemoteMcpToolsRequestConnection) *FetchRemoteMcpToolsRequest
	GetConnection() *FetchRemoteMcpToolsRequestConnection
	SetNetwork(v *FetchRemoteMcpToolsRequestNetwork) *FetchRemoteMcpToolsRequest
	GetNetwork() *FetchRemoteMcpToolsRequestNetwork
}

type FetchRemoteMcpToolsRequest struct {
	// This parameter is required.
	Connection *FetchRemoteMcpToolsRequestConnection `json:"connection,omitempty" xml:"connection,omitempty" type:"Struct"`
	Network    *FetchRemoteMcpToolsRequestNetwork    `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
}

func (s FetchRemoteMcpToolsRequest) String() string {
	return dara.Prettify(s)
}

func (s FetchRemoteMcpToolsRequest) GoString() string {
	return s.String()
}

func (s *FetchRemoteMcpToolsRequest) GetConnection() *FetchRemoteMcpToolsRequestConnection {
	return s.Connection
}

func (s *FetchRemoteMcpToolsRequest) GetNetwork() *FetchRemoteMcpToolsRequestNetwork {
	return s.Network
}

func (s *FetchRemoteMcpToolsRequest) SetConnection(v *FetchRemoteMcpToolsRequestConnection) *FetchRemoteMcpToolsRequest {
	s.Connection = v
	return s
}

func (s *FetchRemoteMcpToolsRequest) SetNetwork(v *FetchRemoteMcpToolsRequestNetwork) *FetchRemoteMcpToolsRequest {
	s.Network = v
	return s
}

func (s *FetchRemoteMcpToolsRequest) Validate() error {
	if s.Connection != nil {
		if err := s.Connection.Validate(); err != nil {
			return err
		}
	}
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FetchRemoteMcpToolsRequestConnection struct {
	Auth *FetchRemoteMcpToolsRequestConnectionAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/mcp
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// Custom
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// example:
	//
	// 5000
	Timeout *int64 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http
	Transport *string `json:"transport,omitempty" xml:"transport,omitempty"`
}

func (s FetchRemoteMcpToolsRequestConnection) String() string {
	return dara.Prettify(s)
}

func (s FetchRemoteMcpToolsRequestConnection) GoString() string {
	return s.String()
}

func (s *FetchRemoteMcpToolsRequestConnection) GetAuth() *FetchRemoteMcpToolsRequestConnectionAuth {
	return s.Auth
}

func (s *FetchRemoteMcpToolsRequestConnection) GetEndpoint() *string {
	return s.Endpoint
}

func (s *FetchRemoteMcpToolsRequestConnection) GetPlatform() *string {
	return s.Platform
}

func (s *FetchRemoteMcpToolsRequestConnection) GetTimeout() *int64 {
	return s.Timeout
}

func (s *FetchRemoteMcpToolsRequestConnection) GetTransport() *string {
	return s.Transport
}

func (s *FetchRemoteMcpToolsRequestConnection) SetAuth(v *FetchRemoteMcpToolsRequestConnectionAuth) *FetchRemoteMcpToolsRequestConnection {
	s.Auth = v
	return s
}

func (s *FetchRemoteMcpToolsRequestConnection) SetEndpoint(v string) *FetchRemoteMcpToolsRequestConnection {
	s.Endpoint = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestConnection) SetPlatform(v string) *FetchRemoteMcpToolsRequestConnection {
	s.Platform = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestConnection) SetTimeout(v int64) *FetchRemoteMcpToolsRequestConnection {
	s.Timeout = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestConnection) SetTransport(v string) *FetchRemoteMcpToolsRequestConnection {
	s.Transport = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestConnection) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FetchRemoteMcpToolsRequestConnectionAuth struct {
	// example:
	//
	// {"token":"example-token"}
	KeyInfo map[string]*string `json:"keyInfo,omitempty" xml:"keyInfo,omitempty"`
	// example:
	//
	// bearer
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FetchRemoteMcpToolsRequestConnectionAuth) String() string {
	return dara.Prettify(s)
}

func (s FetchRemoteMcpToolsRequestConnectionAuth) GoString() string {
	return s.String()
}

func (s *FetchRemoteMcpToolsRequestConnectionAuth) GetKeyInfo() map[string]*string {
	return s.KeyInfo
}

func (s *FetchRemoteMcpToolsRequestConnectionAuth) GetType() *string {
	return s.Type
}

func (s *FetchRemoteMcpToolsRequestConnectionAuth) SetKeyInfo(v map[string]*string) *FetchRemoteMcpToolsRequestConnectionAuth {
	s.KeyInfo = v
	return s
}

func (s *FetchRemoteMcpToolsRequestConnectionAuth) SetType(v string) *FetchRemoteMcpToolsRequestConnectionAuth {
	s.Type = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestConnectionAuth) Validate() error {
	return dara.Validate(s)
}

type FetchRemoteMcpToolsRequestNetwork struct {
	// example:
	//
	// 10.0.0.12
	AccessIp *string `json:"accessIp,omitempty" xml:"accessIp,omitempty"`
	// example:
	//
	// 8080
	AccessPort *int64 `json:"accessPort,omitempty" xml:"accessPort,omitempty"`
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// mcp-xxx
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// example:
	//
	// public
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// sg-xxx
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// example:
	//
	// vsw-xxx
	VswId *string `json:"vswId,omitempty" xml:"vswId,omitempty"`
}

func (s FetchRemoteMcpToolsRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s FetchRemoteMcpToolsRequestNetwork) GoString() string {
	return s.String()
}

func (s *FetchRemoteMcpToolsRequestNetwork) GetAccessIp() *string {
	return s.AccessIp
}

func (s *FetchRemoteMcpToolsRequestNetwork) GetAccessPort() *int64 {
	return s.AccessPort
}

func (s *FetchRemoteMcpToolsRequestNetwork) GetGatewayId() *string {
	return s.GatewayId
}

func (s *FetchRemoteMcpToolsRequestNetwork) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *FetchRemoteMcpToolsRequestNetwork) GetMode() *string {
	return s.Mode
}

func (s *FetchRemoteMcpToolsRequestNetwork) GetRegion() *string {
	return s.Region
}

func (s *FetchRemoteMcpToolsRequestNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *FetchRemoteMcpToolsRequestNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *FetchRemoteMcpToolsRequestNetwork) GetVswId() *string {
	return s.VswId
}

func (s *FetchRemoteMcpToolsRequestNetwork) SetAccessIp(v string) *FetchRemoteMcpToolsRequestNetwork {
	s.AccessIp = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestNetwork) SetAccessPort(v int64) *FetchRemoteMcpToolsRequestNetwork {
	s.AccessPort = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestNetwork) SetGatewayId(v string) *FetchRemoteMcpToolsRequestNetwork {
	s.GatewayId = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestNetwork) SetMcpServerId(v string) *FetchRemoteMcpToolsRequestNetwork {
	s.McpServerId = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestNetwork) SetMode(v string) *FetchRemoteMcpToolsRequestNetwork {
	s.Mode = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestNetwork) SetRegion(v string) *FetchRemoteMcpToolsRequestNetwork {
	s.Region = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestNetwork) SetSecurityGroupId(v string) *FetchRemoteMcpToolsRequestNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestNetwork) SetVpcId(v string) *FetchRemoteMcpToolsRequestNetwork {
	s.VpcId = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestNetwork) SetVswId(v string) *FetchRemoteMcpToolsRequestNetwork {
	s.VswId = &v
	return s
}

func (s *FetchRemoteMcpToolsRequestNetwork) Validate() error {
	return dara.Validate(s)
}

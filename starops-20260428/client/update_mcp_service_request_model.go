// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnection(v *UpdateMcpServiceRequestConnection) *UpdateMcpServiceRequest
	GetConnection() *UpdateMcpServiceRequestConnection
	SetDescription(v string) *UpdateMcpServiceRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateMcpServiceRequest
	GetDisplayName() *string
	SetEnable(v bool) *UpdateMcpServiceRequest
	GetEnable() *bool
	SetNetwork(v *UpdateMcpServiceRequestNetwork) *UpdateMcpServiceRequest
	GetNetwork() *UpdateMcpServiceRequestNetwork
	SetTools(v []*UpdateMcpServiceRequestTools) *UpdateMcpServiceRequest
	GetTools() []*UpdateMcpServiceRequestTools
}

type UpdateMcpServiceRequest struct {
	// The request body parameters.
	//
	// This parameter is required.
	Connection *UpdateMcpServiceRequestConnection `json:"connection,omitempty" xml:"connection,omitempty" type:"Struct"`
	// The description of the MCP service.
	//
	// example:
	//
	// 通过 MCP 调用日志查询工具。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the MCP service.
	//
	// example:
	//
	// 日志查询
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Specifies whether to enable the MCP service.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The request body parameters.
	//
	// This parameter is required.
	Network *UpdateMcpServiceRequestNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The list of MCP tools.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"name":"query_logs"}]
	Tools []*UpdateMcpServiceRequestTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s UpdateMcpServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcpServiceRequest) GetConnection() *UpdateMcpServiceRequestConnection {
	return s.Connection
}

func (s *UpdateMcpServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpServiceRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateMcpServiceRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateMcpServiceRequest) GetNetwork() *UpdateMcpServiceRequestNetwork {
	return s.Network
}

func (s *UpdateMcpServiceRequest) GetTools() []*UpdateMcpServiceRequestTools {
	return s.Tools
}

func (s *UpdateMcpServiceRequest) SetConnection(v *UpdateMcpServiceRequestConnection) *UpdateMcpServiceRequest {
	s.Connection = v
	return s
}

func (s *UpdateMcpServiceRequest) SetDescription(v string) *UpdateMcpServiceRequest {
	s.Description = &v
	return s
}

func (s *UpdateMcpServiceRequest) SetDisplayName(v string) *UpdateMcpServiceRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateMcpServiceRequest) SetEnable(v bool) *UpdateMcpServiceRequest {
	s.Enable = &v
	return s
}

func (s *UpdateMcpServiceRequest) SetNetwork(v *UpdateMcpServiceRequestNetwork) *UpdateMcpServiceRequest {
	s.Network = v
	return s
}

func (s *UpdateMcpServiceRequest) SetTools(v []*UpdateMcpServiceRequestTools) *UpdateMcpServiceRequest {
	s.Tools = v
	return s
}

func (s *UpdateMcpServiceRequest) Validate() error {
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
	if s.Tools != nil {
		for _, item := range s.Tools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateMcpServiceRequestConnection struct {
	// The request body parameters.
	Auth *UpdateMcpServiceRequestConnectionAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The access endpoint of the MCP service.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/mcp
	Endpoint *string            `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Headers  map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The platform type of the MCP service. Valid values: AIGateway and Custom.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The timeout period for requests to the MCP service. Unit: milliseconds.
	//
	// example:
	//
	// 5000
	Timeout *int64 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// The transport protocol of the MCP service. Valid values: http and sse.
	//
	// This parameter is required.
	//
	// example:
	//
	// http
	Transport *string `json:"transport,omitempty" xml:"transport,omitempty"`
}

func (s UpdateMcpServiceRequestConnection) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServiceRequestConnection) GoString() string {
	return s.String()
}

func (s *UpdateMcpServiceRequestConnection) GetAuth() *UpdateMcpServiceRequestConnectionAuth {
	return s.Auth
}

func (s *UpdateMcpServiceRequestConnection) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateMcpServiceRequestConnection) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMcpServiceRequestConnection) GetPlatform() *string {
	return s.Platform
}

func (s *UpdateMcpServiceRequestConnection) GetTimeout() *int64 {
	return s.Timeout
}

func (s *UpdateMcpServiceRequestConnection) GetTransport() *string {
	return s.Transport
}

func (s *UpdateMcpServiceRequestConnection) SetAuth(v *UpdateMcpServiceRequestConnectionAuth) *UpdateMcpServiceRequestConnection {
	s.Auth = v
	return s
}

func (s *UpdateMcpServiceRequestConnection) SetEndpoint(v string) *UpdateMcpServiceRequestConnection {
	s.Endpoint = &v
	return s
}

func (s *UpdateMcpServiceRequestConnection) SetHeaders(v map[string]*string) *UpdateMcpServiceRequestConnection {
	s.Headers = v
	return s
}

func (s *UpdateMcpServiceRequestConnection) SetPlatform(v string) *UpdateMcpServiceRequestConnection {
	s.Platform = &v
	return s
}

func (s *UpdateMcpServiceRequestConnection) SetTimeout(v int64) *UpdateMcpServiceRequestConnection {
	s.Timeout = &v
	return s
}

func (s *UpdateMcpServiceRequestConnection) SetTransport(v string) *UpdateMcpServiceRequestConnection {
	s.Transport = &v
	return s
}

func (s *UpdateMcpServiceRequestConnection) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcpServiceRequestConnectionAuth struct {
	// The request body parameters.
	//
	// example:
	//
	// {"token":"example-token"}
	KeyInfo map[string]*string `json:"keyInfo,omitempty" xml:"keyInfo,omitempty"`
	// The authentication type. Currently, only bearer is supported.
	//
	// example:
	//
	// bearer
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateMcpServiceRequestConnectionAuth) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServiceRequestConnectionAuth) GoString() string {
	return s.String()
}

func (s *UpdateMcpServiceRequestConnectionAuth) GetKeyInfo() map[string]*string {
	return s.KeyInfo
}

func (s *UpdateMcpServiceRequestConnectionAuth) GetType() *string {
	return s.Type
}

func (s *UpdateMcpServiceRequestConnectionAuth) SetKeyInfo(v map[string]*string) *UpdateMcpServiceRequestConnectionAuth {
	s.KeyInfo = v
	return s
}

func (s *UpdateMcpServiceRequestConnectionAuth) SetType(v string) *UpdateMcpServiceRequestConnectionAuth {
	s.Type = &v
	return s
}

func (s *UpdateMcpServiceRequestConnectionAuth) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpServiceRequestNetwork struct {
	// The IP address used to access the MCP service over the VPC network.
	//
	// example:
	//
	// 10.0.0.12
	AccessIp *string `json:"accessIp,omitempty" xml:"accessIp,omitempty"`
	// The port used to access the MCP service over the VPC network. Valid values: 1 to 65535.
	//
	// example:
	//
	// 8080
	AccessPort *int64 `json:"accessPort,omitempty" xml:"accessPort,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// gw-xxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The MCP Server instance ID.
	//
	// example:
	//
	// mcp-xxx
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// The network access mode of the MCP service. Valid values: public and vpc.
	//
	// This parameter is required.
	//
	// example:
	//
	// public
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// The region where the VPC network resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-xxx
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxx
	VswId *string `json:"vswId,omitempty" xml:"vswId,omitempty"`
}

func (s UpdateMcpServiceRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServiceRequestNetwork) GoString() string {
	return s.String()
}

func (s *UpdateMcpServiceRequestNetwork) GetAccessIp() *string {
	return s.AccessIp
}

func (s *UpdateMcpServiceRequestNetwork) GetAccessPort() *int64 {
	return s.AccessPort
}

func (s *UpdateMcpServiceRequestNetwork) GetGatewayId() *string {
	return s.GatewayId
}

func (s *UpdateMcpServiceRequestNetwork) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *UpdateMcpServiceRequestNetwork) GetMode() *string {
	return s.Mode
}

func (s *UpdateMcpServiceRequestNetwork) GetRegion() *string {
	return s.Region
}

func (s *UpdateMcpServiceRequestNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateMcpServiceRequestNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateMcpServiceRequestNetwork) GetVswId() *string {
	return s.VswId
}

func (s *UpdateMcpServiceRequestNetwork) SetAccessIp(v string) *UpdateMcpServiceRequestNetwork {
	s.AccessIp = &v
	return s
}

func (s *UpdateMcpServiceRequestNetwork) SetAccessPort(v int64) *UpdateMcpServiceRequestNetwork {
	s.AccessPort = &v
	return s
}

func (s *UpdateMcpServiceRequestNetwork) SetGatewayId(v string) *UpdateMcpServiceRequestNetwork {
	s.GatewayId = &v
	return s
}

func (s *UpdateMcpServiceRequestNetwork) SetMcpServerId(v string) *UpdateMcpServiceRequestNetwork {
	s.McpServerId = &v
	return s
}

func (s *UpdateMcpServiceRequestNetwork) SetMode(v string) *UpdateMcpServiceRequestNetwork {
	s.Mode = &v
	return s
}

func (s *UpdateMcpServiceRequestNetwork) SetRegion(v string) *UpdateMcpServiceRequestNetwork {
	s.Region = &v
	return s
}

func (s *UpdateMcpServiceRequestNetwork) SetSecurityGroupId(v string) *UpdateMcpServiceRequestNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateMcpServiceRequestNetwork) SetVpcId(v string) *UpdateMcpServiceRequestNetwork {
	s.VpcId = &v
	return s
}

func (s *UpdateMcpServiceRequestNetwork) SetVswId(v string) *UpdateMcpServiceRequestNetwork {
	s.VswId = &v
	return s
}

func (s *UpdateMcpServiceRequestNetwork) Validate() error {
	return dara.Validate(s)
}

type UpdateMcpServiceRequestTools struct {
	// The request body parameters.
	//
	// example:
	//
	// {}
	Annotations map[string]interface{} `json:"annotations,omitempty" xml:"annotations,omitempty"`
	// Specifies whether user confirmation is required before calling the MCP tool.
	//
	// example:
	//
	// false
	Confirm *bool `json:"confirm,omitempty" xml:"confirm,omitempty"`
	// The description of the MCP tool.
	//
	// example:
	//
	// 查询指定日志库中的日志。
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name of the MCP tool.
	//
	// example:
	//
	// 日志查询工具
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// Specifies whether to enable the MCP tool.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The request body parameters.
	//
	// example:
	//
	// {}
	Execution map[string]interface{} `json:"execution,omitempty" xml:"execution,omitempty"`
	// The list of MCP tool icons.
	//
	// example:
	//
	// []
	Icons []map[string]interface{} `json:"icons,omitempty" xml:"icons,omitempty" type:"Repeated"`
	// The request body parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"type":"object","properties":{"query":{"type":"string"}},"required":["query"]}
	InputSchema map[string]interface{} `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	// The name of the MCP tool.
	//
	// This parameter is required.
	//
	// example:
	//
	// query_logs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The request body parameters.
	//
	// example:
	//
	// {"type":"object"}
	OutputSchema map[string]interface{} `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
	// The title of the MCP tool.
	//
	// example:
	//
	// 查询日志
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s UpdateMcpServiceRequestTools) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServiceRequestTools) GoString() string {
	return s.String()
}

func (s *UpdateMcpServiceRequestTools) GetAnnotations() map[string]interface{} {
	return s.Annotations
}

func (s *UpdateMcpServiceRequestTools) GetConfirm() *bool {
	return s.Confirm
}

func (s *UpdateMcpServiceRequestTools) GetDescription() *string {
	return s.Description
}

func (s *UpdateMcpServiceRequestTools) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateMcpServiceRequestTools) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateMcpServiceRequestTools) GetExecution() map[string]interface{} {
	return s.Execution
}

func (s *UpdateMcpServiceRequestTools) GetIcons() []map[string]interface{} {
	return s.Icons
}

func (s *UpdateMcpServiceRequestTools) GetInputSchema() map[string]interface{} {
	return s.InputSchema
}

func (s *UpdateMcpServiceRequestTools) GetName() *string {
	return s.Name
}

func (s *UpdateMcpServiceRequestTools) GetOutputSchema() map[string]interface{} {
	return s.OutputSchema
}

func (s *UpdateMcpServiceRequestTools) GetTitle() *string {
	return s.Title
}

func (s *UpdateMcpServiceRequestTools) SetAnnotations(v map[string]interface{}) *UpdateMcpServiceRequestTools {
	s.Annotations = v
	return s
}

func (s *UpdateMcpServiceRequestTools) SetConfirm(v bool) *UpdateMcpServiceRequestTools {
	s.Confirm = &v
	return s
}

func (s *UpdateMcpServiceRequestTools) SetDescription(v string) *UpdateMcpServiceRequestTools {
	s.Description = &v
	return s
}

func (s *UpdateMcpServiceRequestTools) SetDisplayName(v string) *UpdateMcpServiceRequestTools {
	s.DisplayName = &v
	return s
}

func (s *UpdateMcpServiceRequestTools) SetEnable(v bool) *UpdateMcpServiceRequestTools {
	s.Enable = &v
	return s
}

func (s *UpdateMcpServiceRequestTools) SetExecution(v map[string]interface{}) *UpdateMcpServiceRequestTools {
	s.Execution = v
	return s
}

func (s *UpdateMcpServiceRequestTools) SetIcons(v []map[string]interface{}) *UpdateMcpServiceRequestTools {
	s.Icons = v
	return s
}

func (s *UpdateMcpServiceRequestTools) SetInputSchema(v map[string]interface{}) *UpdateMcpServiceRequestTools {
	s.InputSchema = v
	return s
}

func (s *UpdateMcpServiceRequestTools) SetName(v string) *UpdateMcpServiceRequestTools {
	s.Name = &v
	return s
}

func (s *UpdateMcpServiceRequestTools) SetOutputSchema(v map[string]interface{}) *UpdateMcpServiceRequestTools {
	s.OutputSchema = v
	return s
}

func (s *UpdateMcpServiceRequestTools) SetTitle(v string) *UpdateMcpServiceRequestTools {
	s.Title = &v
	return s
}

func (s *UpdateMcpServiceRequestTools) Validate() error {
	return dara.Validate(s)
}

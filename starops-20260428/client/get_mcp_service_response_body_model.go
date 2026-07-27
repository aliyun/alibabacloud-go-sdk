// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMcpService(v *GetMcpServiceResponseBodyMcpService) *GetMcpServiceResponseBody
	GetMcpService() *GetMcpServiceResponseBodyMcpService
	SetRegionId(v string) *GetMcpServiceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetMcpServiceResponseBody
	GetRequestId() *string
}

type GetMcpServiceResponseBody struct {
	// The MCP service details.
	McpService *GetMcpServiceResponseBodyMcpService `json:"mcpService,omitempty" xml:"mcpService,omitempty" type:"Struct"`
	// The region of the current request.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMcpServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpServiceResponseBody) GetMcpService() *GetMcpServiceResponseBodyMcpService {
	return s.McpService
}

func (s *GetMcpServiceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetMcpServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpServiceResponseBody) SetMcpService(v *GetMcpServiceResponseBodyMcpService) *GetMcpServiceResponseBody {
	s.McpService = v
	return s
}

func (s *GetMcpServiceResponseBody) SetRegionId(v string) *GetMcpServiceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetMcpServiceResponseBody) SetRequestId(v string) *GetMcpServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpServiceResponseBody) Validate() error {
	if s.McpService != nil {
		if err := s.McpService.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpServiceResponseBodyMcpService struct {
	// The MCP service connection configuration.
	Connection *GetMcpServiceResponseBodyMcpServiceConnection `json:"connection,omitempty" xml:"connection,omitempty" type:"Struct"`
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
	// Indicates whether the MCP service is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The service name of the MCP service.
	//
	// example:
	//
	// log-query
	McpServiceName *string `json:"mcpServiceName,omitempty" xml:"mcpServiceName,omitempty"`
	// The network connectivity information.
	Network *GetMcpServiceResponseBodyMcpServiceNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// The list of MCP tools.
	//
	// example:
	//
	// [{"name":"query_logs"}]
	Tools []*GetMcpServiceResponseBodyMcpServiceTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s GetMcpServiceResponseBodyMcpService) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServiceResponseBodyMcpService) GoString() string {
	return s.String()
}

func (s *GetMcpServiceResponseBodyMcpService) GetConnection() *GetMcpServiceResponseBodyMcpServiceConnection {
	return s.Connection
}

func (s *GetMcpServiceResponseBodyMcpService) GetDescription() *string {
	return s.Description
}

func (s *GetMcpServiceResponseBodyMcpService) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetMcpServiceResponseBodyMcpService) GetEnable() *bool {
	return s.Enable
}

func (s *GetMcpServiceResponseBodyMcpService) GetMcpServiceName() *string {
	return s.McpServiceName
}

func (s *GetMcpServiceResponseBodyMcpService) GetNetwork() *GetMcpServiceResponseBodyMcpServiceNetwork {
	return s.Network
}

func (s *GetMcpServiceResponseBodyMcpService) GetTools() []*GetMcpServiceResponseBodyMcpServiceTools {
	return s.Tools
}

func (s *GetMcpServiceResponseBodyMcpService) SetConnection(v *GetMcpServiceResponseBodyMcpServiceConnection) *GetMcpServiceResponseBodyMcpService {
	s.Connection = v
	return s
}

func (s *GetMcpServiceResponseBodyMcpService) SetDescription(v string) *GetMcpServiceResponseBodyMcpService {
	s.Description = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpService) SetDisplayName(v string) *GetMcpServiceResponseBodyMcpService {
	s.DisplayName = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpService) SetEnable(v bool) *GetMcpServiceResponseBodyMcpService {
	s.Enable = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpService) SetMcpServiceName(v string) *GetMcpServiceResponseBodyMcpService {
	s.McpServiceName = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpService) SetNetwork(v *GetMcpServiceResponseBodyMcpServiceNetwork) *GetMcpServiceResponseBodyMcpService {
	s.Network = v
	return s
}

func (s *GetMcpServiceResponseBodyMcpService) SetTools(v []*GetMcpServiceResponseBodyMcpServiceTools) *GetMcpServiceResponseBodyMcpService {
	s.Tools = v
	return s
}

func (s *GetMcpServiceResponseBodyMcpService) Validate() error {
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

type GetMcpServiceResponseBodyMcpServiceConnection struct {
	// The authentication configuration of the MCP service.
	Auth *GetMcpServiceResponseBodyMcpServiceConnectionAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// The access endpoint of the MCP service.
	//
	// example:
	//
	// https://example.com/mcp
	Endpoint *string            `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Headers  map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The MCP service platform type. Valid values: AIGateway and Custom.
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
	// The MCP service transport protocol. Valid values: http and sse.
	//
	// example:
	//
	// http
	Transport *string `json:"transport,omitempty" xml:"transport,omitempty"`
}

func (s GetMcpServiceResponseBodyMcpServiceConnection) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServiceResponseBodyMcpServiceConnection) GoString() string {
	return s.String()
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) GetAuth() *GetMcpServiceResponseBodyMcpServiceConnectionAuth {
	return s.Auth
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) GetPlatform() *string {
	return s.Platform
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) GetTimeout() *int64 {
	return s.Timeout
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) GetTransport() *string {
	return s.Transport
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) SetAuth(v *GetMcpServiceResponseBodyMcpServiceConnectionAuth) *GetMcpServiceResponseBodyMcpServiceConnection {
	s.Auth = v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) SetEndpoint(v string) *GetMcpServiceResponseBodyMcpServiceConnection {
	s.Endpoint = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) SetHeaders(v map[string]*string) *GetMcpServiceResponseBodyMcpServiceConnection {
	s.Headers = v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) SetPlatform(v string) *GetMcpServiceResponseBodyMcpServiceConnection {
	s.Platform = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) SetTimeout(v int64) *GetMcpServiceResponseBodyMcpServiceConnection {
	s.Timeout = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) SetTransport(v string) *GetMcpServiceResponseBodyMcpServiceConnection {
	s.Transport = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceConnection) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpServiceResponseBodyMcpServiceConnectionAuth struct {
	// The key-value information required for authentication.
	//
	// example:
	//
	// {"token":"example-token"}
	KeyInfo map[string]*string `json:"keyInfo,omitempty" xml:"keyInfo,omitempty"`
	// The authentication type. Currently, bearer is supported.
	//
	// example:
	//
	// bearer
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMcpServiceResponseBodyMcpServiceConnectionAuth) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServiceResponseBodyMcpServiceConnectionAuth) GoString() string {
	return s.String()
}

func (s *GetMcpServiceResponseBodyMcpServiceConnectionAuth) GetKeyInfo() map[string]*string {
	return s.KeyInfo
}

func (s *GetMcpServiceResponseBodyMcpServiceConnectionAuth) GetType() *string {
	return s.Type
}

func (s *GetMcpServiceResponseBodyMcpServiceConnectionAuth) SetKeyInfo(v map[string]*string) *GetMcpServiceResponseBodyMcpServiceConnectionAuth {
	s.KeyInfo = v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceConnectionAuth) SetType(v string) *GetMcpServiceResponseBodyMcpServiceConnectionAuth {
	s.Type = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceConnectionAuth) Validate() error {
	return dara.Validate(s)
}

type GetMcpServiceResponseBodyMcpServiceNetwork struct {
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
	// The virtual private cloud (VPC) ID.
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

func (s GetMcpServiceResponseBodyMcpServiceNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServiceResponseBodyMcpServiceNetwork) GoString() string {
	return s.String()
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) GetAccessIp() *string {
	return s.AccessIp
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) GetAccessPort() *int64 {
	return s.AccessPort
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) GetMode() *string {
	return s.Mode
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) GetRegion() *string {
	return s.Region
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) GetVswId() *string {
	return s.VswId
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) SetAccessIp(v string) *GetMcpServiceResponseBodyMcpServiceNetwork {
	s.AccessIp = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) SetAccessPort(v int64) *GetMcpServiceResponseBodyMcpServiceNetwork {
	s.AccessPort = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) SetGatewayId(v string) *GetMcpServiceResponseBodyMcpServiceNetwork {
	s.GatewayId = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) SetMcpServerId(v string) *GetMcpServiceResponseBodyMcpServiceNetwork {
	s.McpServerId = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) SetMode(v string) *GetMcpServiceResponseBodyMcpServiceNetwork {
	s.Mode = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) SetRegion(v string) *GetMcpServiceResponseBodyMcpServiceNetwork {
	s.Region = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) SetSecurityGroupId(v string) *GetMcpServiceResponseBodyMcpServiceNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) SetVpcId(v string) *GetMcpServiceResponseBodyMcpServiceNetwork {
	s.VpcId = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) SetVswId(v string) *GetMcpServiceResponseBodyMcpServiceNetwork {
	s.VswId = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceNetwork) Validate() error {
	return dara.Validate(s)
}

type GetMcpServiceResponseBodyMcpServiceTools struct {
	// The annotation information of the MCP tool.
	//
	// example:
	//
	// {}
	Annotations map[string]interface{} `json:"annotations,omitempty" xml:"annotations,omitempty"`
	// Indicates whether user confirmation is required before the MCP tool is called.
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
	// Indicates whether the MCP tool is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The execution configuration of the MCP tool.
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
	// The JSON Schema of the MCP tool input parameters.
	//
	// example:
	//
	// {"type":"object","properties":{"query":{"type":"string"}},"required":["query"]}
	InputSchema map[string]interface{} `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	// The name of the MCP tool.
	//
	// example:
	//
	// query_logs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The JSON Schema of the MCP tool output.
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

func (s GetMcpServiceResponseBodyMcpServiceTools) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServiceResponseBodyMcpServiceTools) GoString() string {
	return s.String()
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) GetAnnotations() map[string]interface{} {
	return s.Annotations
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) GetConfirm() *bool {
	return s.Confirm
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) GetDescription() *string {
	return s.Description
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) GetEnable() *bool {
	return s.Enable
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) GetExecution() map[string]interface{} {
	return s.Execution
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) GetIcons() []map[string]interface{} {
	return s.Icons
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) GetInputSchema() map[string]interface{} {
	return s.InputSchema
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) GetName() *string {
	return s.Name
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) GetOutputSchema() map[string]interface{} {
	return s.OutputSchema
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) GetTitle() *string {
	return s.Title
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) SetAnnotations(v map[string]interface{}) *GetMcpServiceResponseBodyMcpServiceTools {
	s.Annotations = v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) SetConfirm(v bool) *GetMcpServiceResponseBodyMcpServiceTools {
	s.Confirm = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) SetDescription(v string) *GetMcpServiceResponseBodyMcpServiceTools {
	s.Description = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) SetDisplayName(v string) *GetMcpServiceResponseBodyMcpServiceTools {
	s.DisplayName = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) SetEnable(v bool) *GetMcpServiceResponseBodyMcpServiceTools {
	s.Enable = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) SetExecution(v map[string]interface{}) *GetMcpServiceResponseBodyMcpServiceTools {
	s.Execution = v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) SetIcons(v []map[string]interface{}) *GetMcpServiceResponseBodyMcpServiceTools {
	s.Icons = v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) SetInputSchema(v map[string]interface{}) *GetMcpServiceResponseBodyMcpServiceTools {
	s.InputSchema = v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) SetName(v string) *GetMcpServiceResponseBodyMcpServiceTools {
	s.Name = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) SetOutputSchema(v map[string]interface{}) *GetMcpServiceResponseBodyMcpServiceTools {
	s.OutputSchema = v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) SetTitle(v string) *GetMcpServiceResponseBodyMcpServiceTools {
	s.Title = &v
	return s
}

func (s *GetMcpServiceResponseBodyMcpServiceTools) Validate() error {
	return dara.Validate(s)
}

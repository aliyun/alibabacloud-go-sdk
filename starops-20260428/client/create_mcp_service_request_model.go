// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnection(v *CreateMcpServiceRequestConnection) *CreateMcpServiceRequest
	GetConnection() *CreateMcpServiceRequestConnection
	SetDescription(v string) *CreateMcpServiceRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateMcpServiceRequest
	GetDisplayName() *string
	SetEnable(v bool) *CreateMcpServiceRequest
	GetEnable() *bool
	SetMcpServiceName(v string) *CreateMcpServiceRequest
	GetMcpServiceName() *string
	SetNetwork(v *CreateMcpServiceRequestNetwork) *CreateMcpServiceRequest
	GetNetwork() *CreateMcpServiceRequestNetwork
	SetTools(v []*CreateMcpServiceRequestTools) *CreateMcpServiceRequest
	GetTools() []*CreateMcpServiceRequestTools
}

type CreateMcpServiceRequest struct {
	// This parameter is required.
	Connection *CreateMcpServiceRequestConnection `json:"connection,omitempty" xml:"connection,omitempty" type:"Struct"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// test
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	McpServiceName *string `json:"mcpServiceName,omitempty" xml:"mcpServiceName,omitempty"`
	// This parameter is required.
	Network *CreateMcpServiceRequestNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// This parameter is required.
	Tools []*CreateMcpServiceRequestTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s CreateMcpServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateMcpServiceRequest) GetConnection() *CreateMcpServiceRequestConnection {
	return s.Connection
}

func (s *CreateMcpServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpServiceRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateMcpServiceRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateMcpServiceRequest) GetMcpServiceName() *string {
	return s.McpServiceName
}

func (s *CreateMcpServiceRequest) GetNetwork() *CreateMcpServiceRequestNetwork {
	return s.Network
}

func (s *CreateMcpServiceRequest) GetTools() []*CreateMcpServiceRequestTools {
	return s.Tools
}

func (s *CreateMcpServiceRequest) SetConnection(v *CreateMcpServiceRequestConnection) *CreateMcpServiceRequest {
	s.Connection = v
	return s
}

func (s *CreateMcpServiceRequest) SetDescription(v string) *CreateMcpServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateMcpServiceRequest) SetDisplayName(v string) *CreateMcpServiceRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateMcpServiceRequest) SetEnable(v bool) *CreateMcpServiceRequest {
	s.Enable = &v
	return s
}

func (s *CreateMcpServiceRequest) SetMcpServiceName(v string) *CreateMcpServiceRequest {
	s.McpServiceName = &v
	return s
}

func (s *CreateMcpServiceRequest) SetNetwork(v *CreateMcpServiceRequestNetwork) *CreateMcpServiceRequest {
	s.Network = v
	return s
}

func (s *CreateMcpServiceRequest) SetTools(v []*CreateMcpServiceRequestTools) *CreateMcpServiceRequest {
	s.Tools = v
	return s
}

func (s *CreateMcpServiceRequest) Validate() error {
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

type CreateMcpServiceRequestConnection struct {
	Auth *CreateMcpServiceRequestConnectionAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// http://10.0.1.23:8080/mcp
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Custom
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// example:
	//
	// 10000
	Timeout *int64 `json:"timeout,omitempty" xml:"timeout,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http
	Transport *string `json:"transport,omitempty" xml:"transport,omitempty"`
}

func (s CreateMcpServiceRequestConnection) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServiceRequestConnection) GoString() string {
	return s.String()
}

func (s *CreateMcpServiceRequestConnection) GetAuth() *CreateMcpServiceRequestConnectionAuth {
	return s.Auth
}

func (s *CreateMcpServiceRequestConnection) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateMcpServiceRequestConnection) GetPlatform() *string {
	return s.Platform
}

func (s *CreateMcpServiceRequestConnection) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CreateMcpServiceRequestConnection) GetTransport() *string {
	return s.Transport
}

func (s *CreateMcpServiceRequestConnection) SetAuth(v *CreateMcpServiceRequestConnectionAuth) *CreateMcpServiceRequestConnection {
	s.Auth = v
	return s
}

func (s *CreateMcpServiceRequestConnection) SetEndpoint(v string) *CreateMcpServiceRequestConnection {
	s.Endpoint = &v
	return s
}

func (s *CreateMcpServiceRequestConnection) SetPlatform(v string) *CreateMcpServiceRequestConnection {
	s.Platform = &v
	return s
}

func (s *CreateMcpServiceRequestConnection) SetTimeout(v int64) *CreateMcpServiceRequestConnection {
	s.Timeout = &v
	return s
}

func (s *CreateMcpServiceRequestConnection) SetTransport(v string) *CreateMcpServiceRequestConnection {
	s.Transport = &v
	return s
}

func (s *CreateMcpServiceRequestConnection) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpServiceRequestConnectionAuth struct {
	// key
	KeyInfo map[string]*string `json:"keyInfo,omitempty" xml:"keyInfo,omitempty"`
	// example:
	//
	// bearer
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateMcpServiceRequestConnectionAuth) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServiceRequestConnectionAuth) GoString() string {
	return s.String()
}

func (s *CreateMcpServiceRequestConnectionAuth) GetKeyInfo() map[string]*string {
	return s.KeyInfo
}

func (s *CreateMcpServiceRequestConnectionAuth) GetType() *string {
	return s.Type
}

func (s *CreateMcpServiceRequestConnectionAuth) SetKeyInfo(v map[string]*string) *CreateMcpServiceRequestConnectionAuth {
	s.KeyInfo = v
	return s
}

func (s *CreateMcpServiceRequestConnectionAuth) SetType(v string) *CreateMcpServiceRequestConnectionAuth {
	s.Type = &v
	return s
}

func (s *CreateMcpServiceRequestConnectionAuth) Validate() error {
	return dara.Validate(s)
}

type CreateMcpServiceRequestNetwork struct {
	// example:
	//
	// 100.2.243.1
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
	// mcp-xxxx
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// public
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// sg-xxxx
	SecurityGroupId *string `json:"securityGroupId,omitempty" xml:"securityGroupId,omitempty"`
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
	// example:
	//
	// vsw-xxxx
	VswId *string `json:"vswId,omitempty" xml:"vswId,omitempty"`
}

func (s CreateMcpServiceRequestNetwork) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServiceRequestNetwork) GoString() string {
	return s.String()
}

func (s *CreateMcpServiceRequestNetwork) GetAccessIp() *string {
	return s.AccessIp
}

func (s *CreateMcpServiceRequestNetwork) GetAccessPort() *int64 {
	return s.AccessPort
}

func (s *CreateMcpServiceRequestNetwork) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateMcpServiceRequestNetwork) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *CreateMcpServiceRequestNetwork) GetMode() *string {
	return s.Mode
}

func (s *CreateMcpServiceRequestNetwork) GetRegion() *string {
	return s.Region
}

func (s *CreateMcpServiceRequestNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateMcpServiceRequestNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateMcpServiceRequestNetwork) GetVswId() *string {
	return s.VswId
}

func (s *CreateMcpServiceRequestNetwork) SetAccessIp(v string) *CreateMcpServiceRequestNetwork {
	s.AccessIp = &v
	return s
}

func (s *CreateMcpServiceRequestNetwork) SetAccessPort(v int64) *CreateMcpServiceRequestNetwork {
	s.AccessPort = &v
	return s
}

func (s *CreateMcpServiceRequestNetwork) SetGatewayId(v string) *CreateMcpServiceRequestNetwork {
	s.GatewayId = &v
	return s
}

func (s *CreateMcpServiceRequestNetwork) SetMcpServerId(v string) *CreateMcpServiceRequestNetwork {
	s.McpServerId = &v
	return s
}

func (s *CreateMcpServiceRequestNetwork) SetMode(v string) *CreateMcpServiceRequestNetwork {
	s.Mode = &v
	return s
}

func (s *CreateMcpServiceRequestNetwork) SetRegion(v string) *CreateMcpServiceRequestNetwork {
	s.Region = &v
	return s
}

func (s *CreateMcpServiceRequestNetwork) SetSecurityGroupId(v string) *CreateMcpServiceRequestNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateMcpServiceRequestNetwork) SetVpcId(v string) *CreateMcpServiceRequestNetwork {
	s.VpcId = &v
	return s
}

func (s *CreateMcpServiceRequestNetwork) SetVswId(v string) *CreateMcpServiceRequestNetwork {
	s.VswId = &v
	return s
}

func (s *CreateMcpServiceRequestNetwork) Validate() error {
	return dara.Validate(s)
}

type CreateMcpServiceRequestTools struct {
	Annotations map[string]interface{} `json:"annotations,omitempty" xml:"annotations,omitempty"`
	Confirm     *bool                  `json:"confirm,omitempty" xml:"confirm,omitempty"`
	// example:
	//
	// hahahhahaha
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// ahahah
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// True
	Enable    *bool                    `json:"enable,omitempty" xml:"enable,omitempty"`
	Execution map[string]interface{}   `json:"execution,omitempty" xml:"execution,omitempty"`
	Icons     []map[string]interface{} `json:"icons,omitempty" xml:"icons,omitempty" type:"Repeated"`
	// This parameter is required.
	InputSchema map[string]interface{} `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agentloop-1654218979e79fb55dbddac2
	Name         *string                `json:"name,omitempty" xml:"name,omitempty"`
	OutputSchema map[string]interface{} `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
	// example:
	//
	// 1
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateMcpServiceRequestTools) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServiceRequestTools) GoString() string {
	return s.String()
}

func (s *CreateMcpServiceRequestTools) GetAnnotations() map[string]interface{} {
	return s.Annotations
}

func (s *CreateMcpServiceRequestTools) GetConfirm() *bool {
	return s.Confirm
}

func (s *CreateMcpServiceRequestTools) GetDescription() *string {
	return s.Description
}

func (s *CreateMcpServiceRequestTools) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateMcpServiceRequestTools) GetEnable() *bool {
	return s.Enable
}

func (s *CreateMcpServiceRequestTools) GetExecution() map[string]interface{} {
	return s.Execution
}

func (s *CreateMcpServiceRequestTools) GetIcons() []map[string]interface{} {
	return s.Icons
}

func (s *CreateMcpServiceRequestTools) GetInputSchema() map[string]interface{} {
	return s.InputSchema
}

func (s *CreateMcpServiceRequestTools) GetName() *string {
	return s.Name
}

func (s *CreateMcpServiceRequestTools) GetOutputSchema() map[string]interface{} {
	return s.OutputSchema
}

func (s *CreateMcpServiceRequestTools) GetTitle() *string {
	return s.Title
}

func (s *CreateMcpServiceRequestTools) SetAnnotations(v map[string]interface{}) *CreateMcpServiceRequestTools {
	s.Annotations = v
	return s
}

func (s *CreateMcpServiceRequestTools) SetConfirm(v bool) *CreateMcpServiceRequestTools {
	s.Confirm = &v
	return s
}

func (s *CreateMcpServiceRequestTools) SetDescription(v string) *CreateMcpServiceRequestTools {
	s.Description = &v
	return s
}

func (s *CreateMcpServiceRequestTools) SetDisplayName(v string) *CreateMcpServiceRequestTools {
	s.DisplayName = &v
	return s
}

func (s *CreateMcpServiceRequestTools) SetEnable(v bool) *CreateMcpServiceRequestTools {
	s.Enable = &v
	return s
}

func (s *CreateMcpServiceRequestTools) SetExecution(v map[string]interface{}) *CreateMcpServiceRequestTools {
	s.Execution = v
	return s
}

func (s *CreateMcpServiceRequestTools) SetIcons(v []map[string]interface{}) *CreateMcpServiceRequestTools {
	s.Icons = v
	return s
}

func (s *CreateMcpServiceRequestTools) SetInputSchema(v map[string]interface{}) *CreateMcpServiceRequestTools {
	s.InputSchema = v
	return s
}

func (s *CreateMcpServiceRequestTools) SetName(v string) *CreateMcpServiceRequestTools {
	s.Name = &v
	return s
}

func (s *CreateMcpServiceRequestTools) SetOutputSchema(v map[string]interface{}) *CreateMcpServiceRequestTools {
	s.OutputSchema = v
	return s
}

func (s *CreateMcpServiceRequestTools) SetTitle(v string) *CreateMcpServiceRequestTools {
	s.Title = &v
	return s
}

func (s *CreateMcpServiceRequestTools) Validate() error {
	return dara.Validate(s)
}

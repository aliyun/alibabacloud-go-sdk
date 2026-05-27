// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMcpServicesResponseBody
	GetMaxResults() *int32
	SetMcpServices(v *ListMcpServicesResponseBodyMcpServices) *ListMcpServicesResponseBody
	GetMcpServices() *ListMcpServicesResponseBodyMcpServices
	SetNextToken(v string) *ListMcpServicesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMcpServicesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListMcpServicesResponseBody
	GetTotal() *int64
}

type ListMcpServicesResponseBody struct {
	// example:
	//
	// 20
	MaxResults  *int32                                  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	McpServices *ListMcpServicesResponseBodyMcpServices `json:"mcpServices,omitempty" xml:"mcpServices,omitempty" type:"Struct"`
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMcpServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcpServicesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpServicesResponseBody) GetMcpServices() *ListMcpServicesResponseBodyMcpServices {
	return s.McpServices
}

func (s *ListMcpServicesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcpServicesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListMcpServicesResponseBody) SetMaxResults(v int32) *ListMcpServicesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMcpServicesResponseBody) SetMcpServices(v *ListMcpServicesResponseBodyMcpServices) *ListMcpServicesResponseBody {
	s.McpServices = v
	return s
}

func (s *ListMcpServicesResponseBody) SetNextToken(v string) *ListMcpServicesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMcpServicesResponseBody) SetRequestId(v string) *ListMcpServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcpServicesResponseBody) SetTotal(v int64) *ListMcpServicesResponseBody {
	s.Total = &v
	return s
}

func (s *ListMcpServicesResponseBody) Validate() error {
	if s.McpServices != nil {
		if err := s.McpServices.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcpServicesResponseBodyMcpServices struct {
	McpServiceList []*ListMcpServicesResponseBodyMcpServicesMcpServiceList `json:"mcpServiceList,omitempty" xml:"mcpServiceList,omitempty" type:"Repeated"`
}

func (s ListMcpServicesResponseBodyMcpServices) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServicesResponseBodyMcpServices) GoString() string {
	return s.String()
}

func (s *ListMcpServicesResponseBodyMcpServices) GetMcpServiceList() []*ListMcpServicesResponseBodyMcpServicesMcpServiceList {
	return s.McpServiceList
}

func (s *ListMcpServicesResponseBodyMcpServices) SetMcpServiceList(v []*ListMcpServicesResponseBodyMcpServicesMcpServiceList) *ListMcpServicesResponseBodyMcpServices {
	s.McpServiceList = v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServices) Validate() error {
	if s.McpServiceList != nil {
		for _, item := range s.McpServiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMcpServicesResponseBodyMcpServicesMcpServiceList struct {
	Connection  *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection `json:"connection,omitempty" xml:"connection,omitempty" type:"Struct"`
	Description *string                                                         `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName *string                                                         `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// log-query
	McpServiceName *string                                                      `json:"mcpServiceName,omitempty" xml:"mcpServiceName,omitempty"`
	Network        *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	// example:
	//
	// [{"name":"query_logs"}]
	Tools []*ListMcpServicesResponseBodyMcpServicesMcpServiceListTools `json:"tools,omitempty" xml:"tools,omitempty" type:"Repeated"`
}

func (s ListMcpServicesResponseBodyMcpServicesMcpServiceList) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServicesResponseBodyMcpServicesMcpServiceList) GoString() string {
	return s.String()
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) GetConnection() *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection {
	return s.Connection
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) GetDescription() *string {
	return s.Description
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) GetEnable() *bool {
	return s.Enable
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) GetMcpServiceName() *string {
	return s.McpServiceName
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) GetNetwork() *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork {
	return s.Network
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) GetTools() []*ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	return s.Tools
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) SetConnection(v *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) *ListMcpServicesResponseBodyMcpServicesMcpServiceList {
	s.Connection = v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) SetDescription(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceList {
	s.Description = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) SetDisplayName(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceList {
	s.DisplayName = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) SetEnable(v bool) *ListMcpServicesResponseBodyMcpServicesMcpServiceList {
	s.Enable = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) SetMcpServiceName(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceList {
	s.McpServiceName = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) SetNetwork(v *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) *ListMcpServicesResponseBodyMcpServicesMcpServiceList {
	s.Network = v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) SetTools(v []*ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) *ListMcpServicesResponseBodyMcpServicesMcpServiceList {
	s.Tools = v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceList) Validate() error {
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

type ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection struct {
	Auth *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth `json:"auth,omitempty" xml:"auth,omitempty" type:"Struct"`
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
	// example:
	//
	// http
	Transport *string `json:"transport,omitempty" xml:"transport,omitempty"`
}

func (s ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) GoString() string {
	return s.String()
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) GetAuth() *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth {
	return s.Auth
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) GetPlatform() *string {
	return s.Platform
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) GetTimeout() *int64 {
	return s.Timeout
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) GetTransport() *string {
	return s.Transport
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) SetAuth(v *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth) *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection {
	s.Auth = v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) SetEndpoint(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection {
	s.Endpoint = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) SetPlatform(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection {
	s.Platform = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) SetTimeout(v int64) *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection {
	s.Timeout = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) SetTransport(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection {
	s.Transport = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnection) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth struct {
	// example:
	//
	// {"token":"example-token"}
	KeyInfo map[string]*string `json:"keyInfo,omitempty" xml:"keyInfo,omitempty"`
	// example:
	//
	// bearer
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth) GoString() string {
	return s.String()
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth) GetKeyInfo() map[string]*string {
	return s.KeyInfo
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth) GetType() *string {
	return s.Type
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth) SetKeyInfo(v map[string]*string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth {
	s.KeyInfo = v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth) SetType(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth {
	s.Type = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListConnectionAuth) Validate() error {
	return dara.Validate(s)
}

type ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork struct {
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

func (s ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) GoString() string {
	return s.String()
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) GetAccessIp() *string {
	return s.AccessIp
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) GetAccessPort() *int64 {
	return s.AccessPort
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) GetMode() *string {
	return s.Mode
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) GetRegion() *string {
	return s.Region
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) GetVswId() *string {
	return s.VswId
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) SetAccessIp(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork {
	s.AccessIp = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) SetAccessPort(v int64) *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork {
	s.AccessPort = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) SetGatewayId(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork {
	s.GatewayId = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) SetMcpServerId(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork {
	s.McpServerId = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) SetMode(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork {
	s.Mode = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) SetRegion(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork {
	s.Region = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) SetSecurityGroupId(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) SetVpcId(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork {
	s.VpcId = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) SetVswId(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork {
	s.VswId = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListNetwork) Validate() error {
	return dara.Validate(s)
}

type ListMcpServicesResponseBodyMcpServicesMcpServiceListTools struct {
	// example:
	//
	// {}
	Annotations map[string]interface{} `json:"annotations,omitempty" xml:"annotations,omitempty"`
	// example:
	//
	// false
	Confirm     *bool   `json:"confirm,omitempty" xml:"confirm,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// {}
	Execution map[string]interface{} `json:"execution,omitempty" xml:"execution,omitempty"`
	// example:
	//
	// []
	Icons []map[string]interface{} `json:"icons,omitempty" xml:"icons,omitempty" type:"Repeated"`
	// example:
	//
	// {"type":"object","properties":{"query":{"type":"string"}},"required":["query"]}
	InputSchema map[string]interface{} `json:"inputSchema,omitempty" xml:"inputSchema,omitempty"`
	// example:
	//
	// query_logs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// {"type":"object"}
	OutputSchema map[string]interface{} `json:"outputSchema,omitempty" xml:"outputSchema,omitempty"`
	Title        *string                `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GoString() string {
	return s.String()
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GetAnnotations() map[string]interface{} {
	return s.Annotations
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GetConfirm() *bool {
	return s.Confirm
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GetDescription() *string {
	return s.Description
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GetEnable() *bool {
	return s.Enable
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GetExecution() map[string]interface{} {
	return s.Execution
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GetIcons() []map[string]interface{} {
	return s.Icons
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GetInputSchema() map[string]interface{} {
	return s.InputSchema
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GetName() *string {
	return s.Name
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GetOutputSchema() map[string]interface{} {
	return s.OutputSchema
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) GetTitle() *string {
	return s.Title
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) SetAnnotations(v map[string]interface{}) *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	s.Annotations = v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) SetConfirm(v bool) *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	s.Confirm = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) SetDescription(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	s.Description = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) SetDisplayName(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	s.DisplayName = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) SetEnable(v bool) *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	s.Enable = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) SetExecution(v map[string]interface{}) *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	s.Execution = v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) SetIcons(v []map[string]interface{}) *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	s.Icons = v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) SetInputSchema(v map[string]interface{}) *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	s.InputSchema = v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) SetName(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	s.Name = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) SetOutputSchema(v map[string]interface{}) *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	s.OutputSchema = v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) SetTitle(v string) *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools {
	s.Title = &v
	return s
}

func (s *ListMcpServicesResponseBodyMcpServicesMcpServiceListTools) Validate() error {
	return dara.Validate(s)
}

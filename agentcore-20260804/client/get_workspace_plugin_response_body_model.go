// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspacePluginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkspacePluginResponseBody
	GetCode() *string
	SetData(v *GetWorkspacePluginResponseBodyData) *GetWorkspacePluginResponseBody
	GetData() *GetWorkspacePluginResponseBodyData
	SetHttpStatusCode(v int32) *GetWorkspacePluginResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWorkspacePluginResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkspacePluginResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkspacePluginResponseBody
	GetSuccess() *bool
}

type GetWorkspacePluginResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The plug-in details.
	Data *GetWorkspacePluginResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. An error description is returned if the request fails.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetWorkspacePluginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacePluginResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspacePluginResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkspacePluginResponseBody) GetData() *GetWorkspacePluginResponseBodyData {
	return s.Data
}

func (s *GetWorkspacePluginResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWorkspacePluginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkspacePluginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspacePluginResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkspacePluginResponseBody) SetCode(v string) *GetWorkspacePluginResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkspacePluginResponseBody) SetData(v *GetWorkspacePluginResponseBodyData) *GetWorkspacePluginResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkspacePluginResponseBody) SetHttpStatusCode(v int32) *GetWorkspacePluginResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkspacePluginResponseBody) SetMessage(v string) *GetWorkspacePluginResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkspacePluginResponseBody) SetRequestId(v string) *GetWorkspacePluginResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspacePluginResponseBody) SetSuccess(v bool) *GetWorkspacePluginResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkspacePluginResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspacePluginResponseBodyData struct {
	// The user-configurable properties currently in effect for the plug-in. This field is empty if the plug-in is not installed.
	Config *GetWorkspacePluginResponseBodyDataConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// Indicates whether the plug-in is enabled. The value is true when the status is ENABLED.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The plug-in name.
	//
	// example:
	//
	// collaboration
	PluginName *string `json:"pluginName,omitempty" xml:"pluginName,omitempty"`
	// The plug-in status. Valid values: DISABLED, ENABLING, ENABLED, ENABLE_FAILED, DISABLING, DISABLE_FAILED.
	//
	// example:
	//
	// ENABLED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetWorkspacePluginResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacePluginResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkspacePluginResponseBodyData) GetConfig() *GetWorkspacePluginResponseBodyDataConfig {
	return s.Config
}

func (s *GetWorkspacePluginResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetWorkspacePluginResponseBodyData) GetPluginName() *string {
	return s.PluginName
}

func (s *GetWorkspacePluginResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetWorkspacePluginResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspacePluginResponseBodyData) SetConfig(v *GetWorkspacePluginResponseBodyDataConfig) *GetWorkspacePluginResponseBodyData {
	s.Config = v
	return s
}

func (s *GetWorkspacePluginResponseBodyData) SetEnabled(v bool) *GetWorkspacePluginResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetWorkspacePluginResponseBodyData) SetPluginName(v string) *GetWorkspacePluginResponseBodyData {
	s.PluginName = &v
	return s
}

func (s *GetWorkspacePluginResponseBodyData) SetStatus(v string) *GetWorkspacePluginResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetWorkspacePluginResponseBodyData) SetWorkspaceId(v string) *GetWorkspacePluginResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspacePluginResponseBodyData) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspacePluginResponseBodyDataConfig struct {
	// The AgentLoop plug-in configuration.
	AgentLoop *GetWorkspacePluginResponseBodyDataConfigAgentLoop `json:"agentLoop,omitempty" xml:"agentLoop,omitempty" type:"Struct"`
	// The network configuration of the plug-in, including public network access configuration and VPC configuration.
	Network *GetWorkspacePluginResponseBodyDataConfigNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
}

func (s GetWorkspacePluginResponseBodyDataConfig) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacePluginResponseBodyDataConfig) GoString() string {
	return s.String()
}

func (s *GetWorkspacePluginResponseBodyDataConfig) GetAgentLoop() *GetWorkspacePluginResponseBodyDataConfigAgentLoop {
	return s.AgentLoop
}

func (s *GetWorkspacePluginResponseBodyDataConfig) GetNetwork() *GetWorkspacePluginResponseBodyDataConfigNetwork {
	return s.Network
}

func (s *GetWorkspacePluginResponseBodyDataConfig) SetAgentLoop(v *GetWorkspacePluginResponseBodyDataConfigAgentLoop) *GetWorkspacePluginResponseBodyDataConfig {
	s.AgentLoop = v
	return s
}

func (s *GetWorkspacePluginResponseBodyDataConfig) SetNetwork(v *GetWorkspacePluginResponseBodyDataConfigNetwork) *GetWorkspacePluginResponseBodyDataConfig {
	s.Network = v
	return s
}

func (s *GetWorkspacePluginResponseBodyDataConfig) Validate() error {
	if s.AgentLoop != nil {
		if err := s.AgentLoop.Validate(); err != nil {
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

type GetWorkspacePluginResponseBodyDataConfigAgentLoop struct {
	// The AgentSpace name associated with the AgentLoop plug-in.
	//
	// example:
	//
	// agentcore-ws-123456
	AgentSpaceName *string `json:"agentSpaceName,omitempty" xml:"agentSpaceName,omitempty"`
	// The creation time of the AgentSpace in UTC in RFC 3339 format.
	//
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The region ID where the AgentSpace resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetWorkspacePluginResponseBodyDataConfigAgentLoop) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacePluginResponseBodyDataConfigAgentLoop) GoString() string {
	return s.String()
}

func (s *GetWorkspacePluginResponseBodyDataConfigAgentLoop) GetAgentSpaceName() *string {
	return s.AgentSpaceName
}

func (s *GetWorkspacePluginResponseBodyDataConfigAgentLoop) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetWorkspacePluginResponseBodyDataConfigAgentLoop) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWorkspacePluginResponseBodyDataConfigAgentLoop) SetAgentSpaceName(v string) *GetWorkspacePluginResponseBodyDataConfigAgentLoop {
	s.AgentSpaceName = &v
	return s
}

func (s *GetWorkspacePluginResponseBodyDataConfigAgentLoop) SetCreatedAt(v string) *GetWorkspacePluginResponseBodyDataConfigAgentLoop {
	s.CreatedAt = &v
	return s
}

func (s *GetWorkspacePluginResponseBodyDataConfigAgentLoop) SetRegionId(v string) *GetWorkspacePluginResponseBodyDataConfigAgentLoop {
	s.RegionId = &v
	return s
}

func (s *GetWorkspacePluginResponseBodyDataConfigAgentLoop) Validate() error {
	return dara.Validate(s)
}

type GetWorkspacePluginResponseBodyDataConfigNetwork struct {
	// The public network access configuration.
	Internet *GetWorkspacePluginResponseBodyDataConfigNetworkInternet `json:"internet,omitempty" xml:"internet,omitempty" type:"Struct"`
	// The user VPC configuration.
	Vpc *GetWorkspacePluginResponseBodyDataConfigNetworkVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
}

func (s GetWorkspacePluginResponseBodyDataConfigNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacePluginResponseBodyDataConfigNetwork) GoString() string {
	return s.String()
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetwork) GetInternet() *GetWorkspacePluginResponseBodyDataConfigNetworkInternet {
	return s.Internet
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetwork) GetVpc() *GetWorkspacePluginResponseBodyDataConfigNetworkVpc {
	return s.Vpc
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetwork) SetInternet(v *GetWorkspacePluginResponseBodyDataConfigNetworkInternet) *GetWorkspacePluginResponseBodyDataConfigNetwork {
	s.Internet = v
	return s
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetwork) SetVpc(v *GetWorkspacePluginResponseBodyDataConfigNetworkVpc) *GetWorkspacePluginResponseBodyDataConfigNetwork {
	s.Vpc = v
	return s
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetwork) Validate() error {
	if s.Internet != nil {
		if err := s.Internet.Validate(); err != nil {
			return err
		}
	}
	if s.Vpc != nil {
		if err := s.Vpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspacePluginResponseBodyDataConfigNetworkInternet struct {
	// Indicates whether public network access is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s GetWorkspacePluginResponseBodyDataConfigNetworkInternet) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacePluginResponseBodyDataConfigNetworkInternet) GoString() string {
	return s.String()
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetworkInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetworkInternet) SetEnabled(v bool) *GetWorkspacePluginResponseBodyDataConfigNetworkInternet {
	s.Enabled = &v
	return s
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetworkInternet) Validate() error {
	return dara.Validate(s)
}

type GetWorkspacePluginResponseBodyDataConfigNetworkVpc struct {
	// Indicates whether VPC network access is enabled for the collaboration plug-in.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of vSwitch IDs used for plug-in deployment.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID used for plug-in deployment.
	//
	// example:
	//
	// vpc-bp1example
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetWorkspacePluginResponseBodyDataConfigNetworkVpc) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspacePluginResponseBodyDataConfigNetworkVpc) GoString() string {
	return s.String()
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetworkVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetworkVpc) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetworkVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetworkVpc) SetEnabled(v bool) *GetWorkspacePluginResponseBodyDataConfigNetworkVpc {
	s.Enabled = &v
	return s
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetworkVpc) SetVSwitchIds(v []*string) *GetWorkspacePluginResponseBodyDataConfigNetworkVpc {
	s.VSwitchIds = v
	return s
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetworkVpc) SetVpcId(v string) *GetWorkspacePluginResponseBodyDataConfigNetworkVpc {
	s.VpcId = &v
	return s
}

func (s *GetWorkspacePluginResponseBodyDataConfigNetworkVpc) Validate() error {
	return dara.Validate(s)
}

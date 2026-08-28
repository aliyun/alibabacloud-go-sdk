// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallWorkspacePluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *InstallWorkspacePluginRequestBody) *InstallWorkspacePluginRequest
	GetBody() *InstallWorkspacePluginRequestBody
	SetClientToken(v string) *InstallWorkspacePluginRequest
	GetClientToken() *string
}

type InstallWorkspacePluginRequest struct {
	// The request body for installing a plugin.
	Body *InstallWorkspacePluginRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client idempotency token.
	//
	// example:
	//
	// workspace-plugin-install-20260810-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s InstallWorkspacePluginRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallWorkspacePluginRequest) GoString() string {
	return s.String()
}

func (s *InstallWorkspacePluginRequest) GetBody() *InstallWorkspacePluginRequestBody {
	return s.Body
}

func (s *InstallWorkspacePluginRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *InstallWorkspacePluginRequest) SetBody(v *InstallWorkspacePluginRequestBody) *InstallWorkspacePluginRequest {
	s.Body = v
	return s
}

func (s *InstallWorkspacePluginRequest) SetClientToken(v string) *InstallWorkspacePluginRequest {
	s.ClientToken = &v
	return s
}

func (s *InstallWorkspacePluginRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallWorkspacePluginRequestBody struct {
	// The plugin-specific configuration. The configuration structure is determined by pluginName. Currently, the collaboration plugin supports network.
	Config *InstallWorkspacePluginRequestBodyConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
}

func (s InstallWorkspacePluginRequestBody) String() string {
	return dara.Prettify(s)
}

func (s InstallWorkspacePluginRequestBody) GoString() string {
	return s.String()
}

func (s *InstallWorkspacePluginRequestBody) GetConfig() *InstallWorkspacePluginRequestBodyConfig {
	return s.Config
}

func (s *InstallWorkspacePluginRequestBody) SetConfig(v *InstallWorkspacePluginRequestBodyConfig) *InstallWorkspacePluginRequestBody {
	s.Config = v
	return s
}

func (s *InstallWorkspacePluginRequestBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallWorkspacePluginRequestBodyConfig struct {
	// The network configuration used by the collaboration plugin. If not specified, the server uses the existing network configuration of the workspace.
	Network *InstallWorkspacePluginRequestBodyConfigNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
}

func (s InstallWorkspacePluginRequestBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s InstallWorkspacePluginRequestBodyConfig) GoString() string {
	return s.String()
}

func (s *InstallWorkspacePluginRequestBodyConfig) GetNetwork() *InstallWorkspacePluginRequestBodyConfigNetwork {
	return s.Network
}

func (s *InstallWorkspacePluginRequestBodyConfig) SetNetwork(v *InstallWorkspacePluginRequestBodyConfigNetwork) *InstallWorkspacePluginRequestBodyConfig {
	s.Network = v
	return s
}

func (s *InstallWorkspacePluginRequestBodyConfig) Validate() error {
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InstallWorkspacePluginRequestBodyConfigNetwork struct {
	// Controls whether the collaboration component is allowed to access the public network. This configuration only controls public network access capability and does not expose the component service to the public network.
	Internet *InstallWorkspacePluginRequestBodyConfigNetworkInternet `json:"internet,omitempty" xml:"internet,omitempty" type:"Struct"`
	// The user VPC and vSwitch list used for deploying the collaboration plugin. The zones corresponding to the vSwitches are queried by the server and do not need to be provided by the user.
	Vpc *InstallWorkspacePluginRequestBodyConfigNetworkVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
}

func (s InstallWorkspacePluginRequestBodyConfigNetwork) String() string {
	return dara.Prettify(s)
}

func (s InstallWorkspacePluginRequestBodyConfigNetwork) GoString() string {
	return s.String()
}

func (s *InstallWorkspacePluginRequestBodyConfigNetwork) GetInternet() *InstallWorkspacePluginRequestBodyConfigNetworkInternet {
	return s.Internet
}

func (s *InstallWorkspacePluginRequestBodyConfigNetwork) GetVpc() *InstallWorkspacePluginRequestBodyConfigNetworkVpc {
	return s.Vpc
}

func (s *InstallWorkspacePluginRequestBodyConfigNetwork) SetInternet(v *InstallWorkspacePluginRequestBodyConfigNetworkInternet) *InstallWorkspacePluginRequestBodyConfigNetwork {
	s.Internet = v
	return s
}

func (s *InstallWorkspacePluginRequestBodyConfigNetwork) SetVpc(v *InstallWorkspacePluginRequestBodyConfigNetworkVpc) *InstallWorkspacePluginRequestBodyConfigNetwork {
	s.Vpc = v
	return s
}

func (s *InstallWorkspacePluginRequestBodyConfigNetwork) Validate() error {
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

type InstallWorkspacePluginRequestBodyConfigNetworkInternet struct {
	// Specifies whether to enable public network access. If set to true without a VPC specified, PUB_NET is used. If set to true with a VPC specified, PRIVATE_PUBNET is used. If only a VPC is specified, PRIVATE_NET is used. At least one of public network or VPC must be configured.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s InstallWorkspacePluginRequestBodyConfigNetworkInternet) String() string {
	return dara.Prettify(s)
}

func (s InstallWorkspacePluginRequestBodyConfigNetworkInternet) GoString() string {
	return s.String()
}

func (s *InstallWorkspacePluginRequestBodyConfigNetworkInternet) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallWorkspacePluginRequestBodyConfigNetworkInternet) SetEnabled(v bool) *InstallWorkspacePluginRequestBodyConfigNetworkInternet {
	s.Enabled = &v
	return s
}

func (s *InstallWorkspacePluginRequestBodyConfigNetworkInternet) Validate() error {
	return dara.Validate(s)
}

type InstallWorkspacePluginRequestBodyConfigNetworkVpc struct {
	// Specifies whether the collaboration plugin uses VPC networking. If set to false, vpcId and vSwitchIds are ignored. If set to true, you must provide both vpcId and at least two vSwitchIds.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of vSwitch IDs. The collaboration plugin requires that the vSwitches cover at least two different zones, and all vSwitches must belong to the VPC specified by vpcId.
	//
	// This parameter is required.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp1234567890
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s InstallWorkspacePluginRequestBodyConfigNetworkVpc) String() string {
	return dara.Prettify(s)
}

func (s InstallWorkspacePluginRequestBodyConfigNetworkVpc) GoString() string {
	return s.String()
}

func (s *InstallWorkspacePluginRequestBodyConfigNetworkVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *InstallWorkspacePluginRequestBodyConfigNetworkVpc) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *InstallWorkspacePluginRequestBodyConfigNetworkVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *InstallWorkspacePluginRequestBodyConfigNetworkVpc) SetEnabled(v bool) *InstallWorkspacePluginRequestBodyConfigNetworkVpc {
	s.Enabled = &v
	return s
}

func (s *InstallWorkspacePluginRequestBodyConfigNetworkVpc) SetVSwitchIds(v []*string) *InstallWorkspacePluginRequestBodyConfigNetworkVpc {
	s.VSwitchIds = v
	return s
}

func (s *InstallWorkspacePluginRequestBodyConfigNetworkVpc) SetVpcId(v string) *InstallWorkspacePluginRequestBodyConfigNetworkVpc {
	s.VpcId = &v
	return s
}

func (s *InstallWorkspacePluginRequestBodyConfigNetworkVpc) Validate() error {
	return dara.Validate(s)
}

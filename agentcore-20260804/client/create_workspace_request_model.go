// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateWorkspaceRequestBody) *CreateWorkspaceRequest
	GetBody() *CreateWorkspaceRequestBody
	SetClientToken(v string) *CreateWorkspaceRequest
	GetClientToken() *string
}

type CreateWorkspaceRequest struct {
	// The workspace creation request body.
	Body *CreateWorkspaceRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client idempotency token.
	//
	// example:
	//
	// workspace-create-20260805-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) GetBody() *CreateWorkspaceRequestBody {
	return s.Body
}

func (s *CreateWorkspaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateWorkspaceRequest) SetBody(v *CreateWorkspaceRequestBody) *CreateWorkspaceRequest {
	s.Body = v
	return s
}

func (s *CreateWorkspaceRequest) SetClientToken(v string) *CreateWorkspaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWorkspaceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceRequestBody struct {
	// The name of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// production-agents
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration of the workspace.
	NetworkConfiguration *CreateWorkspaceRequestBodyNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The workspace storage type. PLATFORM indicates platform-managed storage, and PRIVATE indicates user private OSS storage.
	//
	// example:
	//
	// PRIVATE
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
}

func (s CreateWorkspaceRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequestBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequestBody) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceRequestBody) GetNetworkConfiguration() *CreateWorkspaceRequestBodyNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateWorkspaceRequestBody) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateWorkspaceRequestBody) SetName(v string) *CreateWorkspaceRequestBody {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceRequestBody) SetNetworkConfiguration(v *CreateWorkspaceRequestBodyNetworkConfiguration) *CreateWorkspaceRequestBody {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateWorkspaceRequestBody) SetStorageType(v string) *CreateWorkspaceRequestBody {
	s.StorageType = &v
	return s
}

func (s *CreateWorkspaceRequestBody) Validate() error {
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceRequestBodyNetworkConfiguration struct {
	// Specifies whether to allow public outbound access from the platform. If this parameter is not specified, the default value is true. This setting only controls whether the latest versions of managed and custom agents use the public outbound endpoint of the platform. It does not affect fixed historical versions, model or MCP proxy endpoints, or the public outbound endpoint of the user VPC.
	//
	// example:
	//
	// true
	PublicEgressEnabled *bool `json:"publicEgressEnabled,omitempty" xml:"publicEgressEnabled,omitempty"`
	// Specifies whether to allow public inbound access. If this parameter is not specified, the default value is true. If disabled, the public inbound endpoints of all agent endpoints in the workspace are closed, and the original public inbound endpoints are restored when re-enabled. INNER_DIRECT and user VPC endpoints are not affected. Collaboration and Element Web endpoints are not controlled by this parameter.
	//
	// example:
	//
	// true
	PublicIngressEnabled *bool `json:"publicIngressEnabled,omitempty" xml:"publicIngressEnabled,omitempty"`
	// The user VPC network configuration.
	Vpc *CreateWorkspaceRequestBodyNetworkConfigurationVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
}

func (s CreateWorkspaceRequestBodyNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequestBodyNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequestBodyNetworkConfiguration) GetPublicEgressEnabled() *bool {
	return s.PublicEgressEnabled
}

func (s *CreateWorkspaceRequestBodyNetworkConfiguration) GetPublicIngressEnabled() *bool {
	return s.PublicIngressEnabled
}

func (s *CreateWorkspaceRequestBodyNetworkConfiguration) GetVpc() *CreateWorkspaceRequestBodyNetworkConfigurationVpc {
	return s.Vpc
}

func (s *CreateWorkspaceRequestBodyNetworkConfiguration) SetPublicEgressEnabled(v bool) *CreateWorkspaceRequestBodyNetworkConfiguration {
	s.PublicEgressEnabled = &v
	return s
}

func (s *CreateWorkspaceRequestBodyNetworkConfiguration) SetPublicIngressEnabled(v bool) *CreateWorkspaceRequestBodyNetworkConfiguration {
	s.PublicIngressEnabled = &v
	return s
}

func (s *CreateWorkspaceRequestBodyNetworkConfiguration) SetVpc(v *CreateWorkspaceRequestBodyNetworkConfigurationVpc) *CreateWorkspaceRequestBodyNetworkConfiguration {
	s.Vpc = v
	return s
}

func (s *CreateWorkspaceRequestBodyNetworkConfiguration) Validate() error {
	if s.Vpc != nil {
		if err := s.Vpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceRequestBodyNetworkConfigurationVpc struct {
	// Specifies whether to enable the VPC network.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of vSwitch IDs. When the VPC network is enabled, the list must contain at least one vSwitch, and all vSwitches must belong to the VPC specified by VpcId.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The user VPC ID.
	//
	// example:
	//
	// vpc-bp1234567890
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateWorkspaceRequestBodyNetworkConfigurationVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequestBodyNetworkConfigurationVpc) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequestBodyNetworkConfigurationVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateWorkspaceRequestBodyNetworkConfigurationVpc) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateWorkspaceRequestBodyNetworkConfigurationVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateWorkspaceRequestBodyNetworkConfigurationVpc) SetEnabled(v bool) *CreateWorkspaceRequestBodyNetworkConfigurationVpc {
	s.Enabled = &v
	return s
}

func (s *CreateWorkspaceRequestBodyNetworkConfigurationVpc) SetVSwitchIds(v []*string) *CreateWorkspaceRequestBodyNetworkConfigurationVpc {
	s.VSwitchIds = v
	return s
}

func (s *CreateWorkspaceRequestBodyNetworkConfigurationVpc) SetVpcId(v string) *CreateWorkspaceRequestBodyNetworkConfigurationVpc {
	s.VpcId = &v
	return s
}

func (s *CreateWorkspaceRequestBodyNetworkConfigurationVpc) Validate() error {
	return dara.Validate(s)
}

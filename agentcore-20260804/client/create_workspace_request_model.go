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
	// The request body for creating a workspace.
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
	// If specified, the workspace is directly bound to this AgentSpace, and a new one is not created. The CMS workspace to which it belongs is determined by the server reverse lookup result, and CmsWorkspace validation is not required. When the workspace is deleted, it is only unbound, and the AgentSpace is not deleted. The value must match the current account and region, and the reverse lookup must return the associated CMS workspace. Otherwise, InvalidParameter is returned.
	//
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// If specified, the workspace is bound to this CMS workspace, and the platform does not create a new one. When the workspace is deleted, it is only unbound, and the CMS workspace is not deleted. The value must match the current account and region. Otherwise, InvalidParameter is returned.
	//
	// example:
	//
	// my-cms-workspace
	CmsWorkspace *string `json:"cmsWorkspace,omitempty" xml:"cmsWorkspace,omitempty"`
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
	// The storage type of the workspace. PLATFORM indicates platform-managed storage, and PRIVATE indicates user private OSS storage.
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

func (s *CreateWorkspaceRequestBody) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *CreateWorkspaceRequestBody) GetCmsWorkspace() *string {
	return s.CmsWorkspace
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

func (s *CreateWorkspaceRequestBody) SetAgentSpace(v string) *CreateWorkspaceRequestBody {
	s.AgentSpace = &v
	return s
}

func (s *CreateWorkspaceRequestBody) SetCmsWorkspace(v string) *CreateWorkspaceRequestBody {
	s.CmsWorkspace = &v
	return s
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
	// Specifies whether to allow public egress access through the platform. If not specified, this feature is enabled by default. This parameter only controls whether the latest versions of managed and custom agents use the platform public egress. It does not control fixed historical versions, model or MCP proxy egress, or the public egress of the user VPC.
	//
	// example:
	//
	// true
	PublicEgressEnabled *bool `json:"publicEgressEnabled,omitempty" xml:"publicEgressEnabled,omitempty"`
	// Specifies whether to allow public ingress access. If not specified, this feature is enabled by default. If disabled, the public ingress endpoints of all agent endpoints in the workspace are closed. If re-enabled, the original public ingress endpoints are restored. INNER_DIRECT and user VPC ingress endpoints are not affected. Collaboration and Element Web ingress endpoints are not controlled by this parameter.
	//
	// example:
	//
	// true
	PublicIngressEnabled *bool `json:"publicIngressEnabled,omitempty" xml:"publicIngressEnabled,omitempty"`
	// The network configuration of the user VPC.
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
	// The list of vSwitch IDs. When the VPC network is enabled, at least one vSwitch must be included, and all vSwitches must belong to the VPC specified by VpcId.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The ID of the user VPC.
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

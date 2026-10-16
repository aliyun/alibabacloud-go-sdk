// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateWorkspaceRequestBody) *UpdateWorkspaceRequest
	GetBody() *UpdateWorkspaceRequestBody
	SetClientToken(v string) *UpdateWorkspaceRequest
	GetClientToken() *string
}

type UpdateWorkspaceRequest struct {
	// The request body for updating a workspace.
	Body *UpdateWorkspaceRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// The client token used to ensure idempotence.
	//
	// example:
	//
	// workspace-update-20260805-001
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s UpdateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequest) GetBody() *UpdateWorkspaceRequestBody {
	return s.Body
}

func (s *UpdateWorkspaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateWorkspaceRequest) SetBody(v *UpdateWorkspaceRequestBody) *UpdateWorkspaceRequest {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceRequest) SetClientToken(v string) *UpdateWorkspaceRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateWorkspaceRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceRequestBody struct {
	// The updated name of the workspace.
	//
	// example:
	//
	// production-agents-v2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The updated network configuration of the workspace.
	NetworkConfiguration *UpdateWorkspaceRequestBodyNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
}

func (s UpdateWorkspaceRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequestBody) GetName() *string {
	return s.Name
}

func (s *UpdateWorkspaceRequestBody) GetNetworkConfiguration() *UpdateWorkspaceRequestBodyNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateWorkspaceRequestBody) SetName(v string) *UpdateWorkspaceRequestBody {
	s.Name = &v
	return s
}

func (s *UpdateWorkspaceRequestBody) SetNetworkConfiguration(v *UpdateWorkspaceRequestBodyNetworkConfiguration) *UpdateWorkspaceRequestBody {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateWorkspaceRequestBody) Validate() error {
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceRequestBodyNetworkConfiguration struct {
	// Specifies whether to allow public egress access from the platform. If this parameter is not specified, the original value is retained. Disabling this feature asynchronously cascades to disable the public network switch for agents. Re-enabling this feature does not automatically restore the agent switch. This parameter only controls whether the latest versions of managed and custom agents use the public egress of the platform. It does not control fixed historical versions, model or MCP proxy egress, or the public egress of the user VPC.
	//
	// example:
	//
	// true
	PublicEgressEnabled *bool `json:"publicEgressEnabled,omitempty" xml:"publicEgressEnabled,omitempty"`
	// Specifies whether to allow public ingress access. If this parameter is not specified, the original value is retained. This parameter can be modified together with PublicEgressEnabled and takes effect asynchronously through a unified task. After this feature is disabled, the public ingress for all agent endpoints in the workspace is closed. After this feature is re-enabled, the original public ingress is restored. INNER_DIRECT and user VPC ingress are not affected. Collaboration and Element Web ingress are not controlled by this parameter.
	//
	// example:
	//
	// true
	PublicIngressEnabled *bool `json:"publicIngressEnabled,omitempty" xml:"publicIngressEnabled,omitempty"`
	// The VPC network configuration of the user.
	Vpc *UpdateWorkspaceRequestBodyNetworkConfigurationVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
}

func (s UpdateWorkspaceRequestBodyNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequestBodyNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequestBodyNetworkConfiguration) GetPublicEgressEnabled() *bool {
	return s.PublicEgressEnabled
}

func (s *UpdateWorkspaceRequestBodyNetworkConfiguration) GetPublicIngressEnabled() *bool {
	return s.PublicIngressEnabled
}

func (s *UpdateWorkspaceRequestBodyNetworkConfiguration) GetVpc() *UpdateWorkspaceRequestBodyNetworkConfigurationVpc {
	return s.Vpc
}

func (s *UpdateWorkspaceRequestBodyNetworkConfiguration) SetPublicEgressEnabled(v bool) *UpdateWorkspaceRequestBodyNetworkConfiguration {
	s.PublicEgressEnabled = &v
	return s
}

func (s *UpdateWorkspaceRequestBodyNetworkConfiguration) SetPublicIngressEnabled(v bool) *UpdateWorkspaceRequestBodyNetworkConfiguration {
	s.PublicIngressEnabled = &v
	return s
}

func (s *UpdateWorkspaceRequestBodyNetworkConfiguration) SetVpc(v *UpdateWorkspaceRequestBodyNetworkConfigurationVpc) *UpdateWorkspaceRequestBodyNetworkConfiguration {
	s.Vpc = v
	return s
}

func (s *UpdateWorkspaceRequestBodyNetworkConfiguration) Validate() error {
	if s.Vpc != nil {
		if err := s.Vpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceRequestBodyNetworkConfigurationVpc struct {
	// Specifies whether to enable the VPC network.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of vSwitch IDs. When the VPC network is enabled, the list must contain at least one vSwitch, and all vSwitches must belong to the VPC specified by VpcId.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The ID of the user VPC.
	//
	// example:
	//
	// vpc-bp1234567890
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s UpdateWorkspaceRequestBodyNetworkConfigurationVpc) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequestBodyNetworkConfigurationVpc) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequestBodyNetworkConfigurationVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateWorkspaceRequestBodyNetworkConfigurationVpc) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpdateWorkspaceRequestBodyNetworkConfigurationVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateWorkspaceRequestBodyNetworkConfigurationVpc) SetEnabled(v bool) *UpdateWorkspaceRequestBodyNetworkConfigurationVpc {
	s.Enabled = &v
	return s
}

func (s *UpdateWorkspaceRequestBodyNetworkConfigurationVpc) SetVSwitchIds(v []*string) *UpdateWorkspaceRequestBodyNetworkConfigurationVpc {
	s.VSwitchIds = v
	return s
}

func (s *UpdateWorkspaceRequestBodyNetworkConfigurationVpc) SetVpcId(v string) *UpdateWorkspaceRequestBodyNetworkConfigurationVpc {
	s.VpcId = &v
	return s
}

func (s *UpdateWorkspaceRequestBodyNetworkConfigurationVpc) Validate() error {
	return dara.Validate(s)
}

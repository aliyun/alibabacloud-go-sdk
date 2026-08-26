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
	// The client idempotency token.
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
	// The updated workspace name.
	//
	// example:
	//
	// production-agents-v2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The updated workspace network configuration.
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
	// The user VPC network configuration.
	//
	// This parameter is required.
	Vpc *UpdateWorkspaceRequestBodyNetworkConfigurationVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
}

func (s UpdateWorkspaceRequestBodyNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceRequestBodyNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceRequestBodyNetworkConfiguration) GetVpc() *UpdateWorkspaceRequestBodyNetworkConfigurationVpc {
	return s.Vpc
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
	// Specifies whether to enable VPC networking.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of vSwitch IDs. When VPC networking is enabled, at least one vSwitch must be included, and all vSwitches must belong to the VPC specified by VpcId.
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

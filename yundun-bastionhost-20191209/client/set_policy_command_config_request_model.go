// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyCommandConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandConfig(v *SetPolicyCommandConfigRequestCommandConfig) *SetPolicyCommandConfigRequest
	GetCommandConfig() *SetPolicyCommandConfigRequestCommandConfig
	SetInstanceId(v string) *SetPolicyCommandConfigRequest
	GetInstanceId() *string
	SetPolicyId(v string) *SetPolicyCommandConfigRequest
	GetPolicyId() *string
	SetProjectId(v int64) *SetPolicyCommandConfigRequest
	GetProjectId() *int64
	SetRegionId(v string) *SetPolicyCommandConfigRequest
	GetRegionId() *string
}

type SetPolicyCommandConfigRequest struct {
	// The command control settings.
	//
	// > This parameter applies only to Linux hosts.
	//
	// This parameter is required.
	CommandConfig *SetPolicyCommandConfigRequestCommandConfig `json:"CommandConfig,omitempty" xml:"CommandConfig,omitempty" type:"Struct"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-78v1ghxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the control policy that you want to modify.
	//
	// > You can call the [ListPolicies](https://help.aliyun.com/document_detail/2758876.html) operation to query the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 45
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPolicyCommandConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyCommandConfigRequest) GoString() string {
	return s.String()
}

func (s *SetPolicyCommandConfigRequest) GetCommandConfig() *SetPolicyCommandConfigRequestCommandConfig {
	return s.CommandConfig
}

func (s *SetPolicyCommandConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPolicyCommandConfigRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *SetPolicyCommandConfigRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *SetPolicyCommandConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetPolicyCommandConfigRequest) SetCommandConfig(v *SetPolicyCommandConfigRequestCommandConfig) *SetPolicyCommandConfigRequest {
	s.CommandConfig = v
	return s
}

func (s *SetPolicyCommandConfigRequest) SetInstanceId(v string) *SetPolicyCommandConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPolicyCommandConfigRequest) SetPolicyId(v string) *SetPolicyCommandConfigRequest {
	s.PolicyId = &v
	return s
}

func (s *SetPolicyCommandConfigRequest) SetProjectId(v int64) *SetPolicyCommandConfigRequest {
	s.ProjectId = &v
	return s
}

func (s *SetPolicyCommandConfigRequest) SetRegionId(v string) *SetPolicyCommandConfigRequest {
	s.RegionId = &v
	return s
}

func (s *SetPolicyCommandConfigRequest) Validate() error {
	if s.CommandConfig != nil {
		if err := s.CommandConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetPolicyCommandConfigRequestCommandConfig struct {
	// The command approval settings.
	//
	// > Command approval applies to commands that are not subject to command control (whitelist or blacklist). Command control policies take precedence over command approval.
	Approval *SetPolicyCommandConfigRequestCommandConfigApproval `json:"Approval,omitempty" xml:"Approval,omitempty" type:"Struct"`
	// The command control settings.
	//
	// This parameter is required.
	Deny *SetPolicyCommandConfigRequestCommandConfigDeny `json:"Deny,omitempty" xml:"Deny,omitempty" type:"Struct"`
}

func (s SetPolicyCommandConfigRequestCommandConfig) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyCommandConfigRequestCommandConfig) GoString() string {
	return s.String()
}

func (s *SetPolicyCommandConfigRequestCommandConfig) GetApproval() *SetPolicyCommandConfigRequestCommandConfigApproval {
	return s.Approval
}

func (s *SetPolicyCommandConfigRequestCommandConfig) GetDeny() *SetPolicyCommandConfigRequestCommandConfigDeny {
	return s.Deny
}

func (s *SetPolicyCommandConfigRequestCommandConfig) SetApproval(v *SetPolicyCommandConfigRequestCommandConfigApproval) *SetPolicyCommandConfigRequestCommandConfig {
	s.Approval = v
	return s
}

func (s *SetPolicyCommandConfigRequestCommandConfig) SetDeny(v *SetPolicyCommandConfigRequestCommandConfigDeny) *SetPolicyCommandConfigRequestCommandConfig {
	s.Deny = v
	return s
}

func (s *SetPolicyCommandConfigRequestCommandConfig) Validate() error {
	if s.Approval != nil {
		if err := s.Approval.Validate(); err != nil {
			return err
		}
	}
	if s.Deny != nil {
		if err := s.Deny.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetPolicyCommandConfigRequestCommandConfigApproval struct {
	// The list of commands that require approval.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s SetPolicyCommandConfigRequestCommandConfigApproval) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyCommandConfigRequestCommandConfigApproval) GoString() string {
	return s.String()
}

func (s *SetPolicyCommandConfigRequestCommandConfigApproval) GetCommands() []*string {
	return s.Commands
}

func (s *SetPolicyCommandConfigRequestCommandConfigApproval) SetCommands(v []*string) *SetPolicyCommandConfigRequestCommandConfigApproval {
	s.Commands = v
	return s
}

func (s *SetPolicyCommandConfigRequestCommandConfigApproval) Validate() error {
	return dara.Validate(s)
}

type SetPolicyCommandConfigRequestCommandConfigDeny struct {
	// The command control type. Valid values:
	//
	// - **black**: blacklist mode
	//
	// - **white**: whitelist mode
	//
	// This parameter is required.
	//
	// example:
	//
	// white
	AclType *string `json:"AclType,omitempty" xml:"AclType,omitempty"`
	// The list of commands.
	//
	// > If AclType is set to white, Commands is required.
	Commands []*string `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s SetPolicyCommandConfigRequestCommandConfigDeny) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyCommandConfigRequestCommandConfigDeny) GoString() string {
	return s.String()
}

func (s *SetPolicyCommandConfigRequestCommandConfigDeny) GetAclType() *string {
	return s.AclType
}

func (s *SetPolicyCommandConfigRequestCommandConfigDeny) GetCommands() []*string {
	return s.Commands
}

func (s *SetPolicyCommandConfigRequestCommandConfigDeny) SetAclType(v string) *SetPolicyCommandConfigRequestCommandConfigDeny {
	s.AclType = &v
	return s
}

func (s *SetPolicyCommandConfigRequestCommandConfigDeny) SetCommands(v []*string) *SetPolicyCommandConfigRequestCommandConfigDeny {
	s.Commands = v
	return s
}

func (s *SetPolicyCommandConfigRequestCommandConfigDeny) Validate() error {
	return dara.Validate(s)
}

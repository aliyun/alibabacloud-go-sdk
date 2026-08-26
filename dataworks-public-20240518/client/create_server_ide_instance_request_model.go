// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerIdeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CreateServerIdeInstanceRequestCredentialConfig) *CreateServerIdeInstanceRequest
	GetCredentialConfig() *CreateServerIdeInstanceRequestCredentialConfig
	SetCu(v int32) *CreateServerIdeInstanceRequest
	GetCu() *int32
	SetDatasets(v []*CreateServerIdeInstanceRequestDatasets) *CreateServerIdeInstanceRequest
	GetDatasets() []*CreateServerIdeInstanceRequestDatasets
	SetImageId(v string) *CreateServerIdeInstanceRequest
	GetImageId() *string
	SetImageUrl(v string) *CreateServerIdeInstanceRequest
	GetImageUrl() *string
	SetInstanceName(v string) *CreateServerIdeInstanceRequest
	GetInstanceName() *string
	SetOwner(v string) *CreateServerIdeInstanceRequest
	GetOwner() *string
	SetProjectId(v int64) *CreateServerIdeInstanceRequest
	GetProjectId() *int64
	SetResourceGroupId(v string) *CreateServerIdeInstanceRequest
	GetResourceGroupId() *string
	SetUserCommand(v *CreateServerIdeInstanceRequestUserCommand) *CreateServerIdeInstanceRequest
	GetUserCommand() *CreateServerIdeInstanceRequestUserCommand
	SetUserVpc(v *CreateServerIdeInstanceRequestUserVpc) *CreateServerIdeInstanceRequest
	GetUserVpc() *CreateServerIdeInstanceRequestUserVpc
}

type CreateServerIdeInstanceRequest struct {
	// The credential injection configuration for the instance. After this feature is enabled, you can use the default RAM role chain or specify a custom RAM role.
	CredentialConfig *CreateServerIdeInstanceRequestCredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty" type:"Struct"`
	// The number of CUs used by the instance.
	//
	// example:
	//
	// 10
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The list of datasets mounted to the instance.
	Datasets []*CreateServerIdeInstanceRequestDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The image ID. You can call ListServerIdeImages to obtain the image ID.
	//
	// example:
	//
	// System_serveride_notebook_20240822
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image URL. This parameter is required when you use a non-official DataWorks image.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/example/serveride:latest
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The name of the personal development environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// notebook_dev
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The Alibaba Cloud account ID of the user who owns the instance. If this parameter is not specified, the current caller is used by default.
	//
	// example:
	//
	// 20933221576142****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The DataWorks resource group identifier. You can specify the numeric ID of the resource group or the full identifier in the Serverless_res_group_{tenantId}_{resgId} format.
	//
	// This parameter is required.
	//
	// example:
	//
	// Serverless_res_group_123456789012345_9876543210****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The user command configuration to be executed when the instance starts.
	UserCommand *CreateServerIdeInstanceRequestUserCommand `json:"UserCommand,omitempty" xml:"UserCommand,omitempty" type:"Struct"`
	// The Virtual Private Cloud (VPC) configuration used by the instance.
	UserVpc *CreateServerIdeInstanceRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
}

func (s CreateServerIdeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceRequest) GetCredentialConfig() *CreateServerIdeInstanceRequestCredentialConfig {
	return s.CredentialConfig
}

func (s *CreateServerIdeInstanceRequest) GetCu() *int32 {
	return s.Cu
}

func (s *CreateServerIdeInstanceRequest) GetDatasets() []*CreateServerIdeInstanceRequestDatasets {
	return s.Datasets
}

func (s *CreateServerIdeInstanceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateServerIdeInstanceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *CreateServerIdeInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateServerIdeInstanceRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateServerIdeInstanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateServerIdeInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateServerIdeInstanceRequest) GetUserCommand() *CreateServerIdeInstanceRequestUserCommand {
	return s.UserCommand
}

func (s *CreateServerIdeInstanceRequest) GetUserVpc() *CreateServerIdeInstanceRequestUserVpc {
	return s.UserVpc
}

func (s *CreateServerIdeInstanceRequest) SetCredentialConfig(v *CreateServerIdeInstanceRequestCredentialConfig) *CreateServerIdeInstanceRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateServerIdeInstanceRequest) SetCu(v int32) *CreateServerIdeInstanceRequest {
	s.Cu = &v
	return s
}

func (s *CreateServerIdeInstanceRequest) SetDatasets(v []*CreateServerIdeInstanceRequestDatasets) *CreateServerIdeInstanceRequest {
	s.Datasets = v
	return s
}

func (s *CreateServerIdeInstanceRequest) SetImageId(v string) *CreateServerIdeInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateServerIdeInstanceRequest) SetImageUrl(v string) *CreateServerIdeInstanceRequest {
	s.ImageUrl = &v
	return s
}

func (s *CreateServerIdeInstanceRequest) SetInstanceName(v string) *CreateServerIdeInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateServerIdeInstanceRequest) SetOwner(v string) *CreateServerIdeInstanceRequest {
	s.Owner = &v
	return s
}

func (s *CreateServerIdeInstanceRequest) SetProjectId(v int64) *CreateServerIdeInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateServerIdeInstanceRequest) SetResourceGroupId(v string) *CreateServerIdeInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateServerIdeInstanceRequest) SetUserCommand(v *CreateServerIdeInstanceRequestUserCommand) *CreateServerIdeInstanceRequest {
	s.UserCommand = v
	return s
}

func (s *CreateServerIdeInstanceRequest) SetUserVpc(v *CreateServerIdeInstanceRequestUserVpc) *CreateServerIdeInstanceRequest {
	s.UserVpc = v
	return s
}

func (s *CreateServerIdeInstanceRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Datasets != nil {
		for _, item := range s.Datasets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserCommand != nil {
		if err := s.UserCommand.Validate(); err != nil {
			return err
		}
	}
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServerIdeInstanceRequestCredentialConfig struct {
	// The environment variable role key.
	//
	// example:
	//
	// 0
	AliyunEnvRoleKey *string `json:"AliyunEnvRoleKey,omitempty" xml:"AliyunEnvRoleKey,omitempty"`
	// The list of credential configurations.
	Configs []*CreateServerIdeInstanceRequestCredentialConfigConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// Specifies whether to enable credential injection.
	//
	// example:
	//
	// True
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s CreateServerIdeInstanceRequestCredentialConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceRequestCredentialConfig) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceRequestCredentialConfig) GetAliyunEnvRoleKey() *string {
	return s.AliyunEnvRoleKey
}

func (s *CreateServerIdeInstanceRequestCredentialConfig) GetConfigs() []*CreateServerIdeInstanceRequestCredentialConfigConfigs {
	return s.Configs
}

func (s *CreateServerIdeInstanceRequestCredentialConfig) GetEnable() *bool {
	return s.Enable
}

func (s *CreateServerIdeInstanceRequestCredentialConfig) SetAliyunEnvRoleKey(v string) *CreateServerIdeInstanceRequestCredentialConfig {
	s.AliyunEnvRoleKey = &v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfig) SetConfigs(v []*CreateServerIdeInstanceRequestCredentialConfigConfigs) *CreateServerIdeInstanceRequestCredentialConfig {
	s.Configs = v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfig) SetEnable(v bool) *CreateServerIdeInstanceRequestCredentialConfig {
	s.Enable = &v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfig) Validate() error {
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateServerIdeInstanceRequestCredentialConfigConfigs struct {
	// The identifier key of the credential configuration.
	//
	// example:
	//
	// 0
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The list of roles in the credential configuration.
	Roles []*CreateServerIdeInstanceRequestCredentialConfigConfigsRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The credential configuration type. Valid values:
	//
	// - Role: single role assumption.
	//
	// - RoleChain: role chain assumption.
	//
	// example:
	//
	// RoleChain
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateServerIdeInstanceRequestCredentialConfigConfigs) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceRequestCredentialConfigConfigs) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigs) GetKey() *string {
	return s.Key
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigs) GetRoles() []*CreateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	return s.Roles
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigs) GetType() *string {
	return s.Type
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigs) SetKey(v string) *CreateServerIdeInstanceRequestCredentialConfigConfigs {
	s.Key = &v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigs) SetRoles(v []*CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) *CreateServerIdeInstanceRequestCredentialConfigConfigs {
	s.Roles = v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigs) SetType(v string) *CreateServerIdeInstanceRequestCredentialConfigConfigs {
	s.Type = &v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigs) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateServerIdeInstanceRequestCredentialConfigConfigsRoles struct {
	// The Alibaba Cloud account ID of the principal that owns the role to be assumed.
	//
	// example:
	//
	// 123456789012****
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// The policy used to further restrict the permissions of the role.
	//
	// example:
	//
	// {}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the RAM role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/DataWorksRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The role assumption type. Valid values:
	//
	// - service: assumed by a service.
	//
	// - user: assumed by a user.
	//
	// example:
	//
	// service
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The information about the proxied user.
	UserInfo *CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) GetPolicy() *string {
	return s.Policy
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) GetRoleArn() *string {
	return s.RoleArn
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) GetRoleType() *string {
	return s.RoleType
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) GetUserInfo() *CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo {
	return s.UserInfo
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) SetAssumeRoleFor(v string) *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	s.AssumeRoleFor = &v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) SetPolicy(v string) *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	s.Policy = &v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) SetRoleArn(v string) *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	s.RoleArn = &v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) SetRoleType(v string) *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	s.RoleType = &v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) SetUserInfo(v *CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	s.UserInfo = v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRoles) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo struct {
	// The account ID of the proxied user.
	//
	// example:
	//
	// 20933221576142****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The user type. Valid values:
	//
	// - customer: Alibaba Cloud account.
	//
	// - sub: RAM user.
	//
	// - AssumedRoleUser: RAM role.
	//
	// example:
	//
	// sub
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) GetId() *string {
	return s.Id
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) GetType() *string {
	return s.Type
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) SetId(v string) *CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo {
	s.Id = &v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) SetType(v string) *CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo {
	s.Type = &v
	return s
}

func (s *CreateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) Validate() error {
	return dara.Validate(s)
}

type CreateServerIdeInstanceRequestDatasets struct {
	// The custom mount properties of the dataset. The content is passed as mount options.
	//
	// example:
	//
	// {"fs.oss.download.thread.concurrency":"10"}
	ExtOptions *string `json:"ExtOptions,omitempty" xml:"ExtOptions,omitempty"`
	// The dataset identifier.
	//
	// example:
	//
	// d-vsqjvs****rp5l206u
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The mount path of the dataset in the instance.
	//
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// Specifies whether to mount the dataset in read-only mode.
	//
	// example:
	//
	// false
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The URI of the storage service directory for direct mounting.
	//
	// example:
	//
	// oss://example-bucket/data/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The dataset version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateServerIdeInstanceRequestDatasets) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceRequestDatasets) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceRequestDatasets) GetExtOptions() *string {
	return s.ExtOptions
}

func (s *CreateServerIdeInstanceRequestDatasets) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateServerIdeInstanceRequestDatasets) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateServerIdeInstanceRequestDatasets) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *CreateServerIdeInstanceRequestDatasets) GetUri() *string {
	return s.Uri
}

func (s *CreateServerIdeInstanceRequestDatasets) GetVersion() *int32 {
	return s.Version
}

func (s *CreateServerIdeInstanceRequestDatasets) SetExtOptions(v string) *CreateServerIdeInstanceRequestDatasets {
	s.ExtOptions = &v
	return s
}

func (s *CreateServerIdeInstanceRequestDatasets) SetIdentifier(v string) *CreateServerIdeInstanceRequestDatasets {
	s.Identifier = &v
	return s
}

func (s *CreateServerIdeInstanceRequestDatasets) SetMountPath(v string) *CreateServerIdeInstanceRequestDatasets {
	s.MountPath = &v
	return s
}

func (s *CreateServerIdeInstanceRequestDatasets) SetReadOnly(v bool) *CreateServerIdeInstanceRequestDatasets {
	s.ReadOnly = &v
	return s
}

func (s *CreateServerIdeInstanceRequestDatasets) SetUri(v string) *CreateServerIdeInstanceRequestDatasets {
	s.Uri = &v
	return s
}

func (s *CreateServerIdeInstanceRequestDatasets) SetVersion(v int32) *CreateServerIdeInstanceRequestDatasets {
	s.Version = &v
	return s
}

func (s *CreateServerIdeInstanceRequestDatasets) Validate() error {
	return dara.Validate(s)
}

type CreateServerIdeInstanceRequestUserCommand struct {
	// The command configuration to be executed after the instance starts.
	OnStart *CreateServerIdeInstanceRequestUserCommandOnStart `json:"OnStart,omitempty" xml:"OnStart,omitempty" type:"Struct"`
}

func (s CreateServerIdeInstanceRequestUserCommand) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceRequestUserCommand) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceRequestUserCommand) GetOnStart() *CreateServerIdeInstanceRequestUserCommandOnStart {
	return s.OnStart
}

func (s *CreateServerIdeInstanceRequestUserCommand) SetOnStart(v *CreateServerIdeInstanceRequestUserCommandOnStart) *CreateServerIdeInstanceRequestUserCommand {
	s.OnStart = v
	return s
}

func (s *CreateServerIdeInstanceRequestUserCommand) Validate() error {
	if s.OnStart != nil {
		if err := s.OnStart.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServerIdeInstanceRequestUserCommandOnStart struct {
	// The command content to be executed after the instance starts. The maximum length is 1024 characters.
	//
	// example:
	//
	// echo "serveride ready"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s CreateServerIdeInstanceRequestUserCommandOnStart) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceRequestUserCommandOnStart) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceRequestUserCommandOnStart) GetContent() *string {
	return s.Content
}

func (s *CreateServerIdeInstanceRequestUserCommandOnStart) SetContent(v string) *CreateServerIdeInstanceRequestUserCommandOnStart {
	s.Content = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserCommandOnStart) Validate() error {
	return dara.Validate(s)
}

type CreateServerIdeInstanceRequestUserVpc struct {
	// The list of port forwarding configurations.
	ForwardInfos []*CreateServerIdeInstanceRequestUserVpcForwardInfos `json:"ForwardInfos,omitempty" xml:"ForwardInfos,omitempty" type:"Repeated"`
	// The security group ID.
	//
	// example:
	//
	// sg-bp1****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateServerIdeInstanceRequestUserVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceRequestUserVpc) GetForwardInfos() []*CreateServerIdeInstanceRequestUserVpcForwardInfos {
	return s.ForwardInfos
}

func (s *CreateServerIdeInstanceRequestUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateServerIdeInstanceRequestUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateServerIdeInstanceRequestUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateServerIdeInstanceRequestUserVpc) SetForwardInfos(v []*CreateServerIdeInstanceRequestUserVpcForwardInfos) *CreateServerIdeInstanceRequestUserVpc {
	s.ForwardInfos = v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpc) SetSecurityGroupId(v string) *CreateServerIdeInstanceRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpc) SetVSwitchId(v string) *CreateServerIdeInstanceRequestUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpc) SetVpcId(v string) *CreateServerIdeInstanceRequestUserVpc {
	s.VpcId = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpc) Validate() error {
	if s.ForwardInfos != nil {
		for _, item := range s.ForwardInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateServerIdeInstanceRequestUserVpcForwardInfos struct {
	// The list of access types.
	AccessType []*string `json:"AccessType,omitempty" xml:"AccessType,omitempty" type:"Repeated"`
	// The name of the target container.
	//
	// example:
	//
	// dsw-notebook
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The instance ID of the public Elastic IP Address (EIP).
	//
	// example:
	//
	// eip-bp1****
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	// Specifies whether to enable this port forwarding configuration.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The mapped public port.
	//
	// example:
	//
	// 1024
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The target port in the instance container.
	//
	// example:
	//
	// 22
	ForwardPort *string `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// The name of the port forwarding configuration.
	//
	// example:
	//
	// ssh
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The NAT gateway ID.
	//
	// example:
	//
	// ngw-bp1****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The public key used for SSH access.
	//
	// example:
	//
	// ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQ****
	SSHPublicKey *string `json:"SSHPublicKey,omitempty" xml:"SSHPublicKey,omitempty"`
}

func (s CreateServerIdeInstanceRequestUserVpcForwardInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateServerIdeInstanceRequestUserVpcForwardInfos) GoString() string {
	return s.String()
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) GetAccessType() []*string {
	return s.AccessType
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) GetContainerName() *string {
	return s.ContainerName
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) GetEnable() *bool {
	return s.Enable
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) GetForwardPort() *string {
	return s.ForwardPort
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) GetName() *string {
	return s.Name
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) GetSSHPublicKey() *string {
	return s.SSHPublicKey
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) SetAccessType(v []*string) *CreateServerIdeInstanceRequestUserVpcForwardInfos {
	s.AccessType = v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) SetContainerName(v string) *CreateServerIdeInstanceRequestUserVpcForwardInfos {
	s.ContainerName = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) SetEipAllocationId(v string) *CreateServerIdeInstanceRequestUserVpcForwardInfos {
	s.EipAllocationId = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) SetEnable(v bool) *CreateServerIdeInstanceRequestUserVpcForwardInfos {
	s.Enable = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) SetExternalPort(v string) *CreateServerIdeInstanceRequestUserVpcForwardInfos {
	s.ExternalPort = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) SetForwardPort(v string) *CreateServerIdeInstanceRequestUserVpcForwardInfos {
	s.ForwardPort = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) SetName(v string) *CreateServerIdeInstanceRequestUserVpcForwardInfos {
	s.Name = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) SetNatGatewayId(v string) *CreateServerIdeInstanceRequestUserVpcForwardInfos {
	s.NatGatewayId = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) SetSSHPublicKey(v string) *CreateServerIdeInstanceRequestUserVpcForwardInfos {
	s.SSHPublicKey = &v
	return s
}

func (s *CreateServerIdeInstanceRequestUserVpcForwardInfos) Validate() error {
	return dara.Validate(s)
}

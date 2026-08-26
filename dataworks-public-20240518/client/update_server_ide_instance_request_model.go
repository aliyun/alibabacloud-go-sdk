// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServerIdeInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *UpdateServerIdeInstanceRequestCredentialConfig) *UpdateServerIdeInstanceRequest
	GetCredentialConfig() *UpdateServerIdeInstanceRequestCredentialConfig
	SetCu(v int32) *UpdateServerIdeInstanceRequest
	GetCu() *int32
	SetDatasets(v []*UpdateServerIdeInstanceRequestDatasets) *UpdateServerIdeInstanceRequest
	GetDatasets() []*UpdateServerIdeInstanceRequestDatasets
	SetImageId(v string) *UpdateServerIdeInstanceRequest
	GetImageId() *string
	SetImageUrl(v string) *UpdateServerIdeInstanceRequest
	GetImageUrl() *string
	SetInstanceId(v string) *UpdateServerIdeInstanceRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateServerIdeInstanceRequest
	GetInstanceName() *string
	SetProjectId(v int64) *UpdateServerIdeInstanceRequest
	GetProjectId() *int64
	SetUserVpc(v *UpdateServerIdeInstanceRequestUserVpc) *UpdateServerIdeInstanceRequest
	GetUserVpc() *UpdateServerIdeInstanceRequestUserVpc
}

type UpdateServerIdeInstanceRequest struct {
	// The credential injection configuration for the instance. After this feature is enabled, you can use the default RAM role chain or specify a custom RAM role.
	CredentialConfig *UpdateServerIdeInstanceRequestCredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty" type:"Struct"`
	// The number of CUs used by the instance.
	//
	// example:
	//
	// 10
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The list of datasets mounted to the instance.
	Datasets []*UpdateServerIdeInstanceRequestDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The image ID. You can call ListServerIdeImages to obtain the ID.
	//
	// example:
	//
	// System_serveride_notebook_20240822
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image URL. This parameter is required when you use a non-DataWorks official image.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/example/serveride:latest
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The personal development environment instance ID. You can call ListServerIdeInstances to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 699573
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the personal development environment instance.
	//
	// example:
	//
	// notebook_dev
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The VPC configuration used by the instance.
	UserVpc *UpdateServerIdeInstanceRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
}

func (s UpdateServerIdeInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerIdeInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServerIdeInstanceRequest) GetCredentialConfig() *UpdateServerIdeInstanceRequestCredentialConfig {
	return s.CredentialConfig
}

func (s *UpdateServerIdeInstanceRequest) GetCu() *int32 {
	return s.Cu
}

func (s *UpdateServerIdeInstanceRequest) GetDatasets() []*UpdateServerIdeInstanceRequestDatasets {
	return s.Datasets
}

func (s *UpdateServerIdeInstanceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateServerIdeInstanceRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *UpdateServerIdeInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateServerIdeInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateServerIdeInstanceRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateServerIdeInstanceRequest) GetUserVpc() *UpdateServerIdeInstanceRequestUserVpc {
	return s.UserVpc
}

func (s *UpdateServerIdeInstanceRequest) SetCredentialConfig(v *UpdateServerIdeInstanceRequestCredentialConfig) *UpdateServerIdeInstanceRequest {
	s.CredentialConfig = v
	return s
}

func (s *UpdateServerIdeInstanceRequest) SetCu(v int32) *UpdateServerIdeInstanceRequest {
	s.Cu = &v
	return s
}

func (s *UpdateServerIdeInstanceRequest) SetDatasets(v []*UpdateServerIdeInstanceRequestDatasets) *UpdateServerIdeInstanceRequest {
	s.Datasets = v
	return s
}

func (s *UpdateServerIdeInstanceRequest) SetImageId(v string) *UpdateServerIdeInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateServerIdeInstanceRequest) SetImageUrl(v string) *UpdateServerIdeInstanceRequest {
	s.ImageUrl = &v
	return s
}

func (s *UpdateServerIdeInstanceRequest) SetInstanceId(v string) *UpdateServerIdeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateServerIdeInstanceRequest) SetInstanceName(v string) *UpdateServerIdeInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateServerIdeInstanceRequest) SetProjectId(v int64) *UpdateServerIdeInstanceRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateServerIdeInstanceRequest) SetUserVpc(v *UpdateServerIdeInstanceRequestUserVpc) *UpdateServerIdeInstanceRequest {
	s.UserVpc = v
	return s
}

func (s *UpdateServerIdeInstanceRequest) Validate() error {
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
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServerIdeInstanceRequestCredentialConfig struct {
	// The environment variable role key.
	//
	// example:
	//
	// 0
	AliyunEnvRoleKey *string `json:"AliyunEnvRoleKey,omitempty" xml:"AliyunEnvRoleKey,omitempty"`
	// The list of credential configurations.
	Configs []*UpdateServerIdeInstanceRequestCredentialConfigConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// Specifies whether to enable credential injection.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s UpdateServerIdeInstanceRequestCredentialConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerIdeInstanceRequestCredentialConfig) GoString() string {
	return s.String()
}

func (s *UpdateServerIdeInstanceRequestCredentialConfig) GetAliyunEnvRoleKey() *string {
	return s.AliyunEnvRoleKey
}

func (s *UpdateServerIdeInstanceRequestCredentialConfig) GetConfigs() []*UpdateServerIdeInstanceRequestCredentialConfigConfigs {
	return s.Configs
}

func (s *UpdateServerIdeInstanceRequestCredentialConfig) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateServerIdeInstanceRequestCredentialConfig) SetAliyunEnvRoleKey(v string) *UpdateServerIdeInstanceRequestCredentialConfig {
	s.AliyunEnvRoleKey = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfig) SetConfigs(v []*UpdateServerIdeInstanceRequestCredentialConfigConfigs) *UpdateServerIdeInstanceRequestCredentialConfig {
	s.Configs = v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfig) SetEnable(v bool) *UpdateServerIdeInstanceRequestCredentialConfig {
	s.Enable = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfig) Validate() error {
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

type UpdateServerIdeInstanceRequestCredentialConfigConfigs struct {
	// The identifier key of the credential configuration.
	//
	// example:
	//
	// 0
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The list of roles in the credential configuration.
	Roles []*UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
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

func (s UpdateServerIdeInstanceRequestCredentialConfigConfigs) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerIdeInstanceRequestCredentialConfigConfigs) GoString() string {
	return s.String()
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigs) GetKey() *string {
	return s.Key
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigs) GetRoles() []*UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	return s.Roles
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigs) GetType() *string {
	return s.Type
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigs) SetKey(v string) *UpdateServerIdeInstanceRequestCredentialConfigConfigs {
	s.Key = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigs) SetRoles(v []*UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) *UpdateServerIdeInstanceRequestCredentialConfigConfigs {
	s.Roles = v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigs) SetType(v string) *UpdateServerIdeInstanceRequestCredentialConfigConfigs {
	s.Type = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigs) Validate() error {
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

type UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles struct {
	// The Alibaba Cloud account ID of the principal that assumes the role.
	//
	// example:
	//
	// 123456789012****
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// The policy used to further restrict the role permissions.
	//
	// example:
	//
	// {}
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The ARN of the RAM role.
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
	// The information of the delegated user.
	UserInfo *UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) GoString() string {
	return s.String()
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) GetPolicy() *string {
	return s.Policy
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) GetRoleArn() *string {
	return s.RoleArn
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) GetRoleType() *string {
	return s.RoleType
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) GetUserInfo() *UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo {
	return s.UserInfo
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) SetAssumeRoleFor(v string) *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	s.AssumeRoleFor = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) SetPolicy(v string) *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	s.Policy = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) SetRoleArn(v string) *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	s.RoleArn = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) SetRoleType(v string) *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	s.RoleType = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) SetUserInfo(v *UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles {
	s.UserInfo = v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRoles) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo struct {
	// The account ID of the delegated user.
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

func (s UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) GoString() string {
	return s.String()
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) GetId() *string {
	return s.Id
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) GetType() *string {
	return s.Type
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) SetId(v string) *UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo {
	s.Id = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) SetType(v string) *UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo {
	s.Type = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestCredentialConfigConfigsRolesUserInfo) Validate() error {
	return dara.Validate(s)
}

type UpdateServerIdeInstanceRequestDatasets struct {
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
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The storage service directory URI for direct mounting.
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

func (s UpdateServerIdeInstanceRequestDatasets) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerIdeInstanceRequestDatasets) GoString() string {
	return s.String()
}

func (s *UpdateServerIdeInstanceRequestDatasets) GetExtOptions() *string {
	return s.ExtOptions
}

func (s *UpdateServerIdeInstanceRequestDatasets) GetIdentifier() *string {
	return s.Identifier
}

func (s *UpdateServerIdeInstanceRequestDatasets) GetMountPath() *string {
	return s.MountPath
}

func (s *UpdateServerIdeInstanceRequestDatasets) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *UpdateServerIdeInstanceRequestDatasets) GetUri() *string {
	return s.Uri
}

func (s *UpdateServerIdeInstanceRequestDatasets) GetVersion() *int32 {
	return s.Version
}

func (s *UpdateServerIdeInstanceRequestDatasets) SetExtOptions(v string) *UpdateServerIdeInstanceRequestDatasets {
	s.ExtOptions = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestDatasets) SetIdentifier(v string) *UpdateServerIdeInstanceRequestDatasets {
	s.Identifier = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestDatasets) SetMountPath(v string) *UpdateServerIdeInstanceRequestDatasets {
	s.MountPath = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestDatasets) SetReadOnly(v bool) *UpdateServerIdeInstanceRequestDatasets {
	s.ReadOnly = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestDatasets) SetUri(v string) *UpdateServerIdeInstanceRequestDatasets {
	s.Uri = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestDatasets) SetVersion(v int32) *UpdateServerIdeInstanceRequestDatasets {
	s.Version = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestDatasets) Validate() error {
	return dara.Validate(s)
}

type UpdateServerIdeInstanceRequestUserVpc struct {
	// The list of port forwarding configurations.
	ForwardInfos []*UpdateServerIdeInstanceRequestUserVpcForwardInfos `json:"ForwardInfos,omitempty" xml:"ForwardInfos,omitempty" type:"Repeated"`
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

func (s UpdateServerIdeInstanceRequestUserVpc) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerIdeInstanceRequestUserVpc) GoString() string {
	return s.String()
}

func (s *UpdateServerIdeInstanceRequestUserVpc) GetForwardInfos() []*UpdateServerIdeInstanceRequestUserVpcForwardInfos {
	return s.ForwardInfos
}

func (s *UpdateServerIdeInstanceRequestUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *UpdateServerIdeInstanceRequestUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *UpdateServerIdeInstanceRequestUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateServerIdeInstanceRequestUserVpc) SetForwardInfos(v []*UpdateServerIdeInstanceRequestUserVpcForwardInfos) *UpdateServerIdeInstanceRequestUserVpc {
	s.ForwardInfos = v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpc) SetSecurityGroupId(v string) *UpdateServerIdeInstanceRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpc) SetVSwitchId(v string) *UpdateServerIdeInstanceRequestUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpc) SetVpcId(v string) *UpdateServerIdeInstanceRequestUserVpc {
	s.VpcId = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpc) Validate() error {
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

type UpdateServerIdeInstanceRequestUserVpcForwardInfos struct {
	// The list of access types.
	AccessType []*string `json:"AccessType,omitempty" xml:"AccessType,omitempty" type:"Repeated"`
	// The name of the target container.
	//
	// example:
	//
	// dsw-notebook
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The instance ID of the public EIP.
	//
	// example:
	//
	// eip-bp1****
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	// Specifies whether to enable the port forwarding configuration.
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

func (s UpdateServerIdeInstanceRequestUserVpcForwardInfos) String() string {
	return dara.Prettify(s)
}

func (s UpdateServerIdeInstanceRequestUserVpcForwardInfos) GoString() string {
	return s.String()
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) GetAccessType() []*string {
	return s.AccessType
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) GetContainerName() *string {
	return s.ContainerName
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) GetForwardPort() *string {
	return s.ForwardPort
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) GetName() *string {
	return s.Name
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) GetSSHPublicKey() *string {
	return s.SSHPublicKey
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) SetAccessType(v []*string) *UpdateServerIdeInstanceRequestUserVpcForwardInfos {
	s.AccessType = v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) SetContainerName(v string) *UpdateServerIdeInstanceRequestUserVpcForwardInfos {
	s.ContainerName = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) SetEipAllocationId(v string) *UpdateServerIdeInstanceRequestUserVpcForwardInfos {
	s.EipAllocationId = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) SetEnable(v bool) *UpdateServerIdeInstanceRequestUserVpcForwardInfos {
	s.Enable = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) SetExternalPort(v string) *UpdateServerIdeInstanceRequestUserVpcForwardInfos {
	s.ExternalPort = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) SetForwardPort(v string) *UpdateServerIdeInstanceRequestUserVpcForwardInfos {
	s.ForwardPort = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) SetName(v string) *UpdateServerIdeInstanceRequestUserVpcForwardInfos {
	s.Name = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) SetNatGatewayId(v string) *UpdateServerIdeInstanceRequestUserVpcForwardInfos {
	s.NatGatewayId = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) SetSSHPublicKey(v string) *UpdateServerIdeInstanceRequestUserVpcForwardInfos {
	s.SSHPublicKey = &v
	return s
}

func (s *UpdateServerIdeInstanceRequestUserVpcForwardInfos) Validate() error {
	return dara.Validate(s)
}

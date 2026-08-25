// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServerIdeInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v *GetServerIdeInstanceResponseBodyInstance) *GetServerIdeInstanceResponseBody
	GetInstance() *GetServerIdeInstanceResponseBodyInstance
	SetRequestId(v string) *GetServerIdeInstanceResponseBody
	GetRequestId() *string
}

type GetServerIdeInstanceResponseBody struct {
	// The details of the personal development environment instance.
	Instance *GetServerIdeInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetServerIdeInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServerIdeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetServerIdeInstanceResponseBody) GetInstance() *GetServerIdeInstanceResponseBodyInstance {
	return s.Instance
}

func (s *GetServerIdeInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServerIdeInstanceResponseBody) SetInstance(v *GetServerIdeInstanceResponseBodyInstance) *GetServerIdeInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *GetServerIdeInstanceResponseBody) SetRequestId(v string) *GetServerIdeInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServerIdeInstanceResponseBodyInstance struct {
	// The time when the instance was created. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1756000000000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The credential injection configuration of the instance. After this feature is enabled, you can use the default RAM role chain or specify a custom RAM role.
	CredentialConfig *GetServerIdeInstanceResponseBodyInstanceCredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty" type:"Struct"`
	// The number of CUs used by the instance.
	//
	// example:
	//
	// 10
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The list of datasets mounted to the instance.
	Datasets []*GetServerIdeInstanceResponseBodyInstanceDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
	// The reason why the instance entered the failed state.
	//
	// example:
	//
	// ImagePullBackOff
	FailReason *string `json:"FailReason,omitempty" xml:"FailReason,omitempty"`
	// The ID of the image used by the instance.
	//
	// example:
	//
	// System_serveride_notebook_20240822
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// serveride_notebook
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The image URL.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/example/serveride:latest
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The personal development environment instance ID.
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
	// The account ID of the user who owns the instance.
	//
	// example:
	//
	// 20933221576142****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The DataWorks workspace name.
	//
	// example:
	//
	// example_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The internal numeric ID of the resource group.
	//
	// example:
	//
	// 9876543210
	ResourceGroupId *int64 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource group name.
	//
	// example:
	//
	// serverless_group
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The instance status. Valid values: Creating, Starting, Running, Stopping, Stopped, Updating, Deleting, DELETED, Failed, Arrearage, Saving, SaveFailed, and Saved.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the instance was last updated. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1756003600000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The VPC configuration used by the instance.
	UserVpc *GetServerIdeInstanceResponseBodyInstanceUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
}

func (s GetServerIdeInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s GetServerIdeInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetCredentialConfig() *GetServerIdeInstanceResponseBodyInstanceCredentialConfig {
	return s.CredentialConfig
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetCu() *int32 {
	return s.Cu
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetDatasets() []*GetServerIdeInstanceResponseBodyInstanceDatasets {
	return s.Datasets
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetFailReason() *string {
	return s.FailReason
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetImageId() *string {
	return s.ImageId
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetImageName() *string {
	return s.ImageName
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetResourceGroupId() *int64 {
	return s.ResourceGroupId
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetStatus() *string {
	return s.Status
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetServerIdeInstanceResponseBodyInstance) GetUserVpc() *GetServerIdeInstanceResponseBodyInstanceUserVpc {
	return s.UserVpc
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetCreateTime(v int64) *GetServerIdeInstanceResponseBodyInstance {
	s.CreateTime = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetCredentialConfig(v *GetServerIdeInstanceResponseBodyInstanceCredentialConfig) *GetServerIdeInstanceResponseBodyInstance {
	s.CredentialConfig = v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetCu(v int32) *GetServerIdeInstanceResponseBodyInstance {
	s.Cu = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetDatasets(v []*GetServerIdeInstanceResponseBodyInstanceDatasets) *GetServerIdeInstanceResponseBodyInstance {
	s.Datasets = v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetFailReason(v string) *GetServerIdeInstanceResponseBodyInstance {
	s.FailReason = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetImageId(v string) *GetServerIdeInstanceResponseBodyInstance {
	s.ImageId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetImageName(v string) *GetServerIdeInstanceResponseBodyInstance {
	s.ImageName = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetImageUrl(v string) *GetServerIdeInstanceResponseBodyInstance {
	s.ImageUrl = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetInstanceId(v string) *GetServerIdeInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetInstanceName(v string) *GetServerIdeInstanceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetOwnerId(v string) *GetServerIdeInstanceResponseBodyInstance {
	s.OwnerId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetProjectId(v int64) *GetServerIdeInstanceResponseBodyInstance {
	s.ProjectId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetProjectName(v string) *GetServerIdeInstanceResponseBodyInstance {
	s.ProjectName = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetResourceGroupId(v int64) *GetServerIdeInstanceResponseBodyInstance {
	s.ResourceGroupId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetResourceGroupName(v string) *GetServerIdeInstanceResponseBodyInstance {
	s.ResourceGroupName = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetStatus(v string) *GetServerIdeInstanceResponseBodyInstance {
	s.Status = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetUpdateTime(v int64) *GetServerIdeInstanceResponseBodyInstance {
	s.UpdateTime = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) SetUserVpc(v *GetServerIdeInstanceResponseBodyInstanceUserVpc) *GetServerIdeInstanceResponseBodyInstance {
	s.UserVpc = v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstance) Validate() error {
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

type GetServerIdeInstanceResponseBodyInstanceCredentialConfig struct {
	// The environment variable role key.
	//
	// example:
	//
	// 0
	AliyunEnvRoleKey *string `json:"AliyunEnvRoleKey,omitempty" xml:"AliyunEnvRoleKey,omitempty"`
	// The list of credential configurations.
	Configs []*GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// Indicates whether credential injection is enabled.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s GetServerIdeInstanceResponseBodyInstanceCredentialConfig) String() string {
	return dara.Prettify(s)
}

func (s GetServerIdeInstanceResponseBodyInstanceCredentialConfig) GoString() string {
	return s.String()
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfig) GetAliyunEnvRoleKey() *string {
	return s.AliyunEnvRoleKey
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfig) GetConfigs() []*GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs {
	return s.Configs
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfig) GetEnable() *bool {
	return s.Enable
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfig) SetAliyunEnvRoleKey(v string) *GetServerIdeInstanceResponseBodyInstanceCredentialConfig {
	s.AliyunEnvRoleKey = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfig) SetConfigs(v []*GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs) *GetServerIdeInstanceResponseBodyInstanceCredentialConfig {
	s.Configs = v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfig) SetEnable(v bool) *GetServerIdeInstanceResponseBodyInstanceCredentialConfig {
	s.Enable = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfig) Validate() error {
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

type GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs struct {
	// The identifier key of the credential configuration.
	//
	// example:
	//
	// 0
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The list of roles in the credential configuration.
	Roles []*GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// The credential configuration type. Valid values: Role (single role assumption) and RoleChain (role chain assumption).
	//
	// example:
	//
	// RoleChain
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs) String() string {
	return dara.Prettify(s)
}

func (s GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs) GoString() string {
	return s.String()
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs) GetKey() *string {
	return s.Key
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs) GetRoles() []*GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles {
	return s.Roles
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs) GetType() *string {
	return s.Type
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs) SetKey(v string) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs {
	s.Key = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs) SetRoles(v []*GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs {
	s.Roles = v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs) SetType(v string) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs {
	s.Type = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigs) Validate() error {
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

type GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles struct {
	// The Alibaba Cloud account ID of the principal that owns the assumed role.
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
	// The ARN of the RAM role.
	//
	// example:
	//
	// acs:ram::123456789012****:role/DataWorksRole
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The role assumption type. Valid values: service (assumed by a service) and user (assumed by a user).
	//
	// example:
	//
	// service
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The information of the proxied user.
	UserInfo *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) String() string {
	return dara.Prettify(s)
}

func (s GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) GoString() string {
	return s.String()
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) GetPolicy() *string {
	return s.Policy
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) GetRoleArn() *string {
	return s.RoleArn
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) GetRoleType() *string {
	return s.RoleType
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) GetUserInfo() *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo {
	return s.UserInfo
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) SetAssumeRoleFor(v string) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles {
	s.AssumeRoleFor = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) SetPolicy(v string) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles {
	s.Policy = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) SetRoleArn(v string) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles {
	s.RoleArn = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) SetRoleType(v string) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles {
	s.RoleType = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) SetUserInfo(v *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles {
	s.UserInfo = v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRoles) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo struct {
	// The temporary AccessKey ID used for credential injection.
	//
	// example:
	//
	// STS.N*********7
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The account ID of the proxied user.
	//
	// example:
	//
	// 20933221576142****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The temporary security token used for credential injection.
	//
	// example:
	//
	// DFE32G*******
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The user type. Valid values: customer (Alibaba Cloud account), sub (RAM user), and AssumedRoleUser (RAM role).
	//
	// example:
	//
	// sub
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) GoString() string {
	return s.String()
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) GetId() *string {
	return s.Id
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) GetType() *string {
	return s.Type
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) SetAccessKeyId(v string) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo {
	s.AccessKeyId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) SetId(v string) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo {
	s.Id = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) SetSecurityToken(v string) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo {
	s.SecurityToken = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) SetType(v string) *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo {
	s.Type = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceCredentialConfigConfigsRolesUserInfo) Validate() error {
	return dara.Validate(s)
}

type GetServerIdeInstanceResponseBodyInstanceDatasets struct {
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
	// Indicates whether the dataset is mounted in read-only mode.
	ReadOnly *bool `json:"ReadOnly,omitempty" xml:"ReadOnly,omitempty"`
	// The URI of the storage service directory used for direct mounting.
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

func (s GetServerIdeInstanceResponseBodyInstanceDatasets) String() string {
	return dara.Prettify(s)
}

func (s GetServerIdeInstanceResponseBodyInstanceDatasets) GoString() string {
	return s.String()
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) GetExtOptions() *string {
	return s.ExtOptions
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) GetMountPath() *string {
	return s.MountPath
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) GetUri() *string {
	return s.Uri
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) GetVersion() *int32 {
	return s.Version
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) SetExtOptions(v string) *GetServerIdeInstanceResponseBodyInstanceDatasets {
	s.ExtOptions = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) SetIdentifier(v string) *GetServerIdeInstanceResponseBodyInstanceDatasets {
	s.Identifier = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) SetMountPath(v string) *GetServerIdeInstanceResponseBodyInstanceDatasets {
	s.MountPath = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) SetReadOnly(v bool) *GetServerIdeInstanceResponseBodyInstanceDatasets {
	s.ReadOnly = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) SetUri(v string) *GetServerIdeInstanceResponseBodyInstanceDatasets {
	s.Uri = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) SetVersion(v int32) *GetServerIdeInstanceResponseBodyInstanceDatasets {
	s.Version = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceDatasets) Validate() error {
	return dara.Validate(s)
}

type GetServerIdeInstanceResponseBodyInstanceUserVpc struct {
	// The list of port forwarding configurations.
	ForwardInfos []*GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos `json:"ForwardInfos,omitempty" xml:"ForwardInfos,omitempty" type:"Repeated"`
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

func (s GetServerIdeInstanceResponseBodyInstanceUserVpc) String() string {
	return dara.Prettify(s)
}

func (s GetServerIdeInstanceResponseBodyInstanceUserVpc) GoString() string {
	return s.String()
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpc) GetForwardInfos() []*GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos {
	return s.ForwardInfos
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpc) SetForwardInfos(v []*GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) *GetServerIdeInstanceResponseBodyInstanceUserVpc {
	s.ForwardInfos = v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpc) SetSecurityGroupId(v string) *GetServerIdeInstanceResponseBodyInstanceUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpc) SetVSwitchId(v string) *GetServerIdeInstanceResponseBodyInstanceUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpc) SetVpcId(v string) *GetServerIdeInstanceResponseBodyInstanceUserVpc {
	s.VpcId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpc) Validate() error {
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

type GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos struct {
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
	// Indicates whether this port forwarding configuration is enabled.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The mapped public port.
	//
	// example:
	//
	// 1024
	ExternalPort *string `json:"ExternalPort,omitempty" xml:"ExternalPort,omitempty"`
	// The target port inside the instance container.
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

func (s GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) String() string {
	return dara.Prettify(s)
}

func (s GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) GoString() string {
	return s.String()
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) GetAccessType() []*string {
	return s.AccessType
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) GetContainerName() *string {
	return s.ContainerName
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) GetEnable() *bool {
	return s.Enable
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) GetForwardPort() *string {
	return s.ForwardPort
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) GetName() *string {
	return s.Name
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) GetSSHPublicKey() *string {
	return s.SSHPublicKey
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) SetAccessType(v []*string) *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos {
	s.AccessType = v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) SetContainerName(v string) *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos {
	s.ContainerName = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) SetEipAllocationId(v string) *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos {
	s.EipAllocationId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) SetEnable(v bool) *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos {
	s.Enable = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) SetExternalPort(v string) *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos {
	s.ExternalPort = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) SetForwardPort(v string) *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos {
	s.ForwardPort = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) SetName(v string) *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos {
	s.Name = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) SetNatGatewayId(v string) *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos {
	s.NatGatewayId = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) SetSSHPublicKey(v string) *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos {
	s.SSHPublicKey = &v
	return s
}

func (s *GetServerIdeInstanceResponseBodyInstanceUserVpcForwardInfos) Validate() error {
	return dara.Validate(s)
}

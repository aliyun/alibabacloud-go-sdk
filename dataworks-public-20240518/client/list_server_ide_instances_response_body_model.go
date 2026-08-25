// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerIdeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListServerIdeInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListServerIdeInstancesResponseBody
	GetNextToken() *string
	SetPagingInfo(v *ListServerIdeInstancesResponseBodyPagingInfo) *ListServerIdeInstancesResponseBody
	GetPagingInfo() *ListServerIdeInstancesResponseBodyPagingInfo
	SetRequestId(v string) *ListServerIdeInstancesResponseBody
	GetRequestId() *string
}

type ListServerIdeInstancesResponseBody struct {
	// The maximum number of records returned in this response.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. An empty value indicates that no more results are available.
	//
	// example:
	//
	// CAESG****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The pagination information.
	PagingInfo *ListServerIdeInstancesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListServerIdeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServerIdeInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServerIdeInstancesResponseBody) GetPagingInfo() *ListServerIdeInstancesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListServerIdeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServerIdeInstancesResponseBody) SetMaxResults(v int32) *ListServerIdeInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServerIdeInstancesResponseBody) SetNextToken(v string) *ListServerIdeInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServerIdeInstancesResponseBody) SetPagingInfo(v *ListServerIdeInstancesResponseBodyPagingInfo) *ListServerIdeInstancesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListServerIdeInstancesResponseBody) SetRequestId(v string) *ListServerIdeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServerIdeInstancesResponseBodyPagingInfo struct {
	// The list of personal development environment instances.
	Instances []*ListServerIdeInstancesResponseBodyPagingInfoInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records that match the filter conditions.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServerIdeInstancesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesResponseBodyPagingInfo) GetInstances() []*ListServerIdeInstancesResponseBodyPagingInfoInstances {
	return s.Instances
}

func (s *ListServerIdeInstancesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServerIdeInstancesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServerIdeInstancesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServerIdeInstancesResponseBodyPagingInfo) SetInstances(v []*ListServerIdeInstancesResponseBodyPagingInfoInstances) *ListServerIdeInstancesResponseBodyPagingInfo {
	s.Instances = v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfo) SetPageNumber(v int32) *ListServerIdeInstancesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfo) SetPageSize(v int32) *ListServerIdeInstancesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfo) SetTotalCount(v int32) *ListServerIdeInstancesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfo) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServerIdeInstancesResponseBodyPagingInfoInstances struct {
	// The time when the instance was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1756000000000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The credential injection configuration of the instance. After this feature is enabled, you can use the default RAM role chain or specify a custom RAM role.
	CredentialConfig *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty" type:"Struct"`
	// The number of CUs used by the instance.
	//
	// example:
	//
	// 10
	Cu *int32 `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// The list of datasets mounted to the instance.
	Datasets []*ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets `json:"Datasets,omitempty" xml:"Datasets,omitempty" type:"Repeated"`
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
	// The personal development environment instance name.
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
	// The time when the instance was last updated. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1756003600000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The VPC configuration used by the instance.
	UserVpc *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
}

func (s ListServerIdeInstancesResponseBodyPagingInfoInstances) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesResponseBodyPagingInfoInstances) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetCredentialConfig() *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig {
	return s.CredentialConfig
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetCu() *int32 {
	return s.Cu
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetDatasets() []*ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets {
	return s.Datasets
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetFailReason() *string {
	return s.FailReason
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetImageId() *string {
	return s.ImageId
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetImageName() *string {
	return s.ImageName
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetResourceGroupId() *int64 {
	return s.ResourceGroupId
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetStatus() *string {
	return s.Status
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) GetUserVpc() *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc {
	return s.UserVpc
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetCreateTime(v int64) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.CreateTime = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetCredentialConfig(v *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.CredentialConfig = v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetCu(v int32) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.Cu = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetDatasets(v []*ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.Datasets = v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetFailReason(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.FailReason = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetImageId(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.ImageId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetImageName(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.ImageName = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetImageUrl(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.ImageUrl = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetInstanceId(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.InstanceId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetInstanceName(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.InstanceName = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetOwnerId(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.OwnerId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetProjectId(v int64) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.ProjectId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetProjectName(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.ProjectName = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetResourceGroupId(v int64) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetResourceGroupName(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.ResourceGroupName = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetStatus(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.Status = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetUpdateTime(v int64) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.UpdateTime = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) SetUserVpc(v *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) *ListServerIdeInstancesResponseBodyPagingInfoInstances {
	s.UserVpc = v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstances) Validate() error {
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

type ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig struct {
	// The environment variable role key.
	//
	// example:
	//
	// 0
	AliyunEnvRoleKey *string `json:"AliyunEnvRoleKey,omitempty" xml:"AliyunEnvRoleKey,omitempty"`
	// The list of credential configurations.
	Configs []*ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// Indicates whether credential injection is enabled.
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
}

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig) GetAliyunEnvRoleKey() *string {
	return s.AliyunEnvRoleKey
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig) GetConfigs() []*ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs {
	return s.Configs
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig) GetEnable() *bool {
	return s.Enable
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig) SetAliyunEnvRoleKey(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig {
	s.AliyunEnvRoleKey = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig) SetConfigs(v []*ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig {
	s.Configs = v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig) SetEnable(v bool) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig {
	s.Enable = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfig) Validate() error {
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

type ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs struct {
	// The identifier key of the credential configuration.
	//
	// example:
	//
	// 0
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The list of roles in the credential configuration.
	Roles []*ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
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

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs) GetKey() *string {
	return s.Key
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs) GetRoles() []*ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles {
	return s.Roles
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs) GetType() *string {
	return s.Type
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs) SetKey(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs {
	s.Key = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs) SetRoles(v []*ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs {
	s.Roles = v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs) SetType(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs {
	s.Type = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigs) Validate() error {
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

type ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles struct {
	// The Alibaba Cloud account ID of the entity that owns the role to be assumed.
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
	// The information about the delegated user.
	UserInfo *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) GetPolicy() *string {
	return s.Policy
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) GetRoleArn() *string {
	return s.RoleArn
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) GetRoleType() *string {
	return s.RoleType
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) GetUserInfo() *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo {
	return s.UserInfo
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) SetAssumeRoleFor(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles {
	s.AssumeRoleFor = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) SetPolicy(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles {
	s.Policy = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) SetRoleArn(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles {
	s.RoleArn = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) SetRoleType(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles {
	s.RoleType = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) SetUserInfo(v *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles {
	s.UserInfo = v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRoles) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo struct {
	// The temporary AccessKey ID used for credential injection.
	//
	// example:
	//
	// STS.N*********7
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The account ID of the delegated user.
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

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) GetId() *string {
	return s.Id
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) GetType() *string {
	return s.Type
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) SetAccessKeyId(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo {
	s.AccessKeyId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) SetId(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo {
	s.Id = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) SetSecurityToken(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo {
	s.SecurityToken = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) SetType(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo {
	s.Type = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesCredentialConfigConfigsRolesUserInfo) Validate() error {
	return dara.Validate(s)
}

type ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets struct {
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

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) GetExtOptions() *string {
	return s.ExtOptions
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) GetMountPath() *string {
	return s.MountPath
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) GetReadOnly() *bool {
	return s.ReadOnly
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) GetUri() *string {
	return s.Uri
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) GetVersion() *int32 {
	return s.Version
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) SetExtOptions(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets {
	s.ExtOptions = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) SetIdentifier(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets {
	s.Identifier = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) SetMountPath(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets {
	s.MountPath = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) SetReadOnly(v bool) *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets {
	s.ReadOnly = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) SetUri(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets {
	s.Uri = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) SetVersion(v int32) *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets {
	s.Version = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesDatasets) Validate() error {
	return dara.Validate(s)
}

type ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc struct {
	// The list of port forwarding configurations.
	ForwardInfos []*ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos `json:"ForwardInfos,omitempty" xml:"ForwardInfos,omitempty" type:"Repeated"`
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

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) GetForwardInfos() []*ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos {
	return s.ForwardInfos
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) SetForwardInfos(v []*ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc {
	s.ForwardInfos = v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) SetSecurityGroupId(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) SetVSwitchId(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) SetVpcId(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc {
	s.VpcId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpc) Validate() error {
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

type ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos struct {
	// The list of access types.
	AccessType []*string `json:"AccessType,omitempty" xml:"AccessType,omitempty" type:"Repeated"`
	// The name of the target container.
	//
	// example:
	//
	// dsw-notebook
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The instance ID of the elastic IP address (EIP).
	//
	// example:
	//
	// eip-bp1****
	EipAllocationId *string `json:"EipAllocationId,omitempty" xml:"EipAllocationId,omitempty"`
	// Indicates whether the port forwarding configuration is enabled.
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

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) GoString() string {
	return s.String()
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) GetAccessType() []*string {
	return s.AccessType
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) GetContainerName() *string {
	return s.ContainerName
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) GetEipAllocationId() *string {
	return s.EipAllocationId
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) GetEnable() *bool {
	return s.Enable
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) GetExternalPort() *string {
	return s.ExternalPort
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) GetForwardPort() *string {
	return s.ForwardPort
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) GetName() *string {
	return s.Name
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) GetSSHPublicKey() *string {
	return s.SSHPublicKey
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) SetAccessType(v []*string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos {
	s.AccessType = v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) SetContainerName(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos {
	s.ContainerName = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) SetEipAllocationId(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos {
	s.EipAllocationId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) SetEnable(v bool) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos {
	s.Enable = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) SetExternalPort(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos {
	s.ExternalPort = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) SetForwardPort(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos {
	s.ForwardPort = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) SetName(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos {
	s.Name = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) SetNatGatewayId(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos {
	s.NatGatewayId = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) SetSSHPublicKey(v string) *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos {
	s.SSHPublicKey = &v
	return s
}

func (s *ListServerIdeInstancesResponseBodyPagingInfoInstancesUserVpcForwardInfos) Validate() error {
	return dara.Validate(s)
}

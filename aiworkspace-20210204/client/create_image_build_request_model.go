// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageBuildRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateImageBuildRequest
	GetClientToken() *string
	SetAccessibility(v string) *CreateImageBuildRequest
	GetAccessibility() *string
	SetBuildConfig(v *CreateImageBuildRequestBuildConfig) *CreateImageBuildRequest
	GetBuildConfig() *CreateImageBuildRequestBuildConfig
	SetImage(v *CreateImageBuildRequestImage) *CreateImageBuildRequest
	GetImage() *CreateImageBuildRequestImage
	SetImageBuildJobName(v string) *CreateImageBuildRequest
	GetImageBuildJobName() *string
	SetOverwriteImageTag(v bool) *CreateImageBuildRequest
	GetOverwriteImageTag() *bool
	SetRegionId(v string) *CreateImageBuildRequest
	GetRegionId() *string
	SetResource(v *CreateImageBuildRequestResource) *CreateImageBuildRequest
	GetResource() *CreateImageBuildRequestResource
	SetTargetRegistry(v *CreateImageBuildRequestTargetRegistry) *CreateImageBuildRequest
	GetTargetRegistry() *CreateImageBuildRequestTargetRegistry
	SetUserVpc(v *CreateImageBuildRequestUserVpc) *CreateImageBuildRequest
	GetUserVpc() *CreateImageBuildRequestUserVpc
	SetWorkspaceId(v string) *CreateImageBuildRequest
	GetWorkspaceId() *string
}

type CreateImageBuildRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 镜像构建的可见性，可能值： - PUBLIC：当前工作空间所有成员都可以操作。 - PRIVATE：只有创建者可以操作。
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// 构建配置，指定待构建的 Dockerfile 文件内容。
	//
	// This parameter is required.
	BuildConfig *CreateImageBuildRequestBuildConfig `json:"BuildConfig,omitempty" xml:"BuildConfig,omitempty" type:"Struct"`
	// This parameter is required.
	Image *CreateImageBuildRequestImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// example:
	//
	// build-my-image
	ImageBuildJobName *string `json:"ImageBuildJobName,omitempty" xml:"ImageBuildJobName,omitempty"`
	// 是否覆盖更新 ACR 镜像仓库中已存在的镜像 tag。
	OverwriteImageTag *bool `json:"OverwriteImageTag,omitempty" xml:"OverwriteImageTag,omitempty"`
	// 代表region的资源属性字段
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 构建任务运行资源
	//
	// This parameter is required.
	Resource *CreateImageBuildRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// This parameter is required.
	TargetRegistry *CreateImageBuildRequestTargetRegistry `json:"TargetRegistry,omitempty" xml:"TargetRegistry,omitempty" type:"Struct"`
	// 用户专有网络信息。使用企业版 ACR 实例时，此参数必填，指定在用户 ACR 实例的访问控制里已添加的专有网络。
	UserVpc *CreateImageBuildRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// 镜像构建所属的工作空间ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateImageBuildRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageBuildRequest) GoString() string {
	return s.String()
}

func (s *CreateImageBuildRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageBuildRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateImageBuildRequest) GetBuildConfig() *CreateImageBuildRequestBuildConfig {
	return s.BuildConfig
}

func (s *CreateImageBuildRequest) GetImage() *CreateImageBuildRequestImage {
	return s.Image
}

func (s *CreateImageBuildRequest) GetImageBuildJobName() *string {
	return s.ImageBuildJobName
}

func (s *CreateImageBuildRequest) GetOverwriteImageTag() *bool {
	return s.OverwriteImageTag
}

func (s *CreateImageBuildRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageBuildRequest) GetResource() *CreateImageBuildRequestResource {
	return s.Resource
}

func (s *CreateImageBuildRequest) GetTargetRegistry() *CreateImageBuildRequestTargetRegistry {
	return s.TargetRegistry
}

func (s *CreateImageBuildRequest) GetUserVpc() *CreateImageBuildRequestUserVpc {
	return s.UserVpc
}

func (s *CreateImageBuildRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateImageBuildRequest) SetClientToken(v string) *CreateImageBuildRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageBuildRequest) SetAccessibility(v string) *CreateImageBuildRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateImageBuildRequest) SetBuildConfig(v *CreateImageBuildRequestBuildConfig) *CreateImageBuildRequest {
	s.BuildConfig = v
	return s
}

func (s *CreateImageBuildRequest) SetImage(v *CreateImageBuildRequestImage) *CreateImageBuildRequest {
	s.Image = v
	return s
}

func (s *CreateImageBuildRequest) SetImageBuildJobName(v string) *CreateImageBuildRequest {
	s.ImageBuildJobName = &v
	return s
}

func (s *CreateImageBuildRequest) SetOverwriteImageTag(v bool) *CreateImageBuildRequest {
	s.OverwriteImageTag = &v
	return s
}

func (s *CreateImageBuildRequest) SetRegionId(v string) *CreateImageBuildRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageBuildRequest) SetResource(v *CreateImageBuildRequestResource) *CreateImageBuildRequest {
	s.Resource = v
	return s
}

func (s *CreateImageBuildRequest) SetTargetRegistry(v *CreateImageBuildRequestTargetRegistry) *CreateImageBuildRequest {
	s.TargetRegistry = v
	return s
}

func (s *CreateImageBuildRequest) SetUserVpc(v *CreateImageBuildRequestUserVpc) *CreateImageBuildRequest {
	s.UserVpc = v
	return s
}

func (s *CreateImageBuildRequest) SetWorkspaceId(v string) *CreateImageBuildRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateImageBuildRequest) Validate() error {
	if s.BuildConfig != nil {
		if err := s.BuildConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	if s.TargetRegistry != nil {
		if err := s.TargetRegistry.Validate(); err != nil {
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

type CreateImageBuildRequestBuildConfig struct {
	// This parameter is required.
	//
	// example:
	//
	// PackageInstallation
	BuildType *string `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	// Dockerfile文件内容
	//
	// This parameter is required.
	//
	// example:
	//
	// FROM ubuntu:18:04
	//
	// RUN pip3 install numpy==1.19.5
	Dockerfile *string `json:"Dockerfile,omitempty" xml:"Dockerfile,omitempty"`
	// example:
	//
	// {
	//
	//   "user-test-registry-vpc.cn-wulanchabu.cr.aliyuncs.com": {
	//
	//     "Auth": "dXNlcjp0ZXN0"
	//
	//   }
	//
	// }
	RegistryAuths map[string]interface{} `json:"RegistryAuths,omitempty" xml:"RegistryAuths,omitempty"`
}

func (s CreateImageBuildRequestBuildConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateImageBuildRequestBuildConfig) GoString() string {
	return s.String()
}

func (s *CreateImageBuildRequestBuildConfig) GetBuildType() *string {
	return s.BuildType
}

func (s *CreateImageBuildRequestBuildConfig) GetDockerfile() *string {
	return s.Dockerfile
}

func (s *CreateImageBuildRequestBuildConfig) GetRegistryAuths() map[string]interface{} {
	return s.RegistryAuths
}

func (s *CreateImageBuildRequestBuildConfig) SetBuildType(v string) *CreateImageBuildRequestBuildConfig {
	s.BuildType = &v
	return s
}

func (s *CreateImageBuildRequestBuildConfig) SetDockerfile(v string) *CreateImageBuildRequestBuildConfig {
	s.Dockerfile = &v
	return s
}

func (s *CreateImageBuildRequestBuildConfig) SetRegistryAuths(v map[string]interface{}) *CreateImageBuildRequestBuildConfig {
	s.RegistryAuths = v
	return s
}

func (s *CreateImageBuildRequestBuildConfig) Validate() error {
	return dara.Validate(s)
}

type CreateImageBuildRequestImage struct {
	Description *string                               `json:"Description,omitempty" xml:"Description,omitempty"`
	Labels      []*CreateImageBuildRequestImageLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test-v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user-test-registry-vpc.cn-wulanchabu.cr.aliyuncs.com/pai-test/pai-test:test-v1
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateImageBuildRequestImage) String() string {
	return dara.Prettify(s)
}

func (s CreateImageBuildRequestImage) GoString() string {
	return s.String()
}

func (s *CreateImageBuildRequestImage) GetDescription() *string {
	return s.Description
}

func (s *CreateImageBuildRequestImage) GetLabels() []*CreateImageBuildRequestImageLabels {
	return s.Labels
}

func (s *CreateImageBuildRequestImage) GetName() *string {
	return s.Name
}

func (s *CreateImageBuildRequestImage) GetUri() *string {
	return s.Uri
}

func (s *CreateImageBuildRequestImage) SetDescription(v string) *CreateImageBuildRequestImage {
	s.Description = &v
	return s
}

func (s *CreateImageBuildRequestImage) SetLabels(v []*CreateImageBuildRequestImageLabels) *CreateImageBuildRequestImage {
	s.Labels = v
	return s
}

func (s *CreateImageBuildRequestImage) SetName(v string) *CreateImageBuildRequestImage {
	s.Name = &v
	return s
}

func (s *CreateImageBuildRequestImage) SetUri(v string) *CreateImageBuildRequestImage {
	s.Uri = &v
	return s
}

func (s *CreateImageBuildRequestImage) Validate() error {
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateImageBuildRequestImageLabels struct {
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImageBuildRequestImageLabels) String() string {
	return dara.Prettify(s)
}

func (s CreateImageBuildRequestImageLabels) GoString() string {
	return s.String()
}

func (s *CreateImageBuildRequestImageLabels) GetKey() *string {
	return s.Key
}

func (s *CreateImageBuildRequestImageLabels) GetValue() *string {
	return s.Value
}

func (s *CreateImageBuildRequestImageLabels) SetKey(v string) *CreateImageBuildRequestImageLabels {
	s.Key = &v
	return s
}

func (s *CreateImageBuildRequestImageLabels) SetValue(v string) *CreateImageBuildRequestImageLabels {
	s.Value = &v
	return s
}

func (s *CreateImageBuildRequestImageLabels) Validate() error {
	return dara.Validate(s)
}

type CreateImageBuildRequestResource struct {
	// 后付费资源规格
	//
	// example:
	//
	// ecs.c6.large
	EcsSpec        *string                                        `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	ResourceConfig *CreateImageBuildRequestResourceResourceConfig `json:"ResourceConfig,omitempty" xml:"ResourceConfig,omitempty" type:"Struct"`
	// example:
	//
	// quotaadzoqup693z
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// Lingjun
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateImageBuildRequestResource) String() string {
	return dara.Prettify(s)
}

func (s CreateImageBuildRequestResource) GoString() string {
	return s.String()
}

func (s *CreateImageBuildRequestResource) GetEcsSpec() *string {
	return s.EcsSpec
}

func (s *CreateImageBuildRequestResource) GetResourceConfig() *CreateImageBuildRequestResourceResourceConfig {
	return s.ResourceConfig
}

func (s *CreateImageBuildRequestResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateImageBuildRequestResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateImageBuildRequestResource) SetEcsSpec(v string) *CreateImageBuildRequestResource {
	s.EcsSpec = &v
	return s
}

func (s *CreateImageBuildRequestResource) SetResourceConfig(v *CreateImageBuildRequestResourceResourceConfig) *CreateImageBuildRequestResource {
	s.ResourceConfig = v
	return s
}

func (s *CreateImageBuildRequestResource) SetResourceId(v string) *CreateImageBuildRequestResource {
	s.ResourceId = &v
	return s
}

func (s *CreateImageBuildRequestResource) SetResourceType(v string) *CreateImageBuildRequestResource {
	s.ResourceType = &v
	return s
}

func (s *CreateImageBuildRequestResource) Validate() error {
	if s.ResourceConfig != nil {
		if err := s.ResourceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageBuildRequestResourceResourceConfig struct {
	// example:
	//
	// 4
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// 8Gi
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s CreateImageBuildRequestResourceResourceConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateImageBuildRequestResourceResourceConfig) GoString() string {
	return s.String()
}

func (s *CreateImageBuildRequestResourceResourceConfig) GetCPU() *string {
	return s.CPU
}

func (s *CreateImageBuildRequestResourceResourceConfig) GetMemory() *string {
	return s.Memory
}

func (s *CreateImageBuildRequestResourceResourceConfig) SetCPU(v string) *CreateImageBuildRequestResourceResourceConfig {
	s.CPU = &v
	return s
}

func (s *CreateImageBuildRequestResourceResourceConfig) SetMemory(v string) *CreateImageBuildRequestResourceResourceConfig {
	s.Memory = &v
	return s
}

func (s *CreateImageBuildRequestResourceResourceConfig) Validate() error {
	return dara.Validate(s)
}

type CreateImageBuildRequestTargetRegistry struct {
	// example:
	//
	// cri-**abcd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AcrEnterprise
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateImageBuildRequestTargetRegistry) String() string {
	return dara.Prettify(s)
}

func (s CreateImageBuildRequestTargetRegistry) GoString() string {
	return s.String()
}

func (s *CreateImageBuildRequestTargetRegistry) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateImageBuildRequestTargetRegistry) GetType() *string {
	return s.Type
}

func (s *CreateImageBuildRequestTargetRegistry) SetInstanceId(v string) *CreateImageBuildRequestTargetRegistry {
	s.InstanceId = &v
	return s
}

func (s *CreateImageBuildRequestTargetRegistry) SetType(v string) *CreateImageBuildRequestTargetRegistry {
	s.Type = &v
	return s
}

func (s *CreateImageBuildRequestTargetRegistry) Validate() error {
	return dara.Validate(s)
}

type CreateImageBuildRequestUserVpc struct {
	// 默认路由网卡出口
	//
	// example:
	//
	// eth1
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// 扩展网段
	ExtendedCidrs []*string `json:"ExtendedCidrs,omitempty" xml:"ExtendedCidrs,omitempty" type:"Repeated"`
	// 安全组 ID
	//
	// example:
	//
	// sg-abcdef**
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 交换机 ID
	//
	// example:
	//
	// vs-abcdef**
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// 专有网络 ID
	//
	// example:
	//
	// vpc-abcdef**
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateImageBuildRequestUserVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateImageBuildRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateImageBuildRequestUserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *CreateImageBuildRequestUserVpc) GetExtendedCidrs() []*string {
	return s.ExtendedCidrs
}

func (s *CreateImageBuildRequestUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateImageBuildRequestUserVpc) GetSwitchId() *string {
	return s.SwitchId
}

func (s *CreateImageBuildRequestUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateImageBuildRequestUserVpc) SetDefaultRoute(v string) *CreateImageBuildRequestUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *CreateImageBuildRequestUserVpc) SetExtendedCidrs(v []*string) *CreateImageBuildRequestUserVpc {
	s.ExtendedCidrs = v
	return s
}

func (s *CreateImageBuildRequestUserVpc) SetSecurityGroupId(v string) *CreateImageBuildRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateImageBuildRequestUserVpc) SetSwitchId(v string) *CreateImageBuildRequestUserVpc {
	s.SwitchId = &v
	return s
}

func (s *CreateImageBuildRequestUserVpc) SetVpcId(v string) *CreateImageBuildRequestUserVpc {
	s.VpcId = &v
	return s
}

func (s *CreateImageBuildRequestUserVpc) Validate() error {
	return dara.Validate(s)
}

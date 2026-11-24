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
	// The idempotency parameter.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The visibility. Valid values:
	//
	// - **PUBLIC**: Public.
	//
	// - **PRIVATE**: Private.
	//
	// example:
	//
	// PUBLIC
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// **The build configuration, which specifies the content of the Dockerfile to be built.**
	//
	// This parameter is required.
	BuildConfig *CreateImageBuildRequestBuildConfig `json:"BuildConfig,omitempty" xml:"BuildConfig,omitempty" type:"Struct"`
	// The image metadata.
	//
	// This parameter is required.
	Image *CreateImageBuildRequestImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// The name of the image build task.
	//
	// example:
	//
	// build-my-image
	ImageBuildJobName *string `json:"ImageBuildJobName,omitempty" xml:"ImageBuildJobName,omitempty"`
	// Specifies whether to overwrite the existing image version in the image repository.
	OverwriteImageTag *bool `json:"OverwriteImageTag,omitempty" xml:"OverwriteImageTag,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resources used to run the task.
	//
	// This parameter is required.
	Resource *CreateImageBuildRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// **The target registry configuration.**
	//
	// This parameter is required.
	TargetRegistry *CreateImageBuildRequestTargetRegistry `json:"TargetRegistry,omitempty" xml:"TargetRegistry,omitempty" type:"Struct"`
	// The user VPC information. This parameter is required when you use a public resource group.
	UserVpc *CreateImageBuildRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// The workspace ID.
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
	// The build type. Valid values:
	//
	// 	- **PackageInstallation**: Installs software packages based on a specified image.
	//
	// 	- **CustomDockerfile**: Uses a custom Dockerfile.
	//
	// This parameter is required.
	//
	// example:
	//
	// PackageInstallation
	BuildType *string `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	// The list of data sources.
	DataSources []*CreateImageBuildRequestBuildConfigDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The content of the Dockerfile to be built.
	//
	// example:
	//
	// FROM ubuntu:18:04
	//
	// RUN pip3 install numpy==1.19.5
	Dockerfile *string `json:"Dockerfile,omitempty" xml:"Dockerfile,omitempty"`
	// The Dockerfile path, which is a relative path to the mounted context path. Example: docker/Dockerfile. If the mount path is /mnt/data/, the file is located at /mnt/data/docker/Dockerfile.
	//
	// example:
	//
	// docker/Dockerfile
	DockerfilePath *string `json:"DockerfilePath,omitempty" xml:"DockerfilePath,omitempty"`
	// The authentication information for private image repositories. You can specify the authentication information for an ACR image repository that does not belong to the current user. Format: {"user_registry_domain":{"Auth":"base64 encoded auth"}}.
	//
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

func (s *CreateImageBuildRequestBuildConfig) GetDataSources() []*CreateImageBuildRequestBuildConfigDataSources {
	return s.DataSources
}

func (s *CreateImageBuildRequestBuildConfig) GetDockerfile() *string {
	return s.Dockerfile
}

func (s *CreateImageBuildRequestBuildConfig) GetDockerfilePath() *string {
	return s.DockerfilePath
}

func (s *CreateImageBuildRequestBuildConfig) GetRegistryAuths() map[string]interface{} {
	return s.RegistryAuths
}

func (s *CreateImageBuildRequestBuildConfig) SetBuildType(v string) *CreateImageBuildRequestBuildConfig {
	s.BuildType = &v
	return s
}

func (s *CreateImageBuildRequestBuildConfig) SetDataSources(v []*CreateImageBuildRequestBuildConfigDataSources) *CreateImageBuildRequestBuildConfig {
	s.DataSources = v
	return s
}

func (s *CreateImageBuildRequestBuildConfig) SetDockerfile(v string) *CreateImageBuildRequestBuildConfig {
	s.Dockerfile = &v
	return s
}

func (s *CreateImageBuildRequestBuildConfig) SetDockerfilePath(v string) *CreateImageBuildRequestBuildConfig {
	s.DockerfilePath = &v
	return s
}

func (s *CreateImageBuildRequestBuildConfig) SetRegistryAuths(v map[string]interface{}) *CreateImageBuildRequestBuildConfig {
	s.RegistryAuths = v
	return s
}

func (s *CreateImageBuildRequestBuildConfig) Validate() error {
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateImageBuildRequestBuildConfigDataSources struct {
	// The ID of the dataset in the workspace. Example: d-************zbs.
	//
	// example:
	//
	// d-************zbs
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The dataset version. If you specify this parameter, you must also specify DataSourceId. Example: v1.
	//
	// example:
	//
	// v1
	DataSourceVersion *string `json:"DataSourceVersion,omitempty" xml:"DataSourceVersion,omitempty"`
	// The read and write permissions. Valid values: RW (default) and RO.
	//
	// example:
	//
	// RO
	MountAccess *string `json:"MountAccess,omitempty" xml:"MountAccess,omitempty"`
	// The mount path. Example: /mnt/data/.
	//
	// example:
	//
	// /mnt/data/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The mount property for the custom dataset. Only OSS is supported. Example: {"fs.oss.download.thread.concurrency":"10"}.
	//
	// example:
	//
	// {"fs.oss.download.thread.concurrency":"10"}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The storage path, including the endpoint. This parameter is mutually exclusive with DataSourceId. Example: OSS storage path oss://{bucket-name}.oss-{region}-internal.aliyuncs.com/{subPath}/.
	//
	// example:
	//
	// oss://{bucket-name}.oss-{region}-internal.aliyuncs.com/{subPath}/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateImageBuildRequestBuildConfigDataSources) String() string {
	return dara.Prettify(s)
}

func (s CreateImageBuildRequestBuildConfigDataSources) GoString() string {
	return s.String()
}

func (s *CreateImageBuildRequestBuildConfigDataSources) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateImageBuildRequestBuildConfigDataSources) GetDataSourceVersion() *string {
	return s.DataSourceVersion
}

func (s *CreateImageBuildRequestBuildConfigDataSources) GetMountAccess() *string {
	return s.MountAccess
}

func (s *CreateImageBuildRequestBuildConfigDataSources) GetMountPath() *string {
	return s.MountPath
}

func (s *CreateImageBuildRequestBuildConfigDataSources) GetOptions() *string {
	return s.Options
}

func (s *CreateImageBuildRequestBuildConfigDataSources) GetUri() *string {
	return s.Uri
}

func (s *CreateImageBuildRequestBuildConfigDataSources) SetDataSourceId(v string) *CreateImageBuildRequestBuildConfigDataSources {
	s.DataSourceId = &v
	return s
}

func (s *CreateImageBuildRequestBuildConfigDataSources) SetDataSourceVersion(v string) *CreateImageBuildRequestBuildConfigDataSources {
	s.DataSourceVersion = &v
	return s
}

func (s *CreateImageBuildRequestBuildConfigDataSources) SetMountAccess(v string) *CreateImageBuildRequestBuildConfigDataSources {
	s.MountAccess = &v
	return s
}

func (s *CreateImageBuildRequestBuildConfigDataSources) SetMountPath(v string) *CreateImageBuildRequestBuildConfigDataSources {
	s.MountPath = &v
	return s
}

func (s *CreateImageBuildRequestBuildConfigDataSources) SetOptions(v string) *CreateImageBuildRequestBuildConfigDataSources {
	s.Options = &v
	return s
}

func (s *CreateImageBuildRequestBuildConfigDataSources) SetUri(v string) *CreateImageBuildRequestBuildConfigDataSources {
	s.Uri = &v
	return s
}

func (s *CreateImageBuildRequestBuildConfigDataSources) Validate() error {
	return dara.Validate(s)
}

type CreateImageBuildRequestImage struct {
	// The image description.
	//
	// example:
	//
	// Build test image
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image labels.
	Labels []*CreateImageBuildRequestImageLabels `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The image name. The naming rules are as follows:
	//
	// 	- The name must be 1 to 50 characters in length.
	//
	// 	- The name can contain lowercase letters, digits, and hyphens (-), and must start with a letter.
	//
	// 	- The name must be unique within the same workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The image URI.
	//
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
	// The key of the image label.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the image label.
	//
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
	// The pay-as-you-go resource specifications. This parameter is required when you use a public resource group.
	//
	// example:
	//
	// ecs.c6.large
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// The resource configuration. This parameter is required when you use subscription resources and must be left empty when you use a public resource group.
	ResourceConfig *CreateImageBuildRequestResourceResourceConfig `json:"ResourceConfig,omitempty" xml:"ResourceConfig,omitempty" type:"Struct"`
	// The resource quota ID. This parameter applies only to subscription resources and must not be specified when you use pay-as-you-go resources.
	//
	// example:
	//
	// quotaadzoqup693z
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The subscription resource type. Only Lingjun resources are supported. This parameter is required when you use subscription resources.
	//
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
	// The number of CPU cores.
	//
	// example:
	//
	// 4
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The memory capacity.
	//
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
	// The registry instance ID. This parameter is required when you use Container Registry (ACR) as the image service.
	//
	// example:
	//
	// cri-**abcd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The registry type. Only Container Registry Enterprise Edition is supported as the target registry.
	//
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
	// The default routing of the network interface controller (NIC). Valid values:
	//
	// 	- eth1: Uses the user\\"s elastic network interfaces (ENIs) to access external networks through a private NAT gateway. For the configuration method, refer to [Configure a DSW instance to access the Internet through a private NAT gateway](https://www.alibabacloud.com/help/en/pai/user-guide/configure-a-dsw-instance-to-access-the-internet-through-a-private-nat-gateway).
	//
	// example:
	//
	// eth1
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// The extended CIDR blocks.
	//
	// 	- If the vSwitch ID is empty, this parameter is optional and the system automatically retrieves all CIDR blocks under the VPC.
	//
	// 	- If the vSwitch ID is not empty, this parameter is required. We recommend that you specify all CIDR blocks under the VPC.
	ExtendedCidrs []*string `json:"ExtendedCidrs,omitempty" xml:"ExtendedCidrs,omitempty" type:"Repeated"`
	// The security group ID. This parameter is required when you specify the VPC configuration.
	//
	// example:
	//
	// sg-abcdef**
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID. This parameter is optional.
	//
	// example:
	//
	// vs-abcdef**
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The VPC ID. When the build task accesses a user\\"s Container Registry Enterprise Edition instance, you must specify a VPC that is added to the access control list of the instance.
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

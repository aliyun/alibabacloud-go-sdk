// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactBuildProperty(v *UpdateArtifactRequestArtifactBuildProperty) *UpdateArtifactRequest
	GetArtifactBuildProperty() *UpdateArtifactRequestArtifactBuildProperty
	SetArtifactId(v string) *UpdateArtifactRequest
	GetArtifactId() *string
	SetArtifactProperty(v *UpdateArtifactRequestArtifactProperty) *UpdateArtifactRequest
	GetArtifactProperty() *UpdateArtifactRequestArtifactProperty
	SetClientToken(v string) *UpdateArtifactRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateArtifactRequest
	GetDescription() *string
	SetPermissionType(v string) *UpdateArtifactRequest
	GetPermissionType() *string
	SetSupportRegionIds(v []*string) *UpdateArtifactRequest
	GetSupportRegionIds() []*string
	SetVersionName(v string) *UpdateArtifactRequest
	GetVersionName() *string
}

type UpdateArtifactRequest struct {
	// The properties for building the artifact. This is used for managed artifact builds.
	ArtifactBuildProperty *UpdateArtifactRequestArtifactBuildProperty `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty" type:"Struct"`
	// The ID of the artifact.
	//
	// To obtain the artifact ID, call the [ListArtifacts](https://help.aliyun.com/document_detail/469993.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43ae****
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the artifact.
	ArtifactProperty *UpdateArtifactRequestArtifactProperty `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty" type:"Struct"`
	// A client token to ensure the idempotence of the request. Generate a unique token for each request from your client. The **ClientToken*	- can contain only ASCII characters and must be no more than 64 characters long.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the artifact.
	//
	// example:
	//
	// Redhat8_0 image
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The permission type. This parameter is valid for container image artifacts and Helm Chart artifacts. The value can be changed only from \\`Automatic\\` to \\`Public\\`. Valid values:
	//
	// - Public
	//
	// - Automatic
	//
	// example:
	//
	// Public
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The IDs of regions to which the image can be distributed.
	SupportRegionIds []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// The name of the artifact version.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequest) GetArtifactBuildProperty() *UpdateArtifactRequestArtifactBuildProperty {
	return s.ArtifactBuildProperty
}

func (s *UpdateArtifactRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *UpdateArtifactRequest) GetArtifactProperty() *UpdateArtifactRequestArtifactProperty {
	return s.ArtifactProperty
}

func (s *UpdateArtifactRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateArtifactRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateArtifactRequest) GetPermissionType() *string {
	return s.PermissionType
}

func (s *UpdateArtifactRequest) GetSupportRegionIds() []*string {
	return s.SupportRegionIds
}

func (s *UpdateArtifactRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *UpdateArtifactRequest) SetArtifactBuildProperty(v *UpdateArtifactRequestArtifactBuildProperty) *UpdateArtifactRequest {
	s.ArtifactBuildProperty = v
	return s
}

func (s *UpdateArtifactRequest) SetArtifactId(v string) *UpdateArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *UpdateArtifactRequest) SetArtifactProperty(v *UpdateArtifactRequestArtifactProperty) *UpdateArtifactRequest {
	s.ArtifactProperty = v
	return s
}

func (s *UpdateArtifactRequest) SetClientToken(v string) *UpdateArtifactRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateArtifactRequest) SetDescription(v string) *UpdateArtifactRequest {
	s.Description = &v
	return s
}

func (s *UpdateArtifactRequest) SetPermissionType(v string) *UpdateArtifactRequest {
	s.PermissionType = &v
	return s
}

func (s *UpdateArtifactRequest) SetSupportRegionIds(v []*string) *UpdateArtifactRequest {
	s.SupportRegionIds = v
	return s
}

func (s *UpdateArtifactRequest) SetVersionName(v string) *UpdateArtifactRequest {
	s.VersionName = &v
	return s
}

func (s *UpdateArtifactRequest) Validate() error {
	if s.ArtifactBuildProperty != nil {
		if err := s.ArtifactBuildProperty.Validate(); err != nil {
			return err
		}
	}
	if s.ArtifactProperty != nil {
		if err := s.ArtifactProperty.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateArtifactRequestArtifactBuildProperty struct {
	// The build arguments.
	//
	// > This parameter is available only when \\`ArtifactBuildType\\` is set to \\`Dockerfile\\`.
	BuildArgs []*UpdateArtifactRequestArtifactBuildPropertyBuildArgs `json:"BuildArgs,omitempty" xml:"BuildArgs,omitempty" type:"Repeated"`
	// The code repository address.
	//
	// > This parameter is available only when \\`ArtifactBuildType\\` is set to \\`Dockerfile\\` or \\`Buildpacks\\`.
	CodeRepo *UpdateArtifactRequestArtifactBuildPropertyCodeRepo `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty" type:"Struct"`
	// The content of the command.
	//
	// > This parameter is available only for ECS image artifacts.
	//
	// example:
	//
	// echo "start run command"
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command type. Valid values:
	//
	// - RunBatScript: The command is a batch script that runs on a Windows instance.
	//
	// - RunPowerShellScript: The command is a PowerShell script that runs on a Windows instance.
	//
	// - RunShellScript: The command is a shell script that runs on a Linux instance.
	//
	// > This parameter is available only for ECS image artifacts.
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The relative path of the Dockerfile in the code repository.
	//
	// Default value: Dockerfile
	//
	// > This parameter is available only when \\`ArtifactBuildType\\` is set to \\`Dockerfile\\`.
	//
	// example:
	//
	// ./file/Dockerfile
	DockerfilePath *string `json:"DockerfilePath,omitempty" xml:"DockerfilePath,omitempty"`
	// Specifies whether to use a GPU-accelerated instance for the build. By default, a CPU instance is used.
	//
	// example:
	//
	// false
	EnableGpu *bool `json:"EnableGpu,omitempty" xml:"EnableGpu,omitempty"`
	// The ID of the region where the source image is located.
	//
	// > This parameter is available only for ECS image artifacts.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The pull URL of the source container image.
	//
	// Used for \\`docker pull ${SourceContainerImage}\\`.
	//
	// > This parameter is available only when \\`ArtifactBuildType\\` is set to \\`ContainerImage\\`.
	//
	// example:
	//
	// pytorch/pytorch:2.5.1-cuda12.4-cudnn9-devel
	SourceContainerImage *string `json:"SourceContainerImage,omitempty" xml:"SourceContainerImage,omitempty"`
	// The source image ID. The following types are supported:
	//
	// - Image ID: The ID of the ECS image.
	//
	// - OOS common parameter name: The system automatically obtains the corresponding image ID based on the OOS common parameter name.
	//
	// > This parameter is available only for ECS image artifacts.
	//
	// example:
	//
	// Image ID: m-t4nhenrdc38pe4*****
	//
	// ubuntu_22_04_x64_20G_alibase_20240926.vhd
	//
	// OOS public parameter name: aliyun/services/computenest/images/aliyun_3_2104_python_3_11
	SourceImageId *string `json:"SourceImageId,omitempty" xml:"SourceImageId,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// example:
	//
	// 40
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s UpdateArtifactRequestArtifactBuildProperty) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactRequestArtifactBuildProperty) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequestArtifactBuildProperty) GetBuildArgs() []*UpdateArtifactRequestArtifactBuildPropertyBuildArgs {
	return s.BuildArgs
}

func (s *UpdateArtifactRequestArtifactBuildProperty) GetCodeRepo() *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	return s.CodeRepo
}

func (s *UpdateArtifactRequestArtifactBuildProperty) GetCommandContent() *string {
	return s.CommandContent
}

func (s *UpdateArtifactRequestArtifactBuildProperty) GetCommandType() *string {
	return s.CommandType
}

func (s *UpdateArtifactRequestArtifactBuildProperty) GetDockerfilePath() *string {
	return s.DockerfilePath
}

func (s *UpdateArtifactRequestArtifactBuildProperty) GetEnableGpu() *bool {
	return s.EnableGpu
}

func (s *UpdateArtifactRequestArtifactBuildProperty) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateArtifactRequestArtifactBuildProperty) GetSourceContainerImage() *string {
	return s.SourceContainerImage
}

func (s *UpdateArtifactRequestArtifactBuildProperty) GetSourceImageId() *string {
	return s.SourceImageId
}

func (s *UpdateArtifactRequestArtifactBuildProperty) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetBuildArgs(v []*UpdateArtifactRequestArtifactBuildPropertyBuildArgs) *UpdateArtifactRequestArtifactBuildProperty {
	s.BuildArgs = v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetCodeRepo(v *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) *UpdateArtifactRequestArtifactBuildProperty {
	s.CodeRepo = v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetCommandContent(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.CommandContent = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetCommandType(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.CommandType = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetDockerfilePath(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.DockerfilePath = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetEnableGpu(v bool) *UpdateArtifactRequestArtifactBuildProperty {
	s.EnableGpu = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetRegionId(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.RegionId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetSourceContainerImage(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.SourceContainerImage = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetSourceImageId(v string) *UpdateArtifactRequestArtifactBuildProperty {
	s.SourceImageId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) SetSystemDiskSize(v int64) *UpdateArtifactRequestArtifactBuildProperty {
	s.SystemDiskSize = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildProperty) Validate() error {
	if s.BuildArgs != nil {
		for _, item := range s.BuildArgs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CodeRepo != nil {
		if err := s.CodeRepo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateArtifactRequestArtifactBuildPropertyBuildArgs struct {
	// The name of the build argument.
	//
	// example:
	//
	// ENV
	ArgumentName *string `json:"ArgumentName,omitempty" xml:"ArgumentName,omitempty"`
	// The value of the build argument.
	//
	// example:
	//
	// nginx:latest
	ArgumentValue *string `json:"ArgumentValue,omitempty" xml:"ArgumentValue,omitempty"`
}

func (s UpdateArtifactRequestArtifactBuildPropertyBuildArgs) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactRequestArtifactBuildPropertyBuildArgs) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequestArtifactBuildPropertyBuildArgs) GetArgumentName() *string {
	return s.ArgumentName
}

func (s *UpdateArtifactRequestArtifactBuildPropertyBuildArgs) GetArgumentValue() *string {
	return s.ArgumentValue
}

func (s *UpdateArtifactRequestArtifactBuildPropertyBuildArgs) SetArgumentName(v string) *UpdateArtifactRequestArtifactBuildPropertyBuildArgs {
	s.ArgumentName = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyBuildArgs) SetArgumentValue(v string) *UpdateArtifactRequestArtifactBuildPropertyBuildArgs {
	s.ArgumentValue = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyBuildArgs) Validate() error {
	return dara.Validate(s)
}

type UpdateArtifactRequestArtifactBuildPropertyCodeRepo struct {
	// The branch name of the code repository.
	//
	// example:
	//
	// main
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// The endpoint. This parameter is required for a private GitLab deployment.
	//
	// example:
	//
	// http://121.40.25.0
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// 455231
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The owner of the code repository.
	//
	// > This parameter is required only if the code repository is private.
	//
	// example:
	//
	// aliyun-computenest
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The platform of the code repository. Valid values:
	//
	// - github
	//
	// - gitee
	//
	// - gitlab
	//
	// - codeup
	//
	// example:
	//
	// github
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The repository ID.
	//
	// example:
	//
	// 103
	RepoId *int64 `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The repository name.
	//
	// example:
	//
	// aliyun-computenest/quickstart-Lobexxx
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
}

func (s UpdateArtifactRequestArtifactBuildPropertyCodeRepo) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactRequestArtifactBuildPropertyCodeRepo) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) GetBranch() *string {
	return s.Branch
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) GetEndpoint() *string {
	return s.Endpoint
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) GetOrgId() *string {
	return s.OrgId
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) GetOwner() *string {
	return s.Owner
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) GetPlatform() *string {
	return s.Platform
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) GetRepoId() *int64 {
	return s.RepoId
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) GetRepoName() *string {
	return s.RepoName
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetBranch(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Branch = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetEndpoint(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Endpoint = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetOrgId(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.OrgId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetOwner(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Owner = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetPlatform(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Platform = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetRepoId(v int64) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.RepoId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) SetRepoName(v string) *UpdateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.RepoName = &v
	return s
}

func (s *UpdateArtifactRequestArtifactBuildPropertyCodeRepo) Validate() error {
	return dara.Validate(s)
}

type UpdateArtifactRequestArtifactProperty struct {
	// The code of the Alibaba Cloud Marketplace product.
	//
	// You can obtain the product code in the [Alibaba Cloud Marketplace console](https://market.console.aliyun.com/?spm=a2c4g.11186623.0.0.599d6787eMBBxu#/apiTools?_k=d7j8gk).
	//
	// > This parameter is available only for image artifacts.
	//
	// example:
	//
	// cmjj00****
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The version of the Alibaba Cloud Marketplace product.
	//
	// You can view the product version on the [Alibaba Cloud Marketplace page](https://market.aliyun.com/?spm=5176.24779694.0.0.b2144d22sksKM5).
	//
	// > This parameter is available only for image artifacts.
	//
	// example:
	//
	// V1.0
	CommodityVersion *string `json:"CommodityVersion,omitempty" xml:"CommodityVersion,omitempty"`
	// The ID of the image.
	//
	// After you specify a region ID, call the [DescribeImages](https://help.aliyun.com/document_detail/2679797.html) operation to query available image IDs in that region.
	//
	// > This parameter is available only for image artifacts.
	//
	// example:
	//
	// m-0xij191j9cuev6uc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region of the image.
	//
	// > This parameter is available only for image artifacts.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the image repository.
	//
	// To obtain the image repository ID, call the [ListAcrImageRepositories](https://help.aliyun.com/document_detail/2539919.html) operation.
	//
	// > This parameter is available only for container image artifacts and Helm Chart artifacts.
	//
	// example:
	//
	// crr-d8o1nponyc2t****
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the image repository.
	//
	// > This parameter is available only for container image artifacts and Helm Chart artifacts.
	//
	// example:
	//
	// wordpress
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The permission type of the repository. Valid values:
	//
	// - `Public`: public repository
	//
	// - `Private`: private repository
	//
	// > This parameter is available only for container image artifacts and Helm Chart artifacts.
	//
	// example:
	//
	// Public
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The version tag of the image in the repository.
	//
	// To obtain the version tag, call the [ListAcrImageTags](https://help.aliyun.com/document_detail/2539920.html) operation.
	//
	// > This parameter is available only for container image artifacts and Helm Chart artifacts.
	//
	// example:
	//
	// v1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The URL of the file artifact.
	//
	// You can upload the file and obtain its URL in the [Object Storage Service console](https://oss.console.aliyun.com/bucket).
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s UpdateArtifactRequestArtifactProperty) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactRequestArtifactProperty) GoString() string {
	return s.String()
}

func (s *UpdateArtifactRequestArtifactProperty) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *UpdateArtifactRequestArtifactProperty) GetCommodityVersion() *string {
	return s.CommodityVersion
}

func (s *UpdateArtifactRequestArtifactProperty) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateArtifactRequestArtifactProperty) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateArtifactRequestArtifactProperty) GetRepoId() *string {
	return s.RepoId
}

func (s *UpdateArtifactRequestArtifactProperty) GetRepoName() *string {
	return s.RepoName
}

func (s *UpdateArtifactRequestArtifactProperty) GetRepoType() *string {
	return s.RepoType
}

func (s *UpdateArtifactRequestArtifactProperty) GetTag() *string {
	return s.Tag
}

func (s *UpdateArtifactRequestArtifactProperty) GetUrl() *string {
	return s.Url
}

func (s *UpdateArtifactRequestArtifactProperty) SetCommodityCode(v string) *UpdateArtifactRequestArtifactProperty {
	s.CommodityCode = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetCommodityVersion(v string) *UpdateArtifactRequestArtifactProperty {
	s.CommodityVersion = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetImageId(v string) *UpdateArtifactRequestArtifactProperty {
	s.ImageId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetRegionId(v string) *UpdateArtifactRequestArtifactProperty {
	s.RegionId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetRepoId(v string) *UpdateArtifactRequestArtifactProperty {
	s.RepoId = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetRepoName(v string) *UpdateArtifactRequestArtifactProperty {
	s.RepoName = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetRepoType(v string) *UpdateArtifactRequestArtifactProperty {
	s.RepoType = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetTag(v string) *UpdateArtifactRequestArtifactProperty {
	s.Tag = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) SetUrl(v string) *UpdateArtifactRequestArtifactProperty {
	s.Url = &v
	return s
}

func (s *UpdateArtifactRequestArtifactProperty) Validate() error {
	return dara.Validate(s)
}

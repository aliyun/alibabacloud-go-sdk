// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactBuildProperty(v *CreateArtifactRequestArtifactBuildProperty) *CreateArtifactRequest
	GetArtifactBuildProperty() *CreateArtifactRequestArtifactBuildProperty
	SetArtifactBuildType(v string) *CreateArtifactRequest
	GetArtifactBuildType() *string
	SetArtifactId(v string) *CreateArtifactRequest
	GetArtifactId() *string
	SetArtifactProperty(v *CreateArtifactRequestArtifactProperty) *CreateArtifactRequest
	GetArtifactProperty() *CreateArtifactRequestArtifactProperty
	SetArtifactType(v string) *CreateArtifactRequest
	GetArtifactType() *string
	SetClientToken(v string) *CreateArtifactRequest
	GetClientToken() *string
	SetDescription(v string) *CreateArtifactRequest
	GetDescription() *string
	SetName(v string) *CreateArtifactRequest
	GetName() *string
	SetResourceGroupId(v string) *CreateArtifactRequest
	GetResourceGroupId() *string
	SetSupportRegionIds(v []*string) *CreateArtifactRequest
	GetSupportRegionIds() []*string
	SetTag(v []*CreateArtifactRequestTag) *CreateArtifactRequest
	GetTag() []*CreateArtifactRequestTag
	SetVersionName(v string) *CreateArtifactRequest
	GetVersionName() *string
}

type CreateArtifactRequest struct {
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	ArtifactBuildProperty *CreateArtifactRequestArtifactBuildProperty `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty" type:"Struct"`
	// The type of the artifact build task. Valid values:
	//
	// - EcsImage: Build ECS (Elastic Container Service) image.
	//
	// - Dockerfile: Build container image based on Dockerfile.
	//
	// - Buildpacks: Build container image based on Buildpacks.
	//
	// - ContainerImage: Rebuild container image by renaming an existing container image.
	//
	// example:
	//
	// Dockerfile
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The ID of the deployment package.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment object.
	ArtifactProperty *CreateArtifactRequestArtifactProperty `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty" type:"Struct"`
	// The type of the deployment package. Valid values:
	//
	// 	- EcsImage: Elastic Compute Service (ECS) image.
	//
	// 	- AcrImage: container image.
	//
	// 	- File: Object Storage Service (OSS) object.
	//
	// 	- Script: script.
	//
	// This parameter is required.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 10CM943JP0EN9D51H
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the deployment package.
	//
	// example:
	//
	// Test artifact
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzkt5buxxxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The supported regions.
	SupportRegionIds []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// The custom tags.
	Tag []*CreateArtifactRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The version name of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequest) GetArtifactBuildProperty() *CreateArtifactRequestArtifactBuildProperty {
	return s.ArtifactBuildProperty
}

func (s *CreateArtifactRequest) GetArtifactBuildType() *string {
	return s.ArtifactBuildType
}

func (s *CreateArtifactRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *CreateArtifactRequest) GetArtifactProperty() *CreateArtifactRequestArtifactProperty {
	return s.ArtifactProperty
}

func (s *CreateArtifactRequest) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *CreateArtifactRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateArtifactRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateArtifactRequest) GetName() *string {
	return s.Name
}

func (s *CreateArtifactRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateArtifactRequest) GetSupportRegionIds() []*string {
	return s.SupportRegionIds
}

func (s *CreateArtifactRequest) GetTag() []*CreateArtifactRequestTag {
	return s.Tag
}

func (s *CreateArtifactRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateArtifactRequest) SetArtifactBuildProperty(v *CreateArtifactRequestArtifactBuildProperty) *CreateArtifactRequest {
	s.ArtifactBuildProperty = v
	return s
}

func (s *CreateArtifactRequest) SetArtifactBuildType(v string) *CreateArtifactRequest {
	s.ArtifactBuildType = &v
	return s
}

func (s *CreateArtifactRequest) SetArtifactId(v string) *CreateArtifactRequest {
	s.ArtifactId = &v
	return s
}

func (s *CreateArtifactRequest) SetArtifactProperty(v *CreateArtifactRequestArtifactProperty) *CreateArtifactRequest {
	s.ArtifactProperty = v
	return s
}

func (s *CreateArtifactRequest) SetArtifactType(v string) *CreateArtifactRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactRequest) SetClientToken(v string) *CreateArtifactRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateArtifactRequest) SetDescription(v string) *CreateArtifactRequest {
	s.Description = &v
	return s
}

func (s *CreateArtifactRequest) SetName(v string) *CreateArtifactRequest {
	s.Name = &v
	return s
}

func (s *CreateArtifactRequest) SetResourceGroupId(v string) *CreateArtifactRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateArtifactRequest) SetSupportRegionIds(v []*string) *CreateArtifactRequest {
	s.SupportRegionIds = v
	return s
}

func (s *CreateArtifactRequest) SetTag(v []*CreateArtifactRequestTag) *CreateArtifactRequest {
	s.Tag = v
	return s
}

func (s *CreateArtifactRequest) SetVersionName(v string) *CreateArtifactRequest {
	s.VersionName = &v
	return s
}

func (s *CreateArtifactRequest) Validate() error {
	return dara.Validate(s)
}

type CreateArtifactRequestArtifactBuildProperty struct {
	// The build arguments used during the image build process.
	//
	// >  This parameter is available only if the ArtifactBuildType is Dockerfile type.
	BuildArgs []*CreateArtifactRequestArtifactBuildPropertyBuildArgs `json:"BuildArgs,omitempty" xml:"BuildArgs,omitempty" type:"Repeated"`
	// The address of the code repository.
	//
	// >  This parameter is available only if the ArtifactBuildType is Dockerfile or Buildpacks type.
	CodeRepo *CreateArtifactRequestArtifactBuildPropertyCodeRepo `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty" type:"Struct"`
	// The command content.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// echo "start run command"
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command type. Valid values:
	//
	// 	- RunBatScript: batch command, applicable to Windows instances.
	//
	// 	- RunPowerShellScript: PowerShell command, applicable to Windows instances.
	//
	// 	- RunShellScript: shell command, applicable to Linux instances.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The relative path to the Dockerfile within the code repository.
	//
	// >  This parameter is available only if the ArtifactBuildType is Dockerfile type.
	//
	// example:
	//
	// ./file/Dockerfile
	DockerfilePath *string `json:"DockerfilePath,omitempty" xml:"DockerfilePath,omitempty"`
	// Whether GPU is required. CPU instance is used by default.
	//
	// example:
	//
	// false
	EnableGpu *bool `json:"EnableGpu,omitempty" xml:"EnableGpu,omitempty"`
	// The region ID where the source mirror image is located.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The pull location of the source container image. This is used for the command docker pull ${SourceContainerImage}.
	//
	// >  This parameter is available only if the ArtifactBuildType is ContainerImage type.
	//
	// example:
	//
	// pytorch/pytorch:2.5.1-cuda12.4-cudnn9-devel
	SourceContainerImage *string `json:"SourceContainerImage,omitempty" xml:"SourceContainerImage,omitempty"`
	// The source image id. Supported Types:
	//
	// - Image ID: Pass the Image ID of the Ecs image directly.
	//
	// - OOS Common Parameter Name: Obtain the corresponding Image ID automatically by using the OOS common parameter name.
	//
	// >  This parameter is available only if the deployment package is a ecs image type.
	//
	// example:
	//
	// Image ID：m-t4nhenrdc38pe4*****
	//
	// ubuntu_22_04_x64_20G_alibase_20240926.vhd
	//
	// OOS Common Parameter Name：aliyun/services/computenest/images/aliyun_3_2104_python_3_11
	SourceImageId *string `json:"SourceImageId,omitempty" xml:"SourceImageId,omitempty"`
	// The size of the system disk. Unit: GiB.
	//
	// >  The system disk must be at least as large as the image.
	//
	// example:
	//
	// 40
	SystemDiskSize *int64 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s CreateArtifactRequestArtifactBuildProperty) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactRequestArtifactBuildProperty) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestArtifactBuildProperty) GetBuildArgs() []*CreateArtifactRequestArtifactBuildPropertyBuildArgs {
	return s.BuildArgs
}

func (s *CreateArtifactRequestArtifactBuildProperty) GetCodeRepo() *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	return s.CodeRepo
}

func (s *CreateArtifactRequestArtifactBuildProperty) GetCommandContent() *string {
	return s.CommandContent
}

func (s *CreateArtifactRequestArtifactBuildProperty) GetCommandType() *string {
	return s.CommandType
}

func (s *CreateArtifactRequestArtifactBuildProperty) GetDockerfilePath() *string {
	return s.DockerfilePath
}

func (s *CreateArtifactRequestArtifactBuildProperty) GetEnableGpu() *bool {
	return s.EnableGpu
}

func (s *CreateArtifactRequestArtifactBuildProperty) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateArtifactRequestArtifactBuildProperty) GetSourceContainerImage() *string {
	return s.SourceContainerImage
}

func (s *CreateArtifactRequestArtifactBuildProperty) GetSourceImageId() *string {
	return s.SourceImageId
}

func (s *CreateArtifactRequestArtifactBuildProperty) GetSystemDiskSize() *int64 {
	return s.SystemDiskSize
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetBuildArgs(v []*CreateArtifactRequestArtifactBuildPropertyBuildArgs) *CreateArtifactRequestArtifactBuildProperty {
	s.BuildArgs = v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetCodeRepo(v *CreateArtifactRequestArtifactBuildPropertyCodeRepo) *CreateArtifactRequestArtifactBuildProperty {
	s.CodeRepo = v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetCommandContent(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.CommandContent = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetCommandType(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.CommandType = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetDockerfilePath(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.DockerfilePath = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetEnableGpu(v bool) *CreateArtifactRequestArtifactBuildProperty {
	s.EnableGpu = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetRegionId(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.RegionId = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetSourceContainerImage(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.SourceContainerImage = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetSourceImageId(v string) *CreateArtifactRequestArtifactBuildProperty {
	s.SourceImageId = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) SetSystemDiskSize(v int64) *CreateArtifactRequestArtifactBuildProperty {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildProperty) Validate() error {
	return dara.Validate(s)
}

type CreateArtifactRequestArtifactBuildPropertyBuildArgs struct {
	// The name of a specific build argument.
	//
	// example:
	//
	// ENV
	ArgumentName *string `json:"ArgumentName,omitempty" xml:"ArgumentName,omitempty"`
	// The value of a specific build argument.
	//
	// example:
	//
	// nginx:latest
	ArgumentValue *string `json:"ArgumentValue,omitempty" xml:"ArgumentValue,omitempty"`
}

func (s CreateArtifactRequestArtifactBuildPropertyBuildArgs) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactRequestArtifactBuildPropertyBuildArgs) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestArtifactBuildPropertyBuildArgs) GetArgumentName() *string {
	return s.ArgumentName
}

func (s *CreateArtifactRequestArtifactBuildPropertyBuildArgs) GetArgumentValue() *string {
	return s.ArgumentValue
}

func (s *CreateArtifactRequestArtifactBuildPropertyBuildArgs) SetArgumentName(v string) *CreateArtifactRequestArtifactBuildPropertyBuildArgs {
	s.ArgumentName = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyBuildArgs) SetArgumentValue(v string) *CreateArtifactRequestArtifactBuildPropertyBuildArgs {
	s.ArgumentValue = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyBuildArgs) Validate() error {
	return dara.Validate(s)
}

type CreateArtifactRequestArtifactBuildPropertyCodeRepo struct {
	// The name of the branch in the code repository.
	//
	// example:
	//
	// main
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// The endpoint.
	//
	// The URL address used to access the privately deployed GitLab instance.
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
	// >  This parameter is available only if the git repository is private.
	//
	// example:
	//
	// aliyun-computenest
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The platform type. Valid values:
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
	// The name of the repository.
	//
	// example:
	//
	// aliyun-computenest/quickstart-Lobexxx
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
}

func (s CreateArtifactRequestArtifactBuildPropertyCodeRepo) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactRequestArtifactBuildPropertyCodeRepo) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) GetBranch() *string {
	return s.Branch
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) GetOrgId() *string {
	return s.OrgId
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) GetOwner() *string {
	return s.Owner
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) GetPlatform() *string {
	return s.Platform
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) GetRepoId() *int64 {
	return s.RepoId
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) GetRepoName() *string {
	return s.RepoName
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetBranch(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Branch = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetEndpoint(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Endpoint = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetOrgId(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.OrgId = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetOwner(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Owner = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetPlatform(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.Platform = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetRepoId(v int64) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.RepoId = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) SetRepoName(v string) *CreateArtifactRequestArtifactBuildPropertyCodeRepo {
	s.RepoName = &v
	return s
}

func (s *CreateArtifactRequestArtifactBuildPropertyCodeRepo) Validate() error {
	return dara.Validate(s)
}

type CreateArtifactRequestArtifactProperty struct {
	// The commodity code of the service in Alibaba Cloud Marketplace.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// cmjj00xxxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The commodity version of the service in Alibaba Cloud Marketplace.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// V1.0
	CommodityVersion *string `json:"CommodityVersion,omitempty" xml:"CommodityVersion,omitempty"`
	// The image ID.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// m-0xij191j9cuev6xxxxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the image repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// crr-d8o1nponyc2t1gcg
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the image repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// wordpress
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The default repository type. Valid values:
	//
	// 	- `Public`: a public repository.
	//
	// 	- `Private`: a private repository.
	//
	// You can specify the RepoType or Summary parameter. The RepoType parameter is optional.
	//
	// example:
	//
	// Public
	RepoType *string `json:"RepoType,omitempty" xml:"RepoType,omitempty"`
	// The version tag of the image repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// v1
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The object URL of the deployment package.
	//
	// example:
	//
	// https://service-info-private.oss-cn-hangzhou.aliyuncs.com/1309208528xxxxxx/template/2e1ce8fc-xxxx-481c-9e8e-789ba9db487d.json
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateArtifactRequestArtifactProperty) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactRequestArtifactProperty) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestArtifactProperty) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CreateArtifactRequestArtifactProperty) GetCommodityVersion() *string {
	return s.CommodityVersion
}

func (s *CreateArtifactRequestArtifactProperty) GetImageId() *string {
	return s.ImageId
}

func (s *CreateArtifactRequestArtifactProperty) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateArtifactRequestArtifactProperty) GetRepoId() *string {
	return s.RepoId
}

func (s *CreateArtifactRequestArtifactProperty) GetRepoName() *string {
	return s.RepoName
}

func (s *CreateArtifactRequestArtifactProperty) GetRepoType() *string {
	return s.RepoType
}

func (s *CreateArtifactRequestArtifactProperty) GetTag() *string {
	return s.Tag
}

func (s *CreateArtifactRequestArtifactProperty) GetUrl() *string {
	return s.Url
}

func (s *CreateArtifactRequestArtifactProperty) SetCommodityCode(v string) *CreateArtifactRequestArtifactProperty {
	s.CommodityCode = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetCommodityVersion(v string) *CreateArtifactRequestArtifactProperty {
	s.CommodityVersion = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetImageId(v string) *CreateArtifactRequestArtifactProperty {
	s.ImageId = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRegionId(v string) *CreateArtifactRequestArtifactProperty {
	s.RegionId = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRepoId(v string) *CreateArtifactRequestArtifactProperty {
	s.RepoId = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRepoName(v string) *CreateArtifactRequestArtifactProperty {
	s.RepoName = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetRepoType(v string) *CreateArtifactRequestArtifactProperty {
	s.RepoType = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetTag(v string) *CreateArtifactRequestArtifactProperty {
	s.Tag = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) SetUrl(v string) *CreateArtifactRequestArtifactProperty {
	s.Url = &v
	return s
}

func (s *CreateArtifactRequestArtifactProperty) Validate() error {
	return dara.Validate(s)
}

type CreateArtifactRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateArtifactRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactRequestTag) GoString() string {
	return s.String()
}

func (s *CreateArtifactRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateArtifactRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateArtifactRequestTag) SetKey(v string) *CreateArtifactRequestTag {
	s.Key = &v
	return s
}

func (s *CreateArtifactRequestTag) SetValue(v string) *CreateArtifactRequestTag {
	s.Value = &v
	return s
}

func (s *CreateArtifactRequestTag) Validate() error {
	return dara.Validate(s)
}

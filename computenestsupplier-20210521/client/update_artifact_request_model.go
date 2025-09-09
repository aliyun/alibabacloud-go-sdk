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
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	ArtifactBuildProperty *UpdateArtifactRequestArtifactBuildProperty `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty" type:"Struct"`
	// The ID of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment package.
	ArtifactProperty *UpdateArtifactRequestArtifactProperty `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty" type:"Struct"`
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
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Permission fields are applicable to container image artifact and Helm Chart artifact. They can only change from Automatic to Public. Options:
	//
	// Public
	//
	// Automatic
	//
	// example:
	//
	// Public
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The IDs of the regions that support the deployment package.
	SupportRegionIds []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// The version name of the deployment package.
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
	return dara.Validate(s)
}

type UpdateArtifactRequestArtifactBuildProperty struct {
	// The build arguments used during the image build process.
	//
	// >  This parameter is available only if the ArtifactBuildType is Dockerfile type.
	BuildArgs []*UpdateArtifactRequestArtifactBuildPropertyBuildArgs `json:"BuildArgs,omitempty" xml:"BuildArgs,omitempty" type:"Repeated"`
	// The address of the code repository.
	//
	// >  This parameter is available only if the ArtifactBuildType is Dockerfile or Buildpacks type.
	CodeRepo *UpdateArtifactRequestArtifactBuildPropertyCodeRepo `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type UpdateArtifactRequestArtifactBuildPropertyBuildArgs struct {
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
	// m-0xij191j9cuev6ucxxxx
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID.
	//
	// >  This parameter is available only if the deployment package is an image.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Container Registry  repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// crr-yy4g68uhi39ttkm8
	RepoId *string `json:"RepoId,omitempty" xml:"RepoId,omitempty"`
	// The name of the Container Registry repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
	//
	// example:
	//
	// volcanosh/vc-webhook-manager
	RepoName *string `json:"RepoName,omitempty" xml:"RepoName,omitempty"`
	// The type of the repository.Valid values:
	//
	// 	- `Public`: a public repository.
	//
	// 	- `Private`: a private repository.
	//
	// >  This parameter is available only if the deployment package is a container image or of the Helm chart type.
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
	// The URL of the deployment package object.
	//
	//
	// > Note This parameter is available only if the deployment package is an file.
	//
	// example:
	//
	// https://service-info-private.oss-cn-hangzhou.aliyuncs.com/1309208528xxxxxx/template/2e1ce8fc-xxxx-481c-9e8e-789ba9db487d.json
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

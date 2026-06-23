// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListImageVersionsResponseBodyPagingInfo) *ListImageVersionsResponseBody
	GetPagingInfo() *ListImageVersionsResponseBodyPagingInfo
	SetRequestId(v string) *ListImageVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListImageVersionsResponseBody
	GetSuccess() *bool
}

type ListImageVersionsResponseBody struct {
	// The pagination details.
	PagingInfo *ListImageVersionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID, used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListImageVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageVersionsResponseBody) GetPagingInfo() *ListImageVersionsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListImageVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListImageVersionsResponseBody) SetPagingInfo(v *ListImageVersionsResponseBodyPagingInfo) *ListImageVersionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListImageVersionsResponseBody) SetRequestId(v string) *ListImageVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageVersionsResponseBody) SetSuccess(v bool) *ListImageVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListImageVersionsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListImageVersionsResponseBodyPagingInfo struct {
	// A list of image version details.
	ImageVersions []*ListImageVersionsResponseBodyPagingInfoImageVersions `json:"ImageVersions,omitempty" xml:"ImageVersions,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImageVersionsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListImageVersionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListImageVersionsResponseBodyPagingInfo) GetImageVersions() []*ListImageVersionsResponseBodyPagingInfoImageVersions {
	return s.ImageVersions
}

func (s *ListImageVersionsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImageVersionsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImageVersionsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListImageVersionsResponseBodyPagingInfo) SetImageVersions(v []*ListImageVersionsResponseBodyPagingInfoImageVersions) *ListImageVersionsResponseBodyPagingInfo {
	s.ImageVersions = v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListImageVersionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfo) SetPageSize(v int32) *ListImageVersionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListImageVersionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfo) Validate() error {
	if s.ImageVersions != nil {
		for _, item := range s.ImageVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImageVersionsResponseBodyPagingInfoImageVersions struct {
	// The accessibility of the image. Valid values:
	//
	// - `Public`: The image is visible to all users.
	//
	// - `Private`: The image is visible only to the creator.
	//
	// example:
	//
	// Public
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The ID of the VPC associated with the ACR instance.
	//
	// example:
	//
	// vpc-xxx
	AcrAssociatedVpcId *string `json:"AcrAssociatedVpcId,omitempty" xml:"AcrAssociatedVpcId,omitempty"`
	// The ACR endpoint.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com
	AcrEndpoint *string `json:"AcrEndpoint,omitempty" xml:"AcrEndpoint,omitempty"`
	// The ACR instance ID.
	//
	// example:
	//
	// cri-xxx
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The image build configuration.
	BuildConfig *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig `json:"BuildConfig,omitempty" xml:"BuildConfig,omitempty" type:"Struct"`
	// The time when the image was created, as a 64-bit timestamp.
	//
	// example:
	//
	// 1727055811000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the user who created the image.
	//
	// example:
	//
	// 123
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The image description.
	//
	// example:
	//
	// Test image created by xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether MaxCompute synchronization is enabled.
	//
	// example:
	//
	// false
	EnableSyncMaxCompute *bool `json:"EnableSyncMaxCompute,omitempty" xml:"EnableSyncMaxCompute,omitempty"`
	// The image ID.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image tag.
	//
	// example:
	//
	// v1.0.0
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// The image URI.
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/xxx/xxx:tag
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// The VPC URI of the image.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/xxx/xxx:tag
	ImageVpcUri *string `json:"ImageVpcUri,omitempty" xml:"ImageVpcUri,omitempty"`
	// Indicates whether the image is the default image.
	//
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The time when the image was last modified, as a 64-bit timestamp.
	//
	// example:
	//
	// 1727055811000
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The ID of the user who last modified the image.
	//
	// example:
	//
	// 123
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The image name.
	//
	// example:
	//
	// dataworks_image
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// Indicates whether the image is an official image.
	//
	// example:
	//
	// true
	Official *bool `json:"Official,omitempty" xml:"Official,omitempty"`
	// The ID of the image from the provider.
	//
	// example:
	//
	// acr_image_id
	ProviderImageId *string `json:"ProviderImageId,omitempty" xml:"ProviderImageId,omitempty"`
	// The type of the image provider. Valid values:
	//
	// - `ACR`: The image is from an ACR repository.
	//
	// - `DataWorks`: The image is an official image provided by DataWorks.
	//
	// example:
	//
	// ACR
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// The publishing stage of the image. Valid values:
	//
	// - `Untest`: Not tested.
	//
	// - `Testing`: Being tested.
	//
	// - `TestFailed`: Test failed.
	//
	// - `Unpublished`: Not published.
	//
	// - `Publishing`: Being published.
	//
	// - `Published`: Published.
	//
	// - `PublishFailed`: Publishing failed.
	//
	// - `Building`: Being built.
	//
	// - `BuildSuccess`: Build succeeded.
	//
	// - `BuildFailed`: Build failed.
	//
	// - `Accelerating`: Being accelerated.
	//
	// - `AccelerateSuccess`: Acceleration succeeded.
	//
	// - `AccelerateFailed`: Acceleration failed.
	//
	// example:
	//
	// Published
	PublishStage *string `json:"PublishStage,omitempty" xml:"PublishStage,omitempty"`
	// The repository name.
	//
	// example:
	//
	// repo_name
	RepositoryName *string `json:"RepositoryName,omitempty" xml:"RepositoryName,omitempty"`
	// The image size.
	//
	// example:
	//
	// 1GB
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the image. Valid values:
	//
	// - `Disabled`: The image is disabled.
	//
	// - `Expired`: The image is expired.
	//
	// - `Available`: The image is available.
	//
	// - `ReadOnly`: The image is read-only.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The supported modules and task types.
	Supported *ListImageVersionsResponseBodyPagingInfoImageVersionsSupported `json:"Supported,omitempty" xml:"Supported,omitempty" type:"Struct"`
	// The image version.
	//
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListImageVersionsResponseBodyPagingInfoImageVersions) String() string {
	return dara.Prettify(s)
}

func (s ListImageVersionsResponseBodyPagingInfoImageVersions) GoString() string {
	return s.String()
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetAcrAssociatedVpcId() *string {
	return s.AcrAssociatedVpcId
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetAcrEndpoint() *string {
	return s.AcrEndpoint
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetBuildConfig() *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig {
	return s.BuildConfig
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetCreator() *string {
	return s.Creator
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetDescription() *string {
	return s.Description
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetEnableSyncMaxCompute() *bool {
	return s.EnableSyncMaxCompute
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetId() *string {
	return s.Id
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetImageTag() *string {
	return s.ImageTag
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetImageUri() *string {
	return s.ImageUri
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetImageVpcUri() *string {
	return s.ImageVpcUri
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetModifier() *string {
	return s.Modifier
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetName() *string {
	return s.Name
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetNamespace() *string {
	return s.Namespace
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetOfficial() *bool {
	return s.Official
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetProviderImageId() *string {
	return s.ProviderImageId
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetProviderType() *string {
	return s.ProviderType
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetPublishStage() *string {
	return s.PublishStage
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetRepositoryName() *string {
	return s.RepositoryName
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetSize() *string {
	return s.Size
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetStatus() *string {
	return s.Status
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetSupported() *ListImageVersionsResponseBodyPagingInfoImageVersionsSupported {
	return s.Supported
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) GetVersion() *string {
	return s.Version
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetAccessibility(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Accessibility = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetAcrAssociatedVpcId(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.AcrAssociatedVpcId = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetAcrEndpoint(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.AcrEndpoint = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetAcrInstanceId(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.AcrInstanceId = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetBuildConfig(v *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.BuildConfig = v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetCreatedTime(v int64) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.CreatedTime = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetCreator(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Creator = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetDescription(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Description = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetEnableSyncMaxCompute(v bool) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.EnableSyncMaxCompute = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetId(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Id = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetImageTag(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.ImageTag = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetImageUri(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.ImageUri = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetImageVpcUri(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.ImageVpcUri = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetIsDefault(v bool) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.IsDefault = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetLastModifiedTime(v int64) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.LastModifiedTime = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetModifier(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Modifier = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetName(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Name = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetNamespace(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Namespace = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetOfficial(v bool) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Official = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetProviderImageId(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.ProviderImageId = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetProviderType(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.ProviderType = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetPublishStage(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.PublishStage = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetRepositoryName(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.RepositoryName = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetSize(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Size = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetStatus(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Status = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetSupported(v *ListImageVersionsResponseBodyPagingInfoImageVersionsSupported) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Supported = v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) SetVersion(v string) *ListImageVersionsResponseBodyPagingInfoImageVersions {
	s.Version = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersions) Validate() error {
	if s.BuildConfig != nil {
		if err := s.BuildConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Supported != nil {
		if err := s.Supported.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig struct {
	// The build type.
	//
	// example:
	//
	// PackageInstallation
	BuildType *string `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	// The pre-installation scripts.
	PackageInstallationScripts []*ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts `json:"PackageInstallationScripts,omitempty" xml:"PackageInstallationScripts,omitempty" type:"Repeated"`
}

func (s ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig) String() string {
	return dara.Prettify(s)
}

func (s ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig) GoString() string {
	return s.String()
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig) GetBuildType() *string {
	return s.BuildType
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig) GetPackageInstallationScripts() []*ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts {
	return s.PackageInstallationScripts
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig) SetBuildType(v string) *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig {
	s.BuildType = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig) SetPackageInstallationScripts(v []*ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts) *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig {
	s.PackageInstallationScripts = v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfig) Validate() error {
	if s.PackageInstallationScripts != nil {
		for _, item := range s.PackageInstallationScripts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts struct {
	// The script content. Separate multiple package names with a comma (,).
	//
	// example:
	//
	// requests
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The script type.
	//
	// example:
	//
	// Python3
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts) String() string {
	return dara.Prettify(s)
}

func (s ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts) GoString() string {
	return s.String()
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts) GetContent() *string {
	return s.Content
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts) GetType() *string {
	return s.Type
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts) SetContent(v string) *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts {
	s.Content = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts) SetType(v string) *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts {
	s.Type = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsBuildConfigPackageInstallationScripts) Validate() error {
	return dara.Validate(s)
}

type ListImageVersionsResponseBodyPagingInfoImageVersionsSupported struct {
	// The supported module. Valid value:
	//
	// - `Scheduler`: The scheduler module.
	//
	// example:
	//
	// Scheduler
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The supported task types.
	TaskTypes []*string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s ListImageVersionsResponseBodyPagingInfoImageVersionsSupported) String() string {
	return dara.Prettify(s)
}

func (s ListImageVersionsResponseBodyPagingInfoImageVersionsSupported) GoString() string {
	return s.String()
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsSupported) GetModule() *string {
	return s.Module
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsSupported) GetTaskTypes() []*string {
	return s.TaskTypes
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsSupported) SetModule(v string) *ListImageVersionsResponseBodyPagingInfoImageVersionsSupported {
	s.Module = &v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsSupported) SetTaskTypes(v []*string) *ListImageVersionsResponseBodyPagingInfoImageVersionsSupported {
	s.TaskTypes = v
	return s
}

func (s *ListImageVersionsResponseBodyPagingInfoImageVersionsSupported) Validate() error {
	return dara.Validate(s)
}

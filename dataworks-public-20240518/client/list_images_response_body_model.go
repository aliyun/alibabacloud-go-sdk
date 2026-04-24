// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListImagesResponseBodyPagingInfo) *ListImagesResponseBody
	GetPagingInfo() *ListImagesResponseBodyPagingInfo
	SetRequestId(v string) *ListImagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListImagesResponseBody
	GetSuccess() *bool
}

type ListImagesResponseBody struct {
	PagingInfo *ListImagesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) GetPagingInfo() *ListImagesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListImagesResponseBody) SetPagingInfo(v *ListImagesResponseBodyPagingInfo) *ListImagesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetSuccess(v bool) *ListImagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListImagesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListImagesResponseBodyPagingInfo struct {
	ImageList []*ListImagesResponseBodyPagingInfoImageList `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImagesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyPagingInfo) GetImageList() []*ListImagesResponseBodyPagingInfoImageList {
	return s.ImageList
}

func (s *ListImagesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListImagesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListImagesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListImagesResponseBodyPagingInfo) SetImageList(v []*ListImagesResponseBodyPagingInfoImageList) *ListImagesResponseBodyPagingInfo {
	s.ImageList = v
	return s
}

func (s *ListImagesResponseBodyPagingInfo) SetPageNumber(v int32) *ListImagesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfo) SetPageSize(v int32) *ListImagesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfo) SetTotalCount(v int32) *ListImagesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfo) Validate() error {
	if s.ImageList != nil {
		for _, item := range s.ImageList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListImagesResponseBodyPagingInfoImageList struct {
	// example:
	//
	// Public
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// example:
	//
	// vpc-xxx
	AcrAssociatedVpcId *string `json:"AcrAssociatedVpcId,omitempty" xml:"AcrAssociatedVpcId,omitempty"`
	// ACR Endpoint
	//
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com
	AcrEndpoint *string `json:"AcrEndpoint,omitempty" xml:"AcrEndpoint,omitempty"`
	// example:
	//
	// cri-xxx
	AcrInstanceId *string                                               `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	BuildConfig   *ListImagesResponseBodyPagingInfoImageListBuildConfig `json:"BuildConfig,omitempty" xml:"BuildConfig,omitempty" type:"Struct"`
	// example:
	//
	// 1727055811000
	CreatedTime *int64 `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 123
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// Test image created by xxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// false
	EnableSyncMaxCompute *bool `json:"EnableSyncMaxCompute,omitempty" xml:"EnableSyncMaxCompute,omitempty"`
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// v1.0.0
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// example:
	//
	// registry.cn-hangzhou.aliyuncs.com/xxx/xxx:tag
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com/xxx/xxx:tag
	ImageVpcUri *string `json:"ImageVpcUri,omitempty" xml:"ImageVpcUri,omitempty"`
	// example:
	//
	// false
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// 1727055811000
	LastModifiedTime *int64 `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// example:
	//
	// 123
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// dataworks_image
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// example:
	//
	// true
	Official *bool `json:"Official,omitempty" xml:"Official,omitempty"`
	// example:
	//
	// acr_image_id
	ProviderImageId *string `json:"ProviderImageId,omitempty" xml:"ProviderImageId,omitempty"`
	// example:
	//
	// ACR
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// example:
	//
	// Published
	PublishStage *string `json:"PublishStage,omitempty" xml:"PublishStage,omitempty"`
	// example:
	//
	// repo_name
	RepositoryName *string `json:"RepositoryName,omitempty" xml:"RepositoryName,omitempty"`
	// example:
	//
	// 1GB
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// Available
	Status    *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Supported *ListImagesResponseBodyPagingInfoImageListSupported `json:"Supported,omitempty" xml:"Supported,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListImagesResponseBodyPagingInfoImageList) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyPagingInfoImageList) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetAccessibility() *string {
	return s.Accessibility
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetAcrAssociatedVpcId() *string {
	return s.AcrAssociatedVpcId
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetAcrEndpoint() *string {
	return s.AcrEndpoint
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetBuildConfig() *ListImagesResponseBodyPagingInfoImageListBuildConfig {
	return s.BuildConfig
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetCreator() *string {
	return s.Creator
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetDescription() *string {
	return s.Description
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetEnableSyncMaxCompute() *bool {
	return s.EnableSyncMaxCompute
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetId() *string {
	return s.Id
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetImageTag() *string {
	return s.ImageTag
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetImageUri() *string {
	return s.ImageUri
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetImageVpcUri() *string {
	return s.ImageVpcUri
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetModifier() *string {
	return s.Modifier
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetName() *string {
	return s.Name
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetNamespace() *string {
	return s.Namespace
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetOfficial() *bool {
	return s.Official
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetProviderImageId() *string {
	return s.ProviderImageId
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetProviderType() *string {
	return s.ProviderType
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetPublishStage() *string {
	return s.PublishStage
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetRepositoryName() *string {
	return s.RepositoryName
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetSize() *string {
	return s.Size
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetStatus() *string {
	return s.Status
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetSupported() *ListImagesResponseBodyPagingInfoImageListSupported {
	return s.Supported
}

func (s *ListImagesResponseBodyPagingInfoImageList) GetVersion() *string {
	return s.Version
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetAccessibility(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.Accessibility = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetAcrAssociatedVpcId(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.AcrAssociatedVpcId = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetAcrEndpoint(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.AcrEndpoint = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetAcrInstanceId(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.AcrInstanceId = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetBuildConfig(v *ListImagesResponseBodyPagingInfoImageListBuildConfig) *ListImagesResponseBodyPagingInfoImageList {
	s.BuildConfig = v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetCreatedTime(v int64) *ListImagesResponseBodyPagingInfoImageList {
	s.CreatedTime = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetCreator(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.Creator = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetDescription(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetEnableSyncMaxCompute(v bool) *ListImagesResponseBodyPagingInfoImageList {
	s.EnableSyncMaxCompute = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetId(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.Id = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetImageTag(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.ImageTag = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetImageUri(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.ImageUri = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetImageVpcUri(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.ImageVpcUri = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetIsDefault(v bool) *ListImagesResponseBodyPagingInfoImageList {
	s.IsDefault = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetLastModifiedTime(v int64) *ListImagesResponseBodyPagingInfoImageList {
	s.LastModifiedTime = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetModifier(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.Modifier = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetName(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.Name = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetNamespace(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.Namespace = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetOfficial(v bool) *ListImagesResponseBodyPagingInfoImageList {
	s.Official = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetProviderImageId(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.ProviderImageId = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetProviderType(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.ProviderType = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetPublishStage(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.PublishStage = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetRepositoryName(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.RepositoryName = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetSize(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.Size = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetStatus(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.Status = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetSupported(v *ListImagesResponseBodyPagingInfoImageListSupported) *ListImagesResponseBodyPagingInfoImageList {
	s.Supported = v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) SetVersion(v string) *ListImagesResponseBodyPagingInfoImageList {
	s.Version = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageList) Validate() error {
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

type ListImagesResponseBodyPagingInfoImageListBuildConfig struct {
	// example:
	//
	// PackageInstallation
	BuildType                  *string                                                                           `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	PackageInstallationScripts []*ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts `json:"PackageInstallationScripts,omitempty" xml:"PackageInstallationScripts,omitempty" type:"Repeated"`
}

func (s ListImagesResponseBodyPagingInfoImageListBuildConfig) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyPagingInfoImageListBuildConfig) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyPagingInfoImageListBuildConfig) GetBuildType() *string {
	return s.BuildType
}

func (s *ListImagesResponseBodyPagingInfoImageListBuildConfig) GetPackageInstallationScripts() []*ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts {
	return s.PackageInstallationScripts
}

func (s *ListImagesResponseBodyPagingInfoImageListBuildConfig) SetBuildType(v string) *ListImagesResponseBodyPagingInfoImageListBuildConfig {
	s.BuildType = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageListBuildConfig) SetPackageInstallationScripts(v []*ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts) *ListImagesResponseBodyPagingInfoImageListBuildConfig {
	s.PackageInstallationScripts = v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageListBuildConfig) Validate() error {
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

type ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts struct {
	// example:
	//
	// requests
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// Python3
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts) GetContent() *string {
	return s.Content
}

func (s *ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts) GetType() *string {
	return s.Type
}

func (s *ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts) SetContent(v string) *ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts {
	s.Content = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts) SetType(v string) *ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts {
	s.Type = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageListBuildConfigPackageInstallationScripts) Validate() error {
	return dara.Validate(s)
}

type ListImagesResponseBodyPagingInfoImageListSupported struct {
	// example:
	//
	// Scheduler
	Module    *string   `json:"Module,omitempty" xml:"Module,omitempty"`
	TaskTypes []*string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s ListImagesResponseBodyPagingInfoImageListSupported) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyPagingInfoImageListSupported) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyPagingInfoImageListSupported) GetModule() *string {
	return s.Module
}

func (s *ListImagesResponseBodyPagingInfoImageListSupported) GetTaskTypes() []*string {
	return s.TaskTypes
}

func (s *ListImagesResponseBodyPagingInfoImageListSupported) SetModule(v string) *ListImagesResponseBodyPagingInfoImageListSupported {
	s.Module = &v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageListSupported) SetTaskTypes(v []*string) *ListImagesResponseBodyPagingInfoImageListSupported {
	s.TaskTypes = v
	return s
}

func (s *ListImagesResponseBodyPagingInfoImageListSupported) Validate() error {
	return dara.Validate(s)
}

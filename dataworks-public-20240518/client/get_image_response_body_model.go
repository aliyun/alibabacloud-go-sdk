// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v *GetImageResponseBodyImage) *GetImageResponseBody
	GetImage() *GetImageResponseBodyImage
	SetRequestId(v string) *GetImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetImageResponseBody
	GetSuccess() *bool
}

type GetImageResponseBody struct {
	Image *GetImageResponseBodyImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Struct"`
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageResponseBody) GetImage() *GetImageResponseBodyImage {
	return s.Image
}

func (s *GetImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetImageResponseBody) SetImage(v *GetImageResponseBodyImage) *GetImageResponseBody {
	s.Image = v
	return s
}

func (s *GetImageResponseBody) SetRequestId(v string) *GetImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageResponseBody) SetSuccess(v bool) *GetImageResponseBody {
	s.Success = &v
	return s
}

func (s *GetImageResponseBody) Validate() error {
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageResponseBodyImage struct {
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
	AcrInstanceId *string                               `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	BuildConfig   *GetImageResponseBodyImageBuildConfig `json:"BuildConfig,omitempty" xml:"BuildConfig,omitempty" type:"Struct"`
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
	Status    *string                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Supported *GetImageResponseBodyImageSupported `json:"Supported,omitempty" xml:"Supported,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetImageResponseBodyImage) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyImage) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImage) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetImageResponseBodyImage) GetAcrAssociatedVpcId() *string {
	return s.AcrAssociatedVpcId
}

func (s *GetImageResponseBodyImage) GetAcrEndpoint() *string {
	return s.AcrEndpoint
}

func (s *GetImageResponseBodyImage) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *GetImageResponseBodyImage) GetBuildConfig() *GetImageResponseBodyImageBuildConfig {
	return s.BuildConfig
}

func (s *GetImageResponseBodyImage) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *GetImageResponseBodyImage) GetCreator() *string {
	return s.Creator
}

func (s *GetImageResponseBodyImage) GetDescription() *string {
	return s.Description
}

func (s *GetImageResponseBodyImage) GetEnableSyncMaxCompute() *bool {
	return s.EnableSyncMaxCompute
}

func (s *GetImageResponseBodyImage) GetId() *string {
	return s.Id
}

func (s *GetImageResponseBodyImage) GetImageTag() *string {
	return s.ImageTag
}

func (s *GetImageResponseBodyImage) GetImageUri() *string {
	return s.ImageUri
}

func (s *GetImageResponseBodyImage) GetImageVpcUri() *string {
	return s.ImageVpcUri
}

func (s *GetImageResponseBodyImage) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *GetImageResponseBodyImage) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *GetImageResponseBodyImage) GetModifier() *string {
	return s.Modifier
}

func (s *GetImageResponseBodyImage) GetName() *string {
	return s.Name
}

func (s *GetImageResponseBodyImage) GetNamespace() *string {
	return s.Namespace
}

func (s *GetImageResponseBodyImage) GetOfficial() *bool {
	return s.Official
}

func (s *GetImageResponseBodyImage) GetProviderImageId() *string {
	return s.ProviderImageId
}

func (s *GetImageResponseBodyImage) GetProviderType() *string {
	return s.ProviderType
}

func (s *GetImageResponseBodyImage) GetPublishStage() *string {
	return s.PublishStage
}

func (s *GetImageResponseBodyImage) GetRepositoryName() *string {
	return s.RepositoryName
}

func (s *GetImageResponseBodyImage) GetSize() *string {
	return s.Size
}

func (s *GetImageResponseBodyImage) GetStatus() *string {
	return s.Status
}

func (s *GetImageResponseBodyImage) GetSupported() *GetImageResponseBodyImageSupported {
	return s.Supported
}

func (s *GetImageResponseBodyImage) GetVersion() *string {
	return s.Version
}

func (s *GetImageResponseBodyImage) SetAccessibility(v string) *GetImageResponseBodyImage {
	s.Accessibility = &v
	return s
}

func (s *GetImageResponseBodyImage) SetAcrAssociatedVpcId(v string) *GetImageResponseBodyImage {
	s.AcrAssociatedVpcId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetAcrEndpoint(v string) *GetImageResponseBodyImage {
	s.AcrEndpoint = &v
	return s
}

func (s *GetImageResponseBodyImage) SetAcrInstanceId(v string) *GetImageResponseBodyImage {
	s.AcrInstanceId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetBuildConfig(v *GetImageResponseBodyImageBuildConfig) *GetImageResponseBodyImage {
	s.BuildConfig = v
	return s
}

func (s *GetImageResponseBodyImage) SetCreatedTime(v int64) *GetImageResponseBodyImage {
	s.CreatedTime = &v
	return s
}

func (s *GetImageResponseBodyImage) SetCreator(v string) *GetImageResponseBodyImage {
	s.Creator = &v
	return s
}

func (s *GetImageResponseBodyImage) SetDescription(v string) *GetImageResponseBodyImage {
	s.Description = &v
	return s
}

func (s *GetImageResponseBodyImage) SetEnableSyncMaxCompute(v bool) *GetImageResponseBodyImage {
	s.EnableSyncMaxCompute = &v
	return s
}

func (s *GetImageResponseBodyImage) SetId(v string) *GetImageResponseBodyImage {
	s.Id = &v
	return s
}

func (s *GetImageResponseBodyImage) SetImageTag(v string) *GetImageResponseBodyImage {
	s.ImageTag = &v
	return s
}

func (s *GetImageResponseBodyImage) SetImageUri(v string) *GetImageResponseBodyImage {
	s.ImageUri = &v
	return s
}

func (s *GetImageResponseBodyImage) SetImageVpcUri(v string) *GetImageResponseBodyImage {
	s.ImageVpcUri = &v
	return s
}

func (s *GetImageResponseBodyImage) SetIsDefault(v bool) *GetImageResponseBodyImage {
	s.IsDefault = &v
	return s
}

func (s *GetImageResponseBodyImage) SetLastModifiedTime(v int64) *GetImageResponseBodyImage {
	s.LastModifiedTime = &v
	return s
}

func (s *GetImageResponseBodyImage) SetModifier(v string) *GetImageResponseBodyImage {
	s.Modifier = &v
	return s
}

func (s *GetImageResponseBodyImage) SetName(v string) *GetImageResponseBodyImage {
	s.Name = &v
	return s
}

func (s *GetImageResponseBodyImage) SetNamespace(v string) *GetImageResponseBodyImage {
	s.Namespace = &v
	return s
}

func (s *GetImageResponseBodyImage) SetOfficial(v bool) *GetImageResponseBodyImage {
	s.Official = &v
	return s
}

func (s *GetImageResponseBodyImage) SetProviderImageId(v string) *GetImageResponseBodyImage {
	s.ProviderImageId = &v
	return s
}

func (s *GetImageResponseBodyImage) SetProviderType(v string) *GetImageResponseBodyImage {
	s.ProviderType = &v
	return s
}

func (s *GetImageResponseBodyImage) SetPublishStage(v string) *GetImageResponseBodyImage {
	s.PublishStage = &v
	return s
}

func (s *GetImageResponseBodyImage) SetRepositoryName(v string) *GetImageResponseBodyImage {
	s.RepositoryName = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSize(v string) *GetImageResponseBodyImage {
	s.Size = &v
	return s
}

func (s *GetImageResponseBodyImage) SetStatus(v string) *GetImageResponseBodyImage {
	s.Status = &v
	return s
}

func (s *GetImageResponseBodyImage) SetSupported(v *GetImageResponseBodyImageSupported) *GetImageResponseBodyImage {
	s.Supported = v
	return s
}

func (s *GetImageResponseBodyImage) SetVersion(v string) *GetImageResponseBodyImage {
	s.Version = &v
	return s
}

func (s *GetImageResponseBodyImage) Validate() error {
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

type GetImageResponseBodyImageBuildConfig struct {
	// example:
	//
	// PackageInstallation
	BuildType                  *string                                                           `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	PackageInstallationScripts []*GetImageResponseBodyImageBuildConfigPackageInstallationScripts `json:"PackageInstallationScripts,omitempty" xml:"PackageInstallationScripts,omitempty" type:"Repeated"`
}

func (s GetImageResponseBodyImageBuildConfig) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyImageBuildConfig) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageBuildConfig) GetBuildType() *string {
	return s.BuildType
}

func (s *GetImageResponseBodyImageBuildConfig) GetPackageInstallationScripts() []*GetImageResponseBodyImageBuildConfigPackageInstallationScripts {
	return s.PackageInstallationScripts
}

func (s *GetImageResponseBodyImageBuildConfig) SetBuildType(v string) *GetImageResponseBodyImageBuildConfig {
	s.BuildType = &v
	return s
}

func (s *GetImageResponseBodyImageBuildConfig) SetPackageInstallationScripts(v []*GetImageResponseBodyImageBuildConfigPackageInstallationScripts) *GetImageResponseBodyImageBuildConfig {
	s.PackageInstallationScripts = v
	return s
}

func (s *GetImageResponseBodyImageBuildConfig) Validate() error {
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

type GetImageResponseBodyImageBuildConfigPackageInstallationScripts struct {
	// example:
	//
	// requests
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// Python3
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetImageResponseBodyImageBuildConfigPackageInstallationScripts) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyImageBuildConfigPackageInstallationScripts) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageBuildConfigPackageInstallationScripts) GetContent() *string {
	return s.Content
}

func (s *GetImageResponseBodyImageBuildConfigPackageInstallationScripts) GetType() *string {
	return s.Type
}

func (s *GetImageResponseBodyImageBuildConfigPackageInstallationScripts) SetContent(v string) *GetImageResponseBodyImageBuildConfigPackageInstallationScripts {
	s.Content = &v
	return s
}

func (s *GetImageResponseBodyImageBuildConfigPackageInstallationScripts) SetType(v string) *GetImageResponseBodyImageBuildConfigPackageInstallationScripts {
	s.Type = &v
	return s
}

func (s *GetImageResponseBodyImageBuildConfigPackageInstallationScripts) Validate() error {
	return dara.Validate(s)
}

type GetImageResponseBodyImageSupported struct {
	// example:
	//
	// Scheduler
	Module    *string   `json:"Module,omitempty" xml:"Module,omitempty"`
	TaskTypes []*string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s GetImageResponseBodyImageSupported) String() string {
	return dara.Prettify(s)
}

func (s GetImageResponseBodyImageSupported) GoString() string {
	return s.String()
}

func (s *GetImageResponseBodyImageSupported) GetModule() *string {
	return s.Module
}

func (s *GetImageResponseBodyImageSupported) GetTaskTypes() []*string {
	return s.TaskTypes
}

func (s *GetImageResponseBodyImageSupported) SetModule(v string) *GetImageResponseBodyImageSupported {
	s.Module = &v
	return s
}

func (s *GetImageResponseBodyImageSupported) SetTaskTypes(v []*string) *GetImageResponseBodyImageSupported {
	s.TaskTypes = v
	return s
}

func (s *GetImageResponseBodyImageSupported) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateImageRequest
	GetAccessibility() *string
	SetAcrAssociatedVpcId(v string) *CreateImageRequest
	GetAcrAssociatedVpcId() *string
	SetAcrInstanceId(v string) *CreateImageRequest
	GetAcrInstanceId() *string
	SetBuildConfig(v *CreateImageRequestBuildConfig) *CreateImageRequest
	GetBuildConfig() *CreateImageRequestBuildConfig
	SetClientToken(v string) *CreateImageRequest
	GetClientToken() *string
	SetDescription(v string) *CreateImageRequest
	GetDescription() *string
	SetEnableSyncMaxCompute(v bool) *CreateImageRequest
	GetEnableSyncMaxCompute() *bool
	SetImageUri(v string) *CreateImageRequest
	GetImageUri() *string
	SetName(v string) *CreateImageRequest
	GetName() *string
	SetNamespace(v string) *CreateImageRequest
	GetNamespace() *string
	SetProviderImageId(v string) *CreateImageRequest
	GetProviderImageId() *string
	SetProviderType(v string) *CreateImageRequest
	GetProviderType() *string
	SetRepositoryName(v string) *CreateImageRequest
	GetRepositoryName() *string
	SetSupported(v *CreateImageRequestSupported) *CreateImageRequest
	GetSupported() *CreateImageRequestSupported
}

type CreateImageRequest struct {
	// The image visibility. Valid values:
	//
	// - Public: visible to all users.
	//
	// - Private: visible only to the creator.
	//
	// example:
	//
	// Public
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The VPC ID associated with the ACR instance. This parameter is required when referencing an ACR image.
	//
	// example:
	//
	// vpc-xxx
	AcrAssociatedVpcId *string `json:"AcrAssociatedVpcId,omitempty" xml:"AcrAssociatedVpcId,omitempty"`
	// The ACR instance ID. This parameter is required when referencing an ACR image.
	//
	// example:
	//
	// acr_instance_id
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The image build configuration.
	BuildConfig *CreateImageRequestBuildConfig `json:"BuildConfig,omitempty" xml:"BuildConfig,omitempty" type:"Struct"`
	// The client idempotency token.
	//
	// This parameter is required.
	//
	// example:
	//
	// dasfsd-94fqwe-da8d
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The image description, up to 128 characters.
	//
	// example:
	//
	// create by xxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to synchronize the image to MaxCompute. Specify this parameter when referencing an ACR image. Default value: false.
	//
	// example:
	//
	// false
	EnableSyncMaxCompute *bool `json:"EnableSyncMaxCompute,omitempty" xml:"EnableSyncMaxCompute,omitempty"`
	// The image URI. This parameter is required when referencing an ACR image.
	//
	// example:
	//
	// registry-vpc.cn-beijing.cr.aliyuncs.com/namespace/image:0.1.0
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// The image name, which can contain lowercase letters, digits, and underscores (_), up to 128 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// task_image_001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The image namespace. Set this parameter to DataWorks Default when referencing a DataWorks official image.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace_name
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The image ID from the image provider. This parameter is required when referencing a DataWorks official image.
	//
	// example:
	//
	// System_shell_20251201
	ProviderImageId *string `json:"ProviderImageId,omitempty" xml:"ProviderImageId,omitempty"`
	// The image reference data type. Valid values:
	//
	// - ACR: ACR image repository.
	//
	// - DataWorks: DataWorks official image.
	//
	// This parameter is required.
	//
	// example:
	//
	// ACR
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// The image repository name. Set this parameter to DataWorks Default when referencing a DataWorks official image.
	//
	// This parameter is required.
	//
	// example:
	//
	// repo_name
	RepositoryName *string `json:"RepositoryName,omitempty" xml:"RepositoryName,omitempty"`
	// The image sub-purpose.
	//
	// This parameter is required.
	Supported *CreateImageRequestSupported `json:"Supported,omitempty" xml:"Supported,omitempty" type:"Struct"`
}

func (s CreateImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateImageRequest) GetAcrAssociatedVpcId() *string {
	return s.AcrAssociatedVpcId
}

func (s *CreateImageRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CreateImageRequest) GetBuildConfig() *CreateImageRequestBuildConfig {
	return s.BuildConfig
}

func (s *CreateImageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateImageRequest) GetEnableSyncMaxCompute() *bool {
	return s.EnableSyncMaxCompute
}

func (s *CreateImageRequest) GetImageUri() *string {
	return s.ImageUri
}

func (s *CreateImageRequest) GetName() *string {
	return s.Name
}

func (s *CreateImageRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateImageRequest) GetProviderImageId() *string {
	return s.ProviderImageId
}

func (s *CreateImageRequest) GetProviderType() *string {
	return s.ProviderType
}

func (s *CreateImageRequest) GetRepositoryName() *string {
	return s.RepositoryName
}

func (s *CreateImageRequest) GetSupported() *CreateImageRequestSupported {
	return s.Supported
}

func (s *CreateImageRequest) SetAccessibility(v string) *CreateImageRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateImageRequest) SetAcrAssociatedVpcId(v string) *CreateImageRequest {
	s.AcrAssociatedVpcId = &v
	return s
}

func (s *CreateImageRequest) SetAcrInstanceId(v string) *CreateImageRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateImageRequest) SetBuildConfig(v *CreateImageRequestBuildConfig) *CreateImageRequest {
	s.BuildConfig = v
	return s
}

func (s *CreateImageRequest) SetClientToken(v string) *CreateImageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageRequest) SetDescription(v string) *CreateImageRequest {
	s.Description = &v
	return s
}

func (s *CreateImageRequest) SetEnableSyncMaxCompute(v bool) *CreateImageRequest {
	s.EnableSyncMaxCompute = &v
	return s
}

func (s *CreateImageRequest) SetImageUri(v string) *CreateImageRequest {
	s.ImageUri = &v
	return s
}

func (s *CreateImageRequest) SetName(v string) *CreateImageRequest {
	s.Name = &v
	return s
}

func (s *CreateImageRequest) SetNamespace(v string) *CreateImageRequest {
	s.Namespace = &v
	return s
}

func (s *CreateImageRequest) SetProviderImageId(v string) *CreateImageRequest {
	s.ProviderImageId = &v
	return s
}

func (s *CreateImageRequest) SetProviderType(v string) *CreateImageRequest {
	s.ProviderType = &v
	return s
}

func (s *CreateImageRequest) SetRepositoryName(v string) *CreateImageRequest {
	s.RepositoryName = &v
	return s
}

func (s *CreateImageRequest) SetSupported(v *CreateImageRequestSupported) *CreateImageRequest {
	s.Supported = v
	return s
}

func (s *CreateImageRequest) Validate() error {
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

type CreateImageRequestBuildConfig struct {
	// The build type.
	//
	// example:
	//
	// PackageInstallation
	BuildType *string `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	// The list of pre-installation scripts.
	PackageInstallationScripts []*CreateImageRequestBuildConfigPackageInstallationScripts `json:"PackageInstallationScripts,omitempty" xml:"PackageInstallationScripts,omitempty" type:"Repeated"`
}

func (s CreateImageRequestBuildConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequestBuildConfig) GoString() string {
	return s.String()
}

func (s *CreateImageRequestBuildConfig) GetBuildType() *string {
	return s.BuildType
}

func (s *CreateImageRequestBuildConfig) GetPackageInstallationScripts() []*CreateImageRequestBuildConfigPackageInstallationScripts {
	return s.PackageInstallationScripts
}

func (s *CreateImageRequestBuildConfig) SetBuildType(v string) *CreateImageRequestBuildConfig {
	s.BuildType = &v
	return s
}

func (s *CreateImageRequestBuildConfig) SetPackageInstallationScripts(v []*CreateImageRequestBuildConfigPackageInstallationScripts) *CreateImageRequestBuildConfig {
	s.PackageInstallationScripts = v
	return s
}

func (s *CreateImageRequestBuildConfig) Validate() error {
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

type CreateImageRequestBuildConfigPackageInstallationScripts struct {
	// The script content. If the content consists of package names, separate them with commas (,).
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

func (s CreateImageRequestBuildConfigPackageInstallationScripts) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequestBuildConfigPackageInstallationScripts) GoString() string {
	return s.String()
}

func (s *CreateImageRequestBuildConfigPackageInstallationScripts) GetContent() *string {
	return s.Content
}

func (s *CreateImageRequestBuildConfigPackageInstallationScripts) GetType() *string {
	return s.Type
}

func (s *CreateImageRequestBuildConfigPackageInstallationScripts) SetContent(v string) *CreateImageRequestBuildConfigPackageInstallationScripts {
	s.Content = &v
	return s
}

func (s *CreateImageRequestBuildConfigPackageInstallationScripts) SetType(v string) *CreateImageRequestBuildConfigPackageInstallationScripts {
	s.Type = &v
	return s
}

func (s *CreateImageRequestBuildConfigPackageInstallationScripts) Validate() error {
	return dara.Validate(s)
}

type CreateImageRequestSupported struct {
	// The image sub-module. Valid values:
	//
	// - Scheduler: DataStudio.
	//
	// example:
	//
	// Scheduler
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The list of supported node types.
	TaskTypes []*string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s CreateImageRequestSupported) String() string {
	return dara.Prettify(s)
}

func (s CreateImageRequestSupported) GoString() string {
	return s.String()
}

func (s *CreateImageRequestSupported) GetModule() *string {
	return s.Module
}

func (s *CreateImageRequestSupported) GetTaskTypes() []*string {
	return s.TaskTypes
}

func (s *CreateImageRequestSupported) SetModule(v string) *CreateImageRequestSupported {
	s.Module = &v
	return s
}

func (s *CreateImageRequestSupported) SetTaskTypes(v []*string) *CreateImageRequestSupported {
	s.TaskTypes = v
	return s
}

func (s *CreateImageRequestSupported) Validate() error {
	return dara.Validate(s)
}

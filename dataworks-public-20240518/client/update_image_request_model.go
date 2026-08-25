// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *UpdateImageRequest
	GetAccessibility() *string
	SetAcrAssociatedVpcId(v string) *UpdateImageRequest
	GetAcrAssociatedVpcId() *string
	SetAcrInstanceId(v string) *UpdateImageRequest
	GetAcrInstanceId() *string
	SetBuildConfig(v *UpdateImageRequestBuildConfig) *UpdateImageRequest
	GetBuildConfig() *UpdateImageRequestBuildConfig
	SetDescription(v string) *UpdateImageRequest
	GetDescription() *string
	SetId(v string) *UpdateImageRequest
	GetId() *string
	SetImageUri(v string) *UpdateImageRequest
	GetImageUri() *string
	SetName(v string) *UpdateImageRequest
	GetName() *string
	SetNamespace(v string) *UpdateImageRequest
	GetNamespace() *string
	SetProviderImageId(v string) *UpdateImageRequest
	GetProviderImageId() *string
	SetRepositoryName(v string) *UpdateImageRequest
	GetRepositoryName() *string
	SetSupported(v *UpdateImageRequestSupported) *UpdateImageRequest
	GetSupported() *UpdateImageRequestSupported
}

type UpdateImageRequest struct {
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
	// The Container Registry (ACR) instance ID. This parameter is required when referencing an ACR image.
	//
	// example:
	//
	// acr_instance_id
	AcrInstanceId *string `json:"AcrInstanceId,omitempty" xml:"AcrInstanceId,omitempty"`
	// The image build configuration.
	BuildConfig *UpdateImageRequestBuildConfig `json:"BuildConfig,omitempty" xml:"BuildConfig,omitempty" type:"Struct"`
	// The image description.
	//
	// example:
	//
	// create by xxxx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image URI. This parameter is required when referencing an ACR image.
	//
	// example:
	//
	// registry-vpc.cn-beijing.cr.aliyuncs.com/namespace/image:0.1.0
	ImageUri *string `json:"ImageUri,omitempty" xml:"ImageUri,omitempty"`
	// The image name.
	//
	// example:
	//
	// task_image_001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The image namespace. Set this parameter to DataWorks Default when referencing a DataWorks official image.
	//
	// example:
	//
	// namespace_name
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The provider image ID. This parameter is required when referencing a DataWorks official image.
	//
	// example:
	//
	// System_shell_20251201
	ProviderImageId *string `json:"ProviderImageId,omitempty" xml:"ProviderImageId,omitempty"`
	// The image repository name. Set this parameter to DataWorks Default when referencing a DataWorks official image.
	//
	// example:
	//
	// repo_name
	RepositoryName *string `json:"RepositoryName,omitempty" xml:"RepositoryName,omitempty"`
	// The image sub-purpose.
	Supported *UpdateImageRequestSupported `json:"Supported,omitempty" xml:"Supported,omitempty" type:"Struct"`
}

func (s UpdateImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *UpdateImageRequest) GetAcrAssociatedVpcId() *string {
	return s.AcrAssociatedVpcId
}

func (s *UpdateImageRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *UpdateImageRequest) GetBuildConfig() *UpdateImageRequestBuildConfig {
	return s.BuildConfig
}

func (s *UpdateImageRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateImageRequest) GetId() *string {
	return s.Id
}

func (s *UpdateImageRequest) GetImageUri() *string {
	return s.ImageUri
}

func (s *UpdateImageRequest) GetName() *string {
	return s.Name
}

func (s *UpdateImageRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateImageRequest) GetProviderImageId() *string {
	return s.ProviderImageId
}

func (s *UpdateImageRequest) GetRepositoryName() *string {
	return s.RepositoryName
}

func (s *UpdateImageRequest) GetSupported() *UpdateImageRequestSupported {
	return s.Supported
}

func (s *UpdateImageRequest) SetAccessibility(v string) *UpdateImageRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateImageRequest) SetAcrAssociatedVpcId(v string) *UpdateImageRequest {
	s.AcrAssociatedVpcId = &v
	return s
}

func (s *UpdateImageRequest) SetAcrInstanceId(v string) *UpdateImageRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *UpdateImageRequest) SetBuildConfig(v *UpdateImageRequestBuildConfig) *UpdateImageRequest {
	s.BuildConfig = v
	return s
}

func (s *UpdateImageRequest) SetDescription(v string) *UpdateImageRequest {
	s.Description = &v
	return s
}

func (s *UpdateImageRequest) SetId(v string) *UpdateImageRequest {
	s.Id = &v
	return s
}

func (s *UpdateImageRequest) SetImageUri(v string) *UpdateImageRequest {
	s.ImageUri = &v
	return s
}

func (s *UpdateImageRequest) SetName(v string) *UpdateImageRequest {
	s.Name = &v
	return s
}

func (s *UpdateImageRequest) SetNamespace(v string) *UpdateImageRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateImageRequest) SetProviderImageId(v string) *UpdateImageRequest {
	s.ProviderImageId = &v
	return s
}

func (s *UpdateImageRequest) SetRepositoryName(v string) *UpdateImageRequest {
	s.RepositoryName = &v
	return s
}

func (s *UpdateImageRequest) SetSupported(v *UpdateImageRequestSupported) *UpdateImageRequest {
	s.Supported = v
	return s
}

func (s *UpdateImageRequest) Validate() error {
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

type UpdateImageRequestBuildConfig struct {
	// The build type.
	//
	// example:
	//
	// PackageInstallation
	BuildType *string `json:"BuildType,omitempty" xml:"BuildType,omitempty"`
	// The list of pre-installation scripts.
	PackageInstallationScripts []*UpdateImageRequestBuildConfigPackageInstallationScripts `json:"PackageInstallationScripts,omitempty" xml:"PackageInstallationScripts,omitempty" type:"Repeated"`
}

func (s UpdateImageRequestBuildConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageRequestBuildConfig) GoString() string {
	return s.String()
}

func (s *UpdateImageRequestBuildConfig) GetBuildType() *string {
	return s.BuildType
}

func (s *UpdateImageRequestBuildConfig) GetPackageInstallationScripts() []*UpdateImageRequestBuildConfigPackageInstallationScripts {
	return s.PackageInstallationScripts
}

func (s *UpdateImageRequestBuildConfig) SetBuildType(v string) *UpdateImageRequestBuildConfig {
	s.BuildType = &v
	return s
}

func (s *UpdateImageRequestBuildConfig) SetPackageInstallationScripts(v []*UpdateImageRequestBuildConfigPackageInstallationScripts) *UpdateImageRequestBuildConfig {
	s.PackageInstallationScripts = v
	return s
}

func (s *UpdateImageRequestBuildConfig) Validate() error {
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

type UpdateImageRequestBuildConfigPackageInstallationScripts struct {
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

func (s UpdateImageRequestBuildConfigPackageInstallationScripts) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageRequestBuildConfigPackageInstallationScripts) GoString() string {
	return s.String()
}

func (s *UpdateImageRequestBuildConfigPackageInstallationScripts) GetContent() *string {
	return s.Content
}

func (s *UpdateImageRequestBuildConfigPackageInstallationScripts) GetType() *string {
	return s.Type
}

func (s *UpdateImageRequestBuildConfigPackageInstallationScripts) SetContent(v string) *UpdateImageRequestBuildConfigPackageInstallationScripts {
	s.Content = &v
	return s
}

func (s *UpdateImageRequestBuildConfigPackageInstallationScripts) SetType(v string) *UpdateImageRequestBuildConfigPackageInstallationScripts {
	s.Type = &v
	return s
}

func (s *UpdateImageRequestBuildConfigPackageInstallationScripts) Validate() error {
	return dara.Validate(s)
}

type UpdateImageRequestSupported struct {
	// The image sub-module. Valid values:
	//
	// - Scheduler: data development.
	//
	// example:
	//
	// Scheduler
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The list of node types supported by the image.
	TaskTypes []*string `json:"TaskTypes,omitempty" xml:"TaskTypes,omitempty" type:"Repeated"`
}

func (s UpdateImageRequestSupported) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageRequestSupported) GoString() string {
	return s.String()
}

func (s *UpdateImageRequestSupported) GetModule() *string {
	return s.Module
}

func (s *UpdateImageRequestSupported) GetTaskTypes() []*string {
	return s.TaskTypes
}

func (s *UpdateImageRequestSupported) SetModule(v string) *UpdateImageRequestSupported {
	s.Module = &v
	return s
}

func (s *UpdateImageRequestSupported) SetTaskTypes(v []*string) *UpdateImageRequestSupported {
	s.TaskTypes = v
	return s
}

func (s *UpdateImageRequestSupported) Validate() error {
	return dara.Validate(s)
}

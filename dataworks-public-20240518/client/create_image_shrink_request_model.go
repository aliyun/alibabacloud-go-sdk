// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateImageShrinkRequest
	GetAccessibility() *string
	SetAcrAssociatedVpcId(v string) *CreateImageShrinkRequest
	GetAcrAssociatedVpcId() *string
	SetAcrInstanceId(v string) *CreateImageShrinkRequest
	GetAcrInstanceId() *string
	SetBuildConfigShrink(v string) *CreateImageShrinkRequest
	GetBuildConfigShrink() *string
	SetClientToken(v string) *CreateImageShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateImageShrinkRequest
	GetDescription() *string
	SetEnableSyncMaxCompute(v bool) *CreateImageShrinkRequest
	GetEnableSyncMaxCompute() *bool
	SetImageUri(v string) *CreateImageShrinkRequest
	GetImageUri() *string
	SetName(v string) *CreateImageShrinkRequest
	GetName() *string
	SetNamespace(v string) *CreateImageShrinkRequest
	GetNamespace() *string
	SetProviderImageId(v string) *CreateImageShrinkRequest
	GetProviderImageId() *string
	SetProviderType(v string) *CreateImageShrinkRequest
	GetProviderType() *string
	SetRepositoryName(v string) *CreateImageShrinkRequest
	GetRepositoryName() *string
	SetSupportedShrink(v string) *CreateImageShrinkRequest
	GetSupportedShrink() *string
}

type CreateImageShrinkRequest struct {
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
	BuildConfigShrink *string `json:"BuildConfig,omitempty" xml:"BuildConfig,omitempty"`
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
	SupportedShrink *string `json:"Supported,omitempty" xml:"Supported,omitempty"`
}

func (s CreateImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateImageShrinkRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateImageShrinkRequest) GetAcrAssociatedVpcId() *string {
	return s.AcrAssociatedVpcId
}

func (s *CreateImageShrinkRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CreateImageShrinkRequest) GetBuildConfigShrink() *string {
	return s.BuildConfigShrink
}

func (s *CreateImageShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateImageShrinkRequest) GetEnableSyncMaxCompute() *bool {
	return s.EnableSyncMaxCompute
}

func (s *CreateImageShrinkRequest) GetImageUri() *string {
	return s.ImageUri
}

func (s *CreateImageShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateImageShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateImageShrinkRequest) GetProviderImageId() *string {
	return s.ProviderImageId
}

func (s *CreateImageShrinkRequest) GetProviderType() *string {
	return s.ProviderType
}

func (s *CreateImageShrinkRequest) GetRepositoryName() *string {
	return s.RepositoryName
}

func (s *CreateImageShrinkRequest) GetSupportedShrink() *string {
	return s.SupportedShrink
}

func (s *CreateImageShrinkRequest) SetAccessibility(v string) *CreateImageShrinkRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateImageShrinkRequest) SetAcrAssociatedVpcId(v string) *CreateImageShrinkRequest {
	s.AcrAssociatedVpcId = &v
	return s
}

func (s *CreateImageShrinkRequest) SetAcrInstanceId(v string) *CreateImageShrinkRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateImageShrinkRequest) SetBuildConfigShrink(v string) *CreateImageShrinkRequest {
	s.BuildConfigShrink = &v
	return s
}

func (s *CreateImageShrinkRequest) SetClientToken(v string) *CreateImageShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageShrinkRequest) SetDescription(v string) *CreateImageShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateImageShrinkRequest) SetEnableSyncMaxCompute(v bool) *CreateImageShrinkRequest {
	s.EnableSyncMaxCompute = &v
	return s
}

func (s *CreateImageShrinkRequest) SetImageUri(v string) *CreateImageShrinkRequest {
	s.ImageUri = &v
	return s
}

func (s *CreateImageShrinkRequest) SetName(v string) *CreateImageShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateImageShrinkRequest) SetNamespace(v string) *CreateImageShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *CreateImageShrinkRequest) SetProviderImageId(v string) *CreateImageShrinkRequest {
	s.ProviderImageId = &v
	return s
}

func (s *CreateImageShrinkRequest) SetProviderType(v string) *CreateImageShrinkRequest {
	s.ProviderType = &v
	return s
}

func (s *CreateImageShrinkRequest) SetRepositoryName(v string) *CreateImageShrinkRequest {
	s.RepositoryName = &v
	return s
}

func (s *CreateImageShrinkRequest) SetSupportedShrink(v string) *CreateImageShrinkRequest {
	s.SupportedShrink = &v
	return s
}

func (s *CreateImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}

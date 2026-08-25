// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *UpdateImageShrinkRequest
	GetAccessibility() *string
	SetAcrAssociatedVpcId(v string) *UpdateImageShrinkRequest
	GetAcrAssociatedVpcId() *string
	SetAcrInstanceId(v string) *UpdateImageShrinkRequest
	GetAcrInstanceId() *string
	SetBuildConfigShrink(v string) *UpdateImageShrinkRequest
	GetBuildConfigShrink() *string
	SetDescription(v string) *UpdateImageShrinkRequest
	GetDescription() *string
	SetId(v string) *UpdateImageShrinkRequest
	GetId() *string
	SetImageUri(v string) *UpdateImageShrinkRequest
	GetImageUri() *string
	SetName(v string) *UpdateImageShrinkRequest
	GetName() *string
	SetNamespace(v string) *UpdateImageShrinkRequest
	GetNamespace() *string
	SetProviderImageId(v string) *UpdateImageShrinkRequest
	GetProviderImageId() *string
	SetRepositoryName(v string) *UpdateImageShrinkRequest
	GetRepositoryName() *string
	SetSupportedShrink(v string) *UpdateImageShrinkRequest
	GetSupportedShrink() *string
}

type UpdateImageShrinkRequest struct {
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
	BuildConfigShrink *string `json:"BuildConfig,omitempty" xml:"BuildConfig,omitempty"`
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
	SupportedShrink *string `json:"Supported,omitempty" xml:"Supported,omitempty"`
}

func (s UpdateImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageShrinkRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *UpdateImageShrinkRequest) GetAcrAssociatedVpcId() *string {
	return s.AcrAssociatedVpcId
}

func (s *UpdateImageShrinkRequest) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *UpdateImageShrinkRequest) GetBuildConfigShrink() *string {
	return s.BuildConfigShrink
}

func (s *UpdateImageShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateImageShrinkRequest) GetId() *string {
	return s.Id
}

func (s *UpdateImageShrinkRequest) GetImageUri() *string {
	return s.ImageUri
}

func (s *UpdateImageShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateImageShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdateImageShrinkRequest) GetProviderImageId() *string {
	return s.ProviderImageId
}

func (s *UpdateImageShrinkRequest) GetRepositoryName() *string {
	return s.RepositoryName
}

func (s *UpdateImageShrinkRequest) GetSupportedShrink() *string {
	return s.SupportedShrink
}

func (s *UpdateImageShrinkRequest) SetAccessibility(v string) *UpdateImageShrinkRequest {
	s.Accessibility = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetAcrAssociatedVpcId(v string) *UpdateImageShrinkRequest {
	s.AcrAssociatedVpcId = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetAcrInstanceId(v string) *UpdateImageShrinkRequest {
	s.AcrInstanceId = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetBuildConfigShrink(v string) *UpdateImageShrinkRequest {
	s.BuildConfigShrink = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetDescription(v string) *UpdateImageShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetId(v string) *UpdateImageShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetImageUri(v string) *UpdateImageShrinkRequest {
	s.ImageUri = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetName(v string) *UpdateImageShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetNamespace(v string) *UpdateImageShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetProviderImageId(v string) *UpdateImageShrinkRequest {
	s.ProviderImageId = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetRepositoryName(v string) *UpdateImageShrinkRequest {
	s.RepositoryName = &v
	return s
}

func (s *UpdateImageShrinkRequest) SetSupportedShrink(v string) *UpdateImageShrinkRequest {
	s.SupportedShrink = &v
	return s
}

func (s *UpdateImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}

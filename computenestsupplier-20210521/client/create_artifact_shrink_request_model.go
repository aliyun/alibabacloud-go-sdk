// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateArtifactShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactBuildPropertyShrink(v string) *CreateArtifactShrinkRequest
	GetArtifactBuildPropertyShrink() *string
	SetArtifactBuildType(v string) *CreateArtifactShrinkRequest
	GetArtifactBuildType() *string
	SetArtifactId(v string) *CreateArtifactShrinkRequest
	GetArtifactId() *string
	SetArtifactPropertyShrink(v string) *CreateArtifactShrinkRequest
	GetArtifactPropertyShrink() *string
	SetArtifactType(v string) *CreateArtifactShrinkRequest
	GetArtifactType() *string
	SetClientToken(v string) *CreateArtifactShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateArtifactShrinkRequest
	GetDescription() *string
	SetName(v string) *CreateArtifactShrinkRequest
	GetName() *string
	SetResourceGroupId(v string) *CreateArtifactShrinkRequest
	GetResourceGroupId() *string
	SetSupportRegionIds(v []*string) *CreateArtifactShrinkRequest
	GetSupportRegionIds() []*string
	SetTag(v []*CreateArtifactShrinkRequestTag) *CreateArtifactShrinkRequest
	GetTag() []*CreateArtifactShrinkRequestTag
	SetVersionName(v string) *CreateArtifactShrinkRequest
	GetVersionName() *string
}

type CreateArtifactShrinkRequest struct {
	// The content used to build the artifact. This parameter is used for managed artifact builds.
	ArtifactBuildPropertyShrink *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The type of the artifact to be built. Valid values:
	//
	// - EcsImage: builds an ECS image.
	//
	// - Dockerfile: builds a container image based on a Dockerfile.
	//
	// - Buildpacks: builds a container image based on Buildpacks.
	//
	// - ContainerImage: builds a container image by renaming an existing container image.
	//
	// example:
	//
	// Dockerflie
	ArtifactBuildType *string `json:"ArtifactBuildType,omitempty" xml:"ArtifactBuildType,omitempty"`
	// The artifact ID.
	//
	// This parameter is required to create a new version of an existing artifact.
	//
	// You can call the [ListArtifacts](https://help.aliyun.com/document_detail/469993.html) operation to obtain the artifact ID.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43ae****
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The content of the artifact.
	ArtifactPropertyShrink *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// The artifact type.
	//
	// Valid values:
	//
	// - EcsImage: an ECS image artifact.
	//
	// - AcrImage: a container image artifact.
	//
	// - File: an Object Storage Service (OSS) file artifact.
	//
	// - Script: a script artifact.
	//
	// - HelmChart: a Helm chart artifact.
	//
	// This parameter is required.
	//
	// example:
	//
	// EcsImage
	ArtifactType *string `json:"ArtifactType,omitempty" xml:"ArtifactType,omitempty"`
	// Ensures the idempotence of the request.
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
	// The artifact name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Redhat8_5 image
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm2jfvb7b****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The regions where the image can be distributed.
	SupportRegionIds []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// The custom tags.
	Tag []*CreateArtifactShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the artifact version.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateArtifactShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateArtifactShrinkRequest) GetArtifactBuildPropertyShrink() *string {
	return s.ArtifactBuildPropertyShrink
}

func (s *CreateArtifactShrinkRequest) GetArtifactBuildType() *string {
	return s.ArtifactBuildType
}

func (s *CreateArtifactShrinkRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *CreateArtifactShrinkRequest) GetArtifactPropertyShrink() *string {
	return s.ArtifactPropertyShrink
}

func (s *CreateArtifactShrinkRequest) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *CreateArtifactShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateArtifactShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateArtifactShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateArtifactShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateArtifactShrinkRequest) GetSupportRegionIds() []*string {
	return s.SupportRegionIds
}

func (s *CreateArtifactShrinkRequest) GetTag() []*CreateArtifactShrinkRequestTag {
	return s.Tag
}

func (s *CreateArtifactShrinkRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateArtifactShrinkRequest) SetArtifactBuildPropertyShrink(v string) *CreateArtifactShrinkRequest {
	s.ArtifactBuildPropertyShrink = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactBuildType(v string) *CreateArtifactShrinkRequest {
	s.ArtifactBuildType = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactId(v string) *CreateArtifactShrinkRequest {
	s.ArtifactId = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactPropertyShrink(v string) *CreateArtifactShrinkRequest {
	s.ArtifactPropertyShrink = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetArtifactType(v string) *CreateArtifactShrinkRequest {
	s.ArtifactType = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetClientToken(v string) *CreateArtifactShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetDescription(v string) *CreateArtifactShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetName(v string) *CreateArtifactShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetResourceGroupId(v string) *CreateArtifactShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateArtifactShrinkRequest) SetSupportRegionIds(v []*string) *CreateArtifactShrinkRequest {
	s.SupportRegionIds = v
	return s
}

func (s *CreateArtifactShrinkRequest) SetTag(v []*CreateArtifactShrinkRequestTag) *CreateArtifactShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateArtifactShrinkRequest) SetVersionName(v string) *CreateArtifactShrinkRequest {
	s.VersionName = &v
	return s
}

func (s *CreateArtifactShrinkRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateArtifactShrinkRequestTag struct {
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

func (s CreateArtifactShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateArtifactShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateArtifactShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateArtifactShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateArtifactShrinkRequestTag) SetKey(v string) *CreateArtifactShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateArtifactShrinkRequestTag) SetValue(v string) *CreateArtifactShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateArtifactShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}

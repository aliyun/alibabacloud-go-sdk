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
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	ArtifactBuildPropertyShrink *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
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
	ArtifactPropertyShrink *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
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
	Tag []*CreateArtifactShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The version name of the deployment package.
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
	return dara.Validate(s)
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

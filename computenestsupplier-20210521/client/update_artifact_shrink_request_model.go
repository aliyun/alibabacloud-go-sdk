// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateArtifactShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactBuildPropertyShrink(v string) *UpdateArtifactShrinkRequest
	GetArtifactBuildPropertyShrink() *string
	SetArtifactId(v string) *UpdateArtifactShrinkRequest
	GetArtifactId() *string
	SetArtifactPropertyShrink(v string) *UpdateArtifactShrinkRequest
	GetArtifactPropertyShrink() *string
	SetClientToken(v string) *UpdateArtifactShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateArtifactShrinkRequest
	GetDescription() *string
	SetPermissionType(v string) *UpdateArtifactShrinkRequest
	GetPermissionType() *string
	SetSupportRegionIds(v []*string) *UpdateArtifactShrinkRequest
	GetSupportRegionIds() []*string
	SetVersionName(v string) *UpdateArtifactShrinkRequest
	GetVersionName() *string
}

type UpdateArtifactShrinkRequest struct {
	// The properties for building the artifact. This is used for managed artifact builds.
	ArtifactBuildPropertyShrink *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The ID of the artifact.
	//
	// To obtain the artifact ID, call the [ListArtifacts](https://help.aliyun.com/document_detail/469993.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43ae****
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the artifact.
	ArtifactPropertyShrink *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
	// A client token to ensure the idempotence of the request. Generate a unique token for each request from your client. The **ClientToken*	- can contain only ASCII characters and must be no more than 64 characters long.
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
	// The permission type. This parameter is valid for container image artifacts and Helm Chart artifacts. The value can be changed only from \\`Automatic\\` to \\`Public\\`. Valid values:
	//
	// - Public
	//
	// - Automatic
	//
	// example:
	//
	// Public
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The IDs of regions to which the image can be distributed.
	SupportRegionIds []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// The name of the artifact version.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s UpdateArtifactShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateArtifactShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateArtifactShrinkRequest) GetArtifactBuildPropertyShrink() *string {
	return s.ArtifactBuildPropertyShrink
}

func (s *UpdateArtifactShrinkRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *UpdateArtifactShrinkRequest) GetArtifactPropertyShrink() *string {
	return s.ArtifactPropertyShrink
}

func (s *UpdateArtifactShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateArtifactShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateArtifactShrinkRequest) GetPermissionType() *string {
	return s.PermissionType
}

func (s *UpdateArtifactShrinkRequest) GetSupportRegionIds() []*string {
	return s.SupportRegionIds
}

func (s *UpdateArtifactShrinkRequest) GetVersionName() *string {
	return s.VersionName
}

func (s *UpdateArtifactShrinkRequest) SetArtifactBuildPropertyShrink(v string) *UpdateArtifactShrinkRequest {
	s.ArtifactBuildPropertyShrink = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetArtifactId(v string) *UpdateArtifactShrinkRequest {
	s.ArtifactId = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetArtifactPropertyShrink(v string) *UpdateArtifactShrinkRequest {
	s.ArtifactPropertyShrink = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetClientToken(v string) *UpdateArtifactShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetDescription(v string) *UpdateArtifactShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetPermissionType(v string) *UpdateArtifactShrinkRequest {
	s.PermissionType = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetSupportRegionIds(v []*string) *UpdateArtifactShrinkRequest {
	s.SupportRegionIds = v
	return s
}

func (s *UpdateArtifactShrinkRequest) SetVersionName(v string) *UpdateArtifactShrinkRequest {
	s.VersionName = &v
	return s
}

func (s *UpdateArtifactShrinkRequest) Validate() error {
	return dara.Validate(s)
}

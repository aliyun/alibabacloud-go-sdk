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
	// The build properties of the artifact, utilized for hosting and building the deployment package.
	ArtifactBuildPropertyShrink *string `json:"ArtifactBuildProperty,omitempty" xml:"ArtifactBuildProperty,omitempty"`
	// The ID of the deployment package.
	//
	// This parameter is required.
	//
	// example:
	//
	// artifact-eea08d1e2d3a43aexxxx
	ArtifactId *string `json:"ArtifactId,omitempty" xml:"ArtifactId,omitempty"`
	// The properties of the deployment package.
	ArtifactPropertyShrink *string `json:"ArtifactProperty,omitempty" xml:"ArtifactProperty,omitempty"`
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
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Permission fields are applicable to container image artifact and Helm Chart artifact. They can only change from Automatic to Public. Options:
	//
	// Public
	//
	// Automatic
	//
	// example:
	//
	// Public
	PermissionType *string `json:"PermissionType,omitempty" xml:"PermissionType,omitempty"`
	// The IDs of the regions that support the deployment package.
	SupportRegionIds []*string `json:"SupportRegionIds,omitempty" xml:"SupportRegionIds,omitempty" type:"Repeated"`
	// The version name of the deployment package.
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

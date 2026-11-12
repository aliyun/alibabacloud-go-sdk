// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifactRelease interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactRef(v string) *ArtifactRelease
	GetArtifactRef() *string
	SetCreatedAt(v string) *ArtifactRelease
	GetCreatedAt() *string
	SetDescription(v string) *ArtifactRelease
	GetDescription() *string
	SetImage(v string) *ArtifactRelease
	GetImage() *string
	SetType(v string) *ArtifactRelease
	GetType() *string
	SetVersion(v string) *ArtifactRelease
	GetVersion() *string
}

type ArtifactRelease struct {
	// The artifact reference name.
	//
	// example:
	//
	// python-env
	ArtifactRef *string `json:"ArtifactRef,omitempty" xml:"ArtifactRef,omitempty"`
	// The creation time of the artifact release record.
	//
	// example:
	//
	// 2024-01-15T10:30:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The artifact description.
	//
	// example:
	//
	// Container image for model inference
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The container image address used by the artifact.
	//
	// example:
	//
	// registry.cn-shanghai.aliyuncs.com/eas/echo_cn-shanghai:v0.0.1-20210129111320
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The artifact type.
	//
	// example:
	//
	// Image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The artifact version.
	//
	// example:
	//
	// V2.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ArtifactRelease) String() string {
	return dara.Prettify(s)
}

func (s ArtifactRelease) GoString() string {
	return s.String()
}

func (s *ArtifactRelease) GetArtifactRef() *string {
	return s.ArtifactRef
}

func (s *ArtifactRelease) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ArtifactRelease) GetDescription() *string {
	return s.Description
}

func (s *ArtifactRelease) GetImage() *string {
	return s.Image
}

func (s *ArtifactRelease) GetType() *string {
	return s.Type
}

func (s *ArtifactRelease) GetVersion() *string {
	return s.Version
}

func (s *ArtifactRelease) SetArtifactRef(v string) *ArtifactRelease {
	s.ArtifactRef = &v
	return s
}

func (s *ArtifactRelease) SetCreatedAt(v string) *ArtifactRelease {
	s.CreatedAt = &v
	return s
}

func (s *ArtifactRelease) SetDescription(v string) *ArtifactRelease {
	s.Description = &v
	return s
}

func (s *ArtifactRelease) SetImage(v string) *ArtifactRelease {
	s.Image = &v
	return s
}

func (s *ArtifactRelease) SetType(v string) *ArtifactRelease {
	s.Type = &v
	return s
}

func (s *ArtifactRelease) SetVersion(v string) *ArtifactRelease {
	s.Version = &v
	return s
}

func (s *ArtifactRelease) Validate() error {
	return dara.Validate(s)
}

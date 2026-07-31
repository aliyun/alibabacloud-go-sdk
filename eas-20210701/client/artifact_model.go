// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifact interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedAt(v string) *Artifact
	GetCreatedAt() *string
	SetName(v string) *Artifact
	GetName() *string
	SetType(v string) *Artifact
	GetType() *string
	SetVersions(v []*ArtifactVersions) *Artifact
	GetVersions() []*ArtifactVersions
}

type Artifact struct {
	// The creation time.
	//
	// example:
	//
	// 2024-01-15T10:30:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The artifact name.
	//
	// example:
	//
	// foo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The artifact type.
	//
	// example:
	//
	// Image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version list.
	Versions []*ArtifactVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
}

func (s Artifact) String() string {
	return dara.Prettify(s)
}

func (s Artifact) GoString() string {
	return s.String()
}

func (s *Artifact) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *Artifact) GetName() *string {
	return s.Name
}

func (s *Artifact) GetType() *string {
	return s.Type
}

func (s *Artifact) GetVersions() []*ArtifactVersions {
	return s.Versions
}

func (s *Artifact) SetCreatedAt(v string) *Artifact {
	s.CreatedAt = &v
	return s
}

func (s *Artifact) SetName(v string) *Artifact {
	s.Name = &v
	return s
}

func (s *Artifact) SetType(v string) *Artifact {
	s.Type = &v
	return s
}

func (s *Artifact) SetVersions(v []*ArtifactVersions) *Artifact {
	s.Versions = v
	return s
}

func (s *Artifact) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ArtifactVersions struct {
	// The version alias.
	//
	// example:
	//
	// foo
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The template description associated with the version.
	//
	// example:
	//
	// Supports new xx feature
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The version name.
	//
	// example:
	//
	// ali-ahpa-hz
	ReleaseName *string `json:"ReleaseName,omitempty" xml:"ReleaseName,omitempty"`
	// The version number.
	//
	// example:
	//
	// V2.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ArtifactVersions) String() string {
	return dara.Prettify(s)
}

func (s ArtifactVersions) GoString() string {
	return s.String()
}

func (s *ArtifactVersions) GetAlias() *string {
	return s.Alias
}

func (s *ArtifactVersions) GetDescription() *string {
	return s.Description
}

func (s *ArtifactVersions) GetReleaseName() *string {
	return s.ReleaseName
}

func (s *ArtifactVersions) GetVersion() *string {
	return s.Version
}

func (s *ArtifactVersions) SetAlias(v string) *ArtifactVersions {
	s.Alias = &v
	return s
}

func (s *ArtifactVersions) SetDescription(v string) *ArtifactVersions {
	s.Description = &v
	return s
}

func (s *ArtifactVersions) SetReleaseName(v string) *ArtifactVersions {
	s.ReleaseName = &v
	return s
}

func (s *ArtifactVersions) SetVersion(v string) *ArtifactVersions {
	s.Version = &v
	return s
}

func (s *ArtifactVersions) Validate() error {
	return dara.Validate(s)
}

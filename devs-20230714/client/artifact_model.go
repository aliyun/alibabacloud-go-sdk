// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iArtifact interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *Artifact
	GetCreatedTime() *string
	SetDescription(v string) *Artifact
	GetDescription() *string
	SetGeneration(v int32) *Artifact
	GetGeneration() *int32
	SetKind(v string) *Artifact
	GetKind() *string
	SetLabels(v map[string]*string) *Artifact
	GetLabels() map[string]*string
	SetName(v string) *Artifact
	GetName() *string
	SetResourceVersion(v int32) *Artifact
	GetResourceVersion() *int32
	SetSpec(v *ArtifactSpec) *Artifact
	GetSpec() *ArtifactSpec
	SetStatus(v *ArtifactStatus) *Artifact
	GetStatus() *ArtifactStatus
	SetUid(v string) *Artifact
	GetUid() *string
	SetUpdatedTime(v string) *Artifact
	GetUpdatedTime() *string
}

type Artifact struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// test-description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// Artifact
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-artifact
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32          `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *ArtifactSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status          *ArtifactStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	UpdatedTime *string `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
}

func (s Artifact) String() string {
	return dara.Prettify(s)
}

func (s Artifact) GoString() string {
	return s.String()
}

func (s *Artifact) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Artifact) GetDescription() *string {
	return s.Description
}

func (s *Artifact) GetGeneration() *int32 {
	return s.Generation
}

func (s *Artifact) GetKind() *string {
	return s.Kind
}

func (s *Artifact) GetLabels() map[string]*string {
	return s.Labels
}

func (s *Artifact) GetName() *string {
	return s.Name
}

func (s *Artifact) GetResourceVersion() *int32 {
	return s.ResourceVersion
}

func (s *Artifact) GetSpec() *ArtifactSpec {
	return s.Spec
}

func (s *Artifact) GetStatus() *ArtifactStatus {
	return s.Status
}

func (s *Artifact) GetUid() *string {
	return s.Uid
}

func (s *Artifact) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *Artifact) SetCreatedTime(v string) *Artifact {
	s.CreatedTime = &v
	return s
}

func (s *Artifact) SetDescription(v string) *Artifact {
	s.Description = &v
	return s
}

func (s *Artifact) SetGeneration(v int32) *Artifact {
	s.Generation = &v
	return s
}

func (s *Artifact) SetKind(v string) *Artifact {
	s.Kind = &v
	return s
}

func (s *Artifact) SetLabels(v map[string]*string) *Artifact {
	s.Labels = v
	return s
}

func (s *Artifact) SetName(v string) *Artifact {
	s.Name = &v
	return s
}

func (s *Artifact) SetResourceVersion(v int32) *Artifact {
	s.ResourceVersion = &v
	return s
}

func (s *Artifact) SetSpec(v *ArtifactSpec) *Artifact {
	s.Spec = v
	return s
}

func (s *Artifact) SetStatus(v *ArtifactStatus) *Artifact {
	s.Status = v
	return s
}

func (s *Artifact) SetUid(v string) *Artifact {
	s.Uid = &v
	return s
}

func (s *Artifact) SetUpdatedTime(v string) *Artifact {
	s.UpdatedTime = &v
	return s
}

func (s *Artifact) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUdfArtifact interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *UdfArtifact
	GetArtifactType() *string
	SetCreatedAt(v int64) *UdfArtifact
	GetCreatedAt() *int64
	SetCreator(v string) *UdfArtifact
	GetCreator() *string
	SetDependencyJarUris(v []*string) *UdfArtifact
	GetDependencyJarUris() []*string
	SetJarUrl(v string) *UdfArtifact
	GetJarUrl() *string
	SetModifiedAt(v int64) *UdfArtifact
	GetModifiedAt() *int64
	SetName(v string) *UdfArtifact
	GetName() *string
	SetNamespace(v string) *UdfArtifact
	GetNamespace() *string
	SetUdfClasses(v []*UdfClass) *UdfArtifact
	GetUdfClasses() []*UdfClass
}

type UdfArtifact struct {
	ArtifactType      *string     `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	CreatedAt         *int64      `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	Creator           *string     `json:"creator,omitempty" xml:"creator,omitempty"`
	DependencyJarUris []*string   `json:"dependencyJarUris,omitempty" xml:"dependencyJarUris,omitempty" type:"Repeated"`
	JarUrl            *string     `json:"jarUrl,omitempty" xml:"jarUrl,omitempty"`
	ModifiedAt        *int64      `json:"modifiedAt,omitempty" xml:"modifiedAt,omitempty"`
	Name              *string     `json:"name,omitempty" xml:"name,omitempty"`
	Namespace         *string     `json:"namespace,omitempty" xml:"namespace,omitempty"`
	UdfClasses        []*UdfClass `json:"udfClasses,omitempty" xml:"udfClasses,omitempty" type:"Repeated"`
}

func (s UdfArtifact) String() string {
	return dara.Prettify(s)
}

func (s UdfArtifact) GoString() string {
	return s.String()
}

func (s *UdfArtifact) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *UdfArtifact) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *UdfArtifact) GetCreator() *string {
	return s.Creator
}

func (s *UdfArtifact) GetDependencyJarUris() []*string {
	return s.DependencyJarUris
}

func (s *UdfArtifact) GetJarUrl() *string {
	return s.JarUrl
}

func (s *UdfArtifact) GetModifiedAt() *int64 {
	return s.ModifiedAt
}

func (s *UdfArtifact) GetName() *string {
	return s.Name
}

func (s *UdfArtifact) GetNamespace() *string {
	return s.Namespace
}

func (s *UdfArtifact) GetUdfClasses() []*UdfClass {
	return s.UdfClasses
}

func (s *UdfArtifact) SetArtifactType(v string) *UdfArtifact {
	s.ArtifactType = &v
	return s
}

func (s *UdfArtifact) SetCreatedAt(v int64) *UdfArtifact {
	s.CreatedAt = &v
	return s
}

func (s *UdfArtifact) SetCreator(v string) *UdfArtifact {
	s.Creator = &v
	return s
}

func (s *UdfArtifact) SetDependencyJarUris(v []*string) *UdfArtifact {
	s.DependencyJarUris = v
	return s
}

func (s *UdfArtifact) SetJarUrl(v string) *UdfArtifact {
	s.JarUrl = &v
	return s
}

func (s *UdfArtifact) SetModifiedAt(v int64) *UdfArtifact {
	s.ModifiedAt = &v
	return s
}

func (s *UdfArtifact) SetName(v string) *UdfArtifact {
	s.Name = &v
	return s
}

func (s *UdfArtifact) SetNamespace(v string) *UdfArtifact {
	s.Namespace = &v
	return s
}

func (s *UdfArtifact) SetUdfClasses(v []*UdfClass) *UdfArtifact {
	s.UdfClasses = v
	return s
}

func (s *UdfArtifact) Validate() error {
	if s.UdfClasses != nil {
		for _, item := range s.UdfClasses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

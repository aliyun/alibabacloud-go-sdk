// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRepository interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *Repository
	GetCreatedTime() *string
	SetDescription(v string) *Repository
	GetDescription() *string
	SetKind(v string) *Repository
	GetKind() *string
	SetLabels(v map[string]*string) *Repository
	GetLabels() map[string]*string
	SetName(v string) *Repository
	GetName() *string
	SetSpec(v *RepositorySpec) *Repository
	GetSpec() *RepositorySpec
	SetUid(v string) *Repository
	GetUid() *string
}

type Repository struct {
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
	// Repository
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-repository
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	Spec *RepositorySpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Repository) String() string {
	return dara.Prettify(s)
}

func (s Repository) GoString() string {
	return s.String()
}

func (s *Repository) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Repository) GetDescription() *string {
	return s.Description
}

func (s *Repository) GetKind() *string {
	return s.Kind
}

func (s *Repository) GetLabels() map[string]*string {
	return s.Labels
}

func (s *Repository) GetName() *string {
	return s.Name
}

func (s *Repository) GetSpec() *RepositorySpec {
	return s.Spec
}

func (s *Repository) GetUid() *string {
	return s.Uid
}

func (s *Repository) SetCreatedTime(v string) *Repository {
	s.CreatedTime = &v
	return s
}

func (s *Repository) SetDescription(v string) *Repository {
	s.Description = &v
	return s
}

func (s *Repository) SetKind(v string) *Repository {
	s.Kind = &v
	return s
}

func (s *Repository) SetLabels(v map[string]*string) *Repository {
	s.Labels = v
	return s
}

func (s *Repository) SetName(v string) *Repository {
	s.Name = &v
	return s
}

func (s *Repository) SetSpec(v *RepositorySpec) *Repository {
	s.Spec = v
	return s
}

func (s *Repository) SetUid(v string) *Repository {
	s.Uid = &v
	return s
}

func (s *Repository) Validate() error {
	return dara.Validate(s)
}

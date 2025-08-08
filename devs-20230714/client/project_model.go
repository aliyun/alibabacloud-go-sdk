// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProject interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *Project
	GetCreatedTime() *string
	SetDescription(v string) *Project
	GetDescription() *string
	SetKind(v string) *Project
	GetKind() *string
	SetLabels(v map[string]*string) *Project
	GetLabels() map[string]*string
	SetName(v string) *Project
	GetName() *string
	SetStatus(v *ProjectStatus) *Project
	GetStatus() *ProjectStatus
	SetUid(v string) *Project
	GetUid() *string
}

type Project struct {
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
	// Project
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-project
	Name   *string        `json:"name,omitempty" xml:"name,omitempty"`
	Status *ProjectStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Project) String() string {
	return dara.Prettify(s)
}

func (s Project) GoString() string {
	return s.String()
}

func (s *Project) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Project) GetDescription() *string {
	return s.Description
}

func (s *Project) GetKind() *string {
	return s.Kind
}

func (s *Project) GetLabels() map[string]*string {
	return s.Labels
}

func (s *Project) GetName() *string {
	return s.Name
}

func (s *Project) GetStatus() *ProjectStatus {
	return s.Status
}

func (s *Project) GetUid() *string {
	return s.Uid
}

func (s *Project) SetCreatedTime(v string) *Project {
	s.CreatedTime = &v
	return s
}

func (s *Project) SetDescription(v string) *Project {
	s.Description = &v
	return s
}

func (s *Project) SetKind(v string) *Project {
	s.Kind = &v
	return s
}

func (s *Project) SetLabels(v map[string]*string) *Project {
	s.Labels = v
	return s
}

func (s *Project) SetName(v string) *Project {
	s.Name = &v
	return s
}

func (s *Project) SetStatus(v *ProjectStatus) *Project {
	s.Status = v
	return s
}

func (s *Project) SetUid(v string) *Project {
	s.Uid = &v
	return s
}

func (s *Project) Validate() error {
	return dara.Validate(s)
}

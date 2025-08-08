// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *Template
	GetCreatedTime() *string
	SetDescription(v string) *Template
	GetDescription() *string
	SetKind(v string) *Template
	GetKind() *string
	SetLabels(v map[string]*string) *Template
	GetLabels() map[string]*string
	SetName(v string) *Template
	GetName() *string
	SetSpec(v *TemplateSpec) *Template
	GetSpec() *TemplateSpec
	SetStatus(v *TemplateStatus) *Template
	GetStatus() *TemplateStatus
	SetUid(v string) *Template
	GetUid() *string
}

type Template struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// It is a template
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// Template
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// example:
	//
	// demo-template
	Name   *string         `json:"name,omitempty" xml:"name,omitempty"`
	Spec   *TemplateSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status *TemplateStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Template) String() string {
	return dara.Prettify(s)
}

func (s Template) GoString() string {
	return s.String()
}

func (s *Template) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Template) GetDescription() *string {
	return s.Description
}

func (s *Template) GetKind() *string {
	return s.Kind
}

func (s *Template) GetLabels() map[string]*string {
	return s.Labels
}

func (s *Template) GetName() *string {
	return s.Name
}

func (s *Template) GetSpec() *TemplateSpec {
	return s.Spec
}

func (s *Template) GetStatus() *TemplateStatus {
	return s.Status
}

func (s *Template) GetUid() *string {
	return s.Uid
}

func (s *Template) SetCreatedTime(v string) *Template {
	s.CreatedTime = &v
	return s
}

func (s *Template) SetDescription(v string) *Template {
	s.Description = &v
	return s
}

func (s *Template) SetKind(v string) *Template {
	s.Kind = &v
	return s
}

func (s *Template) SetLabels(v map[string]*string) *Template {
	s.Labels = v
	return s
}

func (s *Template) SetName(v string) *Template {
	s.Name = &v
	return s
}

func (s *Template) SetSpec(v *TemplateSpec) *Template {
	s.Spec = v
	return s
}

func (s *Template) SetStatus(v *TemplateStatus) *Template {
	s.Status = v
	return s
}

func (s *Template) SetUid(v string) *Template {
	s.Uid = &v
	return s
}

func (s *Template) Validate() error {
	return dara.Validate(s)
}

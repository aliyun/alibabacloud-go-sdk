// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *TaskTemplate
	GetCreatedTime() *string
	SetDeletionTime(v string) *TaskTemplate
	GetDeletionTime() *string
	SetDescription(v string) *TaskTemplate
	GetDescription() *string
	SetGeneration(v int32) *TaskTemplate
	GetGeneration() *int32
	SetKind(v string) *TaskTemplate
	GetKind() *string
	SetLabels(v map[string]*string) *TaskTemplate
	GetLabels() map[string]*string
	SetName(v string) *TaskTemplate
	GetName() *string
	SetResourceVersion(v int32) *TaskTemplate
	GetResourceVersion() *int32
	SetSpec(v *TaskTemplateSpec) *TaskTemplate
	GetSpec() *TaskTemplateSpec
	SetUid(v string) *TaskTemplate
	GetUid() *string
}

type TaskTemplate struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// 2021-11-19T09:34:38Z
	DeletionTime *string `json:"deletionTime,omitempty" xml:"deletionTime,omitempty"`
	// example:
	//
	// TaskTemplate example.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// TaskTemplate
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-task-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32            `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *TaskTemplateSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s TaskTemplate) String() string {
	return dara.Prettify(s)
}

func (s TaskTemplate) GoString() string {
	return s.String()
}

func (s *TaskTemplate) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *TaskTemplate) GetDeletionTime() *string {
	return s.DeletionTime
}

func (s *TaskTemplate) GetDescription() *string {
	return s.Description
}

func (s *TaskTemplate) GetGeneration() *int32 {
	return s.Generation
}

func (s *TaskTemplate) GetKind() *string {
	return s.Kind
}

func (s *TaskTemplate) GetLabels() map[string]*string {
	return s.Labels
}

func (s *TaskTemplate) GetName() *string {
	return s.Name
}

func (s *TaskTemplate) GetResourceVersion() *int32 {
	return s.ResourceVersion
}

func (s *TaskTemplate) GetSpec() *TaskTemplateSpec {
	return s.Spec
}

func (s *TaskTemplate) GetUid() *string {
	return s.Uid
}

func (s *TaskTemplate) SetCreatedTime(v string) *TaskTemplate {
	s.CreatedTime = &v
	return s
}

func (s *TaskTemplate) SetDeletionTime(v string) *TaskTemplate {
	s.DeletionTime = &v
	return s
}

func (s *TaskTemplate) SetDescription(v string) *TaskTemplate {
	s.Description = &v
	return s
}

func (s *TaskTemplate) SetGeneration(v int32) *TaskTemplate {
	s.Generation = &v
	return s
}

func (s *TaskTemplate) SetKind(v string) *TaskTemplate {
	s.Kind = &v
	return s
}

func (s *TaskTemplate) SetLabels(v map[string]*string) *TaskTemplate {
	s.Labels = v
	return s
}

func (s *TaskTemplate) SetName(v string) *TaskTemplate {
	s.Name = &v
	return s
}

func (s *TaskTemplate) SetResourceVersion(v int32) *TaskTemplate {
	s.ResourceVersion = &v
	return s
}

func (s *TaskTemplate) SetSpec(v *TaskTemplateSpec) *TaskTemplate {
	s.Spec = v
	return s
}

func (s *TaskTemplate) SetUid(v string) *TaskTemplate {
	s.Uid = &v
	return s
}

func (s *TaskTemplate) Validate() error {
	return dara.Validate(s)
}

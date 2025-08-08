// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPipelineTemplate interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *PipelineTemplate
	GetCreatedTime() *string
	SetDeletionTime(v string) *PipelineTemplate
	GetDeletionTime() *string
	SetDescription(v string) *PipelineTemplate
	GetDescription() *string
	SetGeneration(v int32) *PipelineTemplate
	GetGeneration() *int32
	SetKind(v string) *PipelineTemplate
	GetKind() *string
	SetLabels(v map[string]*string) *PipelineTemplate
	GetLabels() map[string]*string
	SetName(v string) *PipelineTemplate
	GetName() *string
	SetResourceVersion(v int32) *PipelineTemplate
	GetResourceVersion() *int32
	SetSpec(v *PipelineTemplateSpec) *PipelineTemplate
	GetSpec() *PipelineTemplateSpec
	SetUid(v string) *PipelineTemplate
	GetUid() *string
}

type PipelineTemplate struct {
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
	// PipelineTemplate example.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// PipelineTemplate
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-pipeline-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32                `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *PipelineTemplateSpec `json:"spec,omitempty" xml:"spec,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s PipelineTemplate) String() string {
	return dara.Prettify(s)
}

func (s PipelineTemplate) GoString() string {
	return s.String()
}

func (s *PipelineTemplate) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *PipelineTemplate) GetDeletionTime() *string {
	return s.DeletionTime
}

func (s *PipelineTemplate) GetDescription() *string {
	return s.Description
}

func (s *PipelineTemplate) GetGeneration() *int32 {
	return s.Generation
}

func (s *PipelineTemplate) GetKind() *string {
	return s.Kind
}

func (s *PipelineTemplate) GetLabels() map[string]*string {
	return s.Labels
}

func (s *PipelineTemplate) GetName() *string {
	return s.Name
}

func (s *PipelineTemplate) GetResourceVersion() *int32 {
	return s.ResourceVersion
}

func (s *PipelineTemplate) GetSpec() *PipelineTemplateSpec {
	return s.Spec
}

func (s *PipelineTemplate) GetUid() *string {
	return s.Uid
}

func (s *PipelineTemplate) SetCreatedTime(v string) *PipelineTemplate {
	s.CreatedTime = &v
	return s
}

func (s *PipelineTemplate) SetDeletionTime(v string) *PipelineTemplate {
	s.DeletionTime = &v
	return s
}

func (s *PipelineTemplate) SetDescription(v string) *PipelineTemplate {
	s.Description = &v
	return s
}

func (s *PipelineTemplate) SetGeneration(v int32) *PipelineTemplate {
	s.Generation = &v
	return s
}

func (s *PipelineTemplate) SetKind(v string) *PipelineTemplate {
	s.Kind = &v
	return s
}

func (s *PipelineTemplate) SetLabels(v map[string]*string) *PipelineTemplate {
	s.Labels = v
	return s
}

func (s *PipelineTemplate) SetName(v string) *PipelineTemplate {
	s.Name = &v
	return s
}

func (s *PipelineTemplate) SetResourceVersion(v int32) *PipelineTemplate {
	s.ResourceVersion = &v
	return s
}

func (s *PipelineTemplate) SetSpec(v *PipelineTemplateSpec) *PipelineTemplate {
	s.Spec = v
	return s
}

func (s *PipelineTemplate) SetUid(v string) *PipelineTemplate {
	s.Uid = &v
	return s
}

func (s *PipelineTemplate) Validate() error {
	return dara.Validate(s)
}

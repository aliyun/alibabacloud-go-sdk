// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPipeline interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *Pipeline
	GetCreatedTime() *string
	SetDescription(v string) *Pipeline
	GetDescription() *string
	SetGeneration(v int32) *Pipeline
	GetGeneration() *int32
	SetKind(v string) *Pipeline
	GetKind() *string
	SetLabels(v map[string]*string) *Pipeline
	GetLabels() map[string]*string
	SetName(v string) *Pipeline
	GetName() *string
	SetResourceVersion(v int32) *Pipeline
	GetResourceVersion() *int32
	SetSpec(v *PipelineSpec) *Pipeline
	GetSpec() *PipelineSpec
	SetStatus(v *PipelineStatus) *Pipeline
	GetStatus() *PipelineStatus
	SetUid(v string) *Pipeline
	GetUid() *string
}

type Pipeline struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// Pipeline example.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// Pipeline
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-pipeline
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32          `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *PipelineSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status          *PipelineStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Pipeline) String() string {
	return dara.Prettify(s)
}

func (s Pipeline) GoString() string {
	return s.String()
}

func (s *Pipeline) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Pipeline) GetDescription() *string {
	return s.Description
}

func (s *Pipeline) GetGeneration() *int32 {
	return s.Generation
}

func (s *Pipeline) GetKind() *string {
	return s.Kind
}

func (s *Pipeline) GetLabels() map[string]*string {
	return s.Labels
}

func (s *Pipeline) GetName() *string {
	return s.Name
}

func (s *Pipeline) GetResourceVersion() *int32 {
	return s.ResourceVersion
}

func (s *Pipeline) GetSpec() *PipelineSpec {
	return s.Spec
}

func (s *Pipeline) GetStatus() *PipelineStatus {
	return s.Status
}

func (s *Pipeline) GetUid() *string {
	return s.Uid
}

func (s *Pipeline) SetCreatedTime(v string) *Pipeline {
	s.CreatedTime = &v
	return s
}

func (s *Pipeline) SetDescription(v string) *Pipeline {
	s.Description = &v
	return s
}

func (s *Pipeline) SetGeneration(v int32) *Pipeline {
	s.Generation = &v
	return s
}

func (s *Pipeline) SetKind(v string) *Pipeline {
	s.Kind = &v
	return s
}

func (s *Pipeline) SetLabels(v map[string]*string) *Pipeline {
	s.Labels = v
	return s
}

func (s *Pipeline) SetName(v string) *Pipeline {
	s.Name = &v
	return s
}

func (s *Pipeline) SetResourceVersion(v int32) *Pipeline {
	s.ResourceVersion = &v
	return s
}

func (s *Pipeline) SetSpec(v *PipelineSpec) *Pipeline {
	s.Spec = v
	return s
}

func (s *Pipeline) SetStatus(v *PipelineStatus) *Pipeline {
	s.Status = v
	return s
}

func (s *Pipeline) SetUid(v string) *Pipeline {
	s.Uid = &v
	return s
}

func (s *Pipeline) Validate() error {
	return dara.Validate(s)
}

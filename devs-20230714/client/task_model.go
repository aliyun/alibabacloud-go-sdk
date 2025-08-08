// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTask interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *Task
	GetCreatedTime() *string
	SetDescription(v string) *Task
	GetDescription() *string
	SetGeneration(v int32) *Task
	GetGeneration() *int32
	SetKind(v string) *Task
	GetKind() *string
	SetLabels(v map[string]*string) *Task
	GetLabels() map[string]*string
	SetName(v string) *Task
	GetName() *string
	SetResourceVersion(v int32) *Task
	GetResourceVersion() *int32
	SetSpec(v *TaskSpec) *Task
	GetSpec() *TaskSpec
	SetStatus(v *TaskStatus) *Task
	GetStatus() *TaskStatus
	SetUid(v string) *Task
	GetUid() *string
}

type Task struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// Task example.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1
	Generation *int32 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// Task
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-task
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	ResourceVersion *int32      `json:"resourceVersion,omitempty" xml:"resourceVersion,omitempty"`
	Spec            *TaskSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status          *TaskStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Task) String() string {
	return dara.Prettify(s)
}

func (s Task) GoString() string {
	return s.String()
}

func (s *Task) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Task) GetDescription() *string {
	return s.Description
}

func (s *Task) GetGeneration() *int32 {
	return s.Generation
}

func (s *Task) GetKind() *string {
	return s.Kind
}

func (s *Task) GetLabels() map[string]*string {
	return s.Labels
}

func (s *Task) GetName() *string {
	return s.Name
}

func (s *Task) GetResourceVersion() *int32 {
	return s.ResourceVersion
}

func (s *Task) GetSpec() *TaskSpec {
	return s.Spec
}

func (s *Task) GetStatus() *TaskStatus {
	return s.Status
}

func (s *Task) GetUid() *string {
	return s.Uid
}

func (s *Task) SetCreatedTime(v string) *Task {
	s.CreatedTime = &v
	return s
}

func (s *Task) SetDescription(v string) *Task {
	s.Description = &v
	return s
}

func (s *Task) SetGeneration(v int32) *Task {
	s.Generation = &v
	return s
}

func (s *Task) SetKind(v string) *Task {
	s.Kind = &v
	return s
}

func (s *Task) SetLabels(v map[string]*string) *Task {
	s.Labels = v
	return s
}

func (s *Task) SetName(v string) *Task {
	s.Name = &v
	return s
}

func (s *Task) SetResourceVersion(v int32) *Task {
	s.ResourceVersion = &v
	return s
}

func (s *Task) SetSpec(v *TaskSpec) *Task {
	s.Spec = v
	return s
}

func (s *Task) SetStatus(v *TaskStatus) *Task {
	s.Status = v
	return s
}

func (s *Task) SetUid(v string) *Task {
	s.Uid = &v
	return s
}

func (s *Task) Validate() error {
	return dara.Validate(s)
}

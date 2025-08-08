// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToolset interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *Toolset
	GetCreatedTime() *string
	SetDescription(v string) *Toolset
	GetDescription() *string
	SetGeneration(v int64) *Toolset
	GetGeneration() *int64
	SetKind(v string) *Toolset
	GetKind() *string
	SetLabels(v map[string]*string) *Toolset
	GetLabels() map[string]*string
	SetName(v string) *Toolset
	GetName() *string
	SetSpec(v *ToolsetSpec) *Toolset
	GetSpec() *ToolsetSpec
	SetStatus(v *ToolsetStatus) *Toolset
	GetStatus() *ToolsetStatus
	SetUid(v string) *Toolset
	GetUid() *string
}

type Toolset struct {
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
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// Toolset
	Kind   *string            `json:"kind,omitempty" xml:"kind,omitempty"`
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-toolset
	Name   *string        `json:"name,omitempty" xml:"name,omitempty"`
	Spec   *ToolsetSpec   `json:"spec,omitempty" xml:"spec,omitempty"`
	Status *ToolsetStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s Toolset) String() string {
	return dara.Prettify(s)
}

func (s Toolset) GoString() string {
	return s.String()
}

func (s *Toolset) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *Toolset) GetDescription() *string {
	return s.Description
}

func (s *Toolset) GetGeneration() *int64 {
	return s.Generation
}

func (s *Toolset) GetKind() *string {
	return s.Kind
}

func (s *Toolset) GetLabels() map[string]*string {
	return s.Labels
}

func (s *Toolset) GetName() *string {
	return s.Name
}

func (s *Toolset) GetSpec() *ToolsetSpec {
	return s.Spec
}

func (s *Toolset) GetStatus() *ToolsetStatus {
	return s.Status
}

func (s *Toolset) GetUid() *string {
	return s.Uid
}

func (s *Toolset) SetCreatedTime(v string) *Toolset {
	s.CreatedTime = &v
	return s
}

func (s *Toolset) SetDescription(v string) *Toolset {
	s.Description = &v
	return s
}

func (s *Toolset) SetGeneration(v int64) *Toolset {
	s.Generation = &v
	return s
}

func (s *Toolset) SetKind(v string) *Toolset {
	s.Kind = &v
	return s
}

func (s *Toolset) SetLabels(v map[string]*string) *Toolset {
	s.Labels = v
	return s
}

func (s *Toolset) SetName(v string) *Toolset {
	s.Name = &v
	return s
}

func (s *Toolset) SetSpec(v *ToolsetSpec) *Toolset {
	s.Spec = v
	return s
}

func (s *Toolset) SetStatus(v *ToolsetStatus) *Toolset {
	s.Status = v
	return s
}

func (s *Toolset) SetUid(v string) *Toolset {
	s.Uid = &v
	return s
}

func (s *Toolset) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelSet interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *ModelSet
	GetCreatedTime() *string
	SetDescription(v string) *ModelSet
	GetDescription() *string
	SetGeneration(v int64) *ModelSet
	GetGeneration() *int64
	SetKind(v string) *ModelSet
	GetKind() *string
	SetLabels(v map[string]*string) *ModelSet
	GetLabels() map[string]*string
	SetName(v string) *ModelSet
	GetName() *string
	SetStatus(v *ModelSetStatus) *ModelSet
	GetStatus() *ModelSetStatus
	SetUid(v string) *ModelSet
	GetUid() *string
}

type ModelSet struct {
	// example:
	//
	// 2021-11-19T09:34:38Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// test-description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Generation  *int64  `json:"generation,omitempty" xml:"generation,omitempty"`
	// example:
	//
	// ModelProvider
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// example:
	//
	// key=value
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my-modelset
	Name   *string         `json:"name,omitempty" xml:"name,omitempty"`
	Status *ModelSetStatus `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1455541096***548
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ModelSet) String() string {
	return dara.Prettify(s)
}

func (s ModelSet) GoString() string {
	return s.String()
}

func (s *ModelSet) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ModelSet) GetDescription() *string {
	return s.Description
}

func (s *ModelSet) GetGeneration() *int64 {
	return s.Generation
}

func (s *ModelSet) GetKind() *string {
	return s.Kind
}

func (s *ModelSet) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ModelSet) GetName() *string {
	return s.Name
}

func (s *ModelSet) GetStatus() *ModelSetStatus {
	return s.Status
}

func (s *ModelSet) GetUid() *string {
	return s.Uid
}

func (s *ModelSet) SetCreatedTime(v string) *ModelSet {
	s.CreatedTime = &v
	return s
}

func (s *ModelSet) SetDescription(v string) *ModelSet {
	s.Description = &v
	return s
}

func (s *ModelSet) SetGeneration(v int64) *ModelSet {
	s.Generation = &v
	return s
}

func (s *ModelSet) SetKind(v string) *ModelSet {
	s.Kind = &v
	return s
}

func (s *ModelSet) SetLabels(v map[string]*string) *ModelSet {
	s.Labels = v
	return s
}

func (s *ModelSet) SetName(v string) *ModelSet {
	s.Name = &v
	return s
}

func (s *ModelSet) SetStatus(v *ModelSetStatus) *ModelSet {
	s.Status = v
	return s
}

func (s *ModelSet) SetUid(v string) *ModelSet {
	s.Uid = &v
	return s
}

func (s *ModelSet) Validate() error {
	return dara.Validate(s)
}

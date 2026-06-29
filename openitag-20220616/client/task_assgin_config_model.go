// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTaskAssginConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAssignCount(v int64) *TaskAssginConfig
	GetAssignCount() *int64
	SetAssignField(v string) *TaskAssginConfig
	GetAssignField() *string
	SetAssignSubTaskCount(v string) *TaskAssginConfig
	GetAssignSubTaskCount() *string
	SetAssignType(v string) *TaskAssginConfig
	GetAssignType() *string
}

type TaskAssginConfig struct {
	// Allocation quantity.
	//
	// example:
	//
	// 2
	AssignCount *int64 `json:"AssignCount,omitempty" xml:"AssignCount,omitempty"`
	// Assign by field.
	//
	// example:
	//
	// label_field
	AssignField *string `json:"AssignField,omitempty" xml:"AssignField,omitempty"`
	// If average allocation is selected, specify the number of job packages.
	//
	// example:
	//
	// 0
	AssignSubTaskCount *string `json:"AssignSubTaskCount,omitempty" xml:"AssignSubTaskCount,omitempty"`
	// The allocation rule for job packages. Valid values:
	//
	// - FIXED_SIZE: Assign by fixed size.
	//
	// - AVG_SIZE: Assign by average quantity.
	//
	// - FIELD_BASE: Assign by imported field.
	//
	// example:
	//
	// FIXED_SIZE
	AssignType *string `json:"AssignType,omitempty" xml:"AssignType,omitempty"`
}

func (s TaskAssginConfig) String() string {
	return dara.Prettify(s)
}

func (s TaskAssginConfig) GoString() string {
	return s.String()
}

func (s *TaskAssginConfig) GetAssignCount() *int64 {
	return s.AssignCount
}

func (s *TaskAssginConfig) GetAssignField() *string {
	return s.AssignField
}

func (s *TaskAssginConfig) GetAssignSubTaskCount() *string {
	return s.AssignSubTaskCount
}

func (s *TaskAssginConfig) GetAssignType() *string {
	return s.AssignType
}

func (s *TaskAssginConfig) SetAssignCount(v int64) *TaskAssginConfig {
	s.AssignCount = &v
	return s
}

func (s *TaskAssginConfig) SetAssignField(v string) *TaskAssginConfig {
	s.AssignField = &v
	return s
}

func (s *TaskAssginConfig) SetAssignSubTaskCount(v string) *TaskAssginConfig {
	s.AssignSubTaskCount = &v
	return s
}

func (s *TaskAssginConfig) SetAssignType(v string) *TaskAssginConfig {
	s.AssignType = &v
	return s
}

func (s *TaskAssginConfig) Validate() error {
	return dara.Validate(s)
}

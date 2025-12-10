// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *ListJobRequest
	GetAll() *bool
	SetCount(v int32) *ListJobRequest
	GetCount() *int32
	SetMarker(v string) *ListJobRequest
	GetMarker() *string
	SetParentName(v string) *ListJobRequest
	GetParentName() *string
}

type ListJobRequest struct {
	// Specifies whether to return subtasks.\\
	//
	// Valid values: true and false.
	//
	// example:
	//
	// true
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// Specifies the number of migration tasks to be returned.\\
	//
	// Valid values: 0 - 1000 (excluding 0).\\
	//
	// Default value: 1000.
	//
	// example:
	//
	// 1000
	Count *int32 `json:"count,omitempty" xml:"count,omitempty"`
	// The marker after which the migration tasks are listed.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// test_marker
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The name of the parent task. If this parameter is specified, all subtasks of the parent task are returned.
	//
	// example:
	//
	// test_parent_job_name
	ParentName *string `json:"parentName,omitempty" xml:"parentName,omitempty"`
}

func (s ListJobRequest) String() string {
	return dara.Prettify(s)
}

func (s ListJobRequest) GoString() string {
	return s.String()
}

func (s *ListJobRequest) GetAll() *bool {
	return s.All
}

func (s *ListJobRequest) GetCount() *int32 {
	return s.Count
}

func (s *ListJobRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListJobRequest) GetParentName() *string {
	return s.ParentName
}

func (s *ListJobRequest) SetAll(v bool) *ListJobRequest {
	s.All = &v
	return s
}

func (s *ListJobRequest) SetCount(v int32) *ListJobRequest {
	s.Count = &v
	return s
}

func (s *ListJobRequest) SetMarker(v string) *ListJobRequest {
	s.Marker = &v
	return s
}

func (s *ListJobRequest) SetParentName(v string) *ListJobRequest {
	s.ParentName = &v
	return s
}

func (s *ListJobRequest) Validate() error {
	return dara.Validate(s)
}

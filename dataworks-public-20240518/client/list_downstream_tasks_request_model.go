// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDownstreamTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ListDownstreamTasksRequest
	GetId() *int64
	SetPageNumber(v int32) *ListDownstreamTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDownstreamTasksRequest
	GetPageSize() *int32
	SetProjectEnv(v string) *ListDownstreamTasksRequest
	GetProjectEnv() *string
}

type ListDownstreamTasksRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The environment of the workspace.
	//
	// Valid values:
	//
	// 	- Prod: production environment
	//
	// 	- Dev: development environment
	//
	// example:
	//
	// Prod
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s ListDownstreamTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDownstreamTasksRequest) GoString() string {
	return s.String()
}

func (s *ListDownstreamTasksRequest) GetId() *int64 {
	return s.Id
}

func (s *ListDownstreamTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDownstreamTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDownstreamTasksRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListDownstreamTasksRequest) SetId(v int64) *ListDownstreamTasksRequest {
	s.Id = &v
	return s
}

func (s *ListDownstreamTasksRequest) SetPageNumber(v int32) *ListDownstreamTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDownstreamTasksRequest) SetPageSize(v int32) *ListDownstreamTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListDownstreamTasksRequest) SetProjectEnv(v string) *ListDownstreamTasksRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListDownstreamTasksRequest) Validate() error {
	return dara.Validate(s)
}

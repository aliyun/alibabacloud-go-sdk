// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpstreamTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ListUpstreamTasksRequest
	GetId() *int64
	SetPageNumber(v int32) *ListUpstreamTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListUpstreamTasksRequest
	GetPageSize() *int32
	SetProjectEnv(v string) *ListUpstreamTasksRequest
	GetProjectEnv() *string
}

type ListUpstreamTasksRequest struct {
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
	// The environment of the workspace. Valid values:
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

func (s ListUpstreamTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUpstreamTasksRequest) GoString() string {
	return s.String()
}

func (s *ListUpstreamTasksRequest) GetId() *int64 {
	return s.Id
}

func (s *ListUpstreamTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListUpstreamTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUpstreamTasksRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListUpstreamTasksRequest) SetId(v int64) *ListUpstreamTasksRequest {
	s.Id = &v
	return s
}

func (s *ListUpstreamTasksRequest) SetPageNumber(v int32) *ListUpstreamTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUpstreamTasksRequest) SetPageSize(v int32) *ListUpstreamTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListUpstreamTasksRequest) SetProjectEnv(v string) *ListUpstreamTasksRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListUpstreamTasksRequest) Validate() error {
	return dara.Validate(s)
}

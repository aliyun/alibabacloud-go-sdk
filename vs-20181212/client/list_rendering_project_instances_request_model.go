// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingProjectInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListRenderingProjectInstancesRequest
	GetEndTime() *string
	SetPageNumber(v int32) *ListRenderingProjectInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRenderingProjectInstancesRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListRenderingProjectInstancesRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *ListRenderingProjectInstancesRequest
	GetRenderingInstanceId() *string
	SetStartTime(v string) *ListRenderingProjectInstancesRequest
	GetStartTime() *string
	SetState(v string) *ListRenderingProjectInstancesRequest
	GetState() *string
}

type ListRenderingProjectInstancesRequest struct {
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Idle
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListRenderingProjectInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingProjectInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListRenderingProjectInstancesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRenderingProjectInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRenderingProjectInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRenderingProjectInstancesRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListRenderingProjectInstancesRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListRenderingProjectInstancesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRenderingProjectInstancesRequest) GetState() *string {
	return s.State
}

func (s *ListRenderingProjectInstancesRequest) SetEndTime(v string) *ListRenderingProjectInstancesRequest {
	s.EndTime = &v
	return s
}

func (s *ListRenderingProjectInstancesRequest) SetPageNumber(v int32) *ListRenderingProjectInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRenderingProjectInstancesRequest) SetPageSize(v int32) *ListRenderingProjectInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRenderingProjectInstancesRequest) SetProjectId(v string) *ListRenderingProjectInstancesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListRenderingProjectInstancesRequest) SetRenderingInstanceId(v string) *ListRenderingProjectInstancesRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListRenderingProjectInstancesRequest) SetStartTime(v string) *ListRenderingProjectInstancesRequest {
	s.StartTime = &v
	return s
}

func (s *ListRenderingProjectInstancesRequest) SetState(v string) *ListRenderingProjectInstancesRequest {
	s.State = &v
	return s
}

func (s *ListRenderingProjectInstancesRequest) Validate() error {
	return dara.Validate(s)
}

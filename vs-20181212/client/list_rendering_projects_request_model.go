// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingProjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListRenderingProjectsRequest
	GetEndTime() *string
	SetPageNumber(v int32) *ListRenderingProjectsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRenderingProjectsRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListRenderingProjectsRequest
	GetProjectId() *string
	SetProjectName(v string) *ListRenderingProjectsRequest
	GetProjectName() *string
	SetStartTime(v string) *ListRenderingProjectsRequest
	GetStartTime() *string
}

type ListRenderingProjectsRequest struct {
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// idata_content
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListRenderingProjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingProjectsRequest) GoString() string {
	return s.String()
}

func (s *ListRenderingProjectsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRenderingProjectsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRenderingProjectsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRenderingProjectsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListRenderingProjectsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListRenderingProjectsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRenderingProjectsRequest) SetEndTime(v string) *ListRenderingProjectsRequest {
	s.EndTime = &v
	return s
}

func (s *ListRenderingProjectsRequest) SetPageNumber(v int32) *ListRenderingProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRenderingProjectsRequest) SetPageSize(v int32) *ListRenderingProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRenderingProjectsRequest) SetProjectId(v string) *ListRenderingProjectsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListRenderingProjectsRequest) SetProjectName(v string) *ListRenderingProjectsRequest {
	s.ProjectName = &v
	return s
}

func (s *ListRenderingProjectsRequest) SetStartTime(v string) *ListRenderingProjectsRequest {
	s.StartTime = &v
	return s
}

func (s *ListRenderingProjectsRequest) Validate() error {
	return dara.Validate(s)
}

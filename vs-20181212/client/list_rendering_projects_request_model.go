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
	// A parameter that specifies the time range to filter by. The time must be in UTC and follow the ISO 8601 format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2024-10-03T02:18:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Page number. The first page is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Project ID
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Project name
	//
	// example:
	//
	// idata_content
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Start time of the time range filter. Use ISO 8601 format in UTC, such as yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2024-09-29T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

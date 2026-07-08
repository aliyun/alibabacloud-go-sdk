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
	// End time of the time range. Use ISO 8601 format in UTC. Format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2024-11-30T02:18:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Page number. Start from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Project ID
	//
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Cloud application service instance ID
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// A parameter for filtering by a time range. The time must be in UTC and formatted according to the ISO 8601 standard as \\`yyyy-MM-ddTHH:mm:ssZ\\`.
	//
	// example:
	//
	// 2024-11-27T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Instance status. Valid values:
	//
	// 1. Idle
	//
	// 2. Locked
	//
	// 3. InUse
	//
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

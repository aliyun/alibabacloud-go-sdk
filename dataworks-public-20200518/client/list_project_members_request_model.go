// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectMembersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListProjectMembersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListProjectMembersRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListProjectMembersRequest
	GetProjectId() *int64
}

type ListProjectMembersRequest struct {
	// The page number. Valid values: 1 to 30. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListProjectMembersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProjectMembersRequest) GoString() string {
	return s.String()
}

func (s *ListProjectMembersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProjectMembersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProjectMembersRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListProjectMembersRequest) SetPageNumber(v int32) *ListProjectMembersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListProjectMembersRequest) SetPageSize(v int32) *ListProjectMembersRequest {
	s.PageSize = &v
	return s
}

func (s *ListProjectMembersRequest) SetProjectId(v int64) *ListProjectMembersRequest {
	s.ProjectId = &v
	return s
}

func (s *ListProjectMembersRequest) Validate() error {
	return dara.Validate(s)
}

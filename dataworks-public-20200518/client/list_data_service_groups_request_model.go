// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupNameKeyword(v string) *ListDataServiceGroupsRequest
	GetGroupNameKeyword() *string
	SetPageNumber(v int32) *ListDataServiceGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataServiceGroupsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataServiceGroupsRequest
	GetProjectId() *int64
	SetTenantId(v int64) *ListDataServiceGroupsRequest
	GetTenantId() *int64
}

type ListDataServiceGroupsRequest struct {
	// The keyword of the business process name. Fuzzy match is supported.
	//
	// example:
	//
	// TestBusinessProcess
	GroupNameKeyword *string `json:"GroupNameKeyword,omitempty" xml:"GroupNameKeyword,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace ID.
	//
	// You can obtain this value from PageResult.ProjectList[].ProjectId in the response of the ListProjects operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// **[Deprecated]*	- The tenant ID.
	//
	// example:
	//
	// 10001
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListDataServiceGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceGroupsRequest) GetGroupNameKeyword() *string {
	return s.GroupNameKeyword
}

func (s *ListDataServiceGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataServiceGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceGroupsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataServiceGroupsRequest) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListDataServiceGroupsRequest) SetGroupNameKeyword(v string) *ListDataServiceGroupsRequest {
	s.GroupNameKeyword = &v
	return s
}

func (s *ListDataServiceGroupsRequest) SetPageNumber(v int32) *ListDataServiceGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataServiceGroupsRequest) SetPageSize(v int32) *ListDataServiceGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceGroupsRequest) SetProjectId(v int64) *ListDataServiceGroupsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataServiceGroupsRequest) SetTenantId(v int64) *ListDataServiceGroupsRequest {
	s.TenantId = &v
	return s
}

func (s *ListDataServiceGroupsRequest) Validate() error {
	return dara.Validate(s)
}

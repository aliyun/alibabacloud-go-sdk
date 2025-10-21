// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWordGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *ListWordGroupRequest
	GetGroupName() *string
	SetOrder(v string) *ListWordGroupRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListWordGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListWordGroupRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListWordGroupRequest
	GetRegionId() *string
	SetSortBy(v string) *ListWordGroupRequest
	GetSortBy() *string
	SetWorkspaceId(v int64) *ListWordGroupRequest
	GetWorkspaceId() *int64
}

type ListWordGroupRequest struct {
	// Keyword group name.
	//
	// example:
	//
	// testGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// When performing a paginated query, sort the specified field in ascending or descending order. Values are as follows:
	//
	// 	- asc: Ascending.
	//
	// 	- desc: Descending.
	//
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, the maximum number of results returned per page.
	//
	// Maximum limit: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Sort field.
	//
	// example:
	//
	// GmtModified
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// Workspace ID.
	//
	// example:
	//
	// 620***
	WorkspaceId *int64 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListWordGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWordGroupRequest) GoString() string {
	return s.String()
}

func (s *ListWordGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListWordGroupRequest) GetOrder() *string {
	return s.Order
}

func (s *ListWordGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListWordGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListWordGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListWordGroupRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListWordGroupRequest) GetWorkspaceId() *int64 {
	return s.WorkspaceId
}

func (s *ListWordGroupRequest) SetGroupName(v string) *ListWordGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ListWordGroupRequest) SetOrder(v string) *ListWordGroupRequest {
	s.Order = &v
	return s
}

func (s *ListWordGroupRequest) SetPageNumber(v int32) *ListWordGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWordGroupRequest) SetPageSize(v int32) *ListWordGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListWordGroupRequest) SetRegionId(v string) *ListWordGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListWordGroupRequest) SetSortBy(v string) *ListWordGroupRequest {
	s.SortBy = &v
	return s
}

func (s *ListWordGroupRequest) SetWorkspaceId(v int64) *ListWordGroupRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListWordGroupRequest) Validate() error {
	return dara.Validate(s)
}

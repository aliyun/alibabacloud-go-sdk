// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStorageProjectsInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAscOrder(v bool) *ListStorageProjectsInfoRequest
	GetAscOrder() *bool
	SetDate(v string) *ListStorageProjectsInfoRequest
	GetDate() *string
	SetOrderColumn(v string) *ListStorageProjectsInfoRequest
	GetOrderColumn() *string
	SetPageNumber(v int64) *ListStorageProjectsInfoRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListStorageProjectsInfoRequest
	GetPageSize() *int64
	SetProjectPrefix(v string) *ListStorageProjectsInfoRequest
	GetProjectPrefix() *string
	SetRecentDays(v int32) *ListStorageProjectsInfoRequest
	GetRecentDays() *int32
	SetRegion(v string) *ListStorageProjectsInfoRequest
	GetRegion() *string
	SetTenantId(v string) *ListStorageProjectsInfoRequest
	GetTenantId() *string
}

type ListStorageProjectsInfoRequest struct {
	// example:
	//
	// false
	AscOrder *bool `json:"ascOrder,omitempty" xml:"ascOrder,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20241205
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// example:
	//
	// totalStorage
	OrderColumn *string `json:"orderColumn,omitempty" xml:"orderColumn,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// odps_project
	ProjectPrefix *string `json:"projectPrefix,omitempty" xml:"projectPrefix,omitempty"`
	// example:
	//
	// 1
	RecentDays *int32 `json:"recentDays,omitempty" xml:"recentDays,omitempty"`
	// example:
	//
	// cn-chengdu
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 28074710977****
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListStorageProjectsInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStorageProjectsInfoRequest) GoString() string {
	return s.String()
}

func (s *ListStorageProjectsInfoRequest) GetAscOrder() *bool {
	return s.AscOrder
}

func (s *ListStorageProjectsInfoRequest) GetDate() *string {
	return s.Date
}

func (s *ListStorageProjectsInfoRequest) GetOrderColumn() *string {
	return s.OrderColumn
}

func (s *ListStorageProjectsInfoRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStorageProjectsInfoRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStorageProjectsInfoRequest) GetProjectPrefix() *string {
	return s.ProjectPrefix
}

func (s *ListStorageProjectsInfoRequest) GetRecentDays() *int32 {
	return s.RecentDays
}

func (s *ListStorageProjectsInfoRequest) GetRegion() *string {
	return s.Region
}

func (s *ListStorageProjectsInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListStorageProjectsInfoRequest) SetAscOrder(v bool) *ListStorageProjectsInfoRequest {
	s.AscOrder = &v
	return s
}

func (s *ListStorageProjectsInfoRequest) SetDate(v string) *ListStorageProjectsInfoRequest {
	s.Date = &v
	return s
}

func (s *ListStorageProjectsInfoRequest) SetOrderColumn(v string) *ListStorageProjectsInfoRequest {
	s.OrderColumn = &v
	return s
}

func (s *ListStorageProjectsInfoRequest) SetPageNumber(v int64) *ListStorageProjectsInfoRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStorageProjectsInfoRequest) SetPageSize(v int64) *ListStorageProjectsInfoRequest {
	s.PageSize = &v
	return s
}

func (s *ListStorageProjectsInfoRequest) SetProjectPrefix(v string) *ListStorageProjectsInfoRequest {
	s.ProjectPrefix = &v
	return s
}

func (s *ListStorageProjectsInfoRequest) SetRecentDays(v int32) *ListStorageProjectsInfoRequest {
	s.RecentDays = &v
	return s
}

func (s *ListStorageProjectsInfoRequest) SetRegion(v string) *ListStorageProjectsInfoRequest {
	s.Region = &v
	return s
}

func (s *ListStorageProjectsInfoRequest) SetTenantId(v string) *ListStorageProjectsInfoRequest {
	s.TenantId = &v
	return s
}

func (s *ListStorageProjectsInfoRequest) Validate() error {
	return dara.Validate(s)
}

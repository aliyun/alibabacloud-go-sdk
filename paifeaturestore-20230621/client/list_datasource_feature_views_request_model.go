// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasourceFeatureViewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *ListDatasourceFeatureViewsRequest
	GetAll() *bool
	SetEndDate(v string) *ListDatasourceFeatureViewsRequest
	GetEndDate() *string
	SetName(v string) *ListDatasourceFeatureViewsRequest
	GetName() *string
	SetOrder(v string) *ListDatasourceFeatureViewsRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListDatasourceFeatureViewsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDatasourceFeatureViewsRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListDatasourceFeatureViewsRequest
	GetProjectId() *string
	SetProjectName(v string) *ListDatasourceFeatureViewsRequest
	GetProjectName() *string
	SetShowStorageUsage(v bool) *ListDatasourceFeatureViewsRequest
	GetShowStorageUsage() *bool
	SetSortBy(v string) *ListDatasourceFeatureViewsRequest
	GetSortBy() *string
	SetStartDate(v string) *ListDatasourceFeatureViewsRequest
	GetStartDate() *string
	SetType(v string) *ListDatasourceFeatureViewsRequest
	GetType() *string
	SetVerbose(v bool) *ListDatasourceFeatureViewsRequest
	GetVerbose() *bool
}

type ListDatasourceFeatureViewsRequest struct {
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// example:
	//
	// 2025-03-19
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// fv
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
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
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// fs_project
	ProjectName      *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ShowStorageUsage *bool   `json:"ShowStorageUsage,omitempty" xml:"ShowStorageUsage,omitempty"`
	// example:
	//
	// GmtCreateTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// 2025-03-14
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// example:
	//
	// Stream
	Type    *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Verbose *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s ListDatasourceFeatureViewsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourceFeatureViewsRequest) GoString() string {
	return s.String()
}

func (s *ListDatasourceFeatureViewsRequest) GetAll() *bool {
	return s.All
}

func (s *ListDatasourceFeatureViewsRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListDatasourceFeatureViewsRequest) GetName() *string {
	return s.Name
}

func (s *ListDatasourceFeatureViewsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListDatasourceFeatureViewsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDatasourceFeatureViewsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDatasourceFeatureViewsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListDatasourceFeatureViewsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDatasourceFeatureViewsRequest) GetShowStorageUsage() *bool {
	return s.ShowStorageUsage
}

func (s *ListDatasourceFeatureViewsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDatasourceFeatureViewsRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *ListDatasourceFeatureViewsRequest) GetType() *string {
	return s.Type
}

func (s *ListDatasourceFeatureViewsRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListDatasourceFeatureViewsRequest) SetAll(v bool) *ListDatasourceFeatureViewsRequest {
	s.All = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetEndDate(v string) *ListDatasourceFeatureViewsRequest {
	s.EndDate = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetName(v string) *ListDatasourceFeatureViewsRequest {
	s.Name = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetOrder(v string) *ListDatasourceFeatureViewsRequest {
	s.Order = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetPageNumber(v int32) *ListDatasourceFeatureViewsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetPageSize(v int32) *ListDatasourceFeatureViewsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetProjectId(v string) *ListDatasourceFeatureViewsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetProjectName(v string) *ListDatasourceFeatureViewsRequest {
	s.ProjectName = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetShowStorageUsage(v bool) *ListDatasourceFeatureViewsRequest {
	s.ShowStorageUsage = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetSortBy(v string) *ListDatasourceFeatureViewsRequest {
	s.SortBy = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetStartDate(v string) *ListDatasourceFeatureViewsRequest {
	s.StartDate = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetType(v string) *ListDatasourceFeatureViewsRequest {
	s.Type = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) SetVerbose(v bool) *ListDatasourceFeatureViewsRequest {
	s.Verbose = &v
	return s
}

func (s *ListDatasourceFeatureViewsRequest) Validate() error {
	return dara.Validate(s)
}

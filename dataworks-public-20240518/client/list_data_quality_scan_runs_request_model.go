// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityScanRunsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeFrom(v int64) *ListDataQualityScanRunsRequest
	GetCreateTimeFrom() *int64
	SetCreateTimeTo(v int64) *ListDataQualityScanRunsRequest
	GetCreateTimeTo() *int64
	SetDataQualityScanId(v int64) *ListDataQualityScanRunsRequest
	GetDataQualityScanId() *int64
	SetFilter(v map[string]interface{}) *ListDataQualityScanRunsRequest
	GetFilter() map[string]interface{}
	SetPageNumber(v int32) *ListDataQualityScanRunsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataQualityScanRunsRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataQualityScanRunsRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListDataQualityScanRunsRequest
	GetSortBy() *string
	SetStatus(v string) *ListDataQualityScanRunsRequest
	GetStatus() *string
}

type ListDataQualityScanRunsRequest struct {
	// The earliest time when the data quality monitor starts to run.
	//
	// example:
	//
	// 1710239005403
	CreateTimeFrom *int64 `json:"CreateTimeFrom,omitempty" xml:"CreateTimeFrom,omitempty"`
	// The latest time when the data quality monitor starts to run.
	//
	// example:
	//
	// 1710239005403
	CreateTimeTo *int64 `json:"CreateTimeTo,omitempty" xml:"CreateTimeTo,omitempty"`
	// The ID of the data quality monitor.
	//
	// example:
	//
	// 10001
	DataQualityScanId *int64                 `json:"DataQualityScanId,omitempty" xml:"DataQualityScanId,omitempty"`
	Filter            map[string]interface{} `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The page number of the results. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The list of sorting fields. Supports fields such as last modified time and creation time. Format: "SortField+SortOrder (Desc/Asc)", where Asc is the default. Valid values:
	//
	// 	- CreateTime (Desc/Asc)
	//
	// 	- Id (Desc/Asc)
	//
	// example:
	//
	// CreateTime Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The status of the data quality check result.
	//
	// 	- Pass
	//
	// 	- Running
	//
	// 	- Error
	//
	// 	- Fail
	//
	// 	- Warn
	//
	// example:
	//
	// Fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDataQualityScanRunsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScanRunsRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityScanRunsRequest) GetCreateTimeFrom() *int64 {
	return s.CreateTimeFrom
}

func (s *ListDataQualityScanRunsRequest) GetCreateTimeTo() *int64 {
	return s.CreateTimeTo
}

func (s *ListDataQualityScanRunsRequest) GetDataQualityScanId() *int64 {
	return s.DataQualityScanId
}

func (s *ListDataQualityScanRunsRequest) GetFilter() map[string]interface{} {
	return s.Filter
}

func (s *ListDataQualityScanRunsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityScanRunsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityScanRunsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityScanRunsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDataQualityScanRunsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDataQualityScanRunsRequest) SetCreateTimeFrom(v int64) *ListDataQualityScanRunsRequest {
	s.CreateTimeFrom = &v
	return s
}

func (s *ListDataQualityScanRunsRequest) SetCreateTimeTo(v int64) *ListDataQualityScanRunsRequest {
	s.CreateTimeTo = &v
	return s
}

func (s *ListDataQualityScanRunsRequest) SetDataQualityScanId(v int64) *ListDataQualityScanRunsRequest {
	s.DataQualityScanId = &v
	return s
}

func (s *ListDataQualityScanRunsRequest) SetFilter(v map[string]interface{}) *ListDataQualityScanRunsRequest {
	s.Filter = v
	return s
}

func (s *ListDataQualityScanRunsRequest) SetPageNumber(v int32) *ListDataQualityScanRunsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityScanRunsRequest) SetPageSize(v int32) *ListDataQualityScanRunsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityScanRunsRequest) SetProjectId(v int64) *ListDataQualityScanRunsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityScanRunsRequest) SetSortBy(v string) *ListDataQualityScanRunsRequest {
	s.SortBy = &v
	return s
}

func (s *ListDataQualityScanRunsRequest) SetStatus(v string) *ListDataQualityScanRunsRequest {
	s.Status = &v
	return s
}

func (s *ListDataQualityScanRunsRequest) Validate() error {
	return dara.Validate(s)
}

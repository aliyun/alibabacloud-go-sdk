// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityScanRunsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeFrom(v int64) *ListDataQualityScanRunsShrinkRequest
	GetCreateTimeFrom() *int64
	SetCreateTimeTo(v int64) *ListDataQualityScanRunsShrinkRequest
	GetCreateTimeTo() *int64
	SetDataQualityScanId(v int64) *ListDataQualityScanRunsShrinkRequest
	GetDataQualityScanId() *int64
	SetFilterShrink(v string) *ListDataQualityScanRunsShrinkRequest
	GetFilterShrink() *string
	SetPageNumber(v int32) *ListDataQualityScanRunsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataQualityScanRunsShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataQualityScanRunsShrinkRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListDataQualityScanRunsShrinkRequest
	GetSortBy() *string
	SetStatus(v string) *ListDataQualityScanRunsShrinkRequest
	GetStatus() *string
}

type ListDataQualityScanRunsShrinkRequest struct {
	// The earliest start time of a data quality scan run to include in the results. Specify the time as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1710239005403
	CreateTimeFrom *int64 `json:"CreateTimeFrom,omitempty" xml:"CreateTimeFrom,omitempty"`
	// The latest start time of a data quality scan run to include in the results. Specify the time as a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1710239005403
	CreateTimeTo *int64 `json:"CreateTimeTo,omitempty" xml:"CreateTimeTo,omitempty"`
	// The ID of the data quality scan.
	//
	// example:
	//
	// 10001
	DataQualityScanId *int64 `json:"DataQualityScanId,omitempty" xml:"DataQualityScanId,omitempty"`
	// An object with advanced filter conditions. The following parameters are supported:
	//
	// - `TaskInstanceId`: The ID of the task instance.
	//
	// - `RunNumber`: The run number of the instance.
	//
	// example:
	//
	// {
	//
	//     "TaskInstanceId": "111",
	//
	//     "RunNumber": "1"
	//
	// }
	FilterShrink *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The page number to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
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
	// The sort field and order for the results. The format is `FieldName Order`. The default order is ascending (Asc). Supported fields:
	//
	// - CreateTime (Desc/Asc)
	//
	// - Id (Desc/Asc)
	//
	// example:
	//
	// CreateTime Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The status of the data quality scan run. Valid values:
	//
	// - Pass
	//
	// - Running
	//
	// - Error
	//
	// - Fail
	//
	// - Warn
	//
	// example:
	//
	// Fail
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDataQualityScanRunsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityScanRunsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityScanRunsShrinkRequest) GetCreateTimeFrom() *int64 {
	return s.CreateTimeFrom
}

func (s *ListDataQualityScanRunsShrinkRequest) GetCreateTimeTo() *int64 {
	return s.CreateTimeTo
}

func (s *ListDataQualityScanRunsShrinkRequest) GetDataQualityScanId() *int64 {
	return s.DataQualityScanId
}

func (s *ListDataQualityScanRunsShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *ListDataQualityScanRunsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityScanRunsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityScanRunsShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityScanRunsShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDataQualityScanRunsShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListDataQualityScanRunsShrinkRequest) SetCreateTimeFrom(v int64) *ListDataQualityScanRunsShrinkRequest {
	s.CreateTimeFrom = &v
	return s
}

func (s *ListDataQualityScanRunsShrinkRequest) SetCreateTimeTo(v int64) *ListDataQualityScanRunsShrinkRequest {
	s.CreateTimeTo = &v
	return s
}

func (s *ListDataQualityScanRunsShrinkRequest) SetDataQualityScanId(v int64) *ListDataQualityScanRunsShrinkRequest {
	s.DataQualityScanId = &v
	return s
}

func (s *ListDataQualityScanRunsShrinkRequest) SetFilterShrink(v string) *ListDataQualityScanRunsShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *ListDataQualityScanRunsShrinkRequest) SetPageNumber(v int32) *ListDataQualityScanRunsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityScanRunsShrinkRequest) SetPageSize(v int32) *ListDataQualityScanRunsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityScanRunsShrinkRequest) SetProjectId(v int64) *ListDataQualityScanRunsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityScanRunsShrinkRequest) SetSortBy(v string) *ListDataQualityScanRunsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListDataQualityScanRunsShrinkRequest) SetStatus(v string) *ListDataQualityScanRunsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListDataQualityScanRunsShrinkRequest) Validate() error {
	return dara.Validate(s)
}

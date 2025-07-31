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
	// example:
	//
	// 1710239005403
	CreateTimeFrom *int64 `json:"CreateTimeFrom,omitempty" xml:"CreateTimeFrom,omitempty"`
	// example:
	//
	// 1710239005403
	CreateTimeTo *int64 `json:"CreateTimeTo,omitempty" xml:"CreateTimeTo,omitempty"`
	// example:
	//
	// 10001
	DataQualityScanId *int64 `json:"DataQualityScanId,omitempty" xml:"DataQualityScanId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// CreateTime Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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

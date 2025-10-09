// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityScanId(v int64) *ListDataQualityAlertRulesRequest
	GetDataQualityScanId() *int64
	SetPageNumber(v int32) *ListDataQualityAlertRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDataQualityAlertRulesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListDataQualityAlertRulesRequest
	GetProjectId() *int64
	SetSortBy(v string) *ListDataQualityAlertRulesRequest
	GetSortBy() *string
}

type ListDataQualityAlertRulesRequest struct {
	// The ID of the data quality monitor that the alert rule targets.
	//
	// example:
	//
	// 10001
	DataQualityScanId *int64 `json:"DataQualityScanId,omitempty" xml:"DataQualityScanId,omitempty"`
	// The page number of the results.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10001
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
}

func (s ListDataQualityAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ListDataQualityAlertRulesRequest) GetDataQualityScanId() *int64 {
	return s.DataQualityScanId
}

func (s *ListDataQualityAlertRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityAlertRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityAlertRulesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityAlertRulesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListDataQualityAlertRulesRequest) SetDataQualityScanId(v int64) *ListDataQualityAlertRulesRequest {
	s.DataQualityScanId = &v
	return s
}

func (s *ListDataQualityAlertRulesRequest) SetPageNumber(v int32) *ListDataQualityAlertRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityAlertRulesRequest) SetPageSize(v int32) *ListDataQualityAlertRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityAlertRulesRequest) SetProjectId(v int64) *ListDataQualityAlertRulesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityAlertRulesRequest) SetSortBy(v string) *ListDataQualityAlertRulesRequest {
	s.SortBy = &v
	return s
}

func (s *ListDataQualityAlertRulesRequest) Validate() error {
	return dara.Validate(s)
}

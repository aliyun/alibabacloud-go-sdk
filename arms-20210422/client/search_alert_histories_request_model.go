// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertHistoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v int64) *SearchAlertHistoriesRequest
	GetAlertId() *int64
	SetAlertType(v int32) *SearchAlertHistoriesRequest
	GetAlertType() *int32
	SetCurrentPage(v int32) *SearchAlertHistoriesRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *SearchAlertHistoriesRequest
	GetEndTime() *int64
	SetPageSize(v int32) *SearchAlertHistoriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *SearchAlertHistoriesRequest
	GetRegionId() *string
	SetStartTime(v int64) *SearchAlertHistoriesRequest
	GetStartTime() *int64
}

type SearchAlertHistoriesRequest struct {
	AlertId     *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertType   *int32 `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize    *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchAlertHistoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertHistoriesRequest) GoString() string {
	return s.String()
}

func (s *SearchAlertHistoriesRequest) GetAlertId() *int64 {
	return s.AlertId
}

func (s *SearchAlertHistoriesRequest) GetAlertType() *int32 {
	return s.AlertType
}

func (s *SearchAlertHistoriesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *SearchAlertHistoriesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SearchAlertHistoriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAlertHistoriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchAlertHistoriesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SearchAlertHistoriesRequest) SetAlertId(v int64) *SearchAlertHistoriesRequest {
	s.AlertId = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetAlertType(v int32) *SearchAlertHistoriesRequest {
	s.AlertType = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetCurrentPage(v int32) *SearchAlertHistoriesRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetEndTime(v int64) *SearchAlertHistoriesRequest {
	s.EndTime = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetPageSize(v int32) *SearchAlertHistoriesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetRegionId(v string) *SearchAlertHistoriesRequest {
	s.RegionId = &v
	return s
}

func (s *SearchAlertHistoriesRequest) SetStartTime(v int64) *SearchAlertHistoriesRequest {
	s.StartTime = &v
	return s
}

func (s *SearchAlertHistoriesRequest) Validate() error {
	return dara.Validate(s)
}

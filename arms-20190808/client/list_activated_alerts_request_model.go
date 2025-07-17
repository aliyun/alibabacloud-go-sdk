// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActivatedAlertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListActivatedAlertsRequest
	GetCurrentPage() *int32
	SetFilter(v string) *ListActivatedAlertsRequest
	GetFilter() *string
	SetPageSize(v int32) *ListActivatedAlertsRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListActivatedAlertsRequest
	GetRegionId() *string
}

type ListActivatedAlertsRequest struct {
	// The number of the page to return. Default value: `1`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The filter condition in the `{"key":"value"}`format. You must specify the `key` and `value` of the filter condition.
	//
	// example:
	//
	// {"alertname":"Container CPU usage is greater than 80%"}
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// The number of entries to return on each page. Default value: `10`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListActivatedAlertsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListActivatedAlertsRequest) GoString() string {
	return s.String()
}

func (s *ListActivatedAlertsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListActivatedAlertsRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListActivatedAlertsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListActivatedAlertsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListActivatedAlertsRequest) SetCurrentPage(v int32) *ListActivatedAlertsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListActivatedAlertsRequest) SetFilter(v string) *ListActivatedAlertsRequest {
	s.Filter = &v
	return s
}

func (s *ListActivatedAlertsRequest) SetPageSize(v int32) *ListActivatedAlertsRequest {
	s.PageSize = &v
	return s
}

func (s *ListActivatedAlertsRequest) SetRegionId(v string) *ListActivatedAlertsRequest {
	s.RegionId = &v
	return s
}

func (s *ListActivatedAlertsRequest) Validate() error {
	return dara.Validate(s)
}

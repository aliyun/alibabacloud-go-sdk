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
	// The ID of the alert rule. You can call the SearchAlertRules operation and view the `Id` parameter in the response. For more information, see [SearchAlertRules](https://help.aliyun.com/document_detail/175825.html).
	//
	// example:
	//
	// 123
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The type of the alert rule. Valid values:
	//
	// 	- `1`: a custom alert rule that is used to monitor drill-down data sets
	//
	// 	- `3`: a custom alert rule that is used to monitor tiled data sets
	//
	// 	- `4`: an alert rule that is used to monitor web pages, including the default alert rule for browser monitoring
	//
	// 	- `5`: an alert rule that is used to monitor applications, including the default alert rule for application monitoring
	//
	// 	- `6`: the default alert rule for browser monitoring
	//
	// 	- `7`: the default alert rule for application monitoring
	//
	// 	- `8`: a Tracing Analysis alert rule
	//
	// 	- `101`: a Prometheus alert rule
	//
	// example:
	//
	// 4
	AlertType *int32 `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The number of the page to return. Default value: `1`.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. The value is a UNIX timestamp of the LONG data type. Unit: milliseconds. The default value is the current time.
	//
	// example:
	//
	// 1579499626000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return on each page. Default value: `10`.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region. Default value: `cn-hangzhou`.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. The value is a UNIX timestamp of the LONG data type. Unit: milliseconds. The default value is 10 minutes before the current time.
	//
	// example:
	//
	// 1595568910000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

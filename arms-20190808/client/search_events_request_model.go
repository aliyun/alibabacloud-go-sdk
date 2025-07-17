// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v int64) *SearchEventsRequest
	GetAlertId() *int64
	SetAlertType(v int32) *SearchEventsRequest
	GetAlertType() *int32
	SetAppType(v string) *SearchEventsRequest
	GetAppType() *string
	SetCurrentPage(v int32) *SearchEventsRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *SearchEventsRequest
	GetEndTime() *int64
	SetIsTrigger(v int32) *SearchEventsRequest
	GetIsTrigger() *int32
	SetPageSize(v int32) *SearchEventsRequest
	GetPageSize() *int32
	SetPid(v string) *SearchEventsRequest
	GetPid() *string
	SetRegionId(v string) *SearchEventsRequest
	GetRegionId() *string
	SetStartTime(v int64) *SearchEventsRequest
	GetStartTime() *int64
}

type SearchEventsRequest struct {
	// The ID of the alert rule. You can call the SearchAlertRules operation and view the `Id` parameter in the response. For more information, see [SearchAlertRules](https://help.aliyun.com/document_detail/175825.html).
	//
	// example:
	//
	// 123
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The type of the alert rule. Valid values:
	//
	// 	- `1`: custom alert rules to monitor drill-down data sets
	//
	// 	- `3`: custom alert rules to monitor tiled data sets
	//
	// 	- `4`: alert rules to monitor the frontend, including the default frontend alert rules
	//
	// 	- `5`: alert rules to monitor applications, including the default application alert rules
	//
	// 	- `6`: the default frontend alert rules
	//
	// 	- `7`: the default application alert rules
	//
	// 	- `8`: Tracing Analysis alert rules
	//
	// 	- `101`: Prometheus alert rules
	//
	// example:
	//
	// 4
	AlertType *int32 `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The type of the application that is associated with the alert rule. Valid values:
	//
	// 	- `TRACE`: application monitoring
	//
	// 	- `RETCODE`: frontend monitoring
	//
	// example:
	//
	// TRACE
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The number of the page to return. Default value: `1`.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The end of the time range to query. Specify a UNIX timestamp of the LONG data type, in milliseconds. The default value is the current time.
	//
	// example:
	//
	// 1595568970000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Specifies whether the alert event is triggered. If you do not set this parameter, all alert events are queried. Valid values:
	//
	// 	- `1`: The event is triggered.
	//
	// 	- `0`: The event is not triggered.
	//
	// example:
	//
	// 1
	IsTrigger *int32 `json:"IsTrigger,omitempty" xml:"IsTrigger,omitempty"`
	// The number of entries to return on each page. Default value: `10`.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The process identifier (PID) of the application that is associated with the alert rule.
	//
	// example:
	//
	// atc889zkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. Specify a UNIX timestamp of the LONG data type, in milliseconds. The default value is 10 minutes before the current time.
	//
	// example:
	//
	// 1595565300000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s SearchEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchEventsRequest) GoString() string {
	return s.String()
}

func (s *SearchEventsRequest) GetAlertId() *int64 {
	return s.AlertId
}

func (s *SearchEventsRequest) GetAlertType() *int32 {
	return s.AlertType
}

func (s *SearchEventsRequest) GetAppType() *string {
	return s.AppType
}

func (s *SearchEventsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *SearchEventsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *SearchEventsRequest) GetIsTrigger() *int32 {
	return s.IsTrigger
}

func (s *SearchEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchEventsRequest) GetPid() *string {
	return s.Pid
}

func (s *SearchEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchEventsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *SearchEventsRequest) SetAlertId(v int64) *SearchEventsRequest {
	s.AlertId = &v
	return s
}

func (s *SearchEventsRequest) SetAlertType(v int32) *SearchEventsRequest {
	s.AlertType = &v
	return s
}

func (s *SearchEventsRequest) SetAppType(v string) *SearchEventsRequest {
	s.AppType = &v
	return s
}

func (s *SearchEventsRequest) SetCurrentPage(v int32) *SearchEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *SearchEventsRequest) SetEndTime(v int64) *SearchEventsRequest {
	s.EndTime = &v
	return s
}

func (s *SearchEventsRequest) SetIsTrigger(v int32) *SearchEventsRequest {
	s.IsTrigger = &v
	return s
}

func (s *SearchEventsRequest) SetPageSize(v int32) *SearchEventsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchEventsRequest) SetPid(v string) *SearchEventsRequest {
	s.Pid = &v
	return s
}

func (s *SearchEventsRequest) SetRegionId(v string) *SearchEventsRequest {
	s.RegionId = &v
	return s
}

func (s *SearchEventsRequest) SetStartTime(v int64) *SearchEventsRequest {
	s.StartTime = &v
	return s
}

func (s *SearchEventsRequest) Validate() error {
	return dara.Validate(s)
}

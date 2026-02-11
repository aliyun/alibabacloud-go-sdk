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
	AlertId     *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	AlertType   *int32  `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AppType     *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsTrigger   *int32  `json:"IsTrigger,omitempty" xml:"IsTrigger,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Pid         *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// This parameter is required.
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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

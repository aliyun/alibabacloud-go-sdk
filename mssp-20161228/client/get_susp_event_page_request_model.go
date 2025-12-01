// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuspEventPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmEndTime(v int64) *GetSuspEventPageRequest
	GetAlarmEndTime() *int64
	SetAlarmStartTime(v int64) *GetSuspEventPageRequest
	GetAlarmStartTime() *int64
	SetCurrentPage(v int32) *GetSuspEventPageRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *GetSuspEventPageRequest
	GetPageSize() *int32
	SetSource(v string) *GetSuspEventPageRequest
	GetSource() *string
	SetStatus(v int32) *GetSuspEventPageRequest
	GetStatus() *int32
}

type GetSuspEventPageRequest struct {
	// Alarm end time.
	//
	// example:
	//
	// 1732515522000
	AlarmEndTime *int64 `json:"AlarmEndTime,omitempty" xml:"AlarmEndTime,omitempty"`
	// Alarm start time.
	//
	// example:
	//
	// 1722515522000
	AlarmStartTime *int64 `json:"AlarmStartTime,omitempty" xml:"AlarmStartTime,omitempty"`
	// Current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Alarm source.
	//
	// example:
	//
	// SUSP_EVENT
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Disposal status.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSuspEventPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventPageRequest) GoString() string {
	return s.String()
}

func (s *GetSuspEventPageRequest) GetAlarmEndTime() *int64 {
	return s.AlarmEndTime
}

func (s *GetSuspEventPageRequest) GetAlarmStartTime() *int64 {
	return s.AlarmStartTime
}

func (s *GetSuspEventPageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetSuspEventPageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSuspEventPageRequest) GetSource() *string {
	return s.Source
}

func (s *GetSuspEventPageRequest) GetStatus() *int32 {
	return s.Status
}

func (s *GetSuspEventPageRequest) SetAlarmEndTime(v int64) *GetSuspEventPageRequest {
	s.AlarmEndTime = &v
	return s
}

func (s *GetSuspEventPageRequest) SetAlarmStartTime(v int64) *GetSuspEventPageRequest {
	s.AlarmStartTime = &v
	return s
}

func (s *GetSuspEventPageRequest) SetCurrentPage(v int32) *GetSuspEventPageRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetSuspEventPageRequest) SetPageSize(v int32) *GetSuspEventPageRequest {
	s.PageSize = &v
	return s
}

func (s *GetSuspEventPageRequest) SetSource(v string) *GetSuspEventPageRequest {
	s.Source = &v
	return s
}

func (s *GetSuspEventPageRequest) SetStatus(v int32) *GetSuspEventPageRequest {
	s.Status = &v
	return s
}

func (s *GetSuspEventPageRequest) Validate() error {
	return dara.Validate(s)
}

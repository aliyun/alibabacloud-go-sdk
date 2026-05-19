// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientEventDashboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v string) *GetFileProtectClientEventDashboardRequest
	GetCurrentPage() *string
	SetEndTime(v int64) *GetFileProtectClientEventDashboardRequest
	GetEndTime() *int64
	SetPageSize(v string) *GetFileProtectClientEventDashboardRequest
	GetPageSize() *string
	SetStartTime(v int64) *GetFileProtectClientEventDashboardRequest
	GetStartTime() *int64
}

type GetFileProtectClientEventDashboardRequest struct {
	// example:
	//
	// 1
	CurrentPage *string `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1656038940435
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1648438617000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetFileProtectClientEventDashboardRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientEventDashboardRequest) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientEventDashboardRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *GetFileProtectClientEventDashboardRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetFileProtectClientEventDashboardRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetFileProtectClientEventDashboardRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetFileProtectClientEventDashboardRequest) SetCurrentPage(v string) *GetFileProtectClientEventDashboardRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetFileProtectClientEventDashboardRequest) SetEndTime(v int64) *GetFileProtectClientEventDashboardRequest {
	s.EndTime = &v
	return s
}

func (s *GetFileProtectClientEventDashboardRequest) SetPageSize(v string) *GetFileProtectClientEventDashboardRequest {
	s.PageSize = &v
	return s
}

func (s *GetFileProtectClientEventDashboardRequest) SetStartTime(v int64) *GetFileProtectClientEventDashboardRequest {
	s.StartTime = &v
	return s
}

func (s *GetFileProtectClientEventDashboardRequest) Validate() error {
	return dara.Validate(s)
}

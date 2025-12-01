// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageServiceCustomerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthStatus(v int32) *PageServiceCustomerRequest
	GetAuthStatus() *int32
	SetCmAuthStatus(v int32) *PageServiceCustomerRequest
	GetCmAuthStatus() *int32
	SetCurrentPage(v int32) *PageServiceCustomerRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *PageServiceCustomerRequest
	GetEndTime() *int64
	SetMonitorAuthStatus(v int32) *PageServiceCustomerRequest
	GetMonitorAuthStatus() *int32
	SetPageSize(v int32) *PageServiceCustomerRequest
	GetPageSize() *int32
	SetStartTime(v int64) *PageServiceCustomerRequest
	GetStartTime() *int64
}

type PageServiceCustomerRequest struct {
	// Authorization status.
	//
	// example:
	//
	// 1
	AuthStatus *int32 `json:"AuthStatus,omitempty" xml:"AuthStatus,omitempty"`
	// Cloud Monitoring - Alert authorization status.
	//
	// example:
	//
	// 1
	CmAuthStatus *int32 `json:"CmAuthStatus,omitempty" xml:"CmAuthStatus,omitempty"`
	// The page number of the query result, default is 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// End time. The format is a Unix timestamp, which is the number of milliseconds since January 1, 1970.
	//
	// example:
	//
	// 1710641101123
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Cloud Security - Alert authorization status.
	//
	// example:
	//
	// 1
	MonitorAuthStatus *int32 `json:"MonitorAuthStatus,omitempty" xml:"MonitorAuthStatus,omitempty"`
	// Number of records per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start time. The format is a Unix timestamp, which is the number of milliseconds since January 1, 1970.
	//
	// example:
	//
	// 1710641101000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s PageServiceCustomerRequest) String() string {
	return dara.Prettify(s)
}

func (s PageServiceCustomerRequest) GoString() string {
	return s.String()
}

func (s *PageServiceCustomerRequest) GetAuthStatus() *int32 {
	return s.AuthStatus
}

func (s *PageServiceCustomerRequest) GetCmAuthStatus() *int32 {
	return s.CmAuthStatus
}

func (s *PageServiceCustomerRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *PageServiceCustomerRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *PageServiceCustomerRequest) GetMonitorAuthStatus() *int32 {
	return s.MonitorAuthStatus
}

func (s *PageServiceCustomerRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PageServiceCustomerRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *PageServiceCustomerRequest) SetAuthStatus(v int32) *PageServiceCustomerRequest {
	s.AuthStatus = &v
	return s
}

func (s *PageServiceCustomerRequest) SetCmAuthStatus(v int32) *PageServiceCustomerRequest {
	s.CmAuthStatus = &v
	return s
}

func (s *PageServiceCustomerRequest) SetCurrentPage(v int32) *PageServiceCustomerRequest {
	s.CurrentPage = &v
	return s
}

func (s *PageServiceCustomerRequest) SetEndTime(v int64) *PageServiceCustomerRequest {
	s.EndTime = &v
	return s
}

func (s *PageServiceCustomerRequest) SetMonitorAuthStatus(v int32) *PageServiceCustomerRequest {
	s.MonitorAuthStatus = &v
	return s
}

func (s *PageServiceCustomerRequest) SetPageSize(v int32) *PageServiceCustomerRequest {
	s.PageSize = &v
	return s
}

func (s *PageServiceCustomerRequest) SetStartTime(v int64) *PageServiceCustomerRequest {
	s.StartTime = &v
	return s
}

func (s *PageServiceCustomerRequest) Validate() error {
	return dara.Validate(s)
}

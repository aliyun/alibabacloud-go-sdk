// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnterpriseAccelerateLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListEnterpriseAccelerateLogsRequest
	GetCurrentPage() *int32
	SetDepartment(v string) *ListEnterpriseAccelerateLogsRequest
	GetDepartment() *string
	SetDstAddr(v string) *ListEnterpriseAccelerateLogsRequest
	GetDstAddr() *string
	SetEndTime(v int64) *ListEnterpriseAccelerateLogsRequest
	GetEndTime() *int64
	SetPageSize(v int32) *ListEnterpriseAccelerateLogsRequest
	GetPageSize() *int32
	SetSearchMode(v string) *ListEnterpriseAccelerateLogsRequest
	GetSearchMode() *string
	SetStartTime(v int64) *ListEnterpriseAccelerateLogsRequest
	GetStartTime() *int64
	SetUsername(v string) *ListEnterpriseAccelerateLogsRequest
	GetUsername() *string
}

type ListEnterpriseAccelerateLogsRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Department  *string `json:"Department,omitempty" xml:"Department,omitempty"`
	DstAddr     *string `json:"DstAddr,omitempty" xml:"DstAddr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1748422694
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchMode *string `json:"SearchMode,omitempty" xml:"SearchMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1748419094
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Username  *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListEnterpriseAccelerateLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnterpriseAccelerateLogsRequest) GoString() string {
	return s.String()
}

func (s *ListEnterpriseAccelerateLogsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListEnterpriseAccelerateLogsRequest) GetDepartment() *string {
	return s.Department
}

func (s *ListEnterpriseAccelerateLogsRequest) GetDstAddr() *string {
	return s.DstAddr
}

func (s *ListEnterpriseAccelerateLogsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListEnterpriseAccelerateLogsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEnterpriseAccelerateLogsRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *ListEnterpriseAccelerateLogsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListEnterpriseAccelerateLogsRequest) GetUsername() *string {
	return s.Username
}

func (s *ListEnterpriseAccelerateLogsRequest) SetCurrentPage(v int32) *ListEnterpriseAccelerateLogsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsRequest) SetDepartment(v string) *ListEnterpriseAccelerateLogsRequest {
	s.Department = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsRequest) SetDstAddr(v string) *ListEnterpriseAccelerateLogsRequest {
	s.DstAddr = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsRequest) SetEndTime(v int64) *ListEnterpriseAccelerateLogsRequest {
	s.EndTime = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsRequest) SetPageSize(v int32) *ListEnterpriseAccelerateLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsRequest) SetSearchMode(v string) *ListEnterpriseAccelerateLogsRequest {
	s.SearchMode = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsRequest) SetStartTime(v int64) *ListEnterpriseAccelerateLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsRequest) SetUsername(v string) *ListEnterpriseAccelerateLogsRequest {
	s.Username = &v
	return s
}

func (s *ListEnterpriseAccelerateLogsRequest) Validate() error {
	return dara.Validate(s)
}

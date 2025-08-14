// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoanTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeLoanTaskListRequest
	GetLang() *string
	SetBatchNo(v string) *DescribeLoanTaskListRequest
	GetBatchNo() *string
	SetCurrentPage(v string) *DescribeLoanTaskListRequest
	GetCurrentPage() *string
	SetMonitorStatus(v string) *DescribeLoanTaskListRequest
	GetMonitorStatus() *string
	SetPageSize(v string) *DescribeLoanTaskListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeLoanTaskListRequest
	GetRegId() *string
}

type DescribeLoanTaskListRequest struct {
	// Set the language type for request and response messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Import batch number.
	//
	// example:
	//
	// 1
	BatchNo *string `json:"batchNo,omitempty" xml:"batchNo,omitempty"`
	// Current page number. Default is: 1.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Task status.
	//
	// example:
	//
	// WAIT
	MonitorStatus *string `json:"monitorStatus,omitempty" xml:"monitorStatus,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 20
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeLoanTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoanTaskListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoanTaskListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeLoanTaskListRequest) GetBatchNo() *string {
	return s.BatchNo
}

func (s *DescribeLoanTaskListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeLoanTaskListRequest) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *DescribeLoanTaskListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeLoanTaskListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeLoanTaskListRequest) SetLang(v string) *DescribeLoanTaskListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLoanTaskListRequest) SetBatchNo(v string) *DescribeLoanTaskListRequest {
	s.BatchNo = &v
	return s
}

func (s *DescribeLoanTaskListRequest) SetCurrentPage(v string) *DescribeLoanTaskListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeLoanTaskListRequest) SetMonitorStatus(v string) *DescribeLoanTaskListRequest {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeLoanTaskListRequest) SetPageSize(v string) *DescribeLoanTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoanTaskListRequest) SetRegId(v string) *DescribeLoanTaskListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeLoanTaskListRequest) Validate() error {
	return dara.Validate(s)
}

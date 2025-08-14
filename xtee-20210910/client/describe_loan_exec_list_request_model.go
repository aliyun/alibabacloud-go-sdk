// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoanExecListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeLoanExecListRequest
	GetLang() *string
	SetBatchNo(v string) *DescribeLoanExecListRequest
	GetBatchNo() *string
	SetCurrentPage(v string) *DescribeLoanExecListRequest
	GetCurrentPage() *string
	SetMonitorObj(v string) *DescribeLoanExecListRequest
	GetMonitorObj() *string
	SetMonitorStatus(v string) *DescribeLoanExecListRequest
	GetMonitorStatus() *string
	SetPageSize(v string) *DescribeLoanExecListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeLoanExecListRequest
	GetRegId() *string
}

type DescribeLoanExecListRequest struct {
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
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Monitoring metric data.
	//
	// example:
	//
	// de_afghcf6411
	MonitorObj *string `json:"monitorObj,omitempty" xml:"monitorObj,omitempty"`
	// Status
	//
	// example:
	//
	// WAIT
	MonitorStatus *string `json:"monitorStatus,omitempty" xml:"monitorStatus,omitempty"`
	// Page size, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeLoanExecListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoanExecListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoanExecListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeLoanExecListRequest) GetBatchNo() *string {
	return s.BatchNo
}

func (s *DescribeLoanExecListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeLoanExecListRequest) GetMonitorObj() *string {
	return s.MonitorObj
}

func (s *DescribeLoanExecListRequest) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *DescribeLoanExecListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeLoanExecListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeLoanExecListRequest) SetLang(v string) *DescribeLoanExecListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeLoanExecListRequest) SetBatchNo(v string) *DescribeLoanExecListRequest {
	s.BatchNo = &v
	return s
}

func (s *DescribeLoanExecListRequest) SetCurrentPage(v string) *DescribeLoanExecListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeLoanExecListRequest) SetMonitorObj(v string) *DescribeLoanExecListRequest {
	s.MonitorObj = &v
	return s
}

func (s *DescribeLoanExecListRequest) SetMonitorStatus(v string) *DescribeLoanExecListRequest {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeLoanExecListRequest) SetPageSize(v string) *DescribeLoanExecListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLoanExecListRequest) SetRegId(v string) *DescribeLoanExecListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeLoanExecListRequest) Validate() error {
	return dara.Validate(s)
}

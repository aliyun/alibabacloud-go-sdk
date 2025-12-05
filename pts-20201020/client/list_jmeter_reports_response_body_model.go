// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListJMeterReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListJMeterReportsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListJMeterReportsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListJMeterReportsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListJMeterReportsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListJMeterReportsResponseBody
	GetPageSize() *int32
	SetReports(v []*ListJMeterReportsResponseBodyReports) *ListJMeterReportsResponseBody
	GetReports() []*ListJMeterReportsResponseBodyReports
	SetRequestId(v string) *ListJMeterReportsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListJMeterReportsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListJMeterReportsResponseBody
	GetTotalCount() *int64
}

type ListJMeterReportsResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, this parameter is left empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of the returned report page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of returned reports.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The reports.
	Reports []*ListJMeterReportsResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of reports returned based on the condition.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListJMeterReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListJMeterReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJMeterReportsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListJMeterReportsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListJMeterReportsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListJMeterReportsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListJMeterReportsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListJMeterReportsResponseBody) GetReports() []*ListJMeterReportsResponseBodyReports {
	return s.Reports
}

func (s *ListJMeterReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListJMeterReportsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListJMeterReportsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListJMeterReportsResponseBody) SetCode(v string) *ListJMeterReportsResponseBody {
	s.Code = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetHttpStatusCode(v int32) *ListJMeterReportsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetMessage(v string) *ListJMeterReportsResponseBody {
	s.Message = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetPageNumber(v int32) *ListJMeterReportsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetPageSize(v int32) *ListJMeterReportsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetReports(v []*ListJMeterReportsResponseBodyReports) *ListJMeterReportsResponseBody {
	s.Reports = v
	return s
}

func (s *ListJMeterReportsResponseBody) SetRequestId(v string) *ListJMeterReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetSuccess(v bool) *ListJMeterReportsResponseBody {
	s.Success = &v
	return s
}

func (s *ListJMeterReportsResponseBody) SetTotalCount(v int64) *ListJMeterReportsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJMeterReportsResponseBody) Validate() error {
	if s.Reports != nil {
		for _, item := range s.Reports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListJMeterReportsResponseBodyReports struct {
	// The start time of the stress testing.
	//
	// example:
	//
	// 1637157073000
	ActualStartTime *int64 `json:"ActualStartTime,omitempty" xml:"ActualStartTime,omitempty"`
	// The stress testing duration.
	//
	// example:
	//
	// 10分钟
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The report ID.
	//
	// example:
	//
	// 7R4RE352
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The report name.
	//
	// example:
	//
	// test
	ReportName *string `json:"ReportName,omitempty" xml:"ReportName,omitempty"`
	// The consumed Virtual User Minutes (VUM).
	//
	// example:
	//
	// 1000
	Vum *int64 `json:"Vum,omitempty" xml:"Vum,omitempty"`
}

func (s ListJMeterReportsResponseBodyReports) String() string {
	return dara.Prettify(s)
}

func (s ListJMeterReportsResponseBodyReports) GoString() string {
	return s.String()
}

func (s *ListJMeterReportsResponseBodyReports) GetActualStartTime() *int64 {
	return s.ActualStartTime
}

func (s *ListJMeterReportsResponseBodyReports) GetDuration() *string {
	return s.Duration
}

func (s *ListJMeterReportsResponseBodyReports) GetReportId() *string {
	return s.ReportId
}

func (s *ListJMeterReportsResponseBodyReports) GetReportName() *string {
	return s.ReportName
}

func (s *ListJMeterReportsResponseBodyReports) GetVum() *int64 {
	return s.Vum
}

func (s *ListJMeterReportsResponseBodyReports) SetActualStartTime(v int64) *ListJMeterReportsResponseBodyReports {
	s.ActualStartTime = &v
	return s
}

func (s *ListJMeterReportsResponseBodyReports) SetDuration(v string) *ListJMeterReportsResponseBodyReports {
	s.Duration = &v
	return s
}

func (s *ListJMeterReportsResponseBodyReports) SetReportId(v string) *ListJMeterReportsResponseBodyReports {
	s.ReportId = &v
	return s
}

func (s *ListJMeterReportsResponseBodyReports) SetReportName(v string) *ListJMeterReportsResponseBodyReports {
	s.ReportName = &v
	return s
}

func (s *ListJMeterReportsResponseBodyReports) SetVum(v int64) *ListJMeterReportsResponseBodyReports {
	s.Vum = &v
	return s
}

func (s *ListJMeterReportsResponseBodyReports) Validate() error {
	return dara.Validate(s)
}

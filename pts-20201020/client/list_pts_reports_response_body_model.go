// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPtsReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPtsReportsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListPtsReportsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPtsReportsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListPtsReportsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPtsReportsResponseBody
	GetPageSize() *int32
	SetReports(v []*ListPtsReportsResponseBodyReports) *ListPtsReportsResponseBody
	GetReports() []*ListPtsReportsResponseBodyReports
	SetRequestId(v string) *ListPtsReportsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPtsReportsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListPtsReportsResponseBody
	GetTotalCount() *int64
}

type ListPtsReportsResponseBody struct {
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
	// The returned message. If the request was successful, an empty string is returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of the returned page. The page number starts from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of reports returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The reports.
	Reports []*ListPtsReportsResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A8E4LR80-15P1-555A-9ZZF-B736AZO5E5ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true false
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

func (s ListPtsReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPtsReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPtsReportsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPtsReportsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPtsReportsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPtsReportsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPtsReportsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPtsReportsResponseBody) GetReports() []*ListPtsReportsResponseBodyReports {
	return s.Reports
}

func (s *ListPtsReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPtsReportsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPtsReportsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPtsReportsResponseBody) SetCode(v string) *ListPtsReportsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPtsReportsResponseBody) SetHttpStatusCode(v int32) *ListPtsReportsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPtsReportsResponseBody) SetMessage(v string) *ListPtsReportsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPtsReportsResponseBody) SetPageNumber(v int32) *ListPtsReportsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPtsReportsResponseBody) SetPageSize(v int32) *ListPtsReportsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPtsReportsResponseBody) SetReports(v []*ListPtsReportsResponseBodyReports) *ListPtsReportsResponseBody {
	s.Reports = v
	return s
}

func (s *ListPtsReportsResponseBody) SetRequestId(v string) *ListPtsReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPtsReportsResponseBody) SetSuccess(v bool) *ListPtsReportsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPtsReportsResponseBody) SetTotalCount(v int64) *ListPtsReportsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPtsReportsResponseBody) Validate() error {
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

type ListPtsReportsResponseBodyReports struct {
	// The timestamp when the stress testing starts. Unit: ms.
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
	// 7RLPM3Y2
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

func (s ListPtsReportsResponseBodyReports) String() string {
	return dara.Prettify(s)
}

func (s ListPtsReportsResponseBodyReports) GoString() string {
	return s.String()
}

func (s *ListPtsReportsResponseBodyReports) GetActualStartTime() *int64 {
	return s.ActualStartTime
}

func (s *ListPtsReportsResponseBodyReports) GetDuration() *string {
	return s.Duration
}

func (s *ListPtsReportsResponseBodyReports) GetReportId() *string {
	return s.ReportId
}

func (s *ListPtsReportsResponseBodyReports) GetReportName() *string {
	return s.ReportName
}

func (s *ListPtsReportsResponseBodyReports) GetVum() *int64 {
	return s.Vum
}

func (s *ListPtsReportsResponseBodyReports) SetActualStartTime(v int64) *ListPtsReportsResponseBodyReports {
	s.ActualStartTime = &v
	return s
}

func (s *ListPtsReportsResponseBodyReports) SetDuration(v string) *ListPtsReportsResponseBodyReports {
	s.Duration = &v
	return s
}

func (s *ListPtsReportsResponseBodyReports) SetReportId(v string) *ListPtsReportsResponseBodyReports {
	s.ReportId = &v
	return s
}

func (s *ListPtsReportsResponseBodyReports) SetReportName(v string) *ListPtsReportsResponseBodyReports {
	s.ReportName = &v
	return s
}

func (s *ListPtsReportsResponseBodyReports) SetVum(v int64) *ListPtsReportsResponseBodyReports {
	s.Vum = &v
	return s
}

func (s *ListPtsReportsResponseBodyReports) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSqlPatternCompareReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCancelTime(v string) *CancelSqlPatternCompareReportResponseBody
	GetCancelTime() *string
	SetCanceled(v bool) *CancelSqlPatternCompareReportResponseBody
	GetCanceled() *bool
	SetReportId(v int64) *CancelSqlPatternCompareReportResponseBody
	GetReportId() *int64
	SetRequestId(v string) *CancelSqlPatternCompareReportResponseBody
	GetRequestId() *string
}

type CancelSqlPatternCompareReportResponseBody struct {
	// The time when the report was first canceled. The time is in UTC in the yyyy-MM-ddTHH:mmZ format.
	//
	// example:
	//
	// 2026-09-08T01:06Z
	CancelTime *string `json:"CancelTime,omitempty" xml:"CancelTime,omitempty"`
	// Indicates whether the report is canceled. The value true is returned when the report is successfully canceled or canceled again.
	//
	// example:
	//
	// true
	Canceled *bool `json:"Canceled,omitempty" xml:"Canceled,omitempty"`
	// The SQL Pattern comparison report ID.
	//
	// example:
	//
	// 1001
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A1B2C3D-4E5F-6789-ABCD-0123456789AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelSqlPatternCompareReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelSqlPatternCompareReportResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSqlPatternCompareReportResponseBody) GetCancelTime() *string {
	return s.CancelTime
}

func (s *CancelSqlPatternCompareReportResponseBody) GetCanceled() *bool {
	return s.Canceled
}

func (s *CancelSqlPatternCompareReportResponseBody) GetReportId() *int64 {
	return s.ReportId
}

func (s *CancelSqlPatternCompareReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelSqlPatternCompareReportResponseBody) SetCancelTime(v string) *CancelSqlPatternCompareReportResponseBody {
	s.CancelTime = &v
	return s
}

func (s *CancelSqlPatternCompareReportResponseBody) SetCanceled(v bool) *CancelSqlPatternCompareReportResponseBody {
	s.Canceled = &v
	return s
}

func (s *CancelSqlPatternCompareReportResponseBody) SetReportId(v int64) *CancelSqlPatternCompareReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *CancelSqlPatternCompareReportResponseBody) SetRequestId(v string) *CancelSqlPatternCompareReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSqlPatternCompareReportResponseBody) Validate() error {
	return dara.Validate(s)
}

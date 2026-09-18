// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSqlPatternCompareReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v int64) *CreateSqlPatternCompareReportResponseBody
	GetReportId() *int64
	SetRequestId(v string) *CreateSqlPatternCompareReportResponseBody
	GetRequestId() *string
}

type CreateSqlPatternCompareReportResponseBody struct {
	// The ID of the created report. This value only indicates that the request has been accepted.
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

func (s CreateSqlPatternCompareReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSqlPatternCompareReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSqlPatternCompareReportResponseBody) GetReportId() *int64 {
	return s.ReportId
}

func (s *CreateSqlPatternCompareReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSqlPatternCompareReportResponseBody) SetReportId(v int64) *CreateSqlPatternCompareReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *CreateSqlPatternCompareReportResponseBody) SetRequestId(v string) *CreateSqlPatternCompareReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSqlPatternCompareReportResponseBody) Validate() error {
	return dara.Validate(s)
}

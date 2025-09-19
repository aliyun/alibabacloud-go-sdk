// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCustomizeReportConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v int64) *SaveCustomizeReportConfigResponseBody
	GetReportId() *int64
	SetRequestId(v string) *SaveCustomizeReportConfigResponseBody
	GetRequestId() *string
}

type SaveCustomizeReportConfigResponseBody struct {
	// The ID of the report.
	//
	// example:
	//
	// 123
	ReportId *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 11472B29-1A1C-5D7F-944B-7CD84319****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveCustomizeReportConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveCustomizeReportConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SaveCustomizeReportConfigResponseBody) GetReportId() *int64 {
	return s.ReportId
}

func (s *SaveCustomizeReportConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveCustomizeReportConfigResponseBody) SetReportId(v int64) *SaveCustomizeReportConfigResponseBody {
	s.ReportId = &v
	return s
}

func (s *SaveCustomizeReportConfigResponseBody) SetRequestId(v string) *SaveCustomizeReportConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveCustomizeReportConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

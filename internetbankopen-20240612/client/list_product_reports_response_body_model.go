// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListProductReportsResponseBodyData) *ListProductReportsResponseBody
	GetData() *ListProductReportsResponseBodyData
	SetRequestId(v string) *ListProductReportsResponseBody
	GetRequestId() *string
}

type ListProductReportsResponseBody struct {
	Data *ListProductReportsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 9634F2B3-B59C-5475-BA9F-96D95271882C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListProductReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductReportsResponseBody) GetData() *ListProductReportsResponseBodyData {
	return s.Data
}

func (s *ListProductReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductReportsResponseBody) SetData(v *ListProductReportsResponseBodyData) *ListProductReportsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductReportsResponseBody) SetRequestId(v string) *ListProductReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductReportsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductReportsResponseBodyData struct {
	Reports []*ListProductReportsResponseBodyDataReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Repeated"`
}

func (s ListProductReportsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProductReportsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductReportsResponseBodyData) GetReports() []*ListProductReportsResponseBodyDataReports {
	return s.Reports
}

func (s *ListProductReportsResponseBodyData) SetReports(v []*ListProductReportsResponseBodyDataReports) *ListProductReportsResponseBodyData {
	s.Reports = v
	return s
}

func (s *ListProductReportsResponseBodyData) Validate() error {
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

type ListProductReportsResponseBodyDataReports struct {
	// example:
	//
	// 2026-11-08T16:00:00Z
	PeriodEnd *string `json:"PeriodEnd,omitempty" xml:"PeriodEnd,omitempty"`
	// example:
	//
	// 2026-11-01T16:00:00Z
	PeriodStart *string `json:"PeriodStart,omitempty" xml:"PeriodStart,omitempty"`
	// example:
	//
	// cloudsec.demo.v1
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// report-3f2b8c1e-6d4a-4b7e-9a1c-2e5f7d8b9c0a
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s ListProductReportsResponseBodyDataReports) String() string {
	return dara.Prettify(s)
}

func (s ListProductReportsResponseBodyDataReports) GoString() string {
	return s.String()
}

func (s *ListProductReportsResponseBodyDataReports) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *ListProductReportsResponseBodyDataReports) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *ListProductReportsResponseBodyDataReports) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListProductReportsResponseBodyDataReports) GetReportId() *string {
	return s.ReportId
}

func (s *ListProductReportsResponseBodyDataReports) SetPeriodEnd(v string) *ListProductReportsResponseBodyDataReports {
	s.PeriodEnd = &v
	return s
}

func (s *ListProductReportsResponseBodyDataReports) SetPeriodStart(v string) *ListProductReportsResponseBodyDataReports {
	s.PeriodStart = &v
	return s
}

func (s *ListProductReportsResponseBodyDataReports) SetProductCode(v string) *ListProductReportsResponseBodyDataReports {
	s.ProductCode = &v
	return s
}

func (s *ListProductReportsResponseBodyDataReports) SetReportId(v string) *ListProductReportsResponseBodyDataReports {
	s.ReportId = &v
	return s
}

func (s *ListProductReportsResponseBodyDataReports) Validate() error {
	return dara.Validate(s)
}

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
	// The query results of the periodic report list, which includes the Reports array. The array is empty if no available reports exist.
	Data *ListProductReportsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identity of the request, which is used for Tracing Analysis and troubleshooting. Provide this identity when troubleshooting exceptions.
	//
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
	// The available report records sorted by period in descending order. By default, up to ten records are returned. If LatestOnly is set to true, at most one record is returned. The array is empty if no reports exist. A single natural week may have multiple report records; use ReportId to distinguish them. The list does not contain download URLs. The actual download status is subject to GetAiReport.
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
	// The end time of the natural week to which the report belongs, exclusive. The natural week ends at 00:00:00 on the following Monday in UTC+8 (Asia/Shanghai). This field is transformed to UTC RFC 3339 format and ends with Z. The example corresponds to 00:00:00 on Monday, November 9, 2026, UTC+8.
	//
	// example:
	//
	// 2026-11-08T16:00:00Z
	PeriodEnd *string `json:"PeriodEnd,omitempty" xml:"PeriodEnd,omitempty"`
	// The start time of the natural week to which the report belongs, inclusive. The natural week is calculated based on UTC+8 (Asia/Shanghai) and starts at 00:00:00 on Monday. This field is transformed to UTC RFC 3339 format and ends with Z. The example corresponds to 00:00:00 on Monday, November 2, 2026, UTC+8.
	//
	// example:
	//
	// 2026-11-01T16:00:00Z
	PeriodStart *string `json:"PeriodStart,omitempty" xml:"PeriodStart,omitempty"`
	// The product code of the data product to which the report belongs. This is consistent with the ProductCode of the current request.
	//
	// example:
	//
	// cloudsec.demo.v1
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The unique ID of the periodic report. Pass this value as the reportId parameter to GetAiReport to query the current status of the report and obtain the download URL when it is available.
	//
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

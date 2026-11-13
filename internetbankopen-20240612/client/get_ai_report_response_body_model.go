// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAiReportResponseBodyData) *GetAiReportResponseBody
	GetData() *GetAiReportResponseBodyData
	SetRequestId(v string) *GetAiReportResponseBody
	GetRequestId() *string
}

type GetAiReportResponseBody struct {
	// The report query result, which contains ReportStatus. The Report object is included only when the report is available.
	Data *GetAiReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique identity of the request, which is used for Tracing Analysis and troubleshooting. Provide this ID when you troubleshoot exceptions.
	//
	// example:
	//
	// 9634F2B3-B59C-5475-BA9F-96D95271882C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAiReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiReportResponseBody) GetData() *GetAiReportResponseBodyData {
	return s.Data
}

func (s *GetAiReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiReportResponseBody) SetData(v *GetAiReportResponseBodyData) *GetAiReportResponseBody {
	s.Data = v
	return s
}

func (s *GetAiReportResponseBody) SetRequestId(v string) *GetAiReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiReportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiReportResponseBodyData struct {
	// The download information of the available report, which is returned only when ReportStatus is AVAILABLE. The report ownership and the corresponding object version are verified before the download.
	Report *GetAiReportResponseBodyDataReport `json:"Report,omitempty" xml:"Report,omitempty" type:"Struct"`
	// The report status. Valid values:
	//
	// - NOT_APPLICABLE: The product does not include a report.
	//
	// - WAITING_DATA: The data call result is not determined.
	//
	// - NO_DATA: No data is found, and no per-use report is generated.
	//
	// - WAITING_PERIOD: The query result for the callId of a periodic product. You must obtain the ReportId by calling ListProductReports and then query the report.
	//
	// - QUEUED: The report task is queued.
	//
	// - RUNNING: The report task is being processed.
	//
	// - AVAILABLE: The report is available, and the Report object is returned in this response.
	//
	// - FAILED: The data call or report generation failed.
	//
	// - EXPIRED: The report content has expired.
	//
	// - UNAVAILABLE: The delivery of the report is suspended.
	//
	// You can query again later by using the same ID when the status is WAITING_DATA, QUEUED, or RUNNING. Repeated queries do not regenerate the report when the status is FAILED. The Report object is not returned unless the status is AVAILABLE.
	//
	// example:
	//
	// AVAILABLE
	ReportStatus *string `json:"ReportStatus,omitempty" xml:"ReportStatus,omitempty"`
}

func (s GetAiReportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAiReportResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAiReportResponseBodyData) GetReport() *GetAiReportResponseBodyDataReport {
	return s.Report
}

func (s *GetAiReportResponseBodyData) GetReportStatus() *string {
	return s.ReportStatus
}

func (s *GetAiReportResponseBodyData) SetReport(v *GetAiReportResponseBodyDataReport) *GetAiReportResponseBodyData {
	s.Report = v
	return s
}

func (s *GetAiReportResponseBodyData) SetReportStatus(v string) *GetAiReportResponseBodyData {
	s.ReportStatus = &v
	return s
}

func (s *GetAiReportResponseBodyData) Validate() error {
	if s.Report != nil {
		if err := s.Report.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAiReportResponseBodyDataReport struct {
	// The expiration time of the report content retention period. The time is in UTC RFC 3339 format, ends with Z, and may contain fractional seconds. No download URL is issued after this time. If the delivery is not suspended, EXPIRED is returned. If the delivery is also suspended, UNAVAILABLE is returned preferentially. This is the content expiration time, not the expiration time of a single URL.
	//
	// example:
	//
	// 2026-12-12T06:20:00Z
	ContentExpiresAt *string `json:"ContentExpiresAt,omitempty" xml:"ContentExpiresAt,omitempty"`
	// The temporary signed download URL. Use the URL before UrlExpiresAt. The validity period of the signed URL does not exceed 300 seconds and is no later than ContentExpiresAt. If the URL expires but the content has not expired, you can query again to obtain a new URL. The example only shows the URL format and is not a real downloadable link. Use the complete URL in the actual response. Do not concatenate the URL yourself or share it publicly.
	//
	// example:
	//
	// https://reports.example.com/report-3f2b8c1e-6d4a-4b7e-9a1c-2e5f7d8b9c0a.pdf
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The unique ID of the report. You can use this ID as the reportId parameter to query the report again. This is the report asset ID, not an internal task ID.
	//
	// example:
	//
	// report-3f2b8c1e-6d4a-4b7e-9a1c-2e5f7d8b9c0a
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The expiration time of the current DownloadUrl. The time is in UTC RFC 3339 format, ends with Z, and may contain fractional seconds. You cannot use the URL after it expires, and this time is no later than ContentExpiresAt.
	//
	// example:
	//
	// 2026-11-12T06:25:00Z
	UrlExpiresAt *string `json:"UrlExpiresAt,omitempty" xml:"UrlExpiresAt,omitempty"`
}

func (s GetAiReportResponseBodyDataReport) String() string {
	return dara.Prettify(s)
}

func (s GetAiReportResponseBodyDataReport) GoString() string {
	return s.String()
}

func (s *GetAiReportResponseBodyDataReport) GetContentExpiresAt() *string {
	return s.ContentExpiresAt
}

func (s *GetAiReportResponseBodyDataReport) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetAiReportResponseBodyDataReport) GetReportId() *string {
	return s.ReportId
}

func (s *GetAiReportResponseBodyDataReport) GetUrlExpiresAt() *string {
	return s.UrlExpiresAt
}

func (s *GetAiReportResponseBodyDataReport) SetContentExpiresAt(v string) *GetAiReportResponseBodyDataReport {
	s.ContentExpiresAt = &v
	return s
}

func (s *GetAiReportResponseBodyDataReport) SetDownloadUrl(v string) *GetAiReportResponseBodyDataReport {
	s.DownloadUrl = &v
	return s
}

func (s *GetAiReportResponseBodyDataReport) SetReportId(v string) *GetAiReportResponseBodyDataReport {
	s.ReportId = &v
	return s
}

func (s *GetAiReportResponseBodyDataReport) SetUrlExpiresAt(v string) *GetAiReportResponseBodyDataReport {
	s.UrlExpiresAt = &v
	return s
}

func (s *GetAiReportResponseBodyDataReport) Validate() error {
	return dara.Validate(s)
}

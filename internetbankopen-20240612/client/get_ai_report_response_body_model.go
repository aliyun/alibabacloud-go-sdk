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
	Data *GetAiReportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	Report *GetAiReportResponseBodyDataReport `json:"Report,omitempty" xml:"Report,omitempty" type:"Struct"`
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
	// example:
	//
	// 2026-12-12T06:20:00Z
	ContentExpiresAt *string `json:"ContentExpiresAt,omitempty" xml:"ContentExpiresAt,omitempty"`
	// example:
	//
	// https://reports.example.com/report-3f2b8c1e-6d4a-4b7e-9a1c-2e5f7d8b9c0a.pdf
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// report-3f2b8c1e-6d4a-4b7e-9a1c-2e5f7d8b9c0a
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
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

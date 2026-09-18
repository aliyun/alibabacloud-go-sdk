// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScannerTaskResultHtmlUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetScannerTaskResultHtmlUrlResponseBodyData) *GetScannerTaskResultHtmlUrlResponseBody
	GetData() *GetScannerTaskResultHtmlUrlResponseBodyData
	SetRequestId(v string) *GetScannerTaskResultHtmlUrlResponseBody
	GetRequestId() *string
}

type GetScannerTaskResultHtmlUrlResponseBody struct {
	// The query result, which contains the temporary download URL for the HTML result report.
	Data *GetScannerTaskResultHtmlUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, which is used for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-B7AB7B5A9C8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetScannerTaskResultHtmlUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScannerTaskResultHtmlUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetScannerTaskResultHtmlUrlResponseBody) GetData() *GetScannerTaskResultHtmlUrlResponseBodyData {
	return s.Data
}

func (s *GetScannerTaskResultHtmlUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScannerTaskResultHtmlUrlResponseBody) SetData(v *GetScannerTaskResultHtmlUrlResponseBodyData) *GetScannerTaskResultHtmlUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetScannerTaskResultHtmlUrlResponseBody) SetRequestId(v string) *GetScannerTaskResultHtmlUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScannerTaskResultHtmlUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScannerTaskResultHtmlUrlResponseBodyData struct {
	// The temporary download URL for the HTML result report. The URL is valid for 2 hours. After the URL expires, call this operation again to obtain a new URL. If the report has not been generated, this value is an empty string. The actual value is a signed temporary URL from object storage that includes signature parameters.
	//
	// example:
	//
	// https://example.com/result.task-abc123def4567.report.html
	ScannerTaskResultHtmlDownloadUrl *string `json:"ScannerTaskResultHtmlDownloadUrl,omitempty" xml:"ScannerTaskResultHtmlDownloadUrl,omitempty"`
}

func (s GetScannerTaskResultHtmlUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetScannerTaskResultHtmlUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetScannerTaskResultHtmlUrlResponseBodyData) GetScannerTaskResultHtmlDownloadUrl() *string {
	return s.ScannerTaskResultHtmlDownloadUrl
}

func (s *GetScannerTaskResultHtmlUrlResponseBodyData) SetScannerTaskResultHtmlDownloadUrl(v string) *GetScannerTaskResultHtmlUrlResponseBodyData {
	s.ScannerTaskResultHtmlDownloadUrl = &v
	return s
}

func (s *GetScannerTaskResultHtmlUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}

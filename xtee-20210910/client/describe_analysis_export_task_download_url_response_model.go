// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnalysisExportTaskDownloadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAnalysisExportTaskDownloadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAnalysisExportTaskDownloadUrlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAnalysisExportTaskDownloadUrlResponseBody) *DescribeAnalysisExportTaskDownloadUrlResponse
	GetBody() *DescribeAnalysisExportTaskDownloadUrlResponseBody
}

type DescribeAnalysisExportTaskDownloadUrlResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAnalysisExportTaskDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAnalysisExportTaskDownloadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnalysisExportTaskDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponse) GetBody() *DescribeAnalysisExportTaskDownloadUrlResponseBody {
	return s.Body
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponse) SetHeaders(v map[string]*string) *DescribeAnalysisExportTaskDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponse) SetStatusCode(v int32) *DescribeAnalysisExportTaskDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponse) SetBody(v *DescribeAnalysisExportTaskDownloadUrlResponseBody) *DescribeAnalysisExportTaskDownloadUrlResponse {
	s.Body = v
	return s
}

func (s *DescribeAnalysisExportTaskDownloadUrlResponse) Validate() error {
	return dara.Validate(s)
}

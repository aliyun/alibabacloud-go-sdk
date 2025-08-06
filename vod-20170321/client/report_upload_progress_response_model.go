// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportUploadProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportUploadProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportUploadProgressResponse
	GetStatusCode() *int32
	SetBody(v *ReportUploadProgressResponseBody) *ReportUploadProgressResponse
	GetBody() *ReportUploadProgressResponseBody
}

type ReportUploadProgressResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportUploadProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportUploadProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportUploadProgressResponse) GoString() string {
	return s.String()
}

func (s *ReportUploadProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportUploadProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportUploadProgressResponse) GetBody() *ReportUploadProgressResponseBody {
	return s.Body
}

func (s *ReportUploadProgressResponse) SetHeaders(v map[string]*string) *ReportUploadProgressResponse {
	s.Headers = v
	return s
}

func (s *ReportUploadProgressResponse) SetStatusCode(v int32) *ReportUploadProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportUploadProgressResponse) SetBody(v *ReportUploadProgressResponseBody) *ReportUploadProgressResponse {
	s.Body = v
	return s
}

func (s *ReportUploadProgressResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticReportListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosticReportListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosticReportListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosticReportListResponseBody) *DescribeDiagnosticReportListResponse
	GetBody() *DescribeDiagnosticReportListResponseBody
}

type DescribeDiagnosticReportListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosticReportListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosticReportListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosticReportListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosticReportListResponse) GetBody() *DescribeDiagnosticReportListResponseBody {
	return s.Body
}

func (s *DescribeDiagnosticReportListResponse) SetHeaders(v map[string]*string) *DescribeDiagnosticReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosticReportListResponse) SetStatusCode(v int32) *DescribeDiagnosticReportListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosticReportListResponse) SetBody(v *DescribeDiagnosticReportListResponseBody) *DescribeDiagnosticReportListResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnosticReportListResponse) Validate() error {
	return dara.Validate(s)
}

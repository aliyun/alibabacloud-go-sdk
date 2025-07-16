// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBenchmarkTaskReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBenchmarkTaskReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBenchmarkTaskReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBenchmarkTaskReportResponseBody) *DescribeBenchmarkTaskReportResponse
	GetBody() *DescribeBenchmarkTaskReportResponseBody
}

type DescribeBenchmarkTaskReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBenchmarkTaskReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBenchmarkTaskReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBenchmarkTaskReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBenchmarkTaskReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBenchmarkTaskReportResponse) GetBody() *DescribeBenchmarkTaskReportResponseBody {
	return s.Body
}

func (s *DescribeBenchmarkTaskReportResponse) SetHeaders(v map[string]*string) *DescribeBenchmarkTaskReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeBenchmarkTaskReportResponse) SetStatusCode(v int32) *DescribeBenchmarkTaskReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBenchmarkTaskReportResponse) SetBody(v *DescribeBenchmarkTaskReportResponseBody) *DescribeBenchmarkTaskReportResponse {
	s.Body = v
	return s
}

func (s *DescribeBenchmarkTaskReportResponse) Validate() error {
	return dara.Validate(s)
}

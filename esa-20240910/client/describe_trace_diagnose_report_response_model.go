// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceDiagnoseReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTraceDiagnoseReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTraceDiagnoseReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTraceDiagnoseReportResponseBody) *DescribeTraceDiagnoseReportResponse
	GetBody() *DescribeTraceDiagnoseReportResponseBody
}

type DescribeTraceDiagnoseReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTraceDiagnoseReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTraceDiagnoseReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceDiagnoseReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeTraceDiagnoseReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTraceDiagnoseReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTraceDiagnoseReportResponse) GetBody() *DescribeTraceDiagnoseReportResponseBody {
	return s.Body
}

func (s *DescribeTraceDiagnoseReportResponse) SetHeaders(v map[string]*string) *DescribeTraceDiagnoseReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeTraceDiagnoseReportResponse) SetStatusCode(v int32) *DescribeTraceDiagnoseReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTraceDiagnoseReportResponse) SetBody(v *DescribeTraceDiagnoseReportResponseBody) *DescribeTraceDiagnoseReportResponse {
	s.Body = v
	return s
}

func (s *DescribeTraceDiagnoseReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

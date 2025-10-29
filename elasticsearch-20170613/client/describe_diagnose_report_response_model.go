// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnoseReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnoseReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnoseReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnoseReportResponseBody) *DescribeDiagnoseReportResponse
	GetBody() *DescribeDiagnoseReportResponseBody
}

type DescribeDiagnoseReportResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnoseReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnoseReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnoseReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnoseReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnoseReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnoseReportResponse) GetBody() *DescribeDiagnoseReportResponseBody {
	return s.Body
}

func (s *DescribeDiagnoseReportResponse) SetHeaders(v map[string]*string) *DescribeDiagnoseReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnoseReportResponse) SetStatusCode(v int32) *DescribeDiagnoseReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnoseReportResponse) SetBody(v *DescribeDiagnoseReportResponseBody) *DescribeDiagnoseReportResponse {
	s.Body = v
	return s
}

func (s *DescribeDiagnoseReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

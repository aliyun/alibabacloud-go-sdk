// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDiagnoseReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnDiagnoseReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnDiagnoseReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnDiagnoseReportResponseBody) *DescribeCdnDiagnoseReportResponse
	GetBody() *DescribeCdnDiagnoseReportResponseBody
}

type DescribeCdnDiagnoseReportResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnDiagnoseReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnDiagnoseReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDiagnoseReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDiagnoseReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnDiagnoseReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnDiagnoseReportResponse) GetBody() *DescribeCdnDiagnoseReportResponseBody {
	return s.Body
}

func (s *DescribeCdnDiagnoseReportResponse) SetHeaders(v map[string]*string) *DescribeCdnDiagnoseReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDiagnoseReportResponse) SetStatusCode(v int32) *DescribeCdnDiagnoseReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDiagnoseReportResponse) SetBody(v *DescribeCdnDiagnoseReportResponseBody) *DescribeCdnDiagnoseReportResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnDiagnoseReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnReportResponseBody) *DescribeDcdnReportResponse
	GetBody() *DescribeDcdnReportResponseBody
}

type DescribeDcdnReportResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnReportResponse) GetBody() *DescribeDcdnReportResponseBody {
	return s.Body
}

func (s *DescribeDcdnReportResponse) SetHeaders(v map[string]*string) *DescribeDcdnReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnReportResponse) SetStatusCode(v int32) *DescribeDcdnReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnReportResponse) SetBody(v *DescribeDcdnReportResponseBody) *DescribeDcdnReportResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnReportResponse) Validate() error {
	return dara.Validate(s)
}

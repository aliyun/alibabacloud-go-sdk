// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnReportListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnReportListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnReportListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnReportListResponseBody) *DescribeDcdnReportListResponse
	GetBody() *DescribeDcdnReportListResponseBody
}

type DescribeDcdnReportListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnReportListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnReportListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnReportListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnReportListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnReportListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnReportListResponse) GetBody() *DescribeDcdnReportListResponseBody {
	return s.Body
}

func (s *DescribeDcdnReportListResponse) SetHeaders(v map[string]*string) *DescribeDcdnReportListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnReportListResponse) SetStatusCode(v int32) *DescribeDcdnReportListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnReportListResponse) SetBody(v *DescribeDcdnReportListResponseBody) *DescribeDcdnReportListResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnReportListResponse) Validate() error {
	return dara.Validate(s)
}

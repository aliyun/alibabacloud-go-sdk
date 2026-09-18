// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlPatternCompareReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlPatternCompareReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlPatternCompareReportResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlPatternCompareReportResponseBody) *DescribeSqlPatternCompareReportResponse
	GetBody() *DescribeSqlPatternCompareReportResponseBody
}

type DescribeSqlPatternCompareReportResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlPatternCompareReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlPatternCompareReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternCompareReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternCompareReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlPatternCompareReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlPatternCompareReportResponse) GetBody() *DescribeSqlPatternCompareReportResponseBody {
	return s.Body
}

func (s *DescribeSqlPatternCompareReportResponse) SetHeaders(v map[string]*string) *DescribeSqlPatternCompareReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlPatternCompareReportResponse) SetStatusCode(v int32) *DescribeSqlPatternCompareReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlPatternCompareReportResponse) SetBody(v *DescribeSqlPatternCompareReportResponseBody) *DescribeSqlPatternCompareReportResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlPatternCompareReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

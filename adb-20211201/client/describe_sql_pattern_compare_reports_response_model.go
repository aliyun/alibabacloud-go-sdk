// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSqlPatternCompareReportsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSqlPatternCompareReportsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSqlPatternCompareReportsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSqlPatternCompareReportsResponseBody) *DescribeSqlPatternCompareReportsResponse
	GetBody() *DescribeSqlPatternCompareReportsResponseBody
}

type DescribeSqlPatternCompareReportsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSqlPatternCompareReportsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSqlPatternCompareReportsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSqlPatternCompareReportsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSqlPatternCompareReportsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSqlPatternCompareReportsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSqlPatternCompareReportsResponse) GetBody() *DescribeSqlPatternCompareReportsResponseBody {
	return s.Body
}

func (s *DescribeSqlPatternCompareReportsResponse) SetHeaders(v map[string]*string) *DescribeSqlPatternCompareReportsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponse) SetStatusCode(v int32) *DescribeSqlPatternCompareReportsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponse) SetBody(v *DescribeSqlPatternCompareReportsResponseBody) *DescribeSqlPatternCompareReportsResponse {
	s.Body = v
	return s
}

func (s *DescribeSqlPatternCompareReportsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

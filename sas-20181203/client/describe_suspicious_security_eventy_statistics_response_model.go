// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSuspiciousSecurityEventyStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSuspiciousSecurityEventyStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSuspiciousSecurityEventyStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSuspiciousSecurityEventyStatisticsResponseBody) *DescribeSuspiciousSecurityEventyStatisticsResponse
	GetBody() *DescribeSuspiciousSecurityEventyStatisticsResponseBody
}

type DescribeSuspiciousSecurityEventyStatisticsResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSuspiciousSecurityEventyStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSuspiciousSecurityEventyStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSuspiciousSecurityEventyStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSuspiciousSecurityEventyStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSuspiciousSecurityEventyStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSuspiciousSecurityEventyStatisticsResponse) GetBody() *DescribeSuspiciousSecurityEventyStatisticsResponseBody {
	return s.Body
}

func (s *DescribeSuspiciousSecurityEventyStatisticsResponse) SetHeaders(v map[string]*string) *DescribeSuspiciousSecurityEventyStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSuspiciousSecurityEventyStatisticsResponse) SetStatusCode(v int32) *DescribeSuspiciousSecurityEventyStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSuspiciousSecurityEventyStatisticsResponse) SetBody(v *DescribeSuspiciousSecurityEventyStatisticsResponseBody) *DescribeSuspiciousSecurityEventyStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeSuspiciousSecurityEventyStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

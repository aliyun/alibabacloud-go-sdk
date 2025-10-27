// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainSecureStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainSecureStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainSecureStatisticsResponseBody) *DescribeDomainSecureStatisticsResponse
	GetBody() *DescribeDomainSecureStatisticsResponseBody
}

type DescribeDomainSecureStatisticsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainSecureStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainSecureStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainSecureStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainSecureStatisticsResponse) GetBody() *DescribeDomainSecureStatisticsResponseBody {
	return s.Body
}

func (s *DescribeDomainSecureStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDomainSecureStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainSecureStatisticsResponse) SetStatusCode(v int32) *DescribeDomainSecureStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainSecureStatisticsResponse) SetBody(v *DescribeDomainSecureStatisticsResponseBody) *DescribeDomainSecureStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainSecureStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

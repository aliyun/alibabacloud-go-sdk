// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohDomainStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDohDomainStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDohDomainStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDohDomainStatisticsResponseBody) *DescribeDohDomainStatisticsResponse
	GetBody() *DescribeDohDomainStatisticsResponseBody
}

type DescribeDohDomainStatisticsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDohDomainStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDohDomainStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohDomainStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohDomainStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDohDomainStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDohDomainStatisticsResponse) GetBody() *DescribeDohDomainStatisticsResponseBody {
	return s.Body
}

func (s *DescribeDohDomainStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDohDomainStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohDomainStatisticsResponse) SetStatusCode(v int32) *DescribeDohDomainStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDohDomainStatisticsResponse) SetBody(v *DescribeDohDomainStatisticsResponseBody) *DescribeDohDomainStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeDohDomainStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

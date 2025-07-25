// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohSubDomainStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDohSubDomainStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDohSubDomainStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDohSubDomainStatisticsResponseBody) *DescribeDohSubDomainStatisticsResponse
	GetBody() *DescribeDohSubDomainStatisticsResponseBody
}

type DescribeDohSubDomainStatisticsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDohSubDomainStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDohSubDomainStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohSubDomainStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohSubDomainStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDohSubDomainStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDohSubDomainStatisticsResponse) GetBody() *DescribeDohSubDomainStatisticsResponseBody {
	return s.Body
}

func (s *DescribeDohSubDomainStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDohSubDomainStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponse) SetStatusCode(v int32) *DescribeDohSubDomainStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponse) SetBody(v *DescribeDohSubDomainStatisticsResponseBody) *DescribeDohSubDomainStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeDohSubDomainStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainStatisticsResponseBody) *DescribeDomainStatisticsResponse
	GetBody() *DescribeDomainStatisticsResponseBody
}

type DescribeDomainStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainStatisticsResponse) GetBody() *DescribeDomainStatisticsResponseBody {
	return s.Body
}

func (s *DescribeDomainStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDomainStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainStatisticsResponse) SetStatusCode(v int32) *DescribeDomainStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainStatisticsResponse) SetBody(v *DescribeDomainStatisticsResponseBody) *DescribeDomainStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

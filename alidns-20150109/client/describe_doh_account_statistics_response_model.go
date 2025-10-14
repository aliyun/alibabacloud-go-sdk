// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDohAccountStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDohAccountStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDohAccountStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDohAccountStatisticsResponseBody) *DescribeDohAccountStatisticsResponse
	GetBody() *DescribeDohAccountStatisticsResponseBody
}

type DescribeDohAccountStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDohAccountStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDohAccountStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDohAccountStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDohAccountStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDohAccountStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDohAccountStatisticsResponse) GetBody() *DescribeDohAccountStatisticsResponseBody {
	return s.Body
}

func (s *DescribeDohAccountStatisticsResponse) SetHeaders(v map[string]*string) *DescribeDohAccountStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDohAccountStatisticsResponse) SetStatusCode(v int32) *DescribeDohAccountStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDohAccountStatisticsResponse) SetBody(v *DescribeDohAccountStatisticsResponseBody) *DescribeDohAccountStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeDohAccountStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

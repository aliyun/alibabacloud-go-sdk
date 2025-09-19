// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTotalStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTotalStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTotalStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTotalStatisticsResponseBody) *DescribeTotalStatisticsResponse
	GetBody() *DescribeTotalStatisticsResponseBody
}

type DescribeTotalStatisticsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTotalStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTotalStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTotalStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTotalStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTotalStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTotalStatisticsResponse) GetBody() *DescribeTotalStatisticsResponseBody {
	return s.Body
}

func (s *DescribeTotalStatisticsResponse) SetHeaders(v map[string]*string) *DescribeTotalStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTotalStatisticsResponse) SetStatusCode(v int32) *DescribeTotalStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTotalStatisticsResponse) SetBody(v *DescribeTotalStatisticsResponseBody) *DescribeTotalStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeTotalStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

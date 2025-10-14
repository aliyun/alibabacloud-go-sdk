// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePdnsThreatStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePdnsThreatStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePdnsThreatStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePdnsThreatStatisticsResponseBody) *DescribePdnsThreatStatisticsResponse
	GetBody() *DescribePdnsThreatStatisticsResponseBody
}

type DescribePdnsThreatStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePdnsThreatStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePdnsThreatStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePdnsThreatStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribePdnsThreatStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePdnsThreatStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePdnsThreatStatisticsResponse) GetBody() *DescribePdnsThreatStatisticsResponseBody {
	return s.Body
}

func (s *DescribePdnsThreatStatisticsResponse) SetHeaders(v map[string]*string) *DescribePdnsThreatStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribePdnsThreatStatisticsResponse) SetStatusCode(v int32) *DescribePdnsThreatStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePdnsThreatStatisticsResponse) SetBody(v *DescribePdnsThreatStatisticsResponseBody) *DescribePdnsThreatStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribePdnsThreatStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

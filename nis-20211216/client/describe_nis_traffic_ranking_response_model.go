// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisTrafficRankingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNisTrafficRankingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNisTrafficRankingResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNisTrafficRankingResponseBody) *DescribeNisTrafficRankingResponse
	GetBody() *DescribeNisTrafficRankingResponseBody
}

type DescribeNisTrafficRankingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNisTrafficRankingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNisTrafficRankingResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisTrafficRankingResponse) GoString() string {
	return s.String()
}

func (s *DescribeNisTrafficRankingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNisTrafficRankingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNisTrafficRankingResponse) GetBody() *DescribeNisTrafficRankingResponseBody {
	return s.Body
}

func (s *DescribeNisTrafficRankingResponse) SetHeaders(v map[string]*string) *DescribeNisTrafficRankingResponse {
	s.Headers = v
	return s
}

func (s *DescribeNisTrafficRankingResponse) SetStatusCode(v int32) *DescribeNisTrafficRankingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNisTrafficRankingResponse) SetBody(v *DescribeNisTrafficRankingResponseBody) *DescribeNisTrafficRankingResponse {
	s.Body = v
	return s
}

func (s *DescribeNisTrafficRankingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

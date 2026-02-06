// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntentStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIntentStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIntentStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIntentStatisticsResponseBody) *DescribeIntentStatisticsResponse
	GetBody() *DescribeIntentStatisticsResponseBody
}

type DescribeIntentStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIntentStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIntentStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeIntentStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIntentStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIntentStatisticsResponse) GetBody() *DescribeIntentStatisticsResponseBody {
	return s.Body
}

func (s *DescribeIntentStatisticsResponse) SetHeaders(v map[string]*string) *DescribeIntentStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeIntentStatisticsResponse) SetStatusCode(v int32) *DescribeIntentStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIntentStatisticsResponse) SetBody(v *DescribeIntentStatisticsResponseBody) *DescribeIntentStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeIntentStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingRiskTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOutgoingRiskTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOutgoingRiskTrendResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOutgoingRiskTrendResponseBody) *DescribeOutgoingRiskTrendResponse
	GetBody() *DescribeOutgoingRiskTrendResponseBody
}

type DescribeOutgoingRiskTrendResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOutgoingRiskTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOutgoingRiskTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingRiskTrendResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingRiskTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOutgoingRiskTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOutgoingRiskTrendResponse) GetBody() *DescribeOutgoingRiskTrendResponseBody {
	return s.Body
}

func (s *DescribeOutgoingRiskTrendResponse) SetHeaders(v map[string]*string) *DescribeOutgoingRiskTrendResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutgoingRiskTrendResponse) SetStatusCode(v int32) *DescribeOutgoingRiskTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutgoingRiskTrendResponse) SetBody(v *DescribeOutgoingRiskTrendResponseBody) *DescribeOutgoingRiskTrendResponse {
	s.Body = v
	return s
}

func (s *DescribeOutgoingRiskTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

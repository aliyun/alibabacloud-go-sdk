// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageBaselineStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeImageBaselineStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeImageBaselineStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeImageBaselineStrategyResponseBody) *DescribeImageBaselineStrategyResponse
	GetBody() *DescribeImageBaselineStrategyResponseBody
}

type DescribeImageBaselineStrategyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeImageBaselineStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeImageBaselineStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageBaselineStrategyResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageBaselineStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeImageBaselineStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeImageBaselineStrategyResponse) GetBody() *DescribeImageBaselineStrategyResponseBody {
	return s.Body
}

func (s *DescribeImageBaselineStrategyResponse) SetHeaders(v map[string]*string) *DescribeImageBaselineStrategyResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageBaselineStrategyResponse) SetStatusCode(v int32) *DescribeImageBaselineStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeImageBaselineStrategyResponse) SetBody(v *DescribeImageBaselineStrategyResponseBody) *DescribeImageBaselineStrategyResponse {
	s.Body = v
	return s
}

func (s *DescribeImageBaselineStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

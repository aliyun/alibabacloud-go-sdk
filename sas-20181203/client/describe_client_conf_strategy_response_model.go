// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientConfStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClientConfStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClientConfStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClientConfStrategyResponseBody) *DescribeClientConfStrategyResponse
	GetBody() *DescribeClientConfStrategyResponseBody
}

type DescribeClientConfStrategyResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClientConfStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClientConfStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientConfStrategyResponse) GoString() string {
	return s.String()
}

func (s *DescribeClientConfStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClientConfStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClientConfStrategyResponse) GetBody() *DescribeClientConfStrategyResponseBody {
	return s.Body
}

func (s *DescribeClientConfStrategyResponse) SetHeaders(v map[string]*string) *DescribeClientConfStrategyResponse {
	s.Headers = v
	return s
}

func (s *DescribeClientConfStrategyResponse) SetStatusCode(v int32) *DescribeClientConfStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClientConfStrategyResponse) SetBody(v *DescribeClientConfStrategyResponseBody) *DescribeClientConfStrategyResponse {
	s.Body = v
	return s
}

func (s *DescribeClientConfStrategyResponse) Validate() error {
	return dara.Validate(s)
}

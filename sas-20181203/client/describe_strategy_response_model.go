// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStrategyResponseBody) *DescribeStrategyResponse
	GetBody() *DescribeStrategyResponseBody
}

type DescribeStrategyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyResponse) GoString() string {
	return s.String()
}

func (s *DescribeStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStrategyResponse) GetBody() *DescribeStrategyResponseBody {
	return s.Body
}

func (s *DescribeStrategyResponse) SetHeaders(v map[string]*string) *DescribeStrategyResponse {
	s.Headers = v
	return s
}

func (s *DescribeStrategyResponse) SetStatusCode(v int32) *DescribeStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStrategyResponse) SetBody(v *DescribeStrategyResponseBody) *DescribeStrategyResponse {
	s.Body = v
	return s
}

func (s *DescribeStrategyResponse) Validate() error {
	return dara.Validate(s)
}

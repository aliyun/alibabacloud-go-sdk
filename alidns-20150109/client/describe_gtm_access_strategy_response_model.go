// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmAccessStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmAccessStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmAccessStrategyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmAccessStrategyResponseBody) *DescribeGtmAccessStrategyResponse
	GetBody() *DescribeGtmAccessStrategyResponseBody
}

type DescribeGtmAccessStrategyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmAccessStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmAccessStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmAccessStrategyResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmAccessStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmAccessStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmAccessStrategyResponse) GetBody() *DescribeGtmAccessStrategyResponseBody {
	return s.Body
}

func (s *DescribeGtmAccessStrategyResponse) SetHeaders(v map[string]*string) *DescribeGtmAccessStrategyResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmAccessStrategyResponse) SetStatusCode(v int32) *DescribeGtmAccessStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmAccessStrategyResponse) SetBody(v *DescribeGtmAccessStrategyResponseBody) *DescribeGtmAccessStrategyResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmAccessStrategyResponse) Validate() error {
	return dara.Validate(s)
}

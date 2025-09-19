// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomizedStrategyTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomizedStrategyTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomizedStrategyTargetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomizedStrategyTargetsResponseBody) *DescribeCustomizedStrategyTargetsResponse
	GetBody() *DescribeCustomizedStrategyTargetsResponseBody
}

type DescribeCustomizedStrategyTargetsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomizedStrategyTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomizedStrategyTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomizedStrategyTargetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomizedStrategyTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomizedStrategyTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomizedStrategyTargetsResponse) GetBody() *DescribeCustomizedStrategyTargetsResponseBody {
	return s.Body
}

func (s *DescribeCustomizedStrategyTargetsResponse) SetHeaders(v map[string]*string) *DescribeCustomizedStrategyTargetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomizedStrategyTargetsResponse) SetStatusCode(v int32) *DescribeCustomizedStrategyTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomizedStrategyTargetsResponse) SetBody(v *DescribeCustomizedStrategyTargetsResponseBody) *DescribeCustomizedStrategyTargetsResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomizedStrategyTargetsResponse) Validate() error {
	return dara.Validate(s)
}

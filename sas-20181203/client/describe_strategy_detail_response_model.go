// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStrategyDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStrategyDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStrategyDetailResponseBody) *DescribeStrategyDetailResponse
	GetBody() *DescribeStrategyDetailResponseBody
}

type DescribeStrategyDetailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStrategyDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStrategyDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeStrategyDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStrategyDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStrategyDetailResponse) GetBody() *DescribeStrategyDetailResponseBody {
	return s.Body
}

func (s *DescribeStrategyDetailResponse) SetHeaders(v map[string]*string) *DescribeStrategyDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeStrategyDetailResponse) SetStatusCode(v int32) *DescribeStrategyDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStrategyDetailResponse) SetBody(v *DescribeStrategyDetailResponseBody) *DescribeStrategyDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeStrategyDetailResponse) Validate() error {
	return dara.Validate(s)
}

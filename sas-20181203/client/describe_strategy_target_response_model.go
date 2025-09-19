// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStrategyTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStrategyTargetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStrategyTargetResponseBody) *DescribeStrategyTargetResponse
	GetBody() *DescribeStrategyTargetResponseBody
}

type DescribeStrategyTargetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStrategyTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStrategyTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyTargetResponse) GoString() string {
	return s.String()
}

func (s *DescribeStrategyTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStrategyTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStrategyTargetResponse) GetBody() *DescribeStrategyTargetResponseBody {
	return s.Body
}

func (s *DescribeStrategyTargetResponse) SetHeaders(v map[string]*string) *DescribeStrategyTargetResponse {
	s.Headers = v
	return s
}

func (s *DescribeStrategyTargetResponse) SetStatusCode(v int32) *DescribeStrategyTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStrategyTargetResponse) SetBody(v *DescribeStrategyTargetResponseBody) *DescribeStrategyTargetResponse {
	s.Body = v
	return s
}

func (s *DescribeStrategyTargetResponse) Validate() error {
	return dara.Validate(s)
}

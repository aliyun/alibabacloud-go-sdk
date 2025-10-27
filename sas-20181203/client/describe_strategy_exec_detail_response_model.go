// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStrategyExecDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeStrategyExecDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeStrategyExecDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeStrategyExecDetailResponseBody) *DescribeStrategyExecDetailResponse
	GetBody() *DescribeStrategyExecDetailResponseBody
}

type DescribeStrategyExecDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeStrategyExecDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeStrategyExecDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeStrategyExecDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeStrategyExecDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeStrategyExecDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeStrategyExecDetailResponse) GetBody() *DescribeStrategyExecDetailResponseBody {
	return s.Body
}

func (s *DescribeStrategyExecDetailResponse) SetHeaders(v map[string]*string) *DescribeStrategyExecDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeStrategyExecDetailResponse) SetStatusCode(v int32) *DescribeStrategyExecDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStrategyExecDetailResponse) SetBody(v *DescribeStrategyExecDetailResponseBody) *DescribeStrategyExecDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeStrategyExecDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

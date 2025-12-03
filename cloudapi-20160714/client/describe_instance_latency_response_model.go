// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceLatencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceLatencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceLatencyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceLatencyResponseBody) *DescribeInstanceLatencyResponse
	GetBody() *DescribeInstanceLatencyResponseBody
}

type DescribeInstanceLatencyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceLatencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceLatencyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceLatencyResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceLatencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceLatencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceLatencyResponse) GetBody() *DescribeInstanceLatencyResponseBody {
	return s.Body
}

func (s *DescribeInstanceLatencyResponse) SetHeaders(v map[string]*string) *DescribeInstanceLatencyResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceLatencyResponse) SetStatusCode(v int32) *DescribeInstanceLatencyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceLatencyResponse) SetBody(v *DescribeInstanceLatencyResponseBody) *DescribeInstanceLatencyResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceLatencyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

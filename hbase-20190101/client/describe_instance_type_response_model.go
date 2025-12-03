// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceTypeResponseBody) *DescribeInstanceTypeResponse
	GetBody() *DescribeInstanceTypeResponseBody
}

type DescribeInstanceTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceTypeResponse) GetBody() *DescribeInstanceTypeResponseBody {
	return s.Body
}

func (s *DescribeInstanceTypeResponse) SetHeaders(v map[string]*string) *DescribeInstanceTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTypeResponse) SetStatusCode(v int32) *DescribeInstanceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceTypeResponse) SetBody(v *DescribeInstanceTypeResponseBody) *DescribeInstanceTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

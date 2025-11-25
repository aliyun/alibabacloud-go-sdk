// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSpecsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceSpecsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceSpecsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceSpecsResponseBody) *DescribeInstanceSpecsResponse
	GetBody() *DescribeInstanceSpecsResponseBody
}

type DescribeInstanceSpecsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSpecsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSpecsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSpecsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceSpecsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceSpecsResponse) GetBody() *DescribeInstanceSpecsResponseBody {
	return s.Body
}

func (s *DescribeInstanceSpecsResponse) SetHeaders(v map[string]*string) *DescribeInstanceSpecsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetStatusCode(v int32) *DescribeInstanceSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSpecsResponse) SetBody(v *DescribeInstanceSpecsResponseBody) *DescribeInstanceSpecsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceSpecsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

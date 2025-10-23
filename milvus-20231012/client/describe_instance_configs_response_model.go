// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceConfigsResponseBody) *DescribeInstanceConfigsResponse
	GetBody() *DescribeInstanceConfigsResponseBody
}

type DescribeInstanceConfigsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceConfigsResponse) GetBody() *DescribeInstanceConfigsResponseBody {
	return s.Body
}

func (s *DescribeInstanceConfigsResponse) SetHeaders(v map[string]*string) *DescribeInstanceConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceConfigsResponse) SetStatusCode(v int32) *DescribeInstanceConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceConfigsResponse) SetBody(v *DescribeInstanceConfigsResponseBody) *DescribeInstanceConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

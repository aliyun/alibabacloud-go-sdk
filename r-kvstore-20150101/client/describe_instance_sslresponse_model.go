// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSSLResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceSSLResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceSSLResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceSSLResponseBody) *DescribeInstanceSSLResponse
	GetBody() *DescribeInstanceSSLResponseBody
}

type DescribeInstanceSSLResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceSSLResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceSSLResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSSLResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceSSLResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceSSLResponse) GetBody() *DescribeInstanceSSLResponseBody {
	return s.Body
}

func (s *DescribeInstanceSSLResponse) SetHeaders(v map[string]*string) *DescribeInstanceSSLResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceSSLResponse) SetStatusCode(v int32) *DescribeInstanceSSLResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceSSLResponse) SetBody(v *DescribeInstanceSSLResponseBody) *DescribeInstanceSSLResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceSSLResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

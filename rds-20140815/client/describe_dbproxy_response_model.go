// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBProxyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBProxyResponseBody) *DescribeDBProxyResponse
	GetBody() *DescribeDBProxyResponseBody
}

type DescribeDBProxyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBProxyResponse) GetBody() *DescribeDBProxyResponseBody {
	return s.Body
}

func (s *DescribeDBProxyResponse) SetHeaders(v map[string]*string) *DescribeDBProxyResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBProxyResponse) SetStatusCode(v int32) *DescribeDBProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBProxyResponse) SetBody(v *DescribeDBProxyResponseBody) *DescribeDBProxyResponse {
	s.Body = v
	return s
}

func (s *DescribeDBProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

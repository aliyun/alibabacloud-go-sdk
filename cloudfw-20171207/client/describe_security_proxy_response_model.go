// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityProxyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityProxyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityProxyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityProxyResponseBody) *DescribeSecurityProxyResponse
	GetBody() *DescribeSecurityProxyResponseBody
}

type DescribeSecurityProxyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityProxyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityProxyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityProxyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityProxyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityProxyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityProxyResponse) GetBody() *DescribeSecurityProxyResponseBody {
	return s.Body
}

func (s *DescribeSecurityProxyResponse) SetHeaders(v map[string]*string) *DescribeSecurityProxyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityProxyResponse) SetStatusCode(v int32) *DescribeSecurityProxyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityProxyResponse) SetBody(v *DescribeSecurityProxyResponseBody) *DescribeSecurityProxyResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityProxyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

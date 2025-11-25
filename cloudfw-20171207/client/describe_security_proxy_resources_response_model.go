// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityProxyResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityProxyResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityProxyResourcesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityProxyResourcesResponseBody) *DescribeSecurityProxyResourcesResponse
	GetBody() *DescribeSecurityProxyResourcesResponseBody
}

type DescribeSecurityProxyResourcesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityProxyResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityProxyResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityProxyResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityProxyResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityProxyResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityProxyResourcesResponse) GetBody() *DescribeSecurityProxyResourcesResponseBody {
	return s.Body
}

func (s *DescribeSecurityProxyResourcesResponse) SetHeaders(v map[string]*string) *DescribeSecurityProxyResourcesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityProxyResourcesResponse) SetStatusCode(v int32) *DescribeSecurityProxyResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityProxyResourcesResponse) SetBody(v *DescribeSecurityProxyResourcesResponseBody) *DescribeSecurityProxyResourcesResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityProxyResourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

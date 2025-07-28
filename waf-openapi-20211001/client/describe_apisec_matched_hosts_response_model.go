// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecMatchedHostsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecMatchedHostsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecMatchedHostsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecMatchedHostsResponseBody) *DescribeApisecMatchedHostsResponse
	GetBody() *DescribeApisecMatchedHostsResponseBody
}

type DescribeApisecMatchedHostsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecMatchedHostsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecMatchedHostsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecMatchedHostsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecMatchedHostsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecMatchedHostsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecMatchedHostsResponse) GetBody() *DescribeApisecMatchedHostsResponseBody {
	return s.Body
}

func (s *DescribeApisecMatchedHostsResponse) SetHeaders(v map[string]*string) *DescribeApisecMatchedHostsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecMatchedHostsResponse) SetStatusCode(v int32) *DescribeApisecMatchedHostsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecMatchedHostsResponse) SetBody(v *DescribeApisecMatchedHostsResponseBody) *DescribeApisecMatchedHostsResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecMatchedHostsResponse) Validate() error {
	return dara.Validate(s)
}

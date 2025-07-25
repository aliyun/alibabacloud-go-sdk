// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceDomainsResponseBody) *DescribeInstanceDomainsResponse
	GetBody() *DescribeInstanceDomainsResponseBody
}

type DescribeInstanceDomainsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceDomainsResponse) GetBody() *DescribeInstanceDomainsResponseBody {
	return s.Body
}

func (s *DescribeInstanceDomainsResponse) SetHeaders(v map[string]*string) *DescribeInstanceDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceDomainsResponse) SetStatusCode(v int32) *DescribeInstanceDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceDomainsResponse) SetBody(v *DescribeInstanceDomainsResponseBody) *DescribeInstanceDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceDomainsResponse) Validate() error {
	return dara.Validate(s)
}

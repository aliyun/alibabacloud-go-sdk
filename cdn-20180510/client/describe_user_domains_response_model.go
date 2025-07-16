// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserDomainsResponseBody) *DescribeUserDomainsResponse
	GetBody() *DescribeUserDomainsResponseBody
}

type DescribeUserDomainsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserDomainsResponse) GetBody() *DescribeUserDomainsResponseBody {
	return s.Body
}

func (s *DescribeUserDomainsResponse) SetHeaders(v map[string]*string) *DescribeUserDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserDomainsResponse) SetStatusCode(v int32) *DescribeUserDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserDomainsResponse) SetBody(v *DescribeUserDomainsResponseBody) *DescribeUserDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeUserDomainsResponse) Validate() error {
	return dara.Validate(s)
}

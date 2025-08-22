// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRDDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRDDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRDDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRDDomainsResponseBody) *DescribeRDDomainsResponse
	GetBody() *DescribeRDDomainsResponseBody
}

type DescribeRDDomainsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRDDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRDDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRDDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRDDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRDDomainsResponse) GetBody() *DescribeRDDomainsResponseBody {
	return s.Body
}

func (s *DescribeRDDomainsResponse) SetHeaders(v map[string]*string) *DescribeRDDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRDDomainsResponse) SetStatusCode(v int32) *DescribeRDDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRDDomainsResponse) SetBody(v *DescribeRDDomainsResponseBody) *DescribeRDDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeRDDomainsResponse) Validate() error {
	return dara.Validate(s)
}

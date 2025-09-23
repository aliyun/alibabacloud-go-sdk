// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimpleDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSimpleDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSimpleDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSimpleDomainsResponseBody) *DescribeSimpleDomainsResponse
	GetBody() *DescribeSimpleDomainsResponseBody
}

type DescribeSimpleDomainsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSimpleDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSimpleDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimpleDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimpleDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSimpleDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSimpleDomainsResponse) GetBody() *DescribeSimpleDomainsResponseBody {
	return s.Body
}

func (s *DescribeSimpleDomainsResponse) SetHeaders(v map[string]*string) *DescribeSimpleDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSimpleDomainsResponse) SetStatusCode(v int32) *DescribeSimpleDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSimpleDomainsResponse) SetBody(v *DescribeSimpleDomainsResponseBody) *DescribeSimpleDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeSimpleDomainsResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainsResponseBody) *DescribeDomainsResponse
	GetBody() *DescribeDomainsResponseBody
}

type DescribeDomainsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainsResponse) GetBody() *DescribeDomainsResponseBody {
	return s.Body
}

func (s *DescribeDomainsResponse) SetHeaders(v map[string]*string) *DescribeDomainsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainsResponse) SetStatusCode(v int32) *DescribeDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainsResponse) SetBody(v *DescribeDomainsResponseBody) *DescribeDomainsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainConfigsResponseBody) *DescribeVsDomainConfigsResponse
	GetBody() *DescribeVsDomainConfigsResponseBody
}

type DescribeVsDomainConfigsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainConfigsResponse) GetBody() *DescribeVsDomainConfigsResponseBody {
	return s.Body
}

func (s *DescribeVsDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeVsDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainConfigsResponse) SetStatusCode(v int32) *DescribeVsDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainConfigsResponse) SetBody(v *DescribeVsDomainConfigsResponseBody) *DescribeVsDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

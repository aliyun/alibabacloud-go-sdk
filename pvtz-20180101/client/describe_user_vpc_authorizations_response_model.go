// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserVpcAuthorizationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserVpcAuthorizationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserVpcAuthorizationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserVpcAuthorizationsResponseBody) *DescribeUserVpcAuthorizationsResponse
	GetBody() *DescribeUserVpcAuthorizationsResponseBody
}

type DescribeUserVpcAuthorizationsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserVpcAuthorizationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserVpcAuthorizationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserVpcAuthorizationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserVpcAuthorizationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserVpcAuthorizationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserVpcAuthorizationsResponse) GetBody() *DescribeUserVpcAuthorizationsResponseBody {
	return s.Body
}

func (s *DescribeUserVpcAuthorizationsResponse) SetHeaders(v map[string]*string) *DescribeUserVpcAuthorizationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponse) SetStatusCode(v int32) *DescribeUserVpcAuthorizationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponse) SetBody(v *DescribeUserVpcAuthorizationsResponseBody) *DescribeUserVpcAuthorizationsResponse {
	s.Body = v
	return s
}

func (s *DescribeUserVpcAuthorizationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

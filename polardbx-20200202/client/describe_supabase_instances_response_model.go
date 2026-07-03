// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupabaseInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupabaseInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupabaseInstancesResponseBody) *DescribeSupabaseInstancesResponse
	GetBody() *DescribeSupabaseInstancesResponseBody
}

type DescribeSupabaseInstancesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupabaseInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupabaseInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupabaseInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupabaseInstancesResponse) GetBody() *DescribeSupabaseInstancesResponseBody {
	return s.Body
}

func (s *DescribeSupabaseInstancesResponse) SetHeaders(v map[string]*string) *DescribeSupabaseInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupabaseInstancesResponse) SetStatusCode(v int32) *DescribeSupabaseInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupabaseInstancesResponse) SetBody(v *DescribeSupabaseInstancesResponseBody) *DescribeSupabaseInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeSupabaseInstancesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

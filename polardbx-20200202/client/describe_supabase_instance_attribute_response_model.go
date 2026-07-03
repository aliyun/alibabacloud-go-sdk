// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseInstanceAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupabaseInstanceAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupabaseInstanceAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupabaseInstanceAttributeResponseBody) *DescribeSupabaseInstanceAttributeResponse
	GetBody() *DescribeSupabaseInstanceAttributeResponseBody
}

type DescribeSupabaseInstanceAttributeResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupabaseInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupabaseInstanceAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstanceAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupabaseInstanceAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupabaseInstanceAttributeResponse) GetBody() *DescribeSupabaseInstanceAttributeResponseBody {
	return s.Body
}

func (s *DescribeSupabaseInstanceAttributeResponse) SetHeaders(v map[string]*string) *DescribeSupabaseInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponse) SetStatusCode(v int32) *DescribeSupabaseInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponse) SetBody(v *DescribeSupabaseInstanceAttributeResponseBody) *DescribeSupabaseInstanceAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeSupabaseInstanceAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

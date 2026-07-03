// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseApiKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupabaseApiKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupabaseApiKeyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupabaseApiKeyResponseBody) *DescribeSupabaseApiKeyResponse
	GetBody() *DescribeSupabaseApiKeyResponseBody
}

type DescribeSupabaseApiKeyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupabaseApiKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupabaseApiKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseApiKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseApiKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupabaseApiKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupabaseApiKeyResponse) GetBody() *DescribeSupabaseApiKeyResponseBody {
	return s.Body
}

func (s *DescribeSupabaseApiKeyResponse) SetHeaders(v map[string]*string) *DescribeSupabaseApiKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupabaseApiKeyResponse) SetStatusCode(v int32) *DescribeSupabaseApiKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupabaseApiKeyResponse) SetBody(v *DescribeSupabaseApiKeyResponseBody) *DescribeSupabaseApiKeyResponse {
	s.Body = v
	return s
}

func (s *DescribeSupabaseApiKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseIpWhitelistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSupabaseIpWhitelistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSupabaseIpWhitelistResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSupabaseIpWhitelistResponseBody) *DescribeSupabaseIpWhitelistResponse
	GetBody() *DescribeSupabaseIpWhitelistResponseBody
}

type DescribeSupabaseIpWhitelistResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSupabaseIpWhitelistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSupabaseIpWhitelistResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseIpWhitelistResponse) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseIpWhitelistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSupabaseIpWhitelistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSupabaseIpWhitelistResponse) GetBody() *DescribeSupabaseIpWhitelistResponseBody {
	return s.Body
}

func (s *DescribeSupabaseIpWhitelistResponse) SetHeaders(v map[string]*string) *DescribeSupabaseIpWhitelistResponse {
	s.Headers = v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponse) SetStatusCode(v int32) *DescribeSupabaseIpWhitelistResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponse) SetBody(v *DescribeSupabaseIpWhitelistResponseBody) *DescribeSupabaseIpWhitelistResponse {
	s.Body = v
	return s
}

func (s *DescribeSupabaseIpWhitelistResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecurityGroupConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecurityGroupConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecurityGroupConfigurationResponseBody) *DescribeSecurityGroupConfigurationResponse
	GetBody() *DescribeSecurityGroupConfigurationResponseBody
}

type DescribeSecurityGroupConfigurationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecurityGroupConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecurityGroupConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecurityGroupConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecurityGroupConfigurationResponse) GetBody() *DescribeSecurityGroupConfigurationResponseBody {
	return s.Body
}

func (s *DescribeSecurityGroupConfigurationResponse) SetHeaders(v map[string]*string) *DescribeSecurityGroupConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponse) SetStatusCode(v int32) *DescribeSecurityGroupConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponse) SetBody(v *DescribeSecurityGroupConfigurationResponseBody) *DescribeSecurityGroupConfigurationResponse {
	s.Body = v
	return s
}

func (s *DescribeSecurityGroupConfigurationResponse) Validate() error {
	return dara.Validate(s)
}

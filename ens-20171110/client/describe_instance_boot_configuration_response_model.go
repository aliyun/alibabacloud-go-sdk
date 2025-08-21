// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceBootConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceBootConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceBootConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceBootConfigurationResponseBody) *DescribeInstanceBootConfigurationResponse
	GetBody() *DescribeInstanceBootConfigurationResponseBody
}

type DescribeInstanceBootConfigurationResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceBootConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceBootConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBootConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBootConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceBootConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceBootConfigurationResponse) GetBody() *DescribeInstanceBootConfigurationResponseBody {
	return s.Body
}

func (s *DescribeInstanceBootConfigurationResponse) SetHeaders(v map[string]*string) *DescribeInstanceBootConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceBootConfigurationResponse) SetStatusCode(v int32) *DescribeInstanceBootConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceBootConfigurationResponse) SetBody(v *DescribeInstanceBootConfigurationResponseBody) *DescribeInstanceBootConfigurationResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceBootConfigurationResponse) Validate() error {
	return dara.Validate(s)
}

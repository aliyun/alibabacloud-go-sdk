// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingInstanceConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRenderingInstanceConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRenderingInstanceConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRenderingInstanceConfigurationResponseBody) *DescribeRenderingInstanceConfigurationResponse
	GetBody() *DescribeRenderingInstanceConfigurationResponseBody
}

type DescribeRenderingInstanceConfigurationResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRenderingInstanceConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRenderingInstanceConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRenderingInstanceConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRenderingInstanceConfigurationResponse) GetBody() *DescribeRenderingInstanceConfigurationResponseBody {
	return s.Body
}

func (s *DescribeRenderingInstanceConfigurationResponse) SetHeaders(v map[string]*string) *DescribeRenderingInstanceConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeRenderingInstanceConfigurationResponse) SetStatusCode(v int32) *DescribeRenderingInstanceConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRenderingInstanceConfigurationResponse) SetBody(v *DescribeRenderingInstanceConfigurationResponseBody) *DescribeRenderingInstanceConfigurationResponse {
	s.Body = v
	return s
}

func (s *DescribeRenderingInstanceConfigurationResponse) Validate() error {
	return dara.Validate(s)
}

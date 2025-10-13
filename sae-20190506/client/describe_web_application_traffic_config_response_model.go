// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebApplicationTrafficConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebApplicationTrafficConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebApplicationTrafficConfigResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationTrafficConfigBody) *DescribeWebApplicationTrafficConfigResponse
	GetBody() *WebApplicationTrafficConfigBody
}

type DescribeWebApplicationTrafficConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationTrafficConfigBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebApplicationTrafficConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebApplicationTrafficConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebApplicationTrafficConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebApplicationTrafficConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebApplicationTrafficConfigResponse) GetBody() *WebApplicationTrafficConfigBody {
	return s.Body
}

func (s *DescribeWebApplicationTrafficConfigResponse) SetHeaders(v map[string]*string) *DescribeWebApplicationTrafficConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebApplicationTrafficConfigResponse) SetStatusCode(v int32) *DescribeWebApplicationTrafficConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebApplicationTrafficConfigResponse) SetBody(v *WebApplicationTrafficConfigBody) *DescribeWebApplicationTrafficConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeWebApplicationTrafficConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

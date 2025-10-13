// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebApplicationScalingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebApplicationScalingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebApplicationScalingConfigResponse
	GetStatusCode() *int32
	SetBody(v *WebApplicationScalingConfigBody) *DescribeWebApplicationScalingConfigResponse
	GetBody() *WebApplicationScalingConfigBody
}

type DescribeWebApplicationScalingConfigResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebApplicationScalingConfigBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebApplicationScalingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebApplicationScalingConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeWebApplicationScalingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebApplicationScalingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebApplicationScalingConfigResponse) GetBody() *WebApplicationScalingConfigBody {
	return s.Body
}

func (s *DescribeWebApplicationScalingConfigResponse) SetHeaders(v map[string]*string) *DescribeWebApplicationScalingConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeWebApplicationScalingConfigResponse) SetStatusCode(v int32) *DescribeWebApplicationScalingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebApplicationScalingConfigResponse) SetBody(v *WebApplicationScalingConfigBody) *DescribeWebApplicationScalingConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeWebApplicationScalingConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

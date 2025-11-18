// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceEndpointsResponseBody) *DescribeInstanceEndpointsResponse
	GetBody() *DescribeInstanceEndpointsResponseBody
}

type DescribeInstanceEndpointsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceEndpointsResponse) GetBody() *DescribeInstanceEndpointsResponseBody {
	return s.Body
}

func (s *DescribeInstanceEndpointsResponse) SetHeaders(v map[string]*string) *DescribeInstanceEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceEndpointsResponse) SetStatusCode(v int32) *DescribeInstanceEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceEndpointsResponse) SetBody(v *DescribeInstanceEndpointsResponseBody) *DescribeInstanceEndpointsResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

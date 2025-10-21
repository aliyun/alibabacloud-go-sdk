// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEndpointsResponseBody) *DescribeEndpointsResponse
	GetBody() *DescribeEndpointsResponseBody
}

type DescribeEndpointsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEndpointsResponse) GetBody() *DescribeEndpointsResponseBody {
	return s.Body
}

func (s *DescribeEndpointsResponse) SetHeaders(v map[string]*string) *DescribeEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointsResponse) SetStatusCode(v int32) *DescribeEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndpointsResponse) SetBody(v *DescribeEndpointsResponseBody) *DescribeEndpointsResponse {
	s.Body = v
	return s
}

func (s *DescribeEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

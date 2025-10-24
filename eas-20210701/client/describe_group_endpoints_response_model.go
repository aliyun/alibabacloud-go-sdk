// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupEndpointsResponseBody) *DescribeGroupEndpointsResponse
	GetBody() *DescribeGroupEndpointsResponseBody
}

type DescribeGroupEndpointsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupEndpointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupEndpointsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupEndpointsResponse) GetBody() *DescribeGroupEndpointsResponseBody {
	return s.Body
}

func (s *DescribeGroupEndpointsResponse) SetHeaders(v map[string]*string) *DescribeGroupEndpointsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupEndpointsResponse) SetStatusCode(v int32) *DescribeGroupEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupEndpointsResponse) SetBody(v *DescribeGroupEndpointsResponseBody) *DescribeGroupEndpointsResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

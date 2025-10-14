// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEndpointListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomEndpointListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomEndpointListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomEndpointListResponseBody) *DescribeCustomEndpointListResponse
	GetBody() *DescribeCustomEndpointListResponseBody
}

type DescribeCustomEndpointListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomEndpointListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomEndpointListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEndpointListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomEndpointListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomEndpointListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomEndpointListResponse) GetBody() *DescribeCustomEndpointListResponseBody {
	return s.Body
}

func (s *DescribeCustomEndpointListResponse) SetHeaders(v map[string]*string) *DescribeCustomEndpointListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomEndpointListResponse) SetStatusCode(v int32) *DescribeCustomEndpointListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomEndpointListResponse) SetBody(v *DescribeCustomEndpointListResponseBody) *DescribeCustomEndpointListResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomEndpointListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

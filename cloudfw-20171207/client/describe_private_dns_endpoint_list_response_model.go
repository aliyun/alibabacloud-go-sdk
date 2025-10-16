// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsEndpointListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrivateDnsEndpointListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrivateDnsEndpointListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrivateDnsEndpointListResponseBody) *DescribePrivateDnsEndpointListResponse
	GetBody() *DescribePrivateDnsEndpointListResponseBody
}

type DescribePrivateDnsEndpointListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrivateDnsEndpointListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrivateDnsEndpointListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsEndpointListResponse) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsEndpointListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrivateDnsEndpointListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrivateDnsEndpointListResponse) GetBody() *DescribePrivateDnsEndpointListResponseBody {
	return s.Body
}

func (s *DescribePrivateDnsEndpointListResponse) SetHeaders(v map[string]*string) *DescribePrivateDnsEndpointListResponse {
	s.Headers = v
	return s
}

func (s *DescribePrivateDnsEndpointListResponse) SetStatusCode(v int32) *DescribePrivateDnsEndpointListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrivateDnsEndpointListResponse) SetBody(v *DescribePrivateDnsEndpointListResponseBody) *DescribePrivateDnsEndpointListResponse {
	s.Body = v
	return s
}

func (s *DescribePrivateDnsEndpointListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

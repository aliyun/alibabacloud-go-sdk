// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrivateDnsEndpointDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePrivateDnsEndpointDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePrivateDnsEndpointDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribePrivateDnsEndpointDetailResponseBody) *DescribePrivateDnsEndpointDetailResponse
	GetBody() *DescribePrivateDnsEndpointDetailResponseBody
}

type DescribePrivateDnsEndpointDetailResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePrivateDnsEndpointDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePrivateDnsEndpointDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePrivateDnsEndpointDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribePrivateDnsEndpointDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePrivateDnsEndpointDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePrivateDnsEndpointDetailResponse) GetBody() *DescribePrivateDnsEndpointDetailResponseBody {
	return s.Body
}

func (s *DescribePrivateDnsEndpointDetailResponse) SetHeaders(v map[string]*string) *DescribePrivateDnsEndpointDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponse) SetStatusCode(v int32) *DescribePrivateDnsEndpointDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponse) SetBody(v *DescribePrivateDnsEndpointDetailResponseBody) *DescribePrivateDnsEndpointDetailResponse {
	s.Body = v
	return s
}

func (s *DescribePrivateDnsEndpointDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResolverAvailableZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeResolverAvailableZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeResolverAvailableZonesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeResolverAvailableZonesResponseBody) *DescribeResolverAvailableZonesResponse
	GetBody() *DescribeResolverAvailableZonesResponseBody
}

type DescribeResolverAvailableZonesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeResolverAvailableZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeResolverAvailableZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeResolverAvailableZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeResolverAvailableZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeResolverAvailableZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeResolverAvailableZonesResponse) GetBody() *DescribeResolverAvailableZonesResponseBody {
	return s.Body
}

func (s *DescribeResolverAvailableZonesResponse) SetHeaders(v map[string]*string) *DescribeResolverAvailableZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeResolverAvailableZonesResponse) SetStatusCode(v int32) *DescribeResolverAvailableZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResolverAvailableZonesResponse) SetBody(v *DescribeResolverAvailableZonesResponseBody) *DescribeResolverAvailableZonesResponse {
	s.Body = v
	return s
}

func (s *DescribeResolverAvailableZonesResponse) Validate() error {
	return dara.Validate(s)
}

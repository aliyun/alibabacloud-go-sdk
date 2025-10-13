// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGreyTagRouteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGreyTagRouteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGreyTagRouteResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGreyTagRouteResponseBody) *DescribeGreyTagRouteResponse
	GetBody() *DescribeGreyTagRouteResponseBody
}

type DescribeGreyTagRouteResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGreyTagRouteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGreyTagRouteResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGreyTagRouteResponse) GoString() string {
	return s.String()
}

func (s *DescribeGreyTagRouteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGreyTagRouteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGreyTagRouteResponse) GetBody() *DescribeGreyTagRouteResponseBody {
	return s.Body
}

func (s *DescribeGreyTagRouteResponse) SetHeaders(v map[string]*string) *DescribeGreyTagRouteResponse {
	s.Headers = v
	return s
}

func (s *DescribeGreyTagRouteResponse) SetStatusCode(v int32) *DescribeGreyTagRouteResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGreyTagRouteResponse) SetBody(v *DescribeGreyTagRouteResponseBody) *DescribeGreyTagRouteResponse {
	s.Body = v
	return s
}

func (s *DescribeGreyTagRouteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

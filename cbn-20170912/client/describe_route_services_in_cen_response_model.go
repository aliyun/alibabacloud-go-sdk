// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteServicesInCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRouteServicesInCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRouteServicesInCenResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRouteServicesInCenResponseBody) *DescribeRouteServicesInCenResponse
	GetBody() *DescribeRouteServicesInCenResponseBody
}

type DescribeRouteServicesInCenResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRouteServicesInCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRouteServicesInCenResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteServicesInCenResponse) GoString() string {
	return s.String()
}

func (s *DescribeRouteServicesInCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRouteServicesInCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRouteServicesInCenResponse) GetBody() *DescribeRouteServicesInCenResponseBody {
	return s.Body
}

func (s *DescribeRouteServicesInCenResponse) SetHeaders(v map[string]*string) *DescribeRouteServicesInCenResponse {
	s.Headers = v
	return s
}

func (s *DescribeRouteServicesInCenResponse) SetStatusCode(v int32) *DescribeRouteServicesInCenResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRouteServicesInCenResponse) SetBody(v *DescribeRouteServicesInCenResponseBody) *DescribeRouteServicesInCenResponse {
	s.Body = v
	return s
}

func (s *DescribeRouteServicesInCenResponse) Validate() error {
	return dara.Validate(s)
}

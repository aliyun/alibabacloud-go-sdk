// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterResourcesListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTransitRouterResourcesListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTransitRouterResourcesListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTransitRouterResourcesListResponseBody) *DescribeTransitRouterResourcesListResponse
	GetBody() *DescribeTransitRouterResourcesListResponseBody
}

type DescribeTransitRouterResourcesListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransitRouterResourcesListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTransitRouterResourcesListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterResourcesListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterResourcesListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTransitRouterResourcesListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTransitRouterResourcesListResponse) GetBody() *DescribeTransitRouterResourcesListResponseBody {
	return s.Body
}

func (s *DescribeTransitRouterResourcesListResponse) SetHeaders(v map[string]*string) *DescribeTransitRouterResourcesListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransitRouterResourcesListResponse) SetStatusCode(v int32) *DescribeTransitRouterResourcesListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransitRouterResourcesListResponse) SetBody(v *DescribeTransitRouterResourcesListResponseBody) *DescribeTransitRouterResourcesListResponse {
	s.Body = v
	return s
}

func (s *DescribeTransitRouterResourcesListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

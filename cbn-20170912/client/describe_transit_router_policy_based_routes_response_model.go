// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterPolicyBasedRoutesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTransitRouterPolicyBasedRoutesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTransitRouterPolicyBasedRoutesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTransitRouterPolicyBasedRoutesResponseBody) *DescribeTransitRouterPolicyBasedRoutesResponse
	GetBody() *DescribeTransitRouterPolicyBasedRoutesResponseBody
}

type DescribeTransitRouterPolicyBasedRoutesResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransitRouterPolicyBasedRoutesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTransitRouterPolicyBasedRoutesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyBasedRoutesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponse) GetBody() *DescribeTransitRouterPolicyBasedRoutesResponseBody {
	return s.Body
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponse) SetHeaders(v map[string]*string) *DescribeTransitRouterPolicyBasedRoutesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponse) SetStatusCode(v int32) *DescribeTransitRouterPolicyBasedRoutesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponse) SetBody(v *DescribeTransitRouterPolicyBasedRoutesResponseBody) *DescribeTransitRouterPolicyBasedRoutesResponse {
	s.Body = v
	return s
}

func (s *DescribeTransitRouterPolicyBasedRoutesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

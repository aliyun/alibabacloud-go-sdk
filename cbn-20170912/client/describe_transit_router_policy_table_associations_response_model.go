// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterPolicyTableAssociationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTransitRouterPolicyTableAssociationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTransitRouterPolicyTableAssociationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTransitRouterPolicyTableAssociationsResponseBody) *DescribeTransitRouterPolicyTableAssociationsResponse
	GetBody() *DescribeTransitRouterPolicyTableAssociationsResponseBody
}

type DescribeTransitRouterPolicyTableAssociationsResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransitRouterPolicyTableAssociationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTransitRouterPolicyTableAssociationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTableAssociationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponse) GetBody() *DescribeTransitRouterPolicyTableAssociationsResponseBody {
	return s.Body
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponse) SetHeaders(v map[string]*string) *DescribeTransitRouterPolicyTableAssociationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponse) SetStatusCode(v int32) *DescribeTransitRouterPolicyTableAssociationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponse) SetBody(v *DescribeTransitRouterPolicyTableAssociationsResponseBody) *DescribeTransitRouterPolicyTableAssociationsResponse {
	s.Body = v
	return s
}

func (s *DescribeTransitRouterPolicyTableAssociationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

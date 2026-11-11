// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransitRouterPolicyTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTransitRouterPolicyTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTransitRouterPolicyTablesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTransitRouterPolicyTablesResponseBody) *DescribeTransitRouterPolicyTablesResponse
	GetBody() *DescribeTransitRouterPolicyTablesResponseBody
}

type DescribeTransitRouterPolicyTablesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTransitRouterPolicyTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTransitRouterPolicyTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransitRouterPolicyTablesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTransitRouterPolicyTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTransitRouterPolicyTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTransitRouterPolicyTablesResponse) GetBody() *DescribeTransitRouterPolicyTablesResponseBody {
	return s.Body
}

func (s *DescribeTransitRouterPolicyTablesResponse) SetHeaders(v map[string]*string) *DescribeTransitRouterPolicyTablesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponse) SetStatusCode(v int32) *DescribeTransitRouterPolicyTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponse) SetBody(v *DescribeTransitRouterPolicyTablesResponseBody) *DescribeTransitRouterPolicyTablesResponse {
	s.Body = v
	return s
}

func (s *DescribeTransitRouterPolicyTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

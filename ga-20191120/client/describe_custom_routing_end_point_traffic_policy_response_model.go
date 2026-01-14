// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomRoutingEndPointTrafficPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomRoutingEndPointTrafficPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomRoutingEndPointTrafficPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) *DescribeCustomRoutingEndPointTrafficPolicyResponse
	GetBody() *DescribeCustomRoutingEndPointTrafficPolicyResponseBody
}

type DescribeCustomRoutingEndPointTrafficPolicyResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomRoutingEndPointTrafficPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomRoutingEndPointTrafficPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponse) GetBody() *DescribeCustomRoutingEndPointTrafficPolicyResponseBody {
	return s.Body
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponse) SetHeaders(v map[string]*string) *DescribeCustomRoutingEndPointTrafficPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponse) SetStatusCode(v int32) *DescribeCustomRoutingEndPointTrafficPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponse) SetBody(v *DescribeCustomRoutingEndPointTrafficPolicyResponseBody) *DescribeCustomRoutingEndPointTrafficPolicyResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomRoutingEndPointTrafficPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

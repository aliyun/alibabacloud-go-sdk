// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLoadBalancerOriginStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLoadBalancerOriginStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLoadBalancerOriginStatusResponse
	GetStatusCode() *int32
	SetBody(v *ListLoadBalancerOriginStatusResponseBody) *ListLoadBalancerOriginStatusResponse
	GetBody() *ListLoadBalancerOriginStatusResponseBody
}

type ListLoadBalancerOriginStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLoadBalancerOriginStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLoadBalancerOriginStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLoadBalancerOriginStatusResponse) GoString() string {
	return s.String()
}

func (s *ListLoadBalancerOriginStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLoadBalancerOriginStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLoadBalancerOriginStatusResponse) GetBody() *ListLoadBalancerOriginStatusResponseBody {
	return s.Body
}

func (s *ListLoadBalancerOriginStatusResponse) SetHeaders(v map[string]*string) *ListLoadBalancerOriginStatusResponse {
	s.Headers = v
	return s
}

func (s *ListLoadBalancerOriginStatusResponse) SetStatusCode(v int32) *ListLoadBalancerOriginStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLoadBalancerOriginStatusResponse) SetBody(v *ListLoadBalancerOriginStatusResponseBody) *ListLoadBalancerOriginStatusResponse {
	s.Body = v
	return s
}

func (s *ListLoadBalancerOriginStatusResponse) Validate() error {
	return dara.Validate(s)
}

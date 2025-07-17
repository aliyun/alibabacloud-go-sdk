// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLoadBalancersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachLoadBalancersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachLoadBalancersResponse
	GetStatusCode() *int32
	SetBody(v *AttachLoadBalancersResponseBody) *AttachLoadBalancersResponse
	GetBody() *AttachLoadBalancersResponseBody
}

type AttachLoadBalancersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachLoadBalancersResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachLoadBalancersResponse) GoString() string {
	return s.String()
}

func (s *AttachLoadBalancersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachLoadBalancersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachLoadBalancersResponse) GetBody() *AttachLoadBalancersResponseBody {
	return s.Body
}

func (s *AttachLoadBalancersResponse) SetHeaders(v map[string]*string) *AttachLoadBalancersResponse {
	s.Headers = v
	return s
}

func (s *AttachLoadBalancersResponse) SetStatusCode(v int32) *AttachLoadBalancersResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachLoadBalancersResponse) SetBody(v *AttachLoadBalancersResponseBody) *AttachLoadBalancersResponse {
	s.Body = v
	return s
}

func (s *AttachLoadBalancersResponse) Validate() error {
	return dara.Validate(s)
}

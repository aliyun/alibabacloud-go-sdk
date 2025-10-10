// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerEditionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLoadBalancerEditionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLoadBalancerEditionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLoadBalancerEditionResponseBody) *UpdateLoadBalancerEditionResponse
	GetBody() *UpdateLoadBalancerEditionResponseBody
}

type UpdateLoadBalancerEditionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLoadBalancerEditionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLoadBalancerEditionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerEditionResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerEditionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLoadBalancerEditionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLoadBalancerEditionResponse) GetBody() *UpdateLoadBalancerEditionResponseBody {
	return s.Body
}

func (s *UpdateLoadBalancerEditionResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerEditionResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerEditionResponse) SetStatusCode(v int32) *UpdateLoadBalancerEditionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerEditionResponse) SetBody(v *UpdateLoadBalancerEditionResponseBody) *UpdateLoadBalancerEditionResponse {
	s.Body = v
	return s
}

func (s *UpdateLoadBalancerEditionResponse) Validate() error {
	return dara.Validate(s)
}

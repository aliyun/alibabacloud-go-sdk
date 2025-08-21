// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLoadBalancerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLoadBalancerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLoadBalancerAttributeResponseBody) *ModifyLoadBalancerAttributeResponse
	GetBody() *ModifyLoadBalancerAttributeResponseBody
}

type ModifyLoadBalancerAttributeResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLoadBalancerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLoadBalancerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLoadBalancerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLoadBalancerAttributeResponse) GetBody() *ModifyLoadBalancerAttributeResponseBody {
	return s.Body
}

func (s *ModifyLoadBalancerAttributeResponse) SetHeaders(v map[string]*string) *ModifyLoadBalancerAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoadBalancerAttributeResponse) SetStatusCode(v int32) *ModifyLoadBalancerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLoadBalancerAttributeResponse) SetBody(v *ModifyLoadBalancerAttributeResponseBody) *ModifyLoadBalancerAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyLoadBalancerAttributeResponse) Validate() error {
	return dara.Validate(s)
}

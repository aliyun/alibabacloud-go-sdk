// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerUDPListenerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLoadBalancerUDPListenerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLoadBalancerUDPListenerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *SetLoadBalancerUDPListenerAttributeResponseBody) *SetLoadBalancerUDPListenerAttributeResponse
	GetBody() *SetLoadBalancerUDPListenerAttributeResponseBody
}

type SetLoadBalancerUDPListenerAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoadBalancerUDPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoadBalancerUDPListenerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerUDPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerUDPListenerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLoadBalancerUDPListenerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLoadBalancerUDPListenerAttributeResponse) GetBody() *SetLoadBalancerUDPListenerAttributeResponseBody {
	return s.Body
}

func (s *SetLoadBalancerUDPListenerAttributeResponse) SetHeaders(v map[string]*string) *SetLoadBalancerUDPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeResponse) SetStatusCode(v int32) *SetLoadBalancerUDPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeResponse) SetBody(v *SetLoadBalancerUDPListenerAttributeResponseBody) *SetLoadBalancerUDPListenerAttributeResponse {
	s.Body = v
	return s
}

func (s *SetLoadBalancerUDPListenerAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

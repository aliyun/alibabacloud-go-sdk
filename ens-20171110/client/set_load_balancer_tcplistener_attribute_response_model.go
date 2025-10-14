// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerTCPListenerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLoadBalancerTCPListenerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLoadBalancerTCPListenerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *SetLoadBalancerTCPListenerAttributeResponseBody) *SetLoadBalancerTCPListenerAttributeResponse
	GetBody() *SetLoadBalancerTCPListenerAttributeResponseBody
}

type SetLoadBalancerTCPListenerAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoadBalancerTCPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoadBalancerTCPListenerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerTCPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerTCPListenerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLoadBalancerTCPListenerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLoadBalancerTCPListenerAttributeResponse) GetBody() *SetLoadBalancerTCPListenerAttributeResponseBody {
	return s.Body
}

func (s *SetLoadBalancerTCPListenerAttributeResponse) SetHeaders(v map[string]*string) *SetLoadBalancerTCPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeResponse) SetStatusCode(v int32) *SetLoadBalancerTCPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeResponse) SetBody(v *SetLoadBalancerTCPListenerAttributeResponseBody) *SetLoadBalancerTCPListenerAttributeResponse {
	s.Body = v
	return s
}

func (s *SetLoadBalancerTCPListenerAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

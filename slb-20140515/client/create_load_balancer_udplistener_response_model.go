// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerUDPListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLoadBalancerUDPListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLoadBalancerUDPListenerResponse
	GetStatusCode() *int32
	SetBody(v *CreateLoadBalancerUDPListenerResponseBody) *CreateLoadBalancerUDPListenerResponse
	GetBody() *CreateLoadBalancerUDPListenerResponseBody
}

type CreateLoadBalancerUDPListenerResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLoadBalancerUDPListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLoadBalancerUDPListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerUDPListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerUDPListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLoadBalancerUDPListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLoadBalancerUDPListenerResponse) GetBody() *CreateLoadBalancerUDPListenerResponseBody {
	return s.Body
}

func (s *CreateLoadBalancerUDPListenerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerUDPListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerUDPListenerResponse) SetStatusCode(v int32) *CreateLoadBalancerUDPListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerUDPListenerResponse) SetBody(v *CreateLoadBalancerUDPListenerResponseBody) *CreateLoadBalancerUDPListenerResponse {
	s.Body = v
	return s
}

func (s *CreateLoadBalancerUDPListenerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

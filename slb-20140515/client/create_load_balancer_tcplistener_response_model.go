// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerTCPListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLoadBalancerTCPListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLoadBalancerTCPListenerResponse
	GetStatusCode() *int32
	SetBody(v *CreateLoadBalancerTCPListenerResponseBody) *CreateLoadBalancerTCPListenerResponse
	GetBody() *CreateLoadBalancerTCPListenerResponseBody
}

type CreateLoadBalancerTCPListenerResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLoadBalancerTCPListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLoadBalancerTCPListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerTCPListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerTCPListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLoadBalancerTCPListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLoadBalancerTCPListenerResponse) GetBody() *CreateLoadBalancerTCPListenerResponseBody {
	return s.Body
}

func (s *CreateLoadBalancerTCPListenerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerTCPListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerTCPListenerResponse) SetStatusCode(v int32) *CreateLoadBalancerTCPListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerTCPListenerResponse) SetBody(v *CreateLoadBalancerTCPListenerResponseBody) *CreateLoadBalancerTCPListenerResponse {
	s.Body = v
	return s
}

func (s *CreateLoadBalancerTCPListenerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

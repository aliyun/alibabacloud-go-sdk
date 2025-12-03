// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopLoadBalancerListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopLoadBalancerListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopLoadBalancerListenerResponse
	GetStatusCode() *int32
	SetBody(v *StopLoadBalancerListenerResponseBody) *StopLoadBalancerListenerResponse
	GetBody() *StopLoadBalancerListenerResponseBody
}

type StopLoadBalancerListenerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopLoadBalancerListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopLoadBalancerListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s StopLoadBalancerListenerResponse) GoString() string {
	return s.String()
}

func (s *StopLoadBalancerListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopLoadBalancerListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopLoadBalancerListenerResponse) GetBody() *StopLoadBalancerListenerResponseBody {
	return s.Body
}

func (s *StopLoadBalancerListenerResponse) SetHeaders(v map[string]*string) *StopLoadBalancerListenerResponse {
	s.Headers = v
	return s
}

func (s *StopLoadBalancerListenerResponse) SetStatusCode(v int32) *StopLoadBalancerListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *StopLoadBalancerListenerResponse) SetBody(v *StopLoadBalancerListenerResponseBody) *StopLoadBalancerListenerResponse {
	s.Body = v
	return s
}

func (s *StopLoadBalancerListenerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

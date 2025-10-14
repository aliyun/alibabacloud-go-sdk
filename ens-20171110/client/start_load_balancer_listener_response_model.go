// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartLoadBalancerListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartLoadBalancerListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartLoadBalancerListenerResponse
	GetStatusCode() *int32
	SetBody(v *StartLoadBalancerListenerResponseBody) *StartLoadBalancerListenerResponse
	GetBody() *StartLoadBalancerListenerResponseBody
}

type StartLoadBalancerListenerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartLoadBalancerListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartLoadBalancerListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s StartLoadBalancerListenerResponse) GoString() string {
	return s.String()
}

func (s *StartLoadBalancerListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartLoadBalancerListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartLoadBalancerListenerResponse) GetBody() *StartLoadBalancerListenerResponseBody {
	return s.Body
}

func (s *StartLoadBalancerListenerResponse) SetHeaders(v map[string]*string) *StartLoadBalancerListenerResponse {
	s.Headers = v
	return s
}

func (s *StartLoadBalancerListenerResponse) SetStatusCode(v int32) *StartLoadBalancerListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *StartLoadBalancerListenerResponse) SetBody(v *StartLoadBalancerListenerResponseBody) *StartLoadBalancerListenerResponse {
	s.Body = v
	return s
}

func (s *StartLoadBalancerListenerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerListenerStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLoadBalancerListenerStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLoadBalancerListenerStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetLoadBalancerListenerStatusResponseBody) *SetLoadBalancerListenerStatusResponse
	GetBody() *SetLoadBalancerListenerStatusResponseBody
}

type SetLoadBalancerListenerStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoadBalancerListenerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoadBalancerListenerStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerListenerStatusResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerListenerStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLoadBalancerListenerStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLoadBalancerListenerStatusResponse) GetBody() *SetLoadBalancerListenerStatusResponseBody {
	return s.Body
}

func (s *SetLoadBalancerListenerStatusResponse) SetHeaders(v map[string]*string) *SetLoadBalancerListenerStatusResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerListenerStatusResponse) SetStatusCode(v int32) *SetLoadBalancerListenerStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerListenerStatusResponse) SetBody(v *SetLoadBalancerListenerStatusResponseBody) *SetLoadBalancerListenerStatusResponse {
	s.Body = v
	return s
}

func (s *SetLoadBalancerListenerStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoadBalancerListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLoadBalancerListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLoadBalancerListenerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLoadBalancerListenerResponseBody) *DeleteLoadBalancerListenerResponse
	GetBody() *DeleteLoadBalancerListenerResponseBody
}

type DeleteLoadBalancerListenerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLoadBalancerListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLoadBalancerListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoadBalancerListenerResponse) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLoadBalancerListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLoadBalancerListenerResponse) GetBody() *DeleteLoadBalancerListenerResponseBody {
	return s.Body
}

func (s *DeleteLoadBalancerListenerResponse) SetHeaders(v map[string]*string) *DeleteLoadBalancerListenerResponse {
	s.Headers = v
	return s
}

func (s *DeleteLoadBalancerListenerResponse) SetStatusCode(v int32) *DeleteLoadBalancerListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLoadBalancerListenerResponse) SetBody(v *DeleteLoadBalancerListenerResponseBody) *DeleteLoadBalancerListenerResponse {
	s.Body = v
	return s
}

func (s *DeleteLoadBalancerListenerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

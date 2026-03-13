// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerHTTPListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLoadBalancerHTTPListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLoadBalancerHTTPListenerResponse
	GetStatusCode() *int32
	SetBody(v *CreateLoadBalancerHTTPListenerResponseBody) *CreateLoadBalancerHTTPListenerResponse
	GetBody() *CreateLoadBalancerHTTPListenerResponseBody
}

type CreateLoadBalancerHTTPListenerResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLoadBalancerHTTPListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLoadBalancerHTTPListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerHTTPListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLoadBalancerHTTPListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLoadBalancerHTTPListenerResponse) GetBody() *CreateLoadBalancerHTTPListenerResponseBody {
	return s.Body
}

func (s *CreateLoadBalancerHTTPListenerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerHTTPListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerHTTPListenerResponse) SetStatusCode(v int32) *CreateLoadBalancerHTTPListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerHTTPListenerResponse) SetBody(v *CreateLoadBalancerHTTPListenerResponseBody) *CreateLoadBalancerHTTPListenerResponse {
	s.Body = v
	return s
}

func (s *CreateLoadBalancerHTTPListenerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

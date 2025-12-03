// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoadBalancerHTTPSListenerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLoadBalancerHTTPSListenerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLoadBalancerHTTPSListenerResponse
	GetStatusCode() *int32
	SetBody(v *CreateLoadBalancerHTTPSListenerResponseBody) *CreateLoadBalancerHTTPSListenerResponse
	GetBody() *CreateLoadBalancerHTTPSListenerResponseBody
}

type CreateLoadBalancerHTTPSListenerResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLoadBalancerHTTPSListenerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLoadBalancerHTTPSListenerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLoadBalancerHTTPSListenerResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadBalancerHTTPSListenerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLoadBalancerHTTPSListenerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLoadBalancerHTTPSListenerResponse) GetBody() *CreateLoadBalancerHTTPSListenerResponseBody {
	return s.Body
}

func (s *CreateLoadBalancerHTTPSListenerResponse) SetHeaders(v map[string]*string) *CreateLoadBalancerHTTPSListenerResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerResponse) SetStatusCode(v int32) *CreateLoadBalancerHTTPSListenerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerResponse) SetBody(v *CreateLoadBalancerHTTPSListenerResponseBody) *CreateLoadBalancerHTTPSListenerResponse {
	s.Body = v
	return s
}

func (s *CreateLoadBalancerHTTPSListenerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

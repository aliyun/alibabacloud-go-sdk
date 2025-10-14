// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerHTTPListenerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLoadBalancerHTTPListenerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLoadBalancerHTTPListenerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *SetLoadBalancerHTTPListenerAttributeResponseBody) *SetLoadBalancerHTTPListenerAttributeResponse
	GetBody() *SetLoadBalancerHTTPListenerAttributeResponseBody
}

type SetLoadBalancerHTTPListenerAttributeResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoadBalancerHTTPListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoadBalancerHTTPListenerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerHTTPListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPListenerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLoadBalancerHTTPListenerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLoadBalancerHTTPListenerAttributeResponse) GetBody() *SetLoadBalancerHTTPListenerAttributeResponseBody {
	return s.Body
}

func (s *SetLoadBalancerHTTPListenerAttributeResponse) SetHeaders(v map[string]*string) *SetLoadBalancerHTTPListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeResponse) SetStatusCode(v int32) *SetLoadBalancerHTTPListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeResponse) SetBody(v *SetLoadBalancerHTTPListenerAttributeResponseBody) *SetLoadBalancerHTTPListenerAttributeResponse {
	s.Body = v
	return s
}

func (s *SetLoadBalancerHTTPListenerAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

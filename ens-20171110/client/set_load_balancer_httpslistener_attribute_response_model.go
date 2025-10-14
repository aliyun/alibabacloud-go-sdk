// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerHTTPSListenerAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLoadBalancerHTTPSListenerAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLoadBalancerHTTPSListenerAttributeResponse
	GetStatusCode() *int32
	SetBody(v *SetLoadBalancerHTTPSListenerAttributeResponseBody) *SetLoadBalancerHTTPSListenerAttributeResponse
	GetBody() *SetLoadBalancerHTTPSListenerAttributeResponseBody
}

type SetLoadBalancerHTTPSListenerAttributeResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoadBalancerHTTPSListenerAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoadBalancerHTTPSListenerAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerHTTPSListenerAttributeResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponse) GetBody() *SetLoadBalancerHTTPSListenerAttributeResponseBody {
	return s.Body
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponse) SetHeaders(v map[string]*string) *SetLoadBalancerHTTPSListenerAttributeResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponse) SetStatusCode(v int32) *SetLoadBalancerHTTPSListenerAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponse) SetBody(v *SetLoadBalancerHTTPSListenerAttributeResponseBody) *SetLoadBalancerHTTPSListenerAttributeResponse {
	s.Body = v
	return s
}

func (s *SetLoadBalancerHTTPSListenerAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

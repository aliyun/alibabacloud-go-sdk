// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoadBalancerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLoadBalancerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLoadBalancerResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLoadBalancerResponseBody) *UpdateLoadBalancerResponse
	GetBody() *UpdateLoadBalancerResponseBody
}

type UpdateLoadBalancerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLoadBalancerResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *UpdateLoadBalancerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLoadBalancerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLoadBalancerResponse) GetBody() *UpdateLoadBalancerResponseBody {
	return s.Body
}

func (s *UpdateLoadBalancerResponse) SetHeaders(v map[string]*string) *UpdateLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *UpdateLoadBalancerResponse) SetStatusCode(v int32) *UpdateLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLoadBalancerResponse) SetBody(v *UpdateLoadBalancerResponseBody) *UpdateLoadBalancerResponse {
	s.Body = v
	return s
}

func (s *UpdateLoadBalancerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

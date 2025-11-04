// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLoadBalancersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachLoadBalancersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachLoadBalancersResponse
	GetStatusCode() *int32
	SetBody(v *DetachLoadBalancersResponseBody) *DetachLoadBalancersResponse
	GetBody() *DetachLoadBalancersResponseBody
}

type DetachLoadBalancersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachLoadBalancersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachLoadBalancersResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachLoadBalancersResponse) GoString() string {
	return s.String()
}

func (s *DetachLoadBalancersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachLoadBalancersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachLoadBalancersResponse) GetBody() *DetachLoadBalancersResponseBody {
	return s.Body
}

func (s *DetachLoadBalancersResponse) SetHeaders(v map[string]*string) *DetachLoadBalancersResponse {
	s.Headers = v
	return s
}

func (s *DetachLoadBalancersResponse) SetStatusCode(v int32) *DetachLoadBalancersResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachLoadBalancersResponse) SetBody(v *DetachLoadBalancersResponseBody) *DetachLoadBalancersResponse {
	s.Body = v
	return s
}

func (s *DetachLoadBalancersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLoadBalancerNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLoadBalancerNameResponse
	GetStatusCode() *int32
	SetBody(v *SetLoadBalancerNameResponseBody) *SetLoadBalancerNameResponse
	GetBody() *SetLoadBalancerNameResponseBody
}

type SetLoadBalancerNameResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoadBalancerNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoadBalancerNameResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerNameResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLoadBalancerNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLoadBalancerNameResponse) GetBody() *SetLoadBalancerNameResponseBody {
	return s.Body
}

func (s *SetLoadBalancerNameResponse) SetHeaders(v map[string]*string) *SetLoadBalancerNameResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerNameResponse) SetStatusCode(v int32) *SetLoadBalancerNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerNameResponse) SetBody(v *SetLoadBalancerNameResponseBody) *SetLoadBalancerNameResponse {
	s.Body = v
	return s
}

func (s *SetLoadBalancerNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

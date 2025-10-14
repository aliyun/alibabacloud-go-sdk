// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoadBalancerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLoadBalancerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLoadBalancerResponse
	GetStatusCode() *int32
	SetBody(v *GetLoadBalancerResponseBody) *GetLoadBalancerResponse
	GetBody() *GetLoadBalancerResponseBody
}

type GetLoadBalancerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoadBalancerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoadBalancerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLoadBalancerResponse) GoString() string {
	return s.String()
}

func (s *GetLoadBalancerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLoadBalancerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLoadBalancerResponse) GetBody() *GetLoadBalancerResponseBody {
	return s.Body
}

func (s *GetLoadBalancerResponse) SetHeaders(v map[string]*string) *GetLoadBalancerResponse {
	s.Headers = v
	return s
}

func (s *GetLoadBalancerResponse) SetStatusCode(v int32) *GetLoadBalancerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoadBalancerResponse) SetBody(v *GetLoadBalancerResponseBody) *GetLoadBalancerResponse {
	s.Body = v
	return s
}

func (s *GetLoadBalancerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLoadBalancerStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLoadBalancerStatusResponse
	GetStatusCode() *int32
	SetBody(v *SetLoadBalancerStatusResponseBody) *SetLoadBalancerStatusResponse
	GetBody() *SetLoadBalancerStatusResponseBody
}

type SetLoadBalancerStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoadBalancerStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoadBalancerStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerStatusResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLoadBalancerStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLoadBalancerStatusResponse) GetBody() *SetLoadBalancerStatusResponseBody {
	return s.Body
}

func (s *SetLoadBalancerStatusResponse) SetHeaders(v map[string]*string) *SetLoadBalancerStatusResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerStatusResponse) SetStatusCode(v int32) *SetLoadBalancerStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerStatusResponse) SetBody(v *SetLoadBalancerStatusResponseBody) *SetLoadBalancerStatusResponse {
	s.Body = v
	return s
}

func (s *SetLoadBalancerStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

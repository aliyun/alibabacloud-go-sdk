// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerDeleteProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLoadBalancerDeleteProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLoadBalancerDeleteProtectionResponse
	GetStatusCode() *int32
	SetBody(v *SetLoadBalancerDeleteProtectionResponseBody) *SetLoadBalancerDeleteProtectionResponse
	GetBody() *SetLoadBalancerDeleteProtectionResponseBody
}

type SetLoadBalancerDeleteProtectionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoadBalancerDeleteProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoadBalancerDeleteProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerDeleteProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerDeleteProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLoadBalancerDeleteProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLoadBalancerDeleteProtectionResponse) GetBody() *SetLoadBalancerDeleteProtectionResponseBody {
	return s.Body
}

func (s *SetLoadBalancerDeleteProtectionResponse) SetHeaders(v map[string]*string) *SetLoadBalancerDeleteProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerDeleteProtectionResponse) SetStatusCode(v int32) *SetLoadBalancerDeleteProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionResponse) SetBody(v *SetLoadBalancerDeleteProtectionResponseBody) *SetLoadBalancerDeleteProtectionResponse {
	s.Body = v
	return s
}

func (s *SetLoadBalancerDeleteProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

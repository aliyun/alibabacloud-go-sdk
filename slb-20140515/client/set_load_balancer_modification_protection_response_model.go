// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerModificationProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetLoadBalancerModificationProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetLoadBalancerModificationProtectionResponse
	GetStatusCode() *int32
	SetBody(v *SetLoadBalancerModificationProtectionResponseBody) *SetLoadBalancerModificationProtectionResponse
	GetBody() *SetLoadBalancerModificationProtectionResponseBody
}

type SetLoadBalancerModificationProtectionResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetLoadBalancerModificationProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetLoadBalancerModificationProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerModificationProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerModificationProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetLoadBalancerModificationProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetLoadBalancerModificationProtectionResponse) GetBody() *SetLoadBalancerModificationProtectionResponseBody {
	return s.Body
}

func (s *SetLoadBalancerModificationProtectionResponse) SetHeaders(v map[string]*string) *SetLoadBalancerModificationProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetLoadBalancerModificationProtectionResponse) SetStatusCode(v int32) *SetLoadBalancerModificationProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetLoadBalancerModificationProtectionResponse) SetBody(v *SetLoadBalancerModificationProtectionResponseBody) *SetLoadBalancerModificationProtectionResponse {
	s.Body = v
	return s
}

func (s *SetLoadBalancerModificationProtectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

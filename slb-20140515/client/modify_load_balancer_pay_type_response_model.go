// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerPayTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLoadBalancerPayTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLoadBalancerPayTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLoadBalancerPayTypeResponseBody) *ModifyLoadBalancerPayTypeResponse
	GetBody() *ModifyLoadBalancerPayTypeResponseBody
}

type ModifyLoadBalancerPayTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLoadBalancerPayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLoadBalancerPayTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerPayTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerPayTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLoadBalancerPayTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLoadBalancerPayTypeResponse) GetBody() *ModifyLoadBalancerPayTypeResponseBody {
	return s.Body
}

func (s *ModifyLoadBalancerPayTypeResponse) SetHeaders(v map[string]*string) *ModifyLoadBalancerPayTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoadBalancerPayTypeResponse) SetStatusCode(v int32) *ModifyLoadBalancerPayTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLoadBalancerPayTypeResponse) SetBody(v *ModifyLoadBalancerPayTypeResponseBody) *ModifyLoadBalancerPayTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyLoadBalancerPayTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerInstanceChargeTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLoadBalancerInstanceChargeTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLoadBalancerInstanceChargeTypeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLoadBalancerInstanceChargeTypeResponseBody) *ModifyLoadBalancerInstanceChargeTypeResponse
	GetBody() *ModifyLoadBalancerInstanceChargeTypeResponseBody
}

type ModifyLoadBalancerInstanceChargeTypeResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLoadBalancerInstanceChargeTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLoadBalancerInstanceChargeTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerInstanceChargeTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponse) GetBody() *ModifyLoadBalancerInstanceChargeTypeResponseBody {
	return s.Body
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponse) SetHeaders(v map[string]*string) *ModifyLoadBalancerInstanceChargeTypeResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponse) SetStatusCode(v int32) *ModifyLoadBalancerInstanceChargeTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponse) SetBody(v *ModifyLoadBalancerInstanceChargeTypeResponseBody) *ModifyLoadBalancerInstanceChargeTypeResponse {
	s.Body = v
	return s
}

func (s *ModifyLoadBalancerInstanceChargeTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

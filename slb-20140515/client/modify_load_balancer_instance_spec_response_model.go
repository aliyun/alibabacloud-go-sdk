// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerInstanceSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLoadBalancerInstanceSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLoadBalancerInstanceSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLoadBalancerInstanceSpecResponseBody) *ModifyLoadBalancerInstanceSpecResponse
	GetBody() *ModifyLoadBalancerInstanceSpecResponseBody
}

type ModifyLoadBalancerInstanceSpecResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLoadBalancerInstanceSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLoadBalancerInstanceSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerInstanceSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInstanceSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLoadBalancerInstanceSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLoadBalancerInstanceSpecResponse) GetBody() *ModifyLoadBalancerInstanceSpecResponseBody {
	return s.Body
}

func (s *ModifyLoadBalancerInstanceSpecResponse) SetHeaders(v map[string]*string) *ModifyLoadBalancerInstanceSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecResponse) SetStatusCode(v int32) *ModifyLoadBalancerInstanceSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecResponse) SetBody(v *ModifyLoadBalancerInstanceSpecResponseBody) *ModifyLoadBalancerInstanceSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyLoadBalancerInstanceSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

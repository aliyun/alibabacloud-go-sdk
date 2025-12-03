// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLoadBalancerInternetSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLoadBalancerInternetSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLoadBalancerInternetSpecResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLoadBalancerInternetSpecResponseBody) *ModifyLoadBalancerInternetSpecResponse
	GetBody() *ModifyLoadBalancerInternetSpecResponseBody
}

type ModifyLoadBalancerInternetSpecResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLoadBalancerInternetSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLoadBalancerInternetSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLoadBalancerInternetSpecResponse) GoString() string {
	return s.String()
}

func (s *ModifyLoadBalancerInternetSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLoadBalancerInternetSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLoadBalancerInternetSpecResponse) GetBody() *ModifyLoadBalancerInternetSpecResponseBody {
	return s.Body
}

func (s *ModifyLoadBalancerInternetSpecResponse) SetHeaders(v map[string]*string) *ModifyLoadBalancerInternetSpecResponse {
	s.Headers = v
	return s
}

func (s *ModifyLoadBalancerInternetSpecResponse) SetStatusCode(v int32) *ModifyLoadBalancerInternetSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLoadBalancerInternetSpecResponse) SetBody(v *ModifyLoadBalancerInternetSpecResponseBody) *ModifyLoadBalancerInternetSpecResponse {
	s.Body = v
	return s
}

func (s *ModifyLoadBalancerInternetSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

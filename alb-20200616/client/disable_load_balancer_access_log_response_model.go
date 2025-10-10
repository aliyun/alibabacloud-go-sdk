// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLoadBalancerAccessLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableLoadBalancerAccessLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableLoadBalancerAccessLogResponse
	GetStatusCode() *int32
	SetBody(v *DisableLoadBalancerAccessLogResponseBody) *DisableLoadBalancerAccessLogResponse
	GetBody() *DisableLoadBalancerAccessLogResponseBody
}

type DisableLoadBalancerAccessLogResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableLoadBalancerAccessLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableLoadBalancerAccessLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableLoadBalancerAccessLogResponse) GoString() string {
	return s.String()
}

func (s *DisableLoadBalancerAccessLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableLoadBalancerAccessLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableLoadBalancerAccessLogResponse) GetBody() *DisableLoadBalancerAccessLogResponseBody {
	return s.Body
}

func (s *DisableLoadBalancerAccessLogResponse) SetHeaders(v map[string]*string) *DisableLoadBalancerAccessLogResponse {
	s.Headers = v
	return s
}

func (s *DisableLoadBalancerAccessLogResponse) SetStatusCode(v int32) *DisableLoadBalancerAccessLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableLoadBalancerAccessLogResponse) SetBody(v *DisableLoadBalancerAccessLogResponseBody) *DisableLoadBalancerAccessLogResponse {
	s.Body = v
	return s
}

func (s *DisableLoadBalancerAccessLogResponse) Validate() error {
	return dara.Validate(s)
}

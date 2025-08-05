// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatFirewallControlPolicyBatchResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNatFirewallControlPolicyBatchResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNatFirewallControlPolicyBatchResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNatFirewallControlPolicyBatchResponseBody) *DeleteNatFirewallControlPolicyBatchResponse
	GetBody() *DeleteNatFirewallControlPolicyBatchResponseBody
}

type DeleteNatFirewallControlPolicyBatchResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNatFirewallControlPolicyBatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNatFirewallControlPolicyBatchResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatFirewallControlPolicyBatchResponse) GoString() string {
	return s.String()
}

func (s *DeleteNatFirewallControlPolicyBatchResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNatFirewallControlPolicyBatchResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNatFirewallControlPolicyBatchResponse) GetBody() *DeleteNatFirewallControlPolicyBatchResponseBody {
	return s.Body
}

func (s *DeleteNatFirewallControlPolicyBatchResponse) SetHeaders(v map[string]*string) *DeleteNatFirewallControlPolicyBatchResponse {
	s.Headers = v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchResponse) SetStatusCode(v int32) *DeleteNatFirewallControlPolicyBatchResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchResponse) SetBody(v *DeleteNatFirewallControlPolicyBatchResponseBody) *DeleteNatFirewallControlPolicyBatchResponse {
	s.Body = v
	return s
}

func (s *DeleteNatFirewallControlPolicyBatchResponse) Validate() error {
	return dara.Validate(s)
}

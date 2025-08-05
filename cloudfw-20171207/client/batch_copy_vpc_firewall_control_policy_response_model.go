// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCopyVpcFirewallControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchCopyVpcFirewallControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchCopyVpcFirewallControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *BatchCopyVpcFirewallControlPolicyResponseBody) *BatchCopyVpcFirewallControlPolicyResponse
	GetBody() *BatchCopyVpcFirewallControlPolicyResponseBody
}

type BatchCopyVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchCopyVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchCopyVpcFirewallControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchCopyVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) GetBody() *BatchCopyVpcFirewallControlPolicyResponseBody {
	return s.Body
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *BatchCopyVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *BatchCopyVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) SetBody(v *BatchCopyVpcFirewallControlPolicyResponseBody) *BatchCopyVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

func (s *BatchCopyVpcFirewallControlPolicyResponse) Validate() error {
	return dara.Validate(s)
}

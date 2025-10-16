// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteVpcFirewallControlPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchDeleteVpcFirewallControlPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchDeleteVpcFirewallControlPolicyResponse
	GetStatusCode() *int32
	SetBody(v *BatchDeleteVpcFirewallControlPolicyResponseBody) *BatchDeleteVpcFirewallControlPolicyResponse
	GetBody() *BatchDeleteVpcFirewallControlPolicyResponseBody
}

type BatchDeleteVpcFirewallControlPolicyResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchDeleteVpcFirewallControlPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchDeleteVpcFirewallControlPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteVpcFirewallControlPolicyResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteVpcFirewallControlPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchDeleteVpcFirewallControlPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchDeleteVpcFirewallControlPolicyResponse) GetBody() *BatchDeleteVpcFirewallControlPolicyResponseBody {
	return s.Body
}

func (s *BatchDeleteVpcFirewallControlPolicyResponse) SetHeaders(v map[string]*string) *BatchDeleteVpcFirewallControlPolicyResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteVpcFirewallControlPolicyResponse) SetStatusCode(v int32) *BatchDeleteVpcFirewallControlPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteVpcFirewallControlPolicyResponse) SetBody(v *BatchDeleteVpcFirewallControlPolicyResponseBody) *BatchDeleteVpcFirewallControlPolicyResponse {
	s.Body = v
	return s
}

func (s *BatchDeleteVpcFirewallControlPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

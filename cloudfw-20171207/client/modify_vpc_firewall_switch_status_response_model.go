// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallSwitchStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcFirewallSwitchStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcFirewallSwitchStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcFirewallSwitchStatusResponseBody) *ModifyVpcFirewallSwitchStatusResponse
	GetBody() *ModifyVpcFirewallSwitchStatusResponseBody
}

type ModifyVpcFirewallSwitchStatusResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallSwitchStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallSwitchStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallSwitchStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcFirewallSwitchStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcFirewallSwitchStatusResponse) GetBody() *ModifyVpcFirewallSwitchStatusResponseBody {
	return s.Body
}

func (s *ModifyVpcFirewallSwitchStatusResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallSwitchStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusResponse) SetStatusCode(v int32) *ModifyVpcFirewallSwitchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusResponse) SetBody(v *ModifyVpcFirewallSwitchStatusResponseBody) *ModifyVpcFirewallSwitchStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcFirewallSwitchStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

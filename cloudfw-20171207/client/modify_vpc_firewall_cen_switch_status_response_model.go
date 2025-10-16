// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallCenSwitchStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcFirewallCenSwitchStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcFirewallCenSwitchStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcFirewallCenSwitchStatusResponseBody) *ModifyVpcFirewallCenSwitchStatusResponse
	GetBody() *ModifyVpcFirewallCenSwitchStatusResponseBody
}

type ModifyVpcFirewallCenSwitchStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallCenSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallCenSwitchStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallCenSwitchStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) GetBody() *ModifyVpcFirewallCenSwitchStatusResponseBody {
	return s.Body
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallCenSwitchStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) SetStatusCode(v int32) *ModifyVpcFirewallCenSwitchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) SetBody(v *ModifyVpcFirewallCenSwitchStatusResponseBody) *ModifyVpcFirewallCenSwitchStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcFirewallCenSwitchStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

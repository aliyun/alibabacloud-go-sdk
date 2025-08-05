// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallConfigureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcFirewallConfigureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcFirewallConfigureResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcFirewallConfigureResponseBody) *ModifyVpcFirewallConfigureResponse
	GetBody() *ModifyVpcFirewallConfigureResponseBody
}

type ModifyVpcFirewallConfigureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallConfigureResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallConfigureResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallConfigureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcFirewallConfigureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcFirewallConfigureResponse) GetBody() *ModifyVpcFirewallConfigureResponseBody {
	return s.Body
}

func (s *ModifyVpcFirewallConfigureResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallConfigureResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallConfigureResponse) SetStatusCode(v int32) *ModifyVpcFirewallConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallConfigureResponse) SetBody(v *ModifyVpcFirewallConfigureResponseBody) *ModifyVpcFirewallConfigureResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcFirewallConfigureResponse) Validate() error {
	return dara.Validate(s)
}

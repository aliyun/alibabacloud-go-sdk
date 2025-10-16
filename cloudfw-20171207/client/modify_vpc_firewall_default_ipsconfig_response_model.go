// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallDefaultIPSConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcFirewallDefaultIPSConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcFirewallDefaultIPSConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcFirewallDefaultIPSConfigResponseBody) *ModifyVpcFirewallDefaultIPSConfigResponse
	GetBody() *ModifyVpcFirewallDefaultIPSConfigResponseBody
}

type ModifyVpcFirewallDefaultIPSConfigResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallDefaultIPSConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallDefaultIPSConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallDefaultIPSConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) GetBody() *ModifyVpcFirewallDefaultIPSConfigResponseBody {
	return s.Body
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallDefaultIPSConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) SetStatusCode(v int32) *ModifyVpcFirewallDefaultIPSConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) SetBody(v *ModifyVpcFirewallDefaultIPSConfigResponseBody) *ModifyVpcFirewallDefaultIPSConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcFirewallDefaultIPSConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

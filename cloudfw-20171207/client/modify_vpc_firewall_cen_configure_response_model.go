// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcFirewallCenConfigureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyVpcFirewallCenConfigureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyVpcFirewallCenConfigureResponse
	GetStatusCode() *int32
	SetBody(v *ModifyVpcFirewallCenConfigureResponseBody) *ModifyVpcFirewallCenConfigureResponse
	GetBody() *ModifyVpcFirewallCenConfigureResponseBody
}

type ModifyVpcFirewallCenConfigureResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyVpcFirewallCenConfigureResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcFirewallCenConfigureResponse) GoString() string {
	return s.String()
}

func (s *ModifyVpcFirewallCenConfigureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyVpcFirewallCenConfigureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyVpcFirewallCenConfigureResponse) GetBody() *ModifyVpcFirewallCenConfigureResponseBody {
	return s.Body
}

func (s *ModifyVpcFirewallCenConfigureResponse) SetHeaders(v map[string]*string) *ModifyVpcFirewallCenConfigureResponse {
	s.Headers = v
	return s
}

func (s *ModifyVpcFirewallCenConfigureResponse) SetStatusCode(v int32) *ModifyVpcFirewallCenConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyVpcFirewallCenConfigureResponse) SetBody(v *ModifyVpcFirewallCenConfigureResponseBody) *ModifyVpcFirewallCenConfigureResponse {
	s.Body = v
	return s
}

func (s *ModifyVpcFirewallCenConfigureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcFirewallConfigureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpcFirewallConfigureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpcFirewallConfigureResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpcFirewallConfigureResponseBody) *DeleteVpcFirewallConfigureResponse
	GetBody() *DeleteVpcFirewallConfigureResponseBody
}

type DeleteVpcFirewallConfigureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcFirewallConfigureResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcFirewallConfigureResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallConfigureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpcFirewallConfigureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpcFirewallConfigureResponse) GetBody() *DeleteVpcFirewallConfigureResponseBody {
	return s.Body
}

func (s *DeleteVpcFirewallConfigureResponse) SetHeaders(v map[string]*string) *DeleteVpcFirewallConfigureResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcFirewallConfigureResponse) SetStatusCode(v int32) *DeleteVpcFirewallConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcFirewallConfigureResponse) SetBody(v *DeleteVpcFirewallConfigureResponseBody) *DeleteVpcFirewallConfigureResponse {
	s.Body = v
	return s
}

func (s *DeleteVpcFirewallConfigureResponse) Validate() error {
	return dara.Validate(s)
}

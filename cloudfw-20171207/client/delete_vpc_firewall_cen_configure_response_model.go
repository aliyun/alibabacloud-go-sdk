// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcFirewallCenConfigureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVpcFirewallCenConfigureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVpcFirewallCenConfigureResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVpcFirewallCenConfigureResponseBody) *DeleteVpcFirewallCenConfigureResponse
	GetBody() *DeleteVpcFirewallCenConfigureResponseBody
}

type DeleteVpcFirewallCenConfigureResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVpcFirewallCenConfigureResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcFirewallCenConfigureResponse) GoString() string {
	return s.String()
}

func (s *DeleteVpcFirewallCenConfigureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVpcFirewallCenConfigureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVpcFirewallCenConfigureResponse) GetBody() *DeleteVpcFirewallCenConfigureResponseBody {
	return s.Body
}

func (s *DeleteVpcFirewallCenConfigureResponse) SetHeaders(v map[string]*string) *DeleteVpcFirewallCenConfigureResponse {
	s.Headers = v
	return s
}

func (s *DeleteVpcFirewallCenConfigureResponse) SetStatusCode(v int32) *DeleteVpcFirewallCenConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVpcFirewallCenConfigureResponse) SetBody(v *DeleteVpcFirewallCenConfigureResponseBody) *DeleteVpcFirewallCenConfigureResponse {
	s.Body = v
	return s
}

func (s *DeleteVpcFirewallCenConfigureResponse) Validate() error {
	return dara.Validate(s)
}

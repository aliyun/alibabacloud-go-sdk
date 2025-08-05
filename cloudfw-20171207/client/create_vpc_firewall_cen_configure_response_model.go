// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallCenConfigureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcFirewallCenConfigureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcFirewallCenConfigureResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcFirewallCenConfigureResponseBody) *CreateVpcFirewallCenConfigureResponse
	GetBody() *CreateVpcFirewallCenConfigureResponseBody
}

type CreateVpcFirewallCenConfigureResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcFirewallCenConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcFirewallCenConfigureResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallCenConfigureResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenConfigureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcFirewallCenConfigureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcFirewallCenConfigureResponse) GetBody() *CreateVpcFirewallCenConfigureResponseBody {
	return s.Body
}

func (s *CreateVpcFirewallCenConfigureResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallCenConfigureResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallCenConfigureResponse) SetStatusCode(v int32) *CreateVpcFirewallCenConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallCenConfigureResponse) SetBody(v *CreateVpcFirewallCenConfigureResponseBody) *CreateVpcFirewallCenConfigureResponse {
	s.Body = v
	return s
}

func (s *CreateVpcFirewallCenConfigureResponse) Validate() error {
	return dara.Validate(s)
}

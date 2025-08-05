// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallCenManualConfigureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcFirewallCenManualConfigureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcFirewallCenManualConfigureResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcFirewallCenManualConfigureResponseBody) *CreateVpcFirewallCenManualConfigureResponse
	GetBody() *CreateVpcFirewallCenManualConfigureResponseBody
}

type CreateVpcFirewallCenManualConfigureResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcFirewallCenManualConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcFirewallCenManualConfigureResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallCenManualConfigureResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallCenManualConfigureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcFirewallCenManualConfigureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcFirewallCenManualConfigureResponse) GetBody() *CreateVpcFirewallCenManualConfigureResponseBody {
	return s.Body
}

func (s *CreateVpcFirewallCenManualConfigureResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallCenManualConfigureResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallCenManualConfigureResponse) SetStatusCode(v int32) *CreateVpcFirewallCenManualConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallCenManualConfigureResponse) SetBody(v *CreateVpcFirewallCenManualConfigureResponseBody) *CreateVpcFirewallCenManualConfigureResponse {
	s.Body = v
	return s
}

func (s *CreateVpcFirewallCenManualConfigureResponse) Validate() error {
	return dara.Validate(s)
}

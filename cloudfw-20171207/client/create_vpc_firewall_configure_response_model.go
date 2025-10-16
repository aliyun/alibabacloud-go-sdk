// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcFirewallConfigureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVpcFirewallConfigureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVpcFirewallConfigureResponse
	GetStatusCode() *int32
	SetBody(v *CreateVpcFirewallConfigureResponseBody) *CreateVpcFirewallConfigureResponse
	GetBody() *CreateVpcFirewallConfigureResponseBody
}

type CreateVpcFirewallConfigureResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVpcFirewallConfigureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVpcFirewallConfigureResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcFirewallConfigureResponse) GoString() string {
	return s.String()
}

func (s *CreateVpcFirewallConfigureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVpcFirewallConfigureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVpcFirewallConfigureResponse) GetBody() *CreateVpcFirewallConfigureResponseBody {
	return s.Body
}

func (s *CreateVpcFirewallConfigureResponse) SetHeaders(v map[string]*string) *CreateVpcFirewallConfigureResponse {
	s.Headers = v
	return s
}

func (s *CreateVpcFirewallConfigureResponse) SetStatusCode(v int32) *CreateVpcFirewallConfigureResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVpcFirewallConfigureResponse) SetBody(v *CreateVpcFirewallConfigureResponseBody) *CreateVpcFirewallConfigureResponse {
	s.Body = v
	return s
}

func (s *CreateVpcFirewallConfigureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

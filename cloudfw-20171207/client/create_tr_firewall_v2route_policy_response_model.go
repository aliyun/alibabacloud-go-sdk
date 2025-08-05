// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrFirewallV2RoutePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrFirewallV2RoutePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrFirewallV2RoutePolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrFirewallV2RoutePolicyResponseBody) *CreateTrFirewallV2RoutePolicyResponse
	GetBody() *CreateTrFirewallV2RoutePolicyResponseBody
}

type CreateTrFirewallV2RoutePolicyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrFirewallV2RoutePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrFirewallV2RoutePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrFirewallV2RoutePolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateTrFirewallV2RoutePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrFirewallV2RoutePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrFirewallV2RoutePolicyResponse) GetBody() *CreateTrFirewallV2RoutePolicyResponseBody {
	return s.Body
}

func (s *CreateTrFirewallV2RoutePolicyResponse) SetHeaders(v map[string]*string) *CreateTrFirewallV2RoutePolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyResponse) SetStatusCode(v int32) *CreateTrFirewallV2RoutePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyResponse) SetBody(v *CreateTrFirewallV2RoutePolicyResponseBody) *CreateTrFirewallV2RoutePolicyResponse {
	s.Body = v
	return s
}

func (s *CreateTrFirewallV2RoutePolicyResponse) Validate() error {
	return dara.Validate(s)
}

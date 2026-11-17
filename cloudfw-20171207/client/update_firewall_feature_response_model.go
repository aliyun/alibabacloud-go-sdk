// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFirewallFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFirewallFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFirewallFeatureResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFirewallFeatureResponseBody) *UpdateFirewallFeatureResponse
	GetBody() *UpdateFirewallFeatureResponseBody
}

type UpdateFirewallFeatureResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFirewallFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFirewallFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFirewallFeatureResponse) GoString() string {
	return s.String()
}

func (s *UpdateFirewallFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFirewallFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFirewallFeatureResponse) GetBody() *UpdateFirewallFeatureResponseBody {
	return s.Body
}

func (s *UpdateFirewallFeatureResponse) SetHeaders(v map[string]*string) *UpdateFirewallFeatureResponse {
	s.Headers = v
	return s
}

func (s *UpdateFirewallFeatureResponse) SetStatusCode(v int32) *UpdateFirewallFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFirewallFeatureResponse) SetBody(v *UpdateFirewallFeatureResponseBody) *UpdateFirewallFeatureResponse {
	s.Body = v
	return s
}

func (s *UpdateFirewallFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFirewallFeatureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetFirewallFeatureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetFirewallFeatureResponse
	GetStatusCode() *int32
	SetBody(v *GetFirewallFeatureResponseBody) *GetFirewallFeatureResponse
	GetBody() *GetFirewallFeatureResponseBody
}

type GetFirewallFeatureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFirewallFeatureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFirewallFeatureResponse) String() string {
	return dara.Prettify(s)
}

func (s GetFirewallFeatureResponse) GoString() string {
	return s.String()
}

func (s *GetFirewallFeatureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetFirewallFeatureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetFirewallFeatureResponse) GetBody() *GetFirewallFeatureResponseBody {
	return s.Body
}

func (s *GetFirewallFeatureResponse) SetHeaders(v map[string]*string) *GetFirewallFeatureResponse {
	s.Headers = v
	return s
}

func (s *GetFirewallFeatureResponse) SetStatusCode(v int32) *GetFirewallFeatureResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFirewallFeatureResponse) SetBody(v *GetFirewallFeatureResponseBody) *GetFirewallFeatureResponse {
	s.Body = v
	return s
}

func (s *GetFirewallFeatureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

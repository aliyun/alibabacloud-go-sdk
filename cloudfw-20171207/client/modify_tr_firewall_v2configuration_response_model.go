// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrFirewallV2ConfigurationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTrFirewallV2ConfigurationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTrFirewallV2ConfigurationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTrFirewallV2ConfigurationResponseBody) *ModifyTrFirewallV2ConfigurationResponse
	GetBody() *ModifyTrFirewallV2ConfigurationResponseBody
}

type ModifyTrFirewallV2ConfigurationResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTrFirewallV2ConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTrFirewallV2ConfigurationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrFirewallV2ConfigurationResponse) GoString() string {
	return s.String()
}

func (s *ModifyTrFirewallV2ConfigurationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTrFirewallV2ConfigurationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTrFirewallV2ConfigurationResponse) GetBody() *ModifyTrFirewallV2ConfigurationResponseBody {
	return s.Body
}

func (s *ModifyTrFirewallV2ConfigurationResponse) SetHeaders(v map[string]*string) *ModifyTrFirewallV2ConfigurationResponse {
	s.Headers = v
	return s
}

func (s *ModifyTrFirewallV2ConfigurationResponse) SetStatusCode(v int32) *ModifyTrFirewallV2ConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTrFirewallV2ConfigurationResponse) SetBody(v *ModifyTrFirewallV2ConfigurationResponseBody) *ModifyTrFirewallV2ConfigurationResponse {
	s.Body = v
	return s
}

func (s *ModifyTrFirewallV2ConfigurationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

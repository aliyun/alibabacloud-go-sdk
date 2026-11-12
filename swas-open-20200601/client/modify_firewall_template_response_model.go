// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFirewallTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyFirewallTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyFirewallTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyFirewallTemplateResponseBody) *ModifyFirewallTemplateResponse
	GetBody() *ModifyFirewallTemplateResponseBody
}

type ModifyFirewallTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFirewallTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyFirewallTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyFirewallTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyFirewallTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyFirewallTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyFirewallTemplateResponse) GetBody() *ModifyFirewallTemplateResponseBody {
	return s.Body
}

func (s *ModifyFirewallTemplateResponse) SetHeaders(v map[string]*string) *ModifyFirewallTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyFirewallTemplateResponse) SetStatusCode(v int32) *ModifyFirewallTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFirewallTemplateResponse) SetBody(v *ModifyFirewallTemplateResponseBody) *ModifyFirewallTemplateResponse {
	s.Body = v
	return s
}

func (s *ModifyFirewallTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

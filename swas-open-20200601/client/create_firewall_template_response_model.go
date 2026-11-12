// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFirewallTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFirewallTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFirewallTemplateResponse
	GetStatusCode() *int32
	SetBody(v *CreateFirewallTemplateResponseBody) *CreateFirewallTemplateResponse
	GetBody() *CreateFirewallTemplateResponseBody
}

type CreateFirewallTemplateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFirewallTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFirewallTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFirewallTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateFirewallTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFirewallTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFirewallTemplateResponse) GetBody() *CreateFirewallTemplateResponseBody {
	return s.Body
}

func (s *CreateFirewallTemplateResponse) SetHeaders(v map[string]*string) *CreateFirewallTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateFirewallTemplateResponse) SetStatusCode(v int32) *CreateFirewallTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFirewallTemplateResponse) SetBody(v *CreateFirewallTemplateResponseBody) *CreateFirewallTemplateResponse {
	s.Body = v
	return s
}

func (s *CreateFirewallTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyFirewallTemplateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyFirewallTemplateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyFirewallTemplateResponse
	GetStatusCode() *int32
	SetBody(v *ApplyFirewallTemplateResponseBody) *ApplyFirewallTemplateResponse
	GetBody() *ApplyFirewallTemplateResponseBody
}

type ApplyFirewallTemplateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyFirewallTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyFirewallTemplateResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyFirewallTemplateResponse) GoString() string {
	return s.String()
}

func (s *ApplyFirewallTemplateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyFirewallTemplateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyFirewallTemplateResponse) GetBody() *ApplyFirewallTemplateResponseBody {
	return s.Body
}

func (s *ApplyFirewallTemplateResponse) SetHeaders(v map[string]*string) *ApplyFirewallTemplateResponse {
	s.Headers = v
	return s
}

func (s *ApplyFirewallTemplateResponse) SetStatusCode(v int32) *ApplyFirewallTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyFirewallTemplateResponse) SetBody(v *ApplyFirewallTemplateResponseBody) *ApplyFirewallTemplateResponse {
	s.Body = v
	return s
}

func (s *ApplyFirewallTemplateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

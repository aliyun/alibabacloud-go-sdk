// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatFirewallPreCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNatFirewallPreCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNatFirewallPreCheckResponse
	GetStatusCode() *int32
	SetBody(v *CreateNatFirewallPreCheckResponseBody) *CreateNatFirewallPreCheckResponse
	GetBody() *CreateNatFirewallPreCheckResponseBody
}

type CreateNatFirewallPreCheckResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNatFirewallPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNatFirewallPreCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNatFirewallPreCheckResponse) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallPreCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNatFirewallPreCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNatFirewallPreCheckResponse) GetBody() *CreateNatFirewallPreCheckResponseBody {
	return s.Body
}

func (s *CreateNatFirewallPreCheckResponse) SetHeaders(v map[string]*string) *CreateNatFirewallPreCheckResponse {
	s.Headers = v
	return s
}

func (s *CreateNatFirewallPreCheckResponse) SetStatusCode(v int32) *CreateNatFirewallPreCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNatFirewallPreCheckResponse) SetBody(v *CreateNatFirewallPreCheckResponseBody) *CreateNatFirewallPreCheckResponse {
	s.Body = v
	return s
}

func (s *CreateNatFirewallPreCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

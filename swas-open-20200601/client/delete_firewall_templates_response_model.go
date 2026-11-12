// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFirewallTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFirewallTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFirewallTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFirewallTemplatesResponseBody) *DeleteFirewallTemplatesResponse
	GetBody() *DeleteFirewallTemplatesResponseBody
}

type DeleteFirewallTemplatesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFirewallTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFirewallTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFirewallTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFirewallTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFirewallTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFirewallTemplatesResponse) GetBody() *DeleteFirewallTemplatesResponseBody {
	return s.Body
}

func (s *DeleteFirewallTemplatesResponse) SetHeaders(v map[string]*string) *DeleteFirewallTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFirewallTemplatesResponse) SetStatusCode(v int32) *DeleteFirewallTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFirewallTemplatesResponse) SetBody(v *DeleteFirewallTemplatesResponseBody) *DeleteFirewallTemplatesResponse {
	s.Body = v
	return s
}

func (s *DeleteFirewallTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

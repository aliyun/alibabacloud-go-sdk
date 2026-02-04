// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindTotpAuthenticatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindTotpAuthenticatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindTotpAuthenticatorResponse
	GetStatusCode() *int32
	SetBody(v *UnbindTotpAuthenticatorResponseBody) *UnbindTotpAuthenticatorResponse
	GetBody() *UnbindTotpAuthenticatorResponseBody
}

type UnbindTotpAuthenticatorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindTotpAuthenticatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindTotpAuthenticatorResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindTotpAuthenticatorResponse) GoString() string {
	return s.String()
}

func (s *UnbindTotpAuthenticatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindTotpAuthenticatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindTotpAuthenticatorResponse) GetBody() *UnbindTotpAuthenticatorResponseBody {
	return s.Body
}

func (s *UnbindTotpAuthenticatorResponse) SetHeaders(v map[string]*string) *UnbindTotpAuthenticatorResponse {
	s.Headers = v
	return s
}

func (s *UnbindTotpAuthenticatorResponse) SetStatusCode(v int32) *UnbindTotpAuthenticatorResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindTotpAuthenticatorResponse) SetBody(v *UnbindTotpAuthenticatorResponseBody) *UnbindTotpAuthenticatorResponse {
	s.Body = v
	return s
}

func (s *UnbindTotpAuthenticatorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

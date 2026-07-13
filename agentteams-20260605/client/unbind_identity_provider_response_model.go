// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *UnbindIdentityProviderResponseBody) *UnbindIdentityProviderResponse
	GetBody() *UnbindIdentityProviderResponseBody
}

type UnbindIdentityProviderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *UnbindIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindIdentityProviderResponse) GetBody() *UnbindIdentityProviderResponseBody {
	return s.Body
}

func (s *UnbindIdentityProviderResponse) SetHeaders(v map[string]*string) *UnbindIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *UnbindIdentityProviderResponse) SetStatusCode(v int32) *UnbindIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindIdentityProviderResponse) SetBody(v *UnbindIdentityProviderResponseBody) *UnbindIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *UnbindIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

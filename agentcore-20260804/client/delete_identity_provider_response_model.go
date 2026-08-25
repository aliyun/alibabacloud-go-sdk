// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIdentityProviderResponseBody) *DeleteIdentityProviderResponse
	GetBody() *DeleteIdentityProviderResponseBody
}

type DeleteIdentityProviderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *DeleteIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIdentityProviderResponse) GetBody() *DeleteIdentityProviderResponseBody {
	return s.Body
}

func (s *DeleteIdentityProviderResponse) SetHeaders(v map[string]*string) *DeleteIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *DeleteIdentityProviderResponse) SetStatusCode(v int32) *DeleteIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIdentityProviderResponse) SetBody(v *DeleteIdentityProviderResponseBody) *DeleteIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *DeleteIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

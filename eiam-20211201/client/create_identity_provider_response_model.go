// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *CreateIdentityProviderResponseBody) *CreateIdentityProviderResponse
	GetBody() *CreateIdentityProviderResponseBody
}

type CreateIdentityProviderResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIdentityProviderResponse) GetBody() *CreateIdentityProviderResponseBody {
	return s.Body
}

func (s *CreateIdentityProviderResponse) SetHeaders(v map[string]*string) *CreateIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateIdentityProviderResponse) SetStatusCode(v int32) *CreateIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIdentityProviderResponse) SetBody(v *CreateIdentityProviderResponseBody) *CreateIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *CreateIdentityProviderResponse) Validate() error {
	return dara.Validate(s)
}

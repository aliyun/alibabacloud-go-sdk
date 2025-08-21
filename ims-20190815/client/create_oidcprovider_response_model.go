// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOIDCProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOIDCProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOIDCProviderResponse
	GetStatusCode() *int32
	SetBody(v *CreateOIDCProviderResponseBody) *CreateOIDCProviderResponse
	GetBody() *CreateOIDCProviderResponseBody
}

type CreateOIDCProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOIDCProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *CreateOIDCProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOIDCProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOIDCProviderResponse) GetBody() *CreateOIDCProviderResponseBody {
	return s.Body
}

func (s *CreateOIDCProviderResponse) SetHeaders(v map[string]*string) *CreateOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *CreateOIDCProviderResponse) SetStatusCode(v int32) *CreateOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOIDCProviderResponse) SetBody(v *CreateOIDCProviderResponseBody) *CreateOIDCProviderResponse {
	s.Body = v
	return s
}

func (s *CreateOIDCProviderResponse) Validate() error {
	return dara.Validate(s)
}

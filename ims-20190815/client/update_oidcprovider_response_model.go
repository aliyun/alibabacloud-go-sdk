// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOIDCProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOIDCProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOIDCProviderResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOIDCProviderResponseBody) *UpdateOIDCProviderResponse
	GetBody() *UpdateOIDCProviderResponseBody
}

type UpdateOIDCProviderResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOIDCProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *UpdateOIDCProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOIDCProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOIDCProviderResponse) GetBody() *UpdateOIDCProviderResponseBody {
	return s.Body
}

func (s *UpdateOIDCProviderResponse) SetHeaders(v map[string]*string) *UpdateOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *UpdateOIDCProviderResponse) SetStatusCode(v int32) *UpdateOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOIDCProviderResponse) SetBody(v *UpdateOIDCProviderResponseBody) *UpdateOIDCProviderResponse {
	s.Body = v
	return s
}

func (s *UpdateOIDCProviderResponse) Validate() error {
	return dara.Validate(s)
}

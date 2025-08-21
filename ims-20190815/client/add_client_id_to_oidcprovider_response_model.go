// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddClientIdToOIDCProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddClientIdToOIDCProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddClientIdToOIDCProviderResponse
	GetStatusCode() *int32
	SetBody(v *AddClientIdToOIDCProviderResponseBody) *AddClientIdToOIDCProviderResponse
	GetBody() *AddClientIdToOIDCProviderResponseBody
}

type AddClientIdToOIDCProviderResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddClientIdToOIDCProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddClientIdToOIDCProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s AddClientIdToOIDCProviderResponse) GoString() string {
	return s.String()
}

func (s *AddClientIdToOIDCProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddClientIdToOIDCProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddClientIdToOIDCProviderResponse) GetBody() *AddClientIdToOIDCProviderResponseBody {
	return s.Body
}

func (s *AddClientIdToOIDCProviderResponse) SetHeaders(v map[string]*string) *AddClientIdToOIDCProviderResponse {
	s.Headers = v
	return s
}

func (s *AddClientIdToOIDCProviderResponse) SetStatusCode(v int32) *AddClientIdToOIDCProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *AddClientIdToOIDCProviderResponse) SetBody(v *AddClientIdToOIDCProviderResponseBody) *AddClientIdToOIDCProviderResponse {
	s.Body = v
	return s
}

func (s *AddClientIdToOIDCProviderResponse) Validate() error {
	return dara.Validate(s)
}

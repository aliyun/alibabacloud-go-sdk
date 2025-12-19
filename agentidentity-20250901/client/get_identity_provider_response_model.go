// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentityProviderResponseBody) *GetIdentityProviderResponse
	GetBody() *GetIdentityProviderResponseBody
}

type GetIdentityProviderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentityProviderResponse) GetBody() *GetIdentityProviderResponseBody {
	return s.Body
}

func (s *GetIdentityProviderResponse) SetHeaders(v map[string]*string) *GetIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityProviderResponse) SetStatusCode(v int32) *GetIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityProviderResponse) SetBody(v *GetIdentityProviderResponseBody) *GetIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *GetIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

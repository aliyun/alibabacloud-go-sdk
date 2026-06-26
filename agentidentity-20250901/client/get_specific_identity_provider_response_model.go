// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpecificIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSpecificIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSpecificIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetSpecificIdentityProviderResponseBody) *GetSpecificIdentityProviderResponse
	GetBody() *GetSpecificIdentityProviderResponseBody
}

type GetSpecificIdentityProviderResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpecificIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpecificIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSpecificIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *GetSpecificIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSpecificIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSpecificIdentityProviderResponse) GetBody() *GetSpecificIdentityProviderResponseBody {
	return s.Body
}

func (s *GetSpecificIdentityProviderResponse) SetHeaders(v map[string]*string) *GetSpecificIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *GetSpecificIdentityProviderResponse) SetStatusCode(v int32) *GetSpecificIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpecificIdentityProviderResponse) SetBody(v *GetSpecificIdentityProviderResponseBody) *GetSpecificIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *GetSpecificIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

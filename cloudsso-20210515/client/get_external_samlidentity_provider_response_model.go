// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExternalSAMLIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExternalSAMLIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExternalSAMLIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetExternalSAMLIdentityProviderResponseBody) *GetExternalSAMLIdentityProviderResponse
	GetBody() *GetExternalSAMLIdentityProviderResponseBody
}

type GetExternalSAMLIdentityProviderResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExternalSAMLIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExternalSAMLIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExternalSAMLIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *GetExternalSAMLIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExternalSAMLIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExternalSAMLIdentityProviderResponse) GetBody() *GetExternalSAMLIdentityProviderResponseBody {
	return s.Body
}

func (s *GetExternalSAMLIdentityProviderResponse) SetHeaders(v map[string]*string) *GetExternalSAMLIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponse) SetStatusCode(v int32) *GetExternalSAMLIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponse) SetBody(v *GetExternalSAMLIdentityProviderResponseBody) *GetExternalSAMLIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *GetExternalSAMLIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

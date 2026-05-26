// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSAMLIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSAMLIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSAMLIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetSAMLIdentityProviderResponseBody) *GetSAMLIdentityProviderResponse
	GetBody() *GetSAMLIdentityProviderResponseBody
}

type GetSAMLIdentityProviderResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSAMLIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSAMLIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSAMLIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *GetSAMLIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSAMLIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSAMLIdentityProviderResponse) GetBody() *GetSAMLIdentityProviderResponseBody {
	return s.Body
}

func (s *GetSAMLIdentityProviderResponse) SetHeaders(v map[string]*string) *GetSAMLIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *GetSAMLIdentityProviderResponse) SetStatusCode(v int32) *GetSAMLIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSAMLIdentityProviderResponse) SetBody(v *GetSAMLIdentityProviderResponseBody) *GetSAMLIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *GetSAMLIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

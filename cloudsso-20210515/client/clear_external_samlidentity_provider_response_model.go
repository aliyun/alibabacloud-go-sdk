// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearExternalSAMLIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearExternalSAMLIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearExternalSAMLIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *ClearExternalSAMLIdentityProviderResponseBody) *ClearExternalSAMLIdentityProviderResponse
	GetBody() *ClearExternalSAMLIdentityProviderResponseBody
}

type ClearExternalSAMLIdentityProviderResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearExternalSAMLIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearExternalSAMLIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearExternalSAMLIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *ClearExternalSAMLIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearExternalSAMLIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearExternalSAMLIdentityProviderResponse) GetBody() *ClearExternalSAMLIdentityProviderResponseBody {
	return s.Body
}

func (s *ClearExternalSAMLIdentityProviderResponse) SetHeaders(v map[string]*string) *ClearExternalSAMLIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *ClearExternalSAMLIdentityProviderResponse) SetStatusCode(v int32) *ClearExternalSAMLIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearExternalSAMLIdentityProviderResponse) SetBody(v *ClearExternalSAMLIdentityProviderResponseBody) *ClearExternalSAMLIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *ClearExternalSAMLIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

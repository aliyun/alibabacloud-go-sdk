// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderServiceProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentityProviderServiceProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentityProviderServiceProviderResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentityProviderServiceProviderResponseBody) *GetIdentityProviderServiceProviderResponse
	GetBody() *GetIdentityProviderServiceProviderResponseBody
}

type GetIdentityProviderServiceProviderResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentityProviderServiceProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityProviderServiceProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderServiceProviderResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderServiceProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentityProviderServiceProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentityProviderServiceProviderResponse) GetBody() *GetIdentityProviderServiceProviderResponseBody {
	return s.Body
}

func (s *GetIdentityProviderServiceProviderResponse) SetHeaders(v map[string]*string) *GetIdentityProviderServiceProviderResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityProviderServiceProviderResponse) SetStatusCode(v int32) *GetIdentityProviderServiceProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityProviderServiceProviderResponse) SetBody(v *GetIdentityProviderServiceProviderResponseBody) *GetIdentityProviderServiceProviderResponse {
	s.Body = v
	return s
}

func (s *GetIdentityProviderServiceProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

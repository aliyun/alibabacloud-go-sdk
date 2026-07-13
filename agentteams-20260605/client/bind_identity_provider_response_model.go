// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindIdentityProviderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindIdentityProviderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindIdentityProviderResponse
	GetStatusCode() *int32
	SetBody(v *BindIdentityProviderResponseBody) *BindIdentityProviderResponse
	GetBody() *BindIdentityProviderResponseBody
}

type BindIdentityProviderResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindIdentityProviderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindIdentityProviderResponse) String() string {
	return dara.Prettify(s)
}

func (s BindIdentityProviderResponse) GoString() string {
	return s.String()
}

func (s *BindIdentityProviderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindIdentityProviderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindIdentityProviderResponse) GetBody() *BindIdentityProviderResponseBody {
	return s.Body
}

func (s *BindIdentityProviderResponse) SetHeaders(v map[string]*string) *BindIdentityProviderResponse {
	s.Headers = v
	return s
}

func (s *BindIdentityProviderResponse) SetStatusCode(v int32) *BindIdentityProviderResponse {
	s.StatusCode = &v
	return s
}

func (s *BindIdentityProviderResponse) SetBody(v *BindIdentityProviderResponseBody) *BindIdentityProviderResponse {
	s.Body = v
	return s
}

func (s *BindIdentityProviderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

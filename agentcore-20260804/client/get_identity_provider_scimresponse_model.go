// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderSCIMResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentityProviderSCIMResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentityProviderSCIMResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentityProviderSCIMResponseBody) *GetIdentityProviderSCIMResponse
	GetBody() *GetIdentityProviderSCIMResponseBody
}

type GetIdentityProviderSCIMResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentityProviderSCIMResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityProviderSCIMResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSCIMResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSCIMResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentityProviderSCIMResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentityProviderSCIMResponse) GetBody() *GetIdentityProviderSCIMResponseBody {
	return s.Body
}

func (s *GetIdentityProviderSCIMResponse) SetHeaders(v map[string]*string) *GetIdentityProviderSCIMResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityProviderSCIMResponse) SetStatusCode(v int32) *GetIdentityProviderSCIMResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityProviderSCIMResponse) SetBody(v *GetIdentityProviderSCIMResponseBody) *GetIdentityProviderSCIMResponse {
	s.Body = v
	return s
}

func (s *GetIdentityProviderSCIMResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

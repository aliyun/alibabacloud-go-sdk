// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepareIdentityProviderSCIMResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PrepareIdentityProviderSCIMResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PrepareIdentityProviderSCIMResponse
	GetStatusCode() *int32
	SetBody(v *PrepareIdentityProviderSCIMResponseBody) *PrepareIdentityProviderSCIMResponse
	GetBody() *PrepareIdentityProviderSCIMResponseBody
}

type PrepareIdentityProviderSCIMResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrepareIdentityProviderSCIMResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrepareIdentityProviderSCIMResponse) String() string {
	return dara.Prettify(s)
}

func (s PrepareIdentityProviderSCIMResponse) GoString() string {
	return s.String()
}

func (s *PrepareIdentityProviderSCIMResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PrepareIdentityProviderSCIMResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PrepareIdentityProviderSCIMResponse) GetBody() *PrepareIdentityProviderSCIMResponseBody {
	return s.Body
}

func (s *PrepareIdentityProviderSCIMResponse) SetHeaders(v map[string]*string) *PrepareIdentityProviderSCIMResponse {
	s.Headers = v
	return s
}

func (s *PrepareIdentityProviderSCIMResponse) SetStatusCode(v int32) *PrepareIdentityProviderSCIMResponse {
	s.StatusCode = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponse) SetBody(v *PrepareIdentityProviderSCIMResponseBody) *PrepareIdentityProviderSCIMResponse {
	s.Body = v
	return s
}

func (s *PrepareIdentityProviderSCIMResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

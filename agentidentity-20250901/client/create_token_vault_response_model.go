// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTokenVaultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTokenVaultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTokenVaultResponse
	GetStatusCode() *int32
	SetBody(v *CreateTokenVaultResponseBody) *CreateTokenVaultResponse
	GetBody() *CreateTokenVaultResponseBody
}

type CreateTokenVaultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTokenVaultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTokenVaultResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTokenVaultResponse) GoString() string {
	return s.String()
}

func (s *CreateTokenVaultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTokenVaultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTokenVaultResponse) GetBody() *CreateTokenVaultResponseBody {
	return s.Body
}

func (s *CreateTokenVaultResponse) SetHeaders(v map[string]*string) *CreateTokenVaultResponse {
	s.Headers = v
	return s
}

func (s *CreateTokenVaultResponse) SetStatusCode(v int32) *CreateTokenVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTokenVaultResponse) SetBody(v *CreateTokenVaultResponseBody) *CreateTokenVaultResponse {
	s.Body = v
	return s
}

func (s *CreateTokenVaultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

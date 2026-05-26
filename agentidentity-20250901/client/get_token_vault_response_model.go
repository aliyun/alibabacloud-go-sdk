// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenVaultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTokenVaultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTokenVaultResponse
	GetStatusCode() *int32
	SetBody(v *GetTokenVaultResponseBody) *GetTokenVaultResponse
	GetBody() *GetTokenVaultResponseBody
}

type GetTokenVaultResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenVaultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenVaultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTokenVaultResponse) GoString() string {
	return s.String()
}

func (s *GetTokenVaultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTokenVaultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTokenVaultResponse) GetBody() *GetTokenVaultResponseBody {
	return s.Body
}

func (s *GetTokenVaultResponse) SetHeaders(v map[string]*string) *GetTokenVaultResponse {
	s.Headers = v
	return s
}

func (s *GetTokenVaultResponse) SetStatusCode(v int32) *GetTokenVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenVaultResponse) SetBody(v *GetTokenVaultResponseBody) *GetTokenVaultResponse {
	s.Body = v
	return s
}

func (s *GetTokenVaultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

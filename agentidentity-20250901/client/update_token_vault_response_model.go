// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTokenVaultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTokenVaultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTokenVaultResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTokenVaultResponseBody) *UpdateTokenVaultResponse
	GetBody() *UpdateTokenVaultResponseBody
}

type UpdateTokenVaultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTokenVaultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTokenVaultResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTokenVaultResponse) GoString() string {
	return s.String()
}

func (s *UpdateTokenVaultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTokenVaultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTokenVaultResponse) GetBody() *UpdateTokenVaultResponseBody {
	return s.Body
}

func (s *UpdateTokenVaultResponse) SetHeaders(v map[string]*string) *UpdateTokenVaultResponse {
	s.Headers = v
	return s
}

func (s *UpdateTokenVaultResponse) SetStatusCode(v int32) *UpdateTokenVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTokenVaultResponse) SetBody(v *UpdateTokenVaultResponseBody) *UpdateTokenVaultResponse {
	s.Body = v
	return s
}

func (s *UpdateTokenVaultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

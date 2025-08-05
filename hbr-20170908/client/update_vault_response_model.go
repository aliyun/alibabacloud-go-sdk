// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVaultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVaultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVaultResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVaultResponseBody) *UpdateVaultResponse
	GetBody() *UpdateVaultResponseBody
}

type UpdateVaultResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVaultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVaultResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVaultResponse) GoString() string {
	return s.String()
}

func (s *UpdateVaultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVaultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVaultResponse) GetBody() *UpdateVaultResponseBody {
	return s.Body
}

func (s *UpdateVaultResponse) SetHeaders(v map[string]*string) *UpdateVaultResponse {
	s.Headers = v
	return s
}

func (s *UpdateVaultResponse) SetStatusCode(v int32) *UpdateVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVaultResponse) SetBody(v *UpdateVaultResponseBody) *UpdateVaultResponse {
	s.Body = v
	return s
}

func (s *UpdateVaultResponse) Validate() error {
	return dara.Validate(s)
}

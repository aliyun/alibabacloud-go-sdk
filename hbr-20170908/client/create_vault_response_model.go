// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVaultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVaultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVaultResponse
	GetStatusCode() *int32
	SetBody(v *CreateVaultResponseBody) *CreateVaultResponse
	GetBody() *CreateVaultResponseBody
}

type CreateVaultResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVaultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVaultResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVaultResponse) GoString() string {
	return s.String()
}

func (s *CreateVaultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVaultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVaultResponse) GetBody() *CreateVaultResponseBody {
	return s.Body
}

func (s *CreateVaultResponse) SetHeaders(v map[string]*string) *CreateVaultResponse {
	s.Headers = v
	return s
}

func (s *CreateVaultResponse) SetStatusCode(v int32) *CreateVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVaultResponse) SetBody(v *CreateVaultResponseBody) *CreateVaultResponse {
	s.Body = v
	return s
}

func (s *CreateVaultResponse) Validate() error {
	return dara.Validate(s)
}

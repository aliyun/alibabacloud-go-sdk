// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTenantKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTenantKeyResponse
	GetStatusCode() *int32
	SetBody(v *CreateTenantKeyResponseBody) *CreateTenantKeyResponse
	GetBody() *CreateTenantKeyResponseBody
}

type CreateTenantKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTenantKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTenantKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantKeyResponse) GoString() string {
	return s.String()
}

func (s *CreateTenantKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTenantKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTenantKeyResponse) GetBody() *CreateTenantKeyResponseBody {
	return s.Body
}

func (s *CreateTenantKeyResponse) SetHeaders(v map[string]*string) *CreateTenantKeyResponse {
	s.Headers = v
	return s
}

func (s *CreateTenantKeyResponse) SetStatusCode(v int32) *CreateTenantKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTenantKeyResponse) SetBody(v *CreateTenantKeyResponseBody) *CreateTenantKeyResponse {
	s.Body = v
	return s
}

func (s *CreateTenantKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

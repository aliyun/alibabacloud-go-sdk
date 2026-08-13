// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTenantDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTenantDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTenantDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *CreateTenantDirectoryResponseBody) *CreateTenantDirectoryResponse
	GetBody() *CreateTenantDirectoryResponseBody
}

type CreateTenantDirectoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTenantDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTenantDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTenantDirectoryResponse) GoString() string {
	return s.String()
}

func (s *CreateTenantDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTenantDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTenantDirectoryResponse) GetBody() *CreateTenantDirectoryResponseBody {
	return s.Body
}

func (s *CreateTenantDirectoryResponse) SetHeaders(v map[string]*string) *CreateTenantDirectoryResponse {
	s.Headers = v
	return s
}

func (s *CreateTenantDirectoryResponse) SetStatusCode(v int32) *CreateTenantDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTenantDirectoryResponse) SetBody(v *CreateTenantDirectoryResponseBody) *CreateTenantDirectoryResponse {
	s.Body = v
	return s
}

func (s *CreateTenantDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

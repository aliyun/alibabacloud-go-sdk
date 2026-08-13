// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTenantDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTenantDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTenantDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTenantDirectoryResponseBody) *UpdateTenantDirectoryResponse
	GetBody() *UpdateTenantDirectoryResponseBody
}

type UpdateTenantDirectoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTenantDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTenantDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTenantDirectoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateTenantDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTenantDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTenantDirectoryResponse) GetBody() *UpdateTenantDirectoryResponseBody {
	return s.Body
}

func (s *UpdateTenantDirectoryResponse) SetHeaders(v map[string]*string) *UpdateTenantDirectoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateTenantDirectoryResponse) SetStatusCode(v int32) *UpdateTenantDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTenantDirectoryResponse) SetBody(v *UpdateTenantDirectoryResponseBody) *UpdateTenantDirectoryResponse {
	s.Body = v
	return s
}

func (s *UpdateTenantDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

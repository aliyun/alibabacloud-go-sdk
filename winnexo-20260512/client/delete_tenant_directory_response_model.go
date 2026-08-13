// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTenantDirectoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTenantDirectoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTenantDirectoryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTenantDirectoryResponseBody) *DeleteTenantDirectoryResponse
	GetBody() *DeleteTenantDirectoryResponseBody
}

type DeleteTenantDirectoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTenantDirectoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTenantDirectoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTenantDirectoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteTenantDirectoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTenantDirectoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTenantDirectoryResponse) GetBody() *DeleteTenantDirectoryResponseBody {
	return s.Body
}

func (s *DeleteTenantDirectoryResponse) SetHeaders(v map[string]*string) *DeleteTenantDirectoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteTenantDirectoryResponse) SetStatusCode(v int32) *DeleteTenantDirectoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTenantDirectoryResponse) SetBody(v *DeleteTenantDirectoryResponseBody) *DeleteTenantDirectoryResponse {
	s.Body = v
	return s
}

func (s *DeleteTenantDirectoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

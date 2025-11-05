// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSchemaPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeSchemaPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeSchemaPermissionResponse
	GetStatusCode() *int32
	SetBody(v *RevokeSchemaPermissionResponseBody) *RevokeSchemaPermissionResponse
	GetBody() *RevokeSchemaPermissionResponseBody
}

type RevokeSchemaPermissionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeSchemaPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeSchemaPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeSchemaPermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokeSchemaPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeSchemaPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeSchemaPermissionResponse) GetBody() *RevokeSchemaPermissionResponseBody {
	return s.Body
}

func (s *RevokeSchemaPermissionResponse) SetHeaders(v map[string]*string) *RevokeSchemaPermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokeSchemaPermissionResponse) SetStatusCode(v int32) *RevokeSchemaPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeSchemaPermissionResponse) SetBody(v *RevokeSchemaPermissionResponseBody) *RevokeSchemaPermissionResponse {
	s.Body = v
	return s
}

func (s *RevokeSchemaPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantSchemaPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantSchemaPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantSchemaPermissionResponse
	GetStatusCode() *int32
	SetBody(v *GrantSchemaPermissionResponseBody) *GrantSchemaPermissionResponse
	GetBody() *GrantSchemaPermissionResponseBody
}

type GrantSchemaPermissionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantSchemaPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantSchemaPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantSchemaPermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantSchemaPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantSchemaPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantSchemaPermissionResponse) GetBody() *GrantSchemaPermissionResponseBody {
	return s.Body
}

func (s *GrantSchemaPermissionResponse) SetHeaders(v map[string]*string) *GrantSchemaPermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantSchemaPermissionResponse) SetStatusCode(v int32) *GrantSchemaPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantSchemaPermissionResponse) SetBody(v *GrantSchemaPermissionResponseBody) *GrantSchemaPermissionResponse {
	s.Body = v
	return s
}

func (s *GrantSchemaPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

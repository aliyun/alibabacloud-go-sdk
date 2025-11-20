// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePermissionResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePermissionResponseBody) *UpdatePermissionResponse
	GetBody() *UpdatePermissionResponseBody
}

type UpdatePermissionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdatePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePermissionResponse) GetBody() *UpdatePermissionResponseBody {
	return s.Body
}

func (s *UpdatePermissionResponse) SetHeaders(v map[string]*string) *UpdatePermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdatePermissionResponse) SetStatusCode(v int32) *UpdatePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePermissionResponse) SetBody(v *UpdatePermissionResponseBody) *UpdatePermissionResponse {
	s.Body = v
	return s
}

func (s *UpdatePermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

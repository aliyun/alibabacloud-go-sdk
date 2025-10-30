// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRowPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRowPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRowPermissionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRowPermissionResponseBody) *UpdateRowPermissionResponse
	GetBody() *UpdateRowPermissionResponseBody
}

type UpdateRowPermissionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRowPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRowPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRowPermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdateRowPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRowPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRowPermissionResponse) GetBody() *UpdateRowPermissionResponseBody {
	return s.Body
}

func (s *UpdateRowPermissionResponse) SetHeaders(v map[string]*string) *UpdateRowPermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdateRowPermissionResponse) SetStatusCode(v int32) *UpdateRowPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRowPermissionResponse) SetBody(v *UpdateRowPermissionResponseBody) *UpdateRowPermissionResponse {
	s.Body = v
	return s
}

func (s *UpdateRowPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

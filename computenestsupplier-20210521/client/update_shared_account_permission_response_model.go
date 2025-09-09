// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSharedAccountPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSharedAccountPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSharedAccountPermissionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSharedAccountPermissionResponseBody) *UpdateSharedAccountPermissionResponse
	GetBody() *UpdateSharedAccountPermissionResponseBody
}

type UpdateSharedAccountPermissionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSharedAccountPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSharedAccountPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSharedAccountPermissionResponse) GoString() string {
	return s.String()
}

func (s *UpdateSharedAccountPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSharedAccountPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSharedAccountPermissionResponse) GetBody() *UpdateSharedAccountPermissionResponseBody {
	return s.Body
}

func (s *UpdateSharedAccountPermissionResponse) SetHeaders(v map[string]*string) *UpdateSharedAccountPermissionResponse {
	s.Headers = v
	return s
}

func (s *UpdateSharedAccountPermissionResponse) SetStatusCode(v int32) *UpdateSharedAccountPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSharedAccountPermissionResponse) SetBody(v *UpdateSharedAccountPermissionResponseBody) *UpdateSharedAccountPermissionResponse {
	s.Body = v
	return s
}

func (s *UpdateSharedAccountPermissionResponse) Validate() error {
	return dara.Validate(s)
}

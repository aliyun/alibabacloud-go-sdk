// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeUserPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeUserPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeUserPermissionResponse
	GetStatusCode() *int32
	SetBody(v *RevokeUserPermissionResponseBody) *RevokeUserPermissionResponse
	GetBody() *RevokeUserPermissionResponseBody
}

type RevokeUserPermissionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeUserPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeUserPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeUserPermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokeUserPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeUserPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeUserPermissionResponse) GetBody() *RevokeUserPermissionResponseBody {
	return s.Body
}

func (s *RevokeUserPermissionResponse) SetHeaders(v map[string]*string) *RevokeUserPermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokeUserPermissionResponse) SetStatusCode(v int32) *RevokeUserPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeUserPermissionResponse) SetBody(v *RevokeUserPermissionResponseBody) *RevokeUserPermissionResponse {
	s.Body = v
	return s
}

func (s *RevokeUserPermissionResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokePermissionResponse
	GetStatusCode() *int32
	SetBody(v *RevokePermissionResponseBody) *RevokePermissionResponse
	GetBody() *RevokePermissionResponseBody
}

type RevokePermissionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokePermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokePermissionResponse) GetBody() *RevokePermissionResponseBody {
	return s.Body
}

func (s *RevokePermissionResponse) SetHeaders(v map[string]*string) *RevokePermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokePermissionResponse) SetStatusCode(v int32) *RevokePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokePermissionResponse) SetBody(v *RevokePermissionResponseBody) *RevokePermissionResponse {
	s.Body = v
	return s
}

func (s *RevokePermissionResponse) Validate() error {
	return dara.Validate(s)
}

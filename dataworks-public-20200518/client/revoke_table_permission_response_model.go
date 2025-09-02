// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeTablePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeTablePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeTablePermissionResponse
	GetStatusCode() *int32
	SetBody(v *RevokeTablePermissionResponseBody) *RevokeTablePermissionResponse
	GetBody() *RevokeTablePermissionResponseBody
}

type RevokeTablePermissionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeTablePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeTablePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeTablePermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokeTablePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeTablePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeTablePermissionResponse) GetBody() *RevokeTablePermissionResponseBody {
	return s.Body
}

func (s *RevokeTablePermissionResponse) SetHeaders(v map[string]*string) *RevokeTablePermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokeTablePermissionResponse) SetStatusCode(v int32) *RevokeTablePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeTablePermissionResponse) SetBody(v *RevokeTablePermissionResponseBody) *RevokeTablePermissionResponse {
	s.Body = v
	return s
}

func (s *RevokeTablePermissionResponse) Validate() error {
	return dara.Validate(s)
}

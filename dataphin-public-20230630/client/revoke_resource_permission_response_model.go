// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourcePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeResourcePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeResourcePermissionResponse
	GetStatusCode() *int32
	SetBody(v *RevokeResourcePermissionResponseBody) *RevokeResourcePermissionResponse
	GetBody() *RevokeResourcePermissionResponseBody
}

type RevokeResourcePermissionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeResourcePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeResourcePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourcePermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeResourcePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeResourcePermissionResponse) GetBody() *RevokeResourcePermissionResponseBody {
	return s.Body
}

func (s *RevokeResourcePermissionResponse) SetHeaders(v map[string]*string) *RevokeResourcePermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokeResourcePermissionResponse) SetStatusCode(v int32) *RevokeResourcePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeResourcePermissionResponse) SetBody(v *RevokeResourcePermissionResponseBody) *RevokeResourcePermissionResponse {
	s.Body = v
	return s
}

func (s *RevokeResourcePermissionResponse) Validate() error {
	return dara.Validate(s)
}

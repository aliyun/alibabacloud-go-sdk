// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeOperatorPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeOperatorPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeOperatorPermissionResponse
	GetStatusCode() *int32
	SetBody(v *RevokeOperatorPermissionResponseBody) *RevokeOperatorPermissionResponse
	GetBody() *RevokeOperatorPermissionResponseBody
}

type RevokeOperatorPermissionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeOperatorPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeOperatorPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeOperatorPermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokeOperatorPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeOperatorPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeOperatorPermissionResponse) GetBody() *RevokeOperatorPermissionResponseBody {
	return s.Body
}

func (s *RevokeOperatorPermissionResponse) SetHeaders(v map[string]*string) *RevokeOperatorPermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokeOperatorPermissionResponse) SetStatusCode(v int32) *RevokeOperatorPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeOperatorPermissionResponse) SetBody(v *RevokeOperatorPermissionResponseBody) *RevokeOperatorPermissionResponse {
	s.Body = v
	return s
}

func (s *RevokeOperatorPermissionResponse) Validate() error {
	return dara.Validate(s)
}

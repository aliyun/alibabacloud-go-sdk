// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMemberDeletionPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetMemberDeletionPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetMemberDeletionPermissionResponse
	GetStatusCode() *int32
	SetBody(v *SetMemberDeletionPermissionResponseBody) *SetMemberDeletionPermissionResponse
	GetBody() *SetMemberDeletionPermissionResponseBody
}

type SetMemberDeletionPermissionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetMemberDeletionPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetMemberDeletionPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetMemberDeletionPermissionResponse) GoString() string {
	return s.String()
}

func (s *SetMemberDeletionPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetMemberDeletionPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetMemberDeletionPermissionResponse) GetBody() *SetMemberDeletionPermissionResponseBody {
	return s.Body
}

func (s *SetMemberDeletionPermissionResponse) SetHeaders(v map[string]*string) *SetMemberDeletionPermissionResponse {
	s.Headers = v
	return s
}

func (s *SetMemberDeletionPermissionResponse) SetStatusCode(v int32) *SetMemberDeletionPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetMemberDeletionPermissionResponse) SetBody(v *SetMemberDeletionPermissionResponseBody) *SetMemberDeletionPermissionResponse {
	s.Body = v
	return s
}

func (s *SetMemberDeletionPermissionResponse) Validate() error {
	return dara.Validate(s)
}

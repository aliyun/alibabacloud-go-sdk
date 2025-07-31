// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantResourcePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantResourcePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantResourcePermissionResponse
	GetStatusCode() *int32
	SetBody(v *GrantResourcePermissionResponseBody) *GrantResourcePermissionResponse
	GetBody() *GrantResourcePermissionResponseBody
}

type GrantResourcePermissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantResourcePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantResourcePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantResourcePermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantResourcePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantResourcePermissionResponse) GetBody() *GrantResourcePermissionResponseBody {
	return s.Body
}

func (s *GrantResourcePermissionResponse) SetHeaders(v map[string]*string) *GrantResourcePermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantResourcePermissionResponse) SetStatusCode(v int32) *GrantResourcePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantResourcePermissionResponse) SetBody(v *GrantResourcePermissionResponseBody) *GrantResourcePermissionResponse {
	s.Body = v
	return s
}

func (s *GrantResourcePermissionResponse) Validate() error {
	return dara.Validate(s)
}

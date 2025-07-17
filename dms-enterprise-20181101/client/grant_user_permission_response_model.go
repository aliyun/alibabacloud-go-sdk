// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantUserPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantUserPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantUserPermissionResponse
	GetStatusCode() *int32
	SetBody(v *GrantUserPermissionResponseBody) *GrantUserPermissionResponse
	GetBody() *GrantUserPermissionResponseBody
}

type GrantUserPermissionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantUserPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantUserPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantUserPermissionResponse) GoString() string {
	return s.String()
}

func (s *GrantUserPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantUserPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantUserPermissionResponse) GetBody() *GrantUserPermissionResponseBody {
	return s.Body
}

func (s *GrantUserPermissionResponse) SetHeaders(v map[string]*string) *GrantUserPermissionResponse {
	s.Headers = v
	return s
}

func (s *GrantUserPermissionResponse) SetStatusCode(v int32) *GrantUserPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantUserPermissionResponse) SetBody(v *GrantUserPermissionResponseBody) *GrantUserPermissionResponse {
	s.Body = v
	return s
}

func (s *GrantUserPermissionResponse) Validate() error {
	return dara.Validate(s)
}

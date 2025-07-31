// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResourcePermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckResourcePermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckResourcePermissionResponse
	GetStatusCode() *int32
	SetBody(v *CheckResourcePermissionResponseBody) *CheckResourcePermissionResponse
	GetBody() *CheckResourcePermissionResponseBody
}

type CheckResourcePermissionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckResourcePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckResourcePermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckResourcePermissionResponse) GoString() string {
	return s.String()
}

func (s *CheckResourcePermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckResourcePermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckResourcePermissionResponse) GetBody() *CheckResourcePermissionResponseBody {
	return s.Body
}

func (s *CheckResourcePermissionResponse) SetHeaders(v map[string]*string) *CheckResourcePermissionResponse {
	s.Headers = v
	return s
}

func (s *CheckResourcePermissionResponse) SetStatusCode(v int32) *CheckResourcePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckResourcePermissionResponse) SetBody(v *CheckResourcePermissionResponseBody) *CheckResourcePermissionResponse {
	s.Body = v
	return s
}

func (s *CheckResourcePermissionResponse) Validate() error {
	return dara.Validate(s)
}

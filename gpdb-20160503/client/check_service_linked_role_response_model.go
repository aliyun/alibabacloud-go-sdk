// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceLinkedRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckServiceLinkedRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckServiceLinkedRoleResponse
	GetStatusCode() *int32
	SetBody(v *CheckServiceLinkedRoleResponseBody) *CheckServiceLinkedRoleResponse
	GetBody() *CheckServiceLinkedRoleResponseBody
}

type CheckServiceLinkedRoleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckServiceLinkedRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckServiceLinkedRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckServiceLinkedRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckServiceLinkedRoleResponse) GetBody() *CheckServiceLinkedRoleResponseBody {
	return s.Body
}

func (s *CheckServiceLinkedRoleResponse) SetHeaders(v map[string]*string) *CheckServiceLinkedRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceLinkedRoleResponse) SetStatusCode(v int32) *CheckServiceLinkedRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceLinkedRoleResponse) SetBody(v *CheckServiceLinkedRoleResponseBody) *CheckServiceLinkedRoleResponse {
	s.Body = v
	return s
}

func (s *CheckServiceLinkedRoleResponse) Validate() error {
	return dara.Validate(s)
}

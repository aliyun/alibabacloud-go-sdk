// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckServiceRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckServiceRoleResponse
	GetStatusCode() *int32
	SetBody(v *CheckServiceRoleResponseBody) *CheckServiceRoleResponse
	GetBody() *CheckServiceRoleResponseBody
}

type CheckServiceRoleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckServiceRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckServiceRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckServiceRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckServiceRoleResponse) GetBody() *CheckServiceRoleResponseBody {
	return s.Body
}

func (s *CheckServiceRoleResponse) SetHeaders(v map[string]*string) *CheckServiceRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckServiceRoleResponse) SetStatusCode(v int32) *CheckServiceRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckServiceRoleResponse) SetBody(v *CheckServiceRoleResponseBody) *CheckServiceRoleResponse {
	s.Body = v
	return s
}

func (s *CheckServiceRoleResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeOperationRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitializeOperationRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitializeOperationRoleResponse
	GetStatusCode() *int32
	SetBody(v *InitializeOperationRoleResponseBody) *InitializeOperationRoleResponse
	GetBody() *InitializeOperationRoleResponseBody
}

type InitializeOperationRoleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitializeOperationRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitializeOperationRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s InitializeOperationRoleResponse) GoString() string {
	return s.String()
}

func (s *InitializeOperationRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitializeOperationRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitializeOperationRoleResponse) GetBody() *InitializeOperationRoleResponseBody {
	return s.Body
}

func (s *InitializeOperationRoleResponse) SetHeaders(v map[string]*string) *InitializeOperationRoleResponse {
	s.Headers = v
	return s
}

func (s *InitializeOperationRoleResponse) SetStatusCode(v int32) *InitializeOperationRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *InitializeOperationRoleResponse) SetBody(v *InitializeOperationRoleResponseBody) *InitializeOperationRoleResponse {
	s.Body = v
	return s
}

func (s *InitializeOperationRoleResponse) Validate() error {
	return dara.Validate(s)
}

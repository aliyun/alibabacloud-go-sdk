// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignRoleResponse
	GetStatusCode() *int32
}

type AssignRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AssignRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignRoleResponse) GoString() string {
	return s.String()
}

func (s *AssignRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignRoleResponse) SetHeaders(v map[string]*string) *AssignRoleResponse {
	s.Headers = v
	return s
}

func (s *AssignRoleResponse) SetStatusCode(v int32) *AssignRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignRoleResponse) Validate() error {
	return dara.Validate(s)
}

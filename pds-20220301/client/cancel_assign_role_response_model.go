// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAssignRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelAssignRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelAssignRoleResponse
	GetStatusCode() *int32
}

type CancelAssignRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CancelAssignRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelAssignRoleResponse) GoString() string {
	return s.String()
}

func (s *CancelAssignRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelAssignRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelAssignRoleResponse) SetHeaders(v map[string]*string) *CancelAssignRoleResponse {
	s.Headers = v
	return s
}

func (s *CancelAssignRoleResponse) SetStatusCode(v int32) *CancelAssignRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelAssignRoleResponse) Validate() error {
	return dara.Validate(s)
}

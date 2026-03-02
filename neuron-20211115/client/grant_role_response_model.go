// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantRoleResponse
	GetStatusCode() *int32
}

type GrantRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GrantRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantRoleResponse) GoString() string {
	return s.String()
}

func (s *GrantRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantRoleResponse) SetHeaders(v map[string]*string) *GrantRoleResponse {
	s.Headers = v
	return s
}

func (s *GrantRoleResponse) SetStatusCode(v int32) *GrantRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantRoleResponse) Validate() error {
	return dara.Validate(s)
}

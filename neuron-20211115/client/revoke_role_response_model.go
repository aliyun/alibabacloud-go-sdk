// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeRoleResponse
	GetStatusCode() *int32
}

type RevokeRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RevokeRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeRoleResponse) GoString() string {
	return s.String()
}

func (s *RevokeRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeRoleResponse) SetHeaders(v map[string]*string) *RevokeRoleResponse {
	s.Headers = v
	return s
}

func (s *RevokeRoleResponse) SetStatusCode(v int32) *RevokeRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeRoleResponse) Validate() error {
	return dara.Validate(s)
}

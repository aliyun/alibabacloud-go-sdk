// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGrantedRolesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGrantedRolesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGrantedRolesResponse
	GetStatusCode() *int32
	SetBody(v *RolePageResult) *ListGrantedRolesResponse
	GetBody() *RolePageResult
}

type ListGrantedRolesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RolePageResult    `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGrantedRolesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGrantedRolesResponse) GoString() string {
	return s.String()
}

func (s *ListGrantedRolesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGrantedRolesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGrantedRolesResponse) GetBody() *RolePageResult {
	return s.Body
}

func (s *ListGrantedRolesResponse) SetHeaders(v map[string]*string) *ListGrantedRolesResponse {
	s.Headers = v
	return s
}

func (s *ListGrantedRolesResponse) SetStatusCode(v int32) *ListGrantedRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGrantedRolesResponse) SetBody(v *RolePageResult) *ListGrantedRolesResponse {
	s.Body = v
	return s
}

func (s *ListGrantedRolesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

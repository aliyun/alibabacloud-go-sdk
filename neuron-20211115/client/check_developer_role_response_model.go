// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDeveloperRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDeveloperRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDeveloperRoleResponse
	GetStatusCode() *int32
	SetBody(v *PermissionResult) *CheckDeveloperRoleResponse
	GetBody() *PermissionResult
}

type CheckDeveloperRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PermissionResult  `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDeveloperRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDeveloperRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckDeveloperRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDeveloperRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDeveloperRoleResponse) GetBody() *PermissionResult {
	return s.Body
}

func (s *CheckDeveloperRoleResponse) SetHeaders(v map[string]*string) *CheckDeveloperRoleResponse {
	s.Headers = v
	return s
}

func (s *CheckDeveloperRoleResponse) SetStatusCode(v int32) *CheckDeveloperRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDeveloperRoleResponse) SetBody(v *PermissionResult) *CheckDeveloperRoleResponse {
	s.Body = v
	return s
}

func (s *CheckDeveloperRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

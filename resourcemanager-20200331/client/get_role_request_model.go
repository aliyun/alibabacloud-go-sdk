// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *GetRoleRequest
	GetLanguage() *string
	SetRoleName(v string) *GetRoleRequest
	GetRoleName() *string
}

type GetRoleRequest struct {
	// The language type. The role description is returned in the specified language. Valid values:
	//
	// - en: English.
	//
	// - zh-CN: Chinese.
	//
	// - ja: Japanese.
	//
	// example:
	//
	// zh-CN
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The role name.
	//
	// The name must be 1 to 64 characters in length and can contain letters, digits, periods (.), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// ECSAdmin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoleRequest) GoString() string {
	return s.String()
}

func (s *GetRoleRequest) GetLanguage() *string {
	return s.Language
}

func (s *GetRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *GetRoleRequest) SetLanguage(v string) *GetRoleRequest {
	s.Language = &v
	return s
}

func (s *GetRoleRequest) SetRoleName(v string) *GetRoleRequest {
	s.RoleName = &v
	return s
}

func (s *GetRoleRequest) Validate() error {
	return dara.Validate(s)
}

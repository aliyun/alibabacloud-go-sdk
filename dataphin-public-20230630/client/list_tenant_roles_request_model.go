// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTenantRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ListTenantRolesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListTenantRolesRequest
	GetOpUserId() *string
}

type ListTenantRolesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListTenantRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTenantRolesRequest) GoString() string {
	return s.String()
}

func (s *ListTenantRolesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListTenantRolesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListTenantRolesRequest) SetOpTenantId(v int64) *ListTenantRolesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListTenantRolesRequest) SetOpUserId(v string) *ListTenantRolesRequest {
	s.OpUserId = &v
	return s
}

func (s *ListTenantRolesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCountTodosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *CountTodosRequest
	GetCollaborationGroupId() *string
	SetTenantId(v string) *CountTodosRequest
	GetTenantId() *string
}

type CountTodosRequest struct {
	// The collaboration workspace ID. This parameter is optional and used for filtering. If this parameter is not specified, all workspaces are included.
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The tenant ID. This is a common parameter. If this parameter is not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CountTodosRequest) String() string {
	return dara.Prettify(s)
}

func (s CountTodosRequest) GoString() string {
	return s.String()
}

func (s *CountTodosRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *CountTodosRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CountTodosRequest) SetCollaborationGroupId(v string) *CountTodosRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *CountTodosRequest) SetTenantId(v string) *CountTodosRequest {
	s.TenantId = &v
	return s
}

func (s *CountTodosRequest) Validate() error {
	return dara.Validate(s)
}

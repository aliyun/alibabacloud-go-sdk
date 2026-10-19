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
	// 协作空间 ID（可选过滤；不传 = 全部空间）
	//
	// example:
	//
	// cg_123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkItemPinnedRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *SetWorkItemPinnedRequest
	GetCollaborationGroupId() *string
	SetPinned(v bool) *SetWorkItemPinnedRequest
	GetPinned() *bool
	SetTenantId(v string) *SetWorkItemPinnedRequest
	GetTenantId() *string
	SetWorkItemId(v string) *SetWorkItemPinnedRequest
	GetWorkItemId() *string
}

type SetWorkItemPinnedRequest struct {
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// 置顶状态
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Pinned *bool `json:"pinned,omitempty" xml:"pinned,omitempty"`
	// 租户ID，公共参数；缺省时使用调用方默认租户
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 工作项 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s SetWorkItemPinnedRequest) String() string {
	return dara.Prettify(s)
}

func (s SetWorkItemPinnedRequest) GoString() string {
	return s.String()
}

func (s *SetWorkItemPinnedRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *SetWorkItemPinnedRequest) GetPinned() *bool {
	return s.Pinned
}

func (s *SetWorkItemPinnedRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SetWorkItemPinnedRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *SetWorkItemPinnedRequest) SetCollaborationGroupId(v string) *SetWorkItemPinnedRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *SetWorkItemPinnedRequest) SetPinned(v bool) *SetWorkItemPinnedRequest {
	s.Pinned = &v
	return s
}

func (s *SetWorkItemPinnedRequest) SetTenantId(v string) *SetWorkItemPinnedRequest {
	s.TenantId = &v
	return s
}

func (s *SetWorkItemPinnedRequest) SetWorkItemId(v string) *SetWorkItemPinnedRequest {
	s.WorkItemId = &v
	return s
}

func (s *SetWorkItemPinnedRequest) Validate() error {
	return dara.Validate(s)
}

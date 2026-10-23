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
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The pin status.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	Pinned *bool `json:"pinned,omitempty" xml:"pinned,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The ID of the work item.
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

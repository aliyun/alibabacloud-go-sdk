// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemArtifactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ListWorkItemArtifactsRequest
	GetCollaborationGroupId() *string
	SetTenantId(v string) *ListWorkItemArtifactsRequest
	GetTenantId() *string
	SetWorkItemId(v string) *ListWorkItemArtifactsRequest
	GetWorkItemId() *string
}

type ListWorkItemArtifactsRequest struct {
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The work item ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s ListWorkItemArtifactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemArtifactsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkItemArtifactsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkItemArtifactsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkItemArtifactsRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ListWorkItemArtifactsRequest) SetCollaborationGroupId(v string) *ListWorkItemArtifactsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkItemArtifactsRequest) SetTenantId(v string) *ListWorkItemArtifactsRequest {
	s.TenantId = &v
	return s
}

func (s *ListWorkItemArtifactsRequest) SetWorkItemId(v string) *ListWorkItemArtifactsRequest {
	s.WorkItemId = &v
	return s
}

func (s *ListWorkItemArtifactsRequest) Validate() error {
	return dara.Validate(s)
}

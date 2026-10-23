// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResolveWorkItemArtifactAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *ResolveWorkItemArtifactAccessRequest
	GetArtifactId() *string
	SetCollaborationGroupId(v string) *ResolveWorkItemArtifactAccessRequest
	GetCollaborationGroupId() *string
	SetTenantId(v string) *ResolveWorkItemArtifactAccessRequest
	GetTenantId() *string
	SetWorkItemId(v string) *ResolveWorkItemArtifactAccessRequest
	GetWorkItemId() *string
}

type ResolveWorkItemArtifactAccessRequest struct {
	// The ID of the artifact.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	ArtifactId *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	// The ID of the collaboration workspace.
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
	// The ID of the work item.
	//
	// This parameter is required.
	//
	// example:
	//
	// wi_example
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s ResolveWorkItemArtifactAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s ResolveWorkItemArtifactAccessRequest) GoString() string {
	return s.String()
}

func (s *ResolveWorkItemArtifactAccessRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *ResolveWorkItemArtifactAccessRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ResolveWorkItemArtifactAccessRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ResolveWorkItemArtifactAccessRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ResolveWorkItemArtifactAccessRequest) SetArtifactId(v string) *ResolveWorkItemArtifactAccessRequest {
	s.ArtifactId = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessRequest) SetCollaborationGroupId(v string) *ResolveWorkItemArtifactAccessRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessRequest) SetTenantId(v string) *ResolveWorkItemArtifactAccessRequest {
	s.TenantId = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessRequest) SetWorkItemId(v string) *ResolveWorkItemArtifactAccessRequest {
	s.WorkItemId = &v
	return s
}

func (s *ResolveWorkItemArtifactAccessRequest) Validate() error {
	return dara.Validate(s)
}

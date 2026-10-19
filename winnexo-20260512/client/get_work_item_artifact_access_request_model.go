// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemArtifactAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactId(v string) *GetWorkItemArtifactAccessRequest
	GetArtifactId() *string
	SetCollaborationGroupId(v string) *GetWorkItemArtifactAccessRequest
	GetCollaborationGroupId() *string
	SetTenantId(v string) *GetWorkItemArtifactAccessRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetWorkItemArtifactAccessRequest
	GetWorkItemId() *string
}

type GetWorkItemArtifactAccessRequest struct {
	// 产物 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	ArtifactId *string `json:"artifactId,omitempty" xml:"artifactId,omitempty"`
	// 协作空间 ID
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
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

func (s GetWorkItemArtifactAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemArtifactAccessRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemArtifactAccessRequest) GetArtifactId() *string {
	return s.ArtifactId
}

func (s *GetWorkItemArtifactAccessRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkItemArtifactAccessRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkItemArtifactAccessRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkItemArtifactAccessRequest) SetArtifactId(v string) *GetWorkItemArtifactAccessRequest {
	s.ArtifactId = &v
	return s
}

func (s *GetWorkItemArtifactAccessRequest) SetCollaborationGroupId(v string) *GetWorkItemArtifactAccessRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkItemArtifactAccessRequest) SetTenantId(v string) *GetWorkItemArtifactAccessRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkItemArtifactAccessRequest) SetWorkItemId(v string) *GetWorkItemArtifactAccessRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkItemArtifactAccessRequest) Validate() error {
	return dara.Validate(s)
}

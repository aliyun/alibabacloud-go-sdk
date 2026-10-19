// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeWorkItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *ResumeWorkItemRequest
	GetCollaborationGroupId() *string
	SetTenantId(v string) *ResumeWorkItemRequest
	GetTenantId() *string
	SetWorkItemId(v string) *ResumeWorkItemRequest
	GetWorkItemId() *string
}

type ResumeWorkItemRequest struct {
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

func (s ResumeWorkItemRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeWorkItemRequest) GoString() string {
	return s.String()
}

func (s *ResumeWorkItemRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ResumeWorkItemRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ResumeWorkItemRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *ResumeWorkItemRequest) SetCollaborationGroupId(v string) *ResumeWorkItemRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ResumeWorkItemRequest) SetTenantId(v string) *ResumeWorkItemRequest {
	s.TenantId = &v
	return s
}

func (s *ResumeWorkItemRequest) SetWorkItemId(v string) *ResumeWorkItemRequest {
	s.WorkItemId = &v
	return s
}

func (s *ResumeWorkItemRequest) Validate() error {
	return dara.Validate(s)
}

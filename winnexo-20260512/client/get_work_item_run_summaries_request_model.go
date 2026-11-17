// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemRunSummariesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCollaborationGroupId(v string) *GetWorkItemRunSummariesRequest
	GetCollaborationGroupId() *string
	SetNodeKey(v string) *GetWorkItemRunSummariesRequest
	GetNodeKey() *string
	SetTenantId(v string) *GetWorkItemRunSummariesRequest
	GetTenantId() *string
	SetWorkItemId(v string) *GetWorkItemRunSummariesRequest
	GetWorkItemId() *string
}

type GetWorkItemRunSummariesRequest struct {
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The node key.
	//
	// example:
	//
	// query.amount
	NodeKey *string `json:"nodeKey,omitempty" xml:"nodeKey,omitempty"`
	// The tenant ID. If not specified, the default tenant of the caller is used.
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
	// wi_0123456789ab4cde
	WorkItemId *string `json:"workItemId,omitempty" xml:"workItemId,omitempty"`
}

func (s GetWorkItemRunSummariesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemRunSummariesRequest) GoString() string {
	return s.String()
}

func (s *GetWorkItemRunSummariesRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *GetWorkItemRunSummariesRequest) GetNodeKey() *string {
	return s.NodeKey
}

func (s *GetWorkItemRunSummariesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkItemRunSummariesRequest) GetWorkItemId() *string {
	return s.WorkItemId
}

func (s *GetWorkItemRunSummariesRequest) SetCollaborationGroupId(v string) *GetWorkItemRunSummariesRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *GetWorkItemRunSummariesRequest) SetNodeKey(v string) *GetWorkItemRunSummariesRequest {
	s.NodeKey = &v
	return s
}

func (s *GetWorkItemRunSummariesRequest) SetTenantId(v string) *GetWorkItemRunSummariesRequest {
	s.TenantId = &v
	return s
}

func (s *GetWorkItemRunSummariesRequest) SetWorkItemId(v string) *GetWorkItemRunSummariesRequest {
	s.WorkItemId = &v
	return s
}

func (s *GetWorkItemRunSummariesRequest) Validate() error {
	return dara.Validate(s)
}

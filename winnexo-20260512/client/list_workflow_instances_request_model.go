// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessId(v string) *ListWorkflowInstancesRequest
	GetBusinessId() *string
	SetCollaborationGroupId(v string) *ListWorkflowInstancesRequest
	GetCollaborationGroupId() *string
	SetInitiator(v string) *ListWorkflowInstancesRequest
	GetInitiator() *string
	SetKeyword(v string) *ListWorkflowInstancesRequest
	GetKeyword() *string
	SetPage(v int64) *ListWorkflowInstancesRequest
	GetPage() *int64
	SetPageSize(v int64) *ListWorkflowInstancesRequest
	GetPageSize() *int64
	SetSortBy(v string) *ListWorkflowInstancesRequest
	GetSortBy() *string
	SetSortOrder(v string) *ListWorkflowInstancesRequest
	GetSortOrder() *string
	SetStatus(v string) *ListWorkflowInstancesRequest
	GetStatus() *string
	SetTenantId(v string) *ListWorkflowInstancesRequest
	GetTenantId() *string
}

type ListWorkflowInstancesRequest struct {
	// The exact requirement ID.
	//
	// example:
	//
	// REQ-1001
	BusinessId *string `json:"businessId,omitempty" xml:"businessId,omitempty"`
	// The collaboration space ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_20261016_abc123
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The exact external initiator tag.
	//
	// example:
	//
	// External initiator Zhang San
	Initiator *string `json:"initiator,omitempty" xml:"initiator,omitempty"`
	// The keyword for the title or requirement ID.
	//
	// example:
	//
	// Quarterly report
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The sort field.
	//
	// example:
	//
	// GMT_CREATE
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The sort order.
	//
	// example:
	//
	// DESC
	SortOrder *string `json:"sortOrder,omitempty" xml:"sortOrder,omitempty"`
	// Filters the instances by status. Valid values: Running, Completed, Failed, and Terminated.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tenant ID. This is a common parameter. If you use winnexo-cli, pass this parameter explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListWorkflowInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkflowInstancesRequest) GetBusinessId() *string {
	return s.BusinessId
}

func (s *ListWorkflowInstancesRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkflowInstancesRequest) GetInitiator() *string {
	return s.Initiator
}

func (s *ListWorkflowInstancesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListWorkflowInstancesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListWorkflowInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWorkflowInstancesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListWorkflowInstancesRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListWorkflowInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListWorkflowInstancesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkflowInstancesRequest) SetBusinessId(v string) *ListWorkflowInstancesRequest {
	s.BusinessId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetCollaborationGroupId(v string) *ListWorkflowInstancesRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetInitiator(v string) *ListWorkflowInstancesRequest {
	s.Initiator = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetKeyword(v string) *ListWorkflowInstancesRequest {
	s.Keyword = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetPage(v int64) *ListWorkflowInstancesRequest {
	s.Page = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetPageSize(v int64) *ListWorkflowInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetSortBy(v string) *ListWorkflowInstancesRequest {
	s.SortBy = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetSortOrder(v string) *ListWorkflowInstancesRequest {
	s.SortOrder = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetStatus(v string) *ListWorkflowInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListWorkflowInstancesRequest) SetTenantId(v string) *ListWorkflowInstancesRequest {
	s.TenantId = &v
	return s
}

func (s *ListWorkflowInstancesRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ListWorkItemsRequest
	GetAppToken() *string
	SetBusinessStatuses(v []*string) *ListWorkItemsRequest
	GetBusinessStatuses() []*string
	SetCollaborationGroupId(v string) *ListWorkItemsRequest
	GetCollaborationGroupId() *string
	SetItemTypes(v []*string) *ListWorkItemsRequest
	GetItemTypes() []*string
	SetKeyword(v string) *ListWorkItemsRequest
	GetKeyword() *string
	SetMaxResults(v int64) *ListWorkItemsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListWorkItemsRequest
	GetNextToken() *string
	SetOpsUserCode(v string) *ListWorkItemsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ListWorkItemsRequest
	GetOpsUserId() *string
	SetPage(v int64) *ListWorkItemsRequest
	GetPage() *int64
	SetPageSize(v int64) *ListWorkItemsRequest
	GetPageSize() *int64
	SetPriorities(v []*string) *ListWorkItemsRequest
	GetPriorities() []*string
	SetScope(v string) *ListWorkItemsRequest
	GetScope() *string
	SetTenantId(v string) *ListWorkItemsRequest
	GetTenantId() *string
}

type ListWorkItemsRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The business status filter, which supports seven states.
	//
	// example:
	//
	// ["PROCESSING"]
	BusinessStatuses []*string `json:"businessStatuses,omitempty" xml:"businessStatuses,omitempty" type:"Repeated"`
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group_example
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The work item type filter.
	//
	// example:
	//
	// ["WORKFLOW"]
	ItemTypes []*string `json:"itemTypes,omitempty" xml:"itemTypes,omitempty" type:"Repeated"`
	// The keyword for fuzzy search in the title or description.
	//
	// example:
	//
	// example
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of entries to return per page. Valid values: 1 to 100. This parameter takes precedence over pageSize if both are specified.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token, which is obtained from the nextToken value returned in the previous response. This parameter takes precedence over page if both are specified.
	//
	// example:
	//
	// example
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The page number. This parameter is ignored if nextToken is specified.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The priority filter.
	//
	// example:
	//
	// ["HIGH"]
	Priorities []*string `json:"priorities,omitempty" xml:"priorities,omitempty" type:"Repeated"`
	// The personal scope filter.
	//
	// example:
	//
	// ASSIGNED_TO_ME
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListWorkItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWorkItemsRequest) GoString() string {
	return s.String()
}

func (s *ListWorkItemsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ListWorkItemsRequest) GetBusinessStatuses() []*string {
	return s.BusinessStatuses
}

func (s *ListWorkItemsRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ListWorkItemsRequest) GetItemTypes() []*string {
	return s.ItemTypes
}

func (s *ListWorkItemsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListWorkItemsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListWorkItemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkItemsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ListWorkItemsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ListWorkItemsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListWorkItemsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWorkItemsRequest) GetPriorities() []*string {
	return s.Priorities
}

func (s *ListWorkItemsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListWorkItemsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListWorkItemsRequest) SetAppToken(v string) *ListWorkItemsRequest {
	s.AppToken = &v
	return s
}

func (s *ListWorkItemsRequest) SetBusinessStatuses(v []*string) *ListWorkItemsRequest {
	s.BusinessStatuses = v
	return s
}

func (s *ListWorkItemsRequest) SetCollaborationGroupId(v string) *ListWorkItemsRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ListWorkItemsRequest) SetItemTypes(v []*string) *ListWorkItemsRequest {
	s.ItemTypes = v
	return s
}

func (s *ListWorkItemsRequest) SetKeyword(v string) *ListWorkItemsRequest {
	s.Keyword = &v
	return s
}

func (s *ListWorkItemsRequest) SetMaxResults(v int64) *ListWorkItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListWorkItemsRequest) SetNextToken(v string) *ListWorkItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListWorkItemsRequest) SetOpsUserCode(v string) *ListWorkItemsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ListWorkItemsRequest) SetOpsUserId(v string) *ListWorkItemsRequest {
	s.OpsUserId = &v
	return s
}

func (s *ListWorkItemsRequest) SetPage(v int64) *ListWorkItemsRequest {
	s.Page = &v
	return s
}

func (s *ListWorkItemsRequest) SetPageSize(v int64) *ListWorkItemsRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkItemsRequest) SetPriorities(v []*string) *ListWorkItemsRequest {
	s.Priorities = v
	return s
}

func (s *ListWorkItemsRequest) SetScope(v string) *ListWorkItemsRequest {
	s.Scope = &v
	return s
}

func (s *ListWorkItemsRequest) SetTenantId(v string) *ListWorkItemsRequest {
	s.TenantId = &v
	return s
}

func (s *ListWorkItemsRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHasResource(v string) *ListQuotasRequest
	GetHasResource() *string
	SetLabels(v string) *ListQuotasRequest
	GetLabels() *string
	SetLayoutMode(v string) *ListQuotasRequest
	GetLayoutMode() *string
	SetOrder(v string) *ListQuotasRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListQuotasRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListQuotasRequest
	GetPageSize() *int32
	SetParentQuotaId(v string) *ListQuotasRequest
	GetParentQuotaId() *string
	SetQuotaIds(v string) *ListQuotasRequest
	GetQuotaIds() *string
	SetQuotaName(v string) *ListQuotasRequest
	GetQuotaName() *string
	SetResourceType(v string) *ListQuotasRequest
	GetResourceType() *string
	SetSortBy(v string) *ListQuotasRequest
	GetSortBy() *string
	SetStatuses(v string) *ListQuotasRequest
	GetStatuses() *string
	SetVerbose(v bool) *ListQuotasRequest
	GetVerbose() *bool
	SetVersions(v string) *ListQuotasRequest
	GetVersions() *string
	SetWorkspaceIds(v string) *ListQuotasRequest
	GetWorkspaceIds() *string
	SetWorkspaceName(v string) *ListQuotasRequest
	GetWorkspaceName() *string
}

type ListQuotasRequest struct {
	HasResource *string `json:"HasResource,omitempty" xml:"HasResource,omitempty"`
	// example:
	//
	// official=true,gpu=false
	Labels     *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	LayoutMode *string `json:"LayoutMode,omitempty" xml:"LayoutMode,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// quotajradxh43rgb
	ParentQuotaId *string `json:"ParentQuotaId,omitempty" xml:"ParentQuotaId,omitempty"`
	// example:
	//
	// quota1ci8g793pgm,quotajradxh43rgb
	QuotaIds *string `json:"QuotaIds,omitempty" xml:"QuotaIds,omitempty"`
	// example:
	//
	// quotajradxh43rgb
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// status
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// Creating
	Statuses *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
	Verbose  *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// example:
	//
	// 1.0
	Versions *string `json:"Versions,omitempty" xml:"Versions,omitempty"`
	// example:
	//
	// 21345,38727
	WorkspaceIds  *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s ListQuotasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListQuotasRequest) GetHasResource() *string {
	return s.HasResource
}

func (s *ListQuotasRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListQuotasRequest) GetLayoutMode() *string {
	return s.LayoutMode
}

func (s *ListQuotasRequest) GetOrder() *string {
	return s.Order
}

func (s *ListQuotasRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListQuotasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQuotasRequest) GetParentQuotaId() *string {
	return s.ParentQuotaId
}

func (s *ListQuotasRequest) GetQuotaIds() *string {
	return s.QuotaIds
}

func (s *ListQuotasRequest) GetQuotaName() *string {
	return s.QuotaName
}

func (s *ListQuotasRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListQuotasRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListQuotasRequest) GetStatuses() *string {
	return s.Statuses
}

func (s *ListQuotasRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListQuotasRequest) GetVersions() *string {
	return s.Versions
}

func (s *ListQuotasRequest) GetWorkspaceIds() *string {
	return s.WorkspaceIds
}

func (s *ListQuotasRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *ListQuotasRequest) SetHasResource(v string) *ListQuotasRequest {
	s.HasResource = &v
	return s
}

func (s *ListQuotasRequest) SetLabels(v string) *ListQuotasRequest {
	s.Labels = &v
	return s
}

func (s *ListQuotasRequest) SetLayoutMode(v string) *ListQuotasRequest {
	s.LayoutMode = &v
	return s
}

func (s *ListQuotasRequest) SetOrder(v string) *ListQuotasRequest {
	s.Order = &v
	return s
}

func (s *ListQuotasRequest) SetPageNumber(v int32) *ListQuotasRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQuotasRequest) SetPageSize(v int32) *ListQuotasRequest {
	s.PageSize = &v
	return s
}

func (s *ListQuotasRequest) SetParentQuotaId(v string) *ListQuotasRequest {
	s.ParentQuotaId = &v
	return s
}

func (s *ListQuotasRequest) SetQuotaIds(v string) *ListQuotasRequest {
	s.QuotaIds = &v
	return s
}

func (s *ListQuotasRequest) SetQuotaName(v string) *ListQuotasRequest {
	s.QuotaName = &v
	return s
}

func (s *ListQuotasRequest) SetResourceType(v string) *ListQuotasRequest {
	s.ResourceType = &v
	return s
}

func (s *ListQuotasRequest) SetSortBy(v string) *ListQuotasRequest {
	s.SortBy = &v
	return s
}

func (s *ListQuotasRequest) SetStatuses(v string) *ListQuotasRequest {
	s.Statuses = &v
	return s
}

func (s *ListQuotasRequest) SetVerbose(v bool) *ListQuotasRequest {
	s.Verbose = &v
	return s
}

func (s *ListQuotasRequest) SetVersions(v string) *ListQuotasRequest {
	s.Versions = &v
	return s
}

func (s *ListQuotasRequest) SetWorkspaceIds(v string) *ListQuotasRequest {
	s.WorkspaceIds = &v
	return s
}

func (s *ListQuotasRequest) SetWorkspaceName(v string) *ListQuotasRequest {
	s.WorkspaceName = &v
	return s
}

func (s *ListQuotasRequest) Validate() error {
	return dara.Validate(s)
}

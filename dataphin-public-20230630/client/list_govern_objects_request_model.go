// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGovernObjectsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListGovernObjectsRequestListQuery) *ListGovernObjectsRequest
	GetListQuery() *ListGovernObjectsRequestListQuery
	SetOpTenantId(v int64) *ListGovernObjectsRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListGovernObjectsRequest
	GetOpUserId() *string
}

type ListGovernObjectsRequest struct {
	// The paged query conditions.
	//
	// This parameter is required.
	ListQuery *ListGovernObjectsRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListGovernObjectsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGovernObjectsRequest) GoString() string {
	return s.String()
}

func (s *ListGovernObjectsRequest) GetListQuery() *ListGovernObjectsRequestListQuery {
	return s.ListQuery
}

func (s *ListGovernObjectsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListGovernObjectsRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListGovernObjectsRequest) SetListQuery(v *ListGovernObjectsRequestListQuery) *ListGovernObjectsRequest {
	s.ListQuery = v
	return s
}

func (s *ListGovernObjectsRequest) SetOpTenantId(v int64) *ListGovernObjectsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListGovernObjectsRequest) SetOpUserId(v string) *ListGovernObjectsRequest {
	s.OpUserId = &v
	return s
}

func (s *ListGovernObjectsRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGovernObjectsRequestListQuery struct {
	// The governance item type. Valid values:
	//
	// - TABLE
	//
	// - DATASOURCE_TABLE
	//
	// - DATASOURCE
	//
	// - INDEX
	//
	// - REALTIME_LOGICAL_TABLE
	//
	// - QD_FEATURE
	//
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	GovernItemType *string `json:"GovernItemType,omitempty" xml:"GovernItemType,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// table_name
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The owner.
	//
	// example:
	//
	// user123
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of project names used to filter results.
	ProjectNames []*string `json:"ProjectNames,omitempty" xml:"ProjectNames,omitempty" type:"Repeated"`
	// The list of governance object statuses. Valid values:
	//
	// - NEW
	//
	// - VERIFY
	//
	// - FINISHED
	//
	// - IGNORE
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	// The view type. Valid values:
	//
	// - ALL
	//
	// - OWNER
	//
	// - PROJECT
	//
	// example:
	//
	// ALL
	ViewType *string `json:"ViewType,omitempty" xml:"ViewType,omitempty"`
}

func (s ListGovernObjectsRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListGovernObjectsRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListGovernObjectsRequestListQuery) GetGovernItemType() *string {
	return s.GovernItemType
}

func (s *ListGovernObjectsRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListGovernObjectsRequestListQuery) GetOwner() *string {
	return s.Owner
}

func (s *ListGovernObjectsRequestListQuery) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGovernObjectsRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGovernObjectsRequestListQuery) GetProjectNames() []*string {
	return s.ProjectNames
}

func (s *ListGovernObjectsRequestListQuery) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListGovernObjectsRequestListQuery) GetViewType() *string {
	return s.ViewType
}

func (s *ListGovernObjectsRequestListQuery) SetGovernItemType(v string) *ListGovernObjectsRequestListQuery {
	s.GovernItemType = &v
	return s
}

func (s *ListGovernObjectsRequestListQuery) SetKeyword(v string) *ListGovernObjectsRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListGovernObjectsRequestListQuery) SetOwner(v string) *ListGovernObjectsRequestListQuery {
	s.Owner = &v
	return s
}

func (s *ListGovernObjectsRequestListQuery) SetPageNumber(v int32) *ListGovernObjectsRequestListQuery {
	s.PageNumber = &v
	return s
}

func (s *ListGovernObjectsRequestListQuery) SetPageSize(v int32) *ListGovernObjectsRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListGovernObjectsRequestListQuery) SetProjectNames(v []*string) *ListGovernObjectsRequestListQuery {
	s.ProjectNames = v
	return s
}

func (s *ListGovernObjectsRequestListQuery) SetStatusList(v []*string) *ListGovernObjectsRequestListQuery {
	s.StatusList = v
	return s
}

func (s *ListGovernObjectsRequestListQuery) SetViewType(v string) *ListGovernObjectsRequestListQuery {
	s.ViewType = &v
	return s
}

func (s *ListGovernObjectsRequestListQuery) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceWithConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListDataSourceWithConfigRequestListQuery) *ListDataSourceWithConfigRequest
	GetListQuery() *ListDataSourceWithConfigRequestListQuery
	SetOpTenantId(v int64) *ListDataSourceWithConfigRequest
	GetOpTenantId() *int64
}

type ListDataSourceWithConfigRequest struct {
	// Paginated query
	//
	// This parameter is required.
	ListQuery *ListDataSourceWithConfigRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// Tenant ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListDataSourceWithConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceWithConfigRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigRequest) GetListQuery() *ListDataSourceWithConfigRequestListQuery {
	return s.ListQuery
}

func (s *ListDataSourceWithConfigRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListDataSourceWithConfigRequest) SetListQuery(v *ListDataSourceWithConfigRequestListQuery) *ListDataSourceWithConfigRequest {
	s.ListQuery = v
	return s
}

func (s *ListDataSourceWithConfigRequest) SetOpTenantId(v int64) *ListDataSourceWithConfigRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListDataSourceWithConfigRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataSourceWithConfigRequestListQuery struct {
	// Data source name
	//
	// example:
	//
	// vcns-test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Data source owner
	OwnerList []*string `json:"OwnerList,omitempty" xml:"OwnerList,omitempty" type:"Repeated"`
	// Page number. The value starts from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// Number of records per page
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Data source scope. Valid values:
	//
	// - STREAMING: Real-time
	//
	// - OFFLINE: Offline
	//
	// - DIP_META_SYNC: Meta warehouse initialization
	//
	// - DATA_DISTILL: Data distillation
	//
	// - CUMPUTE_SOURCE_SHADOW: Shadow compute source
	//
	// - LABEL: Label
	//
	// - ALL: Offline + Real-time
	ScopeList []*string `json:"ScopeList,omitempty" xml:"ScopeList,omitempty" type:"Repeated"`
	// Tag marked when creating the data source
	//
	// example:
	//
	// xx测试
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// Data source type
	TypeList []*string `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s ListDataSourceWithConfigRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceWithConfigRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListDataSourceWithConfigRequestListQuery) GetName() *string {
	return s.Name
}

func (s *ListDataSourceWithConfigRequestListQuery) GetOwnerList() []*string {
	return s.OwnerList
}

func (s *ListDataSourceWithConfigRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListDataSourceWithConfigRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataSourceWithConfigRequestListQuery) GetScopeList() []*string {
	return s.ScopeList
}

func (s *ListDataSourceWithConfigRequestListQuery) GetTag() *string {
	return s.Tag
}

func (s *ListDataSourceWithConfigRequestListQuery) GetTypeList() []*string {
	return s.TypeList
}

func (s *ListDataSourceWithConfigRequestListQuery) SetName(v string) *ListDataSourceWithConfigRequestListQuery {
	s.Name = &v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetOwnerList(v []*string) *ListDataSourceWithConfigRequestListQuery {
	s.OwnerList = v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetPage(v int32) *ListDataSourceWithConfigRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetPageSize(v int32) *ListDataSourceWithConfigRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetScopeList(v []*string) *ListDataSourceWithConfigRequestListQuery {
	s.ScopeList = v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetTag(v string) *ListDataSourceWithConfigRequestListQuery {
	s.Tag = &v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) SetTypeList(v []*string) *ListDataSourceWithConfigRequestListQuery {
	s.TypeList = v
	return s
}

func (s *ListDataSourceWithConfigRequestListQuery) Validate() error {
	return dara.Validate(s)
}

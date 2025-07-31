// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListBizEntitiesRequestListQuery) *ListBizEntitiesRequest
	GetListQuery() *ListBizEntitiesRequestListQuery
	SetOpTenantId(v int64) *ListBizEntitiesRequest
	GetOpTenantId() *int64
}

type ListBizEntitiesRequest struct {
	// This parameter is required.
	ListQuery *ListBizEntitiesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListBizEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBizEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListBizEntitiesRequest) GetListQuery() *ListBizEntitiesRequestListQuery {
	return s.ListQuery
}

func (s *ListBizEntitiesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListBizEntitiesRequest) SetListQuery(v *ListBizEntitiesRequestListQuery) *ListBizEntitiesRequest {
	s.ListQuery = v
	return s
}

func (s *ListBizEntitiesRequest) SetOpTenantId(v int64) *ListBizEntitiesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListBizEntitiesRequest) Validate() error {
	return dara.Validate(s)
}

type ListBizEntitiesRequestListQuery struct {
	FilterCriteria *ListBizEntitiesRequestListQueryFilterCriteria `json:"FilterCriteria,omitempty" xml:"FilterCriteria,omitempty" type:"Struct"`
	// example:
	//
	// object_
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListBizEntitiesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListBizEntitiesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListBizEntitiesRequestListQuery) GetFilterCriteria() *ListBizEntitiesRequestListQueryFilterCriteria {
	return s.FilterCriteria
}

func (s *ListBizEntitiesRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListBizEntitiesRequestListQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListBizEntitiesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBizEntitiesRequestListQuery) SetFilterCriteria(v *ListBizEntitiesRequestListQueryFilterCriteria) *ListBizEntitiesRequestListQuery {
	s.FilterCriteria = v
	return s
}

func (s *ListBizEntitiesRequestListQuery) SetKeyword(v string) *ListBizEntitiesRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListBizEntitiesRequestListQuery) SetPage(v int32) *ListBizEntitiesRequestListQuery {
	s.Page = &v
	return s
}

func (s *ListBizEntitiesRequestListQuery) SetPageSize(v int32) *ListBizEntitiesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListBizEntitiesRequestListQuery) Validate() error {
	return dara.Validate(s)
}

type ListBizEntitiesRequestListQueryFilterCriteria struct {
	BizUnitIdList      []*int64  `json:"BizUnitIdList,omitempty" xml:"BizUnitIdList,omitempty" type:"Repeated"`
	BizUnitNameList    []*string `json:"BizUnitNameList,omitempty" xml:"BizUnitNameList,omitempty" type:"Repeated"`
	DataDomainIdList   []*int64  `json:"DataDomainIdList,omitempty" xml:"DataDomainIdList,omitempty" type:"Repeated"`
	DataDomainNameList []*string `json:"DataDomainNameList,omitempty" xml:"DataDomainNameList,omitempty" type:"Repeated"`
	HasTableRef        *bool     `json:"HasTableRef,omitempty" xml:"HasTableRef,omitempty"`
	OwnerUserIdList    []*string `json:"OwnerUserIdList,omitempty" xml:"OwnerUserIdList,omitempty" type:"Repeated"`
	StatusList         []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	SubTypeList        []*string `json:"SubTypeList,omitempty" xml:"SubTypeList,omitempty" type:"Repeated"`
}

func (s ListBizEntitiesRequestListQueryFilterCriteria) String() string {
	return dara.Prettify(s)
}

func (s ListBizEntitiesRequestListQueryFilterCriteria) GoString() string {
	return s.String()
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) GetBizUnitIdList() []*int64 {
	return s.BizUnitIdList
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) GetBizUnitNameList() []*string {
	return s.BizUnitNameList
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) GetDataDomainIdList() []*int64 {
	return s.DataDomainIdList
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) GetDataDomainNameList() []*string {
	return s.DataDomainNameList
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) GetHasTableRef() *bool {
	return s.HasTableRef
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) GetOwnerUserIdList() []*string {
	return s.OwnerUserIdList
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) GetSubTypeList() []*string {
	return s.SubTypeList
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) SetBizUnitIdList(v []*int64) *ListBizEntitiesRequestListQueryFilterCriteria {
	s.BizUnitIdList = v
	return s
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) SetBizUnitNameList(v []*string) *ListBizEntitiesRequestListQueryFilterCriteria {
	s.BizUnitNameList = v
	return s
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) SetDataDomainIdList(v []*int64) *ListBizEntitiesRequestListQueryFilterCriteria {
	s.DataDomainIdList = v
	return s
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) SetDataDomainNameList(v []*string) *ListBizEntitiesRequestListQueryFilterCriteria {
	s.DataDomainNameList = v
	return s
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) SetHasTableRef(v bool) *ListBizEntitiesRequestListQueryFilterCriteria {
	s.HasTableRef = &v
	return s
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) SetOwnerUserIdList(v []*string) *ListBizEntitiesRequestListQueryFilterCriteria {
	s.OwnerUserIdList = v
	return s
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) SetStatusList(v []*string) *ListBizEntitiesRequestListQueryFilterCriteria {
	s.StatusList = v
	return s
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) SetSubTypeList(v []*string) *ListBizEntitiesRequestListQueryFilterCriteria {
	s.SubTypeList = v
	return s
}

func (s *ListBizEntitiesRequestListQueryFilterCriteria) Validate() error {
	return dara.Validate(s)
}

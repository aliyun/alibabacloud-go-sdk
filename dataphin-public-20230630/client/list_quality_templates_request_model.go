// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQualityTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListQualityTemplatesRequestListQuery) *ListQualityTemplatesRequest
	GetListQuery() *ListQualityTemplatesRequestListQuery
	SetOpTenantId(v int64) *ListQualityTemplatesRequest
	GetOpTenantId() *int64
}

type ListQualityTemplatesRequest struct {
	ListQuery *ListQualityTemplatesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListQualityTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQualityTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListQualityTemplatesRequest) GetListQuery() *ListQualityTemplatesRequestListQuery {
	return s.ListQuery
}

func (s *ListQualityTemplatesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListQualityTemplatesRequest) SetListQuery(v *ListQualityTemplatesRequestListQuery) *ListQualityTemplatesRequest {
	s.ListQuery = v
	return s
}

func (s *ListQualityTemplatesRequest) SetOpTenantId(v int64) *ListQualityTemplatesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListQualityTemplatesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListQualityTemplatesRequestListQuery struct {
	CatalogList      []*string `json:"CatalogList,omitempty" xml:"CatalogList,omitempty" type:"Repeated"`
	CurrentUserOwned *bool     `json:"CurrentUserOwned,omitempty" xml:"CurrentUserOwned,omitempty"`
	// example:
	//
	// abc
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 20
	PageSize                  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SupportDataSourceTypeList []*string `json:"SupportDataSourceTypeList,omitempty" xml:"SupportDataSourceTypeList,omitempty" type:"Repeated"`
	TemplateOwnerList         []*string `json:"TemplateOwnerList,omitempty" xml:"TemplateOwnerList,omitempty" type:"Repeated"`
	TemplateSourceList        []*string `json:"TemplateSourceList,omitempty" xml:"TemplateSourceList,omitempty" type:"Repeated"`
	TemplateTypeList          []*string `json:"TemplateTypeList,omitempty" xml:"TemplateTypeList,omitempty" type:"Repeated"`
	WatchTypeList             []*string `json:"WatchTypeList,omitempty" xml:"WatchTypeList,omitempty" type:"Repeated"`
}

func (s ListQualityTemplatesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListQualityTemplatesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListQualityTemplatesRequestListQuery) GetCatalogList() []*string {
	return s.CatalogList
}

func (s *ListQualityTemplatesRequestListQuery) GetCurrentUserOwned() *bool {
	return s.CurrentUserOwned
}

func (s *ListQualityTemplatesRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListQualityTemplatesRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListQualityTemplatesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListQualityTemplatesRequestListQuery) GetSupportDataSourceTypeList() []*string {
	return s.SupportDataSourceTypeList
}

func (s *ListQualityTemplatesRequestListQuery) GetTemplateOwnerList() []*string {
	return s.TemplateOwnerList
}

func (s *ListQualityTemplatesRequestListQuery) GetTemplateSourceList() []*string {
	return s.TemplateSourceList
}

func (s *ListQualityTemplatesRequestListQuery) GetTemplateTypeList() []*string {
	return s.TemplateTypeList
}

func (s *ListQualityTemplatesRequestListQuery) GetWatchTypeList() []*string {
	return s.WatchTypeList
}

func (s *ListQualityTemplatesRequestListQuery) SetCatalogList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.CatalogList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetCurrentUserOwned(v bool) *ListQualityTemplatesRequestListQuery {
	s.CurrentUserOwned = &v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetKeyword(v string) *ListQualityTemplatesRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetPageNo(v int32) *ListQualityTemplatesRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetPageSize(v int32) *ListQualityTemplatesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetSupportDataSourceTypeList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.SupportDataSourceTypeList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetTemplateOwnerList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.TemplateOwnerList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetTemplateSourceList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.TemplateSourceList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetTemplateTypeList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.TemplateTypeList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) SetWatchTypeList(v []*string) *ListQualityTemplatesRequestListQuery {
	s.WatchTypeList = v
	return s
}

func (s *ListQualityTemplatesRequestListQuery) Validate() error {
	return dara.Validate(s)
}

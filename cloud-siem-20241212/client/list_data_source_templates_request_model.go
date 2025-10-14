// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceTemplateIds(v []*string) *ListDataSourceTemplatesRequest
	GetDataSourceTemplateIds() []*string
	SetLang(v string) *ListDataSourceTemplatesRequest
	GetLang() *string
	SetPageNumber(v string) *ListDataSourceTemplatesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDataSourceTemplatesRequest
	GetPageSize() *string
	SetRegionId(v string) *ListDataSourceTemplatesRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataSourceTemplatesRequest
	GetRoleFor() *int64
}

type ListDataSourceTemplatesRequest struct {
	DataSourceTemplateIds []*string `json:"DataSourceTemplateIds,omitempty" xml:"DataSourceTemplateIds,omitempty" type:"Repeated"`
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1。
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10。
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListDataSourceTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceTemplatesRequest) GetDataSourceTemplateIds() []*string {
	return s.DataSourceTemplateIds
}

func (s *ListDataSourceTemplatesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataSourceTemplatesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataSourceTemplatesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataSourceTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSourceTemplatesRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataSourceTemplatesRequest) SetDataSourceTemplateIds(v []*string) *ListDataSourceTemplatesRequest {
	s.DataSourceTemplateIds = v
	return s
}

func (s *ListDataSourceTemplatesRequest) SetLang(v string) *ListDataSourceTemplatesRequest {
	s.Lang = &v
	return s
}

func (s *ListDataSourceTemplatesRequest) SetPageNumber(v string) *ListDataSourceTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourceTemplatesRequest) SetPageSize(v string) *ListDataSourceTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourceTemplatesRequest) SetRegionId(v string) *ListDataSourceTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSourceTemplatesRequest) SetRoleFor(v int64) *ListDataSourceTemplatesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataSourceTemplatesRequest) Validate() error {
	return dara.Validate(s)
}

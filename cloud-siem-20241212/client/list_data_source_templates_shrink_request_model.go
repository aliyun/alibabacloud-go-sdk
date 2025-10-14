// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourceTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceTemplateIdsShrink(v string) *ListDataSourceTemplatesShrinkRequest
	GetDataSourceTemplateIdsShrink() *string
	SetLang(v string) *ListDataSourceTemplatesShrinkRequest
	GetLang() *string
	SetPageNumber(v string) *ListDataSourceTemplatesShrinkRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDataSourceTemplatesShrinkRequest
	GetPageSize() *string
	SetRegionId(v string) *ListDataSourceTemplatesShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDataSourceTemplatesShrinkRequest
	GetRoleFor() *int64
}

type ListDataSourceTemplatesShrinkRequest struct {
	DataSourceTemplateIdsShrink *string `json:"DataSourceTemplateIds,omitempty" xml:"DataSourceTemplateIds,omitempty"`
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

func (s ListDataSourceTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourceTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceTemplatesShrinkRequest) GetDataSourceTemplateIdsShrink() *string {
	return s.DataSourceTemplateIdsShrink
}

func (s *ListDataSourceTemplatesShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDataSourceTemplatesShrinkRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataSourceTemplatesShrinkRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataSourceTemplatesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataSourceTemplatesShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDataSourceTemplatesShrinkRequest) SetDataSourceTemplateIdsShrink(v string) *ListDataSourceTemplatesShrinkRequest {
	s.DataSourceTemplateIdsShrink = &v
	return s
}

func (s *ListDataSourceTemplatesShrinkRequest) SetLang(v string) *ListDataSourceTemplatesShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListDataSourceTemplatesShrinkRequest) SetPageNumber(v string) *ListDataSourceTemplatesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourceTemplatesShrinkRequest) SetPageSize(v string) *ListDataSourceTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourceTemplatesShrinkRequest) SetRegionId(v string) *ListDataSourceTemplatesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataSourceTemplatesShrinkRequest) SetRoleFor(v int64) *ListDataSourceTemplatesShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDataSourceTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

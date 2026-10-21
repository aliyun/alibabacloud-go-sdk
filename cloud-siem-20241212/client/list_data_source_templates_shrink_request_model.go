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
	// The list of data source template IDs.
	DataSourceTemplateIdsShrink *string `json:"DataSourceTemplateIds,omitempty" xml:"DataSourceTemplateIds,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// This parameter is required in practice. If you do not specify this parameter, the API returns a 400 error (The specified Lang can not be empty.). Valid values: zh and en.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Chinese mainland. Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Outside China. Your assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 173326*******
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

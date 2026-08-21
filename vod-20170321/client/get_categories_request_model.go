// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCategoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCateId(v int64) *GetCategoriesRequest
	GetCateId() *int64
	SetPageNo(v int64) *GetCategoriesRequest
	GetPageNo() *int64
	SetPageSize(v int64) *GetCategoriesRequest
	GetPageSize() *int64
	SetSortBy(v string) *GetCategoriesRequest
	GetSortBy() *string
	SetType(v string) *GetCategoriesRequest
	GetType() *string
}

type GetCategoriesRequest struct {
	// The category ID. If you specify this parameter, the information about the specified category is returned. Only a single category ID is supported. You can obtain the category ID by using the following methods:
	//
	// - Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com) and choose **Configuration Management*	- > **Media Asset Management Configuration*	- > **Category Management*	- to view the category ID.
	//
	// - Obtain the category ID from the response of the [AddCategory](~~AddCategory~~) operation when you create a category.
	//
	// example:
	//
	// 49339****
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The page number of the subcategory list. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page of the subcategory list. Default value: **10**. Maximum value: **100**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The method for sorting the query results. Valid values:
	//
	// - **CreationTime:Desc*	- (default): sorts the results by creation time in descending order.
	//
	// - **CreationTime:Asc**: sorts the results by creation time in ascending order.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The categorization type. If you specify this parameter, a filtered query is performed to return categories of the specified type. Valid values:
	//
	// - **default**: audio, video, and image categorization.
	//
	// - **material**: short video material categorization.
	//
	// example:
	//
	// default
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCategoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCategoriesRequest) GoString() string {
	return s.String()
}

func (s *GetCategoriesRequest) GetCateId() *int64 {
	return s.CateId
}

func (s *GetCategoriesRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetCategoriesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetCategoriesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *GetCategoriesRequest) GetType() *string {
	return s.Type
}

func (s *GetCategoriesRequest) SetCateId(v int64) *GetCategoriesRequest {
	s.CateId = &v
	return s
}

func (s *GetCategoriesRequest) SetPageNo(v int64) *GetCategoriesRequest {
	s.PageNo = &v
	return s
}

func (s *GetCategoriesRequest) SetPageSize(v int64) *GetCategoriesRequest {
	s.PageSize = &v
	return s
}

func (s *GetCategoriesRequest) SetSortBy(v string) *GetCategoriesRequest {
	s.SortBy = &v
	return s
}

func (s *GetCategoriesRequest) SetType(v string) *GetCategoriesRequest {
	s.Type = &v
	return s
}

func (s *GetCategoriesRequest) Validate() error {
	return dara.Validate(s)
}

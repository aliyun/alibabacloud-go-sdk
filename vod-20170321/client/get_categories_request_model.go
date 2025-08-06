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
	// The ID of the category. If you specify this parameter, the system queries the category based on the ID. You can specify only one category ID. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com). Choose **Configuration Management*	- > **Media Management*	- > **Categories**. On the Audio and Video / Image Category or Short Video Material Category tab, view the category ID.
	//
	// 	- Obtain the category ID from the response to the [AddCategory](~~AddCategory~~) operation.
	//
	// example:
	//
	// 49339****
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The number of the page where the subcategories to be returned are listed. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries to return on each page of the subcategory list. Default value: **10**. Maximum value: **100**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting method of the results. Valid values:
	//
	// 	- **CreationTime:Desc*	- (default): The results are sorted in reverse chronological order based on the creation time.
	//
	// 	- **CreationTime:Asc**: The results are sorted in chronological order based on the creation time.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The type of the category. If you specify this parameter, the system queries the category based on the type. Valid values:
	//
	// 	- **default*	- (default): audio, video, and image files
	//
	// 	- **material**: short video materials
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

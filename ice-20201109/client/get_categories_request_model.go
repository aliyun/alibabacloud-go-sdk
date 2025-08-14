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
	// The category ID. You can use one of the following methods to obtain the ID:
	//
	// 	- Log on to the [Intelligent Media Services (IMS) console](https://ims.console.aliyun.com) and choose **Media Asset Management*	- > **Category Management*	- to view the category ID.
	//
	// 	- View the value of CateId returned by the AddCategory operation that you called to create a category.
	//
	// 	- View the value of CateId returned by the GetCategories operation that you called to query a category.
	//
	// example:
	//
	// 33
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The page number. Default value: 1
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Valid values: 10 to 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The sorting rule of results. Valid values:
	//
	// \\- CreationTime:Desc (default): The results are sorted in reverse chronological order based on the creation time.
	//
	// \\- CreationTime:Asc: The results are sorted in chronological order based on the creation time.
	//
	// example:
	//
	// CreationTime:Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The type of the category. Valid values: default and material. A value of default indicates audio, video, and image files. This is the default value. A value of material indicates short video materials.
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

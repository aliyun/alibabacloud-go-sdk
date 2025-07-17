// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryList(v *ListMetaCategoryResponseBodyCategoryList) *ListMetaCategoryResponseBody
	GetCategoryList() *ListMetaCategoryResponseBodyCategoryList
	SetErrorCode(v string) *ListMetaCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListMetaCategoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListMetaCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMetaCategoryResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListMetaCategoryResponseBody
	GetTotalCount() *int64
}

type ListMetaCategoryResponseBody struct {
	CategoryList *ListMetaCategoryResponseBodyCategoryList `json:"CategoryList,omitempty" xml:"CategoryList,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMetaCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaCategoryResponseBody) GetCategoryList() *ListMetaCategoryResponseBodyCategoryList {
	return s.CategoryList
}

func (s *ListMetaCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMetaCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMetaCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetaCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMetaCategoryResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMetaCategoryResponseBody) SetCategoryList(v *ListMetaCategoryResponseBodyCategoryList) *ListMetaCategoryResponseBody {
	s.CategoryList = v
	return s
}

func (s *ListMetaCategoryResponseBody) SetErrorCode(v string) *ListMetaCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMetaCategoryResponseBody) SetErrorMessage(v string) *ListMetaCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMetaCategoryResponseBody) SetRequestId(v string) *ListMetaCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaCategoryResponseBody) SetSuccess(v bool) *ListMetaCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListMetaCategoryResponseBody) SetTotalCount(v int64) *ListMetaCategoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMetaCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMetaCategoryResponseBodyCategoryList struct {
	MetaCategory []*MetaCategory `json:"MetaCategory,omitempty" xml:"MetaCategory,omitempty" type:"Repeated"`
}

func (s ListMetaCategoryResponseBodyCategoryList) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCategoryResponseBodyCategoryList) GoString() string {
	return s.String()
}

func (s *ListMetaCategoryResponseBodyCategoryList) GetMetaCategory() []*MetaCategory {
	return s.MetaCategory
}

func (s *ListMetaCategoryResponseBodyCategoryList) SetMetaCategory(v []*MetaCategory) *ListMetaCategoryResponseBodyCategoryList {
	s.MetaCategory = v
	return s
}

func (s *ListMetaCategoryResponseBodyCategoryList) Validate() error {
	return dara.Validate(s)
}

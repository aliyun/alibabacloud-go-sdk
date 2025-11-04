// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCategoryResponseBody
	GetCode() *string
	SetData(v *ListCategoryResponseBodyData) *ListCategoryResponseBody
	GetData() *ListCategoryResponseBodyData
	SetMessage(v string) *ListCategoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCategoryResponseBody
	GetRequestId() *string
	SetStatus(v string) *ListCategoryResponseBody
	GetStatus() *string
	SetSuccess(v bool) *ListCategoryResponseBody
	GetSuccess() *bool
}

type ListCategoryResponseBody struct {
	// example:
	//
	// success
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCategoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// workspace id is null or invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListCategoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCategoryResponseBody) GetData() *ListCategoryResponseBodyData {
	return s.Data
}

func (s *ListCategoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCategoryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ListCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCategoryResponseBody) SetCode(v string) *ListCategoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListCategoryResponseBody) SetData(v *ListCategoryResponseBodyData) *ListCategoryResponseBody {
	s.Data = v
	return s
}

func (s *ListCategoryResponseBody) SetMessage(v string) *ListCategoryResponseBody {
	s.Message = &v
	return s
}

func (s *ListCategoryResponseBody) SetRequestId(v string) *ListCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCategoryResponseBody) SetStatus(v string) *ListCategoryResponseBody {
	s.Status = &v
	return s
}

func (s *ListCategoryResponseBody) SetSuccess(v bool) *ListCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *ListCategoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCategoryResponseBodyData struct {
	CategoryList []*ListCategoryResponseBodyDataCategoryList `json:"CategoryList,omitempty" xml:"CategoryList,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAALHWGpGoYCcYMxiFfmlhvh7Z4G8jiXR6IjHYd+M9WQVJ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCategoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCategoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCategoryResponseBodyData) GetCategoryList() []*ListCategoryResponseBodyDataCategoryList {
	return s.CategoryList
}

func (s *ListCategoryResponseBodyData) GetHasNext() *bool {
	return s.HasNext
}

func (s *ListCategoryResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCategoryResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCategoryResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCategoryResponseBodyData) SetCategoryList(v []*ListCategoryResponseBodyDataCategoryList) *ListCategoryResponseBodyData {
	s.CategoryList = v
	return s
}

func (s *ListCategoryResponseBodyData) SetHasNext(v bool) *ListCategoryResponseBodyData {
	s.HasNext = &v
	return s
}

func (s *ListCategoryResponseBodyData) SetMaxResults(v int32) *ListCategoryResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListCategoryResponseBodyData) SetNextToken(v string) *ListCategoryResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListCategoryResponseBodyData) SetTotalCount(v int32) *ListCategoryResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCategoryResponseBodyData) Validate() error {
	if s.CategoryList != nil {
		for _, item := range s.CategoryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCategoryResponseBodyDataCategoryList struct {
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	CategoryId   *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// cate_addd11b1b79a74e8bbd675c356a91ee3XXXXXXXX
	ParentCategoryId *string `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
}

func (s ListCategoryResponseBodyDataCategoryList) String() string {
	return dara.Prettify(s)
}

func (s ListCategoryResponseBodyDataCategoryList) GoString() string {
	return s.String()
}

func (s *ListCategoryResponseBodyDataCategoryList) GetCategoryId() *string {
	return s.CategoryId
}

func (s *ListCategoryResponseBodyDataCategoryList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *ListCategoryResponseBodyDataCategoryList) GetCategoryType() *string {
	return s.CategoryType
}

func (s *ListCategoryResponseBodyDataCategoryList) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListCategoryResponseBodyDataCategoryList) GetParentCategoryId() *string {
	return s.ParentCategoryId
}

func (s *ListCategoryResponseBodyDataCategoryList) SetCategoryId(v string) *ListCategoryResponseBodyDataCategoryList {
	s.CategoryId = &v
	return s
}

func (s *ListCategoryResponseBodyDataCategoryList) SetCategoryName(v string) *ListCategoryResponseBodyDataCategoryList {
	s.CategoryName = &v
	return s
}

func (s *ListCategoryResponseBodyDataCategoryList) SetCategoryType(v string) *ListCategoryResponseBodyDataCategoryList {
	s.CategoryType = &v
	return s
}

func (s *ListCategoryResponseBodyDataCategoryList) SetIsDefault(v bool) *ListCategoryResponseBodyDataCategoryList {
	s.IsDefault = &v
	return s
}

func (s *ListCategoryResponseBodyDataCategoryList) SetParentCategoryId(v string) *ListCategoryResponseBodyDataCategoryList {
	s.ParentCategoryId = &v
	return s
}

func (s *ListCategoryResponseBodyDataCategoryList) Validate() error {
	return dara.Validate(s)
}

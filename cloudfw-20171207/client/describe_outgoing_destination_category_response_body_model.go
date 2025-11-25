// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOutgoingDestinationCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeOutgoingDestinationCategoryResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeOutgoingDestinationCategoryResponseBody
	GetTotalCount() *int32
	SetTypeList(v []*DescribeOutgoingDestinationCategoryResponseBodyTypeList) *DescribeOutgoingDestinationCategoryResponseBody
	GetTypeList() []*DescribeOutgoingDestinationCategoryResponseBodyTypeList
}

type DescribeOutgoingDestinationCategoryResponseBody struct {
	// example:
	//
	// C1ED80BC-FFC8-57DB-8151-705DC31****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TypeList   []*DescribeOutgoingDestinationCategoryResponseBodyTypeList `json:"TypeList,omitempty" xml:"TypeList,omitempty" type:"Repeated"`
}

func (s DescribeOutgoingDestinationCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOutgoingDestinationCategoryResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeOutgoingDestinationCategoryResponseBody) GetTypeList() []*DescribeOutgoingDestinationCategoryResponseBodyTypeList {
	return s.TypeList
}

func (s *DescribeOutgoingDestinationCategoryResponseBody) SetRequestId(v string) *DescribeOutgoingDestinationCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponseBody) SetTotalCount(v int32) *DescribeOutgoingDestinationCategoryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponseBody) SetTypeList(v []*DescribeOutgoingDestinationCategoryResponseBodyTypeList) *DescribeOutgoingDestinationCategoryResponseBody {
	s.TypeList = v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponseBody) Validate() error {
	if s.TypeList != nil {
		for _, item := range s.TypeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOutgoingDestinationCategoryResponseBodyTypeList struct {
	CategoryList []*DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList `json:"CategoryList,omitempty" xml:"CategoryList,omitempty" type:"Repeated"`
	TypeDescribe *string                                                                `json:"TypeDescribe,omitempty" xml:"TypeDescribe,omitempty"`
	// example:
	//
	// All
	TypeId   *string `json:"TypeId,omitempty" xml:"TypeId,omitempty"`
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s DescribeOutgoingDestinationCategoryResponseBodyTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationCategoryResponseBodyTypeList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeList) GetCategoryList() []*DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList {
	return s.CategoryList
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeList) GetTypeDescribe() *string {
	return s.TypeDescribe
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeList) GetTypeId() *string {
	return s.TypeId
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeList) GetTypeName() *string {
	return s.TypeName
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeList) SetCategoryList(v []*DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) *DescribeOutgoingDestinationCategoryResponseBodyTypeList {
	s.CategoryList = v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeList) SetTypeDescribe(v string) *DescribeOutgoingDestinationCategoryResponseBodyTypeList {
	s.TypeDescribe = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeList) SetTypeId(v string) *DescribeOutgoingDestinationCategoryResponseBodyTypeList {
	s.TypeId = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeList) SetTypeName(v string) *DescribeOutgoingDestinationCategoryResponseBodyTypeList {
	s.TypeName = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeList) Validate() error {
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

type DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList struct {
	CategoryDescribe *string `json:"CategoryDescribe,omitempty" xml:"CategoryDescribe,omitempty"`
	// example:
	//
	// AliYun
	CategoryId   *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	CategoryName *string `json:"CategoryName,omitempty" xml:"CategoryName,omitempty"`
	// example:
	//
	// Trusted
	ClassId *string `json:"ClassId,omitempty" xml:"ClassId,omitempty"`
}

func (s DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) String() string {
	return dara.Prettify(s)
}

func (s DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) GoString() string {
	return s.String()
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) GetCategoryDescribe() *string {
	return s.CategoryDescribe
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) GetCategoryId() *string {
	return s.CategoryId
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) GetCategoryName() *string {
	return s.CategoryName
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) GetClassId() *string {
	return s.ClassId
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) SetCategoryDescribe(v string) *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList {
	s.CategoryDescribe = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) SetCategoryId(v string) *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList {
	s.CategoryId = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) SetCategoryName(v string) *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList {
	s.CategoryName = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) SetClassId(v string) *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList {
	s.ClassId = &v
	return s
}

func (s *DescribeOutgoingDestinationCategoryResponseBodyTypeListCategoryList) Validate() error {
	return dara.Validate(s)
}

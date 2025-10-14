// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationCategoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNormalizationCategoriesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationCategoriesResponseBody
	GetNextToken() *string
	SetNormalizationCategories(v []*ListNormalizationCategoriesResponseBodyNormalizationCategories) *ListNormalizationCategoriesResponseBody
	GetNormalizationCategories() []*ListNormalizationCategoriesResponseBodyNormalizationCategories
	SetRequestId(v string) *ListNormalizationCategoriesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNormalizationCategoriesResponseBody
	GetTotalCount() *int32
}

type ListNormalizationCategoriesResponseBody struct {
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken               *string                                                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NormalizationCategories []*ListNormalizationCategoriesResponseBodyNormalizationCategories `json:"NormalizationCategories,omitempty" xml:"NormalizationCategories,omitempty" type:"Repeated"`
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 57。
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNormalizationCategoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNormalizationCategoriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationCategoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationCategoriesResponseBody) GetNormalizationCategories() []*ListNormalizationCategoriesResponseBodyNormalizationCategories {
	return s.NormalizationCategories
}

func (s *ListNormalizationCategoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNormalizationCategoriesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNormalizationCategoriesResponseBody) SetMaxResults(v int32) *ListNormalizationCategoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationCategoriesResponseBody) SetNextToken(v string) *ListNormalizationCategoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationCategoriesResponseBody) SetNormalizationCategories(v []*ListNormalizationCategoriesResponseBodyNormalizationCategories) *ListNormalizationCategoriesResponseBody {
	s.NormalizationCategories = v
	return s
}

func (s *ListNormalizationCategoriesResponseBody) SetRequestId(v string) *ListNormalizationCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNormalizationCategoriesResponseBody) SetTotalCount(v int32) *ListNormalizationCategoriesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNormalizationCategoriesResponseBody) Validate() error {
	if s.NormalizationCategories != nil {
		for _, item := range s.NormalizationCategories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNormalizationCategoriesResponseBodyNormalizationCategories struct {
	// example:
	//
	// COMMON_CATEGORY。
	NormalizationCategoryId *string `json:"NormalizationCategoryId,omitempty" xml:"NormalizationCategoryId,omitempty"`
	// example:
	//
	// COMMON_CATEGORY。
	NormalizationCategoryName *string `json:"NormalizationCategoryName,omitempty" xml:"NormalizationCategoryName,omitempty"`
}

func (s ListNormalizationCategoriesResponseBodyNormalizationCategories) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationCategoriesResponseBodyNormalizationCategories) GoString() string {
	return s.String()
}

func (s *ListNormalizationCategoriesResponseBodyNormalizationCategories) GetNormalizationCategoryId() *string {
	return s.NormalizationCategoryId
}

func (s *ListNormalizationCategoriesResponseBodyNormalizationCategories) GetNormalizationCategoryName() *string {
	return s.NormalizationCategoryName
}

func (s *ListNormalizationCategoriesResponseBodyNormalizationCategories) SetNormalizationCategoryId(v string) *ListNormalizationCategoriesResponseBodyNormalizationCategories {
	s.NormalizationCategoryId = &v
	return s
}

func (s *ListNormalizationCategoriesResponseBodyNormalizationCategories) SetNormalizationCategoryName(v string) *ListNormalizationCategoriesResponseBodyNormalizationCategories {
	s.NormalizationCategoryName = &v
	return s
}

func (s *ListNormalizationCategoriesResponseBodyNormalizationCategories) Validate() error {
	return dara.Validate(s)
}

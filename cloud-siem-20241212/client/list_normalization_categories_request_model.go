// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationCategoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListNormalizationCategoriesRequest
	GetLang() *string
	SetMaxResults(v int32) *ListNormalizationCategoriesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationCategoriesRequest
	GetNextToken() *string
	SetNormalizationCategoryType(v string) *ListNormalizationCategoriesRequest
	GetNormalizationCategoryType() *string
	SetRegionId(v string) *ListNormalizationCategoriesRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListNormalizationCategoriesRequest
	GetRoleFor() *int64
}

type ListNormalizationCategoriesRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 50。
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****。
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// entity。
	NormalizationCategoryType *string `json:"NormalizationCategoryType,omitempty" xml:"NormalizationCategoryType,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListNormalizationCategoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationCategoriesRequest) GoString() string {
	return s.String()
}

func (s *ListNormalizationCategoriesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListNormalizationCategoriesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationCategoriesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationCategoriesRequest) GetNormalizationCategoryType() *string {
	return s.NormalizationCategoryType
}

func (s *ListNormalizationCategoriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNormalizationCategoriesRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListNormalizationCategoriesRequest) SetLang(v string) *ListNormalizationCategoriesRequest {
	s.Lang = &v
	return s
}

func (s *ListNormalizationCategoriesRequest) SetMaxResults(v int32) *ListNormalizationCategoriesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationCategoriesRequest) SetNextToken(v string) *ListNormalizationCategoriesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationCategoriesRequest) SetNormalizationCategoryType(v string) *ListNormalizationCategoriesRequest {
	s.NormalizationCategoryType = &v
	return s
}

func (s *ListNormalizationCategoriesRequest) SetRegionId(v string) *ListNormalizationCategoriesRequest {
	s.RegionId = &v
	return s
}

func (s *ListNormalizationCategoriesRequest) SetRoleFor(v int64) *ListNormalizationCategoriesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListNormalizationCategoriesRequest) Validate() error {
	return dara.Validate(s)
}

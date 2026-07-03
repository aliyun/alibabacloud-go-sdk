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
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Set this parameter to the NextToken value returned in the previous API call to retrieve the next page of results. You do not need to specify this parameter for the first query.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the normalization rule category. Valid values:
	//
	// - log
	//
	// - entity
	//
	// example:
	//
	// entity
	NormalizationCategoryType *string `json:"NormalizationCategoryType,omitempty" xml:"NormalizationCategoryType,omitempty"`
	// The region of the Data Management center for threat analysis. Select the region for the Data Management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can use this parameter to switch to the perspective of this member.
	//
	// example:
	//
	// 173326*******
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

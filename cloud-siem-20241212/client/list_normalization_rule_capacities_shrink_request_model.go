// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationRuleCapacitiesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListNormalizationRuleCapacitiesShrinkRequest
	GetLang() *string
	SetMaxResults(v int32) *ListNormalizationRuleCapacitiesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationRuleCapacitiesShrinkRequest
	GetNextToken() *string
	SetNormalizationRuleId(v string) *ListNormalizationRuleCapacitiesShrinkRequest
	GetNormalizationRuleId() *string
	SetNormalizationRuleIdsShrink(v string) *ListNormalizationRuleCapacitiesShrinkRequest
	GetNormalizationRuleIdsShrink() *string
	SetRegionId(v string) *ListNormalizationRuleCapacitiesShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListNormalizationRuleCapacitiesShrinkRequest
	GetRoleFor() *int64
}

type ListNormalizationRuleCapacitiesShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// This parameter is required in practice. If you do not specify this parameter or pass an empty string, the API returns a 400 error (The specified Lang can not be empty.).
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return in this request.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. You do not need to specify this parameter for the first query or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The normalization rule ID.
	//
	// example:
	//
	// nr-z0b2ssjteut85uoh9nzp
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// The list of normalization rule IDs.
	NormalizationRuleIdsShrink *string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty"`
	// The region where the threat analysis data management center resides. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: The assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: The assets reside outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator uses to switch to the view of another member.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListNormalizationRuleCapacitiesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRuleCapacitiesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) GetNormalizationRuleIdsShrink() *string {
	return s.NormalizationRuleIdsShrink
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) SetLang(v string) *ListNormalizationRuleCapacitiesShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) SetMaxResults(v int32) *ListNormalizationRuleCapacitiesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) SetNextToken(v string) *ListNormalizationRuleCapacitiesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) SetNormalizationRuleId(v string) *ListNormalizationRuleCapacitiesShrinkRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) SetNormalizationRuleIdsShrink(v string) *ListNormalizationRuleCapacitiesShrinkRequest {
	s.NormalizationRuleIdsShrink = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) SetRegionId(v string) *ListNormalizationRuleCapacitiesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) SetRoleFor(v int64) *ListNormalizationRuleCapacitiesShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

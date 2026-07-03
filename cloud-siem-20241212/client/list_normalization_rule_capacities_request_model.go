// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationRuleCapacitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListNormalizationRuleCapacitiesRequest
	GetLang() *string
	SetMaxResults(v int32) *ListNormalizationRuleCapacitiesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationRuleCapacitiesRequest
	GetNextToken() *string
	SetNormalizationRuleId(v string) *ListNormalizationRuleCapacitiesRequest
	GetNormalizationRuleId() *string
	SetNormalizationRuleIds(v []*string) *ListNormalizationRuleCapacitiesRequest
	GetNormalizationRuleIds() []*string
	SetRegionId(v string) *ListNormalizationRuleCapacitiesRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListNormalizationRuleCapacitiesRequest
	GetRoleFor() *int64
}

type ListNormalizationRuleCapacitiesRequest struct {
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
	// The token to start the next query. Leave this parameter empty for the first query. If the response is truncated, use the \\`NextToken\\` value from the previous response.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the normalization rule.
	//
	// example:
	//
	// nr-z0b2ssjteut85uoh9nzp
	NormalizationRuleId *string `json:"NormalizationRuleId,omitempty" xml:"NormalizationRuleId,omitempty"`
	// A list of normalization rule IDs.
	NormalizationRuleIds []*string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty" type:"Repeated"`
	// The region of the threat analysis Data Management Center. Select the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets are outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that an administrator uses to switch to a member\\"s perspective.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListNormalizationRuleCapacitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRuleCapacitiesRequest) GoString() string {
	return s.String()
}

func (s *ListNormalizationRuleCapacitiesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListNormalizationRuleCapacitiesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationRuleCapacitiesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationRuleCapacitiesRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *ListNormalizationRuleCapacitiesRequest) GetNormalizationRuleIds() []*string {
	return s.NormalizationRuleIds
}

func (s *ListNormalizationRuleCapacitiesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNormalizationRuleCapacitiesRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListNormalizationRuleCapacitiesRequest) SetLang(v string) *ListNormalizationRuleCapacitiesRequest {
	s.Lang = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesRequest) SetMaxResults(v int32) *ListNormalizationRuleCapacitiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesRequest) SetNextToken(v string) *ListNormalizationRuleCapacitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesRequest) SetNormalizationRuleId(v string) *ListNormalizationRuleCapacitiesRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesRequest) SetNormalizationRuleIds(v []*string) *ListNormalizationRuleCapacitiesRequest {
	s.NormalizationRuleIds = v
	return s
}

func (s *ListNormalizationRuleCapacitiesRequest) SetRegionId(v string) *ListNormalizationRuleCapacitiesRequest {
	s.RegionId = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesRequest) SetRoleFor(v int64) *ListNormalizationRuleCapacitiesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListNormalizationRuleCapacitiesRequest) Validate() error {
	return dara.Validate(s)
}

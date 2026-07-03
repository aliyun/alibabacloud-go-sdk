// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNormalizationRuleVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListNormalizationRuleVersionsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListNormalizationRuleVersionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNormalizationRuleVersionsRequest
	GetNextToken() *string
	SetNormalizationRuleId(v string) *ListNormalizationRuleVersionsRequest
	GetNormalizationRuleId() *string
	SetRegionId(v string) *ListNormalizationRuleVersionsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListNormalizationRuleVersionsRequest
	GetRoleFor() *int64
}

type ListNormalizationRuleVersionsRequest struct {
	// The language of the response messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return for the current request.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. You do not need to specify this parameter for the first request. If a next page exists, set this parameter to the value of NextToken that is returned from the last call.
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
	// The region where the Data Management Center for threat analysis is located. Select the region for the Management Center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets are in a region outside Chinese mainland.
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

func (s ListNormalizationRuleVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNormalizationRuleVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListNormalizationRuleVersionsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListNormalizationRuleVersionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNormalizationRuleVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNormalizationRuleVersionsRequest) GetNormalizationRuleId() *string {
	return s.NormalizationRuleId
}

func (s *ListNormalizationRuleVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNormalizationRuleVersionsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListNormalizationRuleVersionsRequest) SetLang(v string) *ListNormalizationRuleVersionsRequest {
	s.Lang = &v
	return s
}

func (s *ListNormalizationRuleVersionsRequest) SetMaxResults(v int32) *ListNormalizationRuleVersionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNormalizationRuleVersionsRequest) SetNextToken(v string) *ListNormalizationRuleVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNormalizationRuleVersionsRequest) SetNormalizationRuleId(v string) *ListNormalizationRuleVersionsRequest {
	s.NormalizationRuleId = &v
	return s
}

func (s *ListNormalizationRuleVersionsRequest) SetRegionId(v string) *ListNormalizationRuleVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListNormalizationRuleVersionsRequest) SetRoleFor(v int64) *ListNormalizationRuleVersionsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListNormalizationRuleVersionsRequest) Validate() error {
	return dara.Validate(s)
}

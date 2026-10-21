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
	NormalizationRuleIds []*string `json:"NormalizationRuleIds,omitempty" xml:"NormalizationRuleIds,omitempty" type:"Repeated"`
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

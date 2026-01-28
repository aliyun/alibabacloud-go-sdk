// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRules(v []*ListAuthorizationRulesForUserResponseBodyAuthorizationRules) *ListAuthorizationRulesForUserResponseBody
	GetAuthorizationRules() []*ListAuthorizationRulesForUserResponseBodyAuthorizationRules
	SetMaxResults(v int32) *ListAuthorizationRulesForUserResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationRulesForUserResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAuthorizationRulesForUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAuthorizationRulesForUserResponseBody
	GetTotalCount() *int64
}

type ListAuthorizationRulesForUserResponseBody struct {
	AuthorizationRules []*ListAuthorizationRulesForUserResponseBodyAuthorizationRules `json:"AuthorizationRules,omitempty" xml:"AuthorizationRules,omitempty" type:"Repeated"`
	// 分页查询时每页行数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthorizationRulesForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForUserResponseBody) GetAuthorizationRules() []*ListAuthorizationRulesForUserResponseBodyAuthorizationRules {
	return s.AuthorizationRules
}

func (s *ListAuthorizationRulesForUserResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationRulesForUserResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationRulesForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizationRulesForUserResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAuthorizationRulesForUserResponseBody) SetAuthorizationRules(v []*ListAuthorizationRulesForUserResponseBodyAuthorizationRules) *ListAuthorizationRulesForUserResponseBody {
	s.AuthorizationRules = v
	return s
}

func (s *ListAuthorizationRulesForUserResponseBody) SetMaxResults(v int32) *ListAuthorizationRulesForUserResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationRulesForUserResponseBody) SetNextToken(v string) *ListAuthorizationRulesForUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationRulesForUserResponseBody) SetRequestId(v string) *ListAuthorizationRulesForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizationRulesForUserResponseBody) SetTotalCount(v int64) *ListAuthorizationRulesForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizationRulesForUserResponseBody) Validate() error {
	if s.AuthorizationRules != nil {
		for _, item := range s.AuthorizationRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorizationRulesForUserResponseBodyAuthorizationRules struct {
	// 授权规则标识。
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 有效周期。
	ValidityPeriod *ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// 有效期类型，枚举值：permanent（永久），time_bound（自定义时间范围）。
	//
	// example:
	//
	// permanent
	ValidityType *string `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s ListAuthorizationRulesForUserResponseBodyAuthorizationRules) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForUserResponseBodyAuthorizationRules) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRules) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRules) GetValidityPeriod() *ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod {
	return s.ValidityPeriod
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRules) GetValidityType() *string {
	return s.ValidityType
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRules) SetAuthorizationRuleId(v string) *ListAuthorizationRulesForUserResponseBodyAuthorizationRules {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRules) SetInstanceId(v string) *ListAuthorizationRulesForUserResponseBodyAuthorizationRules {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRules) SetValidityPeriod(v *ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod) *ListAuthorizationRulesForUserResponseBodyAuthorizationRules {
	s.ValidityPeriod = v
	return s
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRules) SetValidityType(v string) *ListAuthorizationRulesForUserResponseBodyAuthorizationRules {
	s.ValidityType = &v
	return s
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRules) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod struct {
	// 授权生效结束时间。
	//
	// example:
	//
	// 1704042061000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 授权生效开始时间。
	//
	// example:
	//
	// 1704042061000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod) SetEndTime(v int64) *ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod) SetStartTime(v int64) *ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *ListAuthorizationRulesForUserResponseBodyAuthorizationRulesValidityPeriod) Validate() error {
	return dara.Validate(s)
}

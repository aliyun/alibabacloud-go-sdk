// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesForApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRules(v []*ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) *ListAuthorizationRulesForApplicationResponseBody
	GetAuthorizationRules() []*ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules
	SetMaxResults(v int32) *ListAuthorizationRulesForApplicationResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationRulesForApplicationResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAuthorizationRulesForApplicationResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAuthorizationRulesForApplicationResponseBody
	GetTotalCount() *int64
}

type ListAuthorizationRulesForApplicationResponseBody struct {
	AuthorizationRules []*ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules `json:"AuthorizationRules,omitempty" xml:"AuthorizationRules,omitempty" type:"Repeated"`
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

func (s ListAuthorizationRulesForApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForApplicationResponseBody) GetAuthorizationRules() []*ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules {
	return s.AuthorizationRules
}

func (s *ListAuthorizationRulesForApplicationResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationRulesForApplicationResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationRulesForApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizationRulesForApplicationResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAuthorizationRulesForApplicationResponseBody) SetAuthorizationRules(v []*ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) *ListAuthorizationRulesForApplicationResponseBody {
	s.AuthorizationRules = v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponseBody) SetMaxResults(v int32) *ListAuthorizationRulesForApplicationResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponseBody) SetNextToken(v string) *ListAuthorizationRulesForApplicationResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponseBody) SetRequestId(v string) *ListAuthorizationRulesForApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponseBody) SetTotalCount(v int64) *ListAuthorizationRulesForApplicationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponseBody) Validate() error {
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

type ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules struct {
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
	ValidityPeriod *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// 有效期类型，枚举值：permanent（永久），time_bound（自定义时间范围）。
	//
	// example:
	//
	// permanent
	ValidityType *string `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) GetValidityPeriod() *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod {
	return s.ValidityPeriod
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) GetValidityType() *string {
	return s.ValidityType
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) SetAuthorizationRuleId(v string) *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) SetInstanceId(v string) *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) SetValidityPeriod(v *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod) *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules {
	s.ValidityPeriod = v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) SetValidityType(v string) *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules {
	s.ValidityType = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod struct {
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

func (s ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod) SetEndTime(v int64) *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod) SetStartTime(v int64) *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod) Validate() error {
	return dara.Validate(s)
}

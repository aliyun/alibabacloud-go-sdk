// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizationRulesForGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRules(v []*ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) *ListAuthorizationRulesForGroupResponseBody
	GetAuthorizationRules() []*ListAuthorizationRulesForGroupResponseBodyAuthorizationRules
	SetMaxResults(v int32) *ListAuthorizationRulesForGroupResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAuthorizationRulesForGroupResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAuthorizationRulesForGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAuthorizationRulesForGroupResponseBody
	GetTotalCount() *int64
}

type ListAuthorizationRulesForGroupResponseBody struct {
	AuthorizationRules []*ListAuthorizationRulesForGroupResponseBodyAuthorizationRules `json:"AuthorizationRules,omitempty" xml:"AuthorizationRules,omitempty" type:"Repeated"`
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

func (s ListAuthorizationRulesForGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForGroupResponseBody) GetAuthorizationRules() []*ListAuthorizationRulesForGroupResponseBodyAuthorizationRules {
	return s.AuthorizationRules
}

func (s *ListAuthorizationRulesForGroupResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAuthorizationRulesForGroupResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAuthorizationRulesForGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizationRulesForGroupResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAuthorizationRulesForGroupResponseBody) SetAuthorizationRules(v []*ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) *ListAuthorizationRulesForGroupResponseBody {
	s.AuthorizationRules = v
	return s
}

func (s *ListAuthorizationRulesForGroupResponseBody) SetMaxResults(v int32) *ListAuthorizationRulesForGroupResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAuthorizationRulesForGroupResponseBody) SetNextToken(v string) *ListAuthorizationRulesForGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAuthorizationRulesForGroupResponseBody) SetRequestId(v string) *ListAuthorizationRulesForGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizationRulesForGroupResponseBody) SetTotalCount(v int64) *ListAuthorizationRulesForGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizationRulesForGroupResponseBody) Validate() error {
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

type ListAuthorizationRulesForGroupResponseBodyAuthorizationRules struct {
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
	ValidityPeriod *ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// 有效期类型，枚举值：permanent（永久），time_bound（自定义时间范围）。
	//
	// example:
	//
	// permanent
	ValidityType *string `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) GetValidityPeriod() *ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod {
	return s.ValidityPeriod
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) GetValidityType() *string {
	return s.ValidityType
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) SetAuthorizationRuleId(v string) *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules {
	s.AuthorizationRuleId = &v
	return s
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) SetInstanceId(v string) *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules {
	s.InstanceId = &v
	return s
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) SetValidityPeriod(v *ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod) *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules {
	s.ValidityPeriod = v
	return s
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) SetValidityType(v string) *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules {
	s.ValidityType = &v
	return s
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRules) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod struct {
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

func (s ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod) GoString() string {
	return s.String()
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod) SetEndTime(v int64) *ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod) SetStartTime(v int64) *ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *ListAuthorizationRulesForGroupResponseBodyAuthorizationRulesValidityPeriod) Validate() error {
	return dara.Validate(s)
}

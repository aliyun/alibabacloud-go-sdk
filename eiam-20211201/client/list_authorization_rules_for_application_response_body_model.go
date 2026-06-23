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
	// The list of authorization rules.
	AuthorizationRules []*ListAuthorizationRulesForApplicationResponseBodyAuthorizationRules `json:"AuthorizationRules,omitempty" xml:"AuthorizationRules,omitempty" type:"Repeated"`
	// The number of entries per page for paging.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned in this call. Use this token to query the next page.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
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
	// The authorization rule ID.
	//
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time range of the validity period. This parameter takes effect only when ValidityType is set to time_bound.
	ValidityPeriod *ListAuthorizationRulesForApplicationResponseBodyAuthorizationRulesValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// The validity type of the relationship. Valid values:
	//
	// - permanent: permanent
	//
	// - time_bound: custom time range.
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
	// The end time of the validity period, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1704042061000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the validity period, in UNIX timestamp format. Unit: milliseconds.
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

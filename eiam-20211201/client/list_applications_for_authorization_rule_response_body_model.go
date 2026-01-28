// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListApplicationsForAuthorizationRuleResponseBodyApplications) *ListApplicationsForAuthorizationRuleResponseBody
	GetApplications() []*ListApplicationsForAuthorizationRuleResponseBodyApplications
	SetMaxResults(v int32) *ListApplicationsForAuthorizationRuleResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationsForAuthorizationRuleResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApplicationsForAuthorizationRuleResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationsForAuthorizationRuleResponseBody
	GetTotalCount() *int64
}

type ListApplicationsForAuthorizationRuleResponseBody struct {
	Applications []*ListApplicationsForAuthorizationRuleResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
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

func (s ListApplicationsForAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForAuthorizationRuleResponseBody) GetApplications() []*ListApplicationsForAuthorizationRuleResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsForAuthorizationRuleResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationsForAuthorizationRuleResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsForAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsForAuthorizationRuleResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationsForAuthorizationRuleResponseBody) SetApplications(v []*ListApplicationsForAuthorizationRuleResponseBodyApplications) *ListApplicationsForAuthorizationRuleResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponseBody) SetMaxResults(v int32) *ListApplicationsForAuthorizationRuleResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponseBody) SetNextToken(v string) *ListApplicationsForAuthorizationRuleResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponseBody) SetRequestId(v string) *ListApplicationsForAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponseBody) SetTotalCount(v int64) *ListApplicationsForAuthorizationRuleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponseBody) Validate() error {
	if s.Applications != nil {
		for _, item := range s.Applications {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsForAuthorizationRuleResponseBodyApplications struct {
	// 应用标识。
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 实例ID。
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 有效周期。
	ValidityPeriod *ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod `json:"ValidityPeriod,omitempty" xml:"ValidityPeriod,omitempty" type:"Struct"`
	// 有效期类型，枚举值：permanent（永久），time_bound（自定义时间范围）。
	//
	// example:
	//
	// permanent
	ValidityType *string `json:"ValidityType,omitempty" xml:"ValidityType,omitempty"`
}

func (s ListApplicationsForAuthorizationRuleResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForAuthorizationRuleResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplications) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplications) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplications) GetValidityPeriod() *ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod {
	return s.ValidityPeriod
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplications) GetValidityType() *string {
	return s.ValidityType
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplications) SetApplicationId(v string) *ListApplicationsForAuthorizationRuleResponseBodyApplications {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplications) SetInstanceId(v string) *ListApplicationsForAuthorizationRuleResponseBodyApplications {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplications) SetValidityPeriod(v *ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod) *ListApplicationsForAuthorizationRuleResponseBodyApplications {
	s.ValidityPeriod = v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplications) SetValidityType(v string) *ListApplicationsForAuthorizationRuleResponseBodyApplications {
	s.ValidityType = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplications) Validate() error {
	if s.ValidityPeriod != nil {
		if err := s.ValidityPeriod.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod struct {
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

func (s ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod) GoString() string {
	return s.String()
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod) SetEndTime(v int64) *ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod {
	s.EndTime = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod) SetStartTime(v int64) *ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod {
	s.StartTime = &v
	return s
}

func (s *ListApplicationsForAuthorizationRuleResponseBodyApplicationsValidityPeriod) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterceptionSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInterceptionSummary(v *GetInterceptionSummaryResponseBodyInterceptionSummary) *GetInterceptionSummaryResponseBody
	GetInterceptionSummary() *GetInterceptionSummaryResponseBodyInterceptionSummary
	SetRequestId(v string) *GetInterceptionSummaryResponseBody
	GetRequestId() *string
}

type GetInterceptionSummaryResponseBody struct {
	// The statistics.
	InterceptionSummary *GetInterceptionSummaryResponseBodyInterceptionSummary `json:"InterceptionSummary,omitempty" xml:"InterceptionSummary,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D0760E840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInterceptionSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetInterceptionSummaryResponseBody) GetInterceptionSummary() *GetInterceptionSummaryResponseBodyInterceptionSummary {
	return s.InterceptionSummary
}

func (s *GetInterceptionSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInterceptionSummaryResponseBody) SetInterceptionSummary(v *GetInterceptionSummaryResponseBodyInterceptionSummary) *GetInterceptionSummaryResponseBody {
	s.InterceptionSummary = v
	return s
}

func (s *GetInterceptionSummaryResponseBody) SetRequestId(v string) *GetInterceptionSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInterceptionSummaryResponseBody) Validate() error {
	if s.InterceptionSummary != nil {
		if err := s.InterceptionSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInterceptionSummaryResponseBodyInterceptionSummary struct {
	// The number of clusters that are not protected.
	//
	// example:
	//
	// 0
	CloseClusterCount *int32 `json:"CloseClusterCount,omitempty" xml:"CloseClusterCount,omitempty"`
	// The number of disabled cluster defense rules.
	//
	// example:
	//
	// 0
	CloseRuleCount *int32 `json:"CloseRuleCount,omitempty" xml:"CloseRuleCount,omitempty"`
	// The total number of clusters.
	//
	// example:
	//
	// 0
	ClusterCount *int32 `json:"ClusterCount,omitempty" xml:"ClusterCount,omitempty"`
	// The total number of interception records for the specified cluster.
	//
	// example:
	//
	// 0
	InterceptionCountInDays *int32 `json:"InterceptionCountInDays,omitempty" xml:"InterceptionCountInDays,omitempty"`
	// The number of clusters that are protected.
	//
	// example:
	//
	// 0
	OpenClusterCount *int32 `json:"OpenClusterCount,omitempty" xml:"OpenClusterCount,omitempty"`
	// The number of enabled cluster defense rules.
	//
	// example:
	//
	// 0
	OpenRuleCount *int32 `json:"OpenRuleCount,omitempty" xml:"OpenRuleCount,omitempty"`
	// The number of security risks that are detected in the last 180 days.
	//
	// example:
	//
	// 0
	RiskCount180Day *int64 `json:"RiskCount180Day,omitempty" xml:"RiskCount180Day,omitempty"`
	// The number of security risks that are detected in the last 30 days.
	//
	// example:
	//
	// 0
	RiskCount30Day *int64 `json:"RiskCount30Day,omitempty" xml:"RiskCount30Day,omitempty"`
	// The number of security risks that are detected in the last 24 hours.
	//
	// example:
	//
	// 0
	RiskCountToday *int64 `json:"RiskCountToday,omitempty" xml:"RiskCountToday,omitempty"`
	// The total number of cluster defense rules.
	//
	// example:
	//
	// 0
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
}

func (s GetInterceptionSummaryResponseBodyInterceptionSummary) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionSummaryResponseBodyInterceptionSummary) GoString() string {
	return s.String()
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) GetCloseClusterCount() *int32 {
	return s.CloseClusterCount
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) GetCloseRuleCount() *int32 {
	return s.CloseRuleCount
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) GetClusterCount() *int32 {
	return s.ClusterCount
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) GetInterceptionCountInDays() *int32 {
	return s.InterceptionCountInDays
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) GetOpenClusterCount() *int32 {
	return s.OpenClusterCount
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) GetOpenRuleCount() *int32 {
	return s.OpenRuleCount
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) GetRiskCount180Day() *int64 {
	return s.RiskCount180Day
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) GetRiskCount30Day() *int64 {
	return s.RiskCount30Day
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) GetRiskCountToday() *int64 {
	return s.RiskCountToday
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) SetCloseClusterCount(v int32) *GetInterceptionSummaryResponseBodyInterceptionSummary {
	s.CloseClusterCount = &v
	return s
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) SetCloseRuleCount(v int32) *GetInterceptionSummaryResponseBodyInterceptionSummary {
	s.CloseRuleCount = &v
	return s
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) SetClusterCount(v int32) *GetInterceptionSummaryResponseBodyInterceptionSummary {
	s.ClusterCount = &v
	return s
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) SetInterceptionCountInDays(v int32) *GetInterceptionSummaryResponseBodyInterceptionSummary {
	s.InterceptionCountInDays = &v
	return s
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) SetOpenClusterCount(v int32) *GetInterceptionSummaryResponseBodyInterceptionSummary {
	s.OpenClusterCount = &v
	return s
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) SetOpenRuleCount(v int32) *GetInterceptionSummaryResponseBodyInterceptionSummary {
	s.OpenRuleCount = &v
	return s
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) SetRiskCount180Day(v int64) *GetInterceptionSummaryResponseBodyInterceptionSummary {
	s.RiskCount180Day = &v
	return s
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) SetRiskCount30Day(v int64) *GetInterceptionSummaryResponseBodyInterceptionSummary {
	s.RiskCount30Day = &v
	return s
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) SetRiskCountToday(v int64) *GetInterceptionSummaryResponseBodyInterceptionSummary {
	s.RiskCountToday = &v
	return s
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) SetRuleCount(v int32) *GetInterceptionSummaryResponseBodyInterceptionSummary {
	s.RuleCount = &v
	return s
}

func (s *GetInterceptionSummaryResponseBodyInterceptionSummary) Validate() error {
	return dara.Validate(s)
}

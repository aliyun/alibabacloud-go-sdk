// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterRuleSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterRuleSummary(v *GetClusterRuleSummaryResponseBodyClusterRuleSummary) *GetClusterRuleSummaryResponseBody
	GetClusterRuleSummary() *GetClusterRuleSummaryResponseBodyClusterRuleSummary
	SetRequestId(v string) *GetClusterRuleSummaryResponseBody
	GetRequestId() *string
}

type GetClusterRuleSummaryResponseBody struct {
	// The overall information about the cluster defense rules.
	ClusterRuleSummary *GetClusterRuleSummaryResponseBodyClusterRuleSummary `json:"ClusterRuleSummary,omitempty" xml:"ClusterRuleSummary,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 291B49F9-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterRuleSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRuleSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterRuleSummaryResponseBody) GetClusterRuleSummary() *GetClusterRuleSummaryResponseBodyClusterRuleSummary {
	return s.ClusterRuleSummary
}

func (s *GetClusterRuleSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterRuleSummaryResponseBody) SetClusterRuleSummary(v *GetClusterRuleSummaryResponseBodyClusterRuleSummary) *GetClusterRuleSummaryResponseBody {
	s.ClusterRuleSummary = v
	return s
}

func (s *GetClusterRuleSummaryResponseBody) SetRequestId(v string) *GetClusterRuleSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterRuleSummaryResponseBody) Validate() error {
	if s.ClusterRuleSummary != nil {
		if err := s.ClusterRuleSummary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClusterRuleSummaryResponseBodyClusterRuleSummary struct {
	// The number of disabled cluster defense rules.
	//
	// example:
	//
	// 1
	CloseRuleCount *int32 `json:"CloseRuleCount,omitempty" xml:"CloseRuleCount,omitempty"`
	// The number of alerts that are triggered by the cluster defense rules in the last seven days.
	//
	// example:
	//
	// 10
	InterceptionCount7Day *int64 `json:"InterceptionCount7Day,omitempty" xml:"InterceptionCount7Day,omitempty"`
	// The status of the container firewall feature. Valid values:
	//
	// 	- **0**: disabled
	//
	// 	- **1**: enabled
	//
	// example:
	//
	// 0
	InterceptionSwitch *int32 `json:"InterceptionSwitch,omitempty" xml:"InterceptionSwitch,omitempty"`
	// The interception mode.
	//
	// example:
	//
	// 0
	InterceptionType *int32 `json:"InterceptionType,omitempty" xml:"InterceptionType,omitempty"`
	// The number of enabled cluster defense rules.
	//
	// example:
	//
	// 1
	OpenRuleCount *int32 `json:"OpenRuleCount,omitempty" xml:"OpenRuleCount,omitempty"`
	// The total number of configured cluster defense rules.
	//
	// example:
	//
	// 10
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	// The number of recommended cluster defense rules.
	//
	// example:
	//
	// 1
	SuggestRuleCount *int32 `json:"SuggestRuleCount,omitempty" xml:"SuggestRuleCount,omitempty"`
}

func (s GetClusterRuleSummaryResponseBodyClusterRuleSummary) String() string {
	return dara.Prettify(s)
}

func (s GetClusterRuleSummaryResponseBodyClusterRuleSummary) GoString() string {
	return s.String()
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) GetCloseRuleCount() *int32 {
	return s.CloseRuleCount
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) GetInterceptionCount7Day() *int64 {
	return s.InterceptionCount7Day
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) GetInterceptionSwitch() *int32 {
	return s.InterceptionSwitch
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) GetInterceptionType() *int32 {
	return s.InterceptionType
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) GetOpenRuleCount() *int32 {
	return s.OpenRuleCount
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) GetSuggestRuleCount() *int32 {
	return s.SuggestRuleCount
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) SetCloseRuleCount(v int32) *GetClusterRuleSummaryResponseBodyClusterRuleSummary {
	s.CloseRuleCount = &v
	return s
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) SetInterceptionCount7Day(v int64) *GetClusterRuleSummaryResponseBodyClusterRuleSummary {
	s.InterceptionCount7Day = &v
	return s
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) SetInterceptionSwitch(v int32) *GetClusterRuleSummaryResponseBodyClusterRuleSummary {
	s.InterceptionSwitch = &v
	return s
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) SetInterceptionType(v int32) *GetClusterRuleSummaryResponseBodyClusterRuleSummary {
	s.InterceptionType = &v
	return s
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) SetOpenRuleCount(v int32) *GetClusterRuleSummaryResponseBodyClusterRuleSummary {
	s.OpenRuleCount = &v
	return s
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) SetRuleCount(v int32) *GetClusterRuleSummaryResponseBodyClusterRuleSummary {
	s.RuleCount = &v
	return s
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) SetSuggestRuleCount(v int32) *GetClusterRuleSummaryResponseBodyClusterRuleSummary {
	s.SuggestRuleCount = &v
	return s
}

func (s *GetClusterRuleSummaryResponseBodyClusterRuleSummary) Validate() error {
	return dara.Validate(s)
}

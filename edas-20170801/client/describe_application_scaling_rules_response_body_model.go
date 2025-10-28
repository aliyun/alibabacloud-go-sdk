// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationScalingRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppScalingRules(v *DescribeApplicationScalingRulesResponseBodyAppScalingRules) *DescribeApplicationScalingRulesResponseBody
	GetAppScalingRules() *DescribeApplicationScalingRulesResponseBodyAppScalingRules
	SetCode(v int32) *DescribeApplicationScalingRulesResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeApplicationScalingRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApplicationScalingRulesResponseBody
	GetRequestId() *string
}

type DescribeApplicationScalingRulesResponseBody struct {
	// The auto scaling policies of the application.
	AppScalingRules *DescribeApplicationScalingRulesResponseBodyAppScalingRules `json:"AppScalingRules,omitempty" xml:"AppScalingRules,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// a5281053-08e4-47a5-b2ab-5c0323de7b5a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBody) GetAppScalingRules() *DescribeApplicationScalingRulesResponseBodyAppScalingRules {
	return s.AppScalingRules
}

func (s *DescribeApplicationScalingRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeApplicationScalingRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationScalingRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationScalingRulesResponseBody) SetAppScalingRules(v *DescribeApplicationScalingRulesResponseBodyAppScalingRules) *DescribeApplicationScalingRulesResponseBody {
	s.AppScalingRules = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) SetCode(v int32) *DescribeApplicationScalingRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) SetMessage(v string) *DescribeApplicationScalingRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) SetRequestId(v string) *DescribeApplicationScalingRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBody) Validate() error {
	if s.AppScalingRules != nil {
		if err := s.AppScalingRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationScalingRulesResponseBodyAppScalingRules struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of auto scaling policies returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The information about auto scaling policies.
	Result []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of auto scaling policies.
	//
	// example:
	//
	// 20
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRules) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRules) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRules) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRules) GetResult() []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	return s.Result
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRules) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRules) SetCurrentPage(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRules {
	s.CurrentPage = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRules) SetPageSize(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRules {
	s.PageSize = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRules) SetResult(v []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) *DescribeApplicationScalingRulesResponseBodyAppScalingRules {
	s.Result = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRules) SetTotalSize(v int64) *DescribeApplicationScalingRulesResponseBodyAppScalingRules {
	s.TotalSize = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRules) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult struct {
	// The ID of the application to which the auto scaling policy belongs.
	//
	// example:
	//
	// 78194c76-3dca-418e-a263-cccd1ab4****
	AppId     *string                                                                    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Behaviour *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour `json:"Behaviour,omitempty" xml:"Behaviour,omitempty" type:"Struct"`
	// The time when the auto scaling policy was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 23212323123
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the auto scaling policy was last disabled. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 23212323123
	LastDisableTime *int64 `json:"LastDisableTime,omitempty" xml:"LastDisableTime,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// This parameter is deprecated.
	Metric *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric `json:"Metric,omitempty" xml:"Metric,omitempty" type:"Struct"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// Indicates whether the auto scaling policy is enabled. Valid values:
	//
	// 	- **true**: The auto scaling policy is enabled.
	//
	// 	- **false**: The auto scaling policy is disabled.
	//
	// example:
	//
	// true
	ScaleRuleEnabled *bool `json:"ScaleRuleEnabled,omitempty" xml:"ScaleRuleEnabled,omitempty"`
	// The name of the auto scaling policy.
	//
	// example:
	//
	// cpu-trigger
	ScaleRuleName *string `json:"ScaleRuleName,omitempty" xml:"ScaleRuleName,omitempty"`
	// The type of the auto scaling policy. The value is fixed to trigger.
	//
	// example:
	//
	// trigger
	ScaleRuleType *string `json:"ScaleRuleType,omitempty" xml:"ScaleRuleType,omitempty"`
	// The configurations of the trigger.
	Trigger *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger `json:"Trigger,omitempty" xml:"Trigger,omitempty" type:"Struct"`
	// The time when the auto scaling policy was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 23212323123
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetBehaviour() *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour {
	return s.Behaviour
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetLastDisableTime() *int64 {
	return s.LastDisableTime
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetMetric() *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric {
	return s.Metric
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetScaleRuleEnabled() *bool {
	return s.ScaleRuleEnabled
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetScaleRuleName() *string {
	return s.ScaleRuleName
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetScaleRuleType() *string {
	return s.ScaleRuleType
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetTrigger() *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger {
	return s.Trigger
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetAppId(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetBehaviour(v *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.Behaviour = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetCreateTime(v int64) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.CreateTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetLastDisableTime(v int64) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.LastDisableTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetMaxReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetMetric(v *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.Metric = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetMinReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetScaleRuleEnabled(v bool) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.ScaleRuleEnabled = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetScaleRuleName(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.ScaleRuleName = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetScaleRuleType(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.ScaleRuleType = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetTrigger(v *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.Trigger = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) SetUpdateTime(v int64) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult {
	s.UpdateTime = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResult) Validate() error {
	if s.Behaviour != nil {
		if err := s.Behaviour.Validate(); err != nil {
			return err
		}
	}
	if s.Metric != nil {
		if err := s.Metric.Validate(); err != nil {
			return err
		}
	}
	if s.Trigger != nil {
		if err := s.Trigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour struct {
	ScaleDown *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown `json:"ScaleDown,omitempty" xml:"ScaleDown,omitempty" type:"Struct"`
	ScaleUp   *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp   `json:"ScaleUp,omitempty" xml:"ScaleUp,omitempty" type:"Struct"`
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour) GetScaleDown() *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown {
	return s.ScaleDown
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour) GetScaleUp() *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp {
	return s.ScaleUp
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour) SetScaleDown(v *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour {
	s.ScaleDown = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour) SetScaleUp(v *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour {
	s.ScaleUp = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviour) Validate() error {
	if s.ScaleDown != nil {
		if err := s.ScaleDown.Validate(); err != nil {
			return err
		}
	}
	if s.ScaleUp != nil {
		if err := s.ScaleUp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown struct {
	Policies                   []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	SelectPolicy               *string                                                                                       `json:"SelectPolicy,omitempty" xml:"SelectPolicy,omitempty"`
	StabilizationWindowSeconds *int32                                                                                        `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown) GetPolicies() []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies {
	return s.Policies
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown) GetSelectPolicy() *string {
	return s.SelectPolicy
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown) GetStabilizationWindowSeconds() *int32 {
	return s.StabilizationWindowSeconds
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown) SetPolicies(v []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown {
	s.Policies = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown) SetSelectPolicy(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown {
	s.SelectPolicy = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown) SetStabilizationWindowSeconds(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDown) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies struct {
	PeriodSeconds *int32  `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies) GetType() *string {
	return s.Type
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies) GetValue() *string {
	return s.Value
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies) SetPeriodSeconds(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies {
	s.PeriodSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies) SetType(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies {
	s.Type = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies) SetValue(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies {
	s.Value = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleDownPolicies) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp struct {
	Policies                   []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	SelectPolicy               *string                                                                                     `json:"SelectPolicy,omitempty" xml:"SelectPolicy,omitempty"`
	StabilizationWindowSeconds *int32                                                                                      `json:"StabilizationWindowSeconds,omitempty" xml:"StabilizationWindowSeconds,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp) GetPolicies() []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies {
	return s.Policies
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp) GetSelectPolicy() *string {
	return s.SelectPolicy
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp) GetStabilizationWindowSeconds() *int32 {
	return s.StabilizationWindowSeconds
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp) SetPolicies(v []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp {
	s.Policies = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp) SetSelectPolicy(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp {
	s.SelectPolicy = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp) SetStabilizationWindowSeconds(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp {
	s.StabilizationWindowSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUp) Validate() error {
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies struct {
	PeriodSeconds *int32  `json:"PeriodSeconds,omitempty" xml:"PeriodSeconds,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value         *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies) GetPeriodSeconds() *int32 {
	return s.PeriodSeconds
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies) GetType() *string {
	return s.Type
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies) GetValue() *string {
	return s.Value
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies) SetPeriodSeconds(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies {
	s.PeriodSeconds = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies) SetType(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies {
	s.Type = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies) SetValue(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies {
	s.Value = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultBehaviourScaleUpPolicies) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// This parameter is deprecated.
	Metrics []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric) GetMetrics() []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics {
	return s.Metrics
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric) SetMaxReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric) SetMetrics(v []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric {
	s.Metrics = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric) SetMinReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetric) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	MetricTargetAverageUtilization *int32 `json:"MetricTargetAverageUtilization,omitempty" xml:"MetricTargetAverageUtilization,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// asd
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics) GetMetricTargetAverageUtilization() *int32 {
	return s.MetricTargetAverageUtilization
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics) SetMetricTargetAverageUtilization(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics {
	s.MetricTargetAverageUtilization = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics) SetMetricType(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics {
	s.MetricType = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultMetricMetrics) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger struct {
	// The maximum number of replicas. The upper limit is 1000.
	//
	// example:
	//
	// 122
	MaxReplicas *int32 `json:"MaxReplicas,omitempty" xml:"MaxReplicas,omitempty"`
	// The minimum number of replicas. The lower limit is 0.
	//
	// example:
	//
	// 1
	MinReplicas *int32 `json:"MinReplicas,omitempty" xml:"MinReplicas,omitempty"`
	// The configurations of the trigger.
	Triggers []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers `json:"Triggers,omitempty" xml:"Triggers,omitempty" type:"Repeated"`
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger) GetMaxReplicas() *int32 {
	return s.MaxReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger) GetMinReplicas() *int32 {
	return s.MinReplicas
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger) GetTriggers() []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers {
	return s.Triggers
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger) SetMaxReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger {
	s.MaxReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger) SetMinReplicas(v int32) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger {
	s.MinReplicas = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger) SetTriggers(v []*DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger {
	s.Triggers = v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTrigger) Validate() error {
	if s.Triggers != nil {
		for _, item := range s.Triggers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers struct {
	// The metadata of the trigger.
	//
	// example:
	//
	// {"dryRun":true}
	MetaData *string `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
	// The name of the trigger.
	//
	// example:
	//
	// cron-trigger
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the trigger. Valid values: cron and app_metric.
	//
	// example:
	//
	// cron
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers) GoString() string {
	return s.String()
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers) GetMetaData() *string {
	return s.MetaData
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers) GetName() *string {
	return s.Name
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers) GetType() *string {
	return s.Type
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers) SetMetaData(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers {
	s.MetaData = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers) SetName(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers {
	s.Name = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers) SetType(v string) *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers {
	s.Type = &v
	return s
}

func (s *DescribeApplicationScalingRulesResponseBodyAppScalingRulesResultTriggerTriggers) Validate() error {
	return dara.Validate(s)
}

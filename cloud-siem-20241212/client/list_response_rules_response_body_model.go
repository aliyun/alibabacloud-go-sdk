// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResponseRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResponseRulesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListResponseRulesResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListResponseRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResponseRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListResponseRulesResponseBody
	GetRequestId() *string
	SetResponseRules(v []*ListResponseRulesResponseBodyResponseRules) *ListResponseRulesResponseBody
	GetResponseRules() []*ListResponseRulesResponseBodyResponseRules
	SetTotalCount(v int32) *ListResponseRulesResponseBody
	GetTotalCount() *int32
}

type ListResponseRulesResponseBody struct {
	// The maximum number of records returned in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position from which the current call starts reading. An empty value indicates that all data has been read.
	//
	// This parameter is required.
	//
	// example:
	//
	// AAAAASLVeIxed4466E0LVmGkzwS6hJKd9DGVGMDRM6Lu****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of automatic response rules.
	ResponseRules []*ListResponseRulesResponseBodyResponseRules `json:"ResponseRules,omitempty" xml:"ResponseRules,omitempty" type:"Repeated"`
	// The total number of records that match the request conditions. This parameter is optional and is not returned by default.
	//
	// example:
	//
	// 57
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResponseRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResponseRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResponseRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResponseRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResponseRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResponseRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResponseRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResponseRulesResponseBody) GetResponseRules() []*ListResponseRulesResponseBodyResponseRules {
	return s.ResponseRules
}

func (s *ListResponseRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResponseRulesResponseBody) SetMaxResults(v int32) *ListResponseRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListResponseRulesResponseBody) SetNextToken(v string) *ListResponseRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListResponseRulesResponseBody) SetPageNumber(v int32) *ListResponseRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResponseRulesResponseBody) SetPageSize(v int32) *ListResponseRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResponseRulesResponseBody) SetRequestId(v string) *ListResponseRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResponseRulesResponseBody) SetResponseRules(v []*ListResponseRulesResponseBodyResponseRules) *ListResponseRulesResponseBody {
	s.ResponseRules = v
	return s
}

func (s *ListResponseRulesResponseBody) SetTotalCount(v int32) *ListResponseRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResponseRulesResponseBody) Validate() error {
	if s.ResponseRules != nil {
		for _, item := range s.ResponseRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResponseRulesResponseBodyResponseRules struct {
	// The time when the rule was created.
	//
	// example:
	//
	// 1769843323000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The action configuration of the automatic response rule.
	//
	// example:
	//
	// [{"actionType":"doPlaybook","playbookName":"block waf IP","playbookUuid":"system_aliyun_waf_whole_process_book","disposeParam":{"period":"7d"}}]
	ResponseActionConfig *string `json:"ResponseActionConfig,omitempty" xml:"ResponseActionConfig,omitempty"`
	// The action type of the automatic response rule. Valid values:
	//
	// - doPlaybook: execute a playbook
	//
	// - changeEventStatus: update event status
	//
	// - changeThreatLevel: update event threat level
	//
	// - addEventTag: add an event label
	//
	// - deleteEventTag: delete an event label
	//
	// - alertWhitelist: add alert to whitelist
	//
	// example:
	//
	// doPlaybook
	ResponseActionType *string `json:"ResponseActionType,omitempty" xml:"ResponseActionType,omitempty"`
	// The trigger condition configuration of the rule.
	//
	// example:
	//
	// [{"left":{"value":"threat_level"},"operator":"equals","right":{"value":"suspicious"}}]
	ResponseExecutionCondition *string `json:"ResponseExecutionCondition,omitempty" xml:"ResponseExecutionCondition,omitempty"`
	// The ID of the automatic response rule.
	//
	// example:
	//
	// 403235
	ResponseRuleId *string `json:"ResponseRuleId,omitempty" xml:"ResponseRuleId,omitempty"`
	// The name of the automatic response rule.
	//
	// example:
	//
	// Send Notification When Generating Urgent Incident
	ResponseRuleName *string `json:"ResponseRuleName,omitempty" xml:"ResponseRuleName,omitempty"`
	// The execution priority of the automatic response rule.
	//
	// example:
	//
	// 1
	ResponseRulePriority *int32  `json:"ResponseRulePriority,omitempty" xml:"ResponseRulePriority,omitempty"`
	ResponseRuleRemark   *string `json:"ResponseRuleRemark,omitempty" xml:"ResponseRuleRemark,omitempty"`
	// The status of the automatic response rule. Valid values:
	//
	// - 0: disabled
	//
	// - 100: enabled
	//
	// example:
	//
	// 0
	ResponseRuleStatus *int32 `json:"ResponseRuleStatus,omitempty" xml:"ResponseRuleStatus,omitempty"`
	// The type of the response rule. Valid values:
	//
	// - preset: predefined
	//
	// - custom: custom
	//
	// example:
	//
	// custom
	ResponseRuleType *string `json:"ResponseRuleType,omitempty" xml:"ResponseRuleType,omitempty"`
	// The trigger type of the automatic response rule. Valid values:
	//
	// - event: event occurred
	//
	// - event_update: event updated
	//
	// - alert: alert occurred
	//
	// example:
	//
	// event
	ResponseTriggerType *string `json:"ResponseTriggerType,omitempty" xml:"ResponseTriggerType,omitempty"`
	// The time when the rule was last updated.
	//
	// example:
	//
	// 1769843323000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListResponseRulesResponseBodyResponseRules) String() string {
	return dara.Prettify(s)
}

func (s ListResponseRulesResponseBodyResponseRules) GoString() string {
	return s.String()
}

func (s *ListResponseRulesResponseBodyResponseRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListResponseRulesResponseBodyResponseRules) GetResponseActionConfig() *string {
	return s.ResponseActionConfig
}

func (s *ListResponseRulesResponseBodyResponseRules) GetResponseActionType() *string {
	return s.ResponseActionType
}

func (s *ListResponseRulesResponseBodyResponseRules) GetResponseExecutionCondition() *string {
	return s.ResponseExecutionCondition
}

func (s *ListResponseRulesResponseBodyResponseRules) GetResponseRuleId() *string {
	return s.ResponseRuleId
}

func (s *ListResponseRulesResponseBodyResponseRules) GetResponseRuleName() *string {
	return s.ResponseRuleName
}

func (s *ListResponseRulesResponseBodyResponseRules) GetResponseRulePriority() *int32 {
	return s.ResponseRulePriority
}

func (s *ListResponseRulesResponseBodyResponseRules) GetResponseRuleRemark() *string {
	return s.ResponseRuleRemark
}

func (s *ListResponseRulesResponseBodyResponseRules) GetResponseRuleStatus() *int32 {
	return s.ResponseRuleStatus
}

func (s *ListResponseRulesResponseBodyResponseRules) GetResponseRuleType() *string {
	return s.ResponseRuleType
}

func (s *ListResponseRulesResponseBodyResponseRules) GetResponseTriggerType() *string {
	return s.ResponseTriggerType
}

func (s *ListResponseRulesResponseBodyResponseRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListResponseRulesResponseBodyResponseRules) SetCreateTime(v int64) *ListResponseRulesResponseBodyResponseRules {
	s.CreateTime = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) SetResponseActionConfig(v string) *ListResponseRulesResponseBodyResponseRules {
	s.ResponseActionConfig = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) SetResponseActionType(v string) *ListResponseRulesResponseBodyResponseRules {
	s.ResponseActionType = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) SetResponseExecutionCondition(v string) *ListResponseRulesResponseBodyResponseRules {
	s.ResponseExecutionCondition = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) SetResponseRuleId(v string) *ListResponseRulesResponseBodyResponseRules {
	s.ResponseRuleId = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) SetResponseRuleName(v string) *ListResponseRulesResponseBodyResponseRules {
	s.ResponseRuleName = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) SetResponseRulePriority(v int32) *ListResponseRulesResponseBodyResponseRules {
	s.ResponseRulePriority = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) SetResponseRuleRemark(v string) *ListResponseRulesResponseBodyResponseRules {
	s.ResponseRuleRemark = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) SetResponseRuleStatus(v int32) *ListResponseRulesResponseBodyResponseRules {
	s.ResponseRuleStatus = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) SetResponseRuleType(v string) *ListResponseRulesResponseBodyResponseRules {
	s.ResponseRuleType = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) SetResponseTriggerType(v string) *ListResponseRulesResponseBodyResponseRules {
	s.ResponseTriggerType = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) SetUpdateTime(v int64) *ListResponseRulesResponseBodyResponseRules {
	s.UpdateTime = &v
	return s
}

func (s *ListResponseRulesResponseBodyResponseRules) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDispatchRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDispatchRule(v *DescribeDispatchRuleResponseBodyDispatchRule) *DescribeDispatchRuleResponseBody
	GetDispatchRule() *DescribeDispatchRuleResponseBodyDispatchRule
	SetRequestId(v string) *DescribeDispatchRuleResponseBody
	GetRequestId() *string
}

type DescribeDispatchRuleResponseBody struct {
	// The struct returned.
	DispatchRule *DescribeDispatchRuleResponseBodyDispatchRule `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 34ED024E-9E31-434A-9E4E-D9D15C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDispatchRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDispatchRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBody) GetDispatchRule() *DescribeDispatchRuleResponseBodyDispatchRule {
	return s.DispatchRule
}

func (s *DescribeDispatchRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDispatchRuleResponseBody) SetDispatchRule(v *DescribeDispatchRuleResponseBodyDispatchRule) *DescribeDispatchRuleResponseBody {
	s.DispatchRule = v
	return s
}

func (s *DescribeDispatchRuleResponseBody) SetRequestId(v string) *DescribeDispatchRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDispatchRuleResponseBodyDispatchRule struct {
	// Alarm handling method.
	//
	// CREATE_ALERT: Generate an alert.
	//
	// DISCARD_ALERT: Discard the alarm event, that is, no alarm.
	//
	// example:
	//
	// CREATE_ALERT
	DispatchType *string `json:"DispatchType,omitempty" xml:"DispatchType,omitempty"`
	// The information about groups.
	GroupRules []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules `json:"GroupRules,omitempty" xml:"GroupRules,omitempty" type:"Repeated"`
	// Whether to send recovered alerts.
	//
	// true: send.
	//
	// false: do not send.
	//
	// example:
	//
	// true
	IsRecover *bool `json:"IsRecover,omitempty" xml:"IsRecover,omitempty"`
	// The information about the dispatch rule.
	LabelMatchExpressionGrid *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid `json:"LabelMatchExpressionGrid,omitempty" xml:"LabelMatchExpressionGrid,omitempty" type:"Struct"`
	// The name of the dispatch policy.
	//
	// example:
	//
	// Prometheus Alert
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The collection of notification methods.
	NotifyRules []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules `json:"NotifyRules,omitempty" xml:"NotifyRules,omitempty" type:"Repeated"`
	// The ID of the dispatch rule.
	//
	// example:
	//
	// 10282
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// Indicates whether the dispatch policy is enabled. Valid values:
	//
	// - `true`: enabled
	//
	// - `false`: disabled
	//
	// example:
	//
	// true
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRule) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) GetDispatchType() *string {
	return s.DispatchType
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) GetGroupRules() []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	return s.GroupRules
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) GetIsRecover() *bool {
	return s.IsRecover
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) GetLabelMatchExpressionGrid() *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid {
	return s.LabelMatchExpressionGrid
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) GetName() *string {
	return s.Name
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) GetNotifyRules() []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules {
	return s.NotifyRules
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) GetState() *string {
	return s.State
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetDispatchType(v string) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.DispatchType = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetGroupRules(v []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.GroupRules = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetIsRecover(v bool) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.IsRecover = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetLabelMatchExpressionGrid(v *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.LabelMatchExpressionGrid = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetName(v string) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.Name = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetNotifyRules(v []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.NotifyRules = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetRuleId(v int64) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.RuleId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetState(v string) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.State = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) Validate() error {
	return dara.Validate(s)
}

type DescribeDispatchRuleResponseBodyDispatchRuleGroupRules struct {
	// The ID of the group.
	//
	// example:
	//
	// 1
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The grouping interval.
	//
	// example:
	//
	// 15
	GroupInterval *int64 `json:"GroupInterval,omitempty" xml:"GroupInterval,omitempty"`
	// The waiting time for grouping.
	//
	// example:
	//
	// 10
	GroupWaitTime *int64 `json:"GroupWaitTime,omitempty" xml:"GroupWaitTime,omitempty"`
	// The grouping fields.
	GroupingFields []*string `json:"GroupingFields,omitempty" xml:"GroupingFields,omitempty" type:"Repeated"`
	// The time interval at which a notification is resent for a long-lasting unresolved alert. Unit: seconds.
	//
	// example:
	//
	// 20
	RepeatInterval *int64 `json:"RepeatInterval,omitempty" xml:"RepeatInterval,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) GetGroupInterval() *int64 {
	return s.GroupInterval
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) GetGroupWaitTime() *int64 {
	return s.GroupWaitTime
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) GetGroupingFields() []*string {
	return s.GroupingFields
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) GetRepeatInterval() *int64 {
	return s.RepeatInterval
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupId(v int64) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupInterval(v int64) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupInterval = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupWaitTime(v int64) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupWaitTime = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetGroupingFields(v []*string) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.GroupingFields = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) SetRepeatInterval(v int64) *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	s.RepeatInterval = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) Validate() error {
	return dara.Validate(s)
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid struct {
	// The collection of dispatch rules.
	LabelMatchExpressionGroups []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups `json:"LabelMatchExpressionGroups,omitempty" xml:"LabelMatchExpressionGroups,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) String() string {
	return dara.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) GetLabelMatchExpressionGroups() []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups {
	return s.LabelMatchExpressionGroups
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) SetLabelMatchExpressionGroups(v []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid {
	s.LabelMatchExpressionGroups = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid) Validate() error {
	return dara.Validate(s)
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups struct {
	// The collection of conditions of the dispatch rule.
	LabelMatchExpressions []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions `json:"LabelMatchExpressions,omitempty" xml:"LabelMatchExpressions,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) GetLabelMatchExpressions() []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	return s.LabelMatchExpressions
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) SetLabelMatchExpressions(v []*DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups {
	s.LabelMatchExpressions = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions struct {
	// The key of the tag of the dispatch rule. Valid values:
	//
	// 	- `_aliyun_arms_userid`: user ID
	//
	// 	- `_aliyun_arms_involvedObject_kind`: type of the associated object
	//
	// 	- `_aliyun_arms_involvedObject_id`: ID of the associated object
	//
	// 	- `_aliyun_arms_involvedObject_name`: name of the associated object
	//
	// 	- `_aliyun_arms_alert_name`: alert name
	//
	// 	- `_aliyun_arms_alert_rule_id`: alert rule ID
	//
	// 	- `_aliyun_arms_alert_type`: alert type
	//
	// 	- `_aliyun_arms_alert_level`: alert severity
	//
	// example:
	//
	// _aliyun_arms_involvedObject_kind
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The operator used in the dispatch rule. Valid values:
	//
	// 	- `eq`: equals to.
	//
	// 	- `re`: matches a regular expression.
	//
	// example:
	//
	// eq
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// app
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) String() string {
	return dara.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GetKey() *string {
	return s.Key
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GetOperator() *string {
	return s.Operator
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) GetValue() *string {
	return s.Value
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetKey(v string) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Key = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetOperator(v string) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Operator = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) SetValue(v string) *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	s.Value = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) Validate() error {
	return dara.Validate(s)
}

type DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules struct {
	// The notification method Array.
	NotifyChannels []*string `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	// The collection of alert contacts.
	NotifyObjects []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects `json:"NotifyObjects,omitempty" xml:"NotifyObjects,omitempty" type:"Repeated"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) GetNotifyChannels() []*string {
	return s.NotifyChannels
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) GetNotifyObjects() []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	return s.NotifyObjects
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) SetNotifyChannels(v []*string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules {
	s.NotifyChannels = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) SetNotifyObjects(v []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules {
	s.NotifyObjects = v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules) Validate() error {
	return dara.Validate(s)
}

type DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects struct {
	// The name of the contact or contact group.
	//
	// example:
	//
	// JohnDoe
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the contact or contact group.
	//
	// example:
	//
	// 1
	NotifyObjectId *string `json:"NotifyObjectId,omitempty" xml:"NotifyObjectId,omitempty"`
	// The type of the alert contact. Valid values:
	//
	// - `CONTACT`: contact
	//
	// - `CONTACT_GROUP`: contact group
	//
	// example:
	//
	// CONTACT
	NotifyType *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) String() string {
	return dara.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) GetName() *string {
	return s.Name
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) GetNotifyObjectId() *string {
	return s.NotifyObjectId
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) GetNotifyType() *string {
	return s.NotifyType
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetName(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.Name = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetNotifyObjectId(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.NotifyObjectId = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) SetNotifyType(v string) *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects {
	s.NotifyType = &v
	return s
}

func (s *DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects) Validate() error {
	return dara.Validate(s)
}

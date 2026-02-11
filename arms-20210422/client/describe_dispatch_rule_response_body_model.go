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
	DispatchRule *DescribeDispatchRuleResponseBodyDispatchRule `json:"DispatchRule,omitempty" xml:"DispatchRule,omitempty" type:"Struct"`
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	if s.DispatchRule != nil {
		if err := s.DispatchRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDispatchRuleResponseBodyDispatchRule struct {
	GroupRules               []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules             `json:"GroupRules,omitempty" xml:"GroupRules,omitempty" type:"Repeated"`
	LabelMatchExpressionGrid *DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid `json:"LabelMatchExpressionGrid,omitempty" xml:"LabelMatchExpressionGrid,omitempty" type:"Struct"`
	Name                     *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyRules              []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRules            `json:"NotifyRules,omitempty" xml:"NotifyRules,omitempty" type:"Repeated"`
	RuleId                   *int64                                                                `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	State                    *string                                                               `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeDispatchRuleResponseBodyDispatchRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDispatchRuleResponseBodyDispatchRule) GoString() string {
	return s.String()
}

func (s *DescribeDispatchRuleResponseBodyDispatchRule) GetGroupRules() []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules {
	return s.GroupRules
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

func (s *DescribeDispatchRuleResponseBodyDispatchRule) SetGroupRules(v []*DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) *DescribeDispatchRuleResponseBodyDispatchRule {
	s.GroupRules = v
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
	if s.GroupRules != nil {
		for _, item := range s.GroupRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LabelMatchExpressionGrid != nil {
		if err := s.LabelMatchExpressionGrid.Validate(); err != nil {
			return err
		}
	}
	if s.NotifyRules != nil {
		for _, item := range s.NotifyRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDispatchRuleResponseBodyDispatchRuleGroupRules struct {
	GroupId        *int64    `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupInterval  *int64    `json:"GroupInterval,omitempty" xml:"GroupInterval,omitempty"`
	GroupWaitTime  *int64    `json:"GroupWaitTime,omitempty" xml:"GroupWaitTime,omitempty"`
	GroupingFields []*string `json:"GroupingFields,omitempty" xml:"GroupingFields,omitempty" type:"Repeated"`
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

func (s *DescribeDispatchRuleResponseBodyDispatchRuleGroupRules) Validate() error {
	return dara.Validate(s)
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGrid struct {
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
	if s.LabelMatchExpressionGroups != nil {
		for _, item := range s.LabelMatchExpressionGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroups struct {
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
	if s.LabelMatchExpressions != nil {
		for _, item := range s.LabelMatchExpressions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDispatchRuleResponseBodyDispatchRuleLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions struct {
	Key      *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	NotifyChannels []*string                                                               `json:"NotifyChannels,omitempty" xml:"NotifyChannels,omitempty" type:"Repeated"`
	NotifyObjects  []*DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects `json:"NotifyObjects,omitempty" xml:"NotifyObjects,omitempty" type:"Repeated"`
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
	if s.NotifyObjects != nil {
		for _, item := range s.NotifyObjects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDispatchRuleResponseBodyDispatchRuleNotifyRulesNotifyObjects struct {
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NotifyObjectId *string `json:"NotifyObjectId,omitempty" xml:"NotifyObjectId,omitempty"`
	NotifyType     *string `json:"NotifyType,omitempty" xml:"NotifyType,omitempty"`
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVServerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVServerGroupsResponseBody
	GetRequestId() *string
	SetVServerGroups(v *DescribeVServerGroupsResponseBodyVServerGroups) *DescribeVServerGroupsResponseBody
	GetVServerGroups() *DescribeVServerGroupsResponseBodyVServerGroups
}

type DescribeVServerGroupsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9DEC9C28-AB05-4DDF-9A78-6B08EC9CE18C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The backend servers.
	VServerGroups *DescribeVServerGroupsResponseBodyVServerGroups `json:"VServerGroups,omitempty" xml:"VServerGroups,omitempty" type:"Struct"`
}

func (s DescribeVServerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVServerGroupsResponseBody) GetVServerGroups() *DescribeVServerGroupsResponseBodyVServerGroups {
	return s.VServerGroups
}

func (s *DescribeVServerGroupsResponseBody) SetRequestId(v string) *DescribeVServerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVServerGroupsResponseBody) SetVServerGroups(v *DescribeVServerGroupsResponseBodyVServerGroups) *DescribeVServerGroupsResponseBody {
	s.VServerGroups = v
	return s
}

func (s *DescribeVServerGroupsResponseBody) Validate() error {
	if s.VServerGroups != nil {
		if err := s.VServerGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVServerGroupsResponseBodyVServerGroups struct {
	VServerGroup []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup `json:"VServerGroup,omitempty" xml:"VServerGroup,omitempty" type:"Repeated"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroups) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroups) GetVServerGroup() []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	return s.VServerGroup
}

func (s *DescribeVServerGroupsResponseBodyVServerGroups) SetVServerGroup(v []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) *DescribeVServerGroupsResponseBodyVServerGroups {
	s.VServerGroup = v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroups) Validate() error {
	if s.VServerGroup != nil {
		for _, item := range s.VServerGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup struct {
	// The associated resources.
	AssociatedObjects *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects `json:"AssociatedObjects,omitempty" xml:"AssociatedObjects,omitempty" type:"Struct"`
	// The time when the CLB instance was created. The time follows the `YYYY-MM-DDThh:mm:ssZ` format.
	//
	// example:
	//
	// 2022-08-31T02:49:05Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of servers.
	//
	// This parameter is unavailable by default. To use this parameter, submit a ticket or contact your account manager.
	//
	// example:
	//
	// 1
	ServerCount *int64 `json:"ServerCount,omitempty" xml:"ServerCount,omitempty"`
	// The tags.
	Tags *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The server group ID.
	//
	// example:
	//
	// rsp-0bfuc*****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
	// The server group name.
	//
	// example:
	//
	// Group3
	VServerGroupName *string `json:"VServerGroupName,omitempty" xml:"VServerGroupName,omitempty"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) GetAssociatedObjects() *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects {
	return s.AssociatedObjects
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) GetServerCount() *int64 {
	return s.ServerCount
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) GetTags() *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTags {
	return s.Tags
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) GetVServerGroupName() *string {
	return s.VServerGroupName
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) SetAssociatedObjects(v *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	s.AssociatedObjects = v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) SetCreateTime(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) SetServerCount(v int64) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	s.ServerCount = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) SetTags(v *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTags) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	s.Tags = v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) SetVServerGroupId(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) SetVServerGroupName(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup {
	s.VServerGroupName = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroup) Validate() error {
	if s.AssociatedObjects != nil {
		if err := s.AssociatedObjects.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects struct {
	// The listeners.
	Listeners *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners `json:"Listeners,omitempty" xml:"Listeners,omitempty" type:"Struct"`
	// The forwarding rules.
	Rules *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) GetListeners() *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners {
	return s.Listeners
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) GetRules() *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules {
	return s.Rules
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) SetListeners(v *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects {
	s.Listeners = v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) SetRules(v *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects {
	s.Rules = v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjects) Validate() error {
	if s.Listeners != nil {
		if err := s.Listeners.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners struct {
	Listener []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener `json:"Listener,omitempty" xml:"Listener,omitempty" type:"Repeated"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners) GetListener() []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener {
	return s.Listener
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners) SetListener(v []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners {
	s.Listener = v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListeners) Validate() error {
	if s.Listener != nil {
		for _, item := range s.Listener {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener struct {
	// The listener port.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The listener protocol. Valid values: **tcp**, **udp**, **http**, and **https**.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) GetPort() *int32 {
	return s.Port
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) SetPort(v int32) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener {
	s.Port = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) SetProtocol(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener {
	s.Protocol = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsListenersListener) Validate() error {
	return dara.Validate(s)
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules struct {
	Rule []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules) GetRule() []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule {
	return s.Rule
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules) SetRule(v []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules {
	s.Rule = v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRules) Validate() error {
	if s.Rule != nil {
		for _, item := range s.Rule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule struct {
	// The requested domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the forwarding rule.
	//
	// example:
	//
	// rule-a3x3pg1yohq3lq****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the forwarding rule.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The request URL.
	//
	// example:
	//
	// /example
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) GetDomain() *string {
	return s.Domain
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) GetUrl() *string {
	return s.Url
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) SetDomain(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule {
	s.Domain = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) SetRuleId(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) SetRuleName(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule {
	s.RuleName = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) SetUrl(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule {
	s.Url = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupAssociatedObjectsRulesRule) Validate() error {
	return dara.Validate(s)
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTags struct {
	Tag []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTags) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTags) GetTag() []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag {
	return s.Tag
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTags) SetTag(v []*DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTags {
	s.Tag = v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag) SetTagKey(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag) SetTagValue(v string) *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeVServerGroupsResponseBodyVServerGroupsVServerGroupTagsTag) Validate() error {
	return dara.Validate(s)
}

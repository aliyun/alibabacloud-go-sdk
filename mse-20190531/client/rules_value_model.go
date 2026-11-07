// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRulesValue interface {
	dara.Model
	String() string
	GoString() string
	SetStatus(v int32) *RulesValue
	GetStatus() *int32
	SetRate(v int32) *RulesValue
	GetRate() *int32
	SetEnable(v bool) *RulesValue
	GetEnable() *bool
	SetTag(v string) *RulesValue
	GetTag() *string
	SetName(v string) *RulesValue
	GetName() *string
	SetId(v int64) *RulesValue
	GetId() *int64
	SetInstanceNum(v int32) *RulesValue
	GetInstanceNum() *int32
	SetRules(v *RulesValueRules) *RulesValue
	GetRules() *RulesValueRules
}

type RulesValue struct {
	// The status.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The percentage.
	//
	// example:
	//
	// 10
	Rate *int32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
	// Specifies whether the rule is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The environment of this routing rule.
	//
	// example:
	//
	// gray
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	// The name of the routing rule.
	//
	// example:
	//
	// gray
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 10
	InstanceNum *int32 `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	// The rule information.
	Rules *RulesValueRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s RulesValue) String() string {
	return dara.Prettify(s)
}

func (s RulesValue) GoString() string {
	return s.String()
}

func (s *RulesValue) GetStatus() *int32 {
	return s.Status
}

func (s *RulesValue) GetRate() *int32 {
	return s.Rate
}

func (s *RulesValue) GetEnable() *bool {
	return s.Enable
}

func (s *RulesValue) GetTag() *string {
	return s.Tag
}

func (s *RulesValue) GetName() *string {
	return s.Name
}

func (s *RulesValue) GetId() *int64 {
	return s.Id
}

func (s *RulesValue) GetInstanceNum() *int32 {
	return s.InstanceNum
}

func (s *RulesValue) GetRules() *RulesValueRules {
	return s.Rules
}

func (s *RulesValue) SetStatus(v int32) *RulesValue {
	s.Status = &v
	return s
}

func (s *RulesValue) SetRate(v int32) *RulesValue {
	s.Rate = &v
	return s
}

func (s *RulesValue) SetEnable(v bool) *RulesValue {
	s.Enable = &v
	return s
}

func (s *RulesValue) SetTag(v string) *RulesValue {
	s.Tag = &v
	return s
}

func (s *RulesValue) SetName(v string) *RulesValue {
	s.Name = &v
	return s
}

func (s *RulesValue) SetId(v int64) *RulesValue {
	s.Id = &v
	return s
}

func (s *RulesValue) SetInstanceNum(v int32) *RulesValue {
	s.InstanceNum = &v
	return s
}

func (s *RulesValue) SetRules(v *RulesValueRules) *RulesValue {
	s.Rules = v
	return s
}

func (s *RulesValue) Validate() error {
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RulesValueRules struct {
	// The Spring Cloud rules.
	Springcloud []*RulesValueRulesSpringcloud `json:"springcloud,omitempty" xml:"springcloud,omitempty" type:"Repeated"`
	// The Dubbo rules.
	Dubbo []*RulesValueRulesDubbo `json:"dubbo,omitempty" xml:"dubbo,omitempty" type:"Repeated"`
}

func (s RulesValueRules) String() string {
	return dara.Prettify(s)
}

func (s RulesValueRules) GoString() string {
	return s.String()
}

func (s *RulesValueRules) GetSpringcloud() []*RulesValueRulesSpringcloud {
	return s.Springcloud
}

func (s *RulesValueRules) GetDubbo() []*RulesValueRulesDubbo {
	return s.Dubbo
}

func (s *RulesValueRules) SetSpringcloud(v []*RulesValueRulesSpringcloud) *RulesValueRules {
	s.Springcloud = v
	return s
}

func (s *RulesValueRules) SetDubbo(v []*RulesValueRulesDubbo) *RulesValueRules {
	s.Dubbo = v
	return s
}

func (s *RulesValueRules) Validate() error {
	if s.Springcloud != nil {
		for _, item := range s.Springcloud {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Dubbo != nil {
		for _, item := range s.Dubbo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RulesValueRulesSpringcloud struct {
	// The logical operator used between the following items. Valid values: AND and OR.
	//
	// example:
	//
	// AND
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The list of matching conditions for REST requests.
	RestItems []*RulesValueRulesSpringcloudRestItems `json:"restItems,omitempty" xml:"restItems,omitempty" type:"Repeated"`
	// The policy type.
	//
	// example:
	//
	// CONTENT
	TriggerPolicy *string `json:"triggerPolicy,omitempty" xml:"triggerPolicy,omitempty"`
	// Specifies whether the rule is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// ***@***
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"priority,omitempty" xml:"priority,omitempty"`
	// The list of tags.
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The list of paths.
	Paths []*string `json:"paths,omitempty" xml:"paths,omitempty" type:"Repeated"`
	// The path.
	//
	// example:
	//
	// /hello
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s RulesValueRulesSpringcloud) String() string {
	return dara.Prettify(s)
}

func (s RulesValueRulesSpringcloud) GoString() string {
	return s.String()
}

func (s *RulesValueRulesSpringcloud) GetCondition() *string {
	return s.Condition
}

func (s *RulesValueRulesSpringcloud) GetRestItems() []*RulesValueRulesSpringcloudRestItems {
	return s.RestItems
}

func (s *RulesValueRulesSpringcloud) GetTriggerPolicy() *string {
	return s.TriggerPolicy
}

func (s *RulesValueRulesSpringcloud) GetEnable() *bool {
	return s.Enable
}

func (s *RulesValueRulesSpringcloud) GetAppId() *string {
	return s.AppId
}

func (s *RulesValueRulesSpringcloud) GetPriority() *int32 {
	return s.Priority
}

func (s *RulesValueRulesSpringcloud) GetTags() []*string {
	return s.Tags
}

func (s *RulesValueRulesSpringcloud) GetPaths() []*string {
	return s.Paths
}

func (s *RulesValueRulesSpringcloud) GetPath() *string {
	return s.Path
}

func (s *RulesValueRulesSpringcloud) SetCondition(v string) *RulesValueRulesSpringcloud {
	s.Condition = &v
	return s
}

func (s *RulesValueRulesSpringcloud) SetRestItems(v []*RulesValueRulesSpringcloudRestItems) *RulesValueRulesSpringcloud {
	s.RestItems = v
	return s
}

func (s *RulesValueRulesSpringcloud) SetTriggerPolicy(v string) *RulesValueRulesSpringcloud {
	s.TriggerPolicy = &v
	return s
}

func (s *RulesValueRulesSpringcloud) SetEnable(v bool) *RulesValueRulesSpringcloud {
	s.Enable = &v
	return s
}

func (s *RulesValueRulesSpringcloud) SetAppId(v string) *RulesValueRulesSpringcloud {
	s.AppId = &v
	return s
}

func (s *RulesValueRulesSpringcloud) SetPriority(v int32) *RulesValueRulesSpringcloud {
	s.Priority = &v
	return s
}

func (s *RulesValueRulesSpringcloud) SetTags(v []*string) *RulesValueRulesSpringcloud {
	s.Tags = v
	return s
}

func (s *RulesValueRulesSpringcloud) SetPaths(v []*string) *RulesValueRulesSpringcloud {
	s.Paths = v
	return s
}

func (s *RulesValueRulesSpringcloud) SetPath(v string) *RulesValueRulesSpringcloud {
	s.Path = &v
	return s
}

func (s *RulesValueRulesSpringcloud) Validate() error {
	if s.RestItems != nil {
		for _, item := range s.RestItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RulesValueRulesSpringcloudRestItems struct {
	// The value to be compared for operators such as rawvalue.
	//
	// example:
	//
	// 10
	Datum *string `json:"datum,omitempty" xml:"datum,omitempty"`
	// The operator. A value of rawvalue indicates direct comparison. A value of mode indicates a modulo operation. A value of list indicates a whitelist.
	//
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The field required for the list operator.
	NameList []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	// The comparison operator. Valid values: >=, <=, >, <, and ==.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// The divisor field required for the mod operator.
	//
	// example:
	//
	// 30
	Divisor *int32 `json:"divisor,omitempty" xml:"divisor,omitempty"`
	// The percentage value field required for the percentage operation.
	//
	// example:
	//
	// 30
	Remainder *int32 `json:"remainder,omitempty" xml:"remainder,omitempty"`
	// The proportion rule. For example, 20% indicates that 20% of the traffic is routed to the node corresponding to the tag.
	//
	// example:
	//
	// 20
	Rate *int32 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The type.
	//
	// example:
	//
	// test
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The name.
	//
	// example:
	//
	// key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value.
	//
	// example:
	//
	// test
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RulesValueRulesSpringcloudRestItems) String() string {
	return dara.Prettify(s)
}

func (s RulesValueRulesSpringcloudRestItems) GoString() string {
	return s.String()
}

func (s *RulesValueRulesSpringcloudRestItems) GetDatum() *string {
	return s.Datum
}

func (s *RulesValueRulesSpringcloudRestItems) GetOperator() *string {
	return s.Operator
}

func (s *RulesValueRulesSpringcloudRestItems) GetNameList() []*string {
	return s.NameList
}

func (s *RulesValueRulesSpringcloudRestItems) GetCond() *string {
	return s.Cond
}

func (s *RulesValueRulesSpringcloudRestItems) GetDivisor() *int32 {
	return s.Divisor
}

func (s *RulesValueRulesSpringcloudRestItems) GetRemainder() *int32 {
	return s.Remainder
}

func (s *RulesValueRulesSpringcloudRestItems) GetRate() *int32 {
	return s.Rate
}

func (s *RulesValueRulesSpringcloudRestItems) GetType() *string {
	return s.Type
}

func (s *RulesValueRulesSpringcloudRestItems) GetName() *string {
	return s.Name
}

func (s *RulesValueRulesSpringcloudRestItems) GetValue() interface{} {
	return s.Value
}

func (s *RulesValueRulesSpringcloudRestItems) SetDatum(v string) *RulesValueRulesSpringcloudRestItems {
	s.Datum = &v
	return s
}

func (s *RulesValueRulesSpringcloudRestItems) SetOperator(v string) *RulesValueRulesSpringcloudRestItems {
	s.Operator = &v
	return s
}

func (s *RulesValueRulesSpringcloudRestItems) SetNameList(v []*string) *RulesValueRulesSpringcloudRestItems {
	s.NameList = v
	return s
}

func (s *RulesValueRulesSpringcloudRestItems) SetCond(v string) *RulesValueRulesSpringcloudRestItems {
	s.Cond = &v
	return s
}

func (s *RulesValueRulesSpringcloudRestItems) SetDivisor(v int32) *RulesValueRulesSpringcloudRestItems {
	s.Divisor = &v
	return s
}

func (s *RulesValueRulesSpringcloudRestItems) SetRemainder(v int32) *RulesValueRulesSpringcloudRestItems {
	s.Remainder = &v
	return s
}

func (s *RulesValueRulesSpringcloudRestItems) SetRate(v int32) *RulesValueRulesSpringcloudRestItems {
	s.Rate = &v
	return s
}

func (s *RulesValueRulesSpringcloudRestItems) SetType(v string) *RulesValueRulesSpringcloudRestItems {
	s.Type = &v
	return s
}

func (s *RulesValueRulesSpringcloudRestItems) SetName(v string) *RulesValueRulesSpringcloudRestItems {
	s.Name = &v
	return s
}

func (s *RulesValueRulesSpringcloudRestItems) SetValue(v interface{}) *RulesValueRulesSpringcloudRestItems {
	s.Value = v
	return s
}

func (s *RulesValueRulesSpringcloudRestItems) Validate() error {
	return dara.Validate(s)
}

type RulesValueRulesDubbo struct {
	// The ID of the application.
	//
	// example:
	//
	// ***@***
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// The list of tags.
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The policy type.
	//
	// example:
	//
	// CONTENT
	TriggerPolicy *string `json:"triggerPolicy,omitempty" xml:"triggerPolicy,omitempty"`
	// The service name (interface).
	//
	// example:
	//
	// HelloService
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// The Dubbo service group.
	//
	// example:
	//
	// default
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The Dubbo service version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The method name of the Dubbo service.
	//
	// example:
	//
	// hello
	MethodName *string `json:"methodName,omitempty" xml:"methodName,omitempty"`
	// The list of parameter types.
	ParamTypes []*string `json:"paramTypes,omitempty" xml:"paramTypes,omitempty" type:"Repeated"`
	// The logical operator used between the following items. Valid values: AND and OR.
	//
	// example:
	//
	// AND
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// The list of parameter content.
	ArgumentItems []*RulesValueRulesDubboArgumentItems `json:"argumentItems,omitempty" xml:"argumentItems,omitempty" type:"Repeated"`
}

func (s RulesValueRulesDubbo) String() string {
	return dara.Prettify(s)
}

func (s RulesValueRulesDubbo) GoString() string {
	return s.String()
}

func (s *RulesValueRulesDubbo) GetAppId() *string {
	return s.AppId
}

func (s *RulesValueRulesDubbo) GetTags() []*string {
	return s.Tags
}

func (s *RulesValueRulesDubbo) GetTriggerPolicy() *string {
	return s.TriggerPolicy
}

func (s *RulesValueRulesDubbo) GetServiceName() *string {
	return s.ServiceName
}

func (s *RulesValueRulesDubbo) GetGroup() *string {
	return s.Group
}

func (s *RulesValueRulesDubbo) GetVersion() *string {
	return s.Version
}

func (s *RulesValueRulesDubbo) GetMethodName() *string {
	return s.MethodName
}

func (s *RulesValueRulesDubbo) GetParamTypes() []*string {
	return s.ParamTypes
}

func (s *RulesValueRulesDubbo) GetCondition() *string {
	return s.Condition
}

func (s *RulesValueRulesDubbo) GetArgumentItems() []*RulesValueRulesDubboArgumentItems {
	return s.ArgumentItems
}

func (s *RulesValueRulesDubbo) SetAppId(v string) *RulesValueRulesDubbo {
	s.AppId = &v
	return s
}

func (s *RulesValueRulesDubbo) SetTags(v []*string) *RulesValueRulesDubbo {
	s.Tags = v
	return s
}

func (s *RulesValueRulesDubbo) SetTriggerPolicy(v string) *RulesValueRulesDubbo {
	s.TriggerPolicy = &v
	return s
}

func (s *RulesValueRulesDubbo) SetServiceName(v string) *RulesValueRulesDubbo {
	s.ServiceName = &v
	return s
}

func (s *RulesValueRulesDubbo) SetGroup(v string) *RulesValueRulesDubbo {
	s.Group = &v
	return s
}

func (s *RulesValueRulesDubbo) SetVersion(v string) *RulesValueRulesDubbo {
	s.Version = &v
	return s
}

func (s *RulesValueRulesDubbo) SetMethodName(v string) *RulesValueRulesDubbo {
	s.MethodName = &v
	return s
}

func (s *RulesValueRulesDubbo) SetParamTypes(v []*string) *RulesValueRulesDubbo {
	s.ParamTypes = v
	return s
}

func (s *RulesValueRulesDubbo) SetCondition(v string) *RulesValueRulesDubbo {
	s.Condition = &v
	return s
}

func (s *RulesValueRulesDubbo) SetArgumentItems(v []*RulesValueRulesDubboArgumentItems) *RulesValueRulesDubbo {
	s.ArgumentItems = v
	return s
}

func (s *RulesValueRulesDubbo) Validate() error {
	if s.ArgumentItems != nil {
		for _, item := range s.ArgumentItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RulesValueRulesDubboArgumentItems struct {
	// The operator. A value of rawvalue indicates direct comparison. A value of mode indicates a modulo operation. A value of list indicates a whitelist.
	//
	// example:
	//
	// rawvalue
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The list of names.
	NameList []*string `json:"nameList,omitempty" xml:"nameList,omitempty" type:"Repeated"`
	// The value to be compared for operators such as rawvalue.
	//
	// example:
	//
	// 30
	Datum *string `json:"datum,omitempty" xml:"datum,omitempty"`
	// The comparison operator. Valid values: >=, <=, >, <, and ==.
	//
	// example:
	//
	// ==
	Cond *string `json:"cond,omitempty" xml:"cond,omitempty"`
	// The divisor field required for the mod operator.
	//
	// example:
	//
	// 30
	Divisor *int32 `json:"divisor,omitempty" xml:"divisor,omitempty"`
	// The percentage value field required for the percentage operation.
	//
	// example:
	//
	// 30
	Remainder *int32 `json:"remainder,omitempty" xml:"remainder,omitempty"`
	// The proportion rule. For example, 20% indicates that 20% of the traffic is routed to the node corresponding to the tag.
	//
	// example:
	//
	// 10
	Rate *int32 `json:"rate,omitempty" xml:"rate,omitempty"`
	// The position of the parameter, starting from 0.
	//
	// example:
	//
	// 0
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// The expression. See other field descriptions for details.
	Expr *string `json:"expr,omitempty" xml:"expr,omitempty"`
	// The value to compare. The value obtained from the expression is compared with this value. If the operator is a list whitelist, separate values with commas (,), such as 1,2,3.
	//
	// example:
	//
	// 1,2,3
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s RulesValueRulesDubboArgumentItems) String() string {
	return dara.Prettify(s)
}

func (s RulesValueRulesDubboArgumentItems) GoString() string {
	return s.String()
}

func (s *RulesValueRulesDubboArgumentItems) GetOperator() *string {
	return s.Operator
}

func (s *RulesValueRulesDubboArgumentItems) GetNameList() []*string {
	return s.NameList
}

func (s *RulesValueRulesDubboArgumentItems) GetDatum() *string {
	return s.Datum
}

func (s *RulesValueRulesDubboArgumentItems) GetCond() *string {
	return s.Cond
}

func (s *RulesValueRulesDubboArgumentItems) GetDivisor() *int32 {
	return s.Divisor
}

func (s *RulesValueRulesDubboArgumentItems) GetRemainder() *int32 {
	return s.Remainder
}

func (s *RulesValueRulesDubboArgumentItems) GetRate() *int32 {
	return s.Rate
}

func (s *RulesValueRulesDubboArgumentItems) GetIndex() *int32 {
	return s.Index
}

func (s *RulesValueRulesDubboArgumentItems) GetExpr() *string {
	return s.Expr
}

func (s *RulesValueRulesDubboArgumentItems) GetValue() interface{} {
	return s.Value
}

func (s *RulesValueRulesDubboArgumentItems) SetOperator(v string) *RulesValueRulesDubboArgumentItems {
	s.Operator = &v
	return s
}

func (s *RulesValueRulesDubboArgumentItems) SetNameList(v []*string) *RulesValueRulesDubboArgumentItems {
	s.NameList = v
	return s
}

func (s *RulesValueRulesDubboArgumentItems) SetDatum(v string) *RulesValueRulesDubboArgumentItems {
	s.Datum = &v
	return s
}

func (s *RulesValueRulesDubboArgumentItems) SetCond(v string) *RulesValueRulesDubboArgumentItems {
	s.Cond = &v
	return s
}

func (s *RulesValueRulesDubboArgumentItems) SetDivisor(v int32) *RulesValueRulesDubboArgumentItems {
	s.Divisor = &v
	return s
}

func (s *RulesValueRulesDubboArgumentItems) SetRemainder(v int32) *RulesValueRulesDubboArgumentItems {
	s.Remainder = &v
	return s
}

func (s *RulesValueRulesDubboArgumentItems) SetRate(v int32) *RulesValueRulesDubboArgumentItems {
	s.Rate = &v
	return s
}

func (s *RulesValueRulesDubboArgumentItems) SetIndex(v int32) *RulesValueRulesDubboArgumentItems {
	s.Index = &v
	return s
}

func (s *RulesValueRulesDubboArgumentItems) SetExpr(v string) *RulesValueRulesDubboArgumentItems {
	s.Expr = &v
	return s
}

func (s *RulesValueRulesDubboArgumentItems) SetValue(v interface{}) *RulesValueRulesDubboArgumentItems {
	s.Value = v
	return s
}

func (s *RulesValueRulesDubboArgumentItems) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMonitorGroupDynamicRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v int64) *PutMonitorGroupDynamicRuleRequest
	GetGroupId() *int64
	SetGroupRules(v []*PutMonitorGroupDynamicRuleRequestGroupRules) *PutMonitorGroupDynamicRuleRequest
	GetGroupRules() []*PutMonitorGroupDynamicRuleRequestGroupRules
	SetIsAsync(v bool) *PutMonitorGroupDynamicRuleRequest
	GetIsAsync() *bool
	SetRegionId(v string) *PutMonitorGroupDynamicRuleRequest
	GetRegionId() *string
}

type PutMonitorGroupDynamicRuleRequest struct {
	// The ID of the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// None.
	//
	// This parameter is required.
	GroupRules []*PutMonitorGroupDynamicRuleRequestGroupRules `json:"GroupRules,omitempty" xml:"GroupRules,omitempty" type:"Repeated"`
	// The mode for creating the alert rule. Valid values:
	//
	// 	- true: creates asynchronously
	//
	// 	- false (default): creates synchronously
	//
	// example:
	//
	// false
	IsAsync  *bool   `json:"IsAsync,omitempty" xml:"IsAsync,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PutMonitorGroupDynamicRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s PutMonitorGroupDynamicRuleRequest) GoString() string {
	return s.String()
}

func (s *PutMonitorGroupDynamicRuleRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *PutMonitorGroupDynamicRuleRequest) GetGroupRules() []*PutMonitorGroupDynamicRuleRequestGroupRules {
	return s.GroupRules
}

func (s *PutMonitorGroupDynamicRuleRequest) GetIsAsync() *bool {
	return s.IsAsync
}

func (s *PutMonitorGroupDynamicRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PutMonitorGroupDynamicRuleRequest) SetGroupId(v int64) *PutMonitorGroupDynamicRuleRequest {
	s.GroupId = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequest) SetGroupRules(v []*PutMonitorGroupDynamicRuleRequestGroupRules) *PutMonitorGroupDynamicRuleRequest {
	s.GroupRules = v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequest) SetIsAsync(v bool) *PutMonitorGroupDynamicRuleRequest {
	s.IsAsync = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequest) SetRegionId(v string) *PutMonitorGroupDynamicRuleRequest {
	s.RegionId = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequest) Validate() error {
	if s.GroupRules != nil {
		for _, item := range s.GroupRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutMonitorGroupDynamicRuleRequestGroupRules struct {
	// The cloud service to which the alert rule is applied. Valid values of N: 1 to 3. Valid values:
	//
	// 	- ecs: Elastic Compute Service (ECS)
	//
	// 	- rds: ApsaraDB RDS
	//
	// 	- slb: Server Load Balancer (SLB)
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The logical operator used between conditional expressions in the alert rule. Valid values of N: 1 to 3. Valid values:
	//
	// 	- and: The instances that meet all the conditional expressions are automatically added to the application group.
	//
	// 	- or: The instances that meet one of the conditional expressions are automatically added to the application group.
	//
	// This parameter is required.
	//
	// example:
	//
	// and
	FilterRelation *string `json:"FilterRelation,omitempty" xml:"FilterRelation,omitempty"`
	// None.
	//
	// This parameter is required.
	Filters []*PutMonitorGroupDynamicRuleRequestGroupRulesFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
}

func (s PutMonitorGroupDynamicRuleRequestGroupRules) String() string {
	return dara.Prettify(s)
}

func (s PutMonitorGroupDynamicRuleRequestGroupRules) GoString() string {
	return s.String()
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRules) GetCategory() *string {
	return s.Category
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRules) GetFilterRelation() *string {
	return s.FilterRelation
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRules) GetFilters() []*PutMonitorGroupDynamicRuleRequestGroupRulesFilters {
	return s.Filters
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRules) SetCategory(v string) *PutMonitorGroupDynamicRuleRequestGroupRules {
	s.Category = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRules) SetFilterRelation(v string) *PutMonitorGroupDynamicRuleRequestGroupRules {
	s.FilterRelation = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRules) SetFilters(v []*PutMonitorGroupDynamicRuleRequestGroupRulesFilters) *PutMonitorGroupDynamicRuleRequestGroupRules {
	s.Filters = v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRules) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PutMonitorGroupDynamicRuleRequestGroupRulesFilters struct {
	// The method that is used to filter instances. Valid values of N: 1 to 3. Valid values:
	//
	// 	- contains: contains
	//
	// 	- notContains: does not contain
	//
	// 	- startWith: starts with a prefix
	//
	// 	- endWith: ends with a suffix
	//
	// This parameter is required.
	//
	// example:
	//
	// contains
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The name of the field based on which instances are filtered. Valid values of N: 1 to 3.
	//
	// Only hostnames are supported. Example: hostName.
	//
	// This parameter is required.
	//
	// example:
	//
	// hostName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value to be matched with the specified field. Valid values of N: 1 to 3.
	//
	// This parameter is required.
	//
	// example:
	//
	// nginx
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutMonitorGroupDynamicRuleRequestGroupRulesFilters) String() string {
	return dara.Prettify(s)
}

func (s PutMonitorGroupDynamicRuleRequestGroupRulesFilters) GoString() string {
	return s.String()
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRulesFilters) GetFunction() *string {
	return s.Function
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRulesFilters) GetName() *string {
	return s.Name
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRulesFilters) GetValue() *string {
	return s.Value
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRulesFilters) SetFunction(v string) *PutMonitorGroupDynamicRuleRequestGroupRulesFilters {
	s.Function = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRulesFilters) SetName(v string) *PutMonitorGroupDynamicRuleRequestGroupRulesFilters {
	s.Name = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRulesFilters) SetValue(v string) *PutMonitorGroupDynamicRuleRequestGroupRulesFilters {
	s.Value = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequestGroupRulesFilters) Validate() error {
	return dara.Validate(s)
}

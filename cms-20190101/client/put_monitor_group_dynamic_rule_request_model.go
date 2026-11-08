// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutMonitorGroupDynamicRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *PutMonitorGroupDynamicRuleRequest
	GetClientToken() *string
	SetDryRun(v bool) *PutMonitorGroupDynamicRuleRequest
	GetDryRun() *bool
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
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to only precheck this request. Valid values:
	//
	// - **true**: sends a check request. No security group is associated with the instance. The check items include whether required parameters are specified, the request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. If the check succeeds, an HTTP 2xx status code is returned and the operation is directly performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
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
	// The mode for creating the dynamic matching rule. Valid values:
	//
	// - true: Asynchronous creation.
	//
	// - false (default): Synchronous creation.
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

func (s *PutMonitorGroupDynamicRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *PutMonitorGroupDynamicRuleRequest) GetDryRun() *bool {
	return s.DryRun
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

func (s *PutMonitorGroupDynamicRuleRequest) SetClientToken(v string) *PutMonitorGroupDynamicRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *PutMonitorGroupDynamicRuleRequest) SetDryRun(v bool) *PutMonitorGroupDynamicRuleRequest {
	s.DryRun = &v
	return s
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
	// The type of Alibaba Cloud service for the dynamic alert rule. Valid values of N: 1 to 3. Valid values:
	//
	// - ecs: Elastic Compute Service (ECS) server.
	//
	// - rds: relational database.
	//
	// - slb: load balancing.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The combination condition for dynamic matching rules. Valid values of N: 1 to 3. Valid values:
	//
	// - and: Only instances that meet all matching rules are automatically added to the application group.
	//
	// - or: Instances that meet any matching rule are automatically added to the application group.
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
	// The filter condition for instances. Valid values of N: 1 to 3. Valid values:
	//
	// - contains: Contains.
	//
	// - notContains: Does not contain.
	//
	// - startWith: Prefix.
	//
	// - endWith: Suffix.
	//
	// This parameter is required.
	//
	// example:
	//
	// contains
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The field name for instance matching. Valid values of N: 1 to 3.
	//
	// Currently, only the hostname is supported, such as hostName.
	//
	// This parameter is required.
	//
	// example:
	//
	// hostName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value that meets the matching condition. Valid values of N: 1 to 3.
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

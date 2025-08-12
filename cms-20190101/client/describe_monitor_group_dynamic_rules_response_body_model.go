// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupDynamicRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeMonitorGroupDynamicRulesResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeMonitorGroupDynamicRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMonitorGroupDynamicRulesResponseBody
	GetRequestId() *string
	SetResource(v *DescribeMonitorGroupDynamicRulesResponseBodyResource) *DescribeMonitorGroupDynamicRulesResponseBody
	GetResource() *DescribeMonitorGroupDynamicRulesResponseBodyResource
	SetSuccess(v bool) *DescribeMonitorGroupDynamicRulesResponseBody
	GetSuccess() *bool
}

type DescribeMonitorGroupDynamicRulesResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2170B94A-1576-4D65-900E-2093037CDAF3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resources that are associated with the application group.
	Resource *DescribeMonitorGroupDynamicRulesResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitorGroupDynamicRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) GetResource() *DescribeMonitorGroupDynamicRulesResponseBodyResource {
	return s.Resource
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) SetCode(v int32) *DescribeMonitorGroupDynamicRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) SetMessage(v string) *DescribeMonitorGroupDynamicRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) SetRequestId(v string) *DescribeMonitorGroupDynamicRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) SetResource(v *DescribeMonitorGroupDynamicRulesResponseBodyResource) *DescribeMonitorGroupDynamicRulesResponseBody {
	s.Resource = v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) SetSuccess(v bool) *DescribeMonitorGroupDynamicRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupDynamicRulesResponseBodyResource struct {
	Resource []*DescribeMonitorGroupDynamicRulesResponseBodyResourceResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResource) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResource) GetResource() []*DescribeMonitorGroupDynamicRulesResponseBodyResourceResource {
	return s.Resource
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResource) SetResource(v []*DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) *DescribeMonitorGroupDynamicRulesResponseBodyResource {
	s.Resource = v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResource) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupDynamicRulesResponseBodyResourceResource struct {
	// The type of the cloud service to which the dynamic rule belongs. Valid values:
	//
	// 	- ecs: Elastic Compute Service (ECS)
	//
	// 	- rds: ApsaraDB RDS
	//
	// 	- slb: Server Load Balancer (SLB)
	//
	// example:
	//
	// ecs
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The filter condition. Valid values:
	//
	// 	- and: queries the instances that meet all alert rules.
	//
	// 	- or: queries the instances that meet any alert rule.
	//
	// example:
	//
	// and
	FilterRelation *string `json:"FilterRelation,omitempty" xml:"FilterRelation,omitempty"`
	// The dynamic rules of the application group.
	Filters *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Struct"`
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) GetCategory() *string {
	return s.Category
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) GetFilterRelation() *string {
	return s.FilterRelation
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) GetFilters() *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters {
	return s.Filters
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) SetCategory(v string) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource {
	s.Category = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) SetFilterRelation(v string) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource {
	s.FilterRelation = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) SetFilters(v *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource {
	s.Filters = v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResource) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters struct {
	Filter []*DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters) GetFilter() []*DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter {
	return s.Filter
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters) SetFilter(v []*DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters {
	s.Filter = v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFilters) Validate() error {
	return dara.Validate(s)
}

type DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter struct {
	// The method that is used to filter the instances. Valid values:
	//
	// 	- contains: contains
	//
	// 	- startWith: starts with a prefix
	//
	// 	- endWith: ends with a suffix
	//
	// example:
	//
	// contains
	Function *string `json:"Function,omitempty" xml:"Function,omitempty"`
	// The instance name.
	//
	// example:
	//
	// hostName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the dynamic rule.
	//
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) GetFunction() *string {
	return s.Function
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) GetName() *string {
	return s.Name
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) SetFunction(v string) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter {
	s.Function = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) SetName(v string) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter {
	s.Name = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) SetValue(v string) *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter {
	s.Value = &v
	return s
}

func (s *DescribeMonitorGroupDynamicRulesResponseBodyResourceResourceFiltersFilter) Validate() error {
	return dara.Validate(s)
}

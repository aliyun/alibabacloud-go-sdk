// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeResourceGroupsResponseBody
	GetRequestId() *string
	SetResourceGroup(v []*DescribeResourceGroupsResponseBodyResourceGroup) *DescribeResourceGroupsResponseBody
	GetResourceGroup() []*DescribeResourceGroupsResponseBodyResourceGroup
	SetTotalCount(v string) *DescribeResourceGroupsResponseBody
	GetTotalCount() *string
}

type DescribeResourceGroupsResponseBody struct {
	// example:
	//
	// 68BA1DF7-8814-5AED-B844-F8F7F7****
	RequestId     *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroup []*DescribeResourceGroupsResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Repeated"`
	// example:
	//
	// 7
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeResourceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceGroupsResponseBody) GetResourceGroup() []*DescribeResourceGroupsResponseBodyResourceGroup {
	return s.ResourceGroup
}

func (s *DescribeResourceGroupsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeResourceGroupsResponseBody) SetRequestId(v string) *DescribeResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceGroupsResponseBody) SetResourceGroup(v []*DescribeResourceGroupsResponseBodyResourceGroup) *DescribeResourceGroupsResponseBody {
	s.ResourceGroup = v
	return s
}

func (s *DescribeResourceGroupsResponseBody) SetTotalCount(v string) *DescribeResourceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeResourceGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceGroupsResponseBodyResourceGroup struct {
	AppRules []*DescribeResourceGroupsResponseBodyResourceGroupAppRules `json:"AppRules,omitempty" xml:"AppRules,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	AuthCount *string `json:"AuthCount,omitempty" xml:"AuthCount,omitempty"`
	// example:
	//
	// 2022-11-29T17:25:40.000000000Z
	CreateTime *string                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Policies   []*DescribeResourceGroupsResponseBodyResourceGroupPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// example:
	//
	// 119
	ResourceCount *string `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// example:
	//
	// rg-cyo0il2pzge1***
	ResourceGroupId   *string                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceGroupName *string                                                  `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	Timers            []*DescribeResourceGroupsResponseBodyResourceGroupTimers `json:"Timers,omitempty" xml:"Timers,omitempty" type:"Repeated"`
}

func (s DescribeResourceGroupsResponseBodyResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupsResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) GetAppRules() []*DescribeResourceGroupsResponseBodyResourceGroupAppRules {
	return s.AppRules
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) GetAuthCount() *string {
	return s.AuthCount
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) GetPolicies() []*DescribeResourceGroupsResponseBodyResourceGroupPolicies {
	return s.Policies
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) GetResourceCount() *string {
	return s.ResourceCount
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) GetTimers() []*DescribeResourceGroupsResponseBodyResourceGroupTimers {
	return s.Timers
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetAppRules(v []*DescribeResourceGroupsResponseBodyResourceGroupAppRules) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.AppRules = v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetAuthCount(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.AuthCount = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetCreateTime(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetPolicies(v []*DescribeResourceGroupsResponseBodyResourceGroupPolicies) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.Policies = v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetResourceCount(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.ResourceCount = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetResourceGroupId(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetResourceGroupName(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetTimers(v []*DescribeResourceGroupsResponseBodyResourceGroupTimers) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.Timers = v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceGroupsResponseBodyResourceGroupAppRules struct {
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Type *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeResourceGroupsResponseBodyResourceGroupAppRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupsResponseBodyResourceGroupAppRules) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupAppRules) GetId() *string {
	return s.Id
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupAppRules) GetName() *string {
	return s.Name
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupAppRules) GetType() *int32 {
	return s.Type
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupAppRules) SetId(v string) *DescribeResourceGroupsResponseBodyResourceGroupAppRules {
	s.Id = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupAppRules) SetName(v string) *DescribeResourceGroupsResponseBodyResourceGroupAppRules {
	s.Name = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupAppRules) SetType(v int32) *DescribeResourceGroupsResponseBodyResourceGroupAppRules {
	s.Type = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupAppRules) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceGroupsResponseBodyResourceGroupPolicies struct {
	// example:
	//
	// pl-a8jnatl8kjasb***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// False
	IsDefault *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeResourceGroupsResponseBodyResourceGroupPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupsResponseBodyResourceGroupPolicies) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupPolicies) GetId() *string {
	return s.Id
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupPolicies) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupPolicies) GetName() *string {
	return s.Name
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupPolicies) SetId(v string) *DescribeResourceGroupsResponseBodyResourceGroupPolicies {
	s.Id = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupPolicies) SetIsDefault(v bool) *DescribeResourceGroupsResponseBodyResourceGroupPolicies {
	s.IsDefault = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupPolicies) SetName(v string) *DescribeResourceGroupsResponseBodyResourceGroupPolicies {
	s.Name = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupPolicies) Validate() error {
	return dara.Validate(s)
}

type DescribeResourceGroupsResponseBodyResourceGroupTimers struct {
	BindStatus *string `json:"BindStatus,omitempty" xml:"BindStatus,omitempty"`
	// example:
	//
	// t-asdzx0mbjhg***
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	TimerStatus *string `json:"TimerStatus,omitempty" xml:"TimerStatus,omitempty"`
}

func (s DescribeResourceGroupsResponseBodyResourceGroupTimers) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupsResponseBodyResourceGroupTimers) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) GetBindStatus() *string {
	return s.BindStatus
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) GetId() *string {
	return s.Id
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) GetName() *string {
	return s.Name
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) GetTimerStatus() *string {
	return s.TimerStatus
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) SetBindStatus(v string) *DescribeResourceGroupsResponseBodyResourceGroupTimers {
	s.BindStatus = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) SetId(v string) *DescribeResourceGroupsResponseBodyResourceGroupTimers {
	s.Id = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) SetName(v string) *DescribeResourceGroupsResponseBodyResourceGroupTimers {
	s.Name = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) SetTimerStatus(v string) *DescribeResourceGroupsResponseBodyResourceGroupTimers {
	s.TimerStatus = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) Validate() error {
	return dara.Validate(s)
}

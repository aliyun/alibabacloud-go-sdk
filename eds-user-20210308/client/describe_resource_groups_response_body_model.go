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
	// The ID of the request.
	//
	// example:
	//
	// 68BA1DF7-8814-5AED-B844-F8F7F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of resource group information.
	ResourceGroup []*DescribeResourceGroupsResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Repeated"`
	// The total number of resource groups.
	//
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
	if s.ResourceGroup != nil {
		for _, item := range s.ResourceGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourceGroupsResponseBodyResourceGroup struct {
	// The agent subtype of the DesktopAgent resource group. Valid values: JVSClaw and HermesAgent.
	//
	// example:
	//
	// JVSClaw
	AgentType *string `json:"AgentType,omitempty" xml:"AgentType,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-123456
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The list of application rules associated with the resource group.
	AppRules []*DescribeResourceGroupsResponseBodyResourceGroupAppRules `json:"AppRules,omitempty" xml:"AppRules,omitempty" type:"Repeated"`
	// The number of administrators authorized for the resource group.
	//
	// example:
	//
	// 3
	AuthCount *string `json:"AuthCount,omitempty" xml:"AuthCount,omitempty"`
	// The creation time in ISO 8601 UTC format.
	//
	// example:
	//
	// 2022-11-29T17:25:40.000000000Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// > The policies associated with the resource group.
	//
	// > - The associated policies take effect on the cloud computers in the resource group. If multiple policies exist, they take effect based on their priorities.
	//
	// > - If a cloud computer in the resource group is already assigned another policy, the policy associated with the resource group takes precedence.
	Policies []*DescribeResourceGroupsResponseBodyResourceGroupPolicies `json:"Policies,omitempty" xml:"Policies,omitempty" type:"Repeated"`
	// The resource count in the resource group.
	//
	// example:
	//
	// 119
	ResourceCount *string `json:"ResourceCount,omitempty" xml:"ResourceCount,omitempty"`
	// The description of the resource group. If not set, an empty string is returned.
	//
	// example:
	//
	// Department A resource group
	ResourceGroupDescription *string `json:"ResourceGroupDescription,omitempty" xml:"ResourceGroupDescription,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-cyo0il2pzge1***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// Test resource group
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// > The scheduled tasks associated with the resource group.
	//
	// > - The associated scheduled tasks take effect on the cloud computers in the resource group. If a cloud computer in the resource group is already associated with another scheduled task, the task associated with the resource group takes precedence.
	Timers []*DescribeResourceGroupsResponseBodyResourceGroupTimers `json:"Timers,omitempty" xml:"Timers,omitempty" type:"Repeated"`
}

func (s DescribeResourceGroupsResponseBodyResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceGroupsResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) GetAgentType() *string {
	return s.AgentType
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
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

func (s *DescribeResourceGroupsResponseBodyResourceGroup) GetResourceGroupDescription() *string {
	return s.ResourceGroupDescription
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

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetAgentType(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.AgentType = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetAliyunResourceGroupId(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.AliyunResourceGroupId = &v
	return s
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

func (s *DescribeResourceGroupsResponseBodyResourceGroup) SetResourceGroupDescription(v string) *DescribeResourceGroupsResponseBodyResourceGroup {
	s.ResourceGroupDescription = &v
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
	if s.AppRules != nil {
		for _, item := range s.AppRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Policies != nil {
		for _, item := range s.Policies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Timers != nil {
		for _, item := range s.Timers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResourceGroupsResponseBodyResourceGroupAppRules struct {
	// The ID of the application rule.
	//
	// example:
	//
	// ruleId1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the application rule.
	//
	// example:
	//
	// ruleName1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the application rule.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// The ID of the policy.
	//
	// example:
	//
	// pl-a8jnatl8kjasb***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether the policy is the default policy.
	//
	// example:
	//
	// False
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The policy name.
	//
	// example:
	//
	// Test policy 123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The association status between the scheduled task and the resource group.
	//
	// example:
	//
	// BIND
	BindStatus *string `json:"BindStatus,omitempty" xml:"BindStatus,omitempty"`
	// The ID of the scheduled task.
	//
	// example:
	//
	// t-asdzx0mbjhg***
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the scheduled task.
	//
	// example:
	//
	// Test task 123
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The running status of the scheduled task.
	//
	// example:
	//
	// AVAILABLE
	TimerStatus *string `json:"TimerStatus,omitempty" xml:"TimerStatus,omitempty"`
	// The type of the scheduled task.
	//
	// example:
	//
	// TIMER
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) GetType() *string {
	return s.Type
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

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) SetType(v string) *DescribeResourceGroupsResponseBodyResourceGroupTimers {
	s.Type = &v
	return s
}

func (s *DescribeResourceGroupsResponseBodyResourceGroupTimers) Validate() error {
	return dara.Validate(s)
}

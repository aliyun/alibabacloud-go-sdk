// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListStackGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStackGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListStackGroupsResponseBody
	GetRequestId() *string
	SetStackGroups(v []*ListStackGroupsResponseBodyStackGroups) *ListStackGroupsResponseBody
	GetStackGroups() []*ListStackGroupsResponseBodyStackGroups
	SetTotalCount(v int32) *ListStackGroupsResponseBody
	GetTotalCount() *int32
}

type ListStackGroupsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The stack groups.
	StackGroups []*ListStackGroupsResponseBodyStackGroups `json:"StackGroups,omitempty" xml:"StackGroups,omitempty" type:"Repeated"`
	// The total number of stack groups.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStackGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStackGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStackGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStackGroupsResponseBody) GetStackGroups() []*ListStackGroupsResponseBodyStackGroups {
	return s.StackGroups
}

func (s *ListStackGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListStackGroupsResponseBody) SetPageNumber(v int32) *ListStackGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupsResponseBody) SetPageSize(v int32) *ListStackGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupsResponseBody) SetRequestId(v string) *ListStackGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStackGroupsResponseBody) SetStackGroups(v []*ListStackGroupsResponseBodyStackGroups) *ListStackGroupsResponseBody {
	s.StackGroups = v
	return s
}

func (s *ListStackGroupsResponseBody) SetTotalCount(v int32) *ListStackGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStackGroupsResponseBody) Validate() error {
	if s.StackGroups != nil {
		for _, item := range s.StackGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStackGroupsResponseBodyStackGroups struct {
	// The information about automatic deployment settings.
	AutoDeployment *ListStackGroupsResponseBodyStackGroupsAutoDeployment `json:"AutoDeployment,omitempty" xml:"AutoDeployment,omitempty" type:"Struct"`
	CreateTime     *string                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the stack group.
	//
	// example:
	//
	// My Stack Group
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the most recent successful drift detection was performed on the stack group.
	//
	// example:
	//
	// 2020-02-27T07:47:47
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The permission model of the stack group.
	//
	// Valid values:
	//
	// 	- SELF_MANAGED
	//
	// 	- SERVICE_MANAGED
	//
	// > For more information about the permission models of stack groups, see [Overview](https://help.aliyun.com/document_detail/154578.html).
	//
	// example:
	//
	// SELF_MANAGED
	PermissionModel *string `json:"PermissionModel,omitempty" xml:"PermissionModel,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxcj****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The drift state of the stack group on which the most recent successful drift detection was performed.
	//
	// Valid values:
	//
	// 	- DRIFTED: The stack group has drifted.
	//
	// 	- NOT_CHECKED: No drift detection is performed on the stack group.
	//
	// 	- IN_SYNC: No drifts are detected on the stack group.
	//
	// example:
	//
	// IN_SYNC
	StackGroupDriftStatus *string `json:"StackGroupDriftStatus,omitempty" xml:"StackGroupDriftStatus,omitempty"`
	// The ID of the stack group.
	//
	// example:
	//
	// fd0ddef9-9540-4b42-a464-94f77835****
	StackGroupId *string `json:"StackGroupId,omitempty" xml:"StackGroupId,omitempty"`
	// The name of the stack group.
	//
	// example:
	//
	// MyStackGroup
	StackGroupName *string `json:"StackGroupName,omitempty" xml:"StackGroupName,omitempty"`
	// The state of the stack group.
	//
	// Valid values:
	//
	// 	- ACTIVE
	//
	// 	- DELETED
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the stack group.
	Tags       []*ListStackGroupsResponseBodyStackGroupsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UpdateTime *string                                       `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListStackGroupsResponseBodyStackGroups) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupsResponseBodyStackGroups) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponseBodyStackGroups) GetAutoDeployment() *ListStackGroupsResponseBodyStackGroupsAutoDeployment {
	return s.AutoDeployment
}

func (s *ListStackGroupsResponseBodyStackGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStackGroupsResponseBodyStackGroups) GetDescription() *string {
	return s.Description
}

func (s *ListStackGroupsResponseBodyStackGroups) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *ListStackGroupsResponseBodyStackGroups) GetPermissionModel() *string {
	return s.PermissionModel
}

func (s *ListStackGroupsResponseBodyStackGroups) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListStackGroupsResponseBodyStackGroups) GetStackGroupDriftStatus() *string {
	return s.StackGroupDriftStatus
}

func (s *ListStackGroupsResponseBodyStackGroups) GetStackGroupId() *string {
	return s.StackGroupId
}

func (s *ListStackGroupsResponseBodyStackGroups) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *ListStackGroupsResponseBodyStackGroups) GetStatus() *string {
	return s.Status
}

func (s *ListStackGroupsResponseBodyStackGroups) GetTags() []*ListStackGroupsResponseBodyStackGroupsTags {
	return s.Tags
}

func (s *ListStackGroupsResponseBodyStackGroups) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListStackGroupsResponseBodyStackGroups) SetAutoDeployment(v *ListStackGroupsResponseBodyStackGroupsAutoDeployment) *ListStackGroupsResponseBodyStackGroups {
	s.AutoDeployment = v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetCreateTime(v string) *ListStackGroupsResponseBodyStackGroups {
	s.CreateTime = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetDescription(v string) *ListStackGroupsResponseBodyStackGroups {
	s.Description = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetDriftDetectionTime(v string) *ListStackGroupsResponseBodyStackGroups {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetPermissionModel(v string) *ListStackGroupsResponseBodyStackGroups {
	s.PermissionModel = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetResourceGroupId(v string) *ListStackGroupsResponseBodyStackGroups {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStackGroupDriftStatus(v string) *ListStackGroupsResponseBodyStackGroups {
	s.StackGroupDriftStatus = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStackGroupId(v string) *ListStackGroupsResponseBodyStackGroups {
	s.StackGroupId = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStackGroupName(v string) *ListStackGroupsResponseBodyStackGroups {
	s.StackGroupName = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetStatus(v string) *ListStackGroupsResponseBodyStackGroups {
	s.Status = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetTags(v []*ListStackGroupsResponseBodyStackGroupsTags) *ListStackGroupsResponseBodyStackGroups {
	s.Tags = v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) SetUpdateTime(v string) *ListStackGroupsResponseBodyStackGroups {
	s.UpdateTime = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroups) Validate() error {
	if s.AutoDeployment != nil {
		if err := s.AutoDeployment.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStackGroupsResponseBodyStackGroupsAutoDeployment struct {
	// Indicates whether automatic deployment is enabled.
	//
	// Valid values:
	//
	// 	- true: Automatic deployment is enabled. If you add a member to the folder to which the stack group belongs after automatic deployment is enabled, Resource Orchestration Service (ROS) automatically adds the stack instances in the stack group to the specified region of the member. If you delete the member from the folder, ROS automatically deletes the stack instances in the stack group from the specified region of the member.
	//
	// 	- false: Automatic deployment is disabled. After you disable automatic deployment, the stack instances remain unchanged when you change the member in the folder.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether the stacks within a member are retained when you delete the member from the folder.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// > This parameter is returned only if Enabled is set to true.
	//
	// example:
	//
	// true
	RetainStacksOnAccountRemoval *bool `json:"RetainStacksOnAccountRemoval,omitempty" xml:"RetainStacksOnAccountRemoval,omitempty"`
}

func (s ListStackGroupsResponseBodyStackGroupsAutoDeployment) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupsResponseBodyStackGroupsAutoDeployment) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponseBodyStackGroupsAutoDeployment) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListStackGroupsResponseBodyStackGroupsAutoDeployment) GetRetainStacksOnAccountRemoval() *bool {
	return s.RetainStacksOnAccountRemoval
}

func (s *ListStackGroupsResponseBodyStackGroupsAutoDeployment) SetEnabled(v bool) *ListStackGroupsResponseBodyStackGroupsAutoDeployment {
	s.Enabled = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroupsAutoDeployment) SetRetainStacksOnAccountRemoval(v bool) *ListStackGroupsResponseBodyStackGroupsAutoDeployment {
	s.RetainStacksOnAccountRemoval = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroupsAutoDeployment) Validate() error {
	return dara.Validate(s)
}

type ListStackGroupsResponseBodyStackGroupsTags struct {
	// The key of the tag that is added to the stack group.
	//
	// example:
	//
	// usage1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is added to the stack group.
	//
	// example:
	//
	// test1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListStackGroupsResponseBodyStackGroupsTags) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupsResponseBodyStackGroupsTags) GoString() string {
	return s.String()
}

func (s *ListStackGroupsResponseBodyStackGroupsTags) GetKey() *string {
	return s.Key
}

func (s *ListStackGroupsResponseBodyStackGroupsTags) GetValue() *string {
	return s.Value
}

func (s *ListStackGroupsResponseBodyStackGroupsTags) SetKey(v string) *ListStackGroupsResponseBodyStackGroupsTags {
	s.Key = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroupsTags) SetValue(v string) *ListStackGroupsResponseBodyStackGroupsTags {
	s.Value = &v
	return s
}

func (s *ListStackGroupsResponseBodyStackGroupsTags) Validate() error {
	return dara.Validate(s)
}

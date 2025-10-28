// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStacksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListStacksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListStacksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListStacksResponseBody
	GetRequestId() *string
	SetStacks(v []*ListStacksResponseBodyStacks) *ListStacksResponseBody
	GetStacks() []*ListStacksResponseBodyStacks
	SetTotalCount(v int32) *ListStacksResponseBody
	GetTotalCount() *int32
}

type ListStacksResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FBAC80B4-9C27-529D-BC9C-4155FA5CD7A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the stacks.
	Stacks []*ListStacksResponseBodyStacks `json:"Stacks,omitempty" xml:"Stacks,omitempty" type:"Repeated"`
	// The total number of stacks.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListStacksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListStacksResponseBody) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListStacksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListStacksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStacksResponseBody) GetStacks() []*ListStacksResponseBodyStacks {
	return s.Stacks
}

func (s *ListStacksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListStacksResponseBody) SetPageNumber(v int32) *ListStacksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListStacksResponseBody) SetPageSize(v int32) *ListStacksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListStacksResponseBody) SetRequestId(v string) *ListStacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStacksResponseBody) SetStacks(v []*ListStacksResponseBodyStacks) *ListStacksResponseBody {
	s.Stacks = v
	return s
}

func (s *ListStacksResponseBody) SetTotalCount(v int32) *ListStacksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListStacksResponseBody) Validate() error {
	if s.Stacks != nil {
		for _, item := range s.Stacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStacksResponseBodyStacks struct {
	// The time when the stack was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-10T06:44:36
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether deletion protection is enabled for the stack. Valid values:
	//
	// 	- Enabled: Deletion protection is enabled for the stack.
	//
	// 	- Disabled: Deletion protection is disabled for the stack. In this case, you can delete the stack by using the console or calling the [DeleteStack](https://help.aliyun.com/document_detail/610812.html) operation.
	//
	// >  Deletion protection of a nested stack is the same as that of its root stack.
	//
	// example:
	//
	// Disabled
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// Indicates whether rollback is disabled when the stack fails to be created. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	DisableRollback *bool `json:"DisableRollback,omitempty" xml:"DisableRollback,omitempty"`
	// The time when the most recent successful drift detection was performed on the stack.
	//
	// example:
	//
	// 2022-03-10T06:46:36
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The supplementary information that is returned if an error occurs on a stack operation.
	//
	// >  This parameter is returned only under specific conditions, and is returned together with at least one sub-parameter. For example, an error occurred when an API operation of another Alibaba Cloud service was called.
	OperationInfo *ListStacksResponseBodyStacksOperationInfo `json:"OperationInfo,omitempty" xml:"OperationInfo,omitempty" type:"Struct"`
	// The ID of the parent stack.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf692****
	ParentStackId *string `json:"ParentStackId,omitempty" xml:"ParentStackId,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2frunvw7****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the stack is a managed stack. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The name of the service to which the managed stack belongs.
	//
	// example:
	//
	// ACVS
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The state of the stack on which the most recent successful drift detection was performed. Valid values:
	//
	// 	- DRIFTED: The stack has drifted.
	//
	// 	- NOT_CHECKED: No successful drift detection is performed on the stack.
	//
	// 	- IN_SYNC: The stack is being synchronized.
	//
	// example:
	//
	// IN_SYNC
	StackDriftStatus *string `json:"StackDriftStatus,omitempty" xml:"StackDriftStatus,omitempty"`
	// The stack ID.
	//
	// example:
	//
	// 67805444-a605-45ee-a57f-83908ff6****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The stack name.
	//
	// example:
	//
	// MyStack
	StackName *string `json:"StackName,omitempty" xml:"StackName,omitempty"`
	// The stack type. Valid values:
	//
	// 	- ROS: ROS stack. The stack is created by using a ROS template.
	//
	// 	- Terraform: Terraform stack. The stack is created by using a Terraform template.
	//
	// example:
	//
	// ROS
	StackType *string `json:"StackType,omitempty" xml:"StackType,omitempty"`
	// The state of the stack.
	//
	// example:
	//
	// CREATE_COMPLETE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The reason why the stack is in its current state.
	//
	// example:
	//
	// Stack CREATE completed successfully
	StatusReason *string `json:"StatusReason,omitempty" xml:"StatusReason,omitempty"`
	// The tags of the stack.
	Tags []*ListStacksResponseBodyStacksTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The timeout period for creating the stack. Unit: minutes.
	//
	// example:
	//
	// 60
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	// The time when the stack was updated. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-03-10T07:44:36
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListStacksResponseBodyStacks) String() string {
	return dara.Prettify(s)
}

func (s ListStacksResponseBodyStacks) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBodyStacks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListStacksResponseBodyStacks) GetDeletionProtection() *string {
	return s.DeletionProtection
}

func (s *ListStacksResponseBodyStacks) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *ListStacksResponseBodyStacks) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *ListStacksResponseBodyStacks) GetOperationInfo() *ListStacksResponseBodyStacksOperationInfo {
	return s.OperationInfo
}

func (s *ListStacksResponseBodyStacks) GetParentStackId() *string {
	return s.ParentStackId
}

func (s *ListStacksResponseBodyStacks) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStacksResponseBodyStacks) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListStacksResponseBodyStacks) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *ListStacksResponseBodyStacks) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListStacksResponseBodyStacks) GetStackDriftStatus() *string {
	return s.StackDriftStatus
}

func (s *ListStacksResponseBodyStacks) GetStackId() *string {
	return s.StackId
}

func (s *ListStacksResponseBodyStacks) GetStackName() *string {
	return s.StackName
}

func (s *ListStacksResponseBodyStacks) GetStackType() *string {
	return s.StackType
}

func (s *ListStacksResponseBodyStacks) GetStatus() *string {
	return s.Status
}

func (s *ListStacksResponseBodyStacks) GetStatusReason() *string {
	return s.StatusReason
}

func (s *ListStacksResponseBodyStacks) GetTags() []*ListStacksResponseBodyStacksTags {
	return s.Tags
}

func (s *ListStacksResponseBodyStacks) GetTimeoutInMinutes() *int32 {
	return s.TimeoutInMinutes
}

func (s *ListStacksResponseBodyStacks) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListStacksResponseBodyStacks) SetCreateTime(v string) *ListStacksResponseBodyStacks {
	s.CreateTime = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetDeletionProtection(v string) *ListStacksResponseBodyStacks {
	s.DeletionProtection = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetDisableRollback(v bool) *ListStacksResponseBodyStacks {
	s.DisableRollback = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetDriftDetectionTime(v string) *ListStacksResponseBodyStacks {
	s.DriftDetectionTime = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetOperationInfo(v *ListStacksResponseBodyStacksOperationInfo) *ListStacksResponseBodyStacks {
	s.OperationInfo = v
	return s
}

func (s *ListStacksResponseBodyStacks) SetParentStackId(v string) *ListStacksResponseBodyStacks {
	s.ParentStackId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetRegionId(v string) *ListStacksResponseBodyStacks {
	s.RegionId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetResourceGroupId(v string) *ListStacksResponseBodyStacks {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetServiceManaged(v bool) *ListStacksResponseBodyStacks {
	s.ServiceManaged = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetServiceName(v string) *ListStacksResponseBodyStacks {
	s.ServiceName = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackDriftStatus(v string) *ListStacksResponseBodyStacks {
	s.StackDriftStatus = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackId(v string) *ListStacksResponseBodyStacks {
	s.StackId = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackName(v string) *ListStacksResponseBodyStacks {
	s.StackName = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStackType(v string) *ListStacksResponseBodyStacks {
	s.StackType = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStatus(v string) *ListStacksResponseBodyStacks {
	s.Status = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetStatusReason(v string) *ListStacksResponseBodyStacks {
	s.StatusReason = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetTags(v []*ListStacksResponseBodyStacksTags) *ListStacksResponseBodyStacks {
	s.Tags = v
	return s
}

func (s *ListStacksResponseBodyStacks) SetTimeoutInMinutes(v int32) *ListStacksResponseBodyStacks {
	s.TimeoutInMinutes = &v
	return s
}

func (s *ListStacksResponseBodyStacks) SetUpdateTime(v string) *ListStacksResponseBodyStacks {
	s.UpdateTime = &v
	return s
}

func (s *ListStacksResponseBodyStacks) Validate() error {
	if s.OperationInfo != nil {
		if err := s.OperationInfo.Validate(); err != nil {
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

type ListStacksResponseBodyStacksOperationInfo struct {
	// The name of the API operation that belongs to another Alibaba Cloud service.
	//
	// example:
	//
	// DeleteSecurityGroup
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The error code.
	//
	// example:
	//
	// DependencyViolation
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The logical ID of the resource on which the operation error occurred.
	//
	// example:
	//
	// EcsSecurityGroup
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The error message.
	//
	// example:
	//
	// There is still instance(s) in the specified security group.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request that is initiated to call the API operation of another Alibaba Cloud service.
	//
	// example:
	//
	// 071D6166-3F6B-5C7B-A1F0-0113FBB643A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the resource on which the operation error occurred.
	//
	// example:
	//
	// ALIYUN::ECS::SecurityGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListStacksResponseBodyStacksOperationInfo) String() string {
	return dara.Prettify(s)
}

func (s ListStacksResponseBodyStacksOperationInfo) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBodyStacksOperationInfo) GetAction() *string {
	return s.Action
}

func (s *ListStacksResponseBodyStacksOperationInfo) GetCode() *string {
	return s.Code
}

func (s *ListStacksResponseBodyStacksOperationInfo) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *ListStacksResponseBodyStacksOperationInfo) GetMessage() *string {
	return s.Message
}

func (s *ListStacksResponseBodyStacksOperationInfo) GetRequestId() *string {
	return s.RequestId
}

func (s *ListStacksResponseBodyStacksOperationInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetAction(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.Action = &v
	return s
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetCode(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.Code = &v
	return s
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetLogicalResourceId(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.LogicalResourceId = &v
	return s
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetMessage(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.Message = &v
	return s
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetRequestId(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.RequestId = &v
	return s
}

func (s *ListStacksResponseBodyStacksOperationInfo) SetResourceType(v string) *ListStacksResponseBodyStacksOperationInfo {
	s.ResourceType = &v
	return s
}

func (s *ListStacksResponseBodyStacksOperationInfo) Validate() error {
	return dara.Validate(s)
}

type ListStacksResponseBodyStacksTags struct {
	// The tag key of the stack.
	//
	// example:
	//
	// acs:rm:rgId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the stack.
	//
	// example:
	//
	// rg-aek2frunvw7****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListStacksResponseBodyStacksTags) String() string {
	return dara.Prettify(s)
}

func (s ListStacksResponseBodyStacksTags) GoString() string {
	return s.String()
}

func (s *ListStacksResponseBodyStacksTags) GetKey() *string {
	return s.Key
}

func (s *ListStacksResponseBodyStacksTags) GetValue() *string {
	return s.Value
}

func (s *ListStacksResponseBodyStacksTags) SetKey(v string) *ListStacksResponseBodyStacksTags {
	s.Key = &v
	return s
}

func (s *ListStacksResponseBodyStacksTags) SetValue(v string) *ListStacksResponseBodyStacksTags {
	s.Value = &v
	return s
}

func (s *ListStacksResponseBodyStacksTags) Validate() error {
	return dara.Validate(s)
}

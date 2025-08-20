// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckedStackResourceCount(v int32) *GetStackResponseBody
	GetCheckedStackResourceCount() *int32
	SetCreateTime(v string) *GetStackResponseBody
	GetCreateTime() *string
	SetDeletionProtection(v string) *GetStackResponseBody
	GetDeletionProtection() *string
	SetDescription(v string) *GetStackResponseBody
	GetDescription() *string
	SetDisableRollback(v bool) *GetStackResponseBody
	GetDisableRollback() *bool
	SetDriftDetectionTime(v string) *GetStackResponseBody
	GetDriftDetectionTime() *string
	SetInterface(v string) *GetStackResponseBody
	GetInterface() *string
	SetLog(v *GetStackResponseBodyLog) *GetStackResponseBody
	GetLog() *GetStackResponseBodyLog
	SetNotCheckedStackResourceCount(v int32) *GetStackResponseBody
	GetNotCheckedStackResourceCount() *int32
	SetNotificationURLs(v []*string) *GetStackResponseBody
	GetNotificationURLs() []*string
	SetOperationInfo(v *GetStackResponseBodyOperationInfo) *GetStackResponseBody
	GetOperationInfo() *GetStackResponseBodyOperationInfo
	SetOrderIds(v []*string) *GetStackResponseBody
	GetOrderIds() []*string
	SetOutputs(v []map[string]interface{}) *GetStackResponseBody
	GetOutputs() []map[string]interface{}
	SetParameters(v []*GetStackResponseBodyParameters) *GetStackResponseBody
	GetParameters() []*GetStackResponseBodyParameters
	SetParentStackId(v string) *GetStackResponseBody
	GetParentStackId() *string
	SetRamRoleName(v string) *GetStackResponseBody
	GetRamRoleName() *string
	SetRegionId(v string) *GetStackResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetStackResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetStackResponseBody
	GetResourceGroupId() *string
	SetResourceProgress(v *GetStackResponseBodyResourceProgress) *GetStackResponseBody
	GetResourceProgress() *GetStackResponseBodyResourceProgress
	SetRollbackFailedRootReason(v string) *GetStackResponseBody
	GetRollbackFailedRootReason() *string
	SetRootStackId(v string) *GetStackResponseBody
	GetRootStackId() *string
	SetServiceManaged(v bool) *GetStackResponseBody
	GetServiceManaged() *bool
	SetServiceName(v string) *GetStackResponseBody
	GetServiceName() *string
	SetStackDriftStatus(v string) *GetStackResponseBody
	GetStackDriftStatus() *string
	SetStackId(v string) *GetStackResponseBody
	GetStackId() *string
	SetStackName(v string) *GetStackResponseBody
	GetStackName() *string
	SetStackType(v string) *GetStackResponseBody
	GetStackType() *string
	SetStatus(v string) *GetStackResponseBody
	GetStatus() *string
	SetStatusReason(v string) *GetStackResponseBody
	GetStatusReason() *string
	SetTags(v []*GetStackResponseBodyTags) *GetStackResponseBody
	GetTags() []*GetStackResponseBodyTags
	SetTemplateDescription(v string) *GetStackResponseBody
	GetTemplateDescription() *string
	SetTemplateId(v string) *GetStackResponseBody
	GetTemplateId() *string
	SetTemplateScratchId(v string) *GetStackResponseBody
	GetTemplateScratchId() *string
	SetTemplateURL(v string) *GetStackResponseBody
	GetTemplateURL() *string
	SetTemplateVersion(v string) *GetStackResponseBody
	GetTemplateVersion() *string
	SetTimeoutInMinutes(v int32) *GetStackResponseBody
	GetTimeoutInMinutes() *int32
	SetUpdateTime(v string) *GetStackResponseBody
	GetUpdateTime() *string
}

type GetStackResponseBody struct {
	// The number of resources on which drift detection was performed.
	//
	// >  This parameter is returned only if the most recent drift detection on the stack was successful.
	//
	// example:
	//
	// 1
	CheckedStackResourceCount *int32 `json:"CheckedStackResourceCount,omitempty" xml:"CheckedStackResourceCount,omitempty"`
	// The time when the stack was created. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-09-16T08:21:40
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether deletion protection is enabled for the stack. Valid values:
	//
	// 	- Enabled: Deletion protection is enabled for the stack.
	//
	// 	- Disabled: Deletion protection is disabled for the stack. You can delete the stack by using the ROS console or by calling the DeleteStack operation.
	//
	// >  Deletion protection of a nested stack is the same as deletion protection of its root stack.
	//
	// example:
	//
	// Disabled
	DeletionProtection *string `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of the stack.
	//
	// example:
	//
	// Create a VPC.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	// 2020-09-16T09:21:40
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The description of the console user interface (UI).
	//
	// example:
	//
	// {}
	Interface *string `json:"Interface,omitempty" xml:"Interface,omitempty"`
	// The log of the stack.
	Log *GetStackResponseBodyLog `json:"Log,omitempty" xml:"Log,omitempty" type:"Struct"`
	// The number of resources on which drift detection was not performed.
	//
	// >  This parameter is returned only if the most recent drift detection on the stack was successful.
	//
	// example:
	//
	// 1
	NotCheckedStackResourceCount *int32 `json:"NotCheckedStackResourceCount,omitempty" xml:"NotCheckedStackResourceCount,omitempty"`
	// The callback URLs for receiving stack events.
	NotificationURLs []*string `json:"NotificationURLs,omitempty" xml:"NotificationURLs,omitempty" type:"Repeated"`
	// The supplementary information that is returned if an error occurs on a stack operation.
	//
	// >  This parameter is returned together with at least one sub-parameter and only under specific conditions. For example, the supplementary information is returned when an API operation of another Alibaba Cloud service fails to be called.
	OperationInfo *GetStackResponseBodyOperationInfo `json:"OperationInfo,omitempty" xml:"OperationInfo,omitempty" type:"Struct"`
	// The order IDs. This parameter is returned only if you configured manual payment when you created a subscription stack.
	OrderIds []*string `json:"OrderIds,omitempty" xml:"OrderIds,omitempty" type:"Repeated"`
	// The outputs of the stack.
	Outputs []map[string]interface{} `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
	// The parameters of the stack.
	Parameters []*GetStackResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The ID of the parent stack.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf692****
	ParentStackId *string `json:"ParentStackId,omitempty" xml:"ParentStackId,omitempty"`
	// The name of the Resource Access Management (RAM) role. ROS assumes the RAM role to create the stack and uses the credentials of the role to call the APIs of Alibaba Cloud services.\\
	//
	// ROS assumes the RAM role to perform operations on the stack. If you have permissions to perform operations on the stack, ROS assumes the RAM role even if you do not have permissions to use the RAM role. You must make sure that permissions are granted to the RAM role based on the principle of least privilege.\\
	//
	// If this parameter is not specified, ROS uses the existing role that is associated with the stack. If no roles are available, ROS uses a temporary credential that is generated from the credentials of your account.\\
	//
	// The RAM role name can be up to 64 characters in length.
	//
	// example:
	//
	// test-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource creation progress.
	ResourceProgress *GetStackResponseBodyResourceProgress `json:"ResourceProgress,omitempty" xml:"ResourceProgress,omitempty" type:"Struct"`
	// 当资源栈状态为回滚失败时，该字段展示导致回滚的前一阶段执行失败的原因。
	//
	// example:
	//
	// Resource UPDATE failed: Exception: resources.FailToCreate: FailToCreate: reason
	RollbackFailedRootReason *string `json:"RollbackFailedRootReason,omitempty" xml:"RollbackFailedRootReason,omitempty"`
	// The ID of the root stack. This parameter is returned if the specified stack is a nested stack.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf692****
	RootStackId *string `json:"RootStackId,omitempty" xml:"RootStackId,omitempty"`
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
	// c754d2a4-28f1-46df-b557-9586173a****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
	// The stack name.\\
	//
	// The name can be up to 255 characters in length, and can contain digits, letters, hyphens (-), and underscores (_). The name must start with a digit or letter.
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
	// The state of the stack. Valid values:
	//
	// 	- CREATE_IN_PROGRESS: The stack is being created.
	//
	// 	- CREATE_FAILED: The stack failed to be created.
	//
	// 	- CREATE_COMPLETE: The stack is created.
	//
	// 	- UPDATE_IN_PROGRESS: The stack is being updated.
	//
	// 	- UPDATE_FAILED: The stack failed to be updated.
	//
	// 	- UPDATE_COMPLETE: The stack is updated.
	//
	// 	- DELETE_IN_PROGRESS: The stack is being deleted.
	//
	// 	- DELETE_FAILED: The stack failed to be deleted.
	//
	// 	- CREATE_ROLLBACK_IN_PROGRESS: The resources are being rolled back after the stack failed to be created.
	//
	// 	- CREATE_ROLLBACK_FAILED: The resources failed to be rolled back after the stack failed to be created.
	//
	// 	- CREATE_ROLLBACK_COMPLETE: The resources are rolled back after the stack failed to be created.
	//
	// 	- ROLLBACK_IN_PROGRESS: The resources of the stack are being rolled back.
	//
	// 	- ROLLBACK_FAILED: The resources of the stack failed to be rolled back.
	//
	// 	- ROLLBACK_COMPLETE: The resources of the stack are rolled back.
	//
	// 	- CHECK_IN_PROGRESS: The stack is being validated.
	//
	// 	- CHECK_FAILED: The stack failed to be validated.
	//
	// 	- CHECK_COMPLETE: The stack is validated.
	//
	// 	- REVIEW_IN_PROGRESS: The stack is being reviewed.
	//
	// 	- IMPORT_CREATE_IN_PROGRESS: The stack is being created by using imported resources.
	//
	// 	- IMPORT_CREATE_FAILED: The stack failed to be created by using imported resources.
	//
	// 	- IMPORT_CREATE_COMPLETE: The stack is created by using imported resources.
	//
	// 	- IMPORT_CREATE_ROLLBACK_IN_PROGRESS: The resources are being rolled back after the stack failed to be created by using imported resources.
	//
	// 	- IMPORT_CREATE_ROLLBACK_FAILED: The resources failed to be rolled back after the stack failed to be created by using imported resources.
	//
	// 	- IMPORT_CREATE_ROLLBACK_COMPLETE: The resources are rolled back after the stack failed to be created by using imported resources.
	//
	// 	- IMPORT_UPDATE_IN_PROGRESS: The stack is being updated by using imported resources.
	//
	// 	- IMPORT_UPDATE_FAILED: The stack failed to be updated by using imported resources.
	//
	// 	- IMPORT_UPDATE_COMPLETE: The stack is updated by using imported resources.
	//
	// 	- IMPORT_UPDATE_ROLLBACK_IN_PROGRESS: The resources are being rolled back after the stack failed to be updated by using imported resources.
	//
	// 	- IMPORT_UPDATE_ROLLBACK_FAILED: The resources failed to be rolled back after the stack failed to be updated by using imported resources.
	//
	// 	- IMPORT_UPDATE_ROLLBACK_COMPLETE: The resources are rolled back after the stack failed to be updated by using imported resources.
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
	Tags []*GetStackResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The description of the template.
	//
	// example:
	//
	// Create a VPC.
	TemplateDescription *string `json:"TemplateDescription,omitempty" xml:"TemplateDescription,omitempty"`
	// The template ID. This parameter is returned only if the current stack template is a custom template or shared template.
	//
	// If the template is a shared template, the value of this parameter is the same as the value of TemplateARN.
	//
	// example:
	//
	// a52f81be-496f-4e1c-a286-8852ab54****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The ID of the resource scenario. This parameter is returned only if the current template of the stack is generated from a resource scenario.
	//
	// example:
	//
	// ts-7f7a704cf71c49a6****
	TemplateScratchId *string `json:"TemplateScratchId,omitempty" xml:"TemplateScratchId,omitempty"`
	// The URL of the file that contains the template body. This parameter is returned only if the current template of the stack is from a URL. The URL can point to a template that is located on an HTTP or HTTPS web server or in an Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// oss://ros/template/demo
	TemplateURL *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	// The version of the template. This parameter is returned only if the current stack template is a custom template or shared template.
	//
	// If the template is a shared template, this parameter is returned only if VersionOption is set to AllVersions.
	//
	// Valid values: v1 to v100.
	//
	// example:
	//
	// v1
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
	// The timeout period for creating the stack. Unit: minutes.
	//
	// example:
	//
	// 10
	TimeoutInMinutes *int32 `json:"TimeoutInMinutes,omitempty" xml:"TimeoutInMinutes,omitempty"`
	// The time when the stack was updated. The time follows the ISO 8601 standard in the YYYY-MM-DDThh:mm:ss format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-09-17T08:21:40
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackResponseBody) GetCheckedStackResourceCount() *int32 {
	return s.CheckedStackResourceCount
}

func (s *GetStackResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStackResponseBody) GetDeletionProtection() *string {
	return s.DeletionProtection
}

func (s *GetStackResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetStackResponseBody) GetDisableRollback() *bool {
	return s.DisableRollback
}

func (s *GetStackResponseBody) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *GetStackResponseBody) GetInterface() *string {
	return s.Interface
}

func (s *GetStackResponseBody) GetLog() *GetStackResponseBodyLog {
	return s.Log
}

func (s *GetStackResponseBody) GetNotCheckedStackResourceCount() *int32 {
	return s.NotCheckedStackResourceCount
}

func (s *GetStackResponseBody) GetNotificationURLs() []*string {
	return s.NotificationURLs
}

func (s *GetStackResponseBody) GetOperationInfo() *GetStackResponseBodyOperationInfo {
	return s.OperationInfo
}

func (s *GetStackResponseBody) GetOrderIds() []*string {
	return s.OrderIds
}

func (s *GetStackResponseBody) GetOutputs() []map[string]interface{} {
	return s.Outputs
}

func (s *GetStackResponseBody) GetParameters() []*GetStackResponseBodyParameters {
	return s.Parameters
}

func (s *GetStackResponseBody) GetParentStackId() *string {
	return s.ParentStackId
}

func (s *GetStackResponseBody) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *GetStackResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetStackResponseBody) GetResourceProgress() *GetStackResponseBodyResourceProgress {
	return s.ResourceProgress
}

func (s *GetStackResponseBody) GetRollbackFailedRootReason() *string {
	return s.RollbackFailedRootReason
}

func (s *GetStackResponseBody) GetRootStackId() *string {
	return s.RootStackId
}

func (s *GetStackResponseBody) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *GetStackResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetStackResponseBody) GetStackDriftStatus() *string {
	return s.StackDriftStatus
}

func (s *GetStackResponseBody) GetStackId() *string {
	return s.StackId
}

func (s *GetStackResponseBody) GetStackName() *string {
	return s.StackName
}

func (s *GetStackResponseBody) GetStackType() *string {
	return s.StackType
}

func (s *GetStackResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetStackResponseBody) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetStackResponseBody) GetTags() []*GetStackResponseBodyTags {
	return s.Tags
}

func (s *GetStackResponseBody) GetTemplateDescription() *string {
	return s.TemplateDescription
}

func (s *GetStackResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetStackResponseBody) GetTemplateScratchId() *string {
	return s.TemplateScratchId
}

func (s *GetStackResponseBody) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *GetStackResponseBody) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetStackResponseBody) GetTimeoutInMinutes() *int32 {
	return s.TimeoutInMinutes
}

func (s *GetStackResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetStackResponseBody) SetCheckedStackResourceCount(v int32) *GetStackResponseBody {
	s.CheckedStackResourceCount = &v
	return s
}

func (s *GetStackResponseBody) SetCreateTime(v string) *GetStackResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetStackResponseBody) SetDeletionProtection(v string) *GetStackResponseBody {
	s.DeletionProtection = &v
	return s
}

func (s *GetStackResponseBody) SetDescription(v string) *GetStackResponseBody {
	s.Description = &v
	return s
}

func (s *GetStackResponseBody) SetDisableRollback(v bool) *GetStackResponseBody {
	s.DisableRollback = &v
	return s
}

func (s *GetStackResponseBody) SetDriftDetectionTime(v string) *GetStackResponseBody {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackResponseBody) SetInterface(v string) *GetStackResponseBody {
	s.Interface = &v
	return s
}

func (s *GetStackResponseBody) SetLog(v *GetStackResponseBodyLog) *GetStackResponseBody {
	s.Log = v
	return s
}

func (s *GetStackResponseBody) SetNotCheckedStackResourceCount(v int32) *GetStackResponseBody {
	s.NotCheckedStackResourceCount = &v
	return s
}

func (s *GetStackResponseBody) SetNotificationURLs(v []*string) *GetStackResponseBody {
	s.NotificationURLs = v
	return s
}

func (s *GetStackResponseBody) SetOperationInfo(v *GetStackResponseBodyOperationInfo) *GetStackResponseBody {
	s.OperationInfo = v
	return s
}

func (s *GetStackResponseBody) SetOrderIds(v []*string) *GetStackResponseBody {
	s.OrderIds = v
	return s
}

func (s *GetStackResponseBody) SetOutputs(v []map[string]interface{}) *GetStackResponseBody {
	s.Outputs = v
	return s
}

func (s *GetStackResponseBody) SetParameters(v []*GetStackResponseBodyParameters) *GetStackResponseBody {
	s.Parameters = v
	return s
}

func (s *GetStackResponseBody) SetParentStackId(v string) *GetStackResponseBody {
	s.ParentStackId = &v
	return s
}

func (s *GetStackResponseBody) SetRamRoleName(v string) *GetStackResponseBody {
	s.RamRoleName = &v
	return s
}

func (s *GetStackResponseBody) SetRegionId(v string) *GetStackResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetStackResponseBody) SetRequestId(v string) *GetStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackResponseBody) SetResourceGroupId(v string) *GetStackResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetStackResponseBody) SetResourceProgress(v *GetStackResponseBodyResourceProgress) *GetStackResponseBody {
	s.ResourceProgress = v
	return s
}

func (s *GetStackResponseBody) SetRollbackFailedRootReason(v string) *GetStackResponseBody {
	s.RollbackFailedRootReason = &v
	return s
}

func (s *GetStackResponseBody) SetRootStackId(v string) *GetStackResponseBody {
	s.RootStackId = &v
	return s
}

func (s *GetStackResponseBody) SetServiceManaged(v bool) *GetStackResponseBody {
	s.ServiceManaged = &v
	return s
}

func (s *GetStackResponseBody) SetServiceName(v string) *GetStackResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetStackResponseBody) SetStackDriftStatus(v string) *GetStackResponseBody {
	s.StackDriftStatus = &v
	return s
}

func (s *GetStackResponseBody) SetStackId(v string) *GetStackResponseBody {
	s.StackId = &v
	return s
}

func (s *GetStackResponseBody) SetStackName(v string) *GetStackResponseBody {
	s.StackName = &v
	return s
}

func (s *GetStackResponseBody) SetStackType(v string) *GetStackResponseBody {
	s.StackType = &v
	return s
}

func (s *GetStackResponseBody) SetStatus(v string) *GetStackResponseBody {
	s.Status = &v
	return s
}

func (s *GetStackResponseBody) SetStatusReason(v string) *GetStackResponseBody {
	s.StatusReason = &v
	return s
}

func (s *GetStackResponseBody) SetTags(v []*GetStackResponseBodyTags) *GetStackResponseBody {
	s.Tags = v
	return s
}

func (s *GetStackResponseBody) SetTemplateDescription(v string) *GetStackResponseBody {
	s.TemplateDescription = &v
	return s
}

func (s *GetStackResponseBody) SetTemplateId(v string) *GetStackResponseBody {
	s.TemplateId = &v
	return s
}

func (s *GetStackResponseBody) SetTemplateScratchId(v string) *GetStackResponseBody {
	s.TemplateScratchId = &v
	return s
}

func (s *GetStackResponseBody) SetTemplateURL(v string) *GetStackResponseBody {
	s.TemplateURL = &v
	return s
}

func (s *GetStackResponseBody) SetTemplateVersion(v string) *GetStackResponseBody {
	s.TemplateVersion = &v
	return s
}

func (s *GetStackResponseBody) SetTimeoutInMinutes(v int32) *GetStackResponseBody {
	s.TimeoutInMinutes = &v
	return s
}

func (s *GetStackResponseBody) SetUpdateTime(v string) *GetStackResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetStackResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStackResponseBodyLog struct {
	// The logs of resources in the stack. This parameter is returned if LogOption is set to Resource or All.
	//
	// >  The logs are returned only for resources of specific types, such as the `ALIYUN::ROS::ResourceCleaner` type.
	ResourceLogs []*GetStackResponseBodyLogResourceLogs `json:"ResourceLogs,omitempty" xml:"ResourceLogs,omitempty" type:"Repeated"`
	// The logs generated when the Terraform stack is run. This parameter is returned only for a Terraform stack. This parameter is returned if LogOption is left empty or set to Stack or All.
	//
	// >  This parameter is not returned for a running stack. The logs are generated from the most recent operation on the stack, such as the creation, resumed creation, update, or deletion operation.
	TerraformLogs []*GetStackResponseBodyLogTerraformLogs `json:"TerraformLogs,omitempty" xml:"TerraformLogs,omitempty" type:"Repeated"`
}

func (s GetStackResponseBodyLog) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyLog) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyLog) GetResourceLogs() []*GetStackResponseBodyLogResourceLogs {
	return s.ResourceLogs
}

func (s *GetStackResponseBodyLog) GetTerraformLogs() []*GetStackResponseBodyLogTerraformLogs {
	return s.TerraformLogs
}

func (s *GetStackResponseBodyLog) SetResourceLogs(v []*GetStackResponseBodyLogResourceLogs) *GetStackResponseBodyLog {
	s.ResourceLogs = v
	return s
}

func (s *GetStackResponseBodyLog) SetTerraformLogs(v []*GetStackResponseBodyLogTerraformLogs) *GetStackResponseBodyLog {
	s.TerraformLogs = v
	return s
}

func (s *GetStackResponseBodyLog) Validate() error {
	return dara.Validate(s)
}

type GetStackResponseBodyLogResourceLogs struct {
	// All the logs that are associated with the resources.
	Logs []*GetStackResponseBodyLogResourceLogsLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The name of the resource that is defined in the template.
	//
	// example:
	//
	// MyResourceCleaner
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s GetStackResponseBodyLogResourceLogs) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyLogResourceLogs) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyLogResourceLogs) GetLogs() []*GetStackResponseBodyLogResourceLogsLogs {
	return s.Logs
}

func (s *GetStackResponseBodyLogResourceLogs) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetStackResponseBodyLogResourceLogs) SetLogs(v []*GetStackResponseBodyLogResourceLogsLogs) *GetStackResponseBodyLogResourceLogs {
	s.Logs = v
	return s
}

func (s *GetStackResponseBodyLogResourceLogs) SetResourceName(v string) *GetStackResponseBodyLogResourceLogs {
	s.ResourceName = &v
	return s
}

func (s *GetStackResponseBodyLogResourceLogs) Validate() error {
	return dara.Validate(s)
}

type GetStackResponseBodyLogResourceLogsLogs struct {
	// The content of a resource log.
	//
	// example:
	//
	// []
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The keywords of a resource log.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
}

func (s GetStackResponseBodyLogResourceLogsLogs) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyLogResourceLogsLogs) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyLogResourceLogsLogs) GetContent() *string {
	return s.Content
}

func (s *GetStackResponseBodyLogResourceLogsLogs) GetKeys() []*string {
	return s.Keys
}

func (s *GetStackResponseBodyLogResourceLogsLogs) SetContent(v string) *GetStackResponseBodyLogResourceLogsLogs {
	s.Content = &v
	return s
}

func (s *GetStackResponseBodyLogResourceLogsLogs) SetKeys(v []*string) *GetStackResponseBodyLogResourceLogsLogs {
	s.Keys = v
	return s
}

func (s *GetStackResponseBodyLogResourceLogsLogs) Validate() error {
	return dara.Validate(s)
}

type GetStackResponseBodyLogTerraformLogs struct {
	// The name of the Terraform command that is run. Valid values:
	//
	// 	- apply
	//
	// 	- plan
	//
	// 	- destroy
	//
	// 	- version
	//
	// For more information about Terraform commands, see [Basic CLI Features](https://www.terraform.io/cli/commands).
	//
	// example:
	//
	// apply
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The content of the output stream that is returned after the command is run.
	//
	// example:
	//
	// Apply complete! Resources: 42 added, 0 changed, 0 destroyed.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The output stream. Valid values:
	//
	// 	- stdout: standard output stream
	//
	// 	- stderr: standard error stream
	//
	// example:
	//
	// stdout
	Stream *string `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s GetStackResponseBodyLogTerraformLogs) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyLogTerraformLogs) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyLogTerraformLogs) GetCommand() *string {
	return s.Command
}

func (s *GetStackResponseBodyLogTerraformLogs) GetContent() *string {
	return s.Content
}

func (s *GetStackResponseBodyLogTerraformLogs) GetStream() *string {
	return s.Stream
}

func (s *GetStackResponseBodyLogTerraformLogs) SetCommand(v string) *GetStackResponseBodyLogTerraformLogs {
	s.Command = &v
	return s
}

func (s *GetStackResponseBodyLogTerraformLogs) SetContent(v string) *GetStackResponseBodyLogTerraformLogs {
	s.Content = &v
	return s
}

func (s *GetStackResponseBodyLogTerraformLogs) SetStream(v string) *GetStackResponseBodyLogTerraformLogs {
	s.Stream = &v
	return s
}

func (s *GetStackResponseBodyLogTerraformLogs) Validate() error {
	return dara.Validate(s)
}

type GetStackResponseBodyOperationInfo struct {
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
	// The logical ID of the resource on which the operation error occurs.
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
	// The type of the resource on which the operation error occurs.
	//
	// example:
	//
	// ALIYUN::ECS::SecurityGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetStackResponseBodyOperationInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyOperationInfo) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyOperationInfo) GetAction() *string {
	return s.Action
}

func (s *GetStackResponseBodyOperationInfo) GetCode() *string {
	return s.Code
}

func (s *GetStackResponseBodyOperationInfo) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *GetStackResponseBodyOperationInfo) GetMessage() *string {
	return s.Message
}

func (s *GetStackResponseBodyOperationInfo) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackResponseBodyOperationInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetStackResponseBodyOperationInfo) SetAction(v string) *GetStackResponseBodyOperationInfo {
	s.Action = &v
	return s
}

func (s *GetStackResponseBodyOperationInfo) SetCode(v string) *GetStackResponseBodyOperationInfo {
	s.Code = &v
	return s
}

func (s *GetStackResponseBodyOperationInfo) SetLogicalResourceId(v string) *GetStackResponseBodyOperationInfo {
	s.LogicalResourceId = &v
	return s
}

func (s *GetStackResponseBodyOperationInfo) SetMessage(v string) *GetStackResponseBodyOperationInfo {
	s.Message = &v
	return s
}

func (s *GetStackResponseBodyOperationInfo) SetRequestId(v string) *GetStackResponseBodyOperationInfo {
	s.RequestId = &v
	return s
}

func (s *GetStackResponseBodyOperationInfo) SetResourceType(v string) *GetStackResponseBodyOperationInfo {
	s.ResourceType = &v
	return s
}

func (s *GetStackResponseBodyOperationInfo) Validate() error {
	return dara.Validate(s)
}

type GetStackResponseBodyParameters struct {
	// The parameter name.
	//
	// example:
	//
	// ALIYUN::Region
	ParameterKey *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// cn-hangzhou
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetStackResponseBodyParameters) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetStackResponseBodyParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetStackResponseBodyParameters) SetParameterKey(v string) *GetStackResponseBodyParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetStackResponseBodyParameters) SetParameterValue(v string) *GetStackResponseBodyParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetStackResponseBodyParameters) Validate() error {
	return dara.Validate(s)
}

type GetStackResponseBodyResourceProgress struct {
	// The number of resources that failed to be created.
	//
	// >  This parameter is returned only if `ShowResourceProgress` is set to `EnabledIfCreateStack`.
	//
	// example:
	//
	// 0
	FailedResourceCount *int32 `json:"FailedResourceCount,omitempty" xml:"FailedResourceCount,omitempty"`
	// The number of resources that are being created.
	//
	// >  This parameter is returned only if `ShowResourceProgress` is set to `EnabledIfCreateStack`.
	//
	// example:
	//
	// 1
	InProgressResourceCount *int32 `json:"InProgressResourceCount,omitempty" xml:"InProgressResourceCount,omitempty"`
	// The progress details of resources that are being created.
	//
	// >  This parameter is returned only if `ShowResourceProgress` is set to `EnabledIfCreateStack`.
	InProgressResourceDetails []*GetStackResponseBodyResourceProgressInProgressResourceDetails `json:"InProgressResourceDetails,omitempty" xml:"InProgressResourceDetails,omitempty" type:"Repeated"`
	// The number of resources to be created.
	//
	// >  This parameter is returned only if `ShowResourceProgress` is set to `EnabledIfCreateStack`.
	//
	// example:
	//
	// 0
	PendingResourceCount *int32 `json:"PendingResourceCount,omitempty" xml:"PendingResourceCount,omitempty"`
	// The creation or rollback progress of the stack, in percentage. Valid values: 0 to 100.
	//
	// The value progressively increases from 0 to 100 during a stack creation operation. If the stack is created, the value reaches 100. If the stack fails to be created, a rollback is started for the stack resources, and the value progressively increases from the percentage of the remaining progress (100 - Progress value generated when the stack fails to be created). The value increases to 100 when the stack resources are rolled back. This parameter indicates the creation progress during a stack creation operation and indicates the rollback progress during a stack rollback operation.
	//
	// >  This parameter is returned only if `ShowResourceProgress` is set to `PercentageOnly`.
	//
	// example:
	//
	// 100
	StackActionProgress *float32 `json:"StackActionProgress,omitempty" xml:"StackActionProgress,omitempty"`
	// The overall creation progress of the stack, in percentage. Valid values: 0 to 100.
	//
	// The value progressively increases from 0 to 100 during a stack creation operation. If the stack is created, the value reaches 100. If the stack fails to be created, a rollback is started for the stack resources, and the value progressively decreases. The value decreases to 0 when the stack resources are rolled back. This parameter indicates only the overall creation progress, regardless of whether during a stack creation or rollback operation.
	//
	// >  This parameter is returned only if `ShowResourceProgress` is set to `PercentageOnly`.
	//
	// example:
	//
	// 100
	StackOperationProgress *float32 `json:"StackOperationProgress,omitempty" xml:"StackOperationProgress,omitempty"`
	// The number of resources that are created.
	//
	// >  This parameter is returned only if `ShowResourceProgress` is set to `EnabledIfCreateStack`.
	//
	// example:
	//
	// 1
	SuccessResourceCount *int32 `json:"SuccessResourceCount,omitempty" xml:"SuccessResourceCount,omitempty"`
	// The total number of resources.
	//
	// >  This parameter is returned only if `ShowResourceProgress` is set to `EnabledIfCreateStack`.
	//
	// example:
	//
	// 2
	TotalResourceCount *int32 `json:"TotalResourceCount,omitempty" xml:"TotalResourceCount,omitempty"`
}

func (s GetStackResponseBodyResourceProgress) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyResourceProgress) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyResourceProgress) GetFailedResourceCount() *int32 {
	return s.FailedResourceCount
}

func (s *GetStackResponseBodyResourceProgress) GetInProgressResourceCount() *int32 {
	return s.InProgressResourceCount
}

func (s *GetStackResponseBodyResourceProgress) GetInProgressResourceDetails() []*GetStackResponseBodyResourceProgressInProgressResourceDetails {
	return s.InProgressResourceDetails
}

func (s *GetStackResponseBodyResourceProgress) GetPendingResourceCount() *int32 {
	return s.PendingResourceCount
}

func (s *GetStackResponseBodyResourceProgress) GetStackActionProgress() *float32 {
	return s.StackActionProgress
}

func (s *GetStackResponseBodyResourceProgress) GetStackOperationProgress() *float32 {
	return s.StackOperationProgress
}

func (s *GetStackResponseBodyResourceProgress) GetSuccessResourceCount() *int32 {
	return s.SuccessResourceCount
}

func (s *GetStackResponseBodyResourceProgress) GetTotalResourceCount() *int32 {
	return s.TotalResourceCount
}

func (s *GetStackResponseBodyResourceProgress) SetFailedResourceCount(v int32) *GetStackResponseBodyResourceProgress {
	s.FailedResourceCount = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetInProgressResourceCount(v int32) *GetStackResponseBodyResourceProgress {
	s.InProgressResourceCount = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetInProgressResourceDetails(v []*GetStackResponseBodyResourceProgressInProgressResourceDetails) *GetStackResponseBodyResourceProgress {
	s.InProgressResourceDetails = v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetPendingResourceCount(v int32) *GetStackResponseBodyResourceProgress {
	s.PendingResourceCount = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetStackActionProgress(v float32) *GetStackResponseBodyResourceProgress {
	s.StackActionProgress = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetStackOperationProgress(v float32) *GetStackResponseBodyResourceProgress {
	s.StackOperationProgress = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetSuccessResourceCount(v int32) *GetStackResponseBodyResourceProgress {
	s.SuccessResourceCount = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) SetTotalResourceCount(v int32) *GetStackResponseBodyResourceProgress {
	s.TotalResourceCount = &v
	return s
}

func (s *GetStackResponseBodyResourceProgress) Validate() error {
	return dara.Validate(s)
}

type GetStackResponseBodyResourceProgressInProgressResourceDetails struct {
	// The desired progress value of the resource.
	//
	// example:
	//
	// 10
	ProgressTargetValue *float32 `json:"ProgressTargetValue,omitempty" xml:"ProgressTargetValue,omitempty"`
	// The current progress value of the resource.
	//
	// example:
	//
	// 5
	ProgressValue *float32 `json:"ProgressValue,omitempty" xml:"ProgressValue,omitempty"`
	// The resource name.
	//
	// example:
	//
	// WaitCondition
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::ROS::WaitCondition
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetStackResponseBodyResourceProgressInProgressResourceDetails) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyResourceProgressInProgressResourceDetails) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) GetProgressTargetValue() *float32 {
	return s.ProgressTargetValue
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) GetProgressValue() *float32 {
	return s.ProgressValue
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) GetResourceName() *string {
	return s.ResourceName
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) SetProgressTargetValue(v float32) *GetStackResponseBodyResourceProgressInProgressResourceDetails {
	s.ProgressTargetValue = &v
	return s
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) SetProgressValue(v float32) *GetStackResponseBodyResourceProgressInProgressResourceDetails {
	s.ProgressValue = &v
	return s
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) SetResourceName(v string) *GetStackResponseBodyResourceProgressInProgressResourceDetails {
	s.ResourceName = &v
	return s
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) SetResourceType(v string) *GetStackResponseBodyResourceProgressInProgressResourceDetails {
	s.ResourceType = &v
	return s
}

func (s *GetStackResponseBodyResourceProgressInProgressResourceDetails) Validate() error {
	return dara.Validate(s)
}

type GetStackResponseBodyTags struct {
	// The tag key of the stack.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the stack.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetStackResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetStackResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetStackResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetStackResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetStackResponseBodyTags) SetKey(v string) *GetStackResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetStackResponseBodyTags) SetValue(v string) *GetStackResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetStackResponseBodyTags) Validate() error {
	return dara.Validate(s)
}

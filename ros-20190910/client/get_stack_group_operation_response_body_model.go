// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackGroupOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetStackGroupOperationResponseBody
	GetRequestId() *string
	SetStackGroupOperation(v *GetStackGroupOperationResponseBodyStackGroupOperation) *GetStackGroupOperationResponseBody
	GetStackGroupOperation() *GetStackGroupOperationResponseBodyStackGroupOperation
}

type GetStackGroupOperationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 14A07460-EBE7-47CA-9757-12CC4761D47A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the stack group operation.
	StackGroupOperation *GetStackGroupOperationResponseBodyStackGroupOperation `json:"StackGroupOperation,omitempty" xml:"StackGroupOperation,omitempty" type:"Struct"`
}

func (s GetStackGroupOperationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupOperationResponseBody) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStackGroupOperationResponseBody) GetStackGroupOperation() *GetStackGroupOperationResponseBodyStackGroupOperation {
	return s.StackGroupOperation
}

func (s *GetStackGroupOperationResponseBody) SetRequestId(v string) *GetStackGroupOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStackGroupOperationResponseBody) SetStackGroupOperation(v *GetStackGroupOperationResponseBodyStackGroupOperation) *GetStackGroupOperationResponseBody {
	s.StackGroupOperation = v
	return s
}

func (s *GetStackGroupOperationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStackGroupOperationResponseBodyStackGroupOperation struct {
	// The operation type.
	//
	// Valid values:
	//
	// 	- CREATE
	//
	// 	- UPDATE
	//
	// 	- DELETE
	//
	// 	- DETECT_DRIFT
	//
	// example:
	//
	// DELETE
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The name of the RAM role that you specify for the administrator account when you create the self-managed stack group. ROS assumes the administrator role to perform operations. If this parameter is not specified, the default value AliyunROSStackGroupAdministrationRole is returned.
	//
	// example:
	//
	// AliyunROSStackGroupAdministrationRole
	AdministrationRoleName *string `json:"AdministrationRoleName,omitempty" xml:"AdministrationRoleName,omitempty"`
	// The time when the operation was initiated.
	//
	// example:
	//
	// 2020-01-20T09:22:3
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The destinations to deploy stack instances when the stack is granted service-managed permissions.
	DeploymentTargets *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets `json:"DeploymentTargets,omitempty" xml:"DeploymentTargets,omitempty" type:"Struct"`
	// The time when the operation ended.
	//
	// example:
	//
	// 2020-01-20T09:22:4
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the RAM role that you specify for the execution account when you create the self-managed stack group. The administrator role AliyunROSStackGroupAdministrationRole assumes the execution role to perform operations. If this parameter is not specified, the default value AliyunROSStackGroupExecutionRole is returned.
	//
	// example:
	//
	// AliyunROSStackGroupExecutionRole
	ExecutionRoleName *string `json:"ExecutionRoleName,omitempty" xml:"ExecutionRoleName,omitempty"`
	// The description of the operation.
	//
	// > This parameter is returned only if OperationDescription is specified when the [CreateStackInstances](https://help.aliyun.com/document_detail/151338.html) operation is called to create stack instances.
	//
	// example:
	//
	// Create stack instance in hangzhou
	OperationDescription *string `json:"OperationDescription,omitempty" xml:"OperationDescription,omitempty"`
	// The operation ID.
	//
	// example:
	//
	// 6da106ca-1784-4a6f-a7e1-e723863d****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The operation settings.
	OperationPreferences *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences `json:"OperationPreferences,omitempty" xml:"OperationPreferences,omitempty" type:"Struct"`
	// Indicates whether stacks are retained when the associated stack instances are deleted. When you delete a stack instance, you can choose to delete or retain the stack with which the stack instance is associated.
	//
	// Valid values:
	//
	// 	- true: Stacks are retained when the associated stack instances are deleted.
	//
	// 	- false: Stacks are deleted when the associated stack instances are deleted. Proceed with caution.
	//
	// > This parameter is returned only if you delete stack instances.
	//
	// example:
	//
	// true
	RetainStacks *bool `json:"RetainStacks,omitempty" xml:"RetainStacks,omitempty"`
	// The information about drift detection.
	//
	// > This parameter is returned only if drift detection is performed.
	StackGroupDriftDetectionDetail *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail `json:"StackGroupDriftDetectionDetail,omitempty" xml:"StackGroupDriftDetectionDetail,omitempty" type:"Struct"`
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
	// The state of the operation.
	//
	// Valid values:
	//
	// 	- RUNNING
	//
	// 	- SUCCEEDED
	//
	// 	- FAILED
	//
	// 	- STOPPING
	//
	// 	- STOPPED
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetStackGroupOperationResponseBodyStackGroupOperation) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupOperationResponseBodyStackGroupOperation) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetAction() *string {
	return s.Action
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetAdministrationRoleName() *string {
	return s.AdministrationRoleName
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetDeploymentTargets() *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets {
	return s.DeploymentTargets
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetEndTime() *string {
	return s.EndTime
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetExecutionRoleName() *string {
	return s.ExecutionRoleName
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetOperationDescription() *string {
	return s.OperationDescription
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetOperationId() *string {
	return s.OperationId
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetOperationPreferences() *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	return s.OperationPreferences
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetRetainStacks() *bool {
	return s.RetainStacks
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetStackGroupDriftDetectionDetail() *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	return s.StackGroupDriftDetectionDetail
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetStackGroupId() *string {
	return s.StackGroupId
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetStackGroupName() *string {
	return s.StackGroupName
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) GetStatus() *string {
	return s.Status
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetAction(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.Action = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetAdministrationRoleName(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.AdministrationRoleName = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetCreateTime(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.CreateTime = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetDeploymentTargets(v *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.DeploymentTargets = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetEndTime(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.EndTime = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetExecutionRoleName(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.ExecutionRoleName = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetOperationDescription(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.OperationDescription = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetOperationId(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.OperationId = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetOperationPreferences(v *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.OperationPreferences = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetRetainStacks(v bool) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.RetainStacks = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStackGroupDriftDetectionDetail(v *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.StackGroupDriftDetectionDetail = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStackGroupId(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.StackGroupId = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStackGroupName(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.StackGroupName = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) SetStatus(v string) *GetStackGroupOperationResponseBodyStackGroupOperation {
	s.Status = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperation) Validate() error {
	return dara.Validate(s)
}

type GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets struct {
	// The IDs of the members in the resource directory.
	//
	// > This parameter is returned only if AccountIds is specified when the [UpdateStackInstances](https://help.aliyun.com/document_detail/151716.html) operation is called to update stack instances.
	AccountIds []*string `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The IDs of the folders in the resource directory.
	RdFolderIds []*string `json:"RdFolderIds,omitempty" xml:"RdFolderIds,omitempty" type:"Repeated"`
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) GetAccountIds() []*string {
	return s.AccountIds
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) GetRdFolderIds() []*string {
	return s.RdFolderIds
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) SetAccountIds(v []*string) *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets {
	s.AccountIds = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) SetRdFolderIds(v []*string) *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets {
	s.RdFolderIds = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationDeploymentTargets) Validate() error {
	return dara.Validate(s)
}

type GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences struct {
	// The number of accounts within which stack operation failures are allowed to occur in each region. If the value of this parameter is exceeded in a region, Resource Orchestration Service (ROS) stops the operation in the region. If the operation is stopped in one region, the operation is no longer performed in other regions.
	//
	// Valid values: 0 to 20.
	//
	// > Only one of FailureToleranceCount and FailureTolerancePercentage can be returned.
	//
	// example:
	//
	// 1
	FailureToleranceCount *int32 `json:"FailureToleranceCount,omitempty" xml:"FailureToleranceCount,omitempty"`
	// The percentage of the number of accounts within which stack operation failures are allowed to occur to the total number of accounts in each region. If the value of this parameter is exceeded in a region, ROS stops the operation in the region.
	//
	// Valid values: 0 to 100.
	//
	// > Only one of FailureToleranceCount and FailureTolerancePercentage can be returned.
	//
	// example:
	//
	// 10
	FailureTolerancePercentage *int32 `json:"FailureTolerancePercentage,omitempty" xml:"FailureTolerancePercentage,omitempty"`
	// The maximum number of accounts within which stacks are deployed at the same time in each region.
	//
	// Valid values: 1 to 20.
	//
	// > Only one of MaxConcurrentCount and MaxConcurrentPercentage can be returned.
	//
	// example:
	//
	// 1
	MaxConcurrentCount *int32 `json:"MaxConcurrentCount,omitempty" xml:"MaxConcurrentCount,omitempty"`
	// The percentage of the maximum number of accounts within which stacks are deployed at the same time to the total number of accounts in each region.
	//
	// Valid values: 1 to 100.
	//
	// > Only one of MaxConcurrentCount and MaxConcurrentPercentage can be returned.
	//
	// example:
	//
	// 10
	MaxConcurrentPercentage *int32 `json:"MaxConcurrentPercentage,omitempty" xml:"MaxConcurrentPercentage,omitempty"`
	// The regions in the order of operation execution.
	RegionIdsOrder []*string `json:"RegionIdsOrder,omitempty" xml:"RegionIdsOrder,omitempty" type:"Repeated"`
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) GetFailureToleranceCount() *int32 {
	return s.FailureToleranceCount
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) GetFailureTolerancePercentage() *int32 {
	return s.FailureTolerancePercentage
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) GetMaxConcurrentCount() *int32 {
	return s.MaxConcurrentCount
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) GetMaxConcurrentPercentage() *int32 {
	return s.MaxConcurrentPercentage
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) GetRegionIdsOrder() []*string {
	return s.RegionIdsOrder
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetFailureToleranceCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.FailureToleranceCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetFailureTolerancePercentage(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.FailureTolerancePercentage = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetMaxConcurrentCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.MaxConcurrentCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetMaxConcurrentPercentage(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.MaxConcurrentPercentage = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) SetRegionIdsOrder(v []*string) *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences {
	s.RegionIdsOrder = v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationOperationPreferences) Validate() error {
	return dara.Validate(s)
}

type GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail struct {
	// The number of stack instances for which drift detection was canceled.
	//
	// example:
	//
	// 0
	CancelledStackInstancesCount *int32 `json:"CancelledStackInstancesCount,omitempty" xml:"CancelledStackInstancesCount,omitempty"`
	// The drift detection state.
	//
	// Valid values:
	//
	// 	- COMPLETED: Drift detection is performed on the stack group and all stack instances passed the drift detection.
	//
	// 	- FAILED: Drift detection is performed on the stack group. The number of stack instances that failed the drift detection exceeds the specified threshold.
	//
	// 	- PARTIAL_SUCCESS: Drift detection is performed on the stack group. The number of stack instances that failed the drift detection does not exceed the specified threshold.
	//
	// 	- IN_PROGRESS: Drift detection is being performed on the stack group.
	//
	// 	- STOPPED: Drift detection is canceled for the stack group.
	//
	// example:
	//
	// COMPLETED
	DriftDetectionStatus *string `json:"DriftDetectionStatus,omitempty" xml:"DriftDetectionStatus,omitempty"`
	// The time when drift detection was performed.
	//
	// example:
	//
	// 2020-02-27T07:47:47
	DriftDetectionTime *string `json:"DriftDetectionTime,omitempty" xml:"DriftDetectionTime,omitempty"`
	// The number of stack instances that have drifted.
	//
	// example:
	//
	// 1
	DriftedStackInstancesCount *int32 `json:"DriftedStackInstancesCount,omitempty" xml:"DriftedStackInstancesCount,omitempty"`
	// The number of stack instances that failed drift detection.
	//
	// example:
	//
	// 0
	FailedStackInstancesCount *int32 `json:"FailedStackInstancesCount,omitempty" xml:"FailedStackInstancesCount,omitempty"`
	// The number of stack instances on which drift detection was being performed.
	//
	// example:
	//
	// 0
	InProgressStackInstancesCount *int32 `json:"InProgressStackInstancesCount,omitempty" xml:"InProgressStackInstancesCount,omitempty"`
	// The number of stack instances that were being synchronized.
	//
	// example:
	//
	// 1
	InSyncStackInstancesCount *int32 `json:"InSyncStackInstancesCount,omitempty" xml:"InSyncStackInstancesCount,omitempty"`
	// The drift state of the stack group.
	//
	// Valid values:
	//
	// 	- DRIFTED: At least one stack instance in the stack group has drifted.
	//
	// 	- NOT_CHECKED: No successful drift detection is performed in the stack group.
	//
	// 	- IN_SYNC: All the stack instances in the stack group are being synchronized.
	//
	// example:
	//
	// DRIFTED
	StackGroupDriftStatus *string `json:"StackGroupDriftStatus,omitempty" xml:"StackGroupDriftStatus,omitempty"`
	// The number of stack instances.
	//
	// example:
	//
	// 2
	TotalStackInstancesCount *int32 `json:"TotalStackInstancesCount,omitempty" xml:"TotalStackInstancesCount,omitempty"`
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GetCancelledStackInstancesCount() *int32 {
	return s.CancelledStackInstancesCount
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GetDriftDetectionStatus() *string {
	return s.DriftDetectionStatus
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GetDriftDetectionTime() *string {
	return s.DriftDetectionTime
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GetDriftedStackInstancesCount() *int32 {
	return s.DriftedStackInstancesCount
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GetFailedStackInstancesCount() *int32 {
	return s.FailedStackInstancesCount
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GetInProgressStackInstancesCount() *int32 {
	return s.InProgressStackInstancesCount
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GetInSyncStackInstancesCount() *int32 {
	return s.InSyncStackInstancesCount
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GetStackGroupDriftStatus() *string {
	return s.StackGroupDriftStatus
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) GetTotalStackInstancesCount() *int32 {
	return s.TotalStackInstancesCount
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetCancelledStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.CancelledStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetDriftDetectionStatus(v string) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.DriftDetectionStatus = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetDriftDetectionTime(v string) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.DriftDetectionTime = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetDriftedStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.DriftedStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetFailedStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.FailedStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetInProgressStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.InProgressStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetInSyncStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.InSyncStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetStackGroupDriftStatus(v string) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.StackGroupDriftStatus = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) SetTotalStackInstancesCount(v int32) *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail {
	s.TotalStackInstancesCount = &v
	return s
}

func (s *GetStackGroupOperationResponseBodyStackGroupOperationStackGroupDriftDetectionDetail) Validate() error {
	return dara.Validate(s)
}

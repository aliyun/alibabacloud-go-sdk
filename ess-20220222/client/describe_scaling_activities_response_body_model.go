// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingActivitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeScalingActivitiesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScalingActivitiesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeScalingActivitiesResponseBody
	GetRequestId() *string
	SetScalingActivities(v []*DescribeScalingActivitiesResponseBodyScalingActivities) *DescribeScalingActivitiesResponseBody
	GetScalingActivities() []*DescribeScalingActivitiesResponseBodyScalingActivities
	SetTotalCount(v int32) *DescribeScalingActivitiesResponseBody
	GetTotalCount() *int32
}

type DescribeScalingActivitiesResponseBody struct {
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
	// The ID of the request.
	//
	// example:
	//
	// CC107349-57B7-4405-B1BF-9BF5AF7F2A46
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scaling activities.
	ScalingActivities []*DescribeScalingActivitiesResponseBodyScalingActivities `json:"ScalingActivities,omitempty" xml:"ScalingActivities,omitempty" type:"Repeated"`
	// The total number of scaling activities.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScalingActivitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScalingActivitiesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScalingActivitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScalingActivitiesResponseBody) GetScalingActivities() []*DescribeScalingActivitiesResponseBodyScalingActivities {
	return s.ScalingActivities
}

func (s *DescribeScalingActivitiesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeScalingActivitiesResponseBody) SetPageNumber(v int32) *DescribeScalingActivitiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetPageSize(v int32) *DescribeScalingActivitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetRequestId(v string) *DescribeScalingActivitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetScalingActivities(v []*DescribeScalingActivitiesResponseBodyScalingActivities) *DescribeScalingActivitiesResponseBody {
	s.ScalingActivities = v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) SetTotalCount(v int32) *DescribeScalingActivitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingActivitiesResponseBodyScalingActivities struct {
	// The metadata of the scaling activity.
	//
	// example:
	//
	// {\\"goatscaler.io/managed\\":\\"true\\"}
	ActivityMetadata *string `json:"ActivityMetadata,omitempty" xml:"ActivityMetadata,omitempty"`
	// The total number of instances that are manually added to the scaling group after the scaling activity is complete.
	//
	// example:
	//
	// 0
	AttachedCapacity *string `json:"AttachedCapacity,omitempty" xml:"AttachedCapacity,omitempty"`
	// The total number of instances that are created by Auto Scaling after the scaling activity was complete.
	//
	// example:
	//
	// 2
	AutoCreatedCapacity *string `json:"AutoCreatedCapacity,omitempty" xml:"AutoCreatedCapacity,omitempty"`
	// The reason why the scaling activity was triggered.
	//
	// example:
	//
	// A user requests to execute scaling rule \\"asr-bp12tcnol686y1ik****\\", changing the Total Capacity from \\"1\\" to \\"2\\".
	Cause *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// The number of instances that are created during the scale-out event.
	//
	// example:
	//
	// 1
	CreatedCapacity *int32 `json:"CreatedCapacity,omitempty" xml:"CreatedCapacity,omitempty"`
	// The instances that are created during the scale-out event.
	CreatedInstances []*string `json:"CreatedInstances,omitempty" xml:"CreatedInstances,omitempty" type:"Repeated"`
	// The description of the scaling activity.
	//
	// example:
	//
	// Add \\"1\\" ECS instance
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The number of instances that are released during the scale-in event.
	//
	// example:
	//
	// 1
	DestroyedCapacity *int32 `json:"DestroyedCapacity,omitempty" xml:"DestroyedCapacity,omitempty"`
	// The instances that are released during the scale-in event.
	DestroyedInstances []*string `json:"DestroyedInstances,omitempty" xml:"DestroyedInstances,omitempty" type:"Repeated"`
	// Details of the scaling activity.
	//
	// example:
	//
	// "new ECS instances "i-j6c8ilerw, i-j6c8iler4mx" are created."
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The time when the scaling activity was complete.
	//
	// example:
	//
	// 2020-09-10T09:54Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The error code that is returned when the scaling activity failed.
	//
	// example:
	//
	// OperationDenied.NoStock
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned when the scaling activity failed.
	//
	// example:
	//
	// The specified ECS resource is out of stock in this region. Please try again later.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The error messages that are returned when the scaling activities failed or are partially successful.
	ErrorMessages []*DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages `json:"ErrorMessages,omitempty" xml:"ErrorMessages,omitempty" type:"Repeated"`
	// The ID of the instance refresh task.
	//
	// example:
	//
	// ir-asdf12adsxg*****
	InstanceRefreshTaskId *string `json:"InstanceRefreshTaskId,omitempty" xml:"InstanceRefreshTaskId,omitempty"`
	// The context of the lifecycle hook.
	LifecycleHookContext *DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext `json:"LifecycleHookContext,omitempty" xml:"LifecycleHookContext,omitempty" type:"Struct"`
	// The execution progress of the scaling activity.
	//
	// example:
	//
	// 100
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the scaling activity.
	//
	// example:
	//
	// asa-bp161xudmuxdzofe****
	ScalingActivityId *string `json:"ScalingActivityId,omitempty" xml:"ScalingActivityId,omitempty"`
	// The ID of the scaling group.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	// 	- If you query a scale-out activity, the value of this parameter indicates the number of instances that are created or the number of instances that are started from Economical Mode.
	//
	// 	- If you query a scale-in activity, the value of this parameter indicates the number of instances that are deleted or the number of instances that are stopped in Economical Mode.
	//
	// example:
	//
	// 1
	ScalingInstanceNumber *int32 `json:"ScalingInstanceNumber,omitempty" xml:"ScalingInstanceNumber,omitempty"`
	// The time when the scaling activity was started.
	//
	// example:
	//
	// 2020-09-10T09:54Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of instances that are started from the Economical Mode during the scale-out event.
	//
	// example:
	//
	// 1
	StartedCapacity *int32 `json:"StartedCapacity,omitempty" xml:"StartedCapacity,omitempty"`
	// The instances that are started from the Economical Mode during the scale-out event.
	StartedInstances []*string `json:"StartedInstances,omitempty" xml:"StartedInstances,omitempty" type:"Repeated"`
	// The status of the scaling activity. Valid values:
	//
	// 	- Successful: The scaling activity is successful.
	//
	// 	- Warning: The scaling activity is partially successful.
	//
	// 	- Failed: The scaling activity failed.
	//
	// 	- InProgress: The scaling activity is in progress.
	//
	// 	- Rejected: The request to trigger the scaling activity is rejected.
	//
	// example:
	//
	// Successful
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The status message of the scaling activity.
	//
	// example:
	//
	// \\"1\\" ECS instances are added
	StatusMessage *string `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	// The number of instances that are stopped in the Economical Mode during the scale-in event.
	//
	// example:
	//
	// 1
	StoppedCapacity *int32 `json:"StoppedCapacity,omitempty" xml:"StoppedCapacity,omitempty"`
	// The instances that are stopped in the Economical Mode during the scale-in event.
	StoppedInstances []*string `json:"StoppedInstances,omitempty" xml:"StoppedInstances,omitempty" type:"Repeated"`
	// The total number of instances in the scaling group after the scaling activity was complete.
	//
	// example:
	//
	// 2
	TotalCapacity *string `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	// The ID of the trigger source of the scaling activity. Valid values:
	//
	// 	- If the scaling activity is triggered by an event-triggered task, the ID of the trigger source is the ID of the event-triggered task.
	//
	// 	- If the scaling activity is triggered by calling an API operation, the ID of the trigger source is the ID of the Alibaba Cloud account or Resource Access Management (RAM) user that you use to call the API operation.
	//
	// 	- If the scaling activity is triggered by Auto Scaling, the ID of the trigger source is null.
	//
	// example:
	//
	// 2346366580*****
	TriggerSourceId *string `json:"TriggerSourceId,omitempty" xml:"TriggerSourceId,omitempty"`
	// The type of the trigger source of the scaling activity. Valid values:
	//
	// 	- Cms: The scaling activity is triggered by an event-triggered task.
	//
	// 	- APIs: The scaling activity is triggered by calling an API operation.
	//
	// 	- Ess: The scaling activity is triggered by Auto Scaling.
	//
	// example:
	//
	// Api
	TriggerSourceType *string `json:"TriggerSourceType,omitempty" xml:"TriggerSourceType,omitempty"`
}

func (s DescribeScalingActivitiesResponseBodyScalingActivities) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBodyScalingActivities) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetActivityMetadata() *string {
	return s.ActivityMetadata
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetAttachedCapacity() *string {
	return s.AttachedCapacity
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetAutoCreatedCapacity() *string {
	return s.AutoCreatedCapacity
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetCause() *string {
	return s.Cause
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetCreatedCapacity() *int32 {
	return s.CreatedCapacity
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetCreatedInstances() []*string {
	return s.CreatedInstances
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetDescription() *string {
	return s.Description
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetDestroyedCapacity() *int32 {
	return s.DestroyedCapacity
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetDestroyedInstances() []*string {
	return s.DestroyedInstances
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetDetail() *string {
	return s.Detail
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetErrorMessages() []*DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages {
	return s.ErrorMessages
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetInstanceRefreshTaskId() *string {
	return s.InstanceRefreshTaskId
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetLifecycleHookContext() *DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext {
	return s.LifecycleHookContext
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetProgress() *int32 {
	return s.Progress
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetScalingInstanceNumber() *int32 {
	return s.ScalingInstanceNumber
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetStartedCapacity() *int32 {
	return s.StartedCapacity
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetStartedInstances() []*string {
	return s.StartedInstances
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetStatusCode() *string {
	return s.StatusCode
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetStoppedCapacity() *int32 {
	return s.StoppedCapacity
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetStoppedInstances() []*string {
	return s.StoppedInstances
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetTotalCapacity() *string {
	return s.TotalCapacity
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetTriggerSourceId() *string {
	return s.TriggerSourceId
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) GetTriggerSourceType() *string {
	return s.TriggerSourceType
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetActivityMetadata(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ActivityMetadata = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetAttachedCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.AttachedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetAutoCreatedCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.AutoCreatedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetCause(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.Cause = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetCreatedCapacity(v int32) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.CreatedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetCreatedInstances(v []*string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.CreatedInstances = v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetDescription(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.Description = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetDestroyedCapacity(v int32) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.DestroyedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetDestroyedInstances(v []*string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.DestroyedInstances = v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetDetail(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.Detail = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetEndTime(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.EndTime = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetErrorCode(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ErrorCode = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetErrorMessage(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetErrorMessages(v []*DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ErrorMessages = v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetInstanceRefreshTaskId(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.InstanceRefreshTaskId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetLifecycleHookContext(v *DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.LifecycleHookContext = v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetProgress(v int32) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.Progress = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetScalingActivityId(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ScalingActivityId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetScalingGroupId(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetScalingInstanceNumber(v int32) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.ScalingInstanceNumber = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetStartTime(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.StartTime = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetStartedCapacity(v int32) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.StartedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetStartedInstances(v []*string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.StartedInstances = v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetStatusCode(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.StatusCode = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetStatusMessage(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.StatusMessage = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetStoppedCapacity(v int32) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.StoppedCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetStoppedInstances(v []*string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.StoppedInstances = v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetTotalCapacity(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetTriggerSourceId(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.TriggerSourceId = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) SetTriggerSourceType(v string) *DescribeScalingActivitiesResponseBodyScalingActivities {
	s.TriggerSourceType = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivities) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages struct {
	// The error code that is returned when the scaling activity failed.
	//
	// example:
	//
	// OperationDenied.NoStock
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the scaling activity exception.
	//
	// example:
	//
	// Fail to create instances into scaling group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of the instances included in the failed scaling activities.
	FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" xml:"FailedInstanceIds,omitempty" type:"Repeated"`
	// The error message that is returned when the scaling activity failed or is partially successful.
	//
	// example:
	//
	// The resource is out of stock in the specified zone. Please try other types, or choose other regions and zones.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) GetCode() *string {
	return s.Code
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) GetDescription() *string {
	return s.Description
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) GetFailedInstanceIds() []*string {
	return s.FailedInstanceIds
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) GetMessage() *string {
	return s.Message
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) SetCode(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages {
	s.Code = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) SetDescription(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages {
	s.Description = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) SetFailedInstanceIds(v []*string) *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages {
	s.FailedInstanceIds = v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) SetMessage(v string) *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages {
	s.Message = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesErrorMessages) Validate() error {
	return dara.Validate(s)
}

type DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext struct {
	// Indicates whether all lifecycle hooks are disabled when the scaling activity is triggered. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DisableLifecycleHook *bool `json:"DisableLifecycleHook,omitempty" xml:"DisableLifecycleHook,omitempty"`
	// The IDs of the lifecycle hooks that are disabled.
	IgnoredLifecycleHookIds []*string `json:"IgnoredLifecycleHookIds,omitempty" xml:"IgnoredLifecycleHookIds,omitempty" type:"Repeated"`
}

func (s DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext) GoString() string {
	return s.String()
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext) GetDisableLifecycleHook() *bool {
	return s.DisableLifecycleHook
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext) GetIgnoredLifecycleHookIds() []*string {
	return s.IgnoredLifecycleHookIds
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext) SetDisableLifecycleHook(v bool) *DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext {
	s.DisableLifecycleHook = &v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext) SetIgnoredLifecycleHookIds(v []*string) *DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext {
	s.IgnoredLifecycleHookIds = v
	return s
}

func (s *DescribeScalingActivitiesResponseBodyScalingActivitiesLifecycleHookContext) Validate() error {
	return dara.Validate(s)
}

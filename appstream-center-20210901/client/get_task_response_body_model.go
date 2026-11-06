// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTaskResponseBody
	GetCode() *string
	SetData(v *GetTaskResponseBodyData) *GetTaskResponseBody
	GetData() *GetTaskResponseBodyData
	SetMessage(v string) *GetTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskResponseBody
	GetSuccess() *bool
}

type GetTaskResponseBody struct {
	// The error code. This parameter is returned when the request fails. A successful response may not contain this parameter.
	//
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The task details.
	Data *GetTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error or result message. This parameter is used to describe the cause when the request fails. A successful response may not contain this parameter.
	//
	// example:
	//
	// Internal error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can provide this value when you troubleshoot issues.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was processed successfully. A value of true indicates that the request was processed successfully. To determine whether an asynchronous task is completed, you must query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTaskResponseBody) GetData() *GetTaskResponseBodyData {
	return s.Data
}

func (s *GetTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskResponseBody) SetCode(v string) *GetTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskResponseBody) SetData(v *GetTaskResponseBodyData) *GetTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskResponseBody) SetMessage(v string) *GetTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetSuccess(v bool) *GetTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskResponseBodyData struct {
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 100931896542****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The prompt or execution result description of the current step of the task.
	//
	// example:
	//
	// [code=Builder.Failed,httpStatusCode=<null>,message=Builder environment deploy failed.]
	CurrentTaskMessage *string `json:"CurrentTaskMessage,omitempty" xml:"CurrentTaskMessage,omitempty"`
	// The identifier of the step that the task is currently executing. You can call the ListTaskStep operation to query the image creation and application adaptation steps.
	//
	// example:
	//
	// CREATE_BUILDER_ENV
	CurrentTaskStep *string `json:"CurrentTaskStep,omitempty" xml:"CurrentTaskStep,omitempty"`
	// The ID of the region where the task is actually executed.
	//
	// example:
	//
	// cn-hangzhou
	ExecuteRegion *string `json:"ExecuteRegion,omitempty" xml:"ExecuteRegion,omitempty"`
	// The time when the task was created. The time is in ISO 8601 format, accurate to milliseconds, and the time zone offset is +00:00.
	//
	// example:
	//
	// 2026-09-23T02:28:10.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the task was last modified. The time is in ISO 8601 format, accurate to milliseconds, and the time zone offset is +00:00.
	//
	// example:
	//
	// 2026-09-23T02:32:43.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The compute type of the task creation resource.
	//
	// Valid values:
	//
	// - **CPU**: CPU compute resource
	//
	// - **GPU**: GPU compute resource
	//
	// example:
	//
	// CPU
	PerformanceType *string `json:"PerformanceType,omitempty" xml:"PerformanceType,omitempty"`
	// **Not enabled. The system does not return valid task progress.*	- Determine the task execution status based on the Status, CurrentTaskStep, and TaskStepList parameters.
	//
	// The example "Not applicable (not enabled)" is only a description of the field status, not an actual API response example.
	//
	// example:
	//
	// Not applicable (not enabled)
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The list of resources used for task creation, which corresponds to the creation resources submitted when the task is created.
	ResourceList []*GetTaskResponseBodyDataResourceList `json:"ResourceList,omitempty" xml:"ResourceList,omitempty" type:"Repeated"`
	// The running mode of the task creation environment.
	//
	// Valid values:
	//
	// - **NORMAL**: normal running mode
	//
	// - **COMPATIBLE**: compatible running mode
	//
	// example:
	//
	// NORMAL
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// The overall status of the task.
	//
	// Valid values:
	//
	// - **INIT**: The task is initialized.
	//
	// - **START**: The task is being created or executed.
	//
	// - **SUCCESS**: The task is executed successfully.
	//
	// - **FAILED**: The task fails to be executed.
	//
	// - **DELETED**: The task is deleted.
	//
	// Deleted tasks are not included in the regular valid task query scope.
	//
	// example:
	//
	// START
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The system disk capacity of the task creation environment. Unit: GiB.
	//
	// example:
	//
	// 80
	SystemDiskSize *int32 `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	// The ID of the base system image used in the task creation environment.
	//
	// example:
	//
	// imgc-bp1parent****
	SystemImageId *string `json:"SystemImageId,omitempty" xml:"SystemImageId,omitempty"`
	// The task ID. You can use this ID to call the GetTask operation to query the task status.
	//
	// example:
	//
	// tid-bp1example****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The list of recorded task execution steps and results.
	TaskStepList []*GetTaskResponseBodyDataTaskStepList `json:"TaskStepList,omitempty" xml:"TaskStepList,omitempty" type:"Repeated"`
	// The actual type of the task.
	//
	// Valid values:
	//
	// - **AUTO_ADAPT**: automatic adaptation
	//
	// - **AUTO_CREATE_IMAGE**: automatic image creation
	//
	// - **RETRY_ADAPT**: retry of automatic adaptation
	//
	// - **RETRY_CREATE_IMAGE**: retry of automatic image creation
	//
	// - **AUTO_OTA_IMAGE**: automatic image OTA
	//
	// - **AUTO_CREATE_VHD**: automatic VHD creation
	//
	// - **RETRY_VHD**: retry of VHD creation
	//
	// - **AUTO_CREATE_IMAGE_VHD**: automatic VHD image creation
	//
	// - **RETRY_CREATE_IMAGE_VHD**: retry of VHD image creation
	//
	// - **AUTO_VALIDATE_VHD**: automatic VHD validation
	//
	// - **RETRY_VALIDATE_VHD**: retry of VHD validation
	//
	// - **SYNCHRONIZE_APP_IMAGE**: application image synchronization
	//
	// - **APPLY_VHD_BUILDER**: application for VHD application recording environment
	//
	// - **RETRY_VHD_BUILDER**: retry of application for VHD recording environment
	//
	// - **ADAPT_AUTO**: automatic application adaptation creation
	//
	// - **ADAPT_AUTO_RETRY**: retry of automatic application adaptation creation
	//
	// - **ADAPT_MANUAL**: manual application adaptation creation
	//
	// - **ADAPT_MANUAL_RETRY**: retry of manual application adaptation creation
	//
	// - **IMAGE_AUTO**: automatic image creation
	//
	// - **IMAGE_AUTO_RETRY**: retry of automatic image creation
	//
	// - **IMAGE_AUTO_VALIDATE**: automatic image validation
	//
	// - **IMAGE_AUTO_VALIDATE_RETRY**: retry of automatic image validation
	//
	// - **IMAGE_MANUAL**: manual image creation
	//
	// - **IMAGE_MANUAL_RETRY**: retry of manual image creation
	//
	// - **DESKTOP_CUSTOM_IMAGE**: cloud desktop custom image creation
	//
	// - **DESKTOP_CUSTOM_IMAGE_RETRY**: retry of cloud desktop custom image creation
	//
	// - **DESKTOP_SYNC_IMAGE**: cloud desktop image synchronization across regions
	//
	// - **DESKTOP_SYNC_IMAGE_RETRY**: retry of cloud desktop image synchronization across regions
	//
	// - **IMAGE_DISTRIBUTE**: image distribution to multiple regions
	//
	// - **IMAGE_DISTRIBUTE_RETRY**: retry of image distribution
	//
	// - **IMAGE_FROM_INSTANCE**: image creation based on an instance
	//
	// - **IMAGE_FROM_INSTANCE_RETRY**: retry of image creation based on an instance
	//
	// - **SNAPSHOT_DISTRIBUTE**: snapshot distribution to multiple regions
	//
	// - **SNAPSHOT_DISTRIBUTE_RETRY**: retry of snapshot distribution
	//
	// - **SNAPSHOT_FROM_DISK**: snapshot creation based on a disk
	//
	// - **SNAPSHOT_FROM_DISK_RETRY**: retry of snapshot creation based on a disk
	//
	// - **IMAGE_CLEAN_AND_CREATE**: temporary instance creation based on an image, cleanup, and final image creation
	//
	// - **IMAGE_CLEAN_AND_CREATE_RETRY**: retry of final image creation after cleanup
	//
	// example:
	//
	// IMAGE_MANUAL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyData) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetTaskResponseBodyData) GetCurrentTaskMessage() *string {
	return s.CurrentTaskMessage
}

func (s *GetTaskResponseBodyData) GetCurrentTaskStep() *string {
	return s.CurrentTaskStep
}

func (s *GetTaskResponseBodyData) GetExecuteRegion() *string {
	return s.ExecuteRegion
}

func (s *GetTaskResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTaskResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetTaskResponseBodyData) GetPerformanceType() *string {
	return s.PerformanceType
}

func (s *GetTaskResponseBodyData) GetProgress() *string {
	return s.Progress
}

func (s *GetTaskResponseBodyData) GetResourceList() []*GetTaskResponseBodyDataResourceList {
	return s.ResourceList
}

func (s *GetTaskResponseBodyData) GetRunMode() *string {
	return s.RunMode
}

func (s *GetTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTaskResponseBodyData) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *GetTaskResponseBodyData) GetSystemImageId() *string {
	return s.SystemImageId
}

func (s *GetTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskResponseBodyData) GetTaskStepList() []*GetTaskResponseBodyDataTaskStepList {
	return s.TaskStepList
}

func (s *GetTaskResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *GetTaskResponseBodyData) SetAliUid(v int64) *GetTaskResponseBodyData {
	s.AliUid = &v
	return s
}

func (s *GetTaskResponseBodyData) SetCurrentTaskMessage(v string) *GetTaskResponseBodyData {
	s.CurrentTaskMessage = &v
	return s
}

func (s *GetTaskResponseBodyData) SetCurrentTaskStep(v string) *GetTaskResponseBodyData {
	s.CurrentTaskStep = &v
	return s
}

func (s *GetTaskResponseBodyData) SetExecuteRegion(v string) *GetTaskResponseBodyData {
	s.ExecuteRegion = &v
	return s
}

func (s *GetTaskResponseBodyData) SetGmtCreate(v string) *GetTaskResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetTaskResponseBodyData) SetGmtModified(v string) *GetTaskResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetTaskResponseBodyData) SetPerformanceType(v string) *GetTaskResponseBodyData {
	s.PerformanceType = &v
	return s
}

func (s *GetTaskResponseBodyData) SetProgress(v string) *GetTaskResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetTaskResponseBodyData) SetResourceList(v []*GetTaskResponseBodyDataResourceList) *GetTaskResponseBodyData {
	s.ResourceList = v
	return s
}

func (s *GetTaskResponseBodyData) SetRunMode(v string) *GetTaskResponseBodyData {
	s.RunMode = &v
	return s
}

func (s *GetTaskResponseBodyData) SetStatus(v string) *GetTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyData) SetSystemDiskSize(v int32) *GetTaskResponseBodyData {
	s.SystemDiskSize = &v
	return s
}

func (s *GetTaskResponseBodyData) SetSystemImageId(v string) *GetTaskResponseBodyData {
	s.SystemImageId = &v
	return s
}

func (s *GetTaskResponseBodyData) SetTaskId(v string) *GetTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetTaskResponseBodyData) SetTaskStepList(v []*GetTaskResponseBodyDataTaskStepList) *GetTaskResponseBodyData {
	s.TaskStepList = v
	return s
}

func (s *GetTaskResponseBodyData) SetTaskType(v string) *GetTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetTaskResponseBodyData) Validate() error {
	if s.ResourceList != nil {
		for _, item := range s.ResourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskStepList != nil {
		for _, item := range s.TaskStepList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskResponseBodyDataResourceList struct {
	// The resource content corresponding to ResourceType. For the URL type, specify or return an accessible resource address.
	//
	// example:
	//
	// https://example.com/apps/business/setup.exe
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The provision method of the resource content.
	//
	// Valid values:
	//
	// - **HOSTED**: hosted resource
	//
	// - **TEXT**: text resource
	//
	// - **URL**: resource provided through a URL
	//
	// example:
	//
	// URL
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetTaskResponseBodyDataResourceList) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyDataResourceList) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyDataResourceList) GetResource() *string {
	return s.Resource
}

func (s *GetTaskResponseBodyDataResourceList) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetTaskResponseBodyDataResourceList) SetResource(v string) *GetTaskResponseBodyDataResourceList {
	s.Resource = &v
	return s
}

func (s *GetTaskResponseBodyDataResourceList) SetResourceType(v string) *GetTaskResponseBodyDataResourceList {
	s.ResourceType = &v
	return s
}

func (s *GetTaskResponseBodyDataResourceList) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyDataTaskStepList struct {
	// The end time of the step. The time is a UNIX timestamp calculated from 1970-01-01T00:00:00Z. Unit: milliseconds. This parameter may not be returned if the step has not ended.
	//
	// example:
	//
	// 1790130491000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the step. The time is a UNIX timestamp calculated from 1970-01-01T00:00:00Z. Unit: milliseconds.
	//
	// example:
	//
	// 1790130490000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The execution status of the step.
	//
	// Valid values:
	//
	// - **RUNNING**: The step is being executed.
	//
	// - **SUCCESS**: The step is executed successfully.
	//
	// - **FAILED**: The step fails to be executed.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ISO 8601 string of the step end time, including the time zone offset. In the example, +08:00 indicates UTC+8. This parameter represents the same time as EndTime and is parsed according to the offset in the returned value. This parameter may not be returned if the step has not ended.
	//
	// example:
	//
	// 2026-09-23T10:28:11+08:00
	StepEndTime *string `json:"StepEndTime,omitempty" xml:"StepEndTime,omitempty"`
	// The ISO 8601 string of the step start time, including the time zone offset. In the example, +08:00 indicates UTC+8. This parameter represents the same time as StartTime and is parsed according to the offset in the returned value.
	//
	// example:
	//
	// 2026-09-23T10:28:10+08:00
	StepStartTime *string `json:"StepStartTime,omitempty" xml:"StepStartTime,omitempty"`
	// The prompt or execution result description of this step.
	//
	// example:
	//
	// [code=Builder.Failed,httpStatusCode=<null>,message=Builder environment deploy failed.]
	TaskMessage *string `json:"TaskMessage,omitempty" xml:"TaskMessage,omitempty"`
	// The task step identifier corresponding to this record. You can call the ListTaskStep operation to query the image creation and application adaptation steps.
	//
	// example:
	//
	// CREATE_BUILDER_ENV
	TaskStep *string `json:"TaskStep,omitempty" xml:"TaskStep,omitempty"`
}

func (s GetTaskResponseBodyDataTaskStepList) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyDataTaskStepList) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyDataTaskStepList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetTaskResponseBodyDataTaskStepList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetTaskResponseBodyDataTaskStepList) GetStatus() *string {
	return s.Status
}

func (s *GetTaskResponseBodyDataTaskStepList) GetStepEndTime() *string {
	return s.StepEndTime
}

func (s *GetTaskResponseBodyDataTaskStepList) GetStepStartTime() *string {
	return s.StepStartTime
}

func (s *GetTaskResponseBodyDataTaskStepList) GetTaskMessage() *string {
	return s.TaskMessage
}

func (s *GetTaskResponseBodyDataTaskStepList) GetTaskStep() *string {
	return s.TaskStep
}

func (s *GetTaskResponseBodyDataTaskStepList) SetEndTime(v int64) *GetTaskResponseBodyDataTaskStepList {
	s.EndTime = &v
	return s
}

func (s *GetTaskResponseBodyDataTaskStepList) SetStartTime(v int64) *GetTaskResponseBodyDataTaskStepList {
	s.StartTime = &v
	return s
}

func (s *GetTaskResponseBodyDataTaskStepList) SetStatus(v string) *GetTaskResponseBodyDataTaskStepList {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyDataTaskStepList) SetStepEndTime(v string) *GetTaskResponseBodyDataTaskStepList {
	s.StepEndTime = &v
	return s
}

func (s *GetTaskResponseBodyDataTaskStepList) SetStepStartTime(v string) *GetTaskResponseBodyDataTaskStepList {
	s.StepStartTime = &v
	return s
}

func (s *GetTaskResponseBodyDataTaskStepList) SetTaskMessage(v string) *GetTaskResponseBodyDataTaskStepList {
	s.TaskMessage = &v
	return s
}

func (s *GetTaskResponseBodyDataTaskStepList) SetTaskStep(v string) *GetTaskResponseBodyDataTaskStepList {
	s.TaskStep = &v
	return s
}

func (s *GetTaskResponseBodyDataTaskStepList) Validate() error {
	return dara.Validate(s)
}

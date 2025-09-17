// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTaskResponseBody
	GetRequestId() *string
	SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody
	GetTask() *GetTaskResponseBodyTask
}

type GetTaskResponseBody struct {
	// The ID of a request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the task.
	Task *GetTaskResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
}

func (s GetTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskResponseBody) GetTask() *GetTaskResponseBodyTask {
	return s.Task
}

func (s *GetTaskResponseBody) SetRequestId(v string) *GetTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskResponseBody) SetTask(v *GetTaskResponseBodyTask) *GetTaskResponseBody {
	s.Task = v
	return s
}

func (s *GetTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTask struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2022-10-09T00:46:03Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The error returned for the task.
	Error *GetTaskResponseBodyTaskError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The code of the service.
	//
	// example:
	//
	// ECS
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The ID of the resource.
	//
	// example:
	//
	// i-8vbascjthm7kzhp3****
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The path of the resources. The relative resource ID. The resource path contains the complete resource location (parent resource/child resource).
	//
	// example:
	//
	// Instance/i-8vbascjthm7kzhp3****
	//
	// Instance/r-8vbf5abe31c9c4d4/Account/cctest
	ResourcePath *string `json:"resourcePath,omitempty" xml:"resourcePath,omitempty"`
	// The resource type.
	//
	// example:
	//
	// Instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The task state.
	//
	// Pending
	//
	// Running
	//
	// Succeeded
	//
	// Failed
	//
	// Cancelling
	//
	// Cancelled.
	//
	// example:
	//
	// Succeeded
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the task operation. Valid values: Create, Update, and Delete.
	//
	// example:
	//
	// Create
	TaskAction *string `json:"taskAction,omitempty" xml:"taskAction,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// task-433aead756057fff8189a7ce5****
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s GetTaskResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTaskResponseBodyTask) GetError() *GetTaskResponseBodyTaskError {
	return s.Error
}

func (s *GetTaskResponseBodyTask) GetProduct() *string {
	return s.Product
}

func (s *GetTaskResponseBodyTask) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTaskResponseBodyTask) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetTaskResponseBodyTask) GetResourcePath() *string {
	return s.ResourcePath
}

func (s *GetTaskResponseBodyTask) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetTaskResponseBodyTask) GetStatus() *string {
	return s.Status
}

func (s *GetTaskResponseBodyTask) GetTaskAction() *string {
	return s.TaskAction
}

func (s *GetTaskResponseBodyTask) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskResponseBodyTask) SetCreateTime(v string) *GetTaskResponseBodyTask {
	s.CreateTime = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetError(v *GetTaskResponseBodyTaskError) *GetTaskResponseBodyTask {
	s.Error = v
	return s
}

func (s *GetTaskResponseBodyTask) SetProduct(v string) *GetTaskResponseBodyTask {
	s.Product = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetRegionId(v string) *GetTaskResponseBodyTask {
	s.RegionId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetResourceId(v string) *GetTaskResponseBodyTask {
	s.ResourceId = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetResourcePath(v string) *GetTaskResponseBodyTask {
	s.ResourcePath = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetResourceType(v string) *GetTaskResponseBodyTask {
	s.ResourceType = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetStatus(v string) *GetTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskAction(v string) *GetTaskResponseBodyTask {
	s.TaskAction = &v
	return s
}

func (s *GetTaskResponseBodyTask) SetTaskId(v string) *GetTaskResponseBodyTask {
	s.TaskId = &v
	return s
}

func (s *GetTaskResponseBodyTask) Validate() error {
	return dara.Validate(s)
}

type GetTaskResponseBodyTaskError struct {
	// The error code.
	//
	// example:
	//
	// OperationFailure.OperationFailed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
	// example:
	//
	// {
	//
	//      "requestId": "123****",
	//
	//      "errorCode": "InvalidRamUser.NoPermission",
	//
	//      "errorMsg": "Ram user is not authorized to perform the operation."
	//
	// }
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetTaskResponseBodyTaskError) String() string {
	return dara.Prettify(s)
}

func (s GetTaskResponseBodyTaskError) GoString() string {
	return s.String()
}

func (s *GetTaskResponseBodyTaskError) GetCode() *string {
	return s.Code
}

func (s *GetTaskResponseBodyTaskError) GetMessage() *string {
	return s.Message
}

func (s *GetTaskResponseBodyTaskError) SetCode(v string) *GetTaskResponseBodyTaskError {
	s.Code = &v
	return s
}

func (s *GetTaskResponseBodyTaskError) SetMessage(v string) *GetTaskResponseBodyTaskError {
	s.Message = &v
	return s
}

func (s *GetTaskResponseBodyTaskError) Validate() error {
	return dara.Validate(s)
}

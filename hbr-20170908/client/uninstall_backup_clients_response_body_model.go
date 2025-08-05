// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallBackupClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UninstallBackupClientsResponseBody
	GetCode() *string
	SetInstanceStatuses(v []*UninstallBackupClientsResponseBodyInstanceStatuses) *UninstallBackupClientsResponseBody
	GetInstanceStatuses() []*UninstallBackupClientsResponseBodyInstanceStatuses
	SetMessage(v string) *UninstallBackupClientsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UninstallBackupClientsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UninstallBackupClientsResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *UninstallBackupClientsResponseBody
	GetTaskId() *string
}

type UninstallBackupClientsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status of the ECS instance.
	InstanceStatuses []*UninstallBackupClientsResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Repeated"`
	// The message that is returned. If the request is successful, a value of successful is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the asynchronous job. You can call the DescribeTask operation to query the execution result of the asynchronous job.
	//
	// example:
	//
	// t-*********************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UninstallBackupClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallBackupClientsResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UninstallBackupClientsResponseBody) GetInstanceStatuses() []*UninstallBackupClientsResponseBodyInstanceStatuses {
	return s.InstanceStatuses
}

func (s *UninstallBackupClientsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallBackupClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallBackupClientsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UninstallBackupClientsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UninstallBackupClientsResponseBody) SetCode(v string) *UninstallBackupClientsResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallBackupClientsResponseBody) SetInstanceStatuses(v []*UninstallBackupClientsResponseBodyInstanceStatuses) *UninstallBackupClientsResponseBody {
	s.InstanceStatuses = v
	return s
}

func (s *UninstallBackupClientsResponseBody) SetMessage(v string) *UninstallBackupClientsResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallBackupClientsResponseBody) SetRequestId(v string) *UninstallBackupClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallBackupClientsResponseBody) SetSuccess(v bool) *UninstallBackupClientsResponseBody {
	s.Success = &v
	return s
}

func (s *UninstallBackupClientsResponseBody) SetTaskId(v string) *UninstallBackupClientsResponseBody {
	s.TaskId = &v
	return s
}

func (s *UninstallBackupClientsResponseBody) Validate() error {
	return dara.Validate(s)
}

type UninstallBackupClientsResponseBodyInstanceStatuses struct {
	// The error code. Valid values:
	//
	// 	- If the value is empty, the request is successful.
	//
	// 	- **InstanceNotExists**: The ECS instance does not exist.
	//
	// 	- **InstanceNotRunning**: The ECS instance is not running.
	//
	// 	- **CloudAssistNotRunningOnInstance**: Cloud Assistant is unavailable.
	//
	// example:
	//
	// InstanceNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the ECS instance.
	//
	// example:
	//
	// i-0xi5w***v3j3bh2gj5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether a backup client can be installed on the ECS instance.
	//
	// 	- true: A backup client can be installed on the ECS instance.
	//
	// 	- false: A backup client cannot be installed on the ECS instance.
	//
	// example:
	//
	// true
	ValidInstance *bool `json:"ValidInstance,omitempty" xml:"ValidInstance,omitempty"`
}

func (s UninstallBackupClientsResponseBodyInstanceStatuses) String() string {
	return dara.Prettify(s)
}

func (s UninstallBackupClientsResponseBodyInstanceStatuses) GoString() string {
	return s.String()
}

func (s *UninstallBackupClientsResponseBodyInstanceStatuses) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UninstallBackupClientsResponseBodyInstanceStatuses) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UninstallBackupClientsResponseBodyInstanceStatuses) GetValidInstance() *bool {
	return s.ValidInstance
}

func (s *UninstallBackupClientsResponseBodyInstanceStatuses) SetErrorCode(v string) *UninstallBackupClientsResponseBodyInstanceStatuses {
	s.ErrorCode = &v
	return s
}

func (s *UninstallBackupClientsResponseBodyInstanceStatuses) SetInstanceId(v string) *UninstallBackupClientsResponseBodyInstanceStatuses {
	s.InstanceId = &v
	return s
}

func (s *UninstallBackupClientsResponseBodyInstanceStatuses) SetValidInstance(v bool) *UninstallBackupClientsResponseBodyInstanceStatuses {
	s.ValidInstance = &v
	return s
}

func (s *UninstallBackupClientsResponseBodyInstanceStatuses) Validate() error {
	return dara.Validate(s)
}

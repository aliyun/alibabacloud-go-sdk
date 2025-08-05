// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallBackupClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InstallBackupClientsResponseBody
	GetCode() *string
	SetInstanceStatuses(v []*InstallBackupClientsResponseBodyInstanceStatuses) *InstallBackupClientsResponseBody
	GetInstanceStatuses() []*InstallBackupClientsResponseBodyInstanceStatuses
	SetMessage(v string) *InstallBackupClientsResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallBackupClientsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InstallBackupClientsResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *InstallBackupClientsResponseBody
	GetTaskId() *string
}

type InstallBackupClientsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status of the ECS instance.
	InstanceStatuses []*InstallBackupClientsResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Repeated"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
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
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the asynchronous job. You can call the DescribeTask operation to query the execution result of an asynchronous job.
	//
	// example:
	//
	// t-*********************
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s InstallBackupClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallBackupClientsResponseBody) GoString() string {
	return s.String()
}

func (s *InstallBackupClientsResponseBody) GetCode() *string {
	return s.Code
}

func (s *InstallBackupClientsResponseBody) GetInstanceStatuses() []*InstallBackupClientsResponseBodyInstanceStatuses {
	return s.InstanceStatuses
}

func (s *InstallBackupClientsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallBackupClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallBackupClientsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InstallBackupClientsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *InstallBackupClientsResponseBody) SetCode(v string) *InstallBackupClientsResponseBody {
	s.Code = &v
	return s
}

func (s *InstallBackupClientsResponseBody) SetInstanceStatuses(v []*InstallBackupClientsResponseBodyInstanceStatuses) *InstallBackupClientsResponseBody {
	s.InstanceStatuses = v
	return s
}

func (s *InstallBackupClientsResponseBody) SetMessage(v string) *InstallBackupClientsResponseBody {
	s.Message = &v
	return s
}

func (s *InstallBackupClientsResponseBody) SetRequestId(v string) *InstallBackupClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallBackupClientsResponseBody) SetSuccess(v bool) *InstallBackupClientsResponseBody {
	s.Success = &v
	return s
}

func (s *InstallBackupClientsResponseBody) SetTaskId(v string) *InstallBackupClientsResponseBody {
	s.TaskId = &v
	return s
}

func (s *InstallBackupClientsResponseBody) Validate() error {
	return dara.Validate(s)
}

type InstallBackupClientsResponseBodyInstanceStatuses struct {
	// The error code that is returned. Valid values:
	//
	// 	- If the value is empty, the call is successful.
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
	// Indicates whether an HBR client can be installed on the ECS instance. Valid values:
	//
	// 	- true: An HBR client can be installed on the ECS instance.
	//
	// 	- false: An HBR client cannot be installed on the ECS instance.
	//
	// example:
	//
	// true
	ValidInstance *bool `json:"ValidInstance,omitempty" xml:"ValidInstance,omitempty"`
}

func (s InstallBackupClientsResponseBodyInstanceStatuses) String() string {
	return dara.Prettify(s)
}

func (s InstallBackupClientsResponseBodyInstanceStatuses) GoString() string {
	return s.String()
}

func (s *InstallBackupClientsResponseBodyInstanceStatuses) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *InstallBackupClientsResponseBodyInstanceStatuses) GetInstanceId() *string {
	return s.InstanceId
}

func (s *InstallBackupClientsResponseBodyInstanceStatuses) GetValidInstance() *bool {
	return s.ValidInstance
}

func (s *InstallBackupClientsResponseBodyInstanceStatuses) SetErrorCode(v string) *InstallBackupClientsResponseBodyInstanceStatuses {
	s.ErrorCode = &v
	return s
}

func (s *InstallBackupClientsResponseBodyInstanceStatuses) SetInstanceId(v string) *InstallBackupClientsResponseBodyInstanceStatuses {
	s.InstanceId = &v
	return s
}

func (s *InstallBackupClientsResponseBodyInstanceStatuses) SetValidInstance(v bool) *InstallBackupClientsResponseBodyInstanceStatuses {
	s.ValidInstance = &v
	return s
}

func (s *InstallBackupClientsResponseBodyInstanceStatuses) Validate() error {
	return dara.Validate(s)
}

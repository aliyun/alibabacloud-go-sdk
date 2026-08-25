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
	// The return code. A value of 200 indicates that the operation is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status of the ECS instances.
	InstanceStatuses []*InstallBackupClientsResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Repeated"`
	// The message that is returned. If the request is successful, successful is returned. If the request fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// - true: The request is successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the asynchronous task. Call the DescribeTask operation to query the task result.
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
	if s.InstanceStatuses != nil {
		for _, item := range s.InstanceStatuses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InstallBackupClientsResponseBodyInstanceStatuses struct {
	// The error code. Valid values:
	//
	// - An empty value indicates that the operation is successful.
	//
	// - **InstanceNotExists**: The ECS instance does not exist.
	//
	// - **InstanceNotRunning**: The ECS instance is not in the Running state.
	//
	// - **CloudAssistNotRunningOnInstance**: Cloud Assistant is not available.
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
	// - true: The backup client can be installed.
	//
	// - false: The backup client cannot be installed.
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeBackupClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpgradeBackupClientsResponseBody
	GetCode() *string
	SetInstanceStatuses(v []*UpgradeBackupClientsResponseBodyInstanceStatuses) *UpgradeBackupClientsResponseBody
	GetInstanceStatuses() []*UpgradeBackupClientsResponseBodyInstanceStatuses
	SetMessage(v string) *UpgradeBackupClientsResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeBackupClientsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeBackupClientsResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *UpgradeBackupClientsResponseBody
	GetTaskId() *string
}

type UpgradeBackupClientsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status of the ECS instance. If the status of an ECS instance cannot meet the requirements to install an HBR client and the value of the InstanceIds parameter is greater than 1, an error message is returned based on the value of this parameter.
	InstanceStatuses []*UpgradeBackupClientsResponseBodyInstanceStatuses `json:"InstanceStatuses,omitempty" xml:"InstanceStatuses,omitempty" type:"Repeated"`
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

func (s UpgradeBackupClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeBackupClientsResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeBackupClientsResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeBackupClientsResponseBody) GetInstanceStatuses() []*UpgradeBackupClientsResponseBodyInstanceStatuses {
	return s.InstanceStatuses
}

func (s *UpgradeBackupClientsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeBackupClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeBackupClientsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeBackupClientsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *UpgradeBackupClientsResponseBody) SetCode(v string) *UpgradeBackupClientsResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeBackupClientsResponseBody) SetInstanceStatuses(v []*UpgradeBackupClientsResponseBodyInstanceStatuses) *UpgradeBackupClientsResponseBody {
	s.InstanceStatuses = v
	return s
}

func (s *UpgradeBackupClientsResponseBody) SetMessage(v string) *UpgradeBackupClientsResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeBackupClientsResponseBody) SetRequestId(v string) *UpgradeBackupClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeBackupClientsResponseBody) SetSuccess(v bool) *UpgradeBackupClientsResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeBackupClientsResponseBody) SetTaskId(v string) *UpgradeBackupClientsResponseBody {
	s.TaskId = &v
	return s
}

func (s *UpgradeBackupClientsResponseBody) Validate() error {
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

type UpgradeBackupClientsResponseBodyInstanceStatuses struct {
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

func (s UpgradeBackupClientsResponseBodyInstanceStatuses) String() string {
	return dara.Prettify(s)
}

func (s UpgradeBackupClientsResponseBodyInstanceStatuses) GoString() string {
	return s.String()
}

func (s *UpgradeBackupClientsResponseBodyInstanceStatuses) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpgradeBackupClientsResponseBodyInstanceStatuses) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpgradeBackupClientsResponseBodyInstanceStatuses) GetValidInstance() *bool {
	return s.ValidInstance
}

func (s *UpgradeBackupClientsResponseBodyInstanceStatuses) SetErrorCode(v string) *UpgradeBackupClientsResponseBodyInstanceStatuses {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeBackupClientsResponseBodyInstanceStatuses) SetInstanceId(v string) *UpgradeBackupClientsResponseBodyInstanceStatuses {
	s.InstanceId = &v
	return s
}

func (s *UpgradeBackupClientsResponseBodyInstanceStatuses) SetValidInstance(v bool) *UpgradeBackupClientsResponseBodyInstanceStatuses {
	s.ValidInstance = &v
	return s
}

func (s *UpgradeBackupClientsResponseBodyInstanceStatuses) Validate() error {
	return dara.Validate(s)
}

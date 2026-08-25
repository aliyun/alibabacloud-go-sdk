// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTaskStatusesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListVirusScanTaskStatusesResponseBody
	GetRequestId() *string
	SetTasks(v []*ListVirusScanTaskStatusesResponseBodyTasks) *ListVirusScanTaskStatusesResponseBody
	GetTasks() []*ListVirusScanTaskStatusesResponseBodyTasks
}

type ListVirusScanTaskStatusesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3D7EC0AF-DB2A-5D9C-90EC-F090A6BAAEA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of execution progress for virus scan tasks.
	Tasks []*ListVirusScanTaskStatusesResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListVirusScanTaskStatusesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskStatusesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskStatusesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirusScanTaskStatusesResponseBody) GetTasks() []*ListVirusScanTaskStatusesResponseBodyTasks {
	return s.Tasks
}

func (s *ListVirusScanTaskStatusesResponseBody) SetRequestId(v string) *ListVirusScanTaskStatusesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirusScanTaskStatusesResponseBody) SetTasks(v []*ListVirusScanTaskStatusesResponseBodyTasks) *ListVirusScanTaskStatusesResponseBody {
	s.Tasks = v
	return s
}

func (s *ListVirusScanTaskStatusesResponseBody) Validate() error {
	if s.Tasks != nil {
		for _, item := range s.Tasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusScanTaskStatusesResponseBodyTasks struct {
	// The virus scan task ID.
	//
	// example:
	//
	// v1:1024772
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The execution progress measured by device count.
	TaskStatus *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty" type:"Struct"`
}

func (s ListVirusScanTaskStatusesResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskStatusesResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskStatusesResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ListVirusScanTaskStatusesResponseBodyTasks) GetTaskStatus() *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus {
	return s.TaskStatus
}

func (s *ListVirusScanTaskStatusesResponseBodyTasks) SetTaskId(v string) *ListVirusScanTaskStatusesResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListVirusScanTaskStatusesResponseBodyTasks) SetTaskStatus(v *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) *ListVirusScanTaskStatusesResponseBodyTasks {
	s.TaskStatus = v
	return s
}

func (s *ListVirusScanTaskStatusesResponseBodyTasks) Validate() error {
	if s.TaskStatus != nil {
		if err := s.TaskStatus.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVirusScanTaskStatusesResponseBodyTasksTaskStatus struct {
	// The number of user terminal devices that have received the task.
	//
	// example:
	//
	// 120
	DeviceAckCount *int32 `json:"DeviceAckCount,omitempty" xml:"DeviceAckCount,omitempty"`
	// The number of user terminal devices on which the scan execution failed.
	//
	// example:
	//
	// 5
	DeviceResultFailCount *int32 `json:"DeviceResultFailCount,omitempty" xml:"DeviceResultFailCount,omitempty"`
	// The number of user terminal devices on which the scan was executed successfully.
	//
	// example:
	//
	// 100
	DeviceResultSuccessCount *int32 `json:"DeviceResultSuccessCount,omitempty" xml:"DeviceResultSuccessCount,omitempty"`
	// The number of user terminal devices that are currently executing the scan.
	//
	// example:
	//
	// 15
	DeviceStartCount *int32 `json:"DeviceStartCount,omitempty" xml:"DeviceStartCount,omitempty"`
}

func (s ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) GetDeviceAckCount() *int32 {
	return s.DeviceAckCount
}

func (s *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) GetDeviceResultFailCount() *int32 {
	return s.DeviceResultFailCount
}

func (s *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) GetDeviceResultSuccessCount() *int32 {
	return s.DeviceResultSuccessCount
}

func (s *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) GetDeviceStartCount() *int32 {
	return s.DeviceStartCount
}

func (s *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) SetDeviceAckCount(v int32) *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus {
	s.DeviceAckCount = &v
	return s
}

func (s *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) SetDeviceResultFailCount(v int32) *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus {
	s.DeviceResultFailCount = &v
	return s
}

func (s *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) SetDeviceResultSuccessCount(v int32) *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus {
	s.DeviceResultSuccessCount = &v
	return s
}

func (s *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) SetDeviceStartCount(v int32) *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus {
	s.DeviceStartCount = &v
	return s
}

func (s *ListVirusScanTaskStatusesResponseBodyTasksTaskStatus) Validate() error {
	return dara.Validate(s)
}

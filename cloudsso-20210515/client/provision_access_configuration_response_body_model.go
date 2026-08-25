// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProvisionAccessConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ProvisionAccessConfigurationResponseBody
	GetRequestId() *string
	SetTasks(v []*ProvisionAccessConfigurationResponseBodyTasks) *ProvisionAccessConfigurationResponseBody
	GetTasks() []*ProvisionAccessConfigurationResponseBodyTasks
}

type ProvisionAccessConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DFDC16B2-4509-5FA6-9FA5-3CD35E4292FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	Tasks []*ProvisionAccessConfigurationResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ProvisionAccessConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProvisionAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ProvisionAccessConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProvisionAccessConfigurationResponseBody) GetTasks() []*ProvisionAccessConfigurationResponseBodyTasks {
	return s.Tasks
}

func (s *ProvisionAccessConfigurationResponseBody) SetRequestId(v string) *ProvisionAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBody) SetTasks(v []*ProvisionAccessConfigurationResponseBodyTasks) *ProvisionAccessConfigurationResponseBody {
	s.Tasks = v
	return s
}

func (s *ProvisionAccessConfigurationResponseBody) Validate() error {
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

type ProvisionAccessConfigurationResponseBodyTasks struct {
	// The ID of the access configuration.
	//
	// example:
	//
	// ac-00jhtfl8thteu6uj****
	AccessConfigurationId *string `json:"AccessConfigurationId,omitempty" xml:"AccessConfigurationId,omitempty"`
	// The name of the access configuration.
	//
	// example:
	//
	// ECS-Admin
	AccessConfigurationName *string `json:"AccessConfigurationName,omitempty" xml:"AccessConfigurationName,omitempty"`
	// example:
	//
	// 114240524784****
	OriginTargetId *string `json:"OriginTargetId,omitempty" xml:"OriginTargetId,omitempty"`
	// The task status. Valid values:
	//
	// - InProgress: The task is running.
	//
	// - Success: The task is successful.
	//
	// - Failed: The task failed.
	//
	// example:
	//
	// InProgress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the task object.
	//
	// example:
	//
	// 114240524784****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the task object.
	//
	// example:
	//
	// dev-test
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The path ID of the task object in the resource directory.
	TargetPath *string `json:"TargetPath,omitempty" xml:"TargetPath,omitempty"`
	// The path name of the task object in the resource directory.
	TargetPathName *string `json:"TargetPathName,omitempty" xml:"TargetPathName,omitempty"`
	// The type of the task object. The value is fixed as RD-Account, which indicates the accounts in the resource directory.
	//
	// example:
	//
	// RD-Account
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-shqlhd8uvt280rtm****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. The value is fixed as ProvisionAccessConfiguration, which indicates that an access configuration is provisioned.
	//
	// example:
	//
	// ProvisionAccessConfiguration
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ProvisionAccessConfigurationResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s ProvisionAccessConfigurationResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) GetTargetId() *string {
	return s.TargetId
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) GetTargetName() *string {
	return s.TargetName
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) GetTargetPath() *string {
	return s.TargetPath
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) GetTargetPathName() *string {
	return s.TargetPathName
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) GetTargetType() *string {
	return s.TargetType
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetAccessConfigurationId(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.AccessConfigurationId = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetAccessConfigurationName(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.AccessConfigurationName = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetOriginTargetId(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.OriginTargetId = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetStatus(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTargetId(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TargetId = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTargetName(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TargetName = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTargetPath(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TargetPath = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTargetPathName(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TargetPathName = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTargetType(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TargetType = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTaskId(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) SetTaskType(v string) *ProvisionAccessConfigurationResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *ProvisionAccessConfigurationResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}

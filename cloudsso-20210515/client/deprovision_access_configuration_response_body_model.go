// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeprovisionAccessConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeprovisionAccessConfigurationResponseBody
	GetRequestId() *string
	SetTasks(v []*DeprovisionAccessConfigurationResponseBodyTasks) *DeprovisionAccessConfigurationResponseBody
	GetTasks() []*DeprovisionAccessConfigurationResponseBodyTasks
}

type DeprovisionAccessConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 584FE9D0-D1AC-5B19-A39C-8D244FC0538C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task information.
	Tasks []*DeprovisionAccessConfigurationResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s DeprovisionAccessConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeprovisionAccessConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeprovisionAccessConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeprovisionAccessConfigurationResponseBody) GetTasks() []*DeprovisionAccessConfigurationResponseBodyTasks {
	return s.Tasks
}

func (s *DeprovisionAccessConfigurationResponseBody) SetRequestId(v string) *DeprovisionAccessConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBody) SetTasks(v []*DeprovisionAccessConfigurationResponseBodyTasks) *DeprovisionAccessConfigurationResponseBody {
	s.Tasks = v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBody) Validate() error {
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

type DeprovisionAccessConfigurationResponseBodyTasks struct {
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
	// t-sh0655wnq8pdlrlc****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type. The value is fixed as DeprovisionAccessConfiguration, which indicates that the access configuration is de-provisioned.
	//
	// example:
	//
	// DeprovisionAccessConfiguration
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DeprovisionAccessConfigurationResponseBodyTasks) String() string {
	return dara.Prettify(s)
}

func (s DeprovisionAccessConfigurationResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) GetAccessConfigurationId() *string {
	return s.AccessConfigurationId
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) GetAccessConfigurationName() *string {
	return s.AccessConfigurationName
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) GetOriginTargetId() *string {
	return s.OriginTargetId
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) GetStatus() *string {
	return s.Status
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) GetTargetId() *string {
	return s.TargetId
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) GetTargetName() *string {
	return s.TargetName
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) GetTargetPath() *string {
	return s.TargetPath
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) GetTargetPathName() *string {
	return s.TargetPathName
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) GetTargetType() *string {
	return s.TargetType
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetAccessConfigurationId(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.AccessConfigurationId = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetAccessConfigurationName(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.AccessConfigurationName = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetOriginTargetId(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.OriginTargetId = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetStatus(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTargetId(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TargetId = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTargetName(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TargetName = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTargetPath(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TargetPath = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTargetPathName(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TargetPathName = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTargetType(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TargetType = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTaskId(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) SetTaskType(v string) *DeprovisionAccessConfigurationResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *DeprovisionAccessConfigurationResponseBodyTasks) Validate() error {
	return dara.Validate(s)
}

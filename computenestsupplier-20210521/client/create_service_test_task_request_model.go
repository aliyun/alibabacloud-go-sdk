// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTestTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateServiceTestTaskRequest
	GetRegionId() *string
	SetTaskName(v string) *CreateServiceTestTaskRequest
	GetTaskName() *string
	SetTaskRegionId(v string) *CreateServiceTestTaskRequest
	GetTaskRegionId() *string
	SetTestCaseIds(v []*string) *CreateServiceTestTaskRequest
	GetTestCaseIds() []*string
}

type CreateServiceTestTaskRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// nametest
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The Task Execution Region
	//
	// example:
	//
	// cn-beijing
	TaskRegionId *string `json:"TaskRegionId,omitempty" xml:"TaskRegionId,omitempty"`
	// The service test case ids.
	//
	// This parameter is required.
	TestCaseIds []*string `json:"TestCaseIds,omitempty" xml:"TestCaseIds,omitempty" type:"Repeated"`
}

func (s CreateServiceTestTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTestTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceTestTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceTestTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateServiceTestTaskRequest) GetTaskRegionId() *string {
	return s.TaskRegionId
}

func (s *CreateServiceTestTaskRequest) GetTestCaseIds() []*string {
	return s.TestCaseIds
}

func (s *CreateServiceTestTaskRequest) SetRegionId(v string) *CreateServiceTestTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceTestTaskRequest) SetTaskName(v string) *CreateServiceTestTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateServiceTestTaskRequest) SetTaskRegionId(v string) *CreateServiceTestTaskRequest {
	s.TaskRegionId = &v
	return s
}

func (s *CreateServiceTestTaskRequest) SetTestCaseIds(v []*string) *CreateServiceTestTaskRequest {
	s.TestCaseIds = v
	return s
}

func (s *CreateServiceTestTaskRequest) Validate() error {
	return dara.Validate(s)
}

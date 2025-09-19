// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulCheckTaskStatusDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVulCheckTaskStatusDetailResponseBody
	GetRequestId() *string
	SetTaskStatuses(v []*DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses) *DescribeVulCheckTaskStatusDetailResponseBody
	GetTaskStatuses() []*DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses
	SetTotalCount(v int32) *DescribeVulCheckTaskStatusDetailResponseBody
	GetTotalCount() *int32
}

type DescribeVulCheckTaskStatusDetailResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the status information about the vulnerability scan tasks on the server.
	TaskStatuses []*DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses `json:"TaskStatuses,omitempty" xml:"TaskStatuses,omitempty" type:"Repeated"`
	// The total number of vulnerability scan tasks on the server.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVulCheckTaskStatusDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulCheckTaskStatusDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulCheckTaskStatusDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulCheckTaskStatusDetailResponseBody) GetTaskStatuses() []*DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses {
	return s.TaskStatuses
}

func (s *DescribeVulCheckTaskStatusDetailResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVulCheckTaskStatusDetailResponseBody) SetRequestId(v string) *DescribeVulCheckTaskStatusDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailResponseBody) SetTaskStatuses(v []*DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses) *DescribeVulCheckTaskStatusDetailResponseBody {
	s.TaskStatuses = v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailResponseBody) SetTotalCount(v int32) *DescribeVulCheckTaskStatusDetailResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses struct {
	// The ID of the main task.
	//
	// example:
	//
	// 16190385
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// An array that consists of status information about the vulnerability scan subtask.
	TaskStatusList []*DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList `json:"TaskStatusList,omitempty" xml:"TaskStatusList,omitempty" type:"Repeated"`
}

func (s DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses) GoString() string {
	return s.String()
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses) GetTaskStatusList() []*DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList {
	return s.TaskStatusList
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses) SetTaskId(v string) *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses {
	s.TaskId = &v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses) SetTaskStatusList(v []*DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList) *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses {
	s.TaskStatusList = v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses) Validate() error {
	return dara.Validate(s)
}

type DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList struct {
	// The error code returned.
	//
	// example:
	//
	// push_command_failed
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The status of the subtask. Valid values:
	//
	// 	- **0**: unhandled
	//
	// 	- **1**: collecting
	//
	// 	- **2**: collected
	//
	// 	- **3**: matching
	//
	// 	- **4**: complete
	//
	// example:
	//
	// 4
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability
	//
	// 	- **sca**: vulnerability that is detected based on software component analysis
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList) GoString() string {
	return s.String()
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList) GetCode() *string {
	return s.Code
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList) GetStatus() *string {
	return s.Status
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList) GetType() *string {
	return s.Type
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList) SetCode(v string) *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList {
	s.Code = &v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList) SetStatus(v string) *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList {
	s.Status = &v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList) SetType(v string) *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList {
	s.Type = &v
	return s
}

func (s *DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList) Validate() error {
	return dara.Validate(s)
}

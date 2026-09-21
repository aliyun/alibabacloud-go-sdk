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
	// The request ID.
	//
	// example:
	//
	// BE120DAB-F4E7-4C53-ADC3-A97578ABF384
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of vulnerability task statuses for the server.
	TaskStatuses []*DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses `json:"TaskStatuses,omitempty" xml:"TaskStatuses,omitempty" type:"Repeated"`
	// The total number of vulnerability subtasks for the server.
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
	if s.TaskStatuses != nil {
		for _, item := range s.TaskStatuses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVulCheckTaskStatusDetailResponseBodyTaskStatuses struct {
	// The main task ID.
	//
	// example:
	//
	// 16190385
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The list of vulnerability detection task statuses.
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
	if s.TaskStatusList != nil {
		for _, item := range s.TaskStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVulCheckTaskStatusDetailResponseBodyTaskStatusesTaskStatusList struct {
	// The failure code.
	//
	// example:
	//
	// push_command_failed
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The check completion status. Valid values:
	//
	// - **0**: Unprocessed.
	//
	// - **1**: Collecting.
	//
	// - **2**: Collection completed.
	//
	// - **3**: Matching.
	//
	// - **4**: Completed.
	//
	// example:
	//
	// 4
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vulnerability type. Valid values:
	//
	// - **cve**: Linux software vulnerability
	//
	// - **sys**: Windows system vulnerability
	//
	// - **cms**: Web-CMS vulnerability
	//
	// - **sca**: SCA vulnerability
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

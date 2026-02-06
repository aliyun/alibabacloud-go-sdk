// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTimedScheduleTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeParameterTimedScheduleTaskResponseBody
	GetRequestId() *string
	SetTaskList(v []*DescribeParameterTimedScheduleTaskResponseBodyTaskList) *DescribeParameterTimedScheduleTaskResponseBody
	GetTaskList() []*DescribeParameterTimedScheduleTaskResponseBodyTaskList
}

type DescribeParameterTimedScheduleTaskResponseBody struct {
	// example:
	//
	// A807C95D-410C-5BB5-96C0-C6E09F2C3D36
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskList  []*DescribeParameterTimedScheduleTaskResponseBodyTaskList `json:"TaskList,omitempty" xml:"TaskList,omitempty" type:"Repeated"`
}

func (s DescribeParameterTimedScheduleTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTimedScheduleTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParameterTimedScheduleTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeParameterTimedScheduleTaskResponseBody) GetTaskList() []*DescribeParameterTimedScheduleTaskResponseBodyTaskList {
	return s.TaskList
}

func (s *DescribeParameterTimedScheduleTaskResponseBody) SetRequestId(v string) *DescribeParameterTimedScheduleTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParameterTimedScheduleTaskResponseBody) SetTaskList(v []*DescribeParameterTimedScheduleTaskResponseBodyTaskList) *DescribeParameterTimedScheduleTaskResponseBody {
	s.TaskList = v
	return s
}

func (s *DescribeParameterTimedScheduleTaskResponseBody) Validate() error {
	if s.TaskList != nil {
		for _, item := range s.TaskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeParameterTimedScheduleTaskResponseBodyTaskList struct {
	// example:
	//
	// rm-2ze2za3is7baay****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// {"auto_increment_increment":"1000","back_log":"99"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// PENDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2022-05-06T09:24:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// example:
	//
	// 27056921
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeParameterTimedScheduleTaskResponseBodyTaskList) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTimedScheduleTaskResponseBodyTaskList) GoString() string {
	return s.String()
}

func (s *DescribeParameterTimedScheduleTaskResponseBodyTaskList) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeParameterTimedScheduleTaskResponseBodyTaskList) GetParameters() *string {
	return s.Parameters
}

func (s *DescribeParameterTimedScheduleTaskResponseBodyTaskList) GetStatus() *string {
	return s.Status
}

func (s *DescribeParameterTimedScheduleTaskResponseBodyTaskList) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *DescribeParameterTimedScheduleTaskResponseBodyTaskList) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeParameterTimedScheduleTaskResponseBodyTaskList) SetDBInstanceName(v string) *DescribeParameterTimedScheduleTaskResponseBodyTaskList {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeParameterTimedScheduleTaskResponseBodyTaskList) SetParameters(v string) *DescribeParameterTimedScheduleTaskResponseBodyTaskList {
	s.Parameters = &v
	return s
}

func (s *DescribeParameterTimedScheduleTaskResponseBodyTaskList) SetStatus(v string) *DescribeParameterTimedScheduleTaskResponseBodyTaskList {
	s.Status = &v
	return s
}

func (s *DescribeParameterTimedScheduleTaskResponseBodyTaskList) SetSwitchTime(v string) *DescribeParameterTimedScheduleTaskResponseBodyTaskList {
	s.SwitchTime = &v
	return s
}

func (s *DescribeParameterTimedScheduleTaskResponseBodyTaskList) SetTaskId(v string) *DescribeParameterTimedScheduleTaskResponseBodyTaskList {
	s.TaskId = &v
	return s
}

func (s *DescribeParameterTimedScheduleTaskResponseBodyTaskList) Validate() error {
	return dara.Validate(s)
}

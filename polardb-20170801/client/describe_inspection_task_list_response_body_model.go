// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeInspectionTaskListResponseBody
	GetRequestId() *string
	SetTaskLists(v []*DescribeInspectionTaskListResponseBodyTaskLists) *DescribeInspectionTaskListResponseBody
	GetTaskLists() []*DescribeInspectionTaskListResponseBodyTaskLists
}

type DescribeInspectionTaskListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of inspection tasks.
	TaskLists []*DescribeInspectionTaskListResponseBodyTaskLists `json:"TaskLists,omitempty" xml:"TaskLists,omitempty" type:"Repeated"`
}

func (s DescribeInspectionTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInspectionTaskListResponseBody) GetTaskLists() []*DescribeInspectionTaskListResponseBodyTaskLists {
	return s.TaskLists
}

func (s *DescribeInspectionTaskListResponseBody) SetRequestId(v string) *DescribeInspectionTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInspectionTaskListResponseBody) SetTaskLists(v []*DescribeInspectionTaskListResponseBodyTaskLists) *DescribeInspectionTaskListResponseBody {
	s.TaskLists = v
	return s
}

func (s *DescribeInspectionTaskListResponseBody) Validate() error {
	if s.TaskLists != nil {
		for _, item := range s.TaskLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInspectionTaskListResponseBodyTaskLists struct {
	// The last modification date of the task.
	//
	// example:
	//
	// 2024-06-01
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cron expression for periodically executing the inspection.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The IDs of the PolarDB clusters to be inspected. Multiple IDs are separated by commas.
	//
	// example:
	//
	// pc-001,pc-002
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The expiration date of the inspection task.
	//
	// example:
	//
	// 2025-12-31
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The description of the inspection task.
	//
	// example:
	//
	// Periodic inspection task
	TaskDescription *string `json:"TaskDescription,omitempty" xml:"TaskDescription,omitempty"`
	// The ID of the inspection task.
	//
	// example:
	//
	// task-001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeInspectionTaskListResponseBodyTaskLists) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskListResponseBodyTaskLists) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) GetDBClusterIds() *string {
	return s.DBClusterIds
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) SetCreateTime(v string) *DescribeInspectionTaskListResponseBodyTaskLists {
	s.CreateTime = &v
	return s
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) SetCronExpression(v string) *DescribeInspectionTaskListResponseBodyTaskLists {
	s.CronExpression = &v
	return s
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) SetDBClusterIds(v string) *DescribeInspectionTaskListResponseBodyTaskLists {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) SetExpireTime(v string) *DescribeInspectionTaskListResponseBodyTaskLists {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) SetTaskDescription(v string) *DescribeInspectionTaskListResponseBodyTaskLists {
	s.TaskDescription = &v
	return s
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) SetTaskId(v string) *DescribeInspectionTaskListResponseBodyTaskLists {
	s.TaskId = &v
	return s
}

func (s *DescribeInspectionTaskListResponseBodyTaskLists) Validate() error {
	return dara.Validate(s)
}

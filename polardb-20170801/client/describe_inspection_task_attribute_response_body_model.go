// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTaskAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeInspectionTaskAttributeResponseBody
	GetCreateTime() *string
	SetCronExpression(v string) *DescribeInspectionTaskAttributeResponseBody
	GetCronExpression() *string
	SetDBClusterIds(v string) *DescribeInspectionTaskAttributeResponseBody
	GetDBClusterIds() *string
	SetExpireTime(v string) *DescribeInspectionTaskAttributeResponseBody
	GetExpireTime() *string
	SetModifyTime(v string) *DescribeInspectionTaskAttributeResponseBody
	GetModifyTime() *string
	SetRequestId(v string) *DescribeInspectionTaskAttributeResponseBody
	GetRequestId() *string
	SetTaskDescription(v string) *DescribeInspectionTaskAttributeResponseBody
	GetTaskDescription() *string
	SetTaskId(v string) *DescribeInspectionTaskAttributeResponseBody
	GetTaskId() *string
}

type DescribeInspectionTaskAttributeResponseBody struct {
	// The creation date of the inspection task. The format is YYYY-MM-DD.
	//
	// example:
	//
	// 2024-06-01
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cron expression for the scheduled inspection.
	//
	// example:
	//
	// 0 0 2 	- 	- ?
	CronExpression *string `json:"CronExpression,omitempty" xml:"CronExpression,omitempty"`
	// The IDs of the PolarDB clusters to be inspected. Separate multiple IDs with commas.
	//
	// example:
	//
	// pc-001,pc-002
	DBClusterIds *string `json:"DBClusterIds,omitempty" xml:"DBClusterIds,omitempty"`
	// The expiration date of the inspection task. The format is YYYY-MM-DD.
	//
	// example:
	//
	// 2025-12-31
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The last modification date of the inspection task. The format is YYYY-MM-DD.
	//
	// example:
	//
	// 2024-06-01
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The description of the inspection task.
	//
	// example:
	//
	// Scheduled inspection task
	TaskDescription *string `json:"TaskDescription,omitempty" xml:"TaskDescription,omitempty"`
	// The ID of the inspection task.
	//
	// example:
	//
	// task-001
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeInspectionTaskAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeInspectionTaskAttributeResponseBody) GetCronExpression() *string {
	return s.CronExpression
}

func (s *DescribeInspectionTaskAttributeResponseBody) GetDBClusterIds() *string {
	return s.DBClusterIds
}

func (s *DescribeInspectionTaskAttributeResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeInspectionTaskAttributeResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeInspectionTaskAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInspectionTaskAttributeResponseBody) GetTaskDescription() *string {
	return s.TaskDescription
}

func (s *DescribeInspectionTaskAttributeResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeInspectionTaskAttributeResponseBody) SetCreateTime(v string) *DescribeInspectionTaskAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeInspectionTaskAttributeResponseBody) SetCronExpression(v string) *DescribeInspectionTaskAttributeResponseBody {
	s.CronExpression = &v
	return s
}

func (s *DescribeInspectionTaskAttributeResponseBody) SetDBClusterIds(v string) *DescribeInspectionTaskAttributeResponseBody {
	s.DBClusterIds = &v
	return s
}

func (s *DescribeInspectionTaskAttributeResponseBody) SetExpireTime(v string) *DescribeInspectionTaskAttributeResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInspectionTaskAttributeResponseBody) SetModifyTime(v string) *DescribeInspectionTaskAttributeResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeInspectionTaskAttributeResponseBody) SetRequestId(v string) *DescribeInspectionTaskAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInspectionTaskAttributeResponseBody) SetTaskDescription(v string) *DescribeInspectionTaskAttributeResponseBody {
	s.TaskDescription = &v
	return s
}

func (s *DescribeInspectionTaskAttributeResponseBody) SetTaskId(v string) *DescribeInspectionTaskAttributeResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeInspectionTaskAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}

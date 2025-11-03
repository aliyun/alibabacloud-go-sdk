// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeActiveOperationTaskResponseBodyItems) *DescribeActiveOperationTaskResponseBody
	GetItems() []*DescribeActiveOperationTaskResponseBodyItems
	SetPageNumber(v int32) *DescribeActiveOperationTaskResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeActiveOperationTaskResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeActiveOperationTaskResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeActiveOperationTaskResponseBody
	GetTotalRecordCount() *int32
}

type DescribeActiveOperationTaskResponseBody struct {
	// The list of tasks.
	Items []*DescribeActiveOperationTaskResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6FD47DC4-0750-5524-A89E-E7D559Cxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 0
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeActiveOperationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskResponseBody) GetItems() []*DescribeActiveOperationTaskResponseBodyItems {
	return s.Items
}

func (s *DescribeActiveOperationTaskResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeActiveOperationTaskResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeActiveOperationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActiveOperationTaskResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeActiveOperationTaskResponseBody) SetItems(v []*DescribeActiveOperationTaskResponseBodyItems) *DescribeActiveOperationTaskResponseBody {
	s.Items = v
	return s
}

func (s *DescribeActiveOperationTaskResponseBody) SetPageNumber(v int32) *DescribeActiveOperationTaskResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBody) SetPageSize(v int32) *DescribeActiveOperationTaskResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBody) SetRequestId(v string) *DescribeActiveOperationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBody) SetTotalRecordCount(v int32) *DescribeActiveOperationTaskResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeActiveOperationTaskResponseBodyItems struct {
	// The time when the task was created. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-11-06T13:11:08Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The database type of the instance.
	//
	// example:
	//
	// mongodb
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The deadline before which the time to preform the task can be modified. The time in UTC is displayed in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2023-03-29T13:59:59Z
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 50xx
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dds-bp12721xxxx9b914
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// The time when the task was modified. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-24T08:18:53Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The required preparation period between the task start time and the switchover time. The time is displayed in the HH:mm:ss format.
	//
	// example:
	//
	// 14:00:00
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// The region of the instance.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The result information.
	//
	// example:
	//
	// ""
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The time when the task was preformed. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-20T02:40:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the task. Valid values:
	//
	// - 2: The task is waiting for users to specify a switchover time.
	//
	// - 3: The task is waiting to be performed.
	//
	// - 4: The task is being performed. If the task is in this state, the ModifyActiveOperationTask operation cannot be called to modify the scheduled switchover time.
	//
	// - 5: The task is performed.
	//
	// - 6: The task fails.
	//
	// - 7: The task is canceled.
	//
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the system performs the switchover operation. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2024-11-06T14:11:08Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The task parameters.
	//
	// example:
	//
	// ""
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// The type of the task. Valid values:
	//
	// - rds_apsaradb_ha: master-replica switchover
	//
	// - rds_apsaradb_transfer: instance migration
	//
	// - rds_apsaradb_upgrade: minor version update
	//
	// - all: all types
	//
	// example:
	//
	// rds_apsaradb_ha
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeActiveOperationTaskResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTaskResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetDbType() *string {
	return s.DbType
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetDeadline() *string {
	return s.Deadline
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetId() *int32 {
	return s.Id
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetInsName() *string {
	return s.InsName
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetPrepareInterval() *string {
	return s.PrepareInterval
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetResultInfo() *string {
	return s.ResultInfo
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetTaskParams() *string {
	return s.TaskParams
}

func (s *DescribeActiveOperationTaskResponseBodyItems) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetCreatedTime(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetDbType(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.DbType = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetDeadline(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.Deadline = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetId(v int32) *DescribeActiveOperationTaskResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetInsName(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.InsName = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetModifiedTime(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetPrepareInterval(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.PrepareInterval = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetRegion(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetResultInfo(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.ResultInfo = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetStartTime(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetStatus(v int32) *DescribeActiveOperationTaskResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetSwitchTime(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetTaskParams(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.TaskParams = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) SetTaskType(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTaskResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

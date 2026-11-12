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
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ED2BD080-F9D9-4822-ACBB-AE5D87******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 4
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
	// The time when the task was created. The time is displayed in the `yyyy-MM-ddTHH:mmZ` format (UTC).
	//
	// example:
	//
	// 2020-06-09T22:00:42Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The database engine type.
	//
	// example:
	//
	// polardb_mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 8.0
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// The latest deadline within the adjustable range for task execution time. The time is displayed in the `yyyy-MM-ddTHH:mmZ` format (UTC).
	//
	// example:
	//
	// 2020-06-11T15:59:59Z
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 111111
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// pc-*****************
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// The time when the task was modified. The time is displayed in the `yyyy-MM-ddTHH:mmZ` format (UTC).
	//
	// example:
	//
	// 2020-06-09T22:00:42Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The preparation time before the transient disconnection. The time is displayed in the `HH:mm:ss` format.
	//
	// example:
	//
	// 04:00:00
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// The region where the task is located.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The task execution result.
	//
	// example:
	//
	// 1
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The time when the task is executed in the background. The time is displayed in the `yyyy-MM-ddTHH:mmZ` format (UTC).
	//
	// example:
	//
	// 2020-06-09T18:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The task status.
	//
	// 	- If you set `IsHistory` to **0**, the details of the pending task status are returned. Valid values:
	//
	//     	- **2**: Waiting for the user to set the task start time.
	//
	//     	- **3**: Waiting to be processed.
	//
	//     	- **4**: The task is being processed.
	//
	// 	- If you set `IsHistory` to **1**, the details of the historical task status are returned. Valid values:
	//
	//      	- **5**: The task is completed and executed successfully.
	//
	//      	- **6**: The task is completed but failed.
	//
	//      	- **7**: The task is canceled.
	//
	// example:
	//
	// 5
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the transient disconnection occurs. The time is displayed in the `yyyy-MM-ddTHH:mmZ` format (UTC).
	//
	// example:
	//
	// 2020-06-09T22:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The task parameters.
	//
	// example:
	//
	// {\\"action\\": \\"UpgradeDBVersion\\", \\"check_method\\": \\"check_cpu_usage\\", \\"params\\": {\\"RelativeDbtype\\": \\"maxscale\\", \\"DBInstanceName\\": \\"pc-*************\\", \\"User_ID\\": \\"11111\\", \\"UID\\": \\"1406926474******\\", \\"UpgradePolicy\\": \\"hot\\"}}"
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// The task type.
	//
	// example:
	//
	// rds_apsaradb_upgrade
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

func (s *DescribeActiveOperationTaskResponseBodyItems) GetDbVersion() *string {
	return s.DbVersion
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

func (s *DescribeActiveOperationTaskResponseBodyItems) SetDbVersion(v string) *DescribeActiveOperationTaskResponseBodyItems {
	s.DbVersion = &v
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

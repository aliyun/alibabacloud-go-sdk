// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePendingMaintenanceActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribePendingMaintenanceActionResponseBodyItems) *DescribePendingMaintenanceActionResponseBody
	GetItems() []*DescribePendingMaintenanceActionResponseBodyItems
	SetPageNumber(v int32) *DescribePendingMaintenanceActionResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePendingMaintenanceActionResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribePendingMaintenanceActionResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribePendingMaintenanceActionResponseBody
	GetTotalRecordCount() *int32
}

type DescribePendingMaintenanceActionResponseBody struct {
	// Details about tasks.
	Items []*DescribePendingMaintenanceActionResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// The ID of the request.
	//
	// example:
	//
	// 2F029645-FED9-4FE8-A6D3-488954******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribePendingMaintenanceActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePendingMaintenanceActionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionResponseBody) GetItems() []*DescribePendingMaintenanceActionResponseBodyItems {
	return s.Items
}

func (s *DescribePendingMaintenanceActionResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePendingMaintenanceActionResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePendingMaintenanceActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePendingMaintenanceActionResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribePendingMaintenanceActionResponseBody) SetItems(v []*DescribePendingMaintenanceActionResponseBodyItems) *DescribePendingMaintenanceActionResponseBody {
	s.Items = v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBody) SetPageNumber(v int32) *DescribePendingMaintenanceActionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBody) SetPageSize(v int32) *DescribePendingMaintenanceActionResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBody) SetRequestId(v string) *DescribePendingMaintenanceActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBody) SetTotalRecordCount(v int32) *DescribePendingMaintenanceActionResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePendingMaintenanceActionResponseBodyItems struct {
	// The time when the task was created. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-09T22:00:42Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The type of the database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **PostgreSQL**
	//
	// 	- **Oracle**
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine.
	//
	// 	- Valid values for the MySQL database engine:
	//
	//     	- **5.6**
	//
	//     	- **5.7**
	//
	//     	- **8.0**
	//
	// 	- Valid values for the PostgreSQL database engine:
	//
	//     	- **11**
	//
	//     	- **14**
	//
	// 	- Valid value for the Oracle database engine: **11**
	//
	// example:
	//
	// 8.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The deadline before which the task can be executed. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-11T15:59:59Z
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 111111
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the parameter was modified. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-09T22:00:42Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The preparation time that is required before the pending event is switched. The time follows the `HH:mm:ss` format.
	//
	// example:
	//
	// 04:00:00
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// The region ID of the pending event.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The execution result of the task. Valid values:
	//
	// 	- **manualCancel**: The task is manually canceled.
	//
	// 	- **paramCheckNotPass**: The task fails to pass the parameter check.
	//
	// > This parameter is returned only when the value of the `Status` parameter is **6*	- or **7**. The value 6 indicates that the task is completed but fails to be executed. The value 7 indicates that the task is canceled.
	//
	// example:
	//
	// manualCancel
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The time when the task was executed in the background. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-09T18:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the pending task.
	//
	// 	- If you set the `IsHistory` parameter to **0**, the status of the pending task is returned. Valid values:
	//
	//     	- **2**: The start time of the task is to be specified.
	//
	//     	- **3**: The task is pending.
	//
	//     	- **4**: The task is running. In this case, you cannot modify the execution time.
	//
	// 	- If you set the `IsHistory` parameter to **1**, the details of the historical tasks are returned. Valid values:
	//
	//     	- **5**: The task is completed and executed.
	//
	//     	- **6**: The task is completed but fails to be executed.
	//
	//     	- **7**: The task is canceled.
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the pending event was switched. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-06-09T22:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The type of the pending event.
	//
	// example:
	//
	// DatabaseSoftwareUpgrading
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribePendingMaintenanceActionResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribePendingMaintenanceActionResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetDBType() *string {
	return s.DBType
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetDeadline() *string {
	return s.Deadline
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetId() *int32 {
	return s.Id
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetPrepareInterval() *string {
	return s.PrepareInterval
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetRegion() *string {
	return s.Region
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetResultInfo() *string {
	return s.ResultInfo
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetStatus() *int32 {
	return s.Status
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetCreatedTime(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetDBClusterId(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetDBType(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.DBType = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetDBVersion(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.DBVersion = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetDeadline(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.Deadline = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetId(v int32) *DescribePendingMaintenanceActionResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetModifiedTime(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetPrepareInterval(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.PrepareInterval = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetRegion(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetResultInfo(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.ResultInfo = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetStartTime(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetStatus(v int32) *DescribePendingMaintenanceActionResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetSwitchTime(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) SetTaskType(v string) *DescribePendingMaintenanceActionResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *DescribePendingMaintenanceActionResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

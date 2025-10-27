// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMaintenanceActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeMaintenanceActionResponseBodyItems) *DescribeMaintenanceActionResponseBody
	GetItems() []*DescribeMaintenanceActionResponseBodyItems
	SetPageNumber(v int32) *DescribeMaintenanceActionResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMaintenanceActionResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeMaintenanceActionResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeMaintenanceActionResponseBody
	GetTotalRecordCount() *int32
}

type DescribeMaintenanceActionResponseBody struct {
	// The queried O\\&M events.
	Items []*DescribeMaintenanceActionResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// E774C8A9-8819-4A09-9E91-07C078******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeMaintenanceActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaintenanceActionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionResponseBody) GetItems() []*DescribeMaintenanceActionResponseBodyItems {
	return s.Items
}

func (s *DescribeMaintenanceActionResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMaintenanceActionResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMaintenanceActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMaintenanceActionResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeMaintenanceActionResponseBody) SetItems(v []*DescribeMaintenanceActionResponseBodyItems) *DescribeMaintenanceActionResponseBody {
	s.Items = v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetPageNumber(v int32) *DescribeMaintenanceActionResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetPageSize(v int32) *DescribeMaintenanceActionResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetRequestId(v string) *DescribeMaintenanceActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) SetTotalRecordCount(v int32) *DescribeMaintenanceActionResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBody) Validate() error {
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

type DescribeMaintenanceActionResponseBodyItems struct {
	// The time when the O\\&M event was created. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-30T02:44:27Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The ID of the cluster that is involved in the O\\&M event.
	//
	// example:
	//
	// am-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database engine.
	//
	// example:
	//
	// analyticdb
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 3.0
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The deadline before which the event can be executed. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-07-04T15:59:59Z
	Deadline *string `json:"Deadline,omitempty" xml:"Deadline,omitempty"`
	// The ID of the O\\&M event.
	//
	// example:
	//
	// 11111
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The point in time at which the switchover time of the O\\&M event was modified. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-07-03T06:33:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The preparation time that is required before the pending O\\&M event can be switched. The time is in the `HH:mm:ss` format.
	//
	// example:
	//
	// 02:00:00
	PrepareInterval *string `json:"PrepareInterval,omitempty" xml:"PrepareInterval,omitempty"`
	// The ID of the region where the O\\&M event occurs.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The execution result of the O\\&M event.
	//
	// > This parameter is returned only when the value of `Status` is **FAILED*	- or **CANCEL**.
	//
	// example:
	//
	// autoCancel
	ResultInfo *string `json:"ResultInfo,omitempty" xml:"ResultInfo,omitempty"`
	// The time when the task was executed in the backend. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-07-03T04:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The state of the event.
	//
	// 	- If you set `IsHistory` to **0**, the state of the pending O\\&M event is returned. Valid values:
	//
	//     	- **WAITING_MODIFY**: The start time of the O\\&M event is waiting to be set.
	//
	//     	- **WAITING**: The O\\&M event is waiting to be processed.
	//
	//     	- **PROCESSING**: The O\\&M event is being processed. The switchover time of an event in this state cannot be changed.
	//
	// 	- If you set `IsHistory` to **1**, the state of the historical O\\&M event is returned. Valid values:
	//
	//     	- **SUCCESS**: The event ended and the execution succeeded.
	//
	//     	- **FAILED**: The event ended but the execution failed.
	//
	//     	- **CANCEL**: The event was canceled.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the pending O\\&M event is switched. The time follows the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-07-03T06:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The type of the O\\&M event.
	//
	// example:
	//
	// rds_apsaradb_upgrade
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeMaintenanceActionResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeMaintenanceActionResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetDBType() *string {
	return s.DBType
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetDeadline() *string {
	return s.Deadline
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetId() *int32 {
	return s.Id
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetPrepareInterval() *string {
	return s.PrepareInterval
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetResultInfo() *string {
	return s.ResultInfo
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *DescribeMaintenanceActionResponseBodyItems) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetCreatedTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDBClusterId(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDBType(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.DBType = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDBVersion(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.DBVersion = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetDeadline(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.Deadline = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetId(v int32) *DescribeMaintenanceActionResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetModifiedTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetPrepareInterval(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.PrepareInterval = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetRegion(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.Region = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetResultInfo(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.ResultInfo = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetStartTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetStatus(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetSwitchTime(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) SetTaskType(v string) *DescribeMaintenanceActionResponseBodyItems {
	s.TaskType = &v
	return s
}

func (s *DescribeMaintenanceActionResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

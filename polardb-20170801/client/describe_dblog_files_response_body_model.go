// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBLogFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBLogFilesResponseBody
	GetDBInstanceName() *string
	SetDBInstanceType(v string) *DescribeDBLogFilesResponseBody
	GetDBInstanceType() *string
	SetHaLogItems(v []*DescribeDBLogFilesResponseBodyHaLogItems) *DescribeDBLogFilesResponseBody
	GetHaLogItems() []*DescribeDBLogFilesResponseBodyHaLogItems
	SetHaStatus(v int32) *DescribeDBLogFilesResponseBody
	GetHaStatus() *int32
	SetItemsNumbers(v int32) *DescribeDBLogFilesResponseBody
	GetItemsNumbers() *int32
	SetPageNumber(v int32) *DescribeDBLogFilesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBLogFilesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDBLogFilesResponseBody
	GetRequestId() *string
	SetSwitchListItems(v []*DescribeDBLogFilesResponseBodySwitchListItems) *DescribeDBLogFilesResponseBody
	GetSwitchListItems() []*DescribeDBLogFilesResponseBodySwitchListItems
	SetSwitchLogItems(v []*DescribeDBLogFilesResponseBodySwitchLogItems) *DescribeDBLogFilesResponseBody
	GetSwitchLogItems() []*DescribeDBLogFilesResponseBodySwitchLogItems
	SetTotalRecords(v int32) *DescribeDBLogFilesResponseBody
	GetTotalRecords() *int32
}

type DescribeDBLogFilesResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-****************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The instance type. Valid values:
	//
	// - **polardb_mysql_rw**: read-write instance.
	//
	// - **polardb_mysql_ro**: read-only instance.
	//
	// - **polardb_mysql_standby**: standby instance.
	//
	// example:
	//
	// polardb_mysql_rw
	DBInstanceType *string `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	// A list of failover logs.
	HaLogItems []*DescribeDBLogFilesResponseBodyHaLogItems `json:"HaLogItems,omitempty" xml:"HaLogItems,omitempty" type:"Repeated"`
	// Indicates whether a failover record exists. Valid values:
	//
	// - **1**: No
	//
	// - **0**: Yes
	//
	// example:
	//
	// 1
	HaStatus *int32 `json:"HaStatus,omitempty" xml:"HaStatus,omitempty"`
	// The number of log items on the current page.
	//
	// example:
	//
	// 1
	ItemsNumbers *int32 `json:"ItemsNumbers,omitempty" xml:"ItemsNumbers,omitempty"`
	// The page number. It must be a positive integer that does not exceed the maximum value of the Integer data type. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 5 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of fault simulation records.
	SwitchListItems []*DescribeDBLogFilesResponseBodySwitchListItems `json:"SwitchListItems,omitempty" xml:"SwitchListItems,omitempty" type:"Repeated"`
	// A list of fault simulation logs.
	SwitchLogItems []*DescribeDBLogFilesResponseBodySwitchLogItems `json:"SwitchLogItems,omitempty" xml:"SwitchLogItems,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalRecords *int32 `json:"TotalRecords,omitempty" xml:"TotalRecords,omitempty"`
}

func (s DescribeDBLogFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLogFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBLogFilesResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBLogFilesResponseBody) GetDBInstanceType() *string {
	return s.DBInstanceType
}

func (s *DescribeDBLogFilesResponseBody) GetHaLogItems() []*DescribeDBLogFilesResponseBodyHaLogItems {
	return s.HaLogItems
}

func (s *DescribeDBLogFilesResponseBody) GetHaStatus() *int32 {
	return s.HaStatus
}

func (s *DescribeDBLogFilesResponseBody) GetItemsNumbers() *int32 {
	return s.ItemsNumbers
}

func (s *DescribeDBLogFilesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBLogFilesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBLogFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBLogFilesResponseBody) GetSwitchListItems() []*DescribeDBLogFilesResponseBodySwitchListItems {
	return s.SwitchListItems
}

func (s *DescribeDBLogFilesResponseBody) GetSwitchLogItems() []*DescribeDBLogFilesResponseBodySwitchLogItems {
	return s.SwitchLogItems
}

func (s *DescribeDBLogFilesResponseBody) GetTotalRecords() *int32 {
	return s.TotalRecords
}

func (s *DescribeDBLogFilesResponseBody) SetDBInstanceName(v string) *DescribeDBLogFilesResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBLogFilesResponseBody) SetDBInstanceType(v string) *DescribeDBLogFilesResponseBody {
	s.DBInstanceType = &v
	return s
}

func (s *DescribeDBLogFilesResponseBody) SetHaLogItems(v []*DescribeDBLogFilesResponseBodyHaLogItems) *DescribeDBLogFilesResponseBody {
	s.HaLogItems = v
	return s
}

func (s *DescribeDBLogFilesResponseBody) SetHaStatus(v int32) *DescribeDBLogFilesResponseBody {
	s.HaStatus = &v
	return s
}

func (s *DescribeDBLogFilesResponseBody) SetItemsNumbers(v int32) *DescribeDBLogFilesResponseBody {
	s.ItemsNumbers = &v
	return s
}

func (s *DescribeDBLogFilesResponseBody) SetPageNumber(v int32) *DescribeDBLogFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBLogFilesResponseBody) SetPageSize(v int32) *DescribeDBLogFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDBLogFilesResponseBody) SetRequestId(v string) *DescribeDBLogFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBLogFilesResponseBody) SetSwitchListItems(v []*DescribeDBLogFilesResponseBodySwitchListItems) *DescribeDBLogFilesResponseBody {
	s.SwitchListItems = v
	return s
}

func (s *DescribeDBLogFilesResponseBody) SetSwitchLogItems(v []*DescribeDBLogFilesResponseBodySwitchLogItems) *DescribeDBLogFilesResponseBody {
	s.SwitchLogItems = v
	return s
}

func (s *DescribeDBLogFilesResponseBody) SetTotalRecords(v int32) *DescribeDBLogFilesResponseBody {
	s.TotalRecords = &v
	return s
}

func (s *DescribeDBLogFilesResponseBody) Validate() error {
	if s.HaLogItems != nil {
		for _, item := range s.HaLogItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SwitchListItems != nil {
		for _, item := range s.SwitchListItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SwitchLogItems != nil {
		for _, item := range s.SwitchLogItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBLogFilesResponseBodyHaLogItems struct {
	// The number of affected sessions during the failover.
	//
	// example:
	//
	// 100
	AffectedSessions *int64 `json:"AffectedSessions,omitempty" xml:"AffectedSessions,omitempty"`
	// The instance type before the failover. Valid values:
	//
	// - **polardb_mysql_rw**: read-write instance.
	//
	// - **polardb_mysql_ro**: read-only instance.
	//
	// - **polardb_mysql_standby**: standby instance.
	//
	// example:
	//
	// polardb_mysql_rw
	FromDBType *string `json:"FromDBType,omitempty" xml:"FromDBType,omitempty"`
	// The error code for the failover cause.
	//
	// example:
	//
	// Platform.Ha.AuroraService.ManualOperations
	SwitchCauseCode *string `json:"SwitchCauseCode,omitempty" xml:"SwitchCauseCode,omitempty"`
	// Details about the failover cause.
	//
	// example:
	//
	// Platform.Ha.ManuallyTriggered
	SwitchCauseDetail *string `json:"SwitchCauseDetail,omitempty" xml:"SwitchCauseDetail,omitempty"`
	// The time when the failover was complete. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-05-20T03:09:56Z
	SwitchFinishTime *string `json:"SwitchFinishTime,omitempty" xml:"SwitchFinishTime,omitempty"`
	// The failover log ID.
	//
	// example:
	//
	// e571f897-9b3c-4012-9470-88333832dec4
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The time when the failover started. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-05-20T03:09:45Z
	SwitchStartTime *string `json:"SwitchStartTime,omitempty" xml:"SwitchStartTime,omitempty"`
	// The failover type.
	//
	// example:
	//
	// 0
	SwitchType *int64 `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
	// The total number of sessions during the failover.
	//
	// example:
	//
	// 10000
	TotalSessions *int64 `json:"TotalSessions,omitempty" xml:"TotalSessions,omitempty"`
}

func (s DescribeDBLogFilesResponseBodyHaLogItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLogFilesResponseBodyHaLogItems) GoString() string {
	return s.String()
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) GetAffectedSessions() *int64 {
	return s.AffectedSessions
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) GetFromDBType() *string {
	return s.FromDBType
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) GetSwitchCauseCode() *string {
	return s.SwitchCauseCode
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) GetSwitchCauseDetail() *string {
	return s.SwitchCauseDetail
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) GetSwitchFinishTime() *string {
	return s.SwitchFinishTime
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) GetSwitchId() *string {
	return s.SwitchId
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) GetSwitchStartTime() *string {
	return s.SwitchStartTime
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) GetSwitchType() *int64 {
	return s.SwitchType
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) GetTotalSessions() *int64 {
	return s.TotalSessions
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) SetAffectedSessions(v int64) *DescribeDBLogFilesResponseBodyHaLogItems {
	s.AffectedSessions = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) SetFromDBType(v string) *DescribeDBLogFilesResponseBodyHaLogItems {
	s.FromDBType = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) SetSwitchCauseCode(v string) *DescribeDBLogFilesResponseBodyHaLogItems {
	s.SwitchCauseCode = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) SetSwitchCauseDetail(v string) *DescribeDBLogFilesResponseBodyHaLogItems {
	s.SwitchCauseDetail = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) SetSwitchFinishTime(v string) *DescribeDBLogFilesResponseBodyHaLogItems {
	s.SwitchFinishTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) SetSwitchId(v string) *DescribeDBLogFilesResponseBodyHaLogItems {
	s.SwitchId = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) SetSwitchStartTime(v string) *DescribeDBLogFilesResponseBodyHaLogItems {
	s.SwitchStartTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) SetSwitchType(v int64) *DescribeDBLogFilesResponseBodyHaLogItems {
	s.SwitchType = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) SetTotalSessions(v int64) *DescribeDBLogFilesResponseBodyHaLogItems {
	s.TotalSessions = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodyHaLogItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDBLogFilesResponseBodySwitchListItems struct {
	// The IDs of nodes on which to simulate a fault.
	//
	// > For a node-level fault simulation, specify the ID of a single node. For an availability zone-level fault simulation, you can either omit this parameter or specify the IDs of all nodes in the zone.
	DBNodeCrashList []*string `json:"DBNodeCrashList,omitempty" xml:"DBNodeCrashList,omitempty" type:"Repeated"`
	// The time when the fault simulation was complete. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-02-10T02:25:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time when the system event was complete. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-04-19T02:12:00Z
	EventFinishTime *string `json:"EventFinishTime,omitempty" xml:"EventFinishTime,omitempty"`
	// The time when the system event started. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-04-19T01:12:00Z
	EventStartTime *string `json:"EventStartTime,omitempty" xml:"EventStartTime,omitempty"`
	// The fault injection method. Valid values:
	//
	// - CrashSQLInjection: Injects a fault into the instance by using `Crash SQL`.
	//
	// example:
	//
	// CrashSQLInjection
	FaultInjectionType *string `json:"FaultInjectionType,omitempty" xml:"FaultInjectionType,omitempty"`
	// The fault simulation record ID.
	//
	// example:
	//
	// 23
	SimulateListId *string `json:"SimulateListId,omitempty" xml:"SimulateListId,omitempty"`
	// The fault simulation mode.
	//
	// example:
	//
	// 0
	SimulateMode *string `json:"SimulateMode,omitempty" xml:"SimulateMode,omitempty"`
	// The fault simulation status. Valid values:
	//
	// - **0**: Pending
	//
	// - **1**: Success
	//
	// - **2**: Running
	//
	// - **3**: Failed
	//
	// - **4**: Aborted
	//
	// - **5**: Awaiting rollback
	//
	// example:
	//
	// 2
	SimulateStatus *string `json:"SimulateStatus,omitempty" xml:"SimulateStatus,omitempty"`
	// The fault simulation task ID.
	//
	// example:
	//
	// 23
	SimulateTaskId *string `json:"SimulateTaskId,omitempty" xml:"SimulateTaskId,omitempty"`
	// The time when the fault simulation started. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-02-25T01:05:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A list of fault simulation logs.
	SwitchLogItems []*DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems `json:"SwitchLogItems,omitempty" xml:"SwitchLogItems,omitempty" type:"Repeated"`
	// A list of failover steps.
	SwitchStepItems []*DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems `json:"SwitchStepItems,omitempty" xml:"SwitchStepItems,omitempty" type:"Repeated"`
}

func (s DescribeDBLogFilesResponseBodySwitchListItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLogFilesResponseBodySwitchListItems) GoString() string {
	return s.String()
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetDBNodeCrashList() []*string {
	return s.DBNodeCrashList
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetEventFinishTime() *string {
	return s.EventFinishTime
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetEventStartTime() *string {
	return s.EventStartTime
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetFaultInjectionType() *string {
	return s.FaultInjectionType
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetSimulateListId() *string {
	return s.SimulateListId
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetSimulateMode() *string {
	return s.SimulateMode
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetSimulateStatus() *string {
	return s.SimulateStatus
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetSimulateTaskId() *string {
	return s.SimulateTaskId
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetSwitchLogItems() []*DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems {
	return s.SwitchLogItems
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) GetSwitchStepItems() []*DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems {
	return s.SwitchStepItems
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetDBNodeCrashList(v []*string) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.DBNodeCrashList = v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetEndTime(v string) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.EndTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetEventFinishTime(v string) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.EventFinishTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetEventStartTime(v string) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.EventStartTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetFaultInjectionType(v string) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.FaultInjectionType = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetSimulateListId(v string) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.SimulateListId = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetSimulateMode(v string) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.SimulateMode = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetSimulateStatus(v string) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.SimulateStatus = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetSimulateTaskId(v string) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.SimulateTaskId = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetStartTime(v string) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.StartTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetSwitchLogItems(v []*DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.SwitchLogItems = v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) SetSwitchStepItems(v []*DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) *DescribeDBLogFilesResponseBodySwitchListItems {
	s.SwitchStepItems = v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItems) Validate() error {
	if s.SwitchLogItems != nil {
		for _, item := range s.SwitchLogItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SwitchStepItems != nil {
		for _, item := range s.SwitchStepItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of your clusters, including the cluster IDs.
	//
	// example:
	//
	// pc-*************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The destination database type. Valid values:
	//
	// - **PolarDBMySQL**: A major version upgrade of PolarDB for MySQL.
	//
	// - **RDS**: A migration from RDS to PolarDB for MySQL.
	//
	// example:
	//
	// PolarDBMySQL
	DstDbType *string `json:"DstDbType,omitempty" xml:"DstDbType,omitempty"`
	// The time when the system event was complete. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-04-19T02:12:00Z
	EventFinishTime *string `json:"EventFinishTime,omitempty" xml:"EventFinishTime,omitempty"`
	// The time when the system event started. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-04-19T01:12:00Z
	EventStartTime *string `json:"EventStartTime,omitempty" xml:"EventStartTime,omitempty"`
	// The simulation list ID.
	//
	// example:
	//
	// 96
	SimulateListId *string `json:"SimulateListId,omitempty" xml:"SimulateListId,omitempty"`
	// The simulation log ID.
	//
	// example:
	//
	// 23
	SimulateLogId *string `json:"SimulateLogId,omitempty" xml:"SimulateLogId,omitempty"`
	// The fault simulation status. Valid values:
	//
	// - **0**: Pending
	//
	// - **1**: Success
	//
	// - **2**: Running
	//
	// - **3**: Failed
	//
	// - **4**: Aborted
	//
	// - **5**: Awaiting rollback
	//
	// example:
	//
	// 1
	SimulateStatus *string `json:"SimulateStatus,omitempty" xml:"SimulateStatus,omitempty"`
	// The source database type. Valid values:
	//
	// - **PolarDBMySQL**: A major version upgrade of PolarDB for MySQL.
	//
	// - **RDS**: A migration from RDS to PolarDB for MySQL.
	//
	// example:
	//
	// PolarDBMySQL
	SrcDbType *string `json:"SrcDbType,omitempty" xml:"SrcDbType,omitempty"`
	// A list of fault simulation steps.
	SwitchStepItems []*DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems `json:"SwitchStepItems,omitempty" xml:"SwitchStepItems,omitempty" type:"Repeated"`
}

func (s DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) GoString() string {
	return s.String()
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) GetDstDbType() *string {
	return s.DstDbType
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) GetEventFinishTime() *string {
	return s.EventFinishTime
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) GetEventStartTime() *string {
	return s.EventStartTime
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) GetSimulateListId() *string {
	return s.SimulateListId
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) GetSimulateLogId() *string {
	return s.SimulateLogId
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) GetSimulateStatus() *string {
	return s.SimulateStatus
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) GetSrcDbType() *string {
	return s.SrcDbType
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) GetSwitchStepItems() []*DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems {
	return s.SwitchStepItems
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) SetDBInstanceId(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) SetDstDbType(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems {
	s.DstDbType = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) SetEventFinishTime(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems {
	s.EventFinishTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) SetEventStartTime(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems {
	s.EventStartTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) SetSimulateListId(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems {
	s.SimulateListId = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) SetSimulateLogId(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems {
	s.SimulateLogId = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) SetSimulateStatus(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems {
	s.SimulateStatus = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) SetSrcDbType(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems {
	s.SrcDbType = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) SetSwitchStepItems(v []*DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems {
	s.SwitchStepItems = v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems) Validate() error {
	if s.SwitchStepItems != nil {
		for _, item := range s.SwitchStepItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems struct {
	// The time when the step was complete. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-08-14T02:07:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the step was successful. Valid values:
	//
	// - `true`: The step was successful.
	//
	// - `false`: The step failed.
	//
	// example:
	//
	// true
	IsSuccess *string `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The fault simulation phase. Valid values:
	//
	// - **PolarDB.MySQL.FaultSimulate.Phase.FAULT_INJECTION**: The fault injection phase.
	//
	// - **PolarDB.MySQL.FaultSimulate.Phase.RECOVERY**: The recovery phase.
	//
	// - **PolarDB.MySQL.FaultSimulate.Phase.POST_PROCESS**: The post-processing phase.
	//
	// example:
	//
	// PolarDB.MySQL.FaultSimulate.Phase.FAULT_INJECTION
	SimulatePhase *string `json:"SimulatePhase,omitempty" xml:"SimulatePhase,omitempty"`
	// The time when the step started. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-04-19T02:12:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the current step. You can call the [DescribeHistoryTasks](https://help.aliyun.com/document_detail/2400077.html) operation to query the current step of a specified task. A common value is **do_pause**, which indicates that the system waits for a specified period of time.
	//
	// example:
	//
	// init_task_info
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The duration of the step in milliseconds.
	//
	// example:
	//
	// 1000
	TimeCost *string `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) GoString() string {
	return s.String()
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) GetIsSuccess() *string {
	return s.IsSuccess
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) GetSimulatePhase() *string {
	return s.SimulatePhase
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) GetStepName() *string {
	return s.StepName
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) GetTimeCost() *string {
	return s.TimeCost
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) SetEndTime(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems {
	s.EndTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) SetIsSuccess(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems {
	s.IsSuccess = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) SetSimulatePhase(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems {
	s.SimulatePhase = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) SetStartTime(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems {
	s.StartTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) SetStepName(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems {
	s.StepName = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) SetTimeCost(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems {
	s.TimeCost = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItemsSwitchStepItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems struct {
	// The cluster node ID.
	//
	// > This parameter is returned only when the `Key` parameter in the request is not set to `PolarDBDiskUsage`.
	//
	// example:
	//
	// pi-*************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The time when the step was complete. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-07-23T02:30:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the step was successful. Valid values:
	//
	// - `true`: The step was successful.
	//
	// - `false`: The step failed.
	//
	// example:
	//
	// true
	IsSuccess *string `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The fault simulation phase. Valid values:
	//
	// - **PolarDB.MySQL.FaultSimulate.Phase.FAULT_INJECTION**: The fault injection phase.
	//
	// - **PolarDB.MySQL.FaultSimulate.Phase.RECOVERY**: The recovery phase.
	//
	// - **PolarDB.MySQL.FaultSimulate.Phase.POST_PROCESS**: The post-processing phase.
	//
	// example:
	//
	// PolarDB.MySQL.FaultSimulate.Phase.FAULT_INJECTION
	SimulatePhase *string `json:"SimulatePhase,omitempty" xml:"SimulatePhase,omitempty"`
	// The time when the step started. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-07-16T02:12:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The name of the current step. You can call the [DescribeHistoryTasks](https://help.aliyun.com/document_detail/2400077.html) operation to query the current step of a specified task. A common value is **do_pause**, which indicates that the system waits for a specified period of time.
	//
	// example:
	//
	// init_task_info
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The duration of the step in milliseconds.
	//
	// example:
	//
	// 1000
	TimeCost *string `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) GoString() string {
	return s.String()
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) GetIsSuccess() *string {
	return s.IsSuccess
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) GetSimulatePhase() *string {
	return s.SimulatePhase
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) GetStepName() *string {
	return s.StepName
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) GetTimeCost() *string {
	return s.TimeCost
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) SetDBNodeId(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) SetEndTime(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems {
	s.EndTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) SetIsSuccess(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems {
	s.IsSuccess = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) SetSimulatePhase(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems {
	s.SimulatePhase = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) SetStartTime(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems {
	s.StartTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) SetStepName(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems {
	s.StepName = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) SetTimeCost(v string) *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems {
	s.TimeCost = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems) Validate() error {
	return dara.Validate(s)
}

type DescribeDBLogFilesResponseBodySwitchLogItems struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of your clusters, including the cluster IDs.
	//
	// example:
	//
	// pc-*************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The destination database type. Valid values:
	//
	// - **PolarDBMySQL**: A major version upgrade of PolarDB for MySQL.
	//
	// - **RDS**: A migration from RDS to PolarDB for MySQL.
	//
	// example:
	//
	// PolarDBMySQL
	DstDbType *string `json:"DstDbType,omitempty" xml:"DstDbType,omitempty"`
	// The time when the system event was complete. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-04-19T02:12:00Z
	EventFinishTime *string `json:"EventFinishTime,omitempty" xml:"EventFinishTime,omitempty"`
	// The time when the system event started. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-04-19T01:12:00Z
	EventStartTime *string `json:"EventStartTime,omitempty" xml:"EventStartTime,omitempty"`
	// The simulation list ID.
	//
	// example:
	//
	// 231
	SimulateListId *string `json:"SimulateListId,omitempty" xml:"SimulateListId,omitempty"`
	// The fault simulation status. Valid values:
	//
	// - **0**: Pending
	//
	// - **1**: Success
	//
	// - **2**: Running
	//
	// - **3**: Failed
	//
	// - **4**: Aborted
	//
	// - **5**: Awaiting rollback
	//
	// example:
	//
	// 1
	SimulateStatus *string `json:"SimulateStatus,omitempty" xml:"SimulateStatus,omitempty"`
	// The status code of the fault simulation.
	//
	// example:
	//
	// 0
	Simulatecode *string `json:"Simulatecode,omitempty" xml:"Simulatecode,omitempty"`
	// The source database type. Valid values:
	//
	// - **PolarDBMySQL**: A major version upgrade of PolarDB for MySQL.
	//
	// - **RDS**: A migration from RDS to PolarDB for MySQL.
	//
	// example:
	//
	// PolarDBMySQL
	SrcDbType *string `json:"SrcDbType,omitempty" xml:"SrcDbType,omitempty"`
	// A list of failover steps.
	SwitchStepItems []*DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems `json:"SwitchStepItems,omitempty" xml:"SwitchStepItems,omitempty" type:"Repeated"`
}

func (s DescribeDBLogFilesResponseBodySwitchLogItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLogFilesResponseBodySwitchLogItems) GoString() string {
	return s.String()
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) GetDstDbType() *string {
	return s.DstDbType
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) GetEventFinishTime() *string {
	return s.EventFinishTime
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) GetEventStartTime() *string {
	return s.EventStartTime
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) GetSimulateListId() *string {
	return s.SimulateListId
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) GetSimulateStatus() *string {
	return s.SimulateStatus
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) GetSimulatecode() *string {
	return s.Simulatecode
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) GetSrcDbType() *string {
	return s.SrcDbType
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) GetSwitchStepItems() []*DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems {
	return s.SwitchStepItems
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) SetDBInstanceId(v string) *DescribeDBLogFilesResponseBodySwitchLogItems {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) SetDstDbType(v string) *DescribeDBLogFilesResponseBodySwitchLogItems {
	s.DstDbType = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) SetEventFinishTime(v string) *DescribeDBLogFilesResponseBodySwitchLogItems {
	s.EventFinishTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) SetEventStartTime(v string) *DescribeDBLogFilesResponseBodySwitchLogItems {
	s.EventStartTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) SetSimulateListId(v string) *DescribeDBLogFilesResponseBodySwitchLogItems {
	s.SimulateListId = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) SetSimulateStatus(v string) *DescribeDBLogFilesResponseBodySwitchLogItems {
	s.SimulateStatus = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) SetSimulatecode(v string) *DescribeDBLogFilesResponseBodySwitchLogItems {
	s.Simulatecode = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) SetSrcDbType(v string) *DescribeDBLogFilesResponseBodySwitchLogItems {
	s.SrcDbType = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) SetSwitchStepItems(v []*DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) *DescribeDBLogFilesResponseBodySwitchLogItems {
	s.SwitchStepItems = v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItems) Validate() error {
	if s.SwitchStepItems != nil {
		for _, item := range s.SwitchStepItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems struct {
	// The node ID.
	//
	// > You must specify either the `DBNodeId` or `DBClusterId` parameter. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of your clusters, including the node IDs.
	//
	// example:
	//
	// pi-*************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The time when the step was complete. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2025-03-27T02:27:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the step was successful. Valid values:
	//
	// - `true`: The step was successful.
	//
	// - `false`: The step failed.
	//
	// example:
	//
	// true
	IsSuccess *string `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The fault simulation phase. Valid values:
	//
	// - **PolarDB.MySQL.FaultSimulate.Phase.FAULT_INJECTION**: The fault injection phase.
	//
	// - **PolarDB.MySQL.FaultSimulate.Phase.RECOVERY**: The recovery phase.
	//
	// - **PolarDB.MySQL.FaultSimulate.Phase.POST_PROCESS**: The post-processing phase.
	//
	// example:
	//
	// PolarDB.MySQL.FaultSimulate.Phase.FAULT_INJECTION
	SimulatePhase *string `json:"SimulatePhase,omitempty" xml:"SimulatePhase,omitempty"`
	// The time when the step started. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is displayed in UTC.
	//
	// example:
	//
	// 2024-10-21T02:12:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A message about the execution status of the step.
	//
	// example:
	//
	// This step exec success
	StepMsg *string `json:"StepMsg,omitempty" xml:"StepMsg,omitempty"`
	// The name of the step.
	//
	// example:
	//
	// init_task_info
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// The duration of the step in milliseconds.
	//
	// example:
	//
	// 1000
	TimeCost *string `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) GoString() string {
	return s.String()
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) GetIsSuccess() *string {
	return s.IsSuccess
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) GetSimulatePhase() *string {
	return s.SimulatePhase
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) GetStepMsg() *string {
	return s.StepMsg
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) GetStepName() *string {
	return s.StepName
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) GetTimeCost() *string {
	return s.TimeCost
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) SetDBNodeId(v string) *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) SetEndTime(v string) *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems {
	s.EndTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) SetIsSuccess(v string) *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems {
	s.IsSuccess = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) SetSimulatePhase(v string) *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems {
	s.SimulatePhase = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) SetStartTime(v string) *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems {
	s.StartTime = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) SetStepMsg(v string) *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems {
	s.StepMsg = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) SetStepName(v string) *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems {
	s.StepName = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) SetTimeCost(v string) *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems {
	s.TimeCost = &v
	return s
}

func (s *DescribeDBLogFilesResponseBodySwitchLogItemsSwitchStepItems) Validate() error {
	return dara.Validate(s)
}

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
	// example:
	//
	// pi-****************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// polardb_mysql_rw
	DBInstanceType *string                                     `json:"DBInstanceType,omitempty" xml:"DBInstanceType,omitempty"`
	HaLogItems     []*DescribeDBLogFilesResponseBodyHaLogItems `json:"HaLogItems,omitempty" xml:"HaLogItems,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	HaStatus *int32 `json:"HaStatus,omitempty" xml:"HaStatus,omitempty"`
	// example:
	//
	// 10
	ItemsNumbers *int32 `json:"ItemsNumbers,omitempty" xml:"ItemsNumbers,omitempty"`
	// example:
	//
	// 6
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId       *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SwitchListItems []*DescribeDBLogFilesResponseBodySwitchListItems `json:"SwitchListItems,omitempty" xml:"SwitchListItems,omitempty" type:"Repeated"`
	SwitchLogItems  []*DescribeDBLogFilesResponseBodySwitchLogItems  `json:"SwitchLogItems,omitempty" xml:"SwitchLogItems,omitempty" type:"Repeated"`
	// example:
	//
	// 160
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
	AffectedSessions  *int64  `json:"AffectedSessions,omitempty" xml:"AffectedSessions,omitempty"`
	FromDBType        *string `json:"FromDBType,omitempty" xml:"FromDBType,omitempty"`
	SwitchCauseCode   *string `json:"SwitchCauseCode,omitempty" xml:"SwitchCauseCode,omitempty"`
	SwitchCauseDetail *string `json:"SwitchCauseDetail,omitempty" xml:"SwitchCauseDetail,omitempty"`
	SwitchFinishTime  *string `json:"SwitchFinishTime,omitempty" xml:"SwitchFinishTime,omitempty"`
	// example:
	//
	// e571f897-9b3c-4012-9470-88333832dec4
	SwitchId        *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	SwitchStartTime *string `json:"SwitchStartTime,omitempty" xml:"SwitchStartTime,omitempty"`
	// example:
	//
	// 0
	SwitchType    *int64 `json:"SwitchType,omitempty" xml:"SwitchType,omitempty"`
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
	DBNodeCrashList    []*string                                                       `json:"DBNodeCrashList,omitempty" xml:"DBNodeCrashList,omitempty" type:"Repeated"`
	EndTime            *string                                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EventFinishTime    *string                                                         `json:"EventFinishTime,omitempty" xml:"EventFinishTime,omitempty"`
	EventStartTime     *string                                                         `json:"EventStartTime,omitempty" xml:"EventStartTime,omitempty"`
	FaultInjectionType *string                                                         `json:"FaultInjectionType,omitempty" xml:"FaultInjectionType,omitempty"`
	SimulateListId     *string                                                         `json:"SimulateListId,omitempty" xml:"SimulateListId,omitempty"`
	SimulateMode       *string                                                         `json:"SimulateMode,omitempty" xml:"SimulateMode,omitempty"`
	SimulateStatus     *string                                                         `json:"SimulateStatus,omitempty" xml:"SimulateStatus,omitempty"`
	SimulateTaskId     *string                                                         `json:"SimulateTaskId,omitempty" xml:"SimulateTaskId,omitempty"`
	StartTime          *string                                                         `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SwitchLogItems     []*DescribeDBLogFilesResponseBodySwitchListItemsSwitchLogItems  `json:"SwitchLogItems,omitempty" xml:"SwitchLogItems,omitempty" type:"Repeated"`
	SwitchStepItems    []*DescribeDBLogFilesResponseBodySwitchListItemsSwitchStepItems `json:"SwitchStepItems,omitempty" xml:"SwitchStepItems,omitempty" type:"Repeated"`
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
	DBInstanceId    *string                                                                       `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DstDbType       *string                                                                       `json:"DstDbType,omitempty" xml:"DstDbType,omitempty"`
	EventFinishTime *string                                                                       `json:"EventFinishTime,omitempty" xml:"EventFinishTime,omitempty"`
	EventStartTime  *string                                                                       `json:"EventStartTime,omitempty" xml:"EventStartTime,omitempty"`
	SimulateListId  *string                                                                       `json:"SimulateListId,omitempty" xml:"SimulateListId,omitempty"`
	SimulateLogId   *string                                                                       `json:"SimulateLogId,omitempty" xml:"SimulateLogId,omitempty"`
	SimulateStatus  *string                                                                       `json:"SimulateStatus,omitempty" xml:"SimulateStatus,omitempty"`
	SrcDbType       *string                                                                       `json:"SrcDbType,omitempty" xml:"SrcDbType,omitempty"`
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
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsSuccess     *string `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	SimulatePhase *string `json:"SimulatePhase,omitempty" xml:"SimulatePhase,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StepName      *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	TimeCost      *string `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
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
	DBNodeId      *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsSuccess     *string `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	SimulatePhase *string `json:"SimulatePhase,omitempty" xml:"SimulatePhase,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StepName      *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	TimeCost      *string `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
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
	DBInstanceId    *string                                                        `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DstDbType       *string                                                        `json:"DstDbType,omitempty" xml:"DstDbType,omitempty"`
	EventFinishTime *string                                                        `json:"EventFinishTime,omitempty" xml:"EventFinishTime,omitempty"`
	EventStartTime  *string                                                        `json:"EventStartTime,omitempty" xml:"EventStartTime,omitempty"`
	SimulateListId  *string                                                        `json:"SimulateListId,omitempty" xml:"SimulateListId,omitempty"`
	SimulateStatus  *string                                                        `json:"SimulateStatus,omitempty" xml:"SimulateStatus,omitempty"`
	Simulatecode    *string                                                        `json:"Simulatecode,omitempty" xml:"Simulatecode,omitempty"`
	SrcDbType       *string                                                        `json:"SrcDbType,omitempty" xml:"SrcDbType,omitempty"`
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
	DBNodeId      *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	IsSuccess     *string `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	SimulatePhase *string `json:"SimulatePhase,omitempty" xml:"SimulatePhase,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StepMsg       *string `json:"StepMsg,omitempty" xml:"StepMsg,omitempty"`
	StepName      *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	TimeCost      *string `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
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

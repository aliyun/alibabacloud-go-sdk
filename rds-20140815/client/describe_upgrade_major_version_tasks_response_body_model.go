// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeMajorVersionTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeUpgradeMajorVersionTasksResponseBodyItems) *DescribeUpgradeMajorVersionTasksResponseBody
	GetItems() []*DescribeUpgradeMajorVersionTasksResponseBodyItems
	SetPageNumber(v int32) *DescribeUpgradeMajorVersionTasksResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeUpgradeMajorVersionTasksResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeUpgradeMajorVersionTasksResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeUpgradeMajorVersionTasksResponseBody
	GetTotalRecordCount() *int32
}

type DescribeUpgradeMajorVersionTasksResponseBody struct {
	// The tasks for major engine version upgrades.
	Items []*DescribeUpgradeMajorVersionTasksResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 152E0C6D-B9C3-4468-9F2C-FEF9D9E8417B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeUpgradeMajorVersionTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeMajorVersionTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeMajorVersionTasksResponseBody) GetItems() []*DescribeUpgradeMajorVersionTasksResponseBodyItems {
	return s.Items
}

func (s *DescribeUpgradeMajorVersionTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUpgradeMajorVersionTasksResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeUpgradeMajorVersionTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUpgradeMajorVersionTasksResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeUpgradeMajorVersionTasksResponseBody) SetItems(v []*DescribeUpgradeMajorVersionTasksResponseBodyItems) *DescribeUpgradeMajorVersionTasksResponseBody {
	s.Items = v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBody) SetPageNumber(v int32) *DescribeUpgradeMajorVersionTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBody) SetPageRecordCount(v int32) *DescribeUpgradeMajorVersionTasksResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBody) SetRequestId(v string) *DescribeUpgradeMajorVersionTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBody) SetTotalRecordCount(v int32) *DescribeUpgradeMajorVersionTasksResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBody) Validate() error {
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

type DescribeUpgradeMajorVersionTasksResponseBodyItems struct {
	// The time when the system collects the statistics.
	//
	// Valid values:
	//
	// 	- **After**: The system collects the statistics after a switchover.
	//
	// 	- **Before**: The system collects the statistics before a switchover.
	//
	// example:
	//
	// After
	CollectStatMode *string `json:"CollectStatMode,omitempty" xml:"CollectStatMode,omitempty"`
	// The details of the task.
	//
	// example:
	//
	// 2021-10-27 15:03:05 --- do upgrade precheck on slave succcess.\\n2021-10-27 15:03:11 --- begin to upgrade major version, source instance will locked in readonly mode.\\n2021-10-27 15:03:21 --- upgrade master success.\\n2021-10-27 15:06:10 --- exchange source and target instance dns success.\\n
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The end time of the task.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1614237779000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The status of the task.
	//
	// 	- **Success**: The task is successful.
	//
	// 	- **Failed**: The task failed.
	//
	// 	- **Running**: The task is in the phase in which data is being migrated to a new instance.
	//
	// example:
	//
	// Success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The ID of the original instance.
	//
	// example:
	//
	// pgm-bp1i3kkq7321****
	SourceInsName *string `json:"SourceInsName,omitempty" xml:"SourceInsName,omitempty"`
	// The major engine version of the original instance.
	//
	// example:
	//
	// 11.0
	SourceMajorVersion *string `json:"SourceMajorVersion,omitempty" xml:"SourceMajorVersion,omitempty"`
	// The start time of the task.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1614236007000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The end time of the switching from the original instance to the new instance.
	//
	// Expressed in Unix timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1714237539000
	SwitchEndTime *string `json:"SwitchEndTime,omitempty" xml:"SwitchEndTime,omitempty"`
	// The time at which your workloads are switched over from the original instance to the new instance.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC. Unit: milliseconds.
	//
	// example:
	//
	// 1614237539000
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The ID of the new instance.
	//
	// example:
	//
	// pgm-bp1c0v6d8092****
	TargetInsName *string `json:"TargetInsName,omitempty" xml:"TargetInsName,omitempty"`
	// The major engine version of the new instance. Valid values:
	//
	// 	- **10.0**
	//
	// 	- **11.0**
	//
	// 	- **12.0**
	//
	// 	- **13.0**
	//
	// 	- **14.0**
	//
	// 	- **15.0**
	//
	// example:
	//
	// 12.0
	TargetMajorVersion *string `json:"TargetMajorVersion,omitempty" xml:"TargetMajorVersion,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 342900000
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The upgrade mode.
	//
	// Valid values:
	//
	// 	- **clone**: The system does not migrate data to the new instance and does not switch your workloads over to the new instance.
	//
	// 	- **switch**: The system migrates data to the new instance and switches your workloads over to the new instance.
	//
	// example:
	//
	// switch
	UpgradeMode                  *string `json:"UpgradeMode,omitempty" xml:"UpgradeMode,omitempty"`
	CutOver                      *bool   `json:"cutOver,omitempty" xml:"cutOver,omitempty"`
	TotalLogicRepDelayTime       *int32  `json:"totalLogicRepDelayTime,omitempty" xml:"totalLogicRepDelayTime,omitempty"`
	TotalLogicRepLatencyMB       *int32  `json:"totalLogicRepLatencyMB,omitempty" xml:"totalLogicRepLatencyMB,omitempty"`
	ZeroDownTimeConnectionString *string `json:"zeroDownTimeConnectionString,omitempty" xml:"zeroDownTimeConnectionString,omitempty"`
	ZeroDownTimePort             *int32  `json:"zeroDownTimePort,omitempty" xml:"zeroDownTimePort,omitempty"`
}

func (s DescribeUpgradeMajorVersionTasksResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeMajorVersionTasksResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetCollectStatMode() *string {
	return s.CollectStatMode
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetDetail() *string {
	return s.Detail
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetResult() *string {
	return s.Result
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetSourceInsName() *string {
	return s.SourceInsName
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetSourceMajorVersion() *string {
	return s.SourceMajorVersion
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetSwitchEndTime() *string {
	return s.SwitchEndTime
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetTargetInsName() *string {
	return s.TargetInsName
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetTargetMajorVersion() *string {
	return s.TargetMajorVersion
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetTaskId() *int32 {
	return s.TaskId
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetUpgradeMode() *string {
	return s.UpgradeMode
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetCutOver() *bool {
	return s.CutOver
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetTotalLogicRepDelayTime() *int32 {
	return s.TotalLogicRepDelayTime
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetTotalLogicRepLatencyMB() *int32 {
	return s.TotalLogicRepLatencyMB
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetZeroDownTimeConnectionString() *string {
	return s.ZeroDownTimeConnectionString
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) GetZeroDownTimePort() *int32 {
	return s.ZeroDownTimePort
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetCollectStatMode(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.CollectStatMode = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetDetail(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.Detail = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetEndTime(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.EndTime = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetResult(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.Result = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetSourceInsName(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.SourceInsName = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetSourceMajorVersion(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.SourceMajorVersion = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetStartTime(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.StartTime = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetSwitchEndTime(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.SwitchEndTime = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetSwitchTime(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.SwitchTime = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetTargetInsName(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.TargetInsName = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetTargetMajorVersion(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.TargetMajorVersion = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetTaskId(v int32) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.TaskId = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetUpgradeMode(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.UpgradeMode = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetCutOver(v bool) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.CutOver = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetTotalLogicRepDelayTime(v int32) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.TotalLogicRepDelayTime = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetTotalLogicRepLatencyMB(v int32) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.TotalLogicRepLatencyMB = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetZeroDownTimeConnectionString(v string) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.ZeroDownTimeConnectionString = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) SetZeroDownTimePort(v int32) *DescribeUpgradeMajorVersionTasksResponseBodyItems {
	s.ZeroDownTimePort = &v
	return s
}

func (s *DescribeUpgradeMajorVersionTasksResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

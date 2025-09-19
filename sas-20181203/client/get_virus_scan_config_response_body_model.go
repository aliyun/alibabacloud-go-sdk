// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirusScanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetVirusScanConfigResponseBodyData) *GetVirusScanConfigResponseBody
	GetData() *GetVirusScanConfigResponseBodyData
	SetRequestId(v string) *GetVirusScanConfigResponseBody
	GetRequestId() *string
}

type GetVirusScanConfigResponseBody struct {
	// The data returned if the request was successful.
	Data *GetVirusScanConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 028CF634-5268-5660-9575-48C9ED6BDEFC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVirusScanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetVirusScanConfigResponseBody) GetData() *GetVirusScanConfigResponseBodyData {
	return s.Data
}

func (s *GetVirusScanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVirusScanConfigResponseBody) SetData(v *GetVirusScanConfigResponseBodyData) *GetVirusScanConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetVirusScanConfigResponseBody) SetRequestId(v string) *GetVirusScanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVirusScanConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVirusScanConfigResponseBodyData struct {
	// Extended scan types.
	AdditionType []*string `json:"AdditionType,omitempty" xml:"AdditionType,omitempty" type:"Repeated"`
	// The ID of the task configuration.
	//
	// > You can call the [DescribeCycleTaskList](~~DescribeCycleTaskList~~) operation to query the IDs of task configurations.
	//
	// example:
	//
	// 97a1fed216908e417407344e1505xxxx
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Indicates whether the periodic scan feature is enabled. Valid value:
	//
	// 	- **1**: The feature is enabled
	//
	// 	- **0**: The feature is disabled.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The interval at which virus scan tasks are run.
	//
	// example:
	//
	// 7
	IntervalPeriod *int32 `json:"IntervalPeriod,omitempty" xml:"IntervalPeriod,omitempty"`
	// The unit of the interval at which virus scan tasks are run.
	//
	// 	- The value is fixed as **day**.
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The file paths.
	ScanPath []*string `json:"ScanPath,omitempty" xml:"ScanPath,omitempty" type:"Repeated"`
	// The type of the virus scan task. Valid values:
	//
	// 	- **system**: automatic scan.
	//
	// 	- **user**: custom scan.
	//
	// example:
	//
	// user
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The key that stores the asset information.
	//
	// > You can call the [GetAssetSelectionConfig](~~GetAssetSelectionConfig~~) operation to obtain the key value.
	//
	// example:
	//
	// 345ddbea-a57f-437e-832f-fb7a1202xxxx
	SelectionKey *string `json:"SelectionKey,omitempty" xml:"SelectionKey,omitempty"`
	// The end time of the virus scan task. The time is a time frame.
	//
	// example:
	//
	// 6
	TargetEndTime *int32 `json:"TargetEndTime,omitempty" xml:"TargetEndTime,omitempty"`
	// The start time of the virus scan task. The time is a time frame.
	//
	// example:
	//
	// 0
	TargetStartTime *int32 `json:"TargetStartTime,omitempty" xml:"TargetStartTime,omitempty"`
	// The type of the task. Valid value:
	//
	// 	- **VIRUS_VUL_SCHEDULE_SCAN**: a virus scan task.
	//
	// example:
	//
	// VIRUS_VUL_SCHEDULE_SCAN
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetVirusScanConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVirusScanConfigResponseBodyData) GetAdditionType() []*string {
	return s.AdditionType
}

func (s *GetVirusScanConfigResponseBodyData) GetConfigId() *string {
	return s.ConfigId
}

func (s *GetVirusScanConfigResponseBodyData) GetEnable() *int32 {
	return s.Enable
}

func (s *GetVirusScanConfigResponseBodyData) GetIntervalPeriod() *int32 {
	return s.IntervalPeriod
}

func (s *GetVirusScanConfigResponseBodyData) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *GetVirusScanConfigResponseBodyData) GetScanPath() []*string {
	return s.ScanPath
}

func (s *GetVirusScanConfigResponseBodyData) GetScanType() *string {
	return s.ScanType
}

func (s *GetVirusScanConfigResponseBodyData) GetSelectionKey() *string {
	return s.SelectionKey
}

func (s *GetVirusScanConfigResponseBodyData) GetTargetEndTime() *int32 {
	return s.TargetEndTime
}

func (s *GetVirusScanConfigResponseBodyData) GetTargetStartTime() *int32 {
	return s.TargetStartTime
}

func (s *GetVirusScanConfigResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *GetVirusScanConfigResponseBodyData) SetAdditionType(v []*string) *GetVirusScanConfigResponseBodyData {
	s.AdditionType = v
	return s
}

func (s *GetVirusScanConfigResponseBodyData) SetConfigId(v string) *GetVirusScanConfigResponseBodyData {
	s.ConfigId = &v
	return s
}

func (s *GetVirusScanConfigResponseBodyData) SetEnable(v int32) *GetVirusScanConfigResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetVirusScanConfigResponseBodyData) SetIntervalPeriod(v int32) *GetVirusScanConfigResponseBodyData {
	s.IntervalPeriod = &v
	return s
}

func (s *GetVirusScanConfigResponseBodyData) SetPeriodUnit(v string) *GetVirusScanConfigResponseBodyData {
	s.PeriodUnit = &v
	return s
}

func (s *GetVirusScanConfigResponseBodyData) SetScanPath(v []*string) *GetVirusScanConfigResponseBodyData {
	s.ScanPath = v
	return s
}

func (s *GetVirusScanConfigResponseBodyData) SetScanType(v string) *GetVirusScanConfigResponseBodyData {
	s.ScanType = &v
	return s
}

func (s *GetVirusScanConfigResponseBodyData) SetSelectionKey(v string) *GetVirusScanConfigResponseBodyData {
	s.SelectionKey = &v
	return s
}

func (s *GetVirusScanConfigResponseBodyData) SetTargetEndTime(v int32) *GetVirusScanConfigResponseBodyData {
	s.TargetEndTime = &v
	return s
}

func (s *GetVirusScanConfigResponseBodyData) SetTargetStartTime(v int32) *GetVirusScanConfigResponseBodyData {
	s.TargetStartTime = &v
	return s
}

func (s *GetVirusScanConfigResponseBodyData) SetTaskType(v string) *GetVirusScanConfigResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetVirusScanConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}

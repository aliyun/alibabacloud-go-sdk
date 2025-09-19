// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVirusScanLatestTaskStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetVirusScanLatestTaskStatisticResponseBodyData) *GetVirusScanLatestTaskStatisticResponseBody
	GetData() *GetVirusScanLatestTaskStatisticResponseBodyData
	SetRequestId(v string) *GetVirusScanLatestTaskStatisticResponseBody
	GetRequestId() *string
}

type GetVirusScanLatestTaskStatisticResponseBody struct {
	// The information about the virus scan task.
	Data *GetVirusScanLatestTaskStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7532B7EE-7CE7-5F4D-BF04-B12447DDCAE1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVirusScanLatestTaskStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanLatestTaskStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetVirusScanLatestTaskStatisticResponseBody) GetData() *GetVirusScanLatestTaskStatisticResponseBodyData {
	return s.Data
}

func (s *GetVirusScanLatestTaskStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVirusScanLatestTaskStatisticResponseBody) SetData(v *GetVirusScanLatestTaskStatisticResponseBodyData) *GetVirusScanLatestTaskStatisticResponseBody {
	s.Data = v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBody) SetRequestId(v string) *GetVirusScanLatestTaskStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVirusScanLatestTaskStatisticResponseBodyData struct {
	// The number of machines on which the virus scan task is complete.
	//
	// example:
	//
	// 2
	CompleteMachine *int32 `json:"CompleteMachine,omitempty" xml:"CompleteMachine,omitempty"`
	// The name of the machine.
	//
	// example:
	//
	// testMahine1
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The progress of the virus scan task in percentage.
	//
	// example:
	//
	// 92
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The highest risk level of the detected alerts. Valid values:
	//
	// 	- **high**
	//
	// 	- **medium**
	//
	// 	- **low**
	//
	// example:
	//
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The number of safe machines that are detected.
	//
	// example:
	//
	// 1
	SafeMachine *int32 `json:"SafeMachine,omitempty" xml:"SafeMachine,omitempty"`
	// The number of machines that are scanned.
	//
	// example:
	//
	// 3
	ScanMachine *int32 `json:"ScanMachine,omitempty" xml:"ScanMachine,omitempty"`
	// The paths of files that were scanned. This value is returned only when ScanType is set to user.
	ScanPath []*string `json:"ScanPath,omitempty" xml:"ScanPath,omitempty" type:"Repeated"`
	// The timestamp generated when the virus scan task was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1681145862000
	ScanTime *int64 `json:"ScanTime,omitempty" xml:"ScanTime,omitempty"`
	// The type of the virus scan. Valid values:
	//
	// 	- **system**: automatic scan.
	//
	// 	- **user**: custom scan.
	//
	// example:
	//
	// system
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The status of the virus scan task.
	//
	// **Valid values for a main task**:
	//
	// 	- **0**: The main task is to be started.
	//
	// 	- **10**: The main task is running.
	//
	// 	- **100**: The main task is complete.
	//
	// **Valid values for a subtask**:
	//
	// 	- **0**: The subtask is to be started.
	//
	// 	- **20**: The scan script is sent.
	//
	// 	- **50**: The subtask is running.
	//
	// 	- **100**: The subtask is complete.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of alerts that are detected.
	//
	// example:
	//
	// 2
	SuspiciousCount *int32 `json:"SuspiciousCount,omitempty" xml:"SuspiciousCount,omitempty"`
	// The number of suspicious machines that are detected.
	//
	// example:
	//
	// 2
	SuspiciousMachine *int32 `json:"SuspiciousMachine,omitempty" xml:"SuspiciousMachine,omitempty"`
	// The ID of the virus scan task.
	//
	// example:
	//
	// fc98d58eb56f699d49bf7ebbd6d7****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The number of machines on which the virus scan task was not complete or failed.
	//
	// example:
	//
	// 1
	UnCompleteMachine *int32 `json:"UnCompleteMachine,omitempty" xml:"UnCompleteMachine,omitempty"`
}

func (s GetVirusScanLatestTaskStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVirusScanLatestTaskStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetCompleteMachine() *int32 {
	return s.CompleteMachine
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetMachineName() *string {
	return s.MachineName
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetProgress() *string {
	return s.Progress
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetSafeMachine() *int32 {
	return s.SafeMachine
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetScanMachine() *int32 {
	return s.ScanMachine
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetScanPath() []*string {
	return s.ScanPath
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetScanTime() *int64 {
	return s.ScanTime
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetScanType() *string {
	return s.ScanType
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetSuspiciousCount() *int32 {
	return s.SuspiciousCount
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetSuspiciousMachine() *int32 {
	return s.SuspiciousMachine
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) GetUnCompleteMachine() *int32 {
	return s.UnCompleteMachine
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetCompleteMachine(v int32) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.CompleteMachine = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetMachineName(v string) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.MachineName = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetProgress(v string) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetRiskLevel(v string) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetSafeMachine(v int32) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.SafeMachine = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetScanMachine(v int32) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.ScanMachine = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetScanPath(v []*string) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.ScanPath = v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetScanTime(v int64) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.ScanTime = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetScanType(v string) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.ScanType = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetStatus(v int32) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetSuspiciousCount(v int32) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.SuspiciousCount = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetSuspiciousMachine(v int32) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.SuspiciousMachine = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetTaskId(v string) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) SetUnCompleteMachine(v int32) *GetVirusScanLatestTaskStatisticResponseBodyData {
	s.UnCompleteMachine = &v
	return s
}

func (s *GetVirusScanLatestTaskStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}

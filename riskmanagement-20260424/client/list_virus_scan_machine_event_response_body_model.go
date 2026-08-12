// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanMachineEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVirusScanMachineEventResponseBody
	GetCode() *string
	SetData(v *ListVirusScanMachineEventResponseBodyData) *ListVirusScanMachineEventResponseBody
	GetData() *ListVirusScanMachineEventResponseBodyData
	SetMessage(v string) *ListVirusScanMachineEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListVirusScanMachineEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListVirusScanMachineEventResponseBody
	GetSuccess() *bool
}

type ListVirusScanMachineEventResponseBody struct {
	// code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ListVirusScanMachineEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message information.
	//
	// example:
	//
	// successful‌
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F0AD8096-E7A2-573D-ACF0-7CE9050CDE38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListVirusScanMachineEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVirusScanMachineEventResponseBody) GetData() *ListVirusScanMachineEventResponseBodyData {
	return s.Data
}

func (s *ListVirusScanMachineEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVirusScanMachineEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirusScanMachineEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListVirusScanMachineEventResponseBody) SetCode(v string) *ListVirusScanMachineEventResponseBody {
	s.Code = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBody) SetData(v *ListVirusScanMachineEventResponseBodyData) *ListVirusScanMachineEventResponseBody {
	s.Data = v
	return s
}

func (s *ListVirusScanMachineEventResponseBody) SetMessage(v string) *ListVirusScanMachineEventResponseBody {
	s.Message = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBody) SetRequestId(v string) *ListVirusScanMachineEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBody) SetSuccess(v bool) *ListVirusScanMachineEventResponseBody {
	s.Success = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVirusScanMachineEventResponseBodyData struct {
	// The request ID.
	//
	// example:
	//
	// 1E222AB5-5C2B-50AD-8A96-E704AF80F2A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the latest virus scan task.
	VirusScanLatestTaskStatistic *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic `json:"VirusScanLatestTaskStatistic,omitempty" xml:"VirusScanLatestTaskStatistic,omitempty" type:"Struct"`
	// The virus alerts detected on specific machines during virus scanning.
	VirusScanMachineEventList *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList `json:"VirusScanMachineEventList,omitempty" xml:"VirusScanMachineEventList,omitempty" type:"Struct"`
}

func (s ListVirusScanMachineEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirusScanMachineEventResponseBodyData) GetVirusScanLatestTaskStatistic() *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	return s.VirusScanLatestTaskStatistic
}

func (s *ListVirusScanMachineEventResponseBodyData) GetVirusScanMachineEventList() *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList {
	return s.VirusScanMachineEventList
}

func (s *ListVirusScanMachineEventResponseBodyData) SetRequestId(v string) *ListVirusScanMachineEventResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyData) SetVirusScanLatestTaskStatistic(v *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) *ListVirusScanMachineEventResponseBodyData {
	s.VirusScanLatestTaskStatistic = v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyData) SetVirusScanMachineEventList(v *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList) *ListVirusScanMachineEventResponseBodyData {
	s.VirusScanMachineEventList = v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyData) Validate() error {
	if s.VirusScanLatestTaskStatistic != nil {
		if err := s.VirusScanLatestTaskStatistic.Validate(); err != nil {
			return err
		}
	}
	if s.VirusScanMachineEventList != nil {
		if err := s.VirusScanMachineEventList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic struct {
	// The number of machines that completed scanning.
	//
	// example:
	//
	// 2
	CompleteMachine *int32 `json:"CompleteMachine,omitempty" xml:"CompleteMachine,omitempty"`
	// The server machine name.
	//
	// example:
	//
	// testMahine1
	MachineName *string `json:"MachineName,omitempty" xml:"MachineName,omitempty"`
	// The percentage of the scan task progress.
	//
	// example:
	//
	// 92
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The highest risk level of the detected alerts. Valid values:
	//
	// example:
	//
	// medium
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The number of machines on which no risks were detected.
	//
	// example:
	//
	// 1
	SafeMachine *int32 `json:"SafeMachine,omitempty" xml:"SafeMachine,omitempty"`
	// The number of machines scanned in this virus scan.
	//
	// example:
	//
	// 1
	ScanMachine *int32 `json:"ScanMachine,omitempty" xml:"ScanMachine,omitempty"`
	// The file paths specified for scanning when the user-defined scan type is used.
	ScanPath []*string `json:"ScanPath,omitempty" xml:"ScanPath,omitempty" type:"Repeated"`
	// The scan timestamp, in milliseconds.
	//
	// example:
	//
	// 1681145862000
	ScanTime *int64 `json:"ScanTime,omitempty" xml:"ScanTime,omitempty"`
	// The scan type of this virus scan. Valid values:
	//
	// example:
	//
	// system
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The status of the scan task.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The number of security alerts detected during the scan.
	//
	// example:
	//
	// 0
	SuspiciousCount *int32 `json:"SuspiciousCount,omitempty" xml:"SuspiciousCount,omitempty"`
	// The number of machines on which risks were detected.
	//
	// example:
	//
	// 1
	SuspiciousMachine *int32 `json:"SuspiciousMachine,omitempty" xml:"SuspiciousMachine,omitempty"`
	// The ID of the scan task.
	//
	// example:
	//
	// t-0mqu9dhpi365dp5iyf
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The number of machines that did not complete scanning or failed during scanning.
	//
	// example:
	//
	// 1
	UnCompleteMachine *int32 `json:"UnCompleteMachine,omitempty" xml:"UnCompleteMachine,omitempty"`
}

func (s ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetCompleteMachine() *int32 {
	return s.CompleteMachine
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetMachineName() *string {
	return s.MachineName
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetProgress() *string {
	return s.Progress
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetSafeMachine() *int32 {
	return s.SafeMachine
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetScanMachine() *int32 {
	return s.ScanMachine
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetScanPath() []*string {
	return s.ScanPath
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetScanTime() *int64 {
	return s.ScanTime
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetScanType() *string {
	return s.ScanType
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetStatus() *int32 {
	return s.Status
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetSuspiciousCount() *int32 {
	return s.SuspiciousCount
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetSuspiciousMachine() *int32 {
	return s.SuspiciousMachine
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetTaskId() *string {
	return s.TaskId
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) GetUnCompleteMachine() *int32 {
	return s.UnCompleteMachine
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetCompleteMachine(v int32) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.CompleteMachine = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetMachineName(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.MachineName = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetProgress(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.Progress = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetRiskLevel(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.RiskLevel = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetSafeMachine(v int32) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.SafeMachine = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetScanMachine(v int32) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.ScanMachine = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetScanPath(v []*string) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.ScanPath = v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetScanTime(v int64) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.ScanTime = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetScanType(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.ScanType = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetStatus(v int32) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.Status = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetSuspiciousCount(v int32) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.SuspiciousCount = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetSuspiciousMachine(v int32) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.SuspiciousMachine = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetTaskId(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.TaskId = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) SetUnCompleteMachine(v int32) *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic {
	s.UnCompleteMachine = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanLatestTaskStatistic) Validate() error {
	return dara.Validate(s)
}

type ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList struct {
	// The details of the alert events.
	Data []*ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
}

func (s ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList) GetData() []*ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData {
	return s.Data
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList) GetPageInfo() *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo {
	return s.PageInfo
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList) SetData(v []*ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList {
	s.Data = v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList) SetPageInfo(v *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList {
	s.PageInfo = v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventList) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData struct {
	// The details of the anomalous event.
	Details []*ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The ID of the alert event.
	//
	// example:
	//
	// 123-2CcoavZnCXrJKqk2KQKxp9WGwup
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name (subtype) of the alert event.
	//
	// example:
	//
	// Malicious script code execution.
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The instance name.
	//
	// example:
	//
	// i-wz92q7m5hsbgfhdss***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 47.57.*1.65
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 47.57.*1.65
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The timestamp of the last occurrence, in milliseconds.
	//
	// example:
	//
	// 1682046733628
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	// The risk level of the alert event. Valid values:
	//
	// example:
	//
	// remind
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) GetDetails() []*ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails {
	return s.Details
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) GetEventId() *int64 {
	return s.EventId
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) GetEventName() *string {
	return s.EventName
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) GetLastTimeStamp() *int64 {
	return s.LastTimeStamp
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) GetLevel() *string {
	return s.Level
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) SetDetails(v []*ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData {
	s.Details = v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) SetEventId(v int64) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData {
	s.EventId = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) SetEventName(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData {
	s.EventName = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) SetInstanceName(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData {
	s.InstanceName = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) SetInternetIp(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData {
	s.InternetIp = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) SetIntranetIp(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData {
	s.IntranetIp = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) SetLastTimeStamp(v int64) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData {
	s.LastTimeStamp = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) SetLevel(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData {
	s.Level = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListData) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails struct {
	// The display type of valueDisplay. Valid values:
	//
	// example:
	//
	// download_url
	InfoType *string `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
	// The display name of the alert event.
	//
	// example:
	//
	// Trojan Path
	NameDisplay *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	// The display method of the anomalous event details.
	//
	// example:
	//
	// text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The additional attribute information of the anomalous event, such as the logon time or logon location for abnormal logon alerts, or the trojan file path or trojan type for trojan alerts.
	//
	// example:
	//
	// getopt
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty"`
}

func (s ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) GetInfoType() *string {
	return s.InfoType
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) GetNameDisplay() *string {
	return s.NameDisplay
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) GetType() *string {
	return s.Type
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) GetValueDisplay() *string {
	return s.ValueDisplay
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) SetInfoType(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails {
	s.InfoType = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) SetNameDisplay(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails {
	s.NameDisplay = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) SetType(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails {
	s.Type = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) SetValueDisplay(v string) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails {
	s.ValueDisplay = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListDataDetails) Validate() error {
	return dara.Validate(s)
}

type ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo struct {
	// The page number of the current page in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries displayed per page in a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of alert events returned.
	//
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo) SetCurrentPage(v int32) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo) SetPageSize(v int32) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo) SetTotalCount(v int32) *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataVirusScanMachineEventListPageInfo) Validate() error {
	return dara.Validate(s)
}

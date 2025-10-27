// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListVirusScanTaskResponseBodyList) *ListVirusScanTaskResponseBody
	GetList() []*ListVirusScanTaskResponseBodyList
	SetPageInfo(v *ListVirusScanTaskResponseBodyPageInfo) *ListVirusScanTaskResponseBody
	GetPageInfo() *ListVirusScanTaskResponseBodyPageInfo
	SetRequestId(v string) *ListVirusScanTaskResponseBody
	GetRequestId() *string
}

type ListVirusScanTaskResponseBody struct {
	// The returned virus scan tasks.
	List []*ListVirusScanTaskResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListVirusScanTaskResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-D1BBB3936FA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVirusScanTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskResponseBody) GetList() []*ListVirusScanTaskResponseBodyList {
	return s.List
}

func (s *ListVirusScanTaskResponseBody) GetPageInfo() *ListVirusScanTaskResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListVirusScanTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirusScanTaskResponseBody) SetList(v []*ListVirusScanTaskResponseBodyList) *ListVirusScanTaskResponseBody {
	s.List = v
	return s
}

func (s *ListVirusScanTaskResponseBody) SetPageInfo(v *ListVirusScanTaskResponseBodyPageInfo) *ListVirusScanTaskResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListVirusScanTaskResponseBody) SetRequestId(v string) *ListVirusScanTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirusScanTaskResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
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

type ListVirusScanTaskResponseBodyList struct {
	// The timestamp when the virus scan task ended. Unit: milliseconds.
	//
	// example:
	//
	// 1662343860051
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// sql-test-0****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 8.210.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.25.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The progress of the task in percentage.
	//
	// example:
	//
	// 62
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The information about the file that is scanned.
	ScanPath []*string `json:"ScanPath,omitempty" xml:"ScanPath,omitempty" type:"Repeated"`
	// The type of the virus scan task. Valid values:
	//
	// 	- **system**: automatic scan task
	//
	// 	- **user**: custom scan task
	//
	// example:
	//
	// system
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// The timestamp when the virus scan task started. Unit: milliseconds.
	//
	// example:
	//
	// 1651290987000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the virus scan task. Valid values:
	//
	// 	- **1**: running
	//
	// 	- **2**: complete
	//
	// 	- **3**: failed
	//
	// 	- **4**: timed out
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the virus scan task.
	//
	// example:
	//
	// 2e75557bfa570381f5c516cad9b6xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The name of the virus scan task.
	//
	// 	- The value is fixed as **VIRUS_VUL_SCHEDULE_SCAN**, which indicates a virus scan task.
	//
	// example:
	//
	// VIRUS_VUL_SCHEDULE_SCAN
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListVirusScanTaskResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskResponseBodyList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListVirusScanTaskResponseBodyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListVirusScanTaskResponseBodyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListVirusScanTaskResponseBodyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListVirusScanTaskResponseBodyList) GetProgress() *int32 {
	return s.Progress
}

func (s *ListVirusScanTaskResponseBodyList) GetScanPath() []*string {
	return s.ScanPath
}

func (s *ListVirusScanTaskResponseBodyList) GetScanType() *string {
	return s.ScanType
}

func (s *ListVirusScanTaskResponseBodyList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListVirusScanTaskResponseBodyList) GetStatus() *int32 {
	return s.Status
}

func (s *ListVirusScanTaskResponseBodyList) GetTaskId() *string {
	return s.TaskId
}

func (s *ListVirusScanTaskResponseBodyList) GetTaskName() *string {
	return s.TaskName
}

func (s *ListVirusScanTaskResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *ListVirusScanTaskResponseBodyList) SetEndTime(v int64) *ListVirusScanTaskResponseBodyList {
	s.EndTime = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) SetInstanceName(v string) *ListVirusScanTaskResponseBodyList {
	s.InstanceName = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) SetInternetIp(v string) *ListVirusScanTaskResponseBodyList {
	s.InternetIp = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) SetIntranetIp(v string) *ListVirusScanTaskResponseBodyList {
	s.IntranetIp = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) SetProgress(v int32) *ListVirusScanTaskResponseBodyList {
	s.Progress = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) SetScanPath(v []*string) *ListVirusScanTaskResponseBodyList {
	s.ScanPath = v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) SetScanType(v string) *ListVirusScanTaskResponseBodyList {
	s.ScanType = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) SetStartTime(v int64) *ListVirusScanTaskResponseBodyList {
	s.StartTime = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) SetStatus(v int32) *ListVirusScanTaskResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) SetTaskId(v string) *ListVirusScanTaskResponseBodyList {
	s.TaskId = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) SetTaskName(v string) *ListVirusScanTaskResponseBodyList {
	s.TaskName = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) SetUuid(v string) *ListVirusScanTaskResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListVirusScanTaskResponseBodyPageInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 224
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVirusScanTaskResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanTaskResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListVirusScanTaskResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListVirusScanTaskResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVirusScanTaskResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVirusScanTaskResponseBodyPageInfo) SetCurrentPage(v int32) *ListVirusScanTaskResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyPageInfo) SetPageSize(v int32) *ListVirusScanTaskResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyPageInfo) SetTotalCount(v int32) *ListVirusScanTaskResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListVirusScanTaskResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

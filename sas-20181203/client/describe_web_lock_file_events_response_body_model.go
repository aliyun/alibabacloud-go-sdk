// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebLockFileEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeWebLockFileEventsResponseBody
	GetCurrentPage() *int32
	SetList(v []*DescribeWebLockFileEventsResponseBodyList) *DescribeWebLockFileEventsResponseBody
	GetList() []*DescribeWebLockFileEventsResponseBodyList
	SetPageSize(v int32) *DescribeWebLockFileEventsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeWebLockFileEventsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeWebLockFileEventsResponseBody
	GetTotalCount() *int32
}

type DescribeWebLockFileEventsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 2
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// An array that consists of events on web tamper proofing returned.
	List []*DescribeWebLockFileEventsResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 79CFF74D-E967-5407-8A78-EE03B925FDAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of events on web tamper proofing returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeWebLockFileEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockFileEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebLockFileEventsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeWebLockFileEventsResponseBody) GetList() []*DescribeWebLockFileEventsResponseBodyList {
	return s.List
}

func (s *DescribeWebLockFileEventsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeWebLockFileEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebLockFileEventsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeWebLockFileEventsResponseBody) SetCurrentPage(v int32) *DescribeWebLockFileEventsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBody) SetList(v []*DescribeWebLockFileEventsResponseBodyList) *DescribeWebLockFileEventsResponseBody {
	s.List = v
	return s
}

func (s *DescribeWebLockFileEventsResponseBody) SetPageSize(v int32) *DescribeWebLockFileEventsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBody) SetRequestId(v string) *DescribeWebLockFileEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBody) SetTotalCount(v int32) *DescribeWebLockFileEventsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebLockFileEventsResponseBodyList struct {
	// The number of attempts.
	//
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The timestamp at which the event on web tamper proofing was first detected.
	//
	// example:
	//
	// 1657178400000
	Ds *int64 `json:"Ds,omitempty" xml:"Ds,omitempty"`
	// The name of the event on web tamper proofing.
	//
	// example:
	//
	// modify
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The status of the event on web tamper proofing. Valid values:
	//
	// 	- **1**: unhandled
	//
	// 	- **2**: ignored
	//
	// 	- **4**: deprecated
	//
	// 	- **8**: marked as false positive
	//
	// 	- **10**: added to the whitelist
	//
	// 	- **16**: handling
	//
	// 	- **32**: defended
	//
	// 	- **64**: invalid
	//
	// 	- **128**: deleted
	//
	// 	- **512**: automatically handled
	//
	// example:
	//
	// 1
	EventStatus *string `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The prevention mode. Valid values:
	//
	// 	- **audit**: Interception Mode
	//
	// 	- **web_lock**: Alert Mode
	//
	// example:
	//
	// audit
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The timestamp at which the event on web tamper proofing was last detected.
	//
	// example:
	//
	// 1657178400000
	GmtEvent *int64 `json:"GmtEvent,omitempty" xml:"GmtEvent,omitempty"`
	// The ID of the event on web tamper proofing.
	//
	// example:
	//
	// 3555953980
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// sql-test-001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the affected asset.
	//
	// example:
	//
	// 8.210.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the asset.
	//
	// example:
	//
	// 172.25.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 8.210.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The severity of the event on web tamper proofing. Valid values: **medium**
	//
	// example:
	//
	// medium
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The file path.
	//
	// example:
	//
	// D:\\test-tamper-proofing\\123.html
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// python3.7
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The path to the process.
	//
	// example:
	//
	// C:\\Windows\\explorer.exe
	ProcessPath *string `json:"ProcessPath,omitempty" xml:"ProcessPath,omitempty"`
	// The status of the event on web tamper proofing. Valid values:
	//
	// 	- **1**: unhandled
	//
	// 	- **2**: ignored
	//
	// 	- **4**: deprecated
	//
	// 	- **8**: marked as false positive
	//
	// 	- **10**: added to the whitelist
	//
	// 	- **16**: handling
	//
	// 	- **32**: defended
	//
	// 	- **64**: invalid
	//
	// 	- **128**: deleted
	//
	// 	- **512**: automatically handled
	//
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeWebLockFileEventsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebLockFileEventsResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetCount() *int64 {
	return s.Count
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetDs() *int64 {
	return s.Ds
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetEventName() *string {
	return s.EventName
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetEventStatus() *string {
	return s.EventStatus
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetEventType() *string {
	return s.EventType
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetGmtEvent() *int64 {
	return s.GmtEvent
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetId() *int64 {
	return s.Id
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetIp() *string {
	return s.Ip
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetLevel() *string {
	return s.Level
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetPath() *string {
	return s.Path
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetProcessPath() *string {
	return s.ProcessPath
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetStatus() *string {
	return s.Status
}

func (s *DescribeWebLockFileEventsResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetCount(v int64) *DescribeWebLockFileEventsResponseBodyList {
	s.Count = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetDs(v int64) *DescribeWebLockFileEventsResponseBodyList {
	s.Ds = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetEventName(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.EventName = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetEventStatus(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.EventStatus = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetEventType(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.EventType = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetGmtEvent(v int64) *DescribeWebLockFileEventsResponseBodyList {
	s.GmtEvent = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetId(v int64) *DescribeWebLockFileEventsResponseBodyList {
	s.Id = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetInstanceName(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.InstanceName = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetInternetIp(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.InternetIp = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetIntranetIp(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetIp(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.Ip = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetLevel(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.Level = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetPath(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.Path = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetProcessName(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.ProcessName = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetProcessPath(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.ProcessPath = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetStatus(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.Status = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) SetUuid(v string) *DescribeWebLockFileEventsResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *DescribeWebLockFileEventsResponseBodyList) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventList(v []*ListFileProtectEventResponseBodyEventList) *ListFileProtectEventResponseBody
	GetEventList() []*ListFileProtectEventResponseBodyEventList
	SetPageInfo(v *ListFileProtectEventResponseBodyPageInfo) *ListFileProtectEventResponseBody
	GetPageInfo() *ListFileProtectEventResponseBodyPageInfo
	SetRequestId(v string) *ListFileProtectEventResponseBody
	GetRequestId() *string
}

type ListFileProtectEventResponseBody struct {
	// The events.
	EventList []*ListFileProtectEventResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListFileProtectEventResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 23AD0BD2-8771-5647-819E-6BA51E212F80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFileProtectEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileProtectEventResponseBody) GetEventList() []*ListFileProtectEventResponseBodyEventList {
	return s.EventList
}

func (s *ListFileProtectEventResponseBody) GetPageInfo() *ListFileProtectEventResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListFileProtectEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileProtectEventResponseBody) SetEventList(v []*ListFileProtectEventResponseBodyEventList) *ListFileProtectEventResponseBody {
	s.EventList = v
	return s
}

func (s *ListFileProtectEventResponseBody) SetPageInfo(v *ListFileProtectEventResponseBodyPageInfo) *ListFileProtectEventResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListFileProtectEventResponseBody) SetRequestId(v string) *ListFileProtectEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileProtectEventResponseBody) Validate() error {
	if s.EventList != nil {
		for _, item := range s.EventList {
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

type ListFileProtectEventResponseBodyEventList struct {
	// The severity of the alert. Valid values:
	//
	// 	- 0: no alerts
	//
	// 	- 1: reminder
	//
	// 	- 2: suspicious
	//
	// 	- 3: high-risk
	//
	// example:
	//
	// 1
	AlertLevel *int32 `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// The command line of the event.
	//
	// example:
	//
	// ["touch","/test/aaaa"]
	CmdLine *string `json:"CmdLine,omitempty" xml:"CmdLine,omitempty"`
	// The path to the file that is managed by the process.
	//
	// example:
	//
	// /etc/pam.d/su
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The time when the event was handled.
	//
	// example:
	//
	// 1694576692000
	HandleTime *int64 `json:"HandleTime,omitempty" xml:"HandleTime,omitempty"`
	// The ID of the event.
	//
	// example:
	//
	// 161757
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance name of the server.
	//
	// example:
	//
	// kyy-admin-01
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 121.40.211.194
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.22.5.36
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The time when the event last occurred.
	//
	// example:
	//
	// 1694576692000
	LatestTime *int64 `json:"LatestTime,omitempty" xml:"LatestTime,omitempty"`
	// The operation performed by the process on the file.
	//
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The type of the operating system. Valid values:
	//
	// 	- **windows**: Windows
	//
	// 	- **linux**: Linux
	//
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The path to the process.
	//
	// example:
	//
	// /usr/bin/sshpass
	ProcPath *string `json:"ProcPath,omitempty" xml:"ProcPath,omitempty"`
	// The permissions required to start the process.
	//
	// example:
	//
	// rwxr-xr-x
	ProcPermission *string `json:"ProcPermission,omitempty" xml:"ProcPermission,omitempty"`
	// The ID of the process.
	//
	// example:
	//
	// 52636
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The description.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test-rule-1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- 0: unhandled
	//
	// 	- 1: handled
	//
	// 	- 2: added to the whitelist
	//
	// 	- 3: ignored
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server that is associated with the process.
	//
	// example:
	//
	// 94b44720-d982-4d20-a4e1-80a1a57b5a8e
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListFileProtectEventResponseBodyEventList) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectEventResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *ListFileProtectEventResponseBodyEventList) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *ListFileProtectEventResponseBodyEventList) GetCmdLine() *string {
	return s.CmdLine
}

func (s *ListFileProtectEventResponseBodyEventList) GetFilePath() *string {
	return s.FilePath
}

func (s *ListFileProtectEventResponseBodyEventList) GetHandleTime() *int64 {
	return s.HandleTime
}

func (s *ListFileProtectEventResponseBodyEventList) GetId() *int64 {
	return s.Id
}

func (s *ListFileProtectEventResponseBodyEventList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListFileProtectEventResponseBodyEventList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListFileProtectEventResponseBodyEventList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListFileProtectEventResponseBodyEventList) GetLatestTime() *int64 {
	return s.LatestTime
}

func (s *ListFileProtectEventResponseBodyEventList) GetOperation() *string {
	return s.Operation
}

func (s *ListFileProtectEventResponseBodyEventList) GetPlatform() *string {
	return s.Platform
}

func (s *ListFileProtectEventResponseBodyEventList) GetProcPath() *string {
	return s.ProcPath
}

func (s *ListFileProtectEventResponseBodyEventList) GetProcPermission() *string {
	return s.ProcPermission
}

func (s *ListFileProtectEventResponseBodyEventList) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListFileProtectEventResponseBodyEventList) GetRemark() *string {
	return s.Remark
}

func (s *ListFileProtectEventResponseBodyEventList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListFileProtectEventResponseBodyEventList) GetStatus() *int32 {
	return s.Status
}

func (s *ListFileProtectEventResponseBodyEventList) GetUuid() *string {
	return s.Uuid
}

func (s *ListFileProtectEventResponseBodyEventList) SetAlertLevel(v int32) *ListFileProtectEventResponseBodyEventList {
	s.AlertLevel = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetCmdLine(v string) *ListFileProtectEventResponseBodyEventList {
	s.CmdLine = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetFilePath(v string) *ListFileProtectEventResponseBodyEventList {
	s.FilePath = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetHandleTime(v int64) *ListFileProtectEventResponseBodyEventList {
	s.HandleTime = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetId(v int64) *ListFileProtectEventResponseBodyEventList {
	s.Id = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetInstanceName(v string) *ListFileProtectEventResponseBodyEventList {
	s.InstanceName = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetInternetIp(v string) *ListFileProtectEventResponseBodyEventList {
	s.InternetIp = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetIntranetIp(v string) *ListFileProtectEventResponseBodyEventList {
	s.IntranetIp = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetLatestTime(v int64) *ListFileProtectEventResponseBodyEventList {
	s.LatestTime = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetOperation(v string) *ListFileProtectEventResponseBodyEventList {
	s.Operation = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetPlatform(v string) *ListFileProtectEventResponseBodyEventList {
	s.Platform = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetProcPath(v string) *ListFileProtectEventResponseBodyEventList {
	s.ProcPath = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetProcPermission(v string) *ListFileProtectEventResponseBodyEventList {
	s.ProcPermission = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetProcessId(v string) *ListFileProtectEventResponseBodyEventList {
	s.ProcessId = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetRemark(v string) *ListFileProtectEventResponseBodyEventList {
	s.Remark = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetRuleName(v string) *ListFileProtectEventResponseBodyEventList {
	s.RuleName = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetStatus(v int32) *ListFileProtectEventResponseBodyEventList {
	s.Status = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) SetUuid(v string) *ListFileProtectEventResponseBodyEventList {
	s.Uuid = &v
	return s
}

func (s *ListFileProtectEventResponseBodyEventList) Validate() error {
	return dara.Validate(s)
}

type ListFileProtectEventResponseBodyPageInfo struct {
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
	// 149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileProtectEventResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectEventResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListFileProtectEventResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFileProtectEventResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileProtectEventResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFileProtectEventResponseBodyPageInfo) SetCurrentPage(v int32) *ListFileProtectEventResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListFileProtectEventResponseBodyPageInfo) SetPageSize(v int32) *ListFileProtectEventResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListFileProtectEventResponseBodyPageInfo) SetTotalCount(v int32) *ListFileProtectEventResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListFileProtectEventResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectClientEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventList(v []*ListFileProtectClientEventResponseBodyEventList) *ListFileProtectClientEventResponseBody
	GetEventList() []*ListFileProtectClientEventResponseBodyEventList
	SetPageInfo(v *ListFileProtectClientEventResponseBodyPageInfo) *ListFileProtectClientEventResponseBody
	GetPageInfo() *ListFileProtectClientEventResponseBodyPageInfo
	SetRequestId(v string) *ListFileProtectClientEventResponseBody
	GetRequestId() *string
}

type ListFileProtectClientEventResponseBody struct {
	EventList []*ListFileProtectClientEventResponseBodyEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	PageInfo  *ListFileProtectClientEventResponseBodyPageInfo    `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// ACF97412-FD09-4D1F-994F-34DF12B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFileProtectClientEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientEventResponseBody) GetEventList() []*ListFileProtectClientEventResponseBodyEventList {
	return s.EventList
}

func (s *ListFileProtectClientEventResponseBody) GetPageInfo() *ListFileProtectClientEventResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListFileProtectClientEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFileProtectClientEventResponseBody) SetEventList(v []*ListFileProtectClientEventResponseBodyEventList) *ListFileProtectClientEventResponseBody {
	s.EventList = v
	return s
}

func (s *ListFileProtectClientEventResponseBody) SetPageInfo(v *ListFileProtectClientEventResponseBodyPageInfo) *ListFileProtectClientEventResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListFileProtectClientEventResponseBody) SetRequestId(v string) *ListFileProtectClientEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileProtectClientEventResponseBody) Validate() error {
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

type ListFileProtectClientEventResponseBodyEventList struct {
	// example:
	//
	// 1
	AlertLevel *int32 `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// example:
	//
	// ["touch","/test/aaaa"]
	CmdLine *string `json:"CmdLine,omitempty" xml:"CmdLine,omitempty"`
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// /data/*
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// 1694576692000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// example:
	//
	// 1694576692000
	HandleTime *int64 `json:"HandleTime,omitempty" xml:"HandleTime,omitempty"`
	// example:
	//
	// 3719
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// kyy-admin-01
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 121.40.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// 172.22.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// example:
	//
	// 1694576692000
	LatestTime *int64 `json:"LatestTime,omitempty" xml:"LatestTime,omitempty"`
	// example:
	//
	// DELETE
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// c:/*Unity*
	ProcPath *string `json:"ProcPath,omitempty" xml:"ProcPath,omitempty"`
	// example:
	//
	// rwxr-xr-x
	ProcPermission *string `json:"ProcPermission,omitempty" xml:"ProcPermission,omitempty"`
	// example:
	//
	// 52636
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// block
	RuleAction *string `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// example:
	//
	// test-rule-1
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 94b44720-d982-4d20-a4e1-80a1a57b****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListFileProtectClientEventResponseBodyEventList) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientEventResponseBodyEventList) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetCmdLine() *string {
	return s.CmdLine
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetCount() *int32 {
	return s.Count
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetFilePath() *string {
	return s.FilePath
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetHandleTime() *int64 {
	return s.HandleTime
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetId() *int64 {
	return s.Id
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetLatestTime() *int64 {
	return s.LatestTime
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetOperation() *string {
	return s.Operation
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetPlatform() *string {
	return s.Platform
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetProcPath() *string {
	return s.ProcPath
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetProcPermission() *string {
	return s.ProcPermission
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetProcessId() *string {
	return s.ProcessId
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetRemark() *string {
	return s.Remark
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetRuleAction() *string {
	return s.RuleAction
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetRuleName() *string {
	return s.RuleName
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetStatus() *int32 {
	return s.Status
}

func (s *ListFileProtectClientEventResponseBodyEventList) GetUuid() *string {
	return s.Uuid
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetAlertLevel(v int32) *ListFileProtectClientEventResponseBodyEventList {
	s.AlertLevel = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetCmdLine(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.CmdLine = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetCount(v int32) *ListFileProtectClientEventResponseBodyEventList {
	s.Count = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetFilePath(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.FilePath = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetFirstTime(v int64) *ListFileProtectClientEventResponseBodyEventList {
	s.FirstTime = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetHandleTime(v int64) *ListFileProtectClientEventResponseBodyEventList {
	s.HandleTime = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetId(v int64) *ListFileProtectClientEventResponseBodyEventList {
	s.Id = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetInstanceName(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.InstanceName = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetInternetIp(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.InternetIp = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetIntranetIp(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.IntranetIp = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetLatestTime(v int64) *ListFileProtectClientEventResponseBodyEventList {
	s.LatestTime = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetOperation(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.Operation = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetPlatform(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.Platform = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetProcPath(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.ProcPath = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetProcPermission(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.ProcPermission = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetProcessId(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.ProcessId = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetRemark(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.Remark = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetRuleAction(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.RuleAction = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetRuleName(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.RuleName = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetStatus(v int32) *ListFileProtectClientEventResponseBodyEventList {
	s.Status = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) SetUuid(v string) *ListFileProtectClientEventResponseBodyEventList {
	s.Uuid = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyEventList) Validate() error {
	return dara.Validate(s)
}

type ListFileProtectClientEventResponseBodyPageInfo struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 263
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFileProtectClientEventResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectClientEventResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListFileProtectClientEventResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFileProtectClientEventResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFileProtectClientEventResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFileProtectClientEventResponseBodyPageInfo) SetCurrentPage(v int32) *ListFileProtectClientEventResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyPageInfo) SetPageSize(v int32) *ListFileProtectClientEventResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyPageInfo) SetTotalCount(v int32) *ListFileProtectClientEventResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListFileProtectClientEventResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

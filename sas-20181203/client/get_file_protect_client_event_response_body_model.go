// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectClientEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileProtectClientEventResponseBodyData) *GetFileProtectClientEventResponseBody
	GetData() *GetFileProtectClientEventResponseBodyData
	SetRequestId(v string) *GetFileProtectClientEventResponseBody
	GetRequestId() *string
}

type GetFileProtectClientEventResponseBody struct {
	Data *GetFileProtectClientEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1383B0DB-D5D6-4B0C-9E6B-75939C8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileProtectClientEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientEventResponseBody) GetData() *GetFileProtectClientEventResponseBodyData {
	return s.Data
}

func (s *GetFileProtectClientEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileProtectClientEventResponseBody) SetData(v *GetFileProtectClientEventResponseBodyData) *GetFileProtectClientEventResponseBody {
	s.Data = v
	return s
}

func (s *GetFileProtectClientEventResponseBody) SetRequestId(v string) *GetFileProtectClientEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileProtectClientEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileProtectClientEventResponseBodyData struct {
	// example:
	//
	// 0
	AlertLevel *int32 `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// example:
	//
	// ["touch","/usr/local/aaaa"]
	CmdLine *string `json:"CmdLine,omitempty" xml:"CmdLine,omitempty"`
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// /usr/local
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// example:
	//
	// rwxr-xr-x
	FilePermission *string `json:"FilePermission,omitempty" xml:"FilePermission,omitempty"`
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
	// 3454
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// i-wz92q7m5hsbgfhdss***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 17.16.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// 10.42.XX.XX
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
	// /bin/bash33
	ProcPath *string `json:"ProcPath,omitempty" xml:"ProcPath,omitempty"`
	// example:
	//
	// 3453
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
	// test-000
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 12
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// root
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// 6690a46c-0edb-4663-a641-3629d1a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetFileProtectClientEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectClientEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileProtectClientEventResponseBodyData) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *GetFileProtectClientEventResponseBodyData) GetCmdLine() *string {
	return s.CmdLine
}

func (s *GetFileProtectClientEventResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *GetFileProtectClientEventResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *GetFileProtectClientEventResponseBodyData) GetFilePermission() *string {
	return s.FilePermission
}

func (s *GetFileProtectClientEventResponseBodyData) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *GetFileProtectClientEventResponseBodyData) GetHandleTime() *int64 {
	return s.HandleTime
}

func (s *GetFileProtectClientEventResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetFileProtectClientEventResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetFileProtectClientEventResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *GetFileProtectClientEventResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *GetFileProtectClientEventResponseBodyData) GetLatestTime() *int64 {
	return s.LatestTime
}

func (s *GetFileProtectClientEventResponseBodyData) GetOperation() *string {
	return s.Operation
}

func (s *GetFileProtectClientEventResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *GetFileProtectClientEventResponseBodyData) GetProcPath() *string {
	return s.ProcPath
}

func (s *GetFileProtectClientEventResponseBodyData) GetProcessId() *string {
	return s.ProcessId
}

func (s *GetFileProtectClientEventResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *GetFileProtectClientEventResponseBodyData) GetRuleAction() *string {
	return s.RuleAction
}

func (s *GetFileProtectClientEventResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *GetFileProtectClientEventResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetFileProtectClientEventResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetFileProtectClientEventResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *GetFileProtectClientEventResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *GetFileProtectClientEventResponseBodyData) SetAlertLevel(v int32) *GetFileProtectClientEventResponseBodyData {
	s.AlertLevel = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetCmdLine(v string) *GetFileProtectClientEventResponseBodyData {
	s.CmdLine = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetCount(v int32) *GetFileProtectClientEventResponseBodyData {
	s.Count = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetFilePath(v string) *GetFileProtectClientEventResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetFilePermission(v string) *GetFileProtectClientEventResponseBodyData {
	s.FilePermission = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetFirstTime(v int64) *GetFileProtectClientEventResponseBodyData {
	s.FirstTime = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetHandleTime(v int64) *GetFileProtectClientEventResponseBodyData {
	s.HandleTime = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetId(v int64) *GetFileProtectClientEventResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetInstanceName(v string) *GetFileProtectClientEventResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetInternetIp(v string) *GetFileProtectClientEventResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetIntranetIp(v string) *GetFileProtectClientEventResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetLatestTime(v int64) *GetFileProtectClientEventResponseBodyData {
	s.LatestTime = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetOperation(v string) *GetFileProtectClientEventResponseBodyData {
	s.Operation = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetPlatform(v string) *GetFileProtectClientEventResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetProcPath(v string) *GetFileProtectClientEventResponseBodyData {
	s.ProcPath = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetProcessId(v string) *GetFileProtectClientEventResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetRemark(v string) *GetFileProtectClientEventResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetRuleAction(v string) *GetFileProtectClientEventResponseBodyData {
	s.RuleAction = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetRuleName(v string) *GetFileProtectClientEventResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetStatus(v int32) *GetFileProtectClientEventResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetUserId(v string) *GetFileProtectClientEventResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetUserName(v string) *GetFileProtectClientEventResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) SetUuid(v string) *GetFileProtectClientEventResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetFileProtectClientEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}

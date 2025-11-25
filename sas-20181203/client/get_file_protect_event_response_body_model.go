// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileProtectEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetFileProtectEventResponseBodyData) *GetFileProtectEventResponseBody
	GetData() *GetFileProtectEventResponseBodyData
	SetRequestId(v string) *GetFileProtectEventResponseBody
	GetRequestId() *string
}

type GetFileProtectEventResponseBody struct {
	// The details of the returned data.
	Data *GetFileProtectEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF46038
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFileProtectEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileProtectEventResponseBody) GetData() *GetFileProtectEventResponseBodyData {
	return s.Data
}

func (s *GetFileProtectEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileProtectEventResponseBody) SetData(v *GetFileProtectEventResponseBodyData) *GetFileProtectEventResponseBody {
	s.Data = v
	return s
}

func (s *GetFileProtectEventResponseBody) SetRequestId(v string) *GetFileProtectEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileProtectEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetFileProtectEventResponseBodyData struct {
	// The severity of alerts. Valid values:
	//
	// 	- 0: does not generate alerts
	//
	// 	- 1: sends notifications
	//
	// 	- 2: suspicious
	//
	// 	- 3: high-risk
	//
	// example:
	//
	// 1
	AlertLevel *int32 `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// The event command line.
	//
	// example:
	//
	// ["touch","/usr/local/aaaa"]
	CmdLine *string `json:"CmdLine,omitempty" xml:"CmdLine,omitempty"`
	// The path to the file.
	//
	// example:
	//
	// /usr/local
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The permissions to run the process.
	//
	// example:
	//
	// rwxr-xr-x
	FilePermission *string `json:"FilePermission,omitempty" xml:"FilePermission,omitempty"`
	// The timestamp at which the event was first detected.
	//
	// example:
	//
	// 1694576692000
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
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
	// 55037
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// i-wz92q7m5hsbgfhdss***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the associated instance.
	//
	// example:
	//
	// 172.16.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the associated instance.
	//
	// example:
	//
	// 10.42.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The time when the event last occurred.
	//
	// example:
	//
	// 1694576692000
	LatestTime *int64 `json:"LatestTime,omitempty" xml:"LatestTime,omitempty"`
	// The operation that the process performed on the file.
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
	// /bin/bash33
	ProcPath *string `json:"ProcPath,omitempty" xml:"ProcPath,omitempty"`
	// The process ID of the event.
	//
	// example:
	//
	// 3453
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The remarks.
	//
	// example:
	//
	// ["test"]
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test-000
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the event. Valid values:
	//
	// 	- 0: not handled
	//
	// 	- 1: handled
	//
	// 	- 2: added to the whitelist
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// User ID of the user who started the current process.
	//
	// example:
	//
	// 1001
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username that started the current process.
	//
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// 49f1360f-62c8-4b48-a24c-5cc317656419
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetFileProtectEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFileProtectEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFileProtectEventResponseBodyData) GetAlertLevel() *int32 {
	return s.AlertLevel
}

func (s *GetFileProtectEventResponseBodyData) GetCmdLine() *string {
	return s.CmdLine
}

func (s *GetFileProtectEventResponseBodyData) GetFilePath() *string {
	return s.FilePath
}

func (s *GetFileProtectEventResponseBodyData) GetFilePermission() *string {
	return s.FilePermission
}

func (s *GetFileProtectEventResponseBodyData) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *GetFileProtectEventResponseBodyData) GetHandleTime() *int64 {
	return s.HandleTime
}

func (s *GetFileProtectEventResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetFileProtectEventResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetFileProtectEventResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *GetFileProtectEventResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *GetFileProtectEventResponseBodyData) GetLatestTime() *int64 {
	return s.LatestTime
}

func (s *GetFileProtectEventResponseBodyData) GetOperation() *string {
	return s.Operation
}

func (s *GetFileProtectEventResponseBodyData) GetPlatform() *string {
	return s.Platform
}

func (s *GetFileProtectEventResponseBodyData) GetProcPath() *string {
	return s.ProcPath
}

func (s *GetFileProtectEventResponseBodyData) GetProcessId() *string {
	return s.ProcessId
}

func (s *GetFileProtectEventResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *GetFileProtectEventResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *GetFileProtectEventResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetFileProtectEventResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetFileProtectEventResponseBodyData) GetUserName() *string {
	return s.UserName
}

func (s *GetFileProtectEventResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *GetFileProtectEventResponseBodyData) SetAlertLevel(v int32) *GetFileProtectEventResponseBodyData {
	s.AlertLevel = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetCmdLine(v string) *GetFileProtectEventResponseBodyData {
	s.CmdLine = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetFilePath(v string) *GetFileProtectEventResponseBodyData {
	s.FilePath = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetFilePermission(v string) *GetFileProtectEventResponseBodyData {
	s.FilePermission = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetFirstTime(v int64) *GetFileProtectEventResponseBodyData {
	s.FirstTime = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetHandleTime(v int64) *GetFileProtectEventResponseBodyData {
	s.HandleTime = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetId(v int64) *GetFileProtectEventResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetInstanceName(v string) *GetFileProtectEventResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetInternetIp(v string) *GetFileProtectEventResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetIntranetIp(v string) *GetFileProtectEventResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetLatestTime(v int64) *GetFileProtectEventResponseBodyData {
	s.LatestTime = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetOperation(v string) *GetFileProtectEventResponseBodyData {
	s.Operation = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetPlatform(v string) *GetFileProtectEventResponseBodyData {
	s.Platform = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetProcPath(v string) *GetFileProtectEventResponseBodyData {
	s.ProcPath = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetProcessId(v string) *GetFileProtectEventResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetRemark(v string) *GetFileProtectEventResponseBodyData {
	s.Remark = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetRuleName(v string) *GetFileProtectEventResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetStatus(v int32) *GetFileProtectEventResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetUserId(v string) *GetFileProtectEventResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetUserName(v string) *GetFileProtectEventResponseBodyData {
	s.UserName = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) SetUuid(v string) *GetFileProtectEventResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetFileProtectEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}
